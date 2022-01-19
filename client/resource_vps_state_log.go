package client

// Type for resource Vps.State_log
type ResourceVpsStateLog struct {
	// Pointer to client
	Client *Client

	// Action Vps.State_log#Index
	Index *ActionVpsStateLogIndex
	// Action Vps.State_log#Index
	List *ActionVpsStateLogIndex
	// Action Vps.State_log#Show
	Show *ActionVpsStateLogShow
	// Action Vps.State_log#Show
	Find *ActionVpsStateLogShow
}

func NewResourceVpsStateLog(client *Client) *ResourceVpsStateLog {
	actionIndex := NewActionVpsStateLogIndex(client)
	actionShow := NewActionVpsStateLogShow(client)

	return &ResourceVpsStateLog{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
