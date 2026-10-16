package client

// Type for resource Password_change_log
type ResourcePasswordChangeLog struct {
	// Pointer to client
	Client *Client

	// Action Password_change_log#Index
	Index *ActionPasswordChangeLogIndex
	// Action Password_change_log#Index
	List *ActionPasswordChangeLogIndex
	// Action Password_change_log#Show
	Show *ActionPasswordChangeLogShow
	// Action Password_change_log#Show
	Find *ActionPasswordChangeLogShow
}

func NewResourcePasswordChangeLog(client *Client) *ResourcePasswordChangeLog {
	actionIndex := NewActionPasswordChangeLogIndex(client)
	actionShow := NewActionPasswordChangeLogShow(client)

	return &ResourcePasswordChangeLog{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
