package cloudstack

type Ostype struct {
        ResourceBase
	// the Ostype ID
	Id ID `json:"id"`
	// the Ostype description
	Description NullString `json:"description"`
	// the Ostype isuserdefined
	Isuserdefined NullString `json:"isuserdefined"`
	// the Ostype oscategoryid
	Oscategoryid NullString `json:"oscategoryid"`
}
