package client

// Type for resource Notification_receiver.Target
type ResourceNotificationReceiverTarget struct {
	// Pointer to client
	Client *Client

	// Action Notification_receiver.Target#Create
	Create *ActionNotificationReceiverTargetCreate
	// Action Notification_receiver.Target#Create
	New *ActionNotificationReceiverTargetCreate
	// Action Notification_receiver.Target#Delete
	Delete *ActionNotificationReceiverTargetDelete
	// Action Notification_receiver.Target#Delete
	Destroy *ActionNotificationReceiverTargetDelete
	// Action Notification_receiver.Target#Index
	Index *ActionNotificationReceiverTargetIndex
	// Action Notification_receiver.Target#Index
	List *ActionNotificationReceiverTargetIndex
	// Action Notification_receiver.Target#Show
	Show *ActionNotificationReceiverTargetShow
	// Action Notification_receiver.Target#Show
	Find *ActionNotificationReceiverTargetShow
	// Action Notification_receiver.Target#Update
	Update *ActionNotificationReceiverTargetUpdate
}

func NewResourceNotificationReceiverTarget(client *Client) *ResourceNotificationReceiverTarget {
	actionCreate := NewActionNotificationReceiverTargetCreate(client)
	actionDelete := NewActionNotificationReceiverTargetDelete(client)
	actionIndex := NewActionNotificationReceiverTargetIndex(client)
	actionShow := NewActionNotificationReceiverTargetShow(client)
	actionUpdate := NewActionNotificationReceiverTargetUpdate(client)

	return &ResourceNotificationReceiverTarget{
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
