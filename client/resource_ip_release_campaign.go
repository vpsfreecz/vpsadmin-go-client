package client

// Type for resource Ip_release_campaign
type ResourceIpReleaseCampaign struct {
	// Pointer to client
	Client *Client

	// Resource Ip_release_campaign.Address
	Address *ResourceIpReleaseCampaignAddress
	// Resource Ip_release_campaign.Attempt
	Attempt *ResourceIpReleaseCampaignAttempt
	// Resource Ip_release_campaign.Notice
	Notice *ResourceIpReleaseCampaignNotice
	// Action Ip_release_campaign#Candidates
	Candidates *ActionIpReleaseCampaignCandidates
	// Action Ip_release_campaign#Close
	Close *ActionIpReleaseCampaignClose
	// Action Ip_release_campaign#Create
	Create *ActionIpReleaseCampaignCreate
	// Action Ip_release_campaign#Create
	New *ActionIpReleaseCampaignCreate
	// Action Ip_release_campaign#Exempt
	Exempt *ActionIpReleaseCampaignExempt
	// Action Ip_release_campaign#Index
	Index *ActionIpReleaseCampaignIndex
	// Action Ip_release_campaign#Index
	List *ActionIpReleaseCampaignIndex
	// Action Ip_release_campaign#Notify
	Notify *ActionIpReleaseCampaignNotify
	// Action Ip_release_campaign#Release
	Release *ActionIpReleaseCampaignRelease
	// Action Ip_release_campaign#Show
	Show *ActionIpReleaseCampaignShow
	// Action Ip_release_campaign#Show
	Find *ActionIpReleaseCampaignShow
	// Action Ip_release_campaign#Update
	Update *ActionIpReleaseCampaignUpdate
}

func NewResourceIpReleaseCampaign(client *Client) *ResourceIpReleaseCampaign {
	actionCandidates := NewActionIpReleaseCampaignCandidates(client)
	actionClose := NewActionIpReleaseCampaignClose(client)
	actionCreate := NewActionIpReleaseCampaignCreate(client)
	actionExempt := NewActionIpReleaseCampaignExempt(client)
	actionIndex := NewActionIpReleaseCampaignIndex(client)
	actionNotify := NewActionIpReleaseCampaignNotify(client)
	actionRelease := NewActionIpReleaseCampaignRelease(client)
	actionShow := NewActionIpReleaseCampaignShow(client)
	actionUpdate := NewActionIpReleaseCampaignUpdate(client)

	return &ResourceIpReleaseCampaign{
		Client:     client,
		Address:    NewResourceIpReleaseCampaignAddress(client),
		Attempt:    NewResourceIpReleaseCampaignAttempt(client),
		Notice:     NewResourceIpReleaseCampaignNotice(client),
		Candidates: actionCandidates,
		Close:      actionClose,
		Create:     actionCreate,
		New:        actionCreate,
		Exempt:     actionExempt,
		Index:      actionIndex,
		List:       actionIndex,
		Notify:     actionNotify,
		Release:    actionRelease,
		Show:       actionShow,
		Find:       actionShow,
		Update:     actionUpdate,
	}
}
