package client

// Type for resource Export.Host
type ResourceExportHost struct {
	// Pointer to client
	Client *Client

	// Action Export.Host#Index
	Index *ActionExportHostIndex
	// Action Export.Host#Index
	List *ActionExportHostIndex
	// Action Export.Host#Show
	Show *ActionExportHostShow
	// Action Export.Host#Show
	Find *ActionExportHostShow
	// Action Export.Host#Create
	Create *ActionExportHostCreate
	// Action Export.Host#Create
	New *ActionExportHostCreate
	// Action Export.Host#Update
	Update *ActionExportHostUpdate
	// Action Export.Host#Delete
	Delete *ActionExportHostDelete
	// Action Export.Host#Delete
	Destroy *ActionExportHostDelete
}

func NewResourceExportHost(client *Client) *ResourceExportHost {
	actionIndex := NewActionExportHostIndex(client)
	actionShow := NewActionExportHostShow(client)
	actionCreate := NewActionExportHostCreate(client)
	actionUpdate := NewActionExportHostUpdate(client)
	actionDelete := NewActionExportHostDelete(client)

	return &ResourceExportHost{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Create: actionCreate,
		New: actionCreate,
		Update: actionUpdate,
		Delete: actionDelete,
		Destroy: actionDelete,
	}
}
