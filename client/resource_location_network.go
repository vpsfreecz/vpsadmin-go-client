package client

// Type for resource Location_network
type ResourceLocationNetwork struct {
	// Pointer to client
	Client *Client

	// Action Location_network#Create
	Create *ActionLocationNetworkCreate
	// Action Location_network#Create
	New *ActionLocationNetworkCreate
	// Action Location_network#Delete
	Delete *ActionLocationNetworkDelete
	// Action Location_network#Delete
	Destroy *ActionLocationNetworkDelete
	// Action Location_network#Index
	Index *ActionLocationNetworkIndex
	// Action Location_network#Index
	List *ActionLocationNetworkIndex
	// Action Location_network#Show
	Show *ActionLocationNetworkShow
	// Action Location_network#Show
	Find *ActionLocationNetworkShow
	// Action Location_network#Update
	Update *ActionLocationNetworkUpdate
}

func NewResourceLocationNetwork(client *Client) *ResourceLocationNetwork {
	actionCreate := NewActionLocationNetworkCreate(client)
	actionDelete := NewActionLocationNetworkDelete(client)
	actionIndex := NewActionLocationNetworkIndex(client)
	actionShow := NewActionLocationNetworkShow(client)
	actionUpdate := NewActionLocationNetworkUpdate(client)

	return &ResourceLocationNetwork{
		Client:  client,
		Create:  actionCreate,
		New:     actionCreate,
		Delete:  actionDelete,
		Destroy: actionDelete,
		Index:   actionIndex,
		List:    actionIndex,
		Show:    actionShow,
		Find:    actionShow,
		Update:  actionUpdate,
	}
}
