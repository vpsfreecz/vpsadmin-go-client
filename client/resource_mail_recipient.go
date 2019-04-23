package client

// Type for resource Mail_recipient
type ResourceMailRecipient struct {
	// Pointer to client
	Client *Client

	// Action Mail_recipient#Index
	Index *ActionMailRecipientIndex
	// Action Mail_recipient#Index
	List *ActionMailRecipientIndex
	// Action Mail_recipient#Show
	Show *ActionMailRecipientShow
	// Action Mail_recipient#Show
	Find *ActionMailRecipientShow
	// Action Mail_recipient#Create
	Create *ActionMailRecipientCreate
	// Action Mail_recipient#Create
	New *ActionMailRecipientCreate
	// Action Mail_recipient#Update
	Update *ActionMailRecipientUpdate
	// Action Mail_recipient#Delete
	Delete *ActionMailRecipientDelete
	// Action Mail_recipient#Delete
	Destroy *ActionMailRecipientDelete
}

func NewResourceMailRecipient(client *Client) *ResourceMailRecipient {
	actionIndex := NewActionMailRecipientIndex(client)
	actionShow := NewActionMailRecipientShow(client)
	actionCreate := NewActionMailRecipientCreate(client)
	actionUpdate := NewActionMailRecipientUpdate(client)
	actionDelete := NewActionMailRecipientDelete(client)

	return &ResourceMailRecipient{
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
