package client

// Type for resource Environment
type ResourceEnvironment struct {
	// Pointer to client
	Client *Client

	// Resource Environment.Config_chain
	ConfigChain *ResourceEnvironmentConfigChain
	// Resource Environment.Dataset_plan
	DatasetPlan *ResourceEnvironmentDatasetPlan
	// Action Environment#Index
	Index *ActionEnvironmentIndex
	// Action Environment#Index
	List *ActionEnvironmentIndex
	// Action Environment#Create
	Create *ActionEnvironmentCreate
	// Action Environment#Create
	New *ActionEnvironmentCreate
	// Action Environment#Show
	Show *ActionEnvironmentShow
	// Action Environment#Show
	Find *ActionEnvironmentShow
	// Action Environment#Update
	Update *ActionEnvironmentUpdate
	// Action Environment#Set_maintenance
	SetMaintenance *ActionEnvironmentSetMaintenance
}

func NewResourceEnvironment(client *Client) *ResourceEnvironment {
	actionIndex := NewActionEnvironmentIndex(client)
	actionCreate := NewActionEnvironmentCreate(client)
	actionShow := NewActionEnvironmentShow(client)
	actionUpdate := NewActionEnvironmentUpdate(client)
	actionSetMaintenance := NewActionEnvironmentSetMaintenance(client)

	return &ResourceEnvironment{
		Client: client,
		ConfigChain: NewResourceEnvironmentConfigChain(client),
		DatasetPlan: NewResourceEnvironmentDatasetPlan(client),
		Index: actionIndex,
		List: actionIndex,
		Create: actionCreate,
		New: actionCreate,
		Show: actionShow,
		Find: actionShow,
		Update: actionUpdate,
		SetMaintenance: actionSetMaintenance,
	}
}
