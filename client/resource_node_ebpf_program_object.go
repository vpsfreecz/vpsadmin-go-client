package client

// Type for resource Node_ebpf_program_object
type ResourceNodeEbpfProgramObject struct {
	// Pointer to client
	Client *Client

	// Action Node_ebpf_program_object#Index
	Index *ActionNodeEbpfProgramObjectIndex
	// Action Node_ebpf_program_object#Index
	List *ActionNodeEbpfProgramObjectIndex
}

func NewResourceNodeEbpfProgramObject(client *Client) *ResourceNodeEbpfProgramObject {
	actionIndex := NewActionNodeEbpfProgramObjectIndex(client)

	return &ResourceNodeEbpfProgramObject{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
	}
}
