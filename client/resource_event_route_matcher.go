package client

// Type for resource Event_route.Matcher
type ResourceEventRouteMatcher struct {
	// Pointer to client
	Client *Client

	// Action Event_route.Matcher#Create
	Create *ActionEventRouteMatcherCreate
	// Action Event_route.Matcher#Create
	New *ActionEventRouteMatcherCreate
	// Action Event_route.Matcher#Delete
	Delete *ActionEventRouteMatcherDelete
	// Action Event_route.Matcher#Delete
	Destroy *ActionEventRouteMatcherDelete
	// Action Event_route.Matcher#Index
	Index *ActionEventRouteMatcherIndex
	// Action Event_route.Matcher#Index
	List *ActionEventRouteMatcherIndex
	// Action Event_route.Matcher#Show
	Show *ActionEventRouteMatcherShow
	// Action Event_route.Matcher#Show
	Find *ActionEventRouteMatcherShow
	// Action Event_route.Matcher#Update
	Update *ActionEventRouteMatcherUpdate
}

func NewResourceEventRouteMatcher(client *Client) *ResourceEventRouteMatcher {
	actionCreate := NewActionEventRouteMatcherCreate(client)
	actionDelete := NewActionEventRouteMatcherDelete(client)
	actionIndex := NewActionEventRouteMatcherIndex(client)
	actionShow := NewActionEventRouteMatcherShow(client)
	actionUpdate := NewActionEventRouteMatcherUpdate(client)

	return &ResourceEventRouteMatcher{
		Client:  client,
		Create:  actionCreate,
		New:     actionCreate,
		Delete:  actionDelete,
		Destroy: actionDelete,
		Index:   actionIndex,
		List:    actionIndex,
		Show:    actionShow,
		Find:    actionShow,
		Update:  actionUpdate,
	}
}
