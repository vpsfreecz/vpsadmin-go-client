package client

// Type for resource Dns_server_zone_primary_transfer_state
type ResourceDnsServerZonePrimaryTransferState struct {
	// Pointer to client
	Client *Client

	// Action Dns_server_zone_primary_transfer_state#Index
	Index *ActionDnsServerZonePrimaryTransferStateIndex
	// Action Dns_server_zone_primary_transfer_state#Index
	List *ActionDnsServerZonePrimaryTransferStateIndex
}

func NewResourceDnsServerZonePrimaryTransferState(client *Client) *ResourceDnsServerZonePrimaryTransferState {
	actionIndex := NewActionDnsServerZonePrimaryTransferStateIndex(client)

	return &ResourceDnsServerZonePrimaryTransferState{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
