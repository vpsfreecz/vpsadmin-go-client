package client

// Type for resource Incident_report
type ResourceIncidentReport struct {
	// Pointer to client
	Client *Client

	// Action Incident_report#Create
	Create *ActionIncidentReportCreate
	// Action Incident_report#Create
	New *ActionIncidentReportCreate
	// Action Incident_report#Index
	Index *ActionIncidentReportIndex
	// Action Incident_report#Index
	List *ActionIncidentReportIndex
	// Action Incident_report#Show
	Show *ActionIncidentReportShow
	// Action Incident_report#Show
	Find *ActionIncidentReportShow
}

func NewResourceIncidentReport(client *Client) *ResourceIncidentReport {
	actionCreate := NewActionIncidentReportCreate(client)
	actionIndex := NewActionIncidentReportIndex(client)
	actionShow := NewActionIncidentReportShow(client)

	return &ResourceIncidentReport{
		Client: client,
		Create: actionCreate,
		New:    actionCreate,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
