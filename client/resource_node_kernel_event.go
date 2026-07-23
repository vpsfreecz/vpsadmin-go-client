package client

// Type for resource Node_kernel_event
type ResourceNodeKernelEvent struct {
	// Pointer to client
	Client *Client

	// Action Node_kernel_event#Index
	Index *ActionNodeKernelEventIndex
	// Action Node_kernel_event#Index
	List *ActionNodeKernelEventIndex
	// Action Node_kernel_event#Show
	Show *ActionNodeKernelEventShow
	// Action Node_kernel_event#Show
	Find *ActionNodeKernelEventShow
}

func NewResourceNodeKernelEvent(client *Client) *ResourceNodeKernelEvent {
	actionIndex := NewActionNodeKernelEventIndex(client)
	actionShow := NewActionNodeKernelEventShow(client)

	return &ResourceNodeKernelEvent{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
