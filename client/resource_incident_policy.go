package client

// Type for resource Incident_policy
type ResourceIncidentPolicy struct {
	// Pointer to client
	Client *Client

	// Action Incident_policy#Index
	Index *ActionIncidentPolicyIndex
	// Action Incident_policy#Show
	Show *ActionIncidentPolicyShow
}

func NewResourceIncidentPolicy(client *Client) *ResourceIncidentPolicy {
	actionIndex := NewActionIncidentPolicyIndex(client)
	actionShow := NewActionIncidentPolicyShow(client)

	return &ResourceIncidentPolicy{
		Client: client,
		Index:  actionIndex,
		Show:   actionShow,
	}
}
