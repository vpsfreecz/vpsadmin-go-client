package client

// Type for resource Vps.Feature
type ResourceVpsFeature struct {
	// Pointer to client
	Client *Client

	// Action Vps.Feature#Index
	Index *ActionVpsFeatureIndex
	// Action Vps.Feature#Index
	List *ActionVpsFeatureIndex
	// Action Vps.Feature#Show
	Show *ActionVpsFeatureShow
	// Action Vps.Feature#Show
	Find *ActionVpsFeatureShow
	// Action Vps.Feature#Update
	Update *ActionVpsFeatureUpdate
	// Action Vps.Feature#Update_all
	UpdateAll *ActionVpsFeatureUpdateAll
}

func NewResourceVpsFeature(client *Client) *ResourceVpsFeature {
	actionIndex := NewActionVpsFeatureIndex(client)
	actionShow := NewActionVpsFeatureShow(client)
	actionUpdate := NewActionVpsFeatureUpdate(client)
	actionUpdateAll := NewActionVpsFeatureUpdateAll(client)

	return &ResourceVpsFeature{
		Client:    client,
		Index:     actionIndex,
		List:      actionIndex,
		Show:      actionShow,
		Find:      actionShow,
		Update:    actionUpdate,
		UpdateAll: actionUpdateAll,
	}
}
