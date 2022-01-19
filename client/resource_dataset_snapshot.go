package client

// Type for resource Dataset.Snapshot
type ResourceDatasetSnapshot struct {
	// Pointer to client
	Client *Client

	// Action Dataset.Snapshot#Create
	Create *ActionDatasetSnapshotCreate
	// Action Dataset.Snapshot#Create
	New *ActionDatasetSnapshotCreate
	// Action Dataset.Snapshot#Delete
	Delete *ActionDatasetSnapshotDelete
	// Action Dataset.Snapshot#Delete
	Destroy *ActionDatasetSnapshotDelete
	// Action Dataset.Snapshot#Index
	Index *ActionDatasetSnapshotIndex
	// Action Dataset.Snapshot#Index
	List *ActionDatasetSnapshotIndex
	// Action Dataset.Snapshot#Rollback
	Rollback *ActionDatasetSnapshotRollback
	// Action Dataset.Snapshot#Show
	Show *ActionDatasetSnapshotShow
	// Action Dataset.Snapshot#Show
	Find *ActionDatasetSnapshotShow
}

func NewResourceDatasetSnapshot(client *Client) *ResourceDatasetSnapshot {
	actionCreate := NewActionDatasetSnapshotCreate(client)
	actionDelete := NewActionDatasetSnapshotDelete(client)
	actionIndex := NewActionDatasetSnapshotIndex(client)
	actionRollback := NewActionDatasetSnapshotRollback(client)
	actionShow := NewActionDatasetSnapshotShow(client)

	return &ResourceDatasetSnapshot{
		Client:   client,
		Create:   actionCreate,
		New:      actionCreate,
		Delete:   actionDelete,
		Destroy:  actionDelete,
		Index:    actionIndex,
		List:     actionIndex,
		Rollback: actionRollback,
		Show:     actionShow,
		Find:     actionShow,
	}
}
