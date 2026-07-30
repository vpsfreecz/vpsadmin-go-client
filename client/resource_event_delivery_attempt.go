package client

// Type for resource Event.Delivery.Attempt
type ResourceEventDeliveryAttempt struct {
	// Pointer to client
	Client *Client

	// Action Event.Delivery.Attempt#Index
	Index *ActionEventDeliveryAttemptIndex
	// Action Event.Delivery.Attempt#Index
	List *ActionEventDeliveryAttemptIndex
	// Action Event.Delivery.Attempt#Show
	Show *ActionEventDeliveryAttemptShow
	// Action Event.Delivery.Attempt#Show
	Find *ActionEventDeliveryAttemptShow
}

func NewResourceEventDeliveryAttempt(client *Client) *ResourceEventDeliveryAttempt {
	actionIndex := NewActionEventDeliveryAttemptIndex(client)
	actionShow := NewActionEventDeliveryAttemptShow(client)

	return &ResourceEventDeliveryAttempt{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
