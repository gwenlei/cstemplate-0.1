package cloudstack
type ListHostParameter struct {
	// lists hosts existing in particular cluster
	ClusterId  ID
	// comma separated list of host details requested, value can be a list of [ min, all, capacity, events, stats]
	Details []string
	// if true, list only hosts dedicated to HA
	HaHost NullBool
	// hypervisor type of host: XenServer,KVM,VMware,Hyperv,BareMetal,Simulator
	Hypervisor NullString
	// List by keyword
	Keyword NullString 
	// the name of the host
	Name NullString
	Page      NullNumber
	PageSize  NullNumber
	// the pod ID
	PodId ID
	// list hosts by resource state. Resource state represents current state determined by admin of host, valule can be one of 
	// [Enabled, Disabled, Unmanaged, PrepareForMaintenance, ErrorInMaintenance, Maintenance, Error]
	ResourceState NullString
	// the state of the host
	State NullString
	// the host type
	Type  NullString
	// lists hosts in the same cluster as this VM and flag hosts with enough CPU/RAm to host this VM
	VirtualmachineId NullString
	ZoneId NullString
}

func NewListHostParam() (p *ListHostParameter) {
	p = new(ListHostParameter)
	return p
}

func (c *Client) ListHost(p *ListHostParameter) ([]*Host, error) {
	obj, err := c.Request("listHosts", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*Host), err
}
