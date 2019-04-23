package client

// Type for resource Node.Status
type ResourceNodeStatus struct {
	// Pointer to client
	Client *Client

	// Action Node.Status#Index
	Index *ActionNodeStatusIndex
	// Action Node.Status#Index
	List *ActionNodeStatusIndex
	// Action Node.Status#Show
	Show *ActionNodeStatusShow
	// Action Node.Status#Show
	Find *ActionNodeStatusShow
}

func NewResourceNodeStatus(client *Client) *ResourceNodeStatus {
	actionIndex := NewActionNodeStatusIndex(client)
	actionShow := NewActionNodeStatusShow(client)

	return &ResourceNodeStatus{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
	}
}
