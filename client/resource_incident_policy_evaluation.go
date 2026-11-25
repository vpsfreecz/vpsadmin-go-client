package client

// Type for resource Incident_policy_evaluation
type ResourceIncidentPolicyEvaluation struct {
	// Pointer to client
	Client *Client

	// Action Incident_policy_evaluation#Index
	Index *ActionIncidentPolicyEvaluationIndex
}

func NewResourceIncidentPolicyEvaluation(client *Client) *ResourceIncidentPolicyEvaluation {
	actionIndex := NewActionIncidentPolicyEvaluationIndex(client)

	return &ResourceIncidentPolicyEvaluation{
		Client: client,
		Index:  actionIndex,
	}
}
