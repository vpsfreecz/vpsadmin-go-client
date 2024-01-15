package client

// Type for resource Oauth2_client
type ResourceOauth2Client struct {
	// Pointer to client
	Client *Client

	// Action Oauth2_client#Create
	Create *ActionOauth2ClientCreate
	// Action Oauth2_client#Create
	New *ActionOauth2ClientCreate
	// Action Oauth2_client#Delete
	Delete *ActionOauth2ClientDelete
	// Action Oauth2_client#Delete
	Destroy *ActionOauth2ClientDelete
	// Action Oauth2_client#Index
	Index *ActionOauth2ClientIndex
	// Action Oauth2_client#Index
	List *ActionOauth2ClientIndex
	// Action Oauth2_client#Show
	Show *ActionOauth2ClientShow
	// Action Oauth2_client#Show
	Find *ActionOauth2ClientShow
	// Action Oauth2_client#Update
	Update *ActionOauth2ClientUpdate
}

func NewResourceOauth2Client(client *Client) *ResourceOauth2Client {
	actionCreate := NewActionOauth2ClientCreate(client)
	actionDelete := NewActionOauth2ClientDelete(client)
	actionIndex := NewActionOauth2ClientIndex(client)
	actionShow := NewActionOauth2ClientShow(client)
	actionUpdate := NewActionOauth2ClientUpdate(client)

	return &ResourceOauth2Client{
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
