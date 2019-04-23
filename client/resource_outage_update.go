package client

// Type for resource Outage_update
type ResourceOutageUpdate struct {
	// Pointer to client
	Client *Client

	// Action Outage_update#Index
	Index *ActionOutageUpdateIndex
	// Action Outage_update#Index
	List *ActionOutageUpdateIndex
	// Action Outage_update#Show
	Show *ActionOutageUpdateShow
	// Action Outage_update#Show
	Find *ActionOutageUpdateShow
}

func NewResourceOutageUpdate(client *Client) *ResourceOutageUpdate {
	actionIndex := NewActionOutageUpdateIndex(client)
	actionShow := NewActionOutageUpdateShow(client)

	return &ResourceOutageUpdate{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
	}
}
