package client

// Type for resource Default_object_cluster_resource
type ResourceDefaultObjectClusterResource struct {
	// Pointer to client
	Client *Client

	// Action Default_object_cluster_resource#Create
	Create *ActionDefaultObjectClusterResourceCreate
	// Action Default_object_cluster_resource#Create
	New *ActionDefaultObjectClusterResourceCreate
	// Action Default_object_cluster_resource#Delete
	Delete *ActionDefaultObjectClusterResourceDelete
	// Action Default_object_cluster_resource#Delete
	Destroy *ActionDefaultObjectClusterResourceDelete
	// Action Default_object_cluster_resource#Index
	Index *ActionDefaultObjectClusterResourceIndex
	// Action Default_object_cluster_resource#Index
	List *ActionDefaultObjectClusterResourceIndex
	// Action Default_object_cluster_resource#Show
	Show *ActionDefaultObjectClusterResourceShow
	// Action Default_object_cluster_resource#Show
	Find *ActionDefaultObjectClusterResourceShow
	// Action Default_object_cluster_resource#Update
	Update *ActionDefaultObjectClusterResourceUpdate
}

func NewResourceDefaultObjectClusterResource(client *Client) *ResourceDefaultObjectClusterResource {
	actionCreate := NewActionDefaultObjectClusterResourceCreate(client)
	actionDelete := NewActionDefaultObjectClusterResourceDelete(client)
	actionIndex := NewActionDefaultObjectClusterResourceIndex(client)
	actionShow := NewActionDefaultObjectClusterResourceShow(client)
	actionUpdate := NewActionDefaultObjectClusterResourceUpdate(client)

	return &ResourceDefaultObjectClusterResource{
		Client:  client,
		Create:  actionCreate,
		New:     actionCreate,
		Delete:  actionDelete,
		Destroy: actionDelete,
		Index:   actionIndex,
		List:    actionIndex,
		Show:    actionShow,
		Find:    actionShow,
		Update:  actionUpdate,
	}
}
