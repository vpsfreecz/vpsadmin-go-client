package client

// Type for resource Outage.Handler
type ResourceOutageHandler struct {
	// Pointer to client
	Client *Client

	// Action Outage.Handler#Create
	Create *ActionOutageHandlerCreate
	// Action Outage.Handler#Create
	New *ActionOutageHandlerCreate
	// Action Outage.Handler#Delete
	Delete *ActionOutageHandlerDelete
	// Action Outage.Handler#Delete
	Destroy *ActionOutageHandlerDelete
	// Action Outage.Handler#Index
	Index *ActionOutageHandlerIndex
	// Action Outage.Handler#Index
	List *ActionOutageHandlerIndex
	// Action Outage.Handler#Show
	Show *ActionOutageHandlerShow
	// Action Outage.Handler#Show
	Find *ActionOutageHandlerShow
	// Action Outage.Handler#Update
	Update *ActionOutageHandlerUpdate
}

func NewResourceOutageHandler(client *Client) *ResourceOutageHandler {
	actionCreate := NewActionOutageHandlerCreate(client)
	actionDelete := NewActionOutageHandlerDelete(client)
	actionIndex := NewActionOutageHandlerIndex(client)
	actionShow := NewActionOutageHandlerShow(client)
	actionUpdate := NewActionOutageHandlerUpdate(client)

	return &ResourceOutageHandler{
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
