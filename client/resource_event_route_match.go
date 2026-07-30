package client

// Type for resource Event.Route_match
type ResourceEventRouteMatch struct {
	// Pointer to client
	Client *Client

	// Action Event.Route_match#Index
	Index *ActionEventRouteMatchIndex
	// Action Event.Route_match#Index
	List *ActionEventRouteMatchIndex
	// Action Event.Route_match#Show
	Show *ActionEventRouteMatchShow
	// Action Event.Route_match#Show
	Find *ActionEventRouteMatchShow
}

func NewResourceEventRouteMatch(client *Client) *ResourceEventRouteMatch {
	actionIndex := NewActionEventRouteMatchIndex(client)
	actionShow := NewActionEventRouteMatchShow(client)

	return &ResourceEventRouteMatch{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
