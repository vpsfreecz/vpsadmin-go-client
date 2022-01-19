package client

// Type for resource Dataset.Property_history
type ResourceDatasetPropertyHistory struct {
	// Pointer to client
	Client *Client

	// Action Dataset.Property_history#Index
	Index *ActionDatasetPropertyHistoryIndex
	// Action Dataset.Property_history#Index
	List *ActionDatasetPropertyHistoryIndex
	// Action Dataset.Property_history#Show
	Show *ActionDatasetPropertyHistoryShow
	// Action Dataset.Property_history#Show
	Find *ActionDatasetPropertyHistoryShow
}

func NewResourceDatasetPropertyHistory(client *Client) *ResourceDatasetPropertyHistory {
	actionIndex := NewActionDatasetPropertyHistoryIndex(client)
	actionShow := NewActionDatasetPropertyHistoryShow(client)

	return &ResourceDatasetPropertyHistory{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
