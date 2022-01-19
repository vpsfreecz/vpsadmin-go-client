package client

// Type for resource Vps_outage
type ResourceVpsOutage struct {
	// Pointer to client
	Client *Client

	// Action Vps_outage#Index
	Index *ActionVpsOutageIndex
	// Action Vps_outage#Index
	List *ActionVpsOutageIndex
	// Action Vps_outage#Show
	Show *ActionVpsOutageShow
	// Action Vps_outage#Show
	Find *ActionVpsOutageShow
}

func NewResourceVpsOutage(client *Client) *ResourceVpsOutage {
	actionIndex := NewActionVpsOutageIndex(client)
	actionShow := NewActionVpsOutageShow(client)

	return &ResourceVpsOutage{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
