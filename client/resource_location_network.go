package client

// Type for resource Location_network
type ResourceLocationNetwork struct {
	// Pointer to client
	Client *Client

	// Action Location_network#Index
	Index *ActionLocationNetworkIndex
	// Action Location_network#Index
	List *ActionLocationNetworkIndex
	// Action Location_network#Show
	Show *ActionLocationNetworkShow
	// Action Location_network#Show
	Find *ActionLocationNetworkShow
	// Action Location_network#Create
	Create *ActionLocationNetworkCreate
	// Action Location_network#Create
	New *ActionLocationNetworkCreate
	// Action Location_network#Update
	Update *ActionLocationNetworkUpdate
	// Action Location_network#Delete
	Delete *ActionLocationNetworkDelete
	// Action Location_network#Delete
	Destroy *ActionLocationNetworkDelete
}

func NewResourceLocationNetwork(client *Client) *ResourceLocationNetwork {
	actionIndex := NewActionLocationNetworkIndex(client)
	actionShow := NewActionLocationNetworkShow(client)
	actionCreate := NewActionLocationNetworkCreate(client)
	actionUpdate := NewActionLocationNetworkUpdate(client)
	actionDelete := NewActionLocationNetworkDelete(client)

	return &ResourceLocationNetwork{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Create: actionCreate,
		New: actionCreate,
		Update: actionUpdate,
		Delete: actionDelete,
		Destroy: actionDelete,
	}
}
