package client

// Type for resource Notification_template
type ResourceNotificationTemplate struct {
	// Pointer to client
	Client *Client

	// Resource Notification_template.Variant
	Variant *ResourceNotificationTemplateVariant
	// Action Notification_template#Create
	Create *ActionNotificationTemplateCreate
	// Action Notification_template#Create
	New *ActionNotificationTemplateCreate
	// Action Notification_template#Delete
	Delete *ActionNotificationTemplateDelete
	// Action Notification_template#Delete
	Destroy *ActionNotificationTemplateDelete
	// Action Notification_template#Index
	Index *ActionNotificationTemplateIndex
	// Action Notification_template#Index
	List *ActionNotificationTemplateIndex
	// Action Notification_template#Show
	Show *ActionNotificationTemplateShow
	// Action Notification_template#Show
	Find *ActionNotificationTemplateShow
	// Action Notification_template#Update
	Update *ActionNotificationTemplateUpdate
}

func NewResourceNotificationTemplate(client *Client) *ResourceNotificationTemplate {
	actionCreate := NewActionNotificationTemplateCreate(client)
	actionDelete := NewActionNotificationTemplateDelete(client)
	actionIndex := NewActionNotificationTemplateIndex(client)
	actionShow := NewActionNotificationTemplateShow(client)
	actionUpdate := NewActionNotificationTemplateUpdate(client)

	return &ResourceNotificationTemplate{
		Client:  client,
		Variant: NewResourceNotificationTemplateVariant(client),
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
