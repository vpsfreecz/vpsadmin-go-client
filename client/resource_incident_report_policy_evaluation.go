package client

// Type for resource Incident_report.Policy_evaluation
type ResourceIncidentReportPolicyEvaluation struct {
	// Pointer to client
	Client *Client

	// Action Incident_report.Policy_evaluation#Index
	Index *ActionIncidentReportPolicyEvaluationIndex
}

func NewResourceIncidentReportPolicyEvaluation(client *Client) *ResourceIncidentReportPolicyEvaluation {
	actionIndex := NewActionIncidentReportPolicyEvaluationIndex(client)

	return &ResourceIncidentReportPolicyEvaluation{
		Client: client,
		Index:  actionIndex,
	}
}
