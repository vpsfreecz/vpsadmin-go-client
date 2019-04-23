package client

// Type for resource User_namespace_map.Entry
type ResourceUserNamespaceMapEntry struct {
	// Pointer to client
	Client *Client

	// Action User_namespace_map.Entry#Index
	Index *ActionUserNamespaceMapEntryIndex
	// Action User_namespace_map.Entry#Index
	List *ActionUserNamespaceMapEntryIndex
	// Action User_namespace_map.Entry#Show
	Show *ActionUserNamespaceMapEntryShow
	// Action User_namespace_map.Entry#Show
	Find *ActionUserNamespaceMapEntryShow
	// Action User_namespace_map.Entry#Create
	Create *ActionUserNamespaceMapEntryCreate
	// Action User_namespace_map.Entry#Create
	New *ActionUserNamespaceMapEntryCreate
	// Action User_namespace_map.Entry#Update
	Update *ActionUserNamespaceMapEntryUpdate
	// Action User_namespace_map.Entry#Delete
	Delete *ActionUserNamespaceMapEntryDelete
	// Action User_namespace_map.Entry#Delete
	Destroy *ActionUserNamespaceMapEntryDelete
}

func NewResourceUserNamespaceMapEntry(client *Client) *ResourceUserNamespaceMapEntry {
	actionIndex := NewActionUserNamespaceMapEntryIndex(client)
	actionShow := NewActionUserNamespaceMapEntryShow(client)
	actionCreate := NewActionUserNamespaceMapEntryCreate(client)
	actionUpdate := NewActionUserNamespaceMapEntryUpdate(client)
	actionDelete := NewActionUserNamespaceMapEntryDelete(client)

	return &ResourceUserNamespaceMapEntry{
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
