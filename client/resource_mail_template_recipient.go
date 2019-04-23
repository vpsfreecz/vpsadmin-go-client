package client

// Type for resource Mail_template.Recipient
type ResourceMailTemplateRecipient struct {
	// Pointer to client
	Client *Client

	// Action Mail_template.Recipient#Index
	Index *ActionMailTemplateRecipientIndex
	// Action Mail_template.Recipient#Index
	List *ActionMailTemplateRecipientIndex
	// Action Mail_template.Recipient#Show
	Show *ActionMailTemplateRecipientShow
	// Action Mail_template.Recipient#Show
	Find *ActionMailTemplateRecipientShow
	// Action Mail_template.Recipient#Create
	Create *ActionMailTemplateRecipientCreate
	// Action Mail_template.Recipient#Create
	New *ActionMailTemplateRecipientCreate
	// Action Mail_template.Recipient#Delete
	Delete *ActionMailTemplateRecipientDelete
	// Action Mail_template.Recipient#Delete
	Destroy *ActionMailTemplateRecipientDelete
}

func NewResourceMailTemplateRecipient(client *Client) *ResourceMailTemplateRecipient {
	actionIndex := NewActionMailTemplateRecipientIndex(client)
	actionShow := NewActionMailTemplateRecipientShow(client)
	actionCreate := NewActionMailTemplateRecipientCreate(client)
	actionDelete := NewActionMailTemplateRecipientDelete(client)

	return &ResourceMailTemplateRecipient{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Create: actionCreate,
		New: actionCreate,
		Delete: actionDelete,
		Destroy: actionDelete,
	}
}
