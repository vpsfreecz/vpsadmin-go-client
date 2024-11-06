package client

// Type for resource Dns_tsig_key
type ResourceDnsTsigKey struct {
	// Pointer to client
	Client *Client

	// Action Dns_tsig_key#Create
	Create *ActionDnsTsigKeyCreate
	// Action Dns_tsig_key#Create
	New *ActionDnsTsigKeyCreate
	// Action Dns_tsig_key#Delete
	Delete *ActionDnsTsigKeyDelete
	// Action Dns_tsig_key#Delete
	Destroy *ActionDnsTsigKeyDelete
	// Action Dns_tsig_key#Index
	Index *ActionDnsTsigKeyIndex
	// Action Dns_tsig_key#Index
	List *ActionDnsTsigKeyIndex
	// Action Dns_tsig_key#Show
	Show *ActionDnsTsigKeyShow
	// Action Dns_tsig_key#Show
	Find *ActionDnsTsigKeyShow
}

func NewResourceDnsTsigKey(client *Client) *ResourceDnsTsigKey {
	actionCreate := NewActionDnsTsigKeyCreate(client)
	actionDelete := NewActionDnsTsigKeyDelete(client)
	actionIndex := NewActionDnsTsigKeyIndex(client)
	actionShow := NewActionDnsTsigKeyShow(client)

	return &ResourceDnsTsigKey{
		Client:  client,
		Create:  actionCreate,
		New:     actionCreate,
		Delete:  actionDelete,
		Destroy: actionDelete,
		Index:   actionIndex,
		List:    actionIndex,
		Show:    actionShow,
		Find:    actionShow,
	}
}
