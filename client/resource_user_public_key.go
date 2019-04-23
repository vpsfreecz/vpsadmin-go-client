package client

// Type for resource User.Public_key
type ResourceUserPublicKey struct {
	// Pointer to client
	Client *Client

	// Action User.Public_key#Index
	Index *ActionUserPublicKeyIndex
	// Action User.Public_key#Index
	List *ActionUserPublicKeyIndex
	// Action User.Public_key#Show
	Show *ActionUserPublicKeyShow
	// Action User.Public_key#Show
	Find *ActionUserPublicKeyShow
	// Action User.Public_key#Create
	Create *ActionUserPublicKeyCreate
	// Action User.Public_key#Create
	New *ActionUserPublicKeyCreate
	// Action User.Public_key#Update
	Update *ActionUserPublicKeyUpdate
	// Action User.Public_key#Delete
	Delete *ActionUserPublicKeyDelete
	// Action User.Public_key#Delete
	Destroy *ActionUserPublicKeyDelete
}

func NewResourceUserPublicKey(client *Client) *ResourceUserPublicKey {
	actionIndex := NewActionUserPublicKeyIndex(client)
	actionShow := NewActionUserPublicKeyShow(client)
	actionCreate := NewActionUserPublicKeyCreate(client)
	actionUpdate := NewActionUserPublicKeyUpdate(client)
	actionDelete := NewActionUserPublicKeyDelete(client)

	return &ResourceUserPublicKey{
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
