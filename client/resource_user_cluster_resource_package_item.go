package client

// Type for resource User_cluster_resource_package.Item
type ResourceUserClusterResourcePackageItem struct {
	// Pointer to client
	Client *Client

	// Action User_cluster_resource_package.Item#Index
	Index *ActionUserClusterResourcePackageItemIndex
	// Action User_cluster_resource_package.Item#Index
	List *ActionUserClusterResourcePackageItemIndex
	// Action User_cluster_resource_package.Item#Show
	Show *ActionUserClusterResourcePackageItemShow
	// Action User_cluster_resource_package.Item#Show
	Find *ActionUserClusterResourcePackageItemShow
}

func NewResourceUserClusterResourcePackageItem(client *Client) *ResourceUserClusterResourcePackageItem {
	actionIndex := NewActionUserClusterResourcePackageItemIndex(client)
	actionShow := NewActionUserClusterResourcePackageItemShow(client)

	return &ResourceUserClusterResourcePackageItem{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
