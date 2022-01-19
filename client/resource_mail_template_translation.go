package client

// Type for resource Mail_template.Translation
type ResourceMailTemplateTranslation struct {
	// Pointer to client
	Client *Client

	// Action Mail_template.Translation#Create
	Create *ActionMailTemplateTranslationCreate
	// Action Mail_template.Translation#Create
	New *ActionMailTemplateTranslationCreate
	// Action Mail_template.Translation#Delete
	Delete *ActionMailTemplateTranslationDelete
	// Action Mail_template.Translation#Delete
	Destroy *ActionMailTemplateTranslationDelete
	// Action Mail_template.Translation#Index
	Index *ActionMailTemplateTranslationIndex
	// Action Mail_template.Translation#Index
	List *ActionMailTemplateTranslationIndex
	// Action Mail_template.Translation#Show
	Show *ActionMailTemplateTranslationShow
	// Action Mail_template.Translation#Show
	Find *ActionMailTemplateTranslationShow
	// Action Mail_template.Translation#Update
	Update *ActionMailTemplateTranslationUpdate
}

func NewResourceMailTemplateTranslation(client *Client) *ResourceMailTemplateTranslation {
	actionCreate := NewActionMailTemplateTranslationCreate(client)
	actionDelete := NewActionMailTemplateTranslationDelete(client)
	actionIndex := NewActionMailTemplateTranslationIndex(client)
	actionShow := NewActionMailTemplateTranslationShow(client)
	actionUpdate := NewActionMailTemplateTranslationUpdate(client)

	return &ResourceMailTemplateTranslation{
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
