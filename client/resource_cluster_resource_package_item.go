package client

// Type for resource Cluster_resource_package.Item
type ResourceClusterResourcePackageItem struct {
	// Pointer to client
	Client *Client

	// Action Cluster_resource_package.Item#Create
	Create *ActionClusterResourcePackageItemCreate
	// Action Cluster_resource_package.Item#Create
	New *ActionClusterResourcePackageItemCreate
	// Action Cluster_resource_package.Item#Delete
	Delete *ActionClusterResourcePackageItemDelete
	// Action Cluster_resource_package.Item#Delete
	Destroy *ActionClusterResourcePackageItemDelete
	// Action Cluster_resource_package.Item#Index
	Index *ActionClusterResourcePackageItemIndex
	// Action Cluster_resource_package.Item#Index
	List *ActionClusterResourcePackageItemIndex
	// Action Cluster_resource_package.Item#Show
	Show *ActionClusterResourcePackageItemShow
	// Action Cluster_resource_package.Item#Show
	Find *ActionClusterResourcePackageItemShow
	// Action Cluster_resource_package.Item#Update
	Update *ActionClusterResourcePackageItemUpdate
}

func NewResourceClusterResourcePackageItem(client *Client) *ResourceClusterResourcePackageItem {
	actionCreate := NewActionClusterResourcePackageItemCreate(client)
	actionDelete := NewActionClusterResourcePackageItemDelete(client)
	actionIndex := NewActionClusterResourcePackageItemIndex(client)
	actionShow := NewActionClusterResourcePackageItemShow(client)
	actionUpdate := NewActionClusterResourcePackageItemUpdate(client)

	return &ResourceClusterResourcePackageItem{
		Client: client,
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
