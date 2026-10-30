package client

// Type for resource Node_software_change
type ResourceNodeSoftwareChange struct {
	// Pointer to client
	Client *Client

	// Action Node_software_change#Index
	Index *ActionNodeSoftwareChangeIndex
	// Action Node_software_change#Index
	List *ActionNodeSoftwareChangeIndex
}

func NewResourceNodeSoftwareChange(client *Client) *ResourceNodeSoftwareChange {
	actionIndex := NewActionNodeSoftwareChangeIndex(client)

	return &ResourceNodeSoftwareChange{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
