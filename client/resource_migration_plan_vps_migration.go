package client

// Type for resource Migration_plan.Vps_migration
type ResourceMigrationPlanVpsMigration struct {
	// Pointer to client
	Client *Client

	// Action Migration_plan.Vps_migration#Create
	Create *ActionMigrationPlanVpsMigrationCreate
	// Action Migration_plan.Vps_migration#Create
	New *ActionMigrationPlanVpsMigrationCreate
	// Action Migration_plan.Vps_migration#Index
	Index *ActionMigrationPlanVpsMigrationIndex
	// Action Migration_plan.Vps_migration#Index
	List *ActionMigrationPlanVpsMigrationIndex
	// Action Migration_plan.Vps_migration#Show
	Show *ActionMigrationPlanVpsMigrationShow
	// Action Migration_plan.Vps_migration#Show
	Find *ActionMigrationPlanVpsMigrationShow
}

func NewResourceMigrationPlanVpsMigration(client *Client) *ResourceMigrationPlanVpsMigration {
	actionCreate := NewActionMigrationPlanVpsMigrationCreate(client)
	actionIndex := NewActionMigrationPlanVpsMigrationIndex(client)
	actionShow := NewActionMigrationPlanVpsMigrationShow(client)

	return &ResourceMigrationPlanVpsMigration{
		Client: client,
		Create: actionCreate,
		New: actionCreate,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
	}
}
