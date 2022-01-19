package client

// Type for resource Export
type ResourceExport struct {
	// Pointer to client
	Client *Client

	// Resource Export.Host
	Host *ResourceExportHost
	// Action Export#Create
	Create *ActionExportCreate
	// Action Export#Create
	New *ActionExportCreate
	// Action Export#Delete
	Delete *ActionExportDelete
	// Action Export#Delete
	Destroy *ActionExportDelete
	// Action Export#Index
	Index *ActionExportIndex
	// Action Export#Index
	List *ActionExportIndex
	// Action Export#Show
	Show *ActionExportShow
	// Action Export#Show
	Find *ActionExportShow
	// Action Export#Update
	Update *ActionExportUpdate
}

func NewResourceExport(client *Client) *ResourceExport {
	actionCreate := NewActionExportCreate(client)
	actionDelete := NewActionExportDelete(client)
	actionIndex := NewActionExportIndex(client)
	actionShow := NewActionExportShow(client)
	actionUpdate := NewActionExportUpdate(client)

	return &ResourceExport{
		Client:  client,
		Host:    NewResourceExportHost(client),
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
