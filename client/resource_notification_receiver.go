package client

// Type for resource Notification_receiver
type ResourceNotificationReceiver struct {
	// Pointer to client
	Client *Client

	// Resource Notification_receiver.Target
	Target *ResourceNotificationReceiverTarget
	// Action Notification_receiver#Create
	Create *ActionNotificationReceiverCreate
	// Action Notification_receiver#Create
	New *ActionNotificationReceiverCreate
	// Action Notification_receiver#Delete
	Delete *ActionNotificationReceiverDelete
	// Action Notification_receiver#Delete
	Destroy *ActionNotificationReceiverDelete
	// Action Notification_receiver#Index
	Index *ActionNotificationReceiverIndex
	// Action Notification_receiver#Index
	List *ActionNotificationReceiverIndex
	// Action Notification_receiver#Show
	Show *ActionNotificationReceiverShow
	// Action Notification_receiver#Show
	Find *ActionNotificationReceiverShow
	// Action Notification_receiver#Update
	Update *ActionNotificationReceiverUpdate
}

func NewResourceNotificationReceiver(client *Client) *ResourceNotificationReceiver {
	actionCreate := NewActionNotificationReceiverCreate(client)
	actionDelete := NewActionNotificationReceiverDelete(client)
	actionIndex := NewActionNotificationReceiverIndex(client)
	actionShow := NewActionNotificationReceiverShow(client)
	actionUpdate := NewActionNotificationReceiverUpdate(client)

	return &ResourceNotificationReceiver{
		Client:  client,
		Target:  NewResourceNotificationReceiverTarget(client),
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
