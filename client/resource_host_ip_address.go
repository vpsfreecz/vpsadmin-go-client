package client

// Type for resource Host_ip_address
type ResourceHostIpAddress struct {
	// Pointer to client
	Client *Client

	// Action Host_ip_address#Assign
	Assign *ActionHostIpAddressAssign
	// Action Host_ip_address#Create
	Create *ActionHostIpAddressCreate
	// Action Host_ip_address#Create
	New *ActionHostIpAddressCreate
	// Action Host_ip_address#Delete
	Delete *ActionHostIpAddressDelete
	// Action Host_ip_address#Delete
	Destroy *ActionHostIpAddressDelete
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
	// Action Host_ip_address#Update
	Update *ActionHostIpAddressUpdate
}

func NewResourceHostIpAddress(client *Client) *ResourceHostIpAddress {
	actionAssign := NewActionHostIpAddressAssign(client)
	actionCreate := NewActionHostIpAddressCreate(client)
	actionDelete := NewActionHostIpAddressDelete(client)
	actionFree := NewActionHostIpAddressFree(client)
	actionIndex := NewActionHostIpAddressIndex(client)
	actionShow := NewActionHostIpAddressShow(client)
	actionUpdate := NewActionHostIpAddressUpdate(client)

	return &ResourceHostIpAddress{
		Client:  client,
		Assign:  actionAssign,
		Create:  actionCreate,
		New:     actionCreate,
		Delete:  actionDelete,
		Destroy: actionDelete,
		Free:    actionFree,
		Index:   actionIndex,
		List:    actionIndex,
		Show:    actionShow,
		Find:    actionShow,
		Update:  actionUpdate,
	}
}
