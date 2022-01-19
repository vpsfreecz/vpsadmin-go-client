package client

// Type for resource Action_state
type ResourceActionState struct {
	// Pointer to client
	Client *Client

	// Action Action_state#Cancel
	Cancel *ActionActionStateCancel
	// Action Action_state#Index
	Index *ActionActionStateIndex
	// Action Action_state#Index
	List *ActionActionStateIndex
	// Action Action_state#Poll
	Poll *ActionActionStatePoll
	// Action Action_state#Show
	Show *ActionActionStateShow
	// Action Action_state#Show
	Find *ActionActionStateShow
}

func NewResourceActionState(client *Client) *ResourceActionState {
	actionCancel := NewActionActionStateCancel(client)
	actionIndex := NewActionActionStateIndex(client)
	actionPoll := NewActionActionStatePoll(client)
	actionShow := NewActionActionStateShow(client)

	return &ResourceActionState{
		Client: client,
		Cancel: actionCancel,
		Index: actionIndex,
		List: actionIndex,
		Poll: actionPoll,
		Show: actionShow,
		Find: actionShow,
	}
}
