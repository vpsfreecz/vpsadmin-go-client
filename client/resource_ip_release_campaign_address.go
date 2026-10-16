package client

// Type for resource Ip_release_campaign.Address
type ResourceIpReleaseCampaignAddress struct {
	// Pointer to client
	Client *Client

	// Action Ip_release_campaign.Address#Index
	Index *ActionIpReleaseCampaignAddressIndex
	// Action Ip_release_campaign.Address#Index
	List *ActionIpReleaseCampaignAddressIndex
}

func NewResourceIpReleaseCampaignAddress(client *Client) *ResourceIpReleaseCampaignAddress {
	actionIndex := NewActionIpReleaseCampaignAddressIndex(client)

	return &ResourceIpReleaseCampaignAddress{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
