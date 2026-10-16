package client

// Type for resource Ip_release_campaign.Attempt
type ResourceIpReleaseCampaignAttempt struct {
	// Pointer to client
	Client *Client

	// Action Ip_release_campaign.Attempt#Index
	Index *ActionIpReleaseCampaignAttemptIndex
	// Action Ip_release_campaign.Attempt#Index
	List *ActionIpReleaseCampaignAttemptIndex
	// Action Ip_release_campaign.Attempt#Show
	Show *ActionIpReleaseCampaignAttemptShow
	// Action Ip_release_campaign.Attempt#Show
	Find *ActionIpReleaseCampaignAttemptShow
}

func NewResourceIpReleaseCampaignAttempt(client *Client) *ResourceIpReleaseCampaignAttempt {
	actionIndex := NewActionIpReleaseCampaignAttemptIndex(client)
	actionShow := NewActionIpReleaseCampaignAttemptShow(client)

	return &ResourceIpReleaseCampaignAttempt{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
