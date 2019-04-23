package client

// Type for resource Node
type ResourceNode struct {
	// Pointer to client
	Client *Client

	// Resource Node.Status
	Status *ResourceNodeStatus
	// Action Node#Index
	Index *ActionNodeIndex
	// Action Node#Index
	List *ActionNodeIndex
	// Action Node#Create
	Create *ActionNodeCreate
	// Action Node#Create
	New *ActionNodeCreate
	// Action Node#Overview_list
	OverviewList *ActionNodeOverviewList
	// Action Node#Public_status
	PublicStatus *ActionNodePublicStatus
	// Action Node#Show
	Show *ActionNodeShow
	// Action Node#Show
	Find *ActionNodeShow
	// Action Node#Update
	Update *ActionNodeUpdate
	// Action Node#Evacuate
	Evacuate *ActionNodeEvacuate
	// Action Node#Set_maintenance
	SetMaintenance *ActionNodeSetMaintenance
}

func NewResourceNode(client *Client) *ResourceNode {
	actionIndex := NewActionNodeIndex(client)
	actionCreate := NewActionNodeCreate(client)
	actionOverviewList := NewActionNodeOverviewList(client)
	actionPublicStatus := NewActionNodePublicStatus(client)
	actionShow := NewActionNodeShow(client)
	actionUpdate := NewActionNodeUpdate(client)
	actionEvacuate := NewActionNodeEvacuate(client)
	actionSetMaintenance := NewActionNodeSetMaintenance(client)

	return &ResourceNode{
		Client: client,
		Status: NewResourceNodeStatus(client),
		Index: actionIndex,
		List: actionIndex,
		Create: actionCreate,
		New: actionCreate,
		OverviewList: actionOverviewList,
		PublicStatus: actionPublicStatus,
		Show: actionShow,
		Find: actionShow,
		Update: actionUpdate,
		Evacuate: actionEvacuate,
		SetMaintenance: actionSetMaintenance,
	}
}
