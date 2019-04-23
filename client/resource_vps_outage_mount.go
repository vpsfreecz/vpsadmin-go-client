package client

// Type for resource Vps_outage_mount
type ResourceVpsOutageMount struct {
	// Pointer to client
	Client *Client

	// Action Vps_outage_mount#Index
	Index *ActionVpsOutageMountIndex
	// Action Vps_outage_mount#Index
	List *ActionVpsOutageMountIndex
	// Action Vps_outage_mount#Show
	Show *ActionVpsOutageMountShow
	// Action Vps_outage_mount#Show
	Find *ActionVpsOutageMountShow
}

func NewResourceVpsOutageMount(client *Client) *ResourceVpsOutageMount {
	actionIndex := NewActionVpsOutageMountIndex(client)
	actionShow := NewActionVpsOutageMountShow(client)

	return &ResourceVpsOutageMount{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
	}
}
