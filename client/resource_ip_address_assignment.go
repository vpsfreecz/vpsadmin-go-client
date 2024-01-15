package client

// Type for resource Ip_address_assignment
type ResourceIpAddressAssignment struct {
	// Pointer to client
	Client *Client

	// Action Ip_address_assignment#Index
	Index *ActionIpAddressAssignmentIndex
	// Action Ip_address_assignment#Index
	List *ActionIpAddressAssignmentIndex
	// Action Ip_address_assignment#Show
	Show *ActionIpAddressAssignmentShow
	// Action Ip_address_assignment#Show
	Find *ActionIpAddressAssignmentShow
}

func NewResourceIpAddressAssignment(client *Client) *ResourceIpAddressAssignment {
	actionIndex := NewActionIpAddressAssignmentIndex(client)
	actionShow := NewActionIpAddressAssignmentShow(client)

	return &ResourceIpAddressAssignment{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
