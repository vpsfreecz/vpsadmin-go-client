package client

// Type for resource Dns_zone
type ResourceDnsZone struct {
	// Pointer to client
	Client *Client

	// Action Dns_zone#Create
	Create *ActionDnsZoneCreate
	// Action Dns_zone#Create
	New *ActionDnsZoneCreate
	// Action Dns_zone#Delete
	Delete *ActionDnsZoneDelete
	// Action Dns_zone#Delete
	Destroy *ActionDnsZoneDelete
	// Action Dns_zone#Index
	Index *ActionDnsZoneIndex
	// Action Dns_zone#Index
	List *ActionDnsZoneIndex
	// Action Dns_zone#Show
	Show *ActionDnsZoneShow
	// Action Dns_zone#Show
	Find *ActionDnsZoneShow
	// Action Dns_zone#Update
	Update *ActionDnsZoneUpdate
}

func NewResourceDnsZone(client *Client) *ResourceDnsZone {
	actionCreate := NewActionDnsZoneCreate(client)
	actionDelete := NewActionDnsZoneDelete(client)
	actionIndex := NewActionDnsZoneIndex(client)
	actionShow := NewActionDnsZoneShow(client)
	actionUpdate := NewActionDnsZoneUpdate(client)

	return &ResourceDnsZone{
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
