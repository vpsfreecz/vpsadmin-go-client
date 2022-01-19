package client

// Type for resource User.Cluster_resource
type ResourceUserClusterResource struct {
	// Pointer to client
	Client *Client

	// Action User.Cluster_resource#Create
	Create *ActionUserClusterResourceCreate
	// Action User.Cluster_resource#Create
	New *ActionUserClusterResourceCreate
	// Action User.Cluster_resource#Index
	Index *ActionUserClusterResourceIndex
	// Action User.Cluster_resource#Index
	List *ActionUserClusterResourceIndex
	// Action User.Cluster_resource#Show
	Show *ActionUserClusterResourceShow
	// Action User.Cluster_resource#Show
	Find *ActionUserClusterResourceShow
}

func NewResourceUserClusterResource(client *Client) *ResourceUserClusterResource {
	actionCreate := NewActionUserClusterResourceCreate(client)
	actionIndex := NewActionUserClusterResourceIndex(client)
	actionShow := NewActionUserClusterResourceShow(client)

	return &ResourceUserClusterResource{
		Client: client,
		Create: actionCreate,
		New:    actionCreate,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
