package client

// Type for resource Mailbox.Handler
type ResourceMailboxHandler struct {
	// Pointer to client
	Client *Client

	// Action Mailbox.Handler#Create
	Create *ActionMailboxHandlerCreate
	// Action Mailbox.Handler#Create
	New *ActionMailboxHandlerCreate
	// Action Mailbox.Handler#Delete
	Delete *ActionMailboxHandlerDelete
	// Action Mailbox.Handler#Delete
	Destroy *ActionMailboxHandlerDelete
	// Action Mailbox.Handler#Index
	Index *ActionMailboxHandlerIndex
	// Action Mailbox.Handler#Index
	List *ActionMailboxHandlerIndex
	// Action Mailbox.Handler#Show
	Show *ActionMailboxHandlerShow
	// Action Mailbox.Handler#Show
	Find *ActionMailboxHandlerShow
	// Action Mailbox.Handler#Update
	Update *ActionMailboxHandlerUpdate
}

func NewResourceMailboxHandler(client *Client) *ResourceMailboxHandler {
	actionCreate := NewActionMailboxHandlerCreate(client)
	actionDelete := NewActionMailboxHandlerDelete(client)
	actionIndex := NewActionMailboxHandlerIndex(client)
	actionShow := NewActionMailboxHandlerShow(client)
	actionUpdate := NewActionMailboxHandlerUpdate(client)

	return &ResourceMailboxHandler{
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
