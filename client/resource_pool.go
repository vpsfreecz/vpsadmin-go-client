package client

// Type for resource Pool
type ResourcePool struct {
	// Pointer to client
	Client *Client

	// Action Pool#Create
	Create *ActionPoolCreate
	// Action Pool#Create
	New *ActionPoolCreate
	// Action Pool#Index
	Index *ActionPoolIndex
	// Action Pool#Index
	List *ActionPoolIndex
	// Action Pool#Set_maintenance
	SetMaintenance *ActionPoolSetMaintenance
	// Action Pool#Show
	Show *ActionPoolShow
	// Action Pool#Show
	Find *ActionPoolShow
}

func NewResourcePool(client *Client) *ResourcePool {
	actionCreate := NewActionPoolCreate(client)
	actionIndex := NewActionPoolIndex(client)
	actionSetMaintenance := NewActionPoolSetMaintenance(client)
	actionShow := NewActionPoolShow(client)

	return &ResourcePool{
		Client: client,
		Create: actionCreate,
		New: actionCreate,
		Index: actionIndex,
		List: actionIndex,
		SetMaintenance: actionSetMaintenance,
		Show: actionShow,
		Find: actionShow,
	}
}
