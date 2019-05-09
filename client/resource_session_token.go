package client

// Type for resource Session_token
type ResourceSessionToken struct {
	// Pointer to client
	Client *Client

	// Action Session_token#Index
	Index *ActionSessionTokenIndex
	// Action Session_token#Index
	List *ActionSessionTokenIndex
	// Action Session_token#Create
	Create *ActionSessionTokenCreate
	// Action Session_token#Create
	New *ActionSessionTokenCreate
	// Action Session_token#Show
	Show *ActionSessionTokenShow
	// Action Session_token#Show
	Find *ActionSessionTokenShow
	// Action Session_token#Update
	Update *ActionSessionTokenUpdate
	// Action Session_token#Delete
	Delete *ActionSessionTokenDelete
	// Action Session_token#Delete
	Destroy *ActionSessionTokenDelete
}

func NewResourceSessionToken(client *Client) *ResourceSessionToken {
	actionIndex := NewActionSessionTokenIndex(client)
	actionCreate := NewActionSessionTokenCreate(client)
	actionShow := NewActionSessionTokenShow(client)
	actionUpdate := NewActionSessionTokenUpdate(client)
	actionDelete := NewActionSessionTokenDelete(client)

	return &ResourceSessionToken{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Create: actionCreate,
		New: actionCreate,
		Show: actionShow,
		Find: actionShow,
		Update: actionUpdate,
		Delete: actionDelete,
		Destroy: actionDelete,
	}
}
