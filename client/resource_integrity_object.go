package client

// Type for resource Integrity_object
type ResourceIntegrityObject struct {
	// Pointer to client
	Client *Client

	// Action Integrity_object#Index
	Index *ActionIntegrityObjectIndex
	// Action Integrity_object#Index
	List *ActionIntegrityObjectIndex
	// Action Integrity_object#Show
	Show *ActionIntegrityObjectShow
	// Action Integrity_object#Show
	Find *ActionIntegrityObjectShow
}

func NewResourceIntegrityObject(client *Client) *ResourceIntegrityObject {
	actionIndex := NewActionIntegrityObjectIndex(client)
	actionShow := NewActionIntegrityObjectShow(client)

	return &ResourceIntegrityObject{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
	}
}
