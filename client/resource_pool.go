package client

// Type for resource Pool
type ResourcePool struct {
	// Pointer to client
	Client *Client

	// Action Pool#Index
	Index *ActionPoolIndex
	// Action Pool#Index
	List *ActionPoolIndex
	// Action Pool#Show
	Show *ActionPoolShow
	// Action Pool#Show
	Find *ActionPoolShow
	// Action Pool#Create
	Create *ActionPoolCreate
	// Action Pool#Create
	New *ActionPoolCreate
	// Action Pool#Set_maintenance
	SetMaintenance *ActionPoolSetMaintenance
}

func NewResourcePool(client *Client) *ResourcePool {
	actionIndex := NewActionPoolIndex(client)
	actionShow := NewActionPoolShow(client)
	actionCreate := NewActionPoolCreate(client)
	actionSetMaintenance := NewActionPoolSetMaintenance(client)

	return &ResourcePool{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Create: actionCreate,
		New: actionCreate,
		SetMaintenance: actionSetMaintenance,
	}
}
