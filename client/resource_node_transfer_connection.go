package client

// Type for resource Node_transfer_connection
type ResourceNodeTransferConnection struct {
	// Pointer to client
	Client *Client

	// Action Node_transfer_connection#Create
	Create *ActionNodeTransferConnectionCreate
	// Action Node_transfer_connection#Create
	New *ActionNodeTransferConnectionCreate
	// Action Node_transfer_connection#Delete
	Delete *ActionNodeTransferConnectionDelete
	// Action Node_transfer_connection#Delete
	Destroy *ActionNodeTransferConnectionDelete
	// Action Node_transfer_connection#Index
	Index *ActionNodeTransferConnectionIndex
	// Action Node_transfer_connection#Index
	List *ActionNodeTransferConnectionIndex
	// Action Node_transfer_connection#Show
	Show *ActionNodeTransferConnectionShow
	// Action Node_transfer_connection#Show
	Find *ActionNodeTransferConnectionShow
	// Action Node_transfer_connection#Update
	Update *ActionNodeTransferConnectionUpdate
}

func NewResourceNodeTransferConnection(client *Client) *ResourceNodeTransferConnection {
	actionCreate := NewActionNodeTransferConnectionCreate(client)
	actionDelete := NewActionNodeTransferConnectionDelete(client)
	actionIndex := NewActionNodeTransferConnectionIndex(client)
	actionShow := NewActionNodeTransferConnectionShow(client)
	actionUpdate := NewActionNodeTransferConnectionUpdate(client)

	return &ResourceNodeTransferConnection{
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
