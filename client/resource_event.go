package client

// Type for resource Event
type ResourceEvent struct {
	// Pointer to client
	Client *Client

	// Resource Event.Delivery
	Delivery *ResourceEventDelivery
	// Resource Event.Route_match
	RouteMatch *ResourceEventRouteMatch
	// Action Event#Index
	Index *ActionEventIndex
	// Action Event#Index
	List *ActionEventIndex
	// Action Event#Show
	Show *ActionEventShow
	// Action Event#Show
	Find *ActionEventShow
	// Action Event#Test
	Test *ActionEventTest
}

func NewResourceEvent(client *Client) *ResourceEvent {
	actionIndex := NewActionEventIndex(client)
	actionShow := NewActionEventShow(client)
	actionTest := NewActionEventTest(client)

	return &ResourceEvent{
		Client:     client,
		Delivery:   NewResourceEventDelivery(client),
		RouteMatch: NewResourceEventRouteMatch(client),
		Index:      actionIndex,
		List:       actionIndex,
		Show:       actionShow,
		Find:       actionShow,
		Test:       actionTest,
	}
}
