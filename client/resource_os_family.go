package client

// Type for resource Os_family
type ResourceOsFamily struct {
	// Pointer to client
	Client *Client

	// Action Os_family#Create
	Create *ActionOsFamilyCreate
	// Action Os_family#Create
	New *ActionOsFamilyCreate
	// Action Os_family#Delete
	Delete *ActionOsFamilyDelete
	// Action Os_family#Delete
	Destroy *ActionOsFamilyDelete
	// Action Os_family#Index
	Index *ActionOsFamilyIndex
	// Action Os_family#Index
	List *ActionOsFamilyIndex
	// Action Os_family#Show
	Show *ActionOsFamilyShow
	// Action Os_family#Show
	Find *ActionOsFamilyShow
	// Action Os_family#Update
	Update *ActionOsFamilyUpdate
}

func NewResourceOsFamily(client *Client) *ResourceOsFamily {
	actionCreate := NewActionOsFamilyCreate(client)
	actionDelete := NewActionOsFamilyDelete(client)
	actionIndex := NewActionOsFamilyIndex(client)
	actionShow := NewActionOsFamilyShow(client)
	actionUpdate := NewActionOsFamilyUpdate(client)

	return &ResourceOsFamily{
		Client:  client,
		Create:  actionCreate,
		New:     actionCreate,
		Delete:  actionDelete,
		Destroy: actionDelete,
		Index:   actionIndex,
		List:    actionIndex,
		Show:    actionShow,
		Find:    actionShow,
		Update:  actionUpdate,
	}
}
