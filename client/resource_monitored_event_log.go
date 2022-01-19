package client

// Type for resource Monitored_event.Log
type ResourceMonitoredEventLog struct {
	// Pointer to client
	Client *Client

	// Action Monitored_event.Log#Index
	Index *ActionMonitoredEventLogIndex
	// Action Monitored_event.Log#Index
	List *ActionMonitoredEventLogIndex
	// Action Monitored_event.Log#Show
	Show *ActionMonitoredEventLogShow
	// Action Monitored_event.Log#Show
	Find *ActionMonitoredEventLogShow
}

func NewResourceMonitoredEventLog(client *Client) *ResourceMonitoredEventLog {
	actionIndex := NewActionMonitoredEventLogIndex(client)
	actionShow := NewActionMonitoredEventLogShow(client)

	return &ResourceMonitoredEventLog{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
