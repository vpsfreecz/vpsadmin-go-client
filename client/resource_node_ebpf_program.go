package client

// Type for resource Node_ebpf_program
type ResourceNodeEbpfProgram struct {
	// Pointer to client
	Client *Client

	// Action Node_ebpf_program#Index
	Index *ActionNodeEbpfProgramIndex
	// Action Node_ebpf_program#Index
	List *ActionNodeEbpfProgramIndex
	// Action Node_ebpf_program#Show
	Show *ActionNodeEbpfProgramShow
	// Action Node_ebpf_program#Show
	Find *ActionNodeEbpfProgramShow
}

func NewResourceNodeEbpfProgram(client *Client) *ResourceNodeEbpfProgram {
	actionIndex := NewActionNodeEbpfProgramIndex(client)
	actionShow := NewActionNodeEbpfProgramShow(client)

	return &ResourceNodeEbpfProgram{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
