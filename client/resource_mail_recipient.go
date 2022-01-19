package client

// Type for resource Mail_recipient
type ResourceMailRecipient struct {
	// Pointer to client
	Client *Client

	// Action Mail_recipient#Create
	Create *ActionMailRecipientCreate
	// Action Mail_recipient#Create
	New *ActionMailRecipientCreate
	// Action Mail_recipient#Delete
	Delete *ActionMailRecipientDelete
	// Action Mail_recipient#Delete
	Destroy *ActionMailRecipientDelete
	// Action Mail_recipient#Index
	Index *ActionMailRecipientIndex
	// Action Mail_recipient#Index
	List *ActionMailRecipientIndex
	// Action Mail_recipient#Show
	Show *ActionMailRecipientShow
	// Action Mail_recipient#Show
	Find *ActionMailRecipientShow
	// Action Mail_recipient#Update
	Update *ActionMailRecipientUpdate
}

func NewResourceMailRecipient(client *Client) *ResourceMailRecipient {
	actionCreate := NewActionMailRecipientCreate(client)
	actionDelete := NewActionMailRecipientDelete(client)
	actionIndex := NewActionMailRecipientIndex(client)
	actionShow := NewActionMailRecipientShow(client)
	actionUpdate := NewActionMailRecipientUpdate(client)

	return &ResourceMailRecipient{
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
