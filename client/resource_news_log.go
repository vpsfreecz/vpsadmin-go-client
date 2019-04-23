package client

// Type for resource News_log
type ResourceNewsLog struct {
	// Pointer to client
	Client *Client

	// Action News_log#Index
	Index *ActionNewsLogIndex
	// Action News_log#Index
	List *ActionNewsLogIndex
	// Action News_log#Show
	Show *ActionNewsLogShow
	// Action News_log#Show
	Find *ActionNewsLogShow
	// Action News_log#Create
	Create *ActionNewsLogCreate
	// Action News_log#Create
	New *ActionNewsLogCreate
	// Action News_log#Update
	Update *ActionNewsLogUpdate
	// Action News_log#Delete
	Delete *ActionNewsLogDelete
	// Action News_log#Delete
	Destroy *ActionNewsLogDelete
}

func NewResourceNewsLog(client *Client) *ResourceNewsLog {
	actionIndex := NewActionNewsLogIndex(client)
	actionShow := NewActionNewsLogShow(client)
	actionCreate := NewActionNewsLogCreate(client)
	actionUpdate := NewActionNewsLogUpdate(client)
	actionDelete := NewActionNewsLogDelete(client)

	return &ResourceNewsLog{
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
