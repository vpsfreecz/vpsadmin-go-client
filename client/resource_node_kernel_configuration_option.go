package client

// Type for resource Node_kernel_configuration_option
type ResourceNodeKernelConfigurationOption struct {
	// Pointer to client
	Client *Client

	// Action Node_kernel_configuration_option#Index
	Index *ActionNodeKernelConfigurationOptionIndex
	// Action Node_kernel_configuration_option#Index
	List *ActionNodeKernelConfigurationOptionIndex
}

func NewResourceNodeKernelConfigurationOption(client *Client) *ResourceNodeKernelConfigurationOption {
	actionIndex := NewActionNodeKernelConfigurationOptionIndex(client)

	return &ResourceNodeKernelConfigurationOption{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
