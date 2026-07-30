package client

// Type for resource Node_kernel_history_state
type ResourceNodeKernelHistoryState struct {
	// Pointer to client
	Client *Client

	// Action Node_kernel_history_state#Index
	Index *ActionNodeKernelHistoryStateIndex
	// Action Node_kernel_history_state#Index
	List *ActionNodeKernelHistoryStateIndex
	// Action Node_kernel_history_state#Show
	Show *ActionNodeKernelHistoryStateShow
	// Action Node_kernel_history_state#Show
	Find *ActionNodeKernelHistoryStateShow
}

func NewResourceNodeKernelHistoryState(client *Client) *ResourceNodeKernelHistoryState {
	actionIndex := NewActionNodeKernelHistoryStateIndex(client)
	actionShow := NewActionNodeKernelHistoryStateShow(client)

	return &ResourceNodeKernelHistoryState{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
