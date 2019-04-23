package client

// Type for resource Dataset.Snapshot
type ResourceDatasetSnapshot struct {
	// Pointer to client
	Client *Client

	// Action Dataset.Snapshot#Index
	Index *ActionDatasetSnapshotIndex
	// Action Dataset.Snapshot#Index
	List *ActionDatasetSnapshotIndex
	// Action Dataset.Snapshot#Show
	Show *ActionDatasetSnapshotShow
	// Action Dataset.Snapshot#Show
	Find *ActionDatasetSnapshotShow
	// Action Dataset.Snapshot#Create
	Create *ActionDatasetSnapshotCreate
	// Action Dataset.Snapshot#Create
	New *ActionDatasetSnapshotCreate
	// Action Dataset.Snapshot#Delete
	Delete *ActionDatasetSnapshotDelete
	// Action Dataset.Snapshot#Delete
	Destroy *ActionDatasetSnapshotDelete
	// Action Dataset.Snapshot#Rollback
	Rollback *ActionDatasetSnapshotRollback
}

func NewResourceDatasetSnapshot(client *Client) *ResourceDatasetSnapshot {
	actionIndex := NewActionDatasetSnapshotIndex(client)
	actionShow := NewActionDatasetSnapshotShow(client)
	actionCreate := NewActionDatasetSnapshotCreate(client)
	actionDelete := NewActionDatasetSnapshotDelete(client)
	actionRollback := NewActionDatasetSnapshotRollback(client)

	return &ResourceDatasetSnapshot{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Create: actionCreate,
		New: actionCreate,
		Delete: actionDelete,
		Destroy: actionDelete,
		Rollback: actionRollback,
	}
}
