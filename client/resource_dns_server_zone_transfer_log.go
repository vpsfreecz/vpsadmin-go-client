package client

// Type for resource Dns_server_zone_transfer_log
type ResourceDnsServerZoneTransferLog struct {
	// Pointer to client
	Client *Client

	// Action Dns_server_zone_transfer_log#Index
	Index *ActionDnsServerZoneTransferLogIndex
	// Action Dns_server_zone_transfer_log#Index
	List *ActionDnsServerZoneTransferLogIndex
	// Action Dns_server_zone_transfer_log#Show
	Show *ActionDnsServerZoneTransferLogShow
	// Action Dns_server_zone_transfer_log#Show
	Find *ActionDnsServerZoneTransferLogShow
}

func NewResourceDnsServerZoneTransferLog(client *Client) *ResourceDnsServerZoneTransferLog {
	actionIndex := NewActionDnsServerZoneTransferLogIndex(client)
	actionShow := NewActionDnsServerZoneTransferLogShow(client)

	return &ResourceDnsServerZoneTransferLog{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
