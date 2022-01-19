package client

// Type for resource Export.Host
type ResourceExportHost struct {
	// Pointer to client
	Client *Client

	// Action Export.Host#Create
	Create *ActionExportHostCreate
	// Action Export.Host#Create
	New *ActionExportHostCreate
	// Action Export.Host#Delete
	Delete *ActionExportHostDelete
	// Action Export.Host#Delete
	Destroy *ActionExportHostDelete
	// Action Export.Host#Index
	Index *ActionExportHostIndex
	// Action Export.Host#Index
	List *ActionExportHostIndex
	// Action Export.Host#Show
	Show *ActionExportHostShow
	// Action Export.Host#Show
	Find *ActionExportHostShow
	// Action Export.Host#Update
	Update *ActionExportHostUpdate
}

func NewResourceExportHost(client *Client) *ResourceExportHost {
	actionCreate := NewActionExportHostCreate(client)
	actionDelete := NewActionExportHostDelete(client)
	actionIndex := NewActionExportHostIndex(client)
	actionShow := NewActionExportHostShow(client)
	actionUpdate := NewActionExportHostUpdate(client)

	return &ResourceExportHost{
		Client:  client,
		Create:  actionCreate,
		New:     actionCreate,
		Delete:  actionDelete,
		Destroy: actionDelete,
		Index:   actionIndex,
		List:    actionIndex,
		Show:    actionShow,
		Find:    actionShow,
		Update:  actionUpdate,
	}
}
