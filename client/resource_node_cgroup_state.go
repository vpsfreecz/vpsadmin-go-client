package client

// Type for resource Node_cgroup_state
type ResourceNodeCgroupState struct {
	// Pointer to client
	Client *Client

	// Action Node_cgroup_state#Index
	Index *ActionNodeCgroupStateIndex
	// Action Node_cgroup_state#Index
	List *ActionNodeCgroupStateIndex
	// Action Node_cgroup_state#Show
	Show *ActionNodeCgroupStateShow
	// Action Node_cgroup_state#Show
	Find *ActionNodeCgroupStateShow
}

func NewResourceNodeCgroupState(client *Client) *ResourceNodeCgroupState {
	actionIndex := NewActionNodeCgroupStateIndex(client)
	actionShow := NewActionNodeCgroupStateShow(client)

	return &ResourceNodeCgroupState{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
