package client

// Type for resource Security_advisory_cve
type ResourceSecurityAdvisoryCve struct {
	// Pointer to client
	Client *Client

	// Action Security_advisory_cve#Create
	Create *ActionSecurityAdvisoryCveCreate
	// Action Security_advisory_cve#Create
	New *ActionSecurityAdvisoryCveCreate
	// Action Security_advisory_cve#Delete
	Delete *ActionSecurityAdvisoryCveDelete
	// Action Security_advisory_cve#Delete
	Destroy *ActionSecurityAdvisoryCveDelete
	// Action Security_advisory_cve#Index
	Index *ActionSecurityAdvisoryCveIndex
	// Action Security_advisory_cve#Index
	List *ActionSecurityAdvisoryCveIndex
	// Action Security_advisory_cve#Show
	Show *ActionSecurityAdvisoryCveShow
	// Action Security_advisory_cve#Show
	Find *ActionSecurityAdvisoryCveShow
	// Action Security_advisory_cve#Update
	Update *ActionSecurityAdvisoryCveUpdate
}

func NewResourceSecurityAdvisoryCve(client *Client) *ResourceSecurityAdvisoryCve {
	actionCreate := NewActionSecurityAdvisoryCveCreate(client)
	actionDelete := NewActionSecurityAdvisoryCveDelete(client)
	actionIndex := NewActionSecurityAdvisoryCveIndex(client)
	actionShow := NewActionSecurityAdvisoryCveShow(client)
	actionUpdate := NewActionSecurityAdvisoryCveUpdate(client)

	return &ResourceSecurityAdvisoryCve{
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
