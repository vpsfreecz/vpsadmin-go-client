package client

// Type for resource Dns_server_zone
type ResourceDnsServerZone struct {
	// Pointer to client
	Client *Client

	// Action Dns_server_zone#Create
	Create *ActionDnsServerZoneCreate
	// Action Dns_server_zone#Create
	New *ActionDnsServerZoneCreate
	// Action Dns_server_zone#Delete
	Delete *ActionDnsServerZoneDelete
	// Action Dns_server_zone#Delete
	Destroy *ActionDnsServerZoneDelete
	// Action Dns_server_zone#Index
	Index *ActionDnsServerZoneIndex
	// Action Dns_server_zone#Index
	List *ActionDnsServerZoneIndex
	// Action Dns_server_zone#Show
	Show *ActionDnsServerZoneShow
	// Action Dns_server_zone#Show
	Find *ActionDnsServerZoneShow
}

func NewResourceDnsServerZone(client *Client) *ResourceDnsServerZone {
	actionCreate := NewActionDnsServerZoneCreate(client)
	actionDelete := NewActionDnsServerZoneDelete(client)
	actionIndex := NewActionDnsServerZoneIndex(client)
	actionShow := NewActionDnsServerZoneShow(client)

	return &ResourceDnsServerZone{
		Client:  client,
		Create:  actionCreate,
		New:     actionCreate,
		Delete:  actionDelete,
		Destroy: actionDelete,
		Index:   actionIndex,
		List:    actionIndex,
		Show:    actionShow,
		Find:    actionShow,
	}
}
