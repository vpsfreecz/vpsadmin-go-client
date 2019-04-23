package client

// Type for resource Os_template
type ResourceOsTemplate struct {
	// Pointer to client
	Client *Client

	// Action Os_template#Index
	Index *ActionOsTemplateIndex
	// Action Os_template#Index
	List *ActionOsTemplateIndex
	// Action Os_template#Show
	Show *ActionOsTemplateShow
	// Action Os_template#Show
	Find *ActionOsTemplateShow
	// Action Os_template#Create
	Create *ActionOsTemplateCreate
	// Action Os_template#Create
	New *ActionOsTemplateCreate
	// Action Os_template#Update
	Update *ActionOsTemplateUpdate
	// Action Os_template#Delete
	Delete *ActionOsTemplateDelete
	// Action Os_template#Delete
	Destroy *ActionOsTemplateDelete
}

func NewResourceOsTemplate(client *Client) *ResourceOsTemplate {
	actionIndex := NewActionOsTemplateIndex(client)
	actionShow := NewActionOsTemplateShow(client)
	actionCreate := NewActionOsTemplateCreate(client)
	actionUpdate := NewActionOsTemplateUpdate(client)
	actionDelete := NewActionOsTemplateDelete(client)

	return &ResourceOsTemplate{
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
