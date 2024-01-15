package client

// Type for resource Network_interface_accounting
type ResourceNetworkInterfaceAccounting struct {
	// Pointer to client
	Client *Client

	// Action Network_interface_accounting#Index
	Index *ActionNetworkInterfaceAccountingIndex
	// Action Network_interface_accounting#Index
	List *ActionNetworkInterfaceAccountingIndex
	// Action Network_interface_accounting#User_top
	UserTop *ActionNetworkInterfaceAccountingUserTop
}

func NewResourceNetworkInterfaceAccounting(client *Client) *ResourceNetworkInterfaceAccounting {
	actionIndex := NewActionNetworkInterfaceAccountingIndex(client)
	actionUserTop := NewActionNetworkInterfaceAccountingUserTop(client)

	return &ResourceNetworkInterfaceAccounting{
		Client:  client,
		Index:   actionIndex,
		List:    actionIndex,
		UserTop: actionUserTop,
	}
}
