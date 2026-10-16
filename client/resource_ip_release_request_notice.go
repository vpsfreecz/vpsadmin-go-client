package client

// Type for resource Ip_release_request.Notice
type ResourceIpReleaseRequestNotice struct {
	// Pointer to client
	Client *Client

	// Action Ip_release_request.Notice#Index
	Index *ActionIpReleaseRequestNoticeIndex
	// Action Ip_release_request.Notice#Index
	List *ActionIpReleaseRequestNoticeIndex
}

func NewResourceIpReleaseRequestNotice(client *Client) *ResourceIpReleaseRequestNotice {
	actionIndex := NewActionIpReleaseRequestNoticeIndex(client)

	return &ResourceIpReleaseRequestNotice{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
