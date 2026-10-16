package client

// Type for resource Ip_release_request
type ResourceIpReleaseRequest struct {
	// Pointer to client
	Client *Client

	// Resource Ip_release_request.Address
	Address *ResourceIpReleaseRequestAddress
	// Resource Ip_release_request.Notice
	Notice *ResourceIpReleaseRequestNotice
	// Action Ip_release_request#Index
	Index *ActionIpReleaseRequestIndex
	// Action Ip_release_request#Index
	List *ActionIpReleaseRequestIndex
	// Action Ip_release_request#Keep
	Keep *ActionIpReleaseRequestKeep
	// Action Ip_release_request#Show
	Show *ActionIpReleaseRequestShow
	// Action Ip_release_request#Show
	Find *ActionIpReleaseRequestShow
}

func NewResourceIpReleaseRequest(client *Client) *ResourceIpReleaseRequest {
	actionIndex := NewActionIpReleaseRequestIndex(client)
	actionKeep := NewActionIpReleaseRequestKeep(client)
	actionShow := NewActionIpReleaseRequestShow(client)

	return &ResourceIpReleaseRequest{
		Client:  client,
		Address: NewResourceIpReleaseRequestAddress(client),
		Notice:  NewResourceIpReleaseRequestNotice(client),
		Index:   actionIndex,
		List:    actionIndex,
		Keep:    actionKeep,
		Show:    actionShow,
		Find:    actionShow,
	}
}
