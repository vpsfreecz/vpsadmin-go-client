package client

// Type for resource Dataset_expansion
type ResourceDatasetExpansion struct {
	// Pointer to client
	Client *Client

	// Resource Dataset_expansion.History
	History *ResourceDatasetExpansionHistory
	// Action Dataset_expansion#Create
	Create *ActionDatasetExpansionCreate
	// Action Dataset_expansion#Create
	New *ActionDatasetExpansionCreate
	// Action Dataset_expansion#Index
	Index *ActionDatasetExpansionIndex
	// Action Dataset_expansion#Index
	List *ActionDatasetExpansionIndex
	// Action Dataset_expansion#Register_expanded
	RegisterExpanded *ActionDatasetExpansionRegisterExpanded
	// Action Dataset_expansion#Show
	Show *ActionDatasetExpansionShow
	// Action Dataset_expansion#Show
	Find *ActionDatasetExpansionShow
	// Action Dataset_expansion#Update
	Update *ActionDatasetExpansionUpdate
}

func NewResourceDatasetExpansion(client *Client) *ResourceDatasetExpansion {
	actionCreate := NewActionDatasetExpansionCreate(client)
	actionIndex := NewActionDatasetExpansionIndex(client)
	actionRegisterExpanded := NewActionDatasetExpansionRegisterExpanded(client)
	actionShow := NewActionDatasetExpansionShow(client)
	actionUpdate := NewActionDatasetExpansionUpdate(client)

	return &ResourceDatasetExpansion{
		Client:           client,
		History:          NewResourceDatasetExpansionHistory(client),
		Create:           actionCreate,
		New:              actionCreate,
		Index:            actionIndex,
		List:             actionIndex,
		RegisterExpanded: actionRegisterExpanded,
		Show:             actionShow,
		Find:             actionShow,
		Update:           actionUpdate,
	}
}
