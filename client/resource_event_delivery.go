package client

// Type for resource Event_delivery
type ResourceEventDelivery struct {
	// Pointer to client
	Client *Client

	// Action Event_delivery#Index
	Index *ActionEventDeliveryIndex
	// Action Event_delivery#Index
	List *ActionEventDeliveryIndex
}

func NewResourceEventDelivery(client *Client) *ResourceEventDelivery {
	actionIndex := NewActionEventDeliveryIndex(client)

	return &ResourceEventDelivery{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
