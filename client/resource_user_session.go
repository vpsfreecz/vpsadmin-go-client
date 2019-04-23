package client

// Type for resource User_session
type ResourceUserSession struct {
	// Pointer to client
	Client *Client

	// Action User_session#Index
	Index *ActionUserSessionIndex
	// Action User_session#Index
	List *ActionUserSessionIndex
	// Action User_session#Show
	Show *ActionUserSessionShow
	// Action User_session#Show
	Find *ActionUserSessionShow
}

func NewResourceUserSession(client *Client) *ResourceUserSession {
	actionIndex := NewActionUserSessionIndex(client)
	actionShow := NewActionUserSessionShow(client)

	return &ResourceUserSession{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
	}
}
