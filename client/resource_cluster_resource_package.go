package client

// Type for resource Cluster_resource_package
type ResourceClusterResourcePackage struct {
	// Pointer to client
	Client *Client

	// Resource Cluster_resource_package.Item
	Item *ResourceClusterResourcePackageItem
	// Action Cluster_resource_package#Create
	Create *ActionClusterResourcePackageCreate
	// Action Cluster_resource_package#Create
	New *ActionClusterResourcePackageCreate
	// Action Cluster_resource_package#Delete
	Delete *ActionClusterResourcePackageDelete
	// Action Cluster_resource_package#Delete
	Destroy *ActionClusterResourcePackageDelete
	// Action Cluster_resource_package#Index
	Index *ActionClusterResourcePackageIndex
	// Action Cluster_resource_package#Index
	List *ActionClusterResourcePackageIndex
	// Action Cluster_resource_package#Show
	Show *ActionClusterResourcePackageShow
	// Action Cluster_resource_package#Show
	Find *ActionClusterResourcePackageShow
	// Action Cluster_resource_package#Update
	Update *ActionClusterResourcePackageUpdate
}

func NewResourceClusterResourcePackage(client *Client) *ResourceClusterResourcePackage {
	actionCreate := NewActionClusterResourcePackageCreate(client)
	actionDelete := NewActionClusterResourcePackageDelete(client)
	actionIndex := NewActionClusterResourcePackageIndex(client)
	actionShow := NewActionClusterResourcePackageShow(client)
	actionUpdate := NewActionClusterResourcePackageUpdate(client)

	return &ResourceClusterResourcePackage{
		Client: client,
		Item: NewResourceClusterResourcePackageItem(client),
		Create: actionCreate,
		New: actionCreate,
		Delete: actionDelete,
		Destroy: actionDelete,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Update: actionUpdate,
	}
}
