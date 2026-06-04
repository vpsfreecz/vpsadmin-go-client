package client

// Type for resource Outage_security_advisory
type ResourceOutageSecurityAdvisory struct {
	// Pointer to client
	Client *Client

	// Action Outage_security_advisory#Create
	Create *ActionOutageSecurityAdvisoryCreate
	// Action Outage_security_advisory#Create
	New *ActionOutageSecurityAdvisoryCreate
	// Action Outage_security_advisory#Delete
	Delete *ActionOutageSecurityAdvisoryDelete
	// Action Outage_security_advisory#Delete
	Destroy *ActionOutageSecurityAdvisoryDelete
	// Action Outage_security_advisory#Index
	Index *ActionOutageSecurityAdvisoryIndex
	// Action Outage_security_advisory#Index
	List *ActionOutageSecurityAdvisoryIndex
	// Action Outage_security_advisory#Show
	Show *ActionOutageSecurityAdvisoryShow
	// Action Outage_security_advisory#Show
	Find *ActionOutageSecurityAdvisoryShow
}

func NewResourceOutageSecurityAdvisory(client *Client) *ResourceOutageSecurityAdvisory {
	actionCreate := NewActionOutageSecurityAdvisoryCreate(client)
	actionDelete := NewActionOutageSecurityAdvisoryDelete(client)
	actionIndex := NewActionOutageSecurityAdvisoryIndex(client)
	actionShow := NewActionOutageSecurityAdvisoryShow(client)

	return &ResourceOutageSecurityAdvisory{
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
