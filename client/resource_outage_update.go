package client

// Type for resource Outage_update
type ResourceOutageUpdate struct {
	// Pointer to client
	Client *Client

	// Action Outage_update#Create
	Create *ActionOutageUpdateCreate
	// Action Outage_update#Create
	New *ActionOutageUpdateCreate
	// Action Outage_update#Index
	Index *ActionOutageUpdateIndex
	// Action Outage_update#Index
	List *ActionOutageUpdateIndex
	// Action Outage_update#Show
	Show *ActionOutageUpdateShow
	// Action Outage_update#Show
	Find *ActionOutageUpdateShow
}

func NewResourceOutageUpdate(client *Client) *ResourceOutageUpdate {
	actionCreate := NewActionOutageUpdateCreate(client)
	actionIndex := NewActionOutageUpdateIndex(client)
	actionShow := NewActionOutageUpdateShow(client)

	return &ResourceOutageUpdate{
		Client: client,
		Create: actionCreate,
		New:    actionCreate,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
