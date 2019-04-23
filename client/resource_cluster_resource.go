package client

// Type for resource Cluster_resource
type ResourceClusterResource struct {
	// Pointer to client
	Client *Client

	// Action Cluster_resource#Index
	Index *ActionClusterResourceIndex
	// Action Cluster_resource#Index
	List *ActionClusterResourceIndex
	// Action Cluster_resource#Show
	Show *ActionClusterResourceShow
	// Action Cluster_resource#Show
	Find *ActionClusterResourceShow
	// Action Cluster_resource#Create
	Create *ActionClusterResourceCreate
	// Action Cluster_resource#Create
	New *ActionClusterResourceCreate
	// Action Cluster_resource#Update
	Update *ActionClusterResourceUpdate
}

func NewResourceClusterResource(client *Client) *ResourceClusterResource {
	actionIndex := NewActionClusterResourceIndex(client)
	actionShow := NewActionClusterResourceShow(client)
	actionCreate := NewActionClusterResourceCreate(client)
	actionUpdate := NewActionClusterResourceUpdate(client)

	return &ResourceClusterResource{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Create: actionCreate,
		New: actionCreate,
		Update: actionUpdate,
	}
}
