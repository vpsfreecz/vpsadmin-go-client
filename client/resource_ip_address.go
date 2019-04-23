package client

// Type for resource Ip_address
type ResourceIpAddress struct {
	// Pointer to client
	Client *Client

	// Action Ip_address#Index
	Index *ActionIpAddressIndex
	// Action Ip_address#Index
	List *ActionIpAddressIndex
	// Action Ip_address#Show
	Show *ActionIpAddressShow
	// Action Ip_address#Show
	Find *ActionIpAddressShow
	// Action Ip_address#Create
	Create *ActionIpAddressCreate
	// Action Ip_address#Create
	New *ActionIpAddressCreate
	// Action Ip_address#Update
	Update *ActionIpAddressUpdate
	// Action Ip_address#Assign
	Assign *ActionIpAddressAssign
	// Action Ip_address#Assign_with_host_address
	AssignWithHostAddress *ActionIpAddressAssignWithHostAddress
	// Action Ip_address#Free
	Free *ActionIpAddressFree
}

func NewResourceIpAddress(client *Client) *ResourceIpAddress {
	actionIndex := NewActionIpAddressIndex(client)
	actionShow := NewActionIpAddressShow(client)
	actionCreate := NewActionIpAddressCreate(client)
	actionUpdate := NewActionIpAddressUpdate(client)
	actionAssign := NewActionIpAddressAssign(client)
	actionAssignWithHostAddress := NewActionIpAddressAssignWithHostAddress(client)
	actionFree := NewActionIpAddressFree(client)

	return &ResourceIpAddress{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Create: actionCreate,
		New: actionCreate,
		Update: actionUpdate,
		Assign: actionAssign,
		AssignWithHostAddress: actionAssignWithHostAddress,
		Free: actionFree,
	}
}
