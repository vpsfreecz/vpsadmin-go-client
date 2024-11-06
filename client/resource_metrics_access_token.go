package client

// Type for resource Metrics_access_token
type ResourceMetricsAccessToken struct {
	// Pointer to client
	Client *Client

	// Action Metrics_access_token#Create
	Create *ActionMetricsAccessTokenCreate
	// Action Metrics_access_token#Create
	New *ActionMetricsAccessTokenCreate
	// Action Metrics_access_token#Delete
	Delete *ActionMetricsAccessTokenDelete
	// Action Metrics_access_token#Delete
	Destroy *ActionMetricsAccessTokenDelete
	// Action Metrics_access_token#Index
	Index *ActionMetricsAccessTokenIndex
	// Action Metrics_access_token#Index
	List *ActionMetricsAccessTokenIndex
	// Action Metrics_access_token#Show
	Show *ActionMetricsAccessTokenShow
	// Action Metrics_access_token#Show
	Find *ActionMetricsAccessTokenShow
}

func NewResourceMetricsAccessToken(client *Client) *ResourceMetricsAccessToken {
	actionCreate := NewActionMetricsAccessTokenCreate(client)
	actionDelete := NewActionMetricsAccessTokenDelete(client)
	actionIndex := NewActionMetricsAccessTokenIndex(client)
	actionShow := NewActionMetricsAccessTokenShow(client)

	return &ResourceMetricsAccessToken{
		Client:  client,
		Create:  actionCreate,
		New:     actionCreate,
		Delete:  actionDelete,
		Destroy: actionDelete,
		Index:   actionIndex,
		List:    actionIndex,
		Show:    actionShow,
		Find:    actionShow,
	}
}
