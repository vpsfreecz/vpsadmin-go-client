package client

// Type for resource Language
type ResourceLanguage struct {
	// Pointer to client
	Client *Client

	// Action Language#Index
	Index *ActionLanguageIndex
	// Action Language#Index
	List *ActionLanguageIndex
	// Action Language#Show
	Show *ActionLanguageShow
	// Action Language#Show
	Find *ActionLanguageShow
}

func NewResourceLanguage(client *Client) *ResourceLanguage {
	actionIndex := NewActionLanguageIndex(client)
	actionShow := NewActionLanguageShow(client)

	return &ResourceLanguage{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
