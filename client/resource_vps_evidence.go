package client

// Type for resource Vps_evidence
type ResourceVpsEvidence struct {
	// Pointer to client
	Client *Client

	// Action Vps_evidence#Index
	Index *ActionVpsEvidenceIndex
}

func NewResourceVpsEvidence(client *Client) *ResourceVpsEvidence {
	actionIndex := NewActionVpsEvidenceIndex(client)

	return &ResourceVpsEvidence{
		Client: client,
		Index:  actionIndex,
	}
}
