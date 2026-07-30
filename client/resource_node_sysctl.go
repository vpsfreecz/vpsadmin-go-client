package client

// Type for resource Node_sysctl
type ResourceNodeSysctl struct {
	// Pointer to client
	Client *Client

	// Action Node_sysctl#Index
	Index *ActionNodeSysctlIndex
	// Action Node_sysctl#Index
	List *ActionNodeSysctlIndex
}

func NewResourceNodeSysctl(client *Client) *ResourceNodeSysctl {
	actionIndex := NewActionNodeSysctlIndex(client)

	return &ResourceNodeSysctl{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
