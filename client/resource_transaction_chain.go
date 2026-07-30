package client

// Type for resource Transaction_chain
type ResourceTransactionChain struct {
	// Pointer to client
	Client *Client

	// Action Transaction_chain#Index
	Index *ActionTransactionChainIndex
	// Action Transaction_chain#Index
	List *ActionTransactionChainIndex
	// Action Transaction_chain#Notify_when_done
	NotifyWhenDone *ActionTransactionChainNotifyWhenDone
	// Action Transaction_chain#Show
	Show *ActionTransactionChainShow
	// Action Transaction_chain#Show
	Find *ActionTransactionChainShow
}

func NewResourceTransactionChain(client *Client) *ResourceTransactionChain {
	actionIndex := NewActionTransactionChainIndex(client)
	actionNotifyWhenDone := NewActionTransactionChainNotifyWhenDone(client)
	actionShow := NewActionTransactionChainShow(client)

	return &ResourceTransactionChain{
		Client:         client,
		Index:          actionIndex,
		List:           actionIndex,
		NotifyWhenDone: actionNotifyWhenDone,
		Show:           actionShow,
		Find:           actionShow,
	}
}
