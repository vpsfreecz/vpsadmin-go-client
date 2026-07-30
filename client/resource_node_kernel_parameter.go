package client

// Type for resource Node_kernel_parameter
type ResourceNodeKernelParameter struct {
	// Pointer to client
	Client *Client

	// Action Node_kernel_parameter#Index
	Index *ActionNodeKernelParameterIndex
	// Action Node_kernel_parameter#Index
	List *ActionNodeKernelParameterIndex
}

func NewResourceNodeKernelParameter(client *Client) *ResourceNodeKernelParameter {
	actionIndex := NewActionNodeKernelParameterIndex(client)

	return &ResourceNodeKernelParameter{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
