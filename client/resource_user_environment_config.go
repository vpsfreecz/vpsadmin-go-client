package client

// Type for resource User.Environment_config
type ResourceUserEnvironmentConfig struct {
	// Pointer to client
	Client *Client

	// Action User.Environment_config#Index
	Index *ActionUserEnvironmentConfigIndex
	// Action User.Environment_config#Index
	List *ActionUserEnvironmentConfigIndex
	// Action User.Environment_config#Show
	Show *ActionUserEnvironmentConfigShow
	// Action User.Environment_config#Show
	Find *ActionUserEnvironmentConfigShow
	// Action User.Environment_config#Update
	Update *ActionUserEnvironmentConfigUpdate
}

func NewResourceUserEnvironmentConfig(client *Client) *ResourceUserEnvironmentConfig {
	actionIndex := NewActionUserEnvironmentConfigIndex(client)
	actionShow := NewActionUserEnvironmentConfigShow(client)
	actionUpdate := NewActionUserEnvironmentConfigUpdate(client)

	return &ResourceUserEnvironmentConfig{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Update: actionUpdate,
	}
}
