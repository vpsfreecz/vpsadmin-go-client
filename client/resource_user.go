package client

// Type for resource User
type ResourceUser struct {
	// Pointer to client
	Client *Client

	// Resource User.Environment_config
	EnvironmentConfig *ResourceUserEnvironmentConfig
	// Resource User.Cluster_resource
	ClusterResource *ResourceUserClusterResource
	// Resource User.Public_key
	PublicKey *ResourceUserPublicKey
	// Resource User.Mail_role_recipient
	MailRoleRecipient *ResourceUserMailRoleRecipient
	// Resource User.Mail_template_recipient
	MailTemplateRecipient *ResourceUserMailTemplateRecipient
	// Resource User.State_log
	StateLog *ResourceUserStateLog
	// Action User#Index
	Index *ActionUserIndex
	// Action User#Index
	List *ActionUserIndex
	// Action User#Create
	Create *ActionUserCreate
	// Action User#Create
	New *ActionUserCreate
	// Action User#Current
	Current *ActionUserCurrent
	// Action User#Touch
	Touch *ActionUserTouch
	// Action User#Show
	Show *ActionUserShow
	// Action User#Show
	Find *ActionUserShow
	// Action User#Update
	Update *ActionUserUpdate
	// Action User#Delete
	Delete *ActionUserDelete
	// Action User#Delete
	Destroy *ActionUserDelete
}

func NewResourceUser(client *Client) *ResourceUser {
	actionIndex := NewActionUserIndex(client)
	actionCreate := NewActionUserCreate(client)
	actionCurrent := NewActionUserCurrent(client)
	actionTouch := NewActionUserTouch(client)
	actionShow := NewActionUserShow(client)
	actionUpdate := NewActionUserUpdate(client)
	actionDelete := NewActionUserDelete(client)

	return &ResourceUser{
		Client: client,
		EnvironmentConfig: NewResourceUserEnvironmentConfig(client),
		ClusterResource: NewResourceUserClusterResource(client),
		PublicKey: NewResourceUserPublicKey(client),
		MailRoleRecipient: NewResourceUserMailRoleRecipient(client),
		MailTemplateRecipient: NewResourceUserMailTemplateRecipient(client),
		StateLog: NewResourceUserStateLog(client),
		Index: actionIndex,
		List: actionIndex,
		Create: actionCreate,
		New: actionCreate,
		Current: actionCurrent,
		Touch: actionTouch,
		Show: actionShow,
		Find: actionShow,
		Update: actionUpdate,
		Delete: actionDelete,
		Destroy: actionDelete,
	}
}
