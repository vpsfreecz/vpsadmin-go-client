package client

// Type for resource Mail_template
type ResourceMailTemplate struct {
	// Pointer to client
	Client *Client

	// Resource Mail_template.Recipient
	Recipient *ResourceMailTemplateRecipient
	// Resource Mail_template.Translation
	Translation *ResourceMailTemplateTranslation
	// Action Mail_template#Index
	Index *ActionMailTemplateIndex
	// Action Mail_template#Index
	List *ActionMailTemplateIndex
	// Action Mail_template#Show
	Show *ActionMailTemplateShow
	// Action Mail_template#Show
	Find *ActionMailTemplateShow
	// Action Mail_template#Create
	Create *ActionMailTemplateCreate
	// Action Mail_template#Create
	New *ActionMailTemplateCreate
	// Action Mail_template#Update
	Update *ActionMailTemplateUpdate
	// Action Mail_template#Delete
	Delete *ActionMailTemplateDelete
	// Action Mail_template#Delete
	Destroy *ActionMailTemplateDelete
}

func NewResourceMailTemplate(client *Client) *ResourceMailTemplate {
	actionIndex := NewActionMailTemplateIndex(client)
	actionShow := NewActionMailTemplateShow(client)
	actionCreate := NewActionMailTemplateCreate(client)
	actionUpdate := NewActionMailTemplateUpdate(client)
	actionDelete := NewActionMailTemplateDelete(client)

	return &ResourceMailTemplate{
		Client: client,
		Recipient: NewResourceMailTemplateRecipient(client),
		Translation: NewResourceMailTemplateTranslation(client),
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
