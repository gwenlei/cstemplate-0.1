package cloudstack

type ListRouterParameter struct {
	Account 		NullString 
	ClusterId	ID
	DomainId		ID
	Forvpc		NullBool
	HostId 		ID
	Id 			ID
	IsRecursive NullBool
	Keyword		NullString
	Listall		NullBool
	Name			NullString
	NetworkId	ID
	Page      	NullNumber
	PageSize  	NullNumber
	PodId		ID
	ProjectId	ID
	State		NullString
	Version		NullString
	VpcId		ID
	ZoneId		ID
}

type Router struct {
	ResourceBase
	Id 		ID		`json:"id"`
	Account	ID		`json:"account"`
	Created NullString	`json:"created"`	
	Dns1		NullString	`json:"dns1"`
	Dns2		NullString	`json:"dns2"`
	Domain	NullString	`json:"domain"`
	DomainId	 ID		`json:"domainid"`
	Gateway	NullString	`json:"gateway"`
	GuestIpaddress	NullString	`json:"guestipaddress"`
	GuestMacaddress	NullString	`json:"guestmacaddress"`
	GuestNetmask		NullString	`json:"guestnetmask"`
	GuestNetworkid	ID	`json:"guestnetworkid"`
	HostId	ID	`json:"hostid"`
	HostName		NullString `json:"hostname"`
	Hypervisor	NullString	`json:"hypervisor"`
	Ip6Dns1	NullString	`json:"ip6dns1"`
	Ip6Dns2	NullString	`json:"ip6dns2"`
	IsRedundantRouter NullBool `json:"isredundantrouter"`
	LinkLocalIp	NullString	`json:"linklocalip"`
	LinkLocalMacAddress	NullString `json:"linklocalmacaddress"`
	LinkLocalNetmask		NullString `json:"linklocalnetmask"`
	LinkLocalNetworkId	ID	`json:"linklocalnetworkid"`
	Name		NullString	`json:"name"`
	NetworkDomain	NullString `json:"networkdomain"`
	Podid	ID		`json:"poid"`
	Project NullString	`json:"project"`
	ProjectId	NullString	`json:"projectid"`
	PublicIp		NullString	`json:"publicip"`
	PublicMacAddress NullString	`json;"publicmacaddress"`
	PublicNetmask	NullString	`json:"publicnetmask"`
	PublicNetworkId	ID	`json:"publicnetworkid"`
	RedundantState	NullString	`json:"redundantstate"`
	RequireSupgrade	NullString	`json:"requiresupgrade"`
	Role		NullString	`json:"role"`
	ScriptsVersion	NullString	`json:"scriptsversion"`
	ServiceOfferingId	ID	`json:"serviceofferingid"`
	ServiceOfferingName	NullString	`json:"serviceofferingname"`
	State NullString		`json:"state"`
	TemplateId	ID	`json:"templateid"`
	Version	NullString	`json:"version"`
	VpcId	ID	`json:"vpcid"`
	ZoneId	ID	`json:"zoneid"`
	ZoneName NullString `json:"zonename"`
	Nics []Nic	`json:"nic"`
}

func NewListRouterParam() (p *ListRouterParameter) {
	p = new(ListRouterParameter)
	return p
}

func (c *Client) ListRouter(p *ListRouterParameter) ([]*Router, error) {
	obj, err := c.Request("listRouters", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*Router), err
}