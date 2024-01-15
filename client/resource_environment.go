package client

// Type for resource Environment
type ResourceEnvironment struct {
	// Pointer to client
	Client *Client

	// Resource Environment.Dataset_plan
	DatasetPlan *ResourceEnvironmentDatasetPlan
	// Action Environment#Create
	Create *ActionEnvironmentCreate
	// Action Environment#Create
	New *ActionEnvironmentCreate
	// Action Environment#Index
	Index *ActionEnvironmentIndex
	// Action Environment#Index
	List *ActionEnvironmentIndex
	// Action Environment#Set_maintenance
	SetMaintenance *ActionEnvironmentSetMaintenance
	// Action Environment#Show
	Show *ActionEnvironmentShow
	// Action Environment#Show
	Find *ActionEnvironmentShow
	// Action Environment#Update
	Update *ActionEnvironmentUpdate
}

func NewResourceEnvironment(client *Client) *ResourceEnvironment {
	actionCreate := NewActionEnvironmentCreate(client)
	actionIndex := NewActionEnvironmentIndex(client)
	actionSetMaintenance := NewActionEnvironmentSetMaintenance(client)
	actionShow := NewActionEnvironmentShow(client)
	actionUpdate := NewActionEnvironmentUpdate(client)

	return &ResourceEnvironment{
		Client:         client,
		DatasetPlan:    NewResourceEnvironmentDatasetPlan(client),
		Create:         actionCreate,
		New:            actionCreate,
		Index:          actionIndex,
		List:           actionIndex,
		SetMaintenance: actionSetMaintenance,
		Show:           actionShow,
		Find:           actionShow,
		Update:         actionUpdate,
	}
}
