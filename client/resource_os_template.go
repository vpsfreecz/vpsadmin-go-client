package client

// Type for resource Os_template
type ResourceOsTemplate struct {
	// Pointer to client
	Client *Client

	// Action Os_template#Create
	Create *ActionOsTemplateCreate
	// Action Os_template#Create
	New *ActionOsTemplateCreate
	// Action Os_template#Delete
	Delete *ActionOsTemplateDelete
	// Action Os_template#Delete
	Destroy *ActionOsTemplateDelete
	// Action Os_template#Index
	Index *ActionOsTemplateIndex
	// Action Os_template#Index
	List *ActionOsTemplateIndex
	// Action Os_template#Show
	Show *ActionOsTemplateShow
	// Action Os_template#Show
	Find *ActionOsTemplateShow
	// Action Os_template#Update
	Update *ActionOsTemplateUpdate
}

func NewResourceOsTemplate(client *Client) *ResourceOsTemplate {
	actionCreate := NewActionOsTemplateCreate(client)
	actionDelete := NewActionOsTemplateDelete(client)
	actionIndex := NewActionOsTemplateIndex(client)
	actionShow := NewActionOsTemplateShow(client)
	actionUpdate := NewActionOsTemplateUpdate(client)

	return &ResourceOsTemplate{
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
