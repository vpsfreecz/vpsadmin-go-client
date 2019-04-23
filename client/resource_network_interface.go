package client

// Type for resource Network_interface
type ResourceNetworkInterface struct {
	// Pointer to client
	Client *Client

	// Action Network_interface#Index
	Index *ActionNetworkInterfaceIndex
	// Action Network_interface#Index
	List *ActionNetworkInterfaceIndex
	// Action Network_interface#Show
	Show *ActionNetworkInterfaceShow
	// Action Network_interface#Show
	Find *ActionNetworkInterfaceShow
	// Action Network_interface#Update
	Update *ActionNetworkInterfaceUpdate
}

func NewResourceNetworkInterface(client *Client) *ResourceNetworkInterface {
	actionIndex := NewActionNetworkInterfaceIndex(client)
	actionShow := NewActionNetworkInterfaceShow(client)
	actionUpdate := NewActionNetworkInterfaceUpdate(client)

	return &ResourceNetworkInterface{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Update: actionUpdate,
	}
}
