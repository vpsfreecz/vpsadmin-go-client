package client

// Type for resource Dataset_plan
type ResourceDatasetPlan struct {
	// Pointer to client
	Client *Client

	// Action Dataset_plan#Index
	Index *ActionDatasetPlanIndex
	// Action Dataset_plan#Index
	List *ActionDatasetPlanIndex
	// Action Dataset_plan#Show
	Show *ActionDatasetPlanShow
	// Action Dataset_plan#Show
	Find *ActionDatasetPlanShow
}

func NewResourceDatasetPlan(client *Client) *ResourceDatasetPlan {
	actionIndex := NewActionDatasetPlanIndex(client)
	actionShow := NewActionDatasetPlanShow(client)

	return &ResourceDatasetPlan{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
	}
}
