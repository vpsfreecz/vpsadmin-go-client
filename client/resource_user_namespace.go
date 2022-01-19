package client

// Type for resource User_namespace
type ResourceUserNamespace struct {
	// Pointer to client
	Client *Client

	// Action User_namespace#Index
	Index *ActionUserNamespaceIndex
	// Action User_namespace#Index
	List *ActionUserNamespaceIndex
	// Action User_namespace#Show
	Show *ActionUserNamespaceShow
	// Action User_namespace#Show
	Find *ActionUserNamespaceShow
}

func NewResourceUserNamespace(client *Client) *ResourceUserNamespace {
	actionIndex := NewActionUserNamespaceIndex(client)
	actionShow := NewActionUserNamespaceShow(client)

	return &ResourceUserNamespace{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
