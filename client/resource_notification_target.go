package client

// Type for resource Notification_target
type ResourceNotificationTarget struct {
	// Pointer to client
	Client *Client

	// Action Notification_target#Confirm_email_verification
	ConfirmEmailVerification *ActionNotificationTargetConfirmEmailVerification
	// Action Notification_target#Confirm_sms_verification_code
	ConfirmSmsVerificationCode *ActionNotificationTargetConfirmSmsVerificationCode
	// Action Notification_target#Create
	Create *ActionNotificationTargetCreate
	// Action Notification_target#Create
	New *ActionNotificationTargetCreate
	// Action Notification_target#Create_pairing_token
	CreatePairingToken *ActionNotificationTargetCreatePairingToken
	// Action Notification_target#Delete
	Delete *ActionNotificationTargetDelete
	// Action Notification_target#Delete
	Destroy *ActionNotificationTargetDelete
	// Action Notification_target#Index
	Index *ActionNotificationTargetIndex
	// Action Notification_target#Index
	List *ActionNotificationTargetIndex
	// Action Notification_target#Send_email_verification
	SendEmailVerification *ActionNotificationTargetSendEmailVerification
	// Action Notification_target#Send_sms_verification_code
	SendSmsVerificationCode *ActionNotificationTargetSendSmsVerificationCode
	// Action Notification_target#Show
	Show *ActionNotificationTargetShow
	// Action Notification_target#Show
	Find *ActionNotificationTargetShow
	// Action Notification_target#Update
	Update *ActionNotificationTargetUpdate
}

func NewResourceNotificationTarget(client *Client) *ResourceNotificationTarget {
	actionConfirmEmailVerification := NewActionNotificationTargetConfirmEmailVerification(client)
	actionConfirmSmsVerificationCode := NewActionNotificationTargetConfirmSmsVerificationCode(client)
	actionCreate := NewActionNotificationTargetCreate(client)
	actionCreatePairingToken := NewActionNotificationTargetCreatePairingToken(client)
	actionDelete := NewActionNotificationTargetDelete(client)
	actionIndex := NewActionNotificationTargetIndex(client)
	actionSendEmailVerification := NewActionNotificationTargetSendEmailVerification(client)
	actionSendSmsVerificationCode := NewActionNotificationTargetSendSmsVerificationCode(client)
	actionShow := NewActionNotificationTargetShow(client)
	actionUpdate := NewActionNotificationTargetUpdate(client)

	return &ResourceNotificationTarget{
		Client:                     client,
		ConfirmEmailVerification:   actionConfirmEmailVerification,
		ConfirmSmsVerificationCode: actionConfirmSmsVerificationCode,
		Create:                     actionCreate,
		New:                        actionCreate,
		CreatePairingToken:         actionCreatePairingToken,
		Delete:                     actionDelete,
		Destroy:                    actionDelete,
		Index:                      actionIndex,
		List:                       actionIndex,
		SendEmailVerification:      actionSendEmailVerification,
		SendSmsVerificationCode:    actionSendSmsVerificationCode,
		Show:                       actionShow,
		Find:                       actionShow,
		Update:                     actionUpdate,
	}
}
