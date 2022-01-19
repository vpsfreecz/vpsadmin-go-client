package client

// Type for resource User_request.Change
type ResourceUserRequestChange struct {
	// Pointer to client
	Client *Client

	// Action User_request.Change#Create
	Create *ActionUserRequestChangeCreate
	// Action User_request.Change#Create
	New *ActionUserRequestChangeCreate
	// Action User_request.Change#Index
	Index *ActionUserRequestChangeIndex
	// Action User_request.Change#Index
	List *ActionUserRequestChangeIndex
	// Action User_request.Change#Resolve
	Resolve *ActionUserRequestChangeResolve
	// Action User_request.Change#Show
	Show *ActionUserRequestChangeShow
	// Action User_request.Change#Show
	Find *ActionUserRequestChangeShow
}

func NewResourceUserRequestChange(client *Client) *ResourceUserRequestChange {
	actionCreate := NewActionUserRequestChangeCreate(client)
	actionIndex := NewActionUserRequestChangeIndex(client)
	actionResolve := NewActionUserRequestChangeResolve(client)
	actionShow := NewActionUserRequestChangeShow(client)

	return &ResourceUserRequestChange{
		Client:  client,
		Create:  actionCreate,
		New:     actionCreate,
		Index:   actionIndex,
		List:    actionIndex,
		Resolve: actionResolve,
		Show:    actionShow,
		Find:    actionShow,
	}
}
