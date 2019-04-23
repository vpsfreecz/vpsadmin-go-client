package client

// Type for resource System_config
type ResourceSystemConfig struct {
	// Pointer to client
	Client *Client

	// Action System_config#Index
	Index *ActionSystemConfigIndex
	// Action System_config#Index
	List *ActionSystemConfigIndex
	// Action System_config#Show
	Show *ActionSystemConfigShow
	// Action System_config#Update
	Update *ActionSystemConfigUpdate
}

func NewResourceSystemConfig(client *Client) *ResourceSystemConfig {
	actionIndex := NewActionSystemConfigIndex(client)
	actionShow := NewActionSystemConfigShow(client)
	actionUpdate := NewActionSystemConfigUpdate(client)

	return &ResourceSystemConfig{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Update: actionUpdate,
	}
}
