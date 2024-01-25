package client

// Type for resource Component
type ResourceComponent struct {
	// Pointer to client
	Client *Client

	// Action Component#Index
	Index *ActionComponentIndex
	// Action Component#Index
	List *ActionComponentIndex
	// Action Component#Show
	Show *ActionComponentShow
	// Action Component#Show
	Find *ActionComponentShow
}

func NewResourceComponent(client *Client) *ResourceComponent {
	actionIndex := NewActionComponentIndex(client)
	actionShow := NewActionComponentShow(client)

	return &ResourceComponent{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
