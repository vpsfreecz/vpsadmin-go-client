package client

// Type for resource Security_advisory.Node_status
type ResourceSecurityAdvisoryNodeStatus struct {
	// Pointer to client
	Client *Client

	// Action Security_advisory.Node_status#Create
	Create *ActionSecurityAdvisoryNodeStatusCreate
	// Action Security_advisory.Node_status#Create
	New *ActionSecurityAdvisoryNodeStatusCreate
	// Action Security_advisory.Node_status#Delete
	Delete *ActionSecurityAdvisoryNodeStatusDelete
	// Action Security_advisory.Node_status#Delete
	Destroy *ActionSecurityAdvisoryNodeStatusDelete
	// Action Security_advisory.Node_status#Index
	Index *ActionSecurityAdvisoryNodeStatusIndex
	// Action Security_advisory.Node_status#Index
	List *ActionSecurityAdvisoryNodeStatusIndex
	// Action Security_advisory.Node_status#Update
	Update *ActionSecurityAdvisoryNodeStatusUpdate
}

func NewResourceSecurityAdvisoryNodeStatus(client *Client) *ResourceSecurityAdvisoryNodeStatus {
	actionCreate := NewActionSecurityAdvisoryNodeStatusCreate(client)
	actionDelete := NewActionSecurityAdvisoryNodeStatusDelete(client)
	actionIndex := NewActionSecurityAdvisoryNodeStatusIndex(client)
	actionUpdate := NewActionSecurityAdvisoryNodeStatusUpdate(client)

	return &ResourceSecurityAdvisoryNodeStatus{
		Client:  client,
		Create:  actionCreate,
		New:     actionCreate,
		Delete:  actionDelete,
		Destroy: actionDelete,
		Index:   actionIndex,
		List:    actionIndex,
		Update:  actionUpdate,
	}
}
