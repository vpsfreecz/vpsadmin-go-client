package client

// Type for resource News_log
type ResourceNewsLog struct {
	// Pointer to client
	Client *Client

	// Action News_log#Create
	Create *ActionNewsLogCreate
	// Action News_log#Create
	New *ActionNewsLogCreate
	// Action News_log#Delete
	Delete *ActionNewsLogDelete
	// Action News_log#Delete
	Destroy *ActionNewsLogDelete
	// Action News_log#Index
	Index *ActionNewsLogIndex
	// Action News_log#Index
	List *ActionNewsLogIndex
	// Action News_log#Show
	Show *ActionNewsLogShow
	// Action News_log#Show
	Find *ActionNewsLogShow
	// Action News_log#Update
	Update *ActionNewsLogUpdate
}

func NewResourceNewsLog(client *Client) *ResourceNewsLog {
	actionCreate := NewActionNewsLogCreate(client)
	actionDelete := NewActionNewsLogDelete(client)
	actionIndex := NewActionNewsLogIndex(client)
	actionShow := NewActionNewsLogShow(client)
	actionUpdate := NewActionNewsLogUpdate(client)

	return &ResourceNewsLog{
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
