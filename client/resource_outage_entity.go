package client

// Type for resource Outage.Entity
type ResourceOutageEntity struct {
	// Pointer to client
	Client *Client

	// Action Outage.Entity#Create
	Create *ActionOutageEntityCreate
	// Action Outage.Entity#Create
	New *ActionOutageEntityCreate
	// Action Outage.Entity#Delete
	Delete *ActionOutageEntityDelete
	// Action Outage.Entity#Delete
	Destroy *ActionOutageEntityDelete
	// Action Outage.Entity#Index
	Index *ActionOutageEntityIndex
	// Action Outage.Entity#Index
	List *ActionOutageEntityIndex
	// Action Outage.Entity#Show
	Show *ActionOutageEntityShow
	// Action Outage.Entity#Show
	Find *ActionOutageEntityShow
}

func NewResourceOutageEntity(client *Client) *ResourceOutageEntity {
	actionCreate := NewActionOutageEntityCreate(client)
	actionDelete := NewActionOutageEntityDelete(client)
	actionIndex := NewActionOutageEntityIndex(client)
	actionShow := NewActionOutageEntityShow(client)

	return &ResourceOutageEntity{
		Client:  client,
		Create:  actionCreate,
		New:     actionCreate,
		Delete:  actionDelete,
		Destroy: actionDelete,
		Index:   actionIndex,
		List:    actionIndex,
		Show:    actionShow,
		Find:    actionShow,
	}
}
