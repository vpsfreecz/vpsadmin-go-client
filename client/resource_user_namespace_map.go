package client

// Type for resource User_namespace_map
type ResourceUserNamespaceMap struct {
	// Pointer to client
	Client *Client

	// Resource User_namespace_map.Entry
	Entry *ResourceUserNamespaceMapEntry
	// Action User_namespace_map#Index
	Index *ActionUserNamespaceMapIndex
	// Action User_namespace_map#Index
	List *ActionUserNamespaceMapIndex
	// Action User_namespace_map#Show
	Show *ActionUserNamespaceMapShow
	// Action User_namespace_map#Show
	Find *ActionUserNamespaceMapShow
	// Action User_namespace_map#Create
	Create *ActionUserNamespaceMapCreate
	// Action User_namespace_map#Create
	New *ActionUserNamespaceMapCreate
	// Action User_namespace_map#Update
	Update *ActionUserNamespaceMapUpdate
	// Action User_namespace_map#Delete
	Delete *ActionUserNamespaceMapDelete
	// Action User_namespace_map#Delete
	Destroy *ActionUserNamespaceMapDelete
}

func NewResourceUserNamespaceMap(client *Client) *ResourceUserNamespaceMap {
	actionIndex := NewActionUserNamespaceMapIndex(client)
	actionShow := NewActionUserNamespaceMapShow(client)
	actionCreate := NewActionUserNamespaceMapCreate(client)
	actionUpdate := NewActionUserNamespaceMapUpdate(client)
	actionDelete := NewActionUserNamespaceMapDelete(client)

	return &ResourceUserNamespaceMap{
		Client: client,
		Entry: NewResourceUserNamespaceMapEntry(client),
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
