package client

// Type for resource Network_interface_monitor
type ResourceNetworkInterfaceMonitor struct {
	// Pointer to client
	Client *Client

	// Action Network_interface_monitor#Index
	Index *ActionNetworkInterfaceMonitorIndex
	// Action Network_interface_monitor#Index
	List *ActionNetworkInterfaceMonitorIndex
	// Action Network_interface_monitor#Show
	Show *ActionNetworkInterfaceMonitorShow
	// Action Network_interface_monitor#Show
	Find *ActionNetworkInterfaceMonitorShow
}

func NewResourceNetworkInterfaceMonitor(client *Client) *ResourceNetworkInterfaceMonitor {
	actionIndex := NewActionNetworkInterfaceMonitorIndex(client)
	actionShow := NewActionNetworkInterfaceMonitorShow(client)

	return &ResourceNetworkInterfaceMonitor{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
