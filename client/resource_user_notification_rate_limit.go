package client

// Type for resource User.Notification_rate_limit
type ResourceUserNotificationRateLimit struct {
	// Pointer to client
	Client *Client

	// Action User.Notification_rate_limit#Index
	Index *ActionUserNotificationRateLimitIndex
	// Action User.Notification_rate_limit#Index
	List *ActionUserNotificationRateLimitIndex
	// Action User.Notification_rate_limit#Show
	Show *ActionUserNotificationRateLimitShow
	// Action User.Notification_rate_limit#Show
	Find *ActionUserNotificationRateLimitShow
	// Action User.Notification_rate_limit#Update
	Update *ActionUserNotificationRateLimitUpdate
}

func NewResourceUserNotificationRateLimit(client *Client) *ResourceUserNotificationRateLimit {
	actionIndex := NewActionUserNotificationRateLimitIndex(client)
	actionShow := NewActionUserNotificationRateLimitShow(client)
	actionUpdate := NewActionUserNotificationRateLimitUpdate(client)

	return &ResourceUserNotificationRateLimit{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
		Update: actionUpdate,
	}
}
