package client

// Type for resource Ip_traffic
type ResourceIpTraffic struct {
	// Pointer to client
	Client *Client

	// Action Ip_traffic#Index
	Index *ActionIpTrafficIndex
	// Action Ip_traffic#Index
	List *ActionIpTrafficIndex
	// Action Ip_traffic#User_top
	UserTop *ActionIpTrafficUserTop
	// Action Ip_traffic#Show
	Show *ActionIpTrafficShow
	// Action Ip_traffic#Show
	Find *ActionIpTrafficShow
}

func NewResourceIpTraffic(client *Client) *ResourceIpTraffic {
	actionIndex := NewActionIpTrafficIndex(client)
	actionUserTop := NewActionIpTrafficUserTop(client)
	actionShow := NewActionIpTrafficShow(client)

	return &ResourceIpTraffic{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		UserTop: actionUserTop,
		Show: actionShow,
		Find: actionShow,
	}
}
