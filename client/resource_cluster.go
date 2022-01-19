package client

// Type for resource Cluster
type ResourceCluster struct {
	// Pointer to client
	Client *Client

	// Action Cluster#Full_stats
	FullStats *ActionClusterFullStats
	// Action Cluster#Generate_migration_keys
	GenerateMigrationKeys *ActionClusterGenerateMigrationKeys
	// Action Cluster#Public_stats
	PublicStats *ActionClusterPublicStats
	// Action Cluster#Search
	Search *ActionClusterSearch
	// Action Cluster#Set_maintenance
	SetMaintenance *ActionClusterSetMaintenance
	// Action Cluster#Show
	Show *ActionClusterShow
	// Action Cluster#Show
	Find *ActionClusterShow
}

func NewResourceCluster(client *Client) *ResourceCluster {
	actionFullStats := NewActionClusterFullStats(client)
	actionGenerateMigrationKeys := NewActionClusterGenerateMigrationKeys(client)
	actionPublicStats := NewActionClusterPublicStats(client)
	actionSearch := NewActionClusterSearch(client)
	actionSetMaintenance := NewActionClusterSetMaintenance(client)
	actionShow := NewActionClusterShow(client)

	return &ResourceCluster{
		Client: client,
		FullStats: actionFullStats,
		GenerateMigrationKeys: actionGenerateMigrationKeys,
		PublicStats: actionPublicStats,
		Search: actionSearch,
		SetMaintenance: actionSetMaintenance,
		Show: actionShow,
		Find: actionShow,
	}
}
