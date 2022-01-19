package client

// Type for resource User_cluster_resource_package
type ResourceUserClusterResourcePackage struct {
	// Pointer to client
	Client *Client

	// Resource User_cluster_resource_package.Item
	Item *ResourceUserClusterResourcePackageItem
	// Action User_cluster_resource_package#Create
	Create *ActionUserClusterResourcePackageCreate
	// Action User_cluster_resource_package#Create
	New *ActionUserClusterResourcePackageCreate
	// Action User_cluster_resource_package#Delete
	Delete *ActionUserClusterResourcePackageDelete
	// Action User_cluster_resource_package#Delete
	Destroy *ActionUserClusterResourcePackageDelete
	// Action User_cluster_resource_package#Index
	Index *ActionUserClusterResourcePackageIndex
	// Action User_cluster_resource_package#Index
	List *ActionUserClusterResourcePackageIndex
	// Action User_cluster_resource_package#Show
	Show *ActionUserClusterResourcePackageShow
	// Action User_cluster_resource_package#Show
	Find *ActionUserClusterResourcePackageShow
	// Action User_cluster_resource_package#Update
	Update *ActionUserClusterResourcePackageUpdate
}

func NewResourceUserClusterResourcePackage(client *Client) *ResourceUserClusterResourcePackage {
	actionCreate := NewActionUserClusterResourcePackageCreate(client)
	actionDelete := NewActionUserClusterResourcePackageDelete(client)
	actionIndex := NewActionUserClusterResourcePackageIndex(client)
	actionShow := NewActionUserClusterResourcePackageShow(client)
	actionUpdate := NewActionUserClusterResourcePackageUpdate(client)

	return &ResourceUserClusterResourcePackage{
		Client:  client,
		Item:    NewResourceUserClusterResourcePackageItem(client),
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
