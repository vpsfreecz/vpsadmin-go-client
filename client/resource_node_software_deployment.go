package client

// Type for resource Node_software_deployment
type ResourceNodeSoftwareDeployment struct {
	// Pointer to client
	Client *Client

	// Action Node_software_deployment#Index
	Index *ActionNodeSoftwareDeploymentIndex
	// Action Node_software_deployment#Index
	List *ActionNodeSoftwareDeploymentIndex
}

func NewResourceNodeSoftwareDeployment(client *Client) *ResourceNodeSoftwareDeployment {
	actionIndex := NewActionNodeSoftwareDeploymentIndex(client)

	return &ResourceNodeSoftwareDeployment{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
