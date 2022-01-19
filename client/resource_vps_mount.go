package client

// Type for resource Vps.Mount
type ResourceVpsMount struct {
	// Pointer to client
	Client *Client

	// Action Vps.Mount#Create
	Create *ActionVpsMountCreate
	// Action Vps.Mount#Create
	New *ActionVpsMountCreate
	// Action Vps.Mount#Delete
	Delete *ActionVpsMountDelete
	// Action Vps.Mount#Delete
	Destroy *ActionVpsMountDelete
	// Action Vps.Mount#Index
	Index *ActionVpsMountIndex
	// Action Vps.Mount#Index
	List *ActionVpsMountIndex
	// Action Vps.Mount#Show
	Show *ActionVpsMountShow
	// Action Vps.Mount#Show
	Find *ActionVpsMountShow
	// Action Vps.Mount#Update
	Update *ActionVpsMountUpdate
}

func NewResourceVpsMount(client *Client) *ResourceVpsMount {
	actionCreate := NewActionVpsMountCreate(client)
	actionDelete := NewActionVpsMountDelete(client)
	actionIndex := NewActionVpsMountIndex(client)
	actionShow := NewActionVpsMountShow(client)
	actionUpdate := NewActionVpsMountUpdate(client)

	return &ResourceVpsMount{
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
