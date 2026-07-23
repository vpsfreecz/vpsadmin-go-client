package client

// Type for resource Node_kernel_evidence
type ResourceNodeKernelEvidence struct {
	// Pointer to client
	Client *Client

	// Action Node_kernel_evidence#Index
	Index *ActionNodeKernelEvidenceIndex
	// Action Node_kernel_evidence#Index
	List *ActionNodeKernelEvidenceIndex
	// Action Node_kernel_evidence#Show
	Show *ActionNodeKernelEvidenceShow
	// Action Node_kernel_evidence#Show
	Find *ActionNodeKernelEvidenceShow
}

func NewResourceNodeKernelEvidence(client *Client) *ResourceNodeKernelEvidence {
	actionIndex := NewActionNodeKernelEvidenceIndex(client)
	actionShow := NewActionNodeKernelEvidenceShow(client)

	return &ResourceNodeKernelEvidence{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
