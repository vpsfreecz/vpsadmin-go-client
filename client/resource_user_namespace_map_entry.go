package client

// Type for resource User_namespace_map.Entry
type ResourceUserNamespaceMapEntry struct {
	// Pointer to client
	Client *Client

	// Action User_namespace_map.Entry#Create
	Create *ActionUserNamespaceMapEntryCreate
	// Action User_namespace_map.Entry#Create
	New *ActionUserNamespaceMapEntryCreate
	// Action User_namespace_map.Entry#Delete
	Delete *ActionUserNamespaceMapEntryDelete
	// Action User_namespace_map.Entry#Delete
	Destroy *ActionUserNamespaceMapEntryDelete
	// Action User_namespace_map.Entry#Index
	Index *ActionUserNamespaceMapEntryIndex
	// Action User_namespace_map.Entry#Index
	List *ActionUserNamespaceMapEntryIndex
	// Action User_namespace_map.Entry#Show
	Show *ActionUserNamespaceMapEntryShow
	// Action User_namespace_map.Entry#Show
	Find *ActionUserNamespaceMapEntryShow
	// Action User_namespace_map.Entry#Update
	Update *ActionUserNamespaceMapEntryUpdate
}

func NewResourceUserNamespaceMapEntry(client *Client) *ResourceUserNamespaceMapEntry {
	actionCreate := NewActionUserNamespaceMapEntryCreate(client)
	actionDelete := NewActionUserNamespaceMapEntryDelete(client)
	actionIndex := NewActionUserNamespaceMapEntryIndex(client)
	actionShow := NewActionUserNamespaceMapEntryShow(client)
	actionUpdate := NewActionUserNamespaceMapEntryUpdate(client)

	return &ResourceUserNamespaceMapEntry{
		Client: client,
		Create: actionCreate,
		New: actionCreate,
		Delete: actionDelete,
		Destroy: actionDelete,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Update: actionUpdate,
	}
}
