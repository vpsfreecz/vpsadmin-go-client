package client

// Type for resource User
type ResourceUser struct {
	// Pointer to client
	Client *Client

	// Resource User.Cluster_resource
	ClusterResource *ResourceUserClusterResource
	// Resource User.Environment_config
	EnvironmentConfig *ResourceUserEnvironmentConfig
	// Resource User.Known_device
	KnownDevice *ResourceUserKnownDevice
	// Resource User.Notification_delivery_method
	NotificationDeliveryMethod *ResourceUserNotificationDeliveryMethod
	// Resource User.Notification_rate_limit
	NotificationRateLimit *ResourceUserNotificationRateLimit
	// Resource User.Public_key
	PublicKey *ResourceUserPublicKey
	// Resource User.State_log
	StateLog *ResourceUserStateLog
	// Resource User.Totp_device
	TotpDevice *ResourceUserTotpDevice
	// Resource User.Webauthn_credential
	WebauthnCredential *ResourceUserWebauthnCredential
	// Action User#Available_ips
	AvailableIps *ActionUserAvailableIps
	// Action User#Create
	Create *ActionUserCreate
	// Action User#Create
	New *ActionUserCreate
	// Action User#Current
	Current *ActionUserCurrent
	// Action User#Delete
	Delete *ActionUserDelete
	// Action User#Delete
	Destroy *ActionUserDelete
	// Action User#Get_payment_instructions
	GetPaymentInstructions *ActionUserGetPaymentInstructions
	// Action User#Index
	Index *ActionUserIndex
	// Action User#Index
	List *ActionUserIndex
	// Action User#Show
	Show *ActionUserShow
	// Action User#Show
	Find *ActionUserShow
	// Action User#Touch
	Touch *ActionUserTouch
	// Action User#Update
	Update *ActionUserUpdate
}

func NewResourceUser(client *Client) *ResourceUser {
	actionAvailableIps := NewActionUserAvailableIps(client)
	actionCreate := NewActionUserCreate(client)
	actionCurrent := NewActionUserCurrent(client)
	actionDelete := NewActionUserDelete(client)
	actionGetPaymentInstructions := NewActionUserGetPaymentInstructions(client)
	actionIndex := NewActionUserIndex(client)
	actionShow := NewActionUserShow(client)
	actionTouch := NewActionUserTouch(client)
	actionUpdate := NewActionUserUpdate(client)

	return &ResourceUser{
		Client:                     client,
		ClusterResource:            NewResourceUserClusterResource(client),
		EnvironmentConfig:          NewResourceUserEnvironmentConfig(client),
		KnownDevice:                NewResourceUserKnownDevice(client),
		NotificationDeliveryMethod: NewResourceUserNotificationDeliveryMethod(client),
		NotificationRateLimit:      NewResourceUserNotificationRateLimit(client),
		PublicKey:                  NewResourceUserPublicKey(client),
		StateLog:                   NewResourceUserStateLog(client),
		TotpDevice:                 NewResourceUserTotpDevice(client),
		WebauthnCredential:         NewResourceUserWebauthnCredential(client),
		AvailableIps:               actionAvailableIps,
		Create:                     actionCreate,
		New:                        actionCreate,
		Current:                    actionCurrent,
		Delete:                     actionDelete,
		Destroy:                    actionDelete,
		GetPaymentInstructions:     actionGetPaymentInstructions,
		Index:                      actionIndex,
		List:                       actionIndex,
		Show:                       actionShow,
		Find:                       actionShow,
		Touch:                      actionTouch,
		Update:                     actionUpdate,
	}
}
