package client

// Type for resource Network
type ResourceNetwork struct {
	// Pointer to client
	Client *Client

	// Action Network#Index
	Index *ActionNetworkIndex
	// Action Network#Index
	List *ActionNetworkIndex
	// Action Network#Show
	Show *ActionNetworkShow
	// Action Network#Show
	Find *ActionNetworkShow
	// Action Network#Create
	Create *ActionNetworkCreate
	// Action Network#Create
	New *ActionNetworkCreate
	// Action Network#Update
	Update *ActionNetworkUpdate
	// Action Network#Add_addresses
	AddAddresses *ActionNetworkAddAddresses
}

func NewResourceNetwork(client *Client) *ResourceNetwork {
	actionIndex := NewActionNetworkIndex(client)
	actionShow := NewActionNetworkShow(client)
	actionCreate := NewActionNetworkCreate(client)
	actionUpdate := NewActionNetworkUpdate(client)
	actionAddAddresses := NewActionNetworkAddAddresses(client)

	return &ResourceNetwork{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Create: actionCreate,
		New: actionCreate,
		Update: actionUpdate,
		AddAddresses: actionAddAddresses,
	}
}
