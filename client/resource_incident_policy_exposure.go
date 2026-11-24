package client

// Type for resource Incident_policy_exposure
type ResourceIncidentPolicyExposure struct {
	// Pointer to client
	Client *Client

	// Action Incident_policy_exposure#Index
	Index *ActionIncidentPolicyExposureIndex
}

func NewResourceIncidentPolicyExposure(client *Client) *ResourceIncidentPolicyExposure {
	actionIndex := NewActionIncidentPolicyExposureIndex(client)

	return &ResourceIncidentPolicyExposure{
		Client: client,
		Index:  actionIndex,
	}
}
