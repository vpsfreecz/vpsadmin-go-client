package client

// Type for resource Outage.Handler
type ResourceOutageHandler struct {
	// Pointer to client
	Client *Client

	// Action Outage.Handler#Index
	Index *ActionOutageHandlerIndex
	// Action Outage.Handler#Index
	List *ActionOutageHandlerIndex
	// Action Outage.Handler#Show
	Show *ActionOutageHandlerShow
	// Action Outage.Handler#Show
	Find *ActionOutageHandlerShow
	// Action Outage.Handler#Create
	Create *ActionOutageHandlerCreate
	// Action Outage.Handler#Create
	New *ActionOutageHandlerCreate
	// Action Outage.Handler#Update
	Update *ActionOutageHandlerUpdate
	// Action Outage.Handler#Delete
	Delete *ActionOutageHandlerDelete
	// Action Outage.Handler#Delete
	Destroy *ActionOutageHandlerDelete
}

func NewResourceOutageHandler(client *Client) *ResourceOutageHandler {
	actionIndex := NewActionOutageHandlerIndex(client)
	actionShow := NewActionOutageHandlerShow(client)
	actionCreate := NewActionOutageHandlerCreate(client)
	actionUpdate := NewActionOutageHandlerUpdate(client)
	actionDelete := NewActionOutageHandlerDelete(client)

	return &ResourceOutageHandler{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Create: actionCreate,
		New: actionCreate,
		Update: actionUpdate,
		Delete: actionDelete,
		Destroy: actionDelete,
	}
}
