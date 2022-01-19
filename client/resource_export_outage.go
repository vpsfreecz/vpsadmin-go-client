package client

// Type for resource Export_outage
type ResourceExportOutage struct {
	// Pointer to client
	Client *Client

	// Action Export_outage#Index
	Index *ActionExportOutageIndex
	// Action Export_outage#Index
	List *ActionExportOutageIndex
	// Action Export_outage#Show
	Show *ActionExportOutageShow
	// Action Export_outage#Show
	Find *ActionExportOutageShow
}

func NewResourceExportOutage(client *Client) *ResourceExportOutage {
	actionIndex := NewActionExportOutageIndex(client)
	actionShow := NewActionExportOutageShow(client)

	return &ResourceExportOutage{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
