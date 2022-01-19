package client

// Type for resource Node
type ResourceNode struct {
	// Pointer to client
	Client *Client

	// Resource Node.Status
	Status *ResourceNodeStatus
	// Action Node#Create
	Create *ActionNodeCreate
	// Action Node#Create
	New *ActionNodeCreate
	// Action Node#Evacuate
	Evacuate *ActionNodeEvacuate
	// Action Node#Index
	Index *ActionNodeIndex
	// Action Node#Index
	List *ActionNodeIndex
	// Action Node#Overview_list
	OverviewList *ActionNodeOverviewList
	// Action Node#Public_status
	PublicStatus *ActionNodePublicStatus
	// Action Node#Set_maintenance
	SetMaintenance *ActionNodeSetMaintenance
	// Action Node#Show
	Show *ActionNodeShow
	// Action Node#Show
	Find *ActionNodeShow
	// Action Node#Update
	Update *ActionNodeUpdate
}

func NewResourceNode(client *Client) *ResourceNode {
	actionCreate := NewActionNodeCreate(client)
	actionEvacuate := NewActionNodeEvacuate(client)
	actionIndex := NewActionNodeIndex(client)
	actionOverviewList := NewActionNodeOverviewList(client)
	actionPublicStatus := NewActionNodePublicStatus(client)
	actionSetMaintenance := NewActionNodeSetMaintenance(client)
	actionShow := NewActionNodeShow(client)
	actionUpdate := NewActionNodeUpdate(client)

	return &ResourceNode{
		Client:         client,
		Status:         NewResourceNodeStatus(client),
		Create:         actionCreate,
		New:            actionCreate,
		Evacuate:       actionEvacuate,
		Index:          actionIndex,
		List:           actionIndex,
		OverviewList:   actionOverviewList,
		PublicStatus:   actionPublicStatus,
		SetMaintenance: actionSetMaintenance,
		Show:           actionShow,
		Find:           actionShow,
		Update:         actionUpdate,
	}
}
