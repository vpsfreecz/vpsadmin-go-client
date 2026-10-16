package client

// Type for resource Node_kernel_livepatch
type ResourceNodeKernelLivepatch struct {
	// Pointer to client
	Client *Client

	// Action Node_kernel_livepatch#Index
	Index *ActionNodeKernelLivepatchIndex
	// Action Node_kernel_livepatch#Index
	List *ActionNodeKernelLivepatchIndex
	// Action Node_kernel_livepatch#Show
	Show *ActionNodeKernelLivepatchShow
	// Action Node_kernel_livepatch#Show
	Find *ActionNodeKernelLivepatchShow
}

func NewResourceNodeKernelLivepatch(client *Client) *ResourceNodeKernelLivepatch {
	actionIndex := NewActionNodeKernelLivepatchIndex(client)
	actionShow := NewActionNodeKernelLivepatchShow(client)

	return &ResourceNodeKernelLivepatch{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
