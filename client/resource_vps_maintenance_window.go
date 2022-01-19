package client

// Type for resource Vps.Maintenance_window
type ResourceVpsMaintenanceWindow struct {
	// Pointer to client
	Client *Client

	// Action Vps.Maintenance_window#Index
	Index *ActionVpsMaintenanceWindowIndex
	// Action Vps.Maintenance_window#Index
	List *ActionVpsMaintenanceWindowIndex
	// Action Vps.Maintenance_window#Show
	Show *ActionVpsMaintenanceWindowShow
	// Action Vps.Maintenance_window#Show
	Find *ActionVpsMaintenanceWindowShow
	// Action Vps.Maintenance_window#Update
	Update *ActionVpsMaintenanceWindowUpdate
	// Action Vps.Maintenance_window#Update_all
	UpdateAll *ActionVpsMaintenanceWindowUpdateAll
}

func NewResourceVpsMaintenanceWindow(client *Client) *ResourceVpsMaintenanceWindow {
	actionIndex := NewActionVpsMaintenanceWindowIndex(client)
	actionShow := NewActionVpsMaintenanceWindowShow(client)
	actionUpdate := NewActionVpsMaintenanceWindowUpdate(client)
	actionUpdateAll := NewActionVpsMaintenanceWindowUpdateAll(client)

	return &ResourceVpsMaintenanceWindow{
		Client:    client,
		Index:     actionIndex,
		List:      actionIndex,
		Show:      actionShow,
		Find:      actionShow,
		Update:    actionUpdate,
		UpdateAll: actionUpdateAll,
	}
}
