package client

// Type for resource Event_delivery_group
type ResourceEventDeliveryGroup struct {
	// Pointer to client
	Client *Client

	// Action Event_delivery_group#Index
	Index *ActionEventDeliveryGroupIndex
	// Action Event_delivery_group#Index
	List *ActionEventDeliveryGroupIndex
	// Action Event_delivery_group#Show
	Show *ActionEventDeliveryGroupShow
	// Action Event_delivery_group#Show
	Find *ActionEventDeliveryGroupShow
}

func NewResourceEventDeliveryGroup(client *Client) *ResourceEventDeliveryGroup {
	actionIndex := NewActionEventDeliveryGroupIndex(client)
	actionShow := NewActionEventDeliveryGroupShow(client)

	return &ResourceEventDeliveryGroup{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
