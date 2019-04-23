package client

// Type for resource Mail_log
type ResourceMailLog struct {
	// Pointer to client
	Client *Client

	// Action Mail_log#Index
	Index *ActionMailLogIndex
	// Action Mail_log#Index
	List *ActionMailLogIndex
	// Action Mail_log#Show
	Show *ActionMailLogShow
	// Action Mail_log#Show
	Find *ActionMailLogShow
}

func NewResourceMailLog(client *Client) *ResourceMailLog {
	actionIndex := NewActionMailLogIndex(client)
	actionShow := NewActionMailLogShow(client)

	return &ResourceMailLog{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
	}
}
