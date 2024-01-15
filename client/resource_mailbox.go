package client

// Type for resource Mailbox
type ResourceMailbox struct {
	// Pointer to client
	Client *Client

	// Resource Mailbox.Handler
	Handler *ResourceMailboxHandler
	// Action Mailbox#Create
	Create *ActionMailboxCreate
	// Action Mailbox#Create
	New *ActionMailboxCreate
	// Action Mailbox#Delete
	Delete *ActionMailboxDelete
	// Action Mailbox#Delete
	Destroy *ActionMailboxDelete
	// Action Mailbox#Index
	Index *ActionMailboxIndex
	// Action Mailbox#Index
	List *ActionMailboxIndex
	// Action Mailbox#Show
	Show *ActionMailboxShow
	// Action Mailbox#Show
	Find *ActionMailboxShow
	// Action Mailbox#Update
	Update *ActionMailboxUpdate
}

func NewResourceMailbox(client *Client) *ResourceMailbox {
	actionCreate := NewActionMailboxCreate(client)
	actionDelete := NewActionMailboxDelete(client)
	actionIndex := NewActionMailboxIndex(client)
	actionShow := NewActionMailboxShow(client)
	actionUpdate := NewActionMailboxUpdate(client)

	return &ResourceMailbox{
		Client:  client,
		Handler: NewResourceMailboxHandler(client),
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
