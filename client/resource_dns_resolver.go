package client

// Type for resource Dns_resolver
type ResourceDnsResolver struct {
	// Pointer to client
	Client *Client

	// Action Dns_resolver#Create
	Create *ActionDnsResolverCreate
	// Action Dns_resolver#Create
	New *ActionDnsResolverCreate
	// Action Dns_resolver#Delete
	Delete *ActionDnsResolverDelete
	// Action Dns_resolver#Delete
	Destroy *ActionDnsResolverDelete
	// Action Dns_resolver#Index
	Index *ActionDnsResolverIndex
	// Action Dns_resolver#Index
	List *ActionDnsResolverIndex
	// Action Dns_resolver#Show
	Show *ActionDnsResolverShow
	// Action Dns_resolver#Show
	Find *ActionDnsResolverShow
	// Action Dns_resolver#Update
	Update *ActionDnsResolverUpdate
}

func NewResourceDnsResolver(client *Client) *ResourceDnsResolver {
	actionCreate := NewActionDnsResolverCreate(client)
	actionDelete := NewActionDnsResolverDelete(client)
	actionIndex := NewActionDnsResolverIndex(client)
	actionShow := NewActionDnsResolverShow(client)
	actionUpdate := NewActionDnsResolverUpdate(client)

	return &ResourceDnsResolver{
		Client:  client,
		Create:  actionCreate,
		New:     actionCreate,
		Delete:  actionDelete,
		Destroy: actionDelete,
		Index:   actionIndex,
		List:    actionIndex,
		Show:    actionShow,
		Find:    actionShow,
		Update:  actionUpdate,
	}
}
