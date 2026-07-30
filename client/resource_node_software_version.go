package client

// Type for resource Node_software_version
type ResourceNodeSoftwareVersion struct {
	// Pointer to client
	Client *Client

	// Action Node_software_version#Index
	Index *ActionNodeSoftwareVersionIndex
	// Action Node_software_version#Index
	List *ActionNodeSoftwareVersionIndex
}

func NewResourceNodeSoftwareVersion(client *Client) *ResourceNodeSoftwareVersion {
	actionIndex := NewActionNodeSoftwareVersionIndex(client)

	return &ResourceNodeSoftwareVersion{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
