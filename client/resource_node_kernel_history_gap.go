package client

// Type for resource Node_kernel_history_gap
type ResourceNodeKernelHistoryGap struct {
	// Pointer to client
	Client *Client

	// Action Node_kernel_history_gap#Index
	Index *ActionNodeKernelHistoryGapIndex
	// Action Node_kernel_history_gap#Index
	List *ActionNodeKernelHistoryGapIndex
}

func NewResourceNodeKernelHistoryGap(client *Client) *ResourceNodeKernelHistoryGap {
	actionIndex := NewActionNodeKernelHistoryGapIndex(client)

	return &ResourceNodeKernelHistoryGap{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
