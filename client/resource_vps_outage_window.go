package client

// Type for resource Vps.Outage_window
type ResourceVpsOutageWindow struct {
	// Pointer to client
	Client *Client

	// Action Vps.Outage_window#Index
	Index *ActionVpsOutageWindowIndex
	// Action Vps.Outage_window#Index
	List *ActionVpsOutageWindowIndex
	// Action Vps.Outage_window#Show
	Show *ActionVpsOutageWindowShow
	// Action Vps.Outage_window#Show
	Find *ActionVpsOutageWindowShow
	// Action Vps.Outage_window#Update
	Update *ActionVpsOutageWindowUpdate
	// Action Vps.Outage_window#Update_all
	UpdateAll *ActionVpsOutageWindowUpdateAll
}

func NewResourceVpsOutageWindow(client *Client) *ResourceVpsOutageWindow {
	actionIndex := NewActionVpsOutageWindowIndex(client)
	actionShow := NewActionVpsOutageWindowShow(client)
	actionUpdate := NewActionVpsOutageWindowUpdate(client)
	actionUpdateAll := NewActionVpsOutageWindowUpdateAll(client)

	return &ResourceVpsOutageWindow{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Update: actionUpdate,
		UpdateAll: actionUpdateAll,
	}
}
