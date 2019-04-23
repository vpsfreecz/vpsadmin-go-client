package client

// Type for resource Vps_config
type ResourceVpsConfig struct {
	// Pointer to client
	Client *Client

	// Action Vps_config#Index
	Index *ActionVpsConfigIndex
	// Action Vps_config#Index
	List *ActionVpsConfigIndex
	// Action Vps_config#Create
	Create *ActionVpsConfigCreate
	// Action Vps_config#Create
	New *ActionVpsConfigCreate
	// Action Vps_config#Show
	Show *ActionVpsConfigShow
	// Action Vps_config#Show
	Find *ActionVpsConfigShow
	// Action Vps_config#Update
	Update *ActionVpsConfigUpdate
}

func NewResourceVpsConfig(client *Client) *ResourceVpsConfig {
	actionIndex := NewActionVpsConfigIndex(client)
	actionCreate := NewActionVpsConfigCreate(client)
	actionShow := NewActionVpsConfigShow(client)
	actionUpdate := NewActionVpsConfigUpdate(client)

	return &ResourceVpsConfig{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Create: actionCreate,
		New: actionCreate,
		Show: actionShow,
		Find: actionShow,
		Update: actionUpdate,
	}
}
