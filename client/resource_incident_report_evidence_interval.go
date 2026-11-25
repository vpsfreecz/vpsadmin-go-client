package client

// Type for resource Incident_report.Evidence_interval
type ResourceIncidentReportEvidenceInterval struct {
	// Pointer to client
	Client *Client

	// Action Incident_report.Evidence_interval#Index
	Index *ActionIncidentReportEvidenceIntervalIndex
}

func NewResourceIncidentReportEvidenceInterval(client *Client) *ResourceIncidentReportEvidenceInterval {
	actionIndex := NewActionIncidentReportEvidenceIntervalIndex(client)

	return &ResourceIncidentReportEvidenceInterval{
		Client: client,
		Index:  actionIndex,
	}
}
