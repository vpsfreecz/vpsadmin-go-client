package client

// Type for resource Mail_template.Recipient
type ResourceMailTemplateRecipient struct {
	// Pointer to client
	Client *Client

	// Action Mail_template.Recipient#Create
	Create *ActionMailTemplateRecipientCreate
	// Action Mail_template.Recipient#Create
	New *ActionMailTemplateRecipientCreate
	// Action Mail_template.Recipient#Delete
	Delete *ActionMailTemplateRecipientDelete
	// Action Mail_template.Recipient#Delete
	Destroy *ActionMailTemplateRecipientDelete
	// Action Mail_template.Recipient#Index
	Index *ActionMailTemplateRecipientIndex
	// Action Mail_template.Recipient#Index
	List *ActionMailTemplateRecipientIndex
	// Action Mail_template.Recipient#Show
	Show *ActionMailTemplateRecipientShow
	// Action Mail_template.Recipient#Show
	Find *ActionMailTemplateRecipientShow
}

func NewResourceMailTemplateRecipient(client *Client) *ResourceMailTemplateRecipient {
	actionCreate := NewActionMailTemplateRecipientCreate(client)
	actionDelete := NewActionMailTemplateRecipientDelete(client)
	actionIndex := NewActionMailTemplateRecipientIndex(client)
	actionShow := NewActionMailTemplateRecipientShow(client)

	return &ResourceMailTemplateRecipient{
		Client:  client,
		Create:  actionCreate,
		New:     actionCreate,
		Delete:  actionDelete,
		Destroy: actionDelete,
		Index:   actionIndex,
		List:    actionIndex,
		Show:    actionShow,
		Find:    actionShow,
	}
}
