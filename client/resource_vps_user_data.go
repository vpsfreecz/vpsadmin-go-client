package client

// Type for resource Vps_user_data
type ResourceVpsUserData struct {
	// Pointer to client
	Client *Client

	// Action Vps_user_data#Create
	Create *ActionVpsUserDataCreate
	// Action Vps_user_data#Create
	New *ActionVpsUserDataCreate
	// Action Vps_user_data#Delete
	Delete *ActionVpsUserDataDelete
	// Action Vps_user_data#Delete
	Destroy *ActionVpsUserDataDelete
	// Action Vps_user_data#Deploy
	Deploy *ActionVpsUserDataDeploy
	// Action Vps_user_data#Index
	Index *ActionVpsUserDataIndex
	// Action Vps_user_data#Index
	List *ActionVpsUserDataIndex
	// Action Vps_user_data#Show
	Show *ActionVpsUserDataShow
	// Action Vps_user_data#Show
	Find *ActionVpsUserDataShow
	// Action Vps_user_data#Update
	Update *ActionVpsUserDataUpdate
}

func NewResourceVpsUserData(client *Client) *ResourceVpsUserData {
	actionCreate := NewActionVpsUserDataCreate(client)
	actionDelete := NewActionVpsUserDataDelete(client)
	actionDeploy := NewActionVpsUserDataDeploy(client)
	actionIndex := NewActionVpsUserDataIndex(client)
	actionShow := NewActionVpsUserDataShow(client)
	actionUpdate := NewActionVpsUserDataUpdate(client)

	return &ResourceVpsUserData{
		Client:  client,
		Create:  actionCreate,
		New:     actionCreate,
		Delete:  actionDelete,
		Destroy: actionDelete,
		Deploy:  actionDeploy,
		Index:   actionIndex,
		List:    actionIndex,
		Show:    actionShow,
		Find:    actionShow,
		Update:  actionUpdate,
	}
}
