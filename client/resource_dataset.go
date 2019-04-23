package client

// Type for resource Dataset
type ResourceDataset struct {
	// Pointer to client
	Client *Client

	// Resource Dataset.Snapshot
	Snapshot *ResourceDatasetSnapshot
	// Resource Dataset.Plan
	Plan *ResourceDatasetPlan
	// Resource Dataset.Property_history
	PropertyHistory *ResourceDatasetPropertyHistory
	// Action Dataset#Index
	Index *ActionDatasetIndex
	// Action Dataset#Index
	List *ActionDatasetIndex
	// Action Dataset#Show
	Show *ActionDatasetShow
	// Action Dataset#Show
	Find *ActionDatasetShow
	// Action Dataset#Create
	Create *ActionDatasetCreate
	// Action Dataset#Create
	New *ActionDatasetCreate
	// Action Dataset#Update
	Update *ActionDatasetUpdate
	// Action Dataset#Delete
	Delete *ActionDatasetDelete
	// Action Dataset#Delete
	Destroy *ActionDatasetDelete
	// Action Dataset#Inherit
	Inherit *ActionDatasetInherit
}

func NewResourceDataset(client *Client) *ResourceDataset {
	actionIndex := NewActionDatasetIndex(client)
	actionShow := NewActionDatasetShow(client)
	actionCreate := NewActionDatasetCreate(client)
	actionUpdate := NewActionDatasetUpdate(client)
	actionDelete := NewActionDatasetDelete(client)
	actionInherit := NewActionDatasetInherit(client)

	return &ResourceDataset{
		Client: client,
		Snapshot: NewResourceDatasetSnapshot(client),
		Plan: NewResourceDatasetPlan(client),
		PropertyHistory: NewResourceDatasetPropertyHistory(client),
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Create: actionCreate,
		New: actionCreate,
		Update: actionUpdate,
		Delete: actionDelete,
		Destroy: actionDelete,
		Inherit: actionInherit,
	}
}
