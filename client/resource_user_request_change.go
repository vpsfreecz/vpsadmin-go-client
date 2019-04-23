package client

// Type for resource User_request.Change
type ResourceUserRequestChange struct {
	// Pointer to client
	Client *Client

	// Action User_request.Change#Index
	Index *ActionUserRequestChangeIndex
	// Action User_request.Change#Index
	List *ActionUserRequestChangeIndex
	// Action User_request.Change#Show
	Show *ActionUserRequestChangeShow
	// Action User_request.Change#Show
	Find *ActionUserRequestChangeShow
	// Action User_request.Change#Create
	Create *ActionUserRequestChangeCreate
	// Action User_request.Change#Create
	New *ActionUserRequestChangeCreate
	// Action User_request.Change#Resolve
	Resolve *ActionUserRequestChangeResolve
}

func NewResourceUserRequestChange(client *Client) *ResourceUserRequestChange {
	actionIndex := NewActionUserRequestChangeIndex(client)
	actionShow := NewActionUserRequestChangeShow(client)
	actionCreate := NewActionUserRequestChangeCreate(client)
	actionResolve := NewActionUserRequestChangeResolve(client)

	return &ResourceUserRequestChange{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Create: actionCreate,
		New: actionCreate,
		Resolve: actionResolve,
	}
}
