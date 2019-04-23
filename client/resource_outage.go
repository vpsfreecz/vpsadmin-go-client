package client

// Type for resource Outage
type ResourceOutage struct {
	// Pointer to client
	Client *Client

	// Resource Outage.Entity
	Entity *ResourceOutageEntity
	// Resource Outage.Handler
	Handler *ResourceOutageHandler
	// Action Outage#Index
	Index *ActionOutageIndex
	// Action Outage#Index
	List *ActionOutageIndex
	// Action Outage#Show
	Show *ActionOutageShow
	// Action Outage#Show
	Find *ActionOutageShow
	// Action Outage#Create
	Create *ActionOutageCreate
	// Action Outage#Create
	New *ActionOutageCreate
	// Action Outage#Update
	Update *ActionOutageUpdate
	// Action Outage#Rebuild_affected_vps
	RebuildAffectedVps *ActionOutageRebuildAffectedVps
}

func NewResourceOutage(client *Client) *ResourceOutage {
	actionIndex := NewActionOutageIndex(client)
	actionShow := NewActionOutageShow(client)
	actionCreate := NewActionOutageCreate(client)
	actionUpdate := NewActionOutageUpdate(client)
	actionRebuildAffectedVps := NewActionOutageRebuildAffectedVps(client)

	return &ResourceOutage{
		Client: client,
		Entity: NewResourceOutageEntity(client),
		Handler: NewResourceOutageHandler(client),
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Create: actionCreate,
		New: actionCreate,
		Update: actionUpdate,
		RebuildAffectedVps: actionRebuildAffectedVps,
	}
}
