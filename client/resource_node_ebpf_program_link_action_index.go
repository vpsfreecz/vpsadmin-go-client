package client

import ()

// ActionNodeEbpfProgramLinkIndex is a type for action Node_ebpf_program_link#Index
type ActionNodeEbpfProgramLinkIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeEbpfProgramLinkIndex(client *Client) *ActionNodeEbpfProgramLinkIndex {
	return &ActionNodeEbpfProgramLinkIndex{
		Client: client,
	}
}

// ActionNodeEbpfProgramLinkIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeEbpfProgramLinkIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeEbpfProgramLinkIndexMetaGlobalInput) SetCount(value bool) *ActionNodeEbpfProgramLinkIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeEbpfProgramLinkIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeEbpfProgramLinkIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeEbpfProgramLinkIndexMetaGlobalInput) SetNo(value bool) *ActionNodeEbpfProgramLinkIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeEbpfProgramLinkIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeEbpfProgramLinkIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeEbpfProgramLinkIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeEbpfProgramLinkIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeEbpfProgramLinkIndexInput is a type for action input parameters
type ActionNodeEbpfProgramLinkIndexInput struct {
	From               string "json:\"from\""
	FromId             int64  "json:\"from_id\""
	Limit              int64  "json:\"limit\""
	Name               string "json:\"name\""
	Node               int64  "json:\"node\""
	NodeActive         bool   "json:\"node_active\""
	NodeEbpfProgram    int64  "json:\"node_ebpf_program\""
	NodeKernelEvidence int64  "json:\"node_kernel_evidence\""
	Source             string "json:\"source\""
	To                 string "json:\"to\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionNodeEbpfProgramLinkIndexInput) SetFrom(value string) *ActionNodeEbpfProgramLinkIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNodeEbpfProgramLinkIndexInput) SetFromId(value int64) *ActionNodeEbpfProgramLinkIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeEbpfProgramLinkIndexInput) SetLimit(value int64) *ActionNodeEbpfProgramLinkIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionNodeEbpfProgramLinkIndexInput) SetName(value string) *ActionNodeEbpfProgramLinkIndexInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNodeEbpfProgramLinkIndexInput) SetNode(value int64) *ActionNodeEbpfProgramLinkIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeActive sets parameter NodeActive to value and selects it for sending
func (in *ActionNodeEbpfProgramLinkIndexInput) SetNodeActive(value bool) *ActionNodeEbpfProgramLinkIndexInput {
	in.NodeActive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeActive"] = nil
	return in
}

// SetNodeEbpfProgram sets parameter NodeEbpfProgram to value and selects it for sending
func (in *ActionNodeEbpfProgramLinkIndexInput) SetNodeEbpfProgram(value int64) *ActionNodeEbpfProgramLinkIndexInput {
	in.NodeEbpfProgram = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeEbpfProgram"] = nil
	return in
}

// SetNodeKernelEvidence sets parameter NodeKernelEvidence to value and selects it for sending
func (in *ActionNodeEbpfProgramLinkIndexInput) SetNodeKernelEvidence(value int64) *ActionNodeEbpfProgramLinkIndexInput {
	in.NodeKernelEvidence = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeKernelEvidence"] = nil
	return in
}

// SetSource sets parameter Source to value and selects it for sending
func (in *ActionNodeEbpfProgramLinkIndexInput) SetSource(value string) *ActionNodeEbpfProgramLinkIndexInput {
	in.Source = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Source"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNodeEbpfProgramLinkIndexInput) SetTo(value string) *ActionNodeEbpfProgramLinkIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeEbpfProgramLinkIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeEbpfProgramLinkIndexInput) SelectParameters(params ...string) *ActionNodeEbpfProgramLinkIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeEbpfProgramLinkIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeEbpfProgramLinkIndexInput) UnselectParameters(params ...string) *ActionNodeEbpfProgramLinkIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeEbpfProgramLinkIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeEbpfProgramLinkIndexOutput is a type for action output parameters
type ActionNodeEbpfProgramLinkIndexOutput struct {
	Attached           bool                                "json:\"attached\""
	Id                 int64                               "json:\"id\""
	Name               string                              "json:\"name\""
	Node               *ActionNodeShowOutput               "json:\"node\""
	NodeEbpfProgram    *ActionNodeEbpfProgramShowOutput    "json:\"node_ebpf_program\""
	NodeKernelEvidence *ActionNodeKernelEvidenceShowOutput "json:\"node_kernel_evidence\""
	ObservedAt         string                              "json:\"observed_at\""
	Source             string                              "json:\"source\""
	SourceRevision     string                              "json:\"source_revision\""
}

// Type for action response, including envelope
type ActionNodeEbpfProgramLinkIndexResponse struct {
	Action *ActionNodeEbpfProgramLinkIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeEbpfProgramLinks []*ActionNodeEbpfProgramLinkIndexOutput "json:\"node_ebpf_program_links\""
	}

	// Action output without the namespace
	Output []*ActionNodeEbpfProgramLinkIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeEbpfProgramLinkIndex) Prepare() *ActionNodeEbpfProgramLinkIndexInvocation {
	return &ActionNodeEbpfProgramLinkIndexInvocation{
		Action: action,
		Path:   "/v7.0/node_ebpf_program_links",
	}
}

// ActionNodeEbpfProgramLinkIndexInvocation is used to configure action for invocation
type ActionNodeEbpfProgramLinkIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeEbpfProgramLinkIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeEbpfProgramLinkIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeEbpfProgramLinkIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeEbpfProgramLinkIndexInvocation) NewInput() *ActionNodeEbpfProgramLinkIndexInput {
	inv.Input = &ActionNodeEbpfProgramLinkIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeEbpfProgramLinkIndexInvocation) SetInput(input *ActionNodeEbpfProgramLinkIndexInput) *ActionNodeEbpfProgramLinkIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeEbpfProgramLinkIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeEbpfProgramLinkIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeEbpfProgramLinkIndexInvocation) NewMetaInput() *ActionNodeEbpfProgramLinkIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeEbpfProgramLinkIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeEbpfProgramLinkIndexInvocation) SetMetaInput(input *ActionNodeEbpfProgramLinkIndexMetaGlobalInput) *ActionNodeEbpfProgramLinkIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeEbpfProgramLinkIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeEbpfProgramLinkIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeEbpfProgramLinkIndexInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("From") {
			if !inv.IsParameterNil("From") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.From)
				if !ok {
					verr.Add("from", "not a valid datetime")
				} else {
					inv.Input.From = normalized
				}
			}
		}
		if inv.IsParameterSelected("Node") {
			if !inv.IsParameterNil("Node") {
				if inv.Input.Node < 0 {
					verr.Add("node", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("NodeEbpfProgram") {
			if !inv.IsParameterNil("NodeEbpfProgram") {
				if inv.Input.NodeEbpfProgram < 0 {
					verr.Add("node_ebpf_program", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("NodeKernelEvidence") {
			if !inv.IsParameterNil("NodeKernelEvidence") {
				if inv.Input.NodeKernelEvidence < 0 {
					verr.Add("node_kernel_evidence", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("To") {
			if !inv.IsParameterNil("To") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.To)
				if !ok {
					verr.Add("to", "not a valid datetime")
				} else {
					inv.Input.To = normalized
				}
			}
		}
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeEbpfProgramLinkIndexInvocation) Call() (*ActionNodeEbpfProgramLinkIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeEbpfProgramLinkIndexInvocation) callAsQuery() (*ActionNodeEbpfProgramLinkIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionNodeEbpfProgramLinkIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeEbpfProgramLinks
	}
	return resp, err
}

func (inv *ActionNodeEbpfProgramLinkIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("From") {
			ret["node_ebpf_program_link[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("FromId") {
			ret["node_ebpf_program_link[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["node_ebpf_program_link[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Name") {
			ret["node_ebpf_program_link[name]"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Node") {
			ret["node_ebpf_program_link[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("NodeActive") {
			ret["node_ebpf_program_link[node_active]"] = convertBoolToString(inv.Input.NodeActive)
		}
		if inv.IsParameterSelected("NodeEbpfProgram") {
			ret["node_ebpf_program_link[node_ebpf_program]"] = convertInt64ToString(inv.Input.NodeEbpfProgram)
		}
		if inv.IsParameterSelected("NodeKernelEvidence") {
			ret["node_ebpf_program_link[node_kernel_evidence]"] = convertInt64ToString(inv.Input.NodeKernelEvidence)
		}
		if inv.IsParameterSelected("Source") {
			ret["node_ebpf_program_link[source]"] = inv.Input.Source
		}
		if inv.IsParameterSelected("To") {
			ret["node_ebpf_program_link[to]"] = inv.Input.To
		}
	}

	return nil
}

func (inv *ActionNodeEbpfProgramLinkIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			queryValue, err := convertCustomToString(inv.MetaInput.Includes)
			if err != nil {
				return err
			}
			ret["_meta[includes]"] = queryValue
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}

	return nil
}
