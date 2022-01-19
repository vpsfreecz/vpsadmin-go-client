package client

// Type for resource Host_ip_address
type ResourceHostIpAddress struct {
	// Pointer to client
	Client *Client

	// Action Host_ip_address#Assign
	Assign *ActionHostIpAddressAssign
	// Action Host_ip_address#Free
	Free *ActionHostIpAddressFree
	// Action Host_ip_address#Index
	Index *ActionHostIpAddressIndex
	// Action Host_ip_address#Index
	List *ActionHostIpAddressIndex
	// Action Host_ip_address#Show
	Show *ActionHostIpAddressShow
	// Action Host_ip_address#Show
	Find *ActionHostIpAddressShow
}

func NewResourceHostIpAddress(client *Client) *ResourceHostIpAddress {
	actionAssign := NewActionHostIpAddressAssign(client)
	actionFree := NewActionHostIpAddressFree(client)
	actionIndex := NewActionHostIpAddressIndex(client)
	actionShow := NewActionHostIpAddressShow(client)

	return &ResourceHostIpAddress{
		Client: client,
		Assign: actionAssign,
		Free: actionFree,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
	}
}
