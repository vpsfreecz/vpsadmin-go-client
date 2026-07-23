package client

// Type for resource Node_sysctl_change
type ResourceNodeSysctlChange struct {
	// Pointer to client
	Client *Client

	// Action Node_sysctl_change#Index
	Index *ActionNodeSysctlChangeIndex
	// Action Node_sysctl_change#Index
	List *ActionNodeSysctlChangeIndex
}

func NewResourceNodeSysctlChange(client *Client) *ResourceNodeSysctlChange {
	actionIndex := NewActionNodeSysctlChangeIndex(client)

	return &ResourceNodeSysctlChange{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
