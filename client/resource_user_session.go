package client

// Type for resource User_session
type ResourceUserSession struct {
	// Pointer to client
	Client *Client

	// Action User_session#Close
	Close *ActionUserSessionClose
	// Action User_session#Create
	Create *ActionUserSessionCreate
	// Action User_session#Create
	New *ActionUserSessionCreate
	// Action User_session#Index
	Index *ActionUserSessionIndex
	// Action User_session#Index
	List *ActionUserSessionIndex
	// Action User_session#Show
	Show *ActionUserSessionShow
	// Action User_session#Show
	Find *ActionUserSessionShow
	// Action User_session#Update
	Update *ActionUserSessionUpdate
}

func NewResourceUserSession(client *Client) *ResourceUserSession {
	actionClose := NewActionUserSessionClose(client)
	actionCreate := NewActionUserSessionCreate(client)
	actionIndex := NewActionUserSessionIndex(client)
	actionShow := NewActionUserSessionShow(client)
	actionUpdate := NewActionUserSessionUpdate(client)

	return &ResourceUserSession{
		Client: client,
		Close:  actionClose,
		Create: actionCreate,
		New:    actionCreate,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
		Update: actionUpdate,
	}
}
