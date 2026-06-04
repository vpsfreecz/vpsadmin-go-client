package client

// Type for resource Vps_security_advisory
type ResourceVpsSecurityAdvisory struct {
	// Pointer to client
	Client *Client

	// Action Vps_security_advisory#Index
	Index *ActionVpsSecurityAdvisoryIndex
	// Action Vps_security_advisory#Index
	List *ActionVpsSecurityAdvisoryIndex
}

func NewResourceVpsSecurityAdvisory(client *Client) *ResourceVpsSecurityAdvisory {
	actionIndex := NewActionVpsSecurityAdvisoryIndex(client)

	return &ResourceVpsSecurityAdvisory{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
