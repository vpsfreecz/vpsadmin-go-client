package client

// Type for resource Mail_template
type ResourceMailTemplate struct {
	// Pointer to client
	Client *Client

	// Resource Mail_template.Recipient
	Recipient *ResourceMailTemplateRecipient
	// Resource Mail_template.Translation
	Translation *ResourceMailTemplateTranslation
	// Action Mail_template#Create
	Create *ActionMailTemplateCreate
	// Action Mail_template#Create
	New *ActionMailTemplateCreate
	// Action Mail_template#Delete
	Delete *ActionMailTemplateDelete
	// Action Mail_template#Delete
	Destroy *ActionMailTemplateDelete
	// Action Mail_template#Index
	Index *ActionMailTemplateIndex
	// Action Mail_template#Index
	List *ActionMailTemplateIndex
	// Action Mail_template#Show
	Show *ActionMailTemplateShow
	// Action Mail_template#Show
	Find *ActionMailTemplateShow
	// Action Mail_template#Update
	Update *ActionMailTemplateUpdate
}

func NewResourceMailTemplate(client *Client) *ResourceMailTemplate {
	actionCreate := NewActionMailTemplateCreate(client)
	actionDelete := NewActionMailTemplateDelete(client)
	actionIndex := NewActionMailTemplateIndex(client)
	actionShow := NewActionMailTemplateShow(client)
	actionUpdate := NewActionMailTemplateUpdate(client)

	return &ResourceMailTemplate{
		Client:      client,
		Recipient:   NewResourceMailTemplateRecipient(client),
		Translation: NewResourceMailTemplateTranslation(client),
		Create:      actionCreate,
		New:         actionCreate,
		Delete:      actionDelete,
		Destroy:     actionDelete,
		Index:       actionIndex,
		List:        actionIndex,
		Show:        actionShow,
		Find:        actionShow,
		Update:      actionUpdate,
	}
}
