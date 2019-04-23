package client

// Type for resource Help_box
type ResourceHelpBox struct {
	// Pointer to client
	Client *Client

	// Action Help_box#Index
	Index *ActionHelpBoxIndex
	// Action Help_box#Index
	List *ActionHelpBoxIndex
	// Action Help_box#Show
	Show *ActionHelpBoxShow
	// Action Help_box#Show
	Find *ActionHelpBoxShow
	// Action Help_box#Create
	Create *ActionHelpBoxCreate
	// Action Help_box#Create
	New *ActionHelpBoxCreate
	// Action Help_box#Update
	Update *ActionHelpBoxUpdate
	// Action Help_box#Delete
	Delete *ActionHelpBoxDelete
	// Action Help_box#Delete
	Destroy *ActionHelpBoxDelete
}

func NewResourceHelpBox(client *Client) *ResourceHelpBox {
	actionIndex := NewActionHelpBoxIndex(client)
	actionShow := NewActionHelpBoxShow(client)
	actionCreate := NewActionHelpBoxCreate(client)
	actionUpdate := NewActionHelpBoxUpdate(client)
	actionDelete := NewActionHelpBoxDelete(client)

	return &ResourceHelpBox{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Create: actionCreate,
		New: actionCreate,
		Update: actionUpdate,
		Delete: actionDelete,
		Destroy: actionDelete,
	}
}
