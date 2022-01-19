package client

// Type for resource Help_box
type ResourceHelpBox struct {
	// Pointer to client
	Client *Client

	// Action Help_box#Create
	Create *ActionHelpBoxCreate
	// Action Help_box#Create
	New *ActionHelpBoxCreate
	// Action Help_box#Delete
	Delete *ActionHelpBoxDelete
	// Action Help_box#Delete
	Destroy *ActionHelpBoxDelete
	// Action Help_box#Index
	Index *ActionHelpBoxIndex
	// Action Help_box#Index
	List *ActionHelpBoxIndex
	// Action Help_box#Show
	Show *ActionHelpBoxShow
	// Action Help_box#Show
	Find *ActionHelpBoxShow
	// Action Help_box#Update
	Update *ActionHelpBoxUpdate
}

func NewResourceHelpBox(client *Client) *ResourceHelpBox {
	actionCreate := NewActionHelpBoxCreate(client)
	actionDelete := NewActionHelpBoxDelete(client)
	actionIndex := NewActionHelpBoxIndex(client)
	actionShow := NewActionHelpBoxShow(client)
	actionUpdate := NewActionHelpBoxUpdate(client)

	return &ResourceHelpBox{
		Client: client,
		Create: actionCreate,
		New: actionCreate,
		Delete: actionDelete,
		Destroy: actionDelete,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Update: actionUpdate,
	}
}
