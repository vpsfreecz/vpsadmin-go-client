package client

// Type for resource Transaction_chain
type ResourceTransactionChain struct {
	// Pointer to client
	Client *Client

	// Action Transaction_chain#Index
	Index *ActionTransactionChainIndex
	// Action Transaction_chain#Index
	List *ActionTransactionChainIndex
	// Action Transaction_chain#Show
	Show *ActionTransactionChainShow
	// Action Transaction_chain#Show
	Find *ActionTransactionChainShow
}

func NewResourceTransactionChain(client *Client) *ResourceTransactionChain {
	actionIndex := NewActionTransactionChainIndex(client)
	actionShow := NewActionTransactionChainShow(client)

	return &ResourceTransactionChain{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
