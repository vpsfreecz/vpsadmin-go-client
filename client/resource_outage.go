package client

// Type for resource Outage
type ResourceOutage struct {
	// Pointer to client
	Client *Client

	// Resource Outage.Entity
	Entity *ResourceOutageEntity
	// Resource Outage.Handler
	Handler *ResourceOutageHandler
	// Action Outage#Create
	Create *ActionOutageCreate
	// Action Outage#Create
	New *ActionOutageCreate
	// Action Outage#Index
	Index *ActionOutageIndex
	// Action Outage#Index
	List *ActionOutageIndex
	// Action Outage#Rebuild_affected_vps
	RebuildAffectedVps *ActionOutageRebuildAffectedVps
	// Action Outage#Show
	Show *ActionOutageShow
	// Action Outage#Show
	Find *ActionOutageShow
	// Action Outage#Update
	Update *ActionOutageUpdate
}

func NewResourceOutage(client *Client) *ResourceOutage {
	actionCreate := NewActionOutageCreate(client)
	actionIndex := NewActionOutageIndex(client)
	actionRebuildAffectedVps := NewActionOutageRebuildAffectedVps(client)
	actionShow := NewActionOutageShow(client)
	actionUpdate := NewActionOutageUpdate(client)

	return &ResourceOutage{
		Client: client,
		Entity: NewResourceOutageEntity(client),
		Handler: NewResourceOutageHandler(client),
		Create: actionCreate,
		New: actionCreate,
		Index: actionIndex,
		List: actionIndex,
		RebuildAffectedVps: actionRebuildAffectedVps,
		Show: actionShow,
		Find: actionShow,
		Update: actionUpdate,
	}
}
