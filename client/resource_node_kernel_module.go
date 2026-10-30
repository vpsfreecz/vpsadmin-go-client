package client

// Type for resource Node_kernel_module
type ResourceNodeKernelModule struct {
	// Pointer to client
	Client *Client

	// Action Node_kernel_module#Index
	Index *ActionNodeKernelModuleIndex
	// Action Node_kernel_module#Index
	List *ActionNodeKernelModuleIndex
}

func NewResourceNodeKernelModule(client *Client) *ResourceNodeKernelModule {
	actionIndex := NewActionNodeKernelModuleIndex(client)

	return &ResourceNodeKernelModule{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
