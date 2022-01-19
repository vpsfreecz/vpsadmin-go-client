package client

// Type for resource Network
type ResourceNetwork struct {
	// Pointer to client
	Client *Client

	// Action Network#Add_addresses
	AddAddresses *ActionNetworkAddAddresses
	// Action Network#Create
	Create *ActionNetworkCreate
	// Action Network#Create
	New *ActionNetworkCreate
	// Action Network#Index
	Index *ActionNetworkIndex
	// Action Network#Index
	List *ActionNetworkIndex
	// Action Network#Show
	Show *ActionNetworkShow
	// Action Network#Show
	Find *ActionNetworkShow
	// Action Network#Update
	Update *ActionNetworkUpdate
}

func NewResourceNetwork(client *Client) *ResourceNetwork {
	actionAddAddresses := NewActionNetworkAddAddresses(client)
	actionCreate := NewActionNetworkCreate(client)
	actionIndex := NewActionNetworkIndex(client)
	actionShow := NewActionNetworkShow(client)
	actionUpdate := NewActionNetworkUpdate(client)

	return &ResourceNetwork{
		Client:       client,
		AddAddresses: actionAddAddresses,
		Create:       actionCreate,
		New:          actionCreate,
		Index:        actionIndex,
		List:         actionIndex,
		Show:         actionShow,
		Find:         actionShow,
		Update:       actionUpdate,
	}
}
