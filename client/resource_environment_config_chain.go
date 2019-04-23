package client

// Type for resource Environment.Config_chain
type ResourceEnvironmentConfigChain struct {
	// Pointer to client
	Client *Client

	// Action Environment.Config_chain#Index
	Index *ActionEnvironmentConfigChainIndex
	// Action Environment.Config_chain#Index
	List *ActionEnvironmentConfigChainIndex
	// Action Environment.Config_chain#Replace
	Replace *ActionEnvironmentConfigChainReplace
}

func NewResourceEnvironmentConfigChain(client *Client) *ResourceEnvironmentConfigChain {
	actionIndex := NewActionEnvironmentConfigChainIndex(client)
	actionReplace := NewActionEnvironmentConfigChainReplace(client)

	return &ResourceEnvironmentConfigChain{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Replace: actionReplace,
	}
}
