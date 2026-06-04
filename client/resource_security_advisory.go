package client

// Type for resource Security_advisory
type ResourceSecurityAdvisory struct {
	// Pointer to client
	Client *Client

	// Resource Security_advisory.Node_status
	NodeStatus *ResourceSecurityAdvisoryNodeStatus
	// Action Security_advisory#Create
	Create *ActionSecurityAdvisoryCreate
	// Action Security_advisory#Create
	New *ActionSecurityAdvisoryCreate
	// Action Security_advisory#Index
	Index *ActionSecurityAdvisoryIndex
	// Action Security_advisory#Index
	List *ActionSecurityAdvisoryIndex
	// Action Security_advisory#Publish
	Publish *ActionSecurityAdvisoryPublish
	// Action Security_advisory#Rebuild_affected_vps
	RebuildAffectedVps *ActionSecurityAdvisoryRebuildAffectedVps
	// Action Security_advisory#Show
	Show *ActionSecurityAdvisoryShow
	// Action Security_advisory#Show
	Find *ActionSecurityAdvisoryShow
	// Action Security_advisory#Update
	Update *ActionSecurityAdvisoryUpdate
}

func NewResourceSecurityAdvisory(client *Client) *ResourceSecurityAdvisory {
	actionCreate := NewActionSecurityAdvisoryCreate(client)
	actionIndex := NewActionSecurityAdvisoryIndex(client)
	actionPublish := NewActionSecurityAdvisoryPublish(client)
	actionRebuildAffectedVps := NewActionSecurityAdvisoryRebuildAffectedVps(client)
	actionShow := NewActionSecurityAdvisoryShow(client)
	actionUpdate := NewActionSecurityAdvisoryUpdate(client)

	return &ResourceSecurityAdvisory{
		Client:             client,
		NodeStatus:         NewResourceSecurityAdvisoryNodeStatus(client),
		Create:             actionCreate,
		New:                actionCreate,
		Index:              actionIndex,
		List:               actionIndex,
		Publish:            actionPublish,
		RebuildAffectedVps: actionRebuildAffectedVps,
		Show:               actionShow,
		Find:               actionShow,
		Update:             actionUpdate,
	}
}
