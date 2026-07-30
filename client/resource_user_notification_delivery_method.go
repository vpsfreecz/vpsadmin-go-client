package client

// Type for resource User.Notification_delivery_method
type ResourceUserNotificationDeliveryMethod struct {
	// Pointer to client
	Client *Client

	// Action User.Notification_delivery_method#Index
	Index *ActionUserNotificationDeliveryMethodIndex
	// Action User.Notification_delivery_method#Index
	List *ActionUserNotificationDeliveryMethodIndex
	// Action User.Notification_delivery_method#Show
	Show *ActionUserNotificationDeliveryMethodShow
	// Action User.Notification_delivery_method#Show
	Find *ActionUserNotificationDeliveryMethodShow
	// Action User.Notification_delivery_method#Update
	Update *ActionUserNotificationDeliveryMethodUpdate
}

func NewResourceUserNotificationDeliveryMethod(client *Client) *ResourceUserNotificationDeliveryMethod {
	actionIndex := NewActionUserNotificationDeliveryMethodIndex(client)
	actionShow := NewActionUserNotificationDeliveryMethodShow(client)
	actionUpdate := NewActionUserNotificationDeliveryMethodUpdate(client)

	return &ResourceUserNotificationDeliveryMethod{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
		Update: actionUpdate,
	}
}
