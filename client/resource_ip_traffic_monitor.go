package client

// Type for resource Ip_traffic_monitor
type ResourceIpTrafficMonitor struct {
	// Pointer to client
	Client *Client

	// Action Ip_traffic_monitor#Index
	Index *ActionIpTrafficMonitorIndex
	// Action Ip_traffic_monitor#Index
	List *ActionIpTrafficMonitorIndex
	// Action Ip_traffic_monitor#Show
	Show *ActionIpTrafficMonitorShow
	// Action Ip_traffic_monitor#Show
	Find *ActionIpTrafficMonitorShow
}

func NewResourceIpTrafficMonitor(client *Client) *ResourceIpTrafficMonitor {
	actionIndex := NewActionIpTrafficMonitorIndex(client)
	actionShow := NewActionIpTrafficMonitorShow(client)

	return &ResourceIpTrafficMonitor{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
	}
}
