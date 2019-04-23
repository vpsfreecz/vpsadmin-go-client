package client

// Type for resource Object_history
type ResourceObjectHistory struct {
	// Pointer to client
	Client *Client

	// Action Object_history#Index
	Index *ActionObjectHistoryIndex
	// Action Object_history#Index
	List *ActionObjectHistoryIndex
	// Action Object_history#Show
	Show *ActionObjectHistoryShow
	// Action Object_history#Show
	Find *ActionObjectHistoryShow
}

func NewResourceObjectHistory(client *Client) *ResourceObjectHistory {
	actionIndex := NewActionObjectHistoryIndex(client)
	actionShow := NewActionObjectHistoryShow(client)

	return &ResourceObjectHistory{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
	}
}
