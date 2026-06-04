package client

// Type for resource User_security_advisory
type ResourceUserSecurityAdvisory struct {
	// Pointer to client
	Client *Client

	// Action User_security_advisory#Index
	Index *ActionUserSecurityAdvisoryIndex
	// Action User_security_advisory#Index
	List *ActionUserSecurityAdvisoryIndex
}

func NewResourceUserSecurityAdvisory(client *Client) *ResourceUserSecurityAdvisory {
	actionIndex := NewActionUserSecurityAdvisoryIndex(client)

	return &ResourceUserSecurityAdvisory{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
