package client

// Type for resource Vps.Incident_policy
type ResourceVpsIncidentPolicy struct {
	// Pointer to client
	Client *Client

	// Action Vps.Incident_policy#Index
	Index *ActionVpsIncidentPolicyIndex
	// Action Vps.Incident_policy#Show
	Show *ActionVpsIncidentPolicyShow
	// Action Vps.Incident_policy#Update
	Update *ActionVpsIncidentPolicyUpdate
}

func NewResourceVpsIncidentPolicy(client *Client) *ResourceVpsIncidentPolicy {
	actionIndex := NewActionVpsIncidentPolicyIndex(client)
	actionShow := NewActionVpsIncidentPolicyShow(client)
	actionUpdate := NewActionVpsIncidentPolicyUpdate(client)

	return &ResourceVpsIncidentPolicy{
		Client: client,
		Index:  actionIndex,
		Show:   actionShow,
		Update: actionUpdate,
	}
}
