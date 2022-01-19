package client

// Type for resource User.Public_key
type ResourceUserPublicKey struct {
	// Pointer to client
	Client *Client

	// Action User.Public_key#Create
	Create *ActionUserPublicKeyCreate
	// Action User.Public_key#Create
	New *ActionUserPublicKeyCreate
	// Action User.Public_key#Delete
	Delete *ActionUserPublicKeyDelete
	// Action User.Public_key#Delete
	Destroy *ActionUserPublicKeyDelete
	// Action User.Public_key#Index
	Index *ActionUserPublicKeyIndex
	// Action User.Public_key#Index
	List *ActionUserPublicKeyIndex
	// Action User.Public_key#Show
	Show *ActionUserPublicKeyShow
	// Action User.Public_key#Show
	Find *ActionUserPublicKeyShow
	// Action User.Public_key#Update
	Update *ActionUserPublicKeyUpdate
}

func NewResourceUserPublicKey(client *Client) *ResourceUserPublicKey {
	actionCreate := NewActionUserPublicKeyCreate(client)
	actionDelete := NewActionUserPublicKeyDelete(client)
	actionIndex := NewActionUserPublicKeyIndex(client)
	actionShow := NewActionUserPublicKeyShow(client)
	actionUpdate := NewActionUserPublicKeyUpdate(client)

	return &ResourceUserPublicKey{
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
