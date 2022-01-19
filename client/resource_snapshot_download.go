package client

// Type for resource Snapshot_download
type ResourceSnapshotDownload struct {
	// Pointer to client
	Client *Client

	// Action Snapshot_download#Create
	Create *ActionSnapshotDownloadCreate
	// Action Snapshot_download#Create
	New *ActionSnapshotDownloadCreate
	// Action Snapshot_download#Delete
	Delete *ActionSnapshotDownloadDelete
	// Action Snapshot_download#Delete
	Destroy *ActionSnapshotDownloadDelete
	// Action Snapshot_download#Index
	Index *ActionSnapshotDownloadIndex
	// Action Snapshot_download#Index
	List *ActionSnapshotDownloadIndex
	// Action Snapshot_download#Show
	Show *ActionSnapshotDownloadShow
	// Action Snapshot_download#Show
	Find *ActionSnapshotDownloadShow
}

func NewResourceSnapshotDownload(client *Client) *ResourceSnapshotDownload {
	actionCreate := NewActionSnapshotDownloadCreate(client)
	actionDelete := NewActionSnapshotDownloadDelete(client)
	actionIndex := NewActionSnapshotDownloadIndex(client)
	actionShow := NewActionSnapshotDownloadShow(client)

	return &ResourceSnapshotDownload{
		Client:  client,
		Create:  actionCreate,
		New:     actionCreate,
		Delete:  actionDelete,
		Destroy: actionDelete,
		Index:   actionIndex,
		List:    actionIndex,
		Show:    actionShow,
		Find:    actionShow,
	}
}
