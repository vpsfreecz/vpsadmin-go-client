package client

import ()

// ActionNodeEbpfProgramObjectIndex is a type for action Node_ebpf_program_object#Index
type ActionNodeEbpfProgramObjectIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeEbpfProgramObjectIndex(client *Client) *ActionNodeEbpfProgramObjectIndex {
	return &ActionNodeEbpfProgramObjectIndex{
		Client: client,
	}
}

// ActionNodeEbpfProgramObjectIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeEbpfProgramObjectIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeEbpfProgramObjectIndexMetaGlobalInput) SetCount(value bool) *ActionNodeEbpfProgramObjectIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeEbpfProgramObjectIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeEbpfProgramObjectIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeEbpfProgramObjectIndexMetaGlobalInput) SetNo(value bool) *ActionNodeEbpfProgramObjectIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeEbpfProgramObjectIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeEbpfProgramObjectIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeEbpfProgramObjectIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeEbpfProgramObjectIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeEbpfProgramObjectIndexInput is a type for action input parameters
type ActionNodeEbpfProgramObjectIndexInput struct {
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
func (in *ActionNodeEbpfProgramObjectIndexInput) SetFrom(value string) *ActionNodeEbpfProgramObjectIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNodeEbpfProgramObjectIndexInput) SetFromId(value int64) *ActionNodeEbpfProgramObjectIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeEbpfProgramObjectIndexInput) SetLimit(value int64) *ActionNodeEbpfProgramObjectIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionNodeEbpfProgramObjectIndexInput) SetName(value string) *ActionNodeEbpfProgramObjectIndexInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNodeEbpfProgramObjectIndexInput) SetNode(value int64) *ActionNodeEbpfProgramObjectIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeActive sets parameter NodeActive to value and selects it for sending
func (in *ActionNodeEbpfProgramObjectIndexInput) SetNodeActive(value bool) *ActionNodeEbpfProgramObjectIndexInput {
	in.NodeActive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeActive"] = nil
	return in
}

// SetNodeEbpfProgram sets parameter NodeEbpfProgram to value and selects it for sending
func (in *ActionNodeEbpfProgramObjectIndexInput) SetNodeEbpfProgram(value int64) *ActionNodeEbpfProgramObjectIndexInput {
	in.NodeEbpfProgram = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeEbpfProgram"] = nil
	return in
}

// SetNodeKernelEvidence sets parameter NodeKernelEvidence to value and selects it for sending
func (in *ActionNodeEbpfProgramObjectIndexInput) SetNodeKernelEvidence(value int64) *ActionNodeEbpfProgramObjectIndexInput {
	in.NodeKernelEvidence = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeKernelEvidence"] = nil
	return in
}

// SetSource sets parameter Source to value and selects it for sending
func (in *ActionNodeEbpfProgramObjectIndexInput) SetSource(value string) *ActionNodeEbpfProgramObjectIndexInput {
	in.Source = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Source"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNodeEbpfProgramObjectIndexInput) SetTo(value string) *ActionNodeEbpfProgramObjectIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeEbpfProgramObjectIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeEbpfProgramObjectIndexInput) SelectParameters(params ...string) *ActionNodeEbpfProgramObjectIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeEbpfProgramObjectIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeEbpfProgramObjectIndexInput) UnselectParameters(params ...string) *ActionNodeEbpfProgramObjectIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeEbpfProgramObjectIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeEbpfProgramObjectIndexOutput is a type for action output parameters
type ActionNodeEbpfProgramObjectIndexOutput struct {
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
type ActionNodeEbpfProgramObjectIndexResponse struct {
	Action *ActionNodeEbpfProgramObjectIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeEbpfProgramObjects []*ActionNodeEbpfProgramObjectIndexOutput "json:\"node_ebpf_program_objects\""
	}

	// Action output without the namespace
	Output []*ActionNodeEbpfProgramObjectIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeEbpfProgramObjectIndex) Prepare() *ActionNodeEbpfProgramObjectIndexInvocation {
	return &ActionNodeEbpfProgramObjectIndexInvocation{
		Action: action,
		Path:   "/v7.0/node_ebpf_program_objects",
	}
}

// ActionNodeEbpfProgramObjectIndexInvocation is used to configure action for invocation
type ActionNodeEbpfProgramObjectIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeEbpfProgramObjectIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeEbpfProgramObjectIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeEbpfProgramObjectIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeEbpfProgramObjectIndexInvocation) NewInput() *ActionNodeEbpfProgramObjectIndexInput {
	inv.Input = &ActionNodeEbpfProgramObjectIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeEbpfProgramObjectIndexInvocation) SetInput(input *ActionNodeEbpfProgramObjectIndexInput) *ActionNodeEbpfProgramObjectIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeEbpfProgramObjectIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeEbpfProgramObjectIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeEbpfProgramObjectIndexInvocation) NewMetaInput() *ActionNodeEbpfProgramObjectIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeEbpfProgramObjectIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeEbpfProgramObjectIndexInvocation) SetMetaInput(input *ActionNodeEbpfProgramObjectIndexMetaGlobalInput) *ActionNodeEbpfProgramObjectIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeEbpfProgramObjectIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeEbpfProgramObjectIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeEbpfProgramObjectIndexInvocation) validate() error {
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
func (inv *ActionNodeEbpfProgramObjectIndexInvocation) Call() (*ActionNodeEbpfProgramObjectIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeEbpfProgramObjectIndexInvocation) callAsQuery() (*ActionNodeEbpfProgramObjectIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeEbpfProgramObjectIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeEbpfProgramObjects
	}
	return resp, err
}

func (inv *ActionNodeEbpfProgramObjectIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("From") {
			ret["node_ebpf_program_object[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("FromId") {
			ret["node_ebpf_program_object[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["node_ebpf_program_object[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Name") {
			ret["node_ebpf_program_object[name]"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Node") {
			ret["node_ebpf_program_object[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("NodeActive") {
			ret["node_ebpf_program_object[node_active]"] = convertBoolToString(inv.Input.NodeActive)
		}
		if inv.IsParameterSelected("NodeEbpfProgram") {
			ret["node_ebpf_program_object[node_ebpf_program]"] = convertInt64ToString(inv.Input.NodeEbpfProgram)
		}
		if inv.IsParameterSelected("NodeKernelEvidence") {
			ret["node_ebpf_program_object[node_kernel_evidence]"] = convertInt64ToString(inv.Input.NodeKernelEvidence)
		}
		if inv.IsParameterSelected("Source") {
			ret["node_ebpf_program_object[source]"] = inv.Input.Source
		}
		if inv.IsParameterSelected("To") {
			ret["node_ebpf_program_object[to]"] = inv.Input.To
		}
	}
}

func (inv *ActionNodeEbpfProgramObjectIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
