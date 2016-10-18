package main

import (
	"encoding/json"
	"fmt"
	getopt "github.com/kesselborn/go-getopt"
	goini "github.com/zieckey/goini"
	cloudstack "golang-cloudstack-library"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
        "time"
)

func main() {
	sco := getopt.SubCommandOptions{
		getopt.Options{
			"manipulate templates in cloudstack",
			getopt.Definitions{
				{"ini|i", "specified ini file", getopt.Optional | getopt.ExampleIsDefault, "/etc/cstemplate.ini"},
				{"command", "command to execute", getopt.IsSubCommand, ""},
			},
		},
		getopt.SubCommands{
			"register": {
				"register template",
				getopt.Definitions{
					{"name", "template name", getopt.IsArg | getopt.Optional, ""},
					{"url", "template url", getopt.IsArg | getopt.Optional, ""},
					{"ostype|o", "template ostype", getopt.Optional | getopt.ExampleIsDefault, "centos%6.5%64"},
					{"passwordenabled|p", "template passwordenabled", getopt.Optional | getopt.ExampleIsDefault, "true"},
				},
			},
			"delete": {
				"delete template",
				getopt.Definitions{
					{"id", "template id", getopt.IsArg | getopt.Optional, ""},
				},
			},
			"list": {
				"list templates",
				getopt.Definitions{
					{"keyword|k", "keyword to search", getopt.Optional | getopt.ExampleIsDefault, "all"},
					{"columns|c", "columns to show", getopt.Optional | getopt.ExampleIsDefault, "id,name"},
				},
			},
		},
	}

	scope, options, arguments, _, e := sco.ParseCommandLine()

	help, wantsHelp := options["help"]

	if e != nil || wantsHelp {
		exit_code := 0

		switch {
		case wantsHelp && help.String == "usage":
			fmt.Print(sco.Usage())
		case wantsHelp && help.String == "help":
			fmt.Print(sco.Help())
		default:
			fmt.Println("**** Error: ", e.Error(), "\n", sco.Help())
			exit_code = e.ErrorCode
		}
		os.Exit(exit_code)
	}

	log.SetOutput(ioutil.Discard)
	ini := goini.New()
	err := ini.ParseFile(options["ini"].String)
	if err != nil {
		fmt.Printf("parse INI file %v failed : %v\n", os.Args[1], err.Error())
		return
	}

	client := callclient(ini)
	if scope == "register" {
		callregister(ini, client, arguments, options)
	} else if scope == "delete" {
		calldelete(ini, client, arguments)
	} else if scope == "list" {
		calllist(ini, client, arguments, options)
	}

}

func callclient(ini *goini.INI) (client *cloudstack.Client) {
	endpoints, _ := ini.SectionGet("main", "endpoint")
	apikey, _ := ini.SectionGet("main", "apikey")
	secretkey, _ := ini.SectionGet("main", "secretkey")
	username, _ := ini.SectionGet("main", "username")
	password, _ := ini.SectionGet("main", "password")
	endpoint, _ := url.Parse(endpoints)
	client, _ = cloudstack.NewClient(endpoint, apikey, secretkey, username, password)
	return client
}

func callregister(ini *goini.INI, client *cloudstack.Client, arguments []string, options map[string]getopt.OptionValue) {
	zonename, _ := ini.SectionGet("main", "zonename")
	format, _ := ini.SectionGet("main", "format")
	hypervisor, _ := ini.SectionGet("main", "hypervisor")
	ostype, _ := ini.SectionGet("main", "ostype")
	passwordenabled, _ := ini.SectionGetBool("main", "passwordenabled")
	registermap := make(map[string]string)
	if options["ostype"].Set {
		ostype = options["ostype"].String
	}
	if options["passwordenabled"].Set && options["passwordenabled"].String == "false" {
		passwordenabled = false
	}
	if len(arguments) == 2 {
		registermap[arguments[0]] = arguments[1]
	} else {
		registermap, _ = ini.GetKvmap("register")
	}
	params := cloudstack.NewListOstypesParameter(strings.Replace(ostype, "%", "%25", -1))
	ostypes, _ := client.ListOstypes(params)
	var ostypeid string
	if len(ostypes) > 0 {
		ostypeid = ostypes[0].Id.String()
	} else {
		fmt.Println("ostype is not exist")
		os.Exit(1)
	}

	params1 := cloudstack.NewListZonesParameter(strings.Replace(zonename, "%", "%25", -1))
	zones, _ := client.ListZones(params1)
	var zoneid string
	if len(zones) > 0 {
		zoneid = zones[0].Id.String()
	} else {
		fmt.Println("zonename is not exist")
		os.Exit(1)
	}

	// registering a new template.
	for k, v := range registermap {
		url := strings.Replace(v, ":", "%3A", -1)
		url = strings.Replace(url, "/", "%2F", -1)
		fmt.Println(k, v)
		params2 := cloudstack.NewRegisterTemplateParameter(k, format, hypervisor, k, ostypeid, url, zoneid)
		params2.IsPublic.Set(true)
		params2.PasswordEnabled.Set(passwordenabled)
		templates, err := client.RegisterTemplate(params2)
		if err == nil {
			fmt.Println("return template id : ", templates[0].Id.String())
                        callstatus(client,templates[0].Id.String())
		} else {
			fmt.Println(err.Error())
		}
	}
	fmt.Println("register ", len(registermap), " templates.")
}

func calldelete(ini *goini.INI, client *cloudstack.Client, arguments []string) {
	deletemap := make(map[string]string)
	if len(arguments) > 0 {
		for k, v := range arguments {
			deletemap["id"+strconv.Itoa(k)] = v
		}
	} else {
		deletemap, _ = ini.GetKvmap("delete")
	}
	for _, v := range deletemap {
		params := cloudstack.NewDeleteTemplateParameter(v)
		templates, err := client.DeleteTemplate(params)
		if err == nil {
			fmt.Println(v)
			b, _ := json.MarshalIndent(templates, "", "    ")
			fmt.Println(string(b))
		} else {
			fmt.Println(err.Error())
		}
	}
	fmt.Println("delete ", len(deletemap), " templates.")
}

func calllist(ini *goini.INI, client *cloudstack.Client, arguments []string, options map[string]getopt.OptionValue) {
	mapa, _ := ini.GetKvmap("list")
	if options["columns"].Set {
		for k, _ := range mapa {
			if k != "keyword" {
				mapa[k] = "false"
			}
		}
		for _, v := range strings.Split(options["columns"].String, ",") {
			mapa[v] = "true"
		}
	}
	keyword := mapa["keyword"]
	if options["keyword"].Set {
		keyword = options["keyword"].String
	}
	params := cloudstack.NewListTemplatesParameter("all")
	if keyword != "all" {
		params.Keyword.Set(strings.Replace(keyword, "%", "%25", -1))
	}
	templates, _ := client.ListTemplates(params)
	fmt.Println("total:", len(templates))
	num := []int{}
	for k, v := range templates {
		val := reflect.ValueOf(v).Elem()
		if k == 0 {
			for i := 0; i < val.NumField(); i++ {
				if mapa[strings.ToLower(val.Type().Field(i).Name)] == "true" {
					num = append(num, i)
				}
			}
			for _, i := range num {
				//fmt.Printf("|"+strconv.Itoa(i)+val.Type().Field(i).Name)
				fmt.Printf(" | " + val.Type().Field(i).Name)
			}
			fmt.Printf(" |\n ")
			for _, _ = range num {
				fmt.Printf("|----")
			}
			fmt.Printf("|\n")
		}
		//fmt.Printf("|"+val.Field(1).Field(0).MethodByName("String").Call([]reflect.Value{})[0].String()+"|\n")
		for _, i := range num {
			if val.Type().Field(i).Name == "ResourceBase" || val.Type().Field(i).Name == "Tags" {
				//fmt.Printf("|"+strconv.Itoa(i)+val.Type().Field(i).Name)
				fmt.Printf(" | " + val.Type().Field(i).Name)
			} else if val.Type().Field(i).Name == "ID" {
				//fmt.Printf("|"+strconv.Itoa(i)+val.Field(i).Field(0).Field(0).MethodByName("String").Call([]reflect.Value{})[0].String())
				fmt.Printf(" | " + val.Field(i).Field(0).Field(0).MethodByName("String").Call([]reflect.Value{})[0].String())
			} else {
				//fmt.Printf("|"+strconv.Itoa(i)+val.Field(i).Field(0).MethodByName("String").Call([]reflect.Value{})[0].String())
				fmt.Printf(" | " + val.Field(i).Field(0).MethodByName("String").Call([]reflect.Value{})[0].String())
			}
		}
		fmt.Printf(" |\n")
	}
}

func callstatus(client *cloudstack.Client,id string) {
	params := cloudstack.NewListTemplatesParameter("all")
	params.Id.Set(id)
        for {
            templates, _ := client.ListTemplates(params)
            if len(templates)==0 {
              fmt.Println("id not exist ",id)
              break
            }
            if templates[0].IsReady.Bool(){
              fmt.Println("IsReady:",templates[0].IsReady.Bool())
              break
            }
            fmt.Println("IsReady:",templates[0].IsReady.Bool(),", wait 30s" )
            time.Sleep(30 * time.Second)
        }
}

