package client

// Type for resource Integrity_fact
type ResourceIntegrityFact struct {
	// Pointer to client
	Client *Client

	// Action Integrity_fact#Index
	Index *ActionIntegrityFactIndex
	// Action Integrity_fact#Index
	List *ActionIntegrityFactIndex
	// Action Integrity_fact#Show
	Show *ActionIntegrityFactShow
	// Action Integrity_fact#Show
	Find *ActionIntegrityFactShow
}

func NewResourceIntegrityFact(client *Client) *ResourceIntegrityFact {
	actionIndex := NewActionIntegrityFactIndex(client)
	actionShow := NewActionIntegrityFactShow(client)

	return &ResourceIntegrityFact{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
	}
}
