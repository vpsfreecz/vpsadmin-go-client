package client

// Type for resource Node.Kernel_history
type ResourceNodeKernelHistory struct {
	// Pointer to client
	Client *Client

	// Action Node.Kernel_history#Index
	Index *ActionNodeKernelHistoryIndex
	// Action Node.Kernel_history#Index
	List *ActionNodeKernelHistoryIndex
}

func NewResourceNodeKernelHistory(client *Client) *ResourceNodeKernelHistory {
	actionIndex := NewActionNodeKernelHistoryIndex(client)

	return &ResourceNodeKernelHistory{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
