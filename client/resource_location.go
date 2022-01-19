package client

// Type for resource Location
type ResourceLocation struct {
	// Pointer to client
	Client *Client

	// Action Location#Create
	Create *ActionLocationCreate
	// Action Location#Create
	New *ActionLocationCreate
	// Action Location#Index
	Index *ActionLocationIndex
	// Action Location#Index
	List *ActionLocationIndex
	// Action Location#Set_maintenance
	SetMaintenance *ActionLocationSetMaintenance
	// Action Location#Show
	Show *ActionLocationShow
	// Action Location#Show
	Find *ActionLocationShow
	// Action Location#Update
	Update *ActionLocationUpdate
}

func NewResourceLocation(client *Client) *ResourceLocation {
	actionCreate := NewActionLocationCreate(client)
	actionIndex := NewActionLocationIndex(client)
	actionSetMaintenance := NewActionLocationSetMaintenance(client)
	actionShow := NewActionLocationShow(client)
	actionUpdate := NewActionLocationUpdate(client)

	return &ResourceLocation{
		Client: client,
		Create: actionCreate,
		New: actionCreate,
		Index: actionIndex,
		List: actionIndex,
		SetMaintenance: actionSetMaintenance,
		Show: actionShow,
		Find: actionShow,
		Update: actionUpdate,
	}
}
