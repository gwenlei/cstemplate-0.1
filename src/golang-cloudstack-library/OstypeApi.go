package cloudstack


// ListOstypes represents the paramter of ListOstypes
type ListOstypesParameter struct {
	// the Ostype ID
	Id ID 
	// the Ostype description
	Description NullString 
	// the Ostype isuserdefined
	Keyword NullString
	// the Ostype oscategoryid
	Oscategoryid NullString 
}

func NewListOstypesParameter(Keyword string) (p *ListOstypesParameter) {
	p = new(ListOstypesParameter)
        p.Keyword.Set(Keyword)
	return p
}

// List all ostypes.
func (c *Client) ListOstypes(p *ListOstypesParameter) ([]*Ostype, error) {
	obj, err := c.Request("listOstypes", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*Ostype), err
}


