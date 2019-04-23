package client

// Type for resource Environment.Dataset_plan
type ResourceEnvironmentDatasetPlan struct {
	// Pointer to client
	Client *Client

	// Action Environment.Dataset_plan#Index
	Index *ActionEnvironmentDatasetPlanIndex
	// Action Environment.Dataset_plan#Index
	List *ActionEnvironmentDatasetPlanIndex
	// Action Environment.Dataset_plan#Show
	Show *ActionEnvironmentDatasetPlanShow
	// Action Environment.Dataset_plan#Show
	Find *ActionEnvironmentDatasetPlanShow
}

func NewResourceEnvironmentDatasetPlan(client *Client) *ResourceEnvironmentDatasetPlan {
	actionIndex := NewActionEnvironmentDatasetPlanIndex(client)
	actionShow := NewActionEnvironmentDatasetPlanShow(client)

	return &ResourceEnvironmentDatasetPlan{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
	}
}
