package client

// Type for resource Node_kernel_livepatch_patch
type ResourceNodeKernelLivepatchPatch struct {
	// Pointer to client
	Client *Client

	// Action Node_kernel_livepatch_patch#Index
	Index *ActionNodeKernelLivepatchPatchIndex
	// Action Node_kernel_livepatch_patch#Index
	List *ActionNodeKernelLivepatchPatchIndex
}

func NewResourceNodeKernelLivepatchPatch(client *Client) *ResourceNodeKernelLivepatchPatch {
	actionIndex := NewActionNodeKernelLivepatchPatchIndex(client)

	return &ResourceNodeKernelLivepatchPatch{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
