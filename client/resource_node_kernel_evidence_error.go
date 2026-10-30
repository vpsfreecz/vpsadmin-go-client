package client

// Type for resource Node_kernel_evidence_error
type ResourceNodeKernelEvidenceError struct {
	// Pointer to client
	Client *Client

	// Action Node_kernel_evidence_error#Index
	Index *ActionNodeKernelEvidenceErrorIndex
	// Action Node_kernel_evidence_error#Index
	List *ActionNodeKernelEvidenceErrorIndex
}

func NewResourceNodeKernelEvidenceError(client *Client) *ResourceNodeKernelEvidenceError {
	actionIndex := NewActionNodeKernelEvidenceErrorIndex(client)

	return &ResourceNodeKernelEvidenceError{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
