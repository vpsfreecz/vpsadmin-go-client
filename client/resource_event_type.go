package client

// Type for resource Event_type
type ResourceEventType struct {
	// Pointer to client
	Client *Client

	// Action Event_type#Index
	Index *ActionEventTypeIndex
	// Action Event_type#Index
	List *ActionEventTypeIndex
}

func NewResourceEventType(client *Client) *ResourceEventType {
	actionIndex := NewActionEventTypeIndex(client)

	return &ResourceEventType{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
