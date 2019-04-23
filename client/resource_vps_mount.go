package client

// Type for resource Vps.Mount
type ResourceVpsMount struct {
	// Pointer to client
	Client *Client

	// Action Vps.Mount#Index
	Index *ActionVpsMountIndex
	// Action Vps.Mount#Index
	List *ActionVpsMountIndex
	// Action Vps.Mount#Show
	Show *ActionVpsMountShow
	// Action Vps.Mount#Show
	Find *ActionVpsMountShow
	// Action Vps.Mount#Create
	Create *ActionVpsMountCreate
	// Action Vps.Mount#Create
	New *ActionVpsMountCreate
	// Action Vps.Mount#Update
	Update *ActionVpsMountUpdate
	// Action Vps.Mount#Delete
	Delete *ActionVpsMountDelete
	// Action Vps.Mount#Delete
	Destroy *ActionVpsMountDelete
}

func NewResourceVpsMount(client *Client) *ResourceVpsMount {
	actionIndex := NewActionVpsMountIndex(client)
	actionShow := NewActionVpsMountShow(client)
	actionCreate := NewActionVpsMountCreate(client)
	actionUpdate := NewActionVpsMountUpdate(client)
	actionDelete := NewActionVpsMountDelete(client)

	return &ResourceVpsMount{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Create: actionCreate,
		New: actionCreate,
		Update: actionUpdate,
		Delete: actionDelete,
		Destroy: actionDelete,
	}
}
