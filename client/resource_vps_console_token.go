package client

// Type for resource Vps.Console_token
type ResourceVpsConsoleToken struct {
	// Pointer to client
	Client *Client

	// Action Vps.Console_token#Create
	Create *ActionVpsConsoleTokenCreate
	// Action Vps.Console_token#Create
	New *ActionVpsConsoleTokenCreate
	// Action Vps.Console_token#Show
	Show *ActionVpsConsoleTokenShow
	// Action Vps.Console_token#Show
	Find *ActionVpsConsoleTokenShow
	// Action Vps.Console_token#Delete
	Delete *ActionVpsConsoleTokenDelete
	// Action Vps.Console_token#Delete
	Destroy *ActionVpsConsoleTokenDelete
}

func NewResourceVpsConsoleToken(client *Client) *ResourceVpsConsoleToken {
	actionCreate := NewActionVpsConsoleTokenCreate(client)
	actionShow := NewActionVpsConsoleTokenShow(client)
	actionDelete := NewActionVpsConsoleTokenDelete(client)

	return &ResourceVpsConsoleToken{
		Client: client,
		Create: actionCreate,
		New: actionCreate,
		Show: actionShow,
		Find: actionShow,
		Delete: actionDelete,
		Destroy: actionDelete,
	}
}
