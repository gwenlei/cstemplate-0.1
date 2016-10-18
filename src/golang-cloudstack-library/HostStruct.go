package cloudstack

type Host struct {
	ResourceBase
	Id 	ID `json:"id"`
	Name NullString `json:"name"`
	State NullString `json:"state"`
	Disconnected NullString `json:"disconnected"`
	Type NullString `json:"type"`
	IP NullString `json:"ipaddress"`
	ZoneId ID `json:"zoneid"`
	ZoneName NullString `json:"zonename"`
	PodId ID `json:"podid"`
	PodName NullString `json:"podname"`
	Version NullString `json:"version"`
	Hypervisor NullString `json:"hypervisor"`
	CPUSockets NullNumber 	`json:"cpusockets"`
	CPUNumber NullNumber 	`json:"cpunumber"`
	CPUSpeed  NullNumber	`json:"cpuspeed"`
	CPUAllocated NullString `json:"cpuallocated"`
	CPUWithoverProvisioning NullNumber `json:"cpuwithoverprovisioning"`
	Memorytotal NullNumber `json:"memorytotal"`
	MemoryAllocated NullNumber `json:"memoryallocated"`
	Capabilities	NullString `json:"capabilities"`
	LastPinged NullString `json:"lastpinged"`
	ManagementServerId NullString `json:"managementserverid"`
	ClusterId ID 	`json:"clusterid"`
	ClusterName NullString 	`json:"clustername"`
	ClusterType NullString 	`json:"clustertype"`
	IsLocalStorageActive NullBool `json:"islocalstorageactive"`
	Created NullString 		`json:"created"`
	Events NullString	`json:"events"`
	HosTags NullString `json:"hosttag"`
	ResourceState NullString `json"resourcestate"`
	Hahost NullBool	`json:"hahost"`
//	Details NullString 		`json:"details"`
}
