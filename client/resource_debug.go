package client

// Type for resource Debug
type ResourceDebug struct {
	// Pointer to client
	Client *Client

	// Action Debug#Array_top
	ArrayTop *ActionDebugArrayTop
	// Action Debug#Hash_top
	HashTop *ActionDebugHashTop
	// Action Debug#List_object_counts
	ListObjectCounts *ActionDebugListObjectCounts
}

func NewResourceDebug(client *Client) *ResourceDebug {
	actionArrayTop := NewActionDebugArrayTop(client)
	actionHashTop := NewActionDebugHashTop(client)
	actionListObjectCounts := NewActionDebugListObjectCounts(client)

	return &ResourceDebug{
		Client:           client,
		ArrayTop:         actionArrayTop,
		HashTop:          actionHashTop,
		ListObjectCounts: actionListObjectCounts,
	}
}
