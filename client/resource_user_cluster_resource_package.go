package client

// Type for resource User_cluster_resource_package
type ResourceUserClusterResourcePackage struct {
	// Pointer to client
	Client *Client

	// Resource User_cluster_resource_package.Item
	Item *ResourceUserClusterResourcePackageItem
	// Action User_cluster_resource_package#Index
	Index *ActionUserClusterResourcePackageIndex
	// Action User_cluster_resource_package#Index
	List *ActionUserClusterResourcePackageIndex
	// Action User_cluster_resource_package#Show
	Show *ActionUserClusterResourcePackageShow
	// Action User_cluster_resource_package#Show
	Find *ActionUserClusterResourcePackageShow
	// Action User_cluster_resource_package#Create
	Create *ActionUserClusterResourcePackageCreate
	// Action User_cluster_resource_package#Create
	New *ActionUserClusterResourcePackageCreate
	// Action User_cluster_resource_package#Update
	Update *ActionUserClusterResourcePackageUpdate
	// Action User_cluster_resource_package#Delete
	Delete *ActionUserClusterResourcePackageDelete
	// Action User_cluster_resource_package#Delete
	Destroy *ActionUserClusterResourcePackageDelete
}

func NewResourceUserClusterResourcePackage(client *Client) *ResourceUserClusterResourcePackage {
	actionIndex := NewActionUserClusterResourcePackageIndex(client)
	actionShow := NewActionUserClusterResourcePackageShow(client)
	actionCreate := NewActionUserClusterResourcePackageCreate(client)
	actionUpdate := NewActionUserClusterResourcePackageUpdate(client)
	actionDelete := NewActionUserClusterResourcePackageDelete(client)

	return &ResourceUserClusterResourcePackage{
		Client: client,
		Item: NewResourceUserClusterResourcePackageItem(client),
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Create: actionCreate,
		New: actionCreate,
		Update: actionUpdate,
		Delete: actionDelete,
		Destroy: actionDelete,
	}
}
