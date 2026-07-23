package client

// Type for resource Notification_template.Variant
type ResourceNotificationTemplateVariant struct {
	// Pointer to client
	Client *Client

	// Action Notification_template.Variant#Create
	Create *ActionNotificationTemplateVariantCreate
	// Action Notification_template.Variant#Create
	New *ActionNotificationTemplateVariantCreate
	// Action Notification_template.Variant#Delete
	Delete *ActionNotificationTemplateVariantDelete
	// Action Notification_template.Variant#Delete
	Destroy *ActionNotificationTemplateVariantDelete
	// Action Notification_template.Variant#Index
	Index *ActionNotificationTemplateVariantIndex
	// Action Notification_template.Variant#Index
	List *ActionNotificationTemplateVariantIndex
	// Action Notification_template.Variant#Show
	Show *ActionNotificationTemplateVariantShow
	// Action Notification_template.Variant#Show
	Find *ActionNotificationTemplateVariantShow
	// Action Notification_template.Variant#Update
	Update *ActionNotificationTemplateVariantUpdate
}

func NewResourceNotificationTemplateVariant(client *Client) *ResourceNotificationTemplateVariant {
	actionCreate := NewActionNotificationTemplateVariantCreate(client)
	actionDelete := NewActionNotificationTemplateVariantDelete(client)
	actionIndex := NewActionNotificationTemplateVariantIndex(client)
	actionShow := NewActionNotificationTemplateVariantShow(client)
	actionUpdate := NewActionNotificationTemplateVariantUpdate(client)

	return &ResourceNotificationTemplateVariant{
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
