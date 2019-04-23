package client

// Type for resource Auth_token
type ResourceAuthToken struct {
	// Pointer to client
	Client *Client

	// Action Auth_token#Index
	Index *ActionAuthTokenIndex
	// Action Auth_token#Index
	List *ActionAuthTokenIndex
	// Action Auth_token#Create
	Create *ActionAuthTokenCreate
	// Action Auth_token#Create
	New *ActionAuthTokenCreate
	// Action Auth_token#Show
	Show *ActionAuthTokenShow
	// Action Auth_token#Show
	Find *ActionAuthTokenShow
	// Action Auth_token#Update
	Update *ActionAuthTokenUpdate
	// Action Auth_token#Delete
	Delete *ActionAuthTokenDelete
	// Action Auth_token#Delete
	Destroy *ActionAuthTokenDelete
}

func NewResourceAuthToken(client *Client) *ResourceAuthToken {
	actionIndex := NewActionAuthTokenIndex(client)
	actionCreate := NewActionAuthTokenCreate(client)
	actionShow := NewActionAuthTokenShow(client)
	actionUpdate := NewActionAuthTokenUpdate(client)
	actionDelete := NewActionAuthTokenDelete(client)

	return &ResourceAuthToken{
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
