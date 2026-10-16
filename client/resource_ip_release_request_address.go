package client

// Type for resource Ip_release_request.Address
type ResourceIpReleaseRequestAddress struct {
	// Pointer to client
	Client *Client

	// Action Ip_release_request.Address#Exempt
	Exempt *ActionIpReleaseRequestAddressExempt
	// Action Ip_release_request.Address#Index
	Index *ActionIpReleaseRequestAddressIndex
	// Action Ip_release_request.Address#Index
	List *ActionIpReleaseRequestAddressIndex
}

func NewResourceIpReleaseRequestAddress(client *Client) *ResourceIpReleaseRequestAddress {
	actionExempt := NewActionIpReleaseRequestAddressExempt(client)
	actionIndex := NewActionIpReleaseRequestAddressIndex(client)

	return &ResourceIpReleaseRequestAddress{
		Client: client,
		Exempt: actionExempt,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
