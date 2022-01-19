package client

// Type for resource Monitored_event
type ResourceMonitoredEvent struct {
	// Pointer to client
	Client *Client

	// Resource Monitored_event.Log
	Log *ResourceMonitoredEventLog
	// Action Monitored_event#Acknowledge
	Acknowledge *ActionMonitoredEventAcknowledge
	// Action Monitored_event#Acknowledge
	Ack *ActionMonitoredEventAcknowledge
	// Action Monitored_event#Ignore
	Ignore *ActionMonitoredEventIgnore
	// Action Monitored_event#Index
	Index *ActionMonitoredEventIndex
	// Action Monitored_event#Index
	List *ActionMonitoredEventIndex
	// Action Monitored_event#Show
	Show *ActionMonitoredEventShow
	// Action Monitored_event#Show
	Find *ActionMonitoredEventShow
}

func NewResourceMonitoredEvent(client *Client) *ResourceMonitoredEvent {
	actionAcknowledge := NewActionMonitoredEventAcknowledge(client)
	actionIgnore := NewActionMonitoredEventIgnore(client)
	actionIndex := NewActionMonitoredEventIndex(client)
	actionShow := NewActionMonitoredEventShow(client)

	return &ResourceMonitoredEvent{
		Client: client,
		Log: NewResourceMonitoredEventLog(client),
		Acknowledge: actionAcknowledge,
		Ack: actionAcknowledge,
		Ignore: actionIgnore,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
	}
}
