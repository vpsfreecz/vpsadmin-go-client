package client

// Type for resource Dataset
type ResourceDataset struct {
	// Pointer to client
	Client *Client

	// Resource Dataset.Plan
	Plan *ResourceDatasetPlan
	// Resource Dataset.Property_history
	PropertyHistory *ResourceDatasetPropertyHistory
	// Resource Dataset.Snapshot
	Snapshot *ResourceDatasetSnapshot
	// Action Dataset#Create
	Create *ActionDatasetCreate
	// Action Dataset#Create
	New *ActionDatasetCreate
	// Action Dataset#Delete
	Delete *ActionDatasetDelete
	// Action Dataset#Delete
	Destroy *ActionDatasetDelete
	// Action Dataset#Index
	Index *ActionDatasetIndex
	// Action Dataset#Index
	List *ActionDatasetIndex
	// Action Dataset#Inherit
	Inherit *ActionDatasetInherit
	// Action Dataset#Show
	Show *ActionDatasetShow
	// Action Dataset#Show
	Find *ActionDatasetShow
	// Action Dataset#Update
	Update *ActionDatasetUpdate
}

func NewResourceDataset(client *Client) *ResourceDataset {
	actionCreate := NewActionDatasetCreate(client)
	actionDelete := NewActionDatasetDelete(client)
	actionIndex := NewActionDatasetIndex(client)
	actionInherit := NewActionDatasetInherit(client)
	actionShow := NewActionDatasetShow(client)
	actionUpdate := NewActionDatasetUpdate(client)

	return &ResourceDataset{
		Client: client,
		Plan: NewResourceDatasetPlan(client),
		PropertyHistory: NewResourceDatasetPropertyHistory(client),
		Snapshot: NewResourceDatasetSnapshot(client),
		Create: actionCreate,
		New: actionCreate,
		Delete: actionDelete,
		Destroy: actionDelete,
		Index: actionIndex,
		List: actionIndex,
		Inherit: actionInherit,
		Show: actionShow,
		Find: actionShow,
		Update: actionUpdate,
	}
}
