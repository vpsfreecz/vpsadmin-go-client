package client

// Type for resource Security_advisory_update
type ResourceSecurityAdvisoryUpdate struct {
	// Pointer to client
	Client *Client

	// Action Security_advisory_update#Create
	Create *ActionSecurityAdvisoryUpdateCreate
	// Action Security_advisory_update#Create
	New *ActionSecurityAdvisoryUpdateCreate
	// Action Security_advisory_update#Delete
	Delete *ActionSecurityAdvisoryUpdateDelete
	// Action Security_advisory_update#Delete
	Destroy *ActionSecurityAdvisoryUpdateDelete
	// Action Security_advisory_update#Index
	Index *ActionSecurityAdvisoryUpdateIndex
	// Action Security_advisory_update#Index
	List *ActionSecurityAdvisoryUpdateIndex
	// Action Security_advisory_update#Show
	Show *ActionSecurityAdvisoryUpdateShow
	// Action Security_advisory_update#Show
	Find *ActionSecurityAdvisoryUpdateShow
	// Action Security_advisory_update#Update
	Update *ActionSecurityAdvisoryUpdateUpdate
}

func NewResourceSecurityAdvisoryUpdate(client *Client) *ResourceSecurityAdvisoryUpdate {
	actionCreate := NewActionSecurityAdvisoryUpdateCreate(client)
	actionDelete := NewActionSecurityAdvisoryUpdateDelete(client)
	actionIndex := NewActionSecurityAdvisoryUpdateIndex(client)
	actionShow := NewActionSecurityAdvisoryUpdateShow(client)
	actionUpdate := NewActionSecurityAdvisoryUpdateUpdate(client)

	return &ResourceSecurityAdvisoryUpdate{
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
