package client

// Type for resource Dns_server
type ResourceDnsServer struct {
	// Pointer to client
	Client *Client

	// Action Dns_server#Create
	Create *ActionDnsServerCreate
	// Action Dns_server#Create
	New *ActionDnsServerCreate
	// Action Dns_server#Delete
	Delete *ActionDnsServerDelete
	// Action Dns_server#Delete
	Destroy *ActionDnsServerDelete
	// Action Dns_server#Index
	Index *ActionDnsServerIndex
	// Action Dns_server#Index
	List *ActionDnsServerIndex
	// Action Dns_server#Show
	Show *ActionDnsServerShow
	// Action Dns_server#Show
	Find *ActionDnsServerShow
	// Action Dns_server#Update
	Update *ActionDnsServerUpdate
}

func NewResourceDnsServer(client *Client) *ResourceDnsServer {
	actionCreate := NewActionDnsServerCreate(client)
	actionDelete := NewActionDnsServerDelete(client)
	actionIndex := NewActionDnsServerIndex(client)
	actionShow := NewActionDnsServerShow(client)
	actionUpdate := NewActionDnsServerUpdate(client)

	return &ResourceDnsServer{
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
