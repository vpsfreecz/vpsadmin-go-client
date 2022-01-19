package client

// Type for resource Integrity_check
type ResourceIntegrityCheck struct {
	// Pointer to client
	Client *Client

	// Action Integrity_check#Create
	Create *ActionIntegrityCheckCreate
	// Action Integrity_check#Create
	New *ActionIntegrityCheckCreate
	// Action Integrity_check#Index
	Index *ActionIntegrityCheckIndex
	// Action Integrity_check#Index
	List *ActionIntegrityCheckIndex
	// Action Integrity_check#Show
	Show *ActionIntegrityCheckShow
	// Action Integrity_check#Show
	Find *ActionIntegrityCheckShow
}

func NewResourceIntegrityCheck(client *Client) *ResourceIntegrityCheck {
	actionCreate := NewActionIntegrityCheckCreate(client)
	actionIndex := NewActionIntegrityCheckIndex(client)
	actionShow := NewActionIntegrityCheckShow(client)

	return &ResourceIntegrityCheck{
		Client: client,
		Create: actionCreate,
		New: actionCreate,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
	}
}
