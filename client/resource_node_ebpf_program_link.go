package client

// Type for resource Node_ebpf_program_link
type ResourceNodeEbpfProgramLink struct {
	// Pointer to client
	Client *Client

	// Action Node_ebpf_program_link#Index
	Index *ActionNodeEbpfProgramLinkIndex
	// Action Node_ebpf_program_link#Index
	List *ActionNodeEbpfProgramLinkIndex
}

func NewResourceNodeEbpfProgramLink(client *Client) *ResourceNodeEbpfProgramLink {
	actionIndex := NewActionNodeEbpfProgramLinkIndex(client)

	return &ResourceNodeEbpfProgramLink{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
