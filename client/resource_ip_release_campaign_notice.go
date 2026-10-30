package client

// Type for resource Ip_release_campaign.Notice
type ResourceIpReleaseCampaignNotice struct {
	// Pointer to client
	Client *Client

	// Action Ip_release_campaign.Notice#Index
	Index *ActionIpReleaseCampaignNoticeIndex
	// Action Ip_release_campaign.Notice#Index
	List *ActionIpReleaseCampaignNoticeIndex
}

func NewResourceIpReleaseCampaignNotice(client *Client) *ResourceIpReleaseCampaignNotice {
	actionIndex := NewActionIpReleaseCampaignNoticeIndex(client)

	return &ResourceIpReleaseCampaignNotice{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
