package client

// Type for resource Ip_address
type ResourceIpAddress struct {
	// Pointer to client
	Client *Client

	// Action Ip_address#Assign
	Assign *ActionIpAddressAssign
	// Action Ip_address#Assign_with_host_address
	AssignWithHostAddress *ActionIpAddressAssignWithHostAddress
	// Action Ip_address#Create
	Create *ActionIpAddressCreate
	// Action Ip_address#Create
	New *ActionIpAddressCreate
	// Action Ip_address#Free
	Free *ActionIpAddressFree
	// Action Ip_address#Index
	Index *ActionIpAddressIndex
	// Action Ip_address#Index
	List *ActionIpAddressIndex
	// Action Ip_address#Show
	Show *ActionIpAddressShow
	// Action Ip_address#Show
	Find *ActionIpAddressShow
	// Action Ip_address#Update
	Update *ActionIpAddressUpdate
}

func NewResourceIpAddress(client *Client) *ResourceIpAddress {
	actionAssign := NewActionIpAddressAssign(client)
	actionAssignWithHostAddress := NewActionIpAddressAssignWithHostAddress(client)
	actionCreate := NewActionIpAddressCreate(client)
	actionFree := NewActionIpAddressFree(client)
	actionIndex := NewActionIpAddressIndex(client)
	actionShow := NewActionIpAddressShow(client)
	actionUpdate := NewActionIpAddressUpdate(client)

	return &ResourceIpAddress{
		Client:                client,
		Assign:                actionAssign,
		AssignWithHostAddress: actionAssignWithHostAddress,
		Create:                actionCreate,
		New:                   actionCreate,
		Free:                  actionFree,
		Index:                 actionIndex,
		List:                  actionIndex,
		Show:                  actionShow,
		Find:                  actionShow,
		Update:                actionUpdate,
	}
}
