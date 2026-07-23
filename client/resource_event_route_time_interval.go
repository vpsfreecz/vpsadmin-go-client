package client

// Type for resource Event_route.Time_interval
type ResourceEventRouteTimeInterval struct {
	// Pointer to client
	Client *Client

	// Action Event_route.Time_interval#Create
	Create *ActionEventRouteTimeIntervalCreate
	// Action Event_route.Time_interval#Create
	New *ActionEventRouteTimeIntervalCreate
	// Action Event_route.Time_interval#Delete
	Delete *ActionEventRouteTimeIntervalDelete
	// Action Event_route.Time_interval#Delete
	Destroy *ActionEventRouteTimeIntervalDelete
	// Action Event_route.Time_interval#Index
	Index *ActionEventRouteTimeIntervalIndex
	// Action Event_route.Time_interval#Index
	List *ActionEventRouteTimeIntervalIndex
	// Action Event_route.Time_interval#Show
	Show *ActionEventRouteTimeIntervalShow
	// Action Event_route.Time_interval#Show
	Find *ActionEventRouteTimeIntervalShow
	// Action Event_route.Time_interval#Update
	Update *ActionEventRouteTimeIntervalUpdate
}

func NewResourceEventRouteTimeInterval(client *Client) *ResourceEventRouteTimeInterval {
	actionCreate := NewActionEventRouteTimeIntervalCreate(client)
	actionDelete := NewActionEventRouteTimeIntervalDelete(client)
	actionIndex := NewActionEventRouteTimeIntervalIndex(client)
	actionShow := NewActionEventRouteTimeIntervalShow(client)
	actionUpdate := NewActionEventRouteTimeIntervalUpdate(client)

	return &ResourceEventRouteTimeInterval{
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
