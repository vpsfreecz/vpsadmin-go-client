package client

// Type for resource Export
type ResourceExport struct {
	// Pointer to client
	Client *Client

	// Resource Export.Host
	Host *ResourceExportHost
	// Action Export#Index
	Index *ActionExportIndex
	// Action Export#Index
	List *ActionExportIndex
	// Action Export#Show
	Show *ActionExportShow
	// Action Export#Show
	Find *ActionExportShow
	// Action Export#Create
	Create *ActionExportCreate
	// Action Export#Create
	New *ActionExportCreate
	// Action Export#Update
	Update *ActionExportUpdate
	// Action Export#Delete
	Delete *ActionExportDelete
	// Action Export#Delete
	Destroy *ActionExportDelete
}

func NewResourceExport(client *Client) *ResourceExport {
	actionIndex := NewActionExportIndex(client)
	actionShow := NewActionExportShow(client)
	actionCreate := NewActionExportCreate(client)
	actionUpdate := NewActionExportUpdate(client)
	actionDelete := NewActionExportDelete(client)

	return &ResourceExport{
		Client: client,
		Host: NewResourceExportHost(client),
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
