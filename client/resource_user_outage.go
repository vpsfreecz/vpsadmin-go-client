package client

// Type for resource User_outage
type ResourceUserOutage struct {
	// Pointer to client
	Client *Client

	// Action User_outage#Index
	Index *ActionUserOutageIndex
	// Action User_outage#Index
	List *ActionUserOutageIndex
	// Action User_outage#Show
	Show *ActionUserOutageShow
	// Action User_outage#Show
	Find *ActionUserOutageShow
}

func NewResourceUserOutage(client *Client) *ResourceUserOutage {
	actionIndex := NewActionUserOutageIndex(client)
	actionShow := NewActionUserOutageShow(client)

	return &ResourceUserOutage{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
	}
}
