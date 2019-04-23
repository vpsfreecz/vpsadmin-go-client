package client

// Type for resource Outage.Entity
type ResourceOutageEntity struct {
	// Pointer to client
	Client *Client

	// Action Outage.Entity#Index
	Index *ActionOutageEntityIndex
	// Action Outage.Entity#Index
	List *ActionOutageEntityIndex
	// Action Outage.Entity#Show
	Show *ActionOutageEntityShow
	// Action Outage.Entity#Show
	Find *ActionOutageEntityShow
	// Action Outage.Entity#Create
	Create *ActionOutageEntityCreate
	// Action Outage.Entity#Create
	New *ActionOutageEntityCreate
	// Action Outage.Entity#Delete
	Delete *ActionOutageEntityDelete
	// Action Outage.Entity#Delete
	Destroy *ActionOutageEntityDelete
}

func NewResourceOutageEntity(client *Client) *ResourceOutageEntity {
	actionIndex := NewActionOutageEntityIndex(client)
	actionShow := NewActionOutageEntityShow(client)
	actionCreate := NewActionOutageEntityCreate(client)
	actionDelete := NewActionOutageEntityDelete(client)

	return &ResourceOutageEntity{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Create: actionCreate,
		New: actionCreate,
		Delete: actionDelete,
		Destroy: actionDelete,
	}
}
