package client

// Type for resource User.State_log
type ResourceUserStateLog struct {
	// Pointer to client
	Client *Client

	// Action User.State_log#Index
	Index *ActionUserStateLogIndex
	// Action User.State_log#Index
	List *ActionUserStateLogIndex
	// Action User.State_log#Show
	Show *ActionUserStateLogShow
	// Action User.State_log#Show
	Find *ActionUserStateLogShow
}

func NewResourceUserStateLog(client *Client) *ResourceUserStateLog {
	actionIndex := NewActionUserStateLogIndex(client)
	actionShow := NewActionUserStateLogShow(client)

	return &ResourceUserStateLog{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
