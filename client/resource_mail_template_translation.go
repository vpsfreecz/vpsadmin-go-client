package client

// Type for resource Mail_template.Translation
type ResourceMailTemplateTranslation struct {
	// Pointer to client
	Client *Client

	// Action Mail_template.Translation#Index
	Index *ActionMailTemplateTranslationIndex
	// Action Mail_template.Translation#Index
	List *ActionMailTemplateTranslationIndex
	// Action Mail_template.Translation#Show
	Show *ActionMailTemplateTranslationShow
	// Action Mail_template.Translation#Show
	Find *ActionMailTemplateTranslationShow
	// Action Mail_template.Translation#Create
	Create *ActionMailTemplateTranslationCreate
	// Action Mail_template.Translation#Create
	New *ActionMailTemplateTranslationCreate
	// Action Mail_template.Translation#Update
	Update *ActionMailTemplateTranslationUpdate
	// Action Mail_template.Translation#Delete
	Delete *ActionMailTemplateTranslationDelete
	// Action Mail_template.Translation#Delete
	Destroy *ActionMailTemplateTranslationDelete
}

func NewResourceMailTemplateTranslation(client *Client) *ResourceMailTemplateTranslation {
	actionIndex := NewActionMailTemplateTranslationIndex(client)
	actionShow := NewActionMailTemplateTranslationShow(client)
	actionCreate := NewActionMailTemplateTranslationCreate(client)
	actionUpdate := NewActionMailTemplateTranslationUpdate(client)
	actionDelete := NewActionMailTemplateTranslationDelete(client)

	return &ResourceMailTemplateTranslation{
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
