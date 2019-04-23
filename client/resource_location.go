package client

// Type for resource Location
type ResourceLocation struct {
	// Pointer to client
	Client *Client

	// Action Location#Index
	Index *ActionLocationIndex
	// Action Location#Index
	List *ActionLocationIndex
	// Action Location#Create
	Create *ActionLocationCreate
	// Action Location#Create
	New *ActionLocationCreate
	// Action Location#Show
	Show *ActionLocationShow
	// Action Location#Show
	Find *ActionLocationShow
	// Action Location#Update
	Update *ActionLocationUpdate
	// Action Location#Set_maintenance
	SetMaintenance *ActionLocationSetMaintenance
}

func NewResourceLocation(client *Client) *ResourceLocation {
	actionIndex := NewActionLocationIndex(client)
	actionCreate := NewActionLocationCreate(client)
	actionShow := NewActionLocationShow(client)
	actionUpdate := NewActionLocationUpdate(client)
	actionSetMaintenance := NewActionLocationSetMaintenance(client)

	return &ResourceLocation{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Create: actionCreate,
		New: actionCreate,
		Show: actionShow,
		Find: actionShow,
		Update: actionUpdate,
		SetMaintenance: actionSetMaintenance,
	}
}
