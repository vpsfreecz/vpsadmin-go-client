package client

// Type for resource Vps.Status
type ResourceVpsStatus struct {
	// Pointer to client
	Client *Client

	// Action Vps.Status#Index
	Index *ActionVpsStatusIndex
	// Action Vps.Status#Index
	List *ActionVpsStatusIndex
	// Action Vps.Status#Show
	Show *ActionVpsStatusShow
	// Action Vps.Status#Show
	Find *ActionVpsStatusShow
}

func NewResourceVpsStatus(client *Client) *ResourceVpsStatus {
	actionIndex := NewActionVpsStatusIndex(client)
	actionShow := NewActionVpsStatusShow(client)

	return &ResourceVpsStatus{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
