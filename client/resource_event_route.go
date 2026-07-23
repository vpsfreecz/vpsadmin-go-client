package client

// Type for resource Event_route
type ResourceEventRoute struct {
	// Pointer to client
	Client *Client

	// Resource Event_route.Matcher
	Matcher *ResourceEventRouteMatcher
	// Resource Event_route.Time_interval
	TimeInterval *ResourceEventRouteTimeInterval
	// Action Event_route#Create
	Create *ActionEventRouteCreate
	// Action Event_route#Create
	New *ActionEventRouteCreate
	// Action Event_route#Delete
	Delete *ActionEventRouteDelete
	// Action Event_route#Delete
	Destroy *ActionEventRouteDelete
	// Action Event_route#Index
	Index *ActionEventRouteIndex
	// Action Event_route#Index
	List *ActionEventRouteIndex
	// Action Event_route#Show
	Show *ActionEventRouteShow
	// Action Event_route#Show
	Find *ActionEventRouteShow
	// Action Event_route#Update
	Update *ActionEventRouteUpdate
}

func NewResourceEventRoute(client *Client) *ResourceEventRoute {
	actionCreate := NewActionEventRouteCreate(client)
	actionDelete := NewActionEventRouteDelete(client)
	actionIndex := NewActionEventRouteIndex(client)
	actionShow := NewActionEventRouteShow(client)
	actionUpdate := NewActionEventRouteUpdate(client)

	return &ResourceEventRoute{
		Client:       client,
		Matcher:      NewResourceEventRouteMatcher(client),
		TimeInterval: NewResourceEventRouteTimeInterval(client),
		Create:       actionCreate,
		New:          actionCreate,
		Delete:       actionDelete,
		Destroy:      actionDelete,
		Index:        actionIndex,
		List:         actionIndex,
		Show:         actionShow,
		Find:         actionShow,
		Update:       actionUpdate,
	}
}
