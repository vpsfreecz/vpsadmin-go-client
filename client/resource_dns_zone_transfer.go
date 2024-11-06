package client

// Type for resource Dns_zone_transfer
type ResourceDnsZoneTransfer struct {
	// Pointer to client
	Client *Client

	// Action Dns_zone_transfer#Create
	Create *ActionDnsZoneTransferCreate
	// Action Dns_zone_transfer#Create
	New *ActionDnsZoneTransferCreate
	// Action Dns_zone_transfer#Delete
	Delete *ActionDnsZoneTransferDelete
	// Action Dns_zone_transfer#Delete
	Destroy *ActionDnsZoneTransferDelete
	// Action Dns_zone_transfer#Index
	Index *ActionDnsZoneTransferIndex
	// Action Dns_zone_transfer#Index
	List *ActionDnsZoneTransferIndex
	// Action Dns_zone_transfer#Show
	Show *ActionDnsZoneTransferShow
	// Action Dns_zone_transfer#Show
	Find *ActionDnsZoneTransferShow
}

func NewResourceDnsZoneTransfer(client *Client) *ResourceDnsZoneTransfer {
	actionCreate := NewActionDnsZoneTransferCreate(client)
	actionDelete := NewActionDnsZoneTransferDelete(client)
	actionIndex := NewActionDnsZoneTransferIndex(client)
	actionShow := NewActionDnsZoneTransferShow(client)

	return &ResourceDnsZoneTransfer{
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
