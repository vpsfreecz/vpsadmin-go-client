package client

// Type for resource Cluster
type ResourceCluster struct {
	// Pointer to client
	Client *Client

	// Action Cluster#Show
	Show *ActionClusterShow
	// Action Cluster#Show
	Find *ActionClusterShow
	// Action Cluster#Public_stats
	PublicStats *ActionClusterPublicStats
	// Action Cluster#Full_stats
	FullStats *ActionClusterFullStats
	// Action Cluster#Search
	Search *ActionClusterSearch
	// Action Cluster#Authorize_migration_keys
	AuthorizeMigrationKeys *ActionClusterAuthorizeMigrationKeys
	// Action Cluster#Set_maintenance
	SetMaintenance *ActionClusterSetMaintenance
}

func NewResourceCluster(client *Client) *ResourceCluster {
	actionShow := NewActionClusterShow(client)
	actionPublicStats := NewActionClusterPublicStats(client)
	actionFullStats := NewActionClusterFullStats(client)
	actionSearch := NewActionClusterSearch(client)
	actionAuthorizeMigrationKeys := NewActionClusterAuthorizeMigrationKeys(client)
	actionSetMaintenance := NewActionClusterSetMaintenance(client)

	return &ResourceCluster{
		Client: client,
		Show: actionShow,
		Find: actionShow,
		PublicStats: actionPublicStats,
		FullStats: actionFullStats,
		Search: actionSearch,
		AuthorizeMigrationKeys: actionAuthorizeMigrationKeys,
		SetMaintenance: actionSetMaintenance,
	}
}
