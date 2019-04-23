package client

// Type for resource Dns_resolver
type ResourceDnsResolver struct {
	// Pointer to client
	Client *Client

	// Action Dns_resolver#Index
	Index *ActionDnsResolverIndex
	// Action Dns_resolver#Index
	List *ActionDnsResolverIndex
	// Action Dns_resolver#Show
	Show *ActionDnsResolverShow
	// Action Dns_resolver#Show
	Find *ActionDnsResolverShow
	// Action Dns_resolver#Create
	Create *ActionDnsResolverCreate
	// Action Dns_resolver#Create
	New *ActionDnsResolverCreate
	// Action Dns_resolver#Update
	Update *ActionDnsResolverUpdate
	// Action Dns_resolver#Delete
	Delete *ActionDnsResolverDelete
	// Action Dns_resolver#Delete
	Destroy *ActionDnsResolverDelete
}

func NewResourceDnsResolver(client *Client) *ResourceDnsResolver {
	actionIndex := NewActionDnsResolverIndex(client)
	actionShow := NewActionDnsResolverShow(client)
	actionCreate := NewActionDnsResolverCreate(client)
	actionUpdate := NewActionDnsResolverUpdate(client)
	actionDelete := NewActionDnsResolverDelete(client)

	return &ResourceDnsResolver{
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
