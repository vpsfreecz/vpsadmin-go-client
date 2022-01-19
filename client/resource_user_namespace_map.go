package client

// Type for resource User_namespace_map
type ResourceUserNamespaceMap struct {
	// Pointer to client
	Client *Client

	// Resource User_namespace_map.Entry
	Entry *ResourceUserNamespaceMapEntry
	// Action User_namespace_map#Create
	Create *ActionUserNamespaceMapCreate
	// Action User_namespace_map#Create
	New *ActionUserNamespaceMapCreate
	// Action User_namespace_map#Delete
	Delete *ActionUserNamespaceMapDelete
	// Action User_namespace_map#Delete
	Destroy *ActionUserNamespaceMapDelete
	// Action User_namespace_map#Index
	Index *ActionUserNamespaceMapIndex
	// Action User_namespace_map#Index
	List *ActionUserNamespaceMapIndex
	// Action User_namespace_map#Show
	Show *ActionUserNamespaceMapShow
	// Action User_namespace_map#Show
	Find *ActionUserNamespaceMapShow
	// Action User_namespace_map#Update
	Update *ActionUserNamespaceMapUpdate
}

func NewResourceUserNamespaceMap(client *Client) *ResourceUserNamespaceMap {
	actionCreate := NewActionUserNamespaceMapCreate(client)
	actionDelete := NewActionUserNamespaceMapDelete(client)
	actionIndex := NewActionUserNamespaceMapIndex(client)
	actionShow := NewActionUserNamespaceMapShow(client)
	actionUpdate := NewActionUserNamespaceMapUpdate(client)

	return &ResourceUserNamespaceMap{
		Client:  client,
		Entry:   NewResourceUserNamespaceMapEntry(client),
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
