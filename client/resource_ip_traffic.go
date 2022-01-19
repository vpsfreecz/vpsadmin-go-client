package client

// Type for resource Ip_traffic
type ResourceIpTraffic struct {
	// Pointer to client
	Client *Client

	// Action Ip_traffic#Index
	Index *ActionIpTrafficIndex
	// Action Ip_traffic#Index
	List *ActionIpTrafficIndex
	// Action Ip_traffic#Show
	Show *ActionIpTrafficShow
	// Action Ip_traffic#Show
	Find *ActionIpTrafficShow
	// Action Ip_traffic#User_top
	UserTop *ActionIpTrafficUserTop
}

func NewResourceIpTraffic(client *Client) *ResourceIpTraffic {
	actionIndex := NewActionIpTrafficIndex(client)
	actionShow := NewActionIpTrafficShow(client)
	actionUserTop := NewActionIpTrafficUserTop(client)

	return &ResourceIpTraffic{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		UserTop: actionUserTop,
	}
}
