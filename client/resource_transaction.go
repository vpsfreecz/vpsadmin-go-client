package client

// Type for resource Transaction
type ResourceTransaction struct {
	// Pointer to client
	Client *Client

	// Action Transaction#Index
	Index *ActionTransactionIndex
	// Action Transaction#Index
	List *ActionTransactionIndex
	// Action Transaction#Show
	Show *ActionTransactionShow
	// Action Transaction#Show
	Find *ActionTransactionShow
}

func NewResourceTransaction(client *Client) *ResourceTransaction {
	actionIndex := NewActionTransactionIndex(client)
	actionShow := NewActionTransactionShow(client)

	return &ResourceTransaction{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
	}
}
