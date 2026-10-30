package client

// Type for resource Node_system_state
type ResourceNodeSystemState struct {
	// Pointer to client
	Client *Client

	// Action Node_system_state#Index
	Index *ActionNodeSystemStateIndex
	// Action Node_system_state#Index
	List *ActionNodeSystemStateIndex
	// Action Node_system_state#Show
	Show *ActionNodeSystemStateShow
	// Action Node_system_state#Show
	Find *ActionNodeSystemStateShow
}

func NewResourceNodeSystemState(client *Client) *ResourceNodeSystemState {
	actionIndex := NewActionNodeSystemStateIndex(client)
	actionShow := NewActionNodeSystemStateShow(client)

	return &ResourceNodeSystemState{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
