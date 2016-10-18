package cloudstack

type ListSystemVmsParam struct {
	// the host ID of the system VM
	HostId	ID	
	// the ID of the system VM
	Id 	ID	
	// List by keyword
	Keywork	NullString
	// the name of the system VM
	Name		NullString
	Page		NullNumber
	PageSize		NullNumber
	// the Pod ID of the system VM
	PodId	ID
	State	NullString
	// the storage ID where vm's volumes belong to
	StorageId ID 
	SystemVmType NullString
	ZoneId	ID
}

type SystemVm struct {
	Id ID `json:"id"`
	ActiveViewerSessions NullNumber	`json:"activeviewersessions"`
	Created NullString	`json:"created"`
	Dns1 NullString	`json:"dns1"`
	Dns2	 NullString	`json:"dns2"`
	Gateway	NullString	`json:"gateway"`
	HostId	ID	`json:"hostid"`
	HostName	 NullString	`json:"hostname"`
	Hypervisor	NullString `json:"hypervisor"`
	JobId	ID	`json:"jobid"`
	JobStatus	NullString 	`json:"jobstatus"`
	LinkLocalIp	NullString	`json:"linklocalip"`
	LinkLocalMacAddress NullString	`json:"linklocalmacaddress"`
	LinkLocalNetmask		NullString	`json:"linklocalnetmask"`
	Name		NullString	`json:"name"`
	NetworkDomain	NullString	`json:"networkdomain"`
	PodId	ID	`json:"podid"`
	PrivateIp	NullString	`json:"privateip"`
	PrivateMacAddress	NullString	`json:"privatemacaddress"`
	PrivateNetmask	NullString	`json:"privatenetmask"`
	PublicIp		NullString	`json:"publicip"`
	PublicMacAddress		NullString	`json:"publicmacaddress"`
	PublicNetmask	NullString	`json:"publicnetmask"`
	State	NullString	`json:"systemvmtype"`
	TemplateId	NullString	`json:"templateid"`
	ZoneId	ID	`json:"zoneid"`
	ZoneName		NullString	`json:"zonename"`
}


func NewListSystemVmsParam() (p *ListSystemVmsParam) {
	p = new(ListSystemVmsParam)
	return p
}

func (c *Client) ListSystemVms(p *ListSystemVmsParam) ([]*SystemVm, error) {
	obj, err := c.Request("listSystemVms", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*SystemVm), err
}
