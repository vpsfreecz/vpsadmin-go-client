package client

// Type for resource User
type ResourceUser struct {
	// Pointer to client
	Client *Client

	// Resource User.Cluster_resource
	ClusterResource *ResourceUserClusterResource
	// Resource User.Environment_config
	EnvironmentConfig *ResourceUserEnvironmentConfig
	// Resource User.Mail_role_recipient
	MailRoleRecipient *ResourceUserMailRoleRecipient
	// Resource User.Mail_template_recipient
	MailTemplateRecipient *ResourceUserMailTemplateRecipient
	// Resource User.Public_key
	PublicKey *ResourceUserPublicKey
	// Resource User.State_log
	StateLog *ResourceUserStateLog
	// Resource User.Totp_device
	TotpDevice *ResourceUserTotpDevice
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
	actionCreate := NewActionUserCreate(client)
	actionCurrent := NewActionUserCurrent(client)
	actionDelete := NewActionUserDelete(client)
	actionGetPaymentInstructions := NewActionUserGetPaymentInstructions(client)
	actionIndex := NewActionUserIndex(client)
	actionShow := NewActionUserShow(client)
	actionTouch := NewActionUserTouch(client)
	actionUpdate := NewActionUserUpdate(client)

	return &ResourceUser{
		Client:                 client,
		ClusterResource:        NewResourceUserClusterResource(client),
		EnvironmentConfig:      NewResourceUserEnvironmentConfig(client),
		MailRoleRecipient:      NewResourceUserMailRoleRecipient(client),
		MailTemplateRecipient:  NewResourceUserMailTemplateRecipient(client),
		PublicKey:              NewResourceUserPublicKey(client),
		StateLog:               NewResourceUserStateLog(client),
		TotpDevice:             NewResourceUserTotpDevice(client),
		Create:                 actionCreate,
		New:                    actionCreate,
		Current:                actionCurrent,
		Delete:                 actionDelete,
		Destroy:                actionDelete,
		GetPaymentInstructions: actionGetPaymentInstructions,
		Index:                  actionIndex,
		List:                   actionIndex,
		Show:                   actionShow,
		Find:                   actionShow,
		Touch:                  actionTouch,
		Update:                 actionUpdate,
	}
}
