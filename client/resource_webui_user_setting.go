package client

// Type for resource Webui_user_setting
type ResourceWebuiUserSetting struct {
	// Pointer to client
	Client *Client

	// Action Webui_user_setting#Delete
	Delete *ActionWebuiUserSettingDelete
	// Action Webui_user_setting#Index
	Index *ActionWebuiUserSettingIndex
	// Action Webui_user_setting#Index
	List *ActionWebuiUserSettingIndex
	// Action Webui_user_setting#Set
	Set *ActionWebuiUserSettingSet
	// Action Webui_user_setting#Show
	Show *ActionWebuiUserSettingShow
}

func NewResourceWebuiUserSetting(client *Client) *ResourceWebuiUserSetting {
	actionDelete := NewActionWebuiUserSettingDelete(client)
	actionIndex := NewActionWebuiUserSettingIndex(client)
	actionSet := NewActionWebuiUserSettingSet(client)
	actionShow := NewActionWebuiUserSettingShow(client)

	return &ResourceWebuiUserSetting{
		Client: client,
		Delete: actionDelete,
		Index:  actionIndex,
		List:   actionIndex,
		Set:    actionSet,
		Show:   actionShow,
	}
}
