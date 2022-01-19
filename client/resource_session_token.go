package client

// Type for resource Session_token
type ResourceSessionToken struct {
	// Pointer to client
	Client *Client

	// Action Session_token#Create
	Create *ActionSessionTokenCreate
	// Action Session_token#Create
	New *ActionSessionTokenCreate
	// Action Session_token#Delete
	Delete *ActionSessionTokenDelete
	// Action Session_token#Delete
	Destroy *ActionSessionTokenDelete
	// Action Session_token#Index
	Index *ActionSessionTokenIndex
	// Action Session_token#Index
	List *ActionSessionTokenIndex
	// Action Session_token#Show
	Show *ActionSessionTokenShow
	// Action Session_token#Show
	Find *ActionSessionTokenShow
	// Action Session_token#Update
	Update *ActionSessionTokenUpdate
}

func NewResourceSessionToken(client *Client) *ResourceSessionToken {
	actionCreate := NewActionSessionTokenCreate(client)
	actionDelete := NewActionSessionTokenDelete(client)
	actionIndex := NewActionSessionTokenIndex(client)
	actionShow := NewActionSessionTokenShow(client)
	actionUpdate := NewActionSessionTokenUpdate(client)

	return &ResourceSessionToken{
		Client: client,
		Create: actionCreate,
		New: actionCreate,
		Delete: actionDelete,
		Destroy: actionDelete,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Update: actionUpdate,
	}
}
