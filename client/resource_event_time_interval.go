package client

// Type for resource Event_time_interval
type ResourceEventTimeInterval struct {
	// Pointer to client
	Client *Client

	// Action Event_time_interval#Create
	Create *ActionEventTimeIntervalCreate
	// Action Event_time_interval#Create
	New *ActionEventTimeIntervalCreate
	// Action Event_time_interval#Delete
	Delete *ActionEventTimeIntervalDelete
	// Action Event_time_interval#Delete
	Destroy *ActionEventTimeIntervalDelete
	// Action Event_time_interval#Index
	Index *ActionEventTimeIntervalIndex
	// Action Event_time_interval#Index
	List *ActionEventTimeIntervalIndex
	// Action Event_time_interval#Show
	Show *ActionEventTimeIntervalShow
	// Action Event_time_interval#Show
	Find *ActionEventTimeIntervalShow
	// Action Event_time_interval#Update
	Update *ActionEventTimeIntervalUpdate
}

func NewResourceEventTimeInterval(client *Client) *ResourceEventTimeInterval {
	actionCreate := NewActionEventTimeIntervalCreate(client)
	actionDelete := NewActionEventTimeIntervalDelete(client)
	actionIndex := NewActionEventTimeIntervalIndex(client)
	actionShow := NewActionEventTimeIntervalShow(client)
	actionUpdate := NewActionEventTimeIntervalUpdate(client)

	return &ResourceEventTimeInterval{
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
