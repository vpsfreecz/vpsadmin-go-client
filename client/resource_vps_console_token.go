package client

// Type for resource Vps.Console_token
type ResourceVpsConsoleToken struct {
	// Pointer to client
	Client *Client

	// Action Vps.Console_token#Create
	Create *ActionVpsConsoleTokenCreate
	// Action Vps.Console_token#Create
	New *ActionVpsConsoleTokenCreate
	// Action Vps.Console_token#Delete
	Delete *ActionVpsConsoleTokenDelete
	// Action Vps.Console_token#Delete
	Destroy *ActionVpsConsoleTokenDelete
	// Action Vps.Console_token#Show
	Show *ActionVpsConsoleTokenShow
	// Action Vps.Console_token#Show
	Find *ActionVpsConsoleTokenShow
}

func NewResourceVpsConsoleToken(client *Client) *ResourceVpsConsoleToken {
	actionCreate := NewActionVpsConsoleTokenCreate(client)
	actionDelete := NewActionVpsConsoleTokenDelete(client)
	actionShow := NewActionVpsConsoleTokenShow(client)

	return &ResourceVpsConsoleToken{
		Client: client,
		Create: actionCreate,
		New: actionCreate,
		Delete: actionDelete,
		Destroy: actionDelete,
		Show: actionShow,
		Find: actionShow,
	}
}
