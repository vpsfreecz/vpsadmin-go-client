package client

// Type for resource Dataset_expansion.History
type ResourceDatasetExpansionHistory struct {
	// Pointer to client
	Client *Client

	// Action Dataset_expansion.History#Create
	Create *ActionDatasetExpansionHistoryCreate
	// Action Dataset_expansion.History#Create
	New *ActionDatasetExpansionHistoryCreate
	// Action Dataset_expansion.History#Index
	Index *ActionDatasetExpansionHistoryIndex
	// Action Dataset_expansion.History#Index
	List *ActionDatasetExpansionHistoryIndex
	// Action Dataset_expansion.History#Show
	Show *ActionDatasetExpansionHistoryShow
	// Action Dataset_expansion.History#Show
	Find *ActionDatasetExpansionHistoryShow
}

func NewResourceDatasetExpansionHistory(client *Client) *ResourceDatasetExpansionHistory {
	actionCreate := NewActionDatasetExpansionHistoryCreate(client)
	actionIndex := NewActionDatasetExpansionHistoryIndex(client)
	actionShow := NewActionDatasetExpansionHistoryShow(client)

	return &ResourceDatasetExpansionHistory{
		Client: client,
		Create: actionCreate,
		New:    actionCreate,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
