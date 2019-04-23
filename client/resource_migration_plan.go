package client

// Type for resource Migration_plan
type ResourceMigrationPlan struct {
	// Pointer to client
	Client *Client

	// Resource Migration_plan.Vps_migration
	VpsMigration *ResourceMigrationPlanVpsMigration
	// Action Migration_plan#Index
	Index *ActionMigrationPlanIndex
	// Action Migration_plan#Index
	List *ActionMigrationPlanIndex
	// Action Migration_plan#Show
	Show *ActionMigrationPlanShow
	// Action Migration_plan#Show
	Find *ActionMigrationPlanShow
	// Action Migration_plan#Create
	Create *ActionMigrationPlanCreate
	// Action Migration_plan#Create
	New *ActionMigrationPlanCreate
	// Action Migration_plan#Start
	Start *ActionMigrationPlanStart
	// Action Migration_plan#Cancel
	Cancel *ActionMigrationPlanCancel
	// Action Migration_plan#Delete
	Delete *ActionMigrationPlanDelete
	// Action Migration_plan#Delete
	Destroy *ActionMigrationPlanDelete
}

func NewResourceMigrationPlan(client *Client) *ResourceMigrationPlan {
	actionIndex := NewActionMigrationPlanIndex(client)
	actionShow := NewActionMigrationPlanShow(client)
	actionCreate := NewActionMigrationPlanCreate(client)
	actionStart := NewActionMigrationPlanStart(client)
	actionCancel := NewActionMigrationPlanCancel(client)
	actionDelete := NewActionMigrationPlanDelete(client)

	return &ResourceMigrationPlan{
		Client: client,
		VpsMigration: NewResourceMigrationPlanVpsMigration(client),
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Create: actionCreate,
		New: actionCreate,
		Start: actionStart,
		Cancel: actionCancel,
		Delete: actionDelete,
		Destroy: actionDelete,
	}
}
