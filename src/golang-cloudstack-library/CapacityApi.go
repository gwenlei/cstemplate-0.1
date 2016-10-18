package cloudstack

type ListCapacityParameter struct {
	// lists capacity by the Cluster ID
	ClusterId ID
	// recalculate capacities and fetch the latest
	FetchLatest NullBool
	// keyword
	Keyword NullString
	// list capacity by the Pod ID
	PodId  ID
	SortBy ID
	// lists capacity by type* CAPACITY_TYPE_MEMORY = 0* CAPACITY_TYPE_CPU = 1* CAPACITY_TYPE_STORAGE = 2* CAPACITY_TYPE_STORAGE_ALLOCATED = 3* CAPACITY_TYPE_VIRTUAL_NETWORK_PUBLIC_IP = 4* CAPACITY_TYPE_PRIVATE_IP = 5* CAPACITY_TYPE_SECONDARY_STORAGE = 6* CAPACITY_TYPE_VLAN = 7* CAPACITY_TYPE_DIRECT_ATTACHED_PUBLIC_IP = 8* CAPACITY_TYPE_LOCAL_STORAGE = 9.
	Type NullNumber
	// lists capacity by the Zone ID
	ZoneId ID
}

func NewListCapacityParamete() (p *ListCapacityParameter) {
	p = new(ListCapacityParameter)
	return p
}

// Lists zones
func (c *Client) ListCapacity(p *ListCapacityParameter) ([]*Capacity, error) {
	obj, err := c.Request("listCapacity", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*Capacity), err
}
