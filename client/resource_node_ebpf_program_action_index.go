package client

import ()

// ActionNodeEbpfProgramIndex is a type for action Node_ebpf_program#Index
type ActionNodeEbpfProgramIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeEbpfProgramIndex(client *Client) *ActionNodeEbpfProgramIndex {
	return &ActionNodeEbpfProgramIndex{
		Client: client,
	}
}

// ActionNodeEbpfProgramIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeEbpfProgramIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeEbpfProgramIndexMetaGlobalInput) SetCount(value bool) *ActionNodeEbpfProgramIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeEbpfProgramIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeEbpfProgramIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeEbpfProgramIndexMetaGlobalInput) SetNo(value bool) *ActionNodeEbpfProgramIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeEbpfProgramIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeEbpfProgramIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeEbpfProgramIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeEbpfProgramIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeEbpfProgramIndexInput is a type for action input parameters
type ActionNodeEbpfProgramIndexInput struct {
	Active             bool   "json:\"active\""
	From               string "json:\"from\""
	FromId             int64  "json:\"from_id\""
	Limit              int64  "json:\"limit\""
	Name               string "json:\"name\""
	Node               int64  "json:\"node\""
	NodeActive         bool   "json:\"node_active\""
	NodeKernelEvidence int64  "json:\"node_kernel_evidence\""
	Source             string "json:\"source\""
	To                 string "json:\"to\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetActive sets parameter Active to value and selects it for sending
func (in *ActionNodeEbpfProgramIndexInput) SetActive(value bool) *ActionNodeEbpfProgramIndexInput {
	in.Active = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Active"] = nil
	return in
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionNodeEbpfProgramIndexInput) SetFrom(value string) *ActionNodeEbpfProgramIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNodeEbpfProgramIndexInput) SetFromId(value int64) *ActionNodeEbpfProgramIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeEbpfProgramIndexInput) SetLimit(value int64) *ActionNodeEbpfProgramIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionNodeEbpfProgramIndexInput) SetName(value string) *ActionNodeEbpfProgramIndexInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNodeEbpfProgramIndexInput) SetNode(value int64) *ActionNodeEbpfProgramIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeActive sets parameter NodeActive to value and selects it for sending
func (in *ActionNodeEbpfProgramIndexInput) SetNodeActive(value bool) *ActionNodeEbpfProgramIndexInput {
	in.NodeActive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeActive"] = nil
	return in
}

// SetNodeKernelEvidence sets parameter NodeKernelEvidence to value and selects it for sending
func (in *ActionNodeEbpfProgramIndexInput) SetNodeKernelEvidence(value int64) *ActionNodeEbpfProgramIndexInput {
	in.NodeKernelEvidence = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeKernelEvidence"] = nil
	return in
}

// SetSource sets parameter Source to value and selects it for sending
func (in *ActionNodeEbpfProgramIndexInput) SetSource(value string) *ActionNodeEbpfProgramIndexInput {
	in.Source = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Source"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNodeEbpfProgramIndexInput) SetTo(value string) *ActionNodeEbpfProgramIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeEbpfProgramIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeEbpfProgramIndexInput) SelectParameters(params ...string) *ActionNodeEbpfProgramIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeEbpfProgramIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeEbpfProgramIndexInput) UnselectParameters(params ...string) *ActionNodeEbpfProgramIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeEbpfProgramIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeEbpfProgramIndexOutput is a type for action output parameters
type ActionNodeEbpfProgramIndexOutput struct {
	Active             bool                                "json:\"active\""
	AttachedAt         string                              "json:\"attached_at\""
	Description        string                              "json:\"description\""
	Digest             string                              "json:\"digest\""
	Id                 int64                               "json:\"id\""
	Name               string                              "json:\"name\""
	Node               *ActionNodeShowOutput               "json:\"node\""
	NodeKernelEvidence *ActionNodeKernelEvidenceShowOutput "json:\"node_kernel_evidence\""
	ObservedAt         string                              "json:\"observed_at\""
	Revision           string                              "json:\"revision\""
	SinceKernel        string                              "json:\"since_kernel\""
	Source             string                              "json:\"source\""
	SourceRevision     string                              "json:\"source_revision\""
	UntilKernel        string                              "json:\"until_kernel\""
	VerifiedAt         string                              "json:\"verified_at\""
}

// Type for action response, including envelope
type ActionNodeEbpfProgramIndexResponse struct {
	Action *ActionNodeEbpfProgramIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeEbpfPrograms []*ActionNodeEbpfProgramIndexOutput "json:\"node_ebpf_programs\""
	}

	// Action output without the namespace
	Output []*ActionNodeEbpfProgramIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeEbpfProgramIndex) Prepare() *ActionNodeEbpfProgramIndexInvocation {
	return &ActionNodeEbpfProgramIndexInvocation{
		Action: action,
		Path:   "/v7.0/node_ebpf_programs",
	}
}

// ActionNodeEbpfProgramIndexInvocation is used to configure action for invocation
type ActionNodeEbpfProgramIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeEbpfProgramIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeEbpfProgramIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeEbpfProgramIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeEbpfProgramIndexInvocation) NewInput() *ActionNodeEbpfProgramIndexInput {
	inv.Input = &ActionNodeEbpfProgramIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeEbpfProgramIndexInvocation) SetInput(input *ActionNodeEbpfProgramIndexInput) *ActionNodeEbpfProgramIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeEbpfProgramIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeEbpfProgramIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeEbpfProgramIndexInvocation) NewMetaInput() *ActionNodeEbpfProgramIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeEbpfProgramIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeEbpfProgramIndexInvocation) SetMetaInput(input *ActionNodeEbpfProgramIndexMetaGlobalInput) *ActionNodeEbpfProgramIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeEbpfProgramIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeEbpfProgramIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeEbpfProgramIndexInvocation) validate() error {
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
func (inv *ActionNodeEbpfProgramIndexInvocation) Call() (*ActionNodeEbpfProgramIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeEbpfProgramIndexInvocation) callAsQuery() (*ActionNodeEbpfProgramIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionNodeEbpfProgramIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeEbpfPrograms
	}
	return resp, err
}

func (inv *ActionNodeEbpfProgramIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("Active") {
			ret["node_ebpf_program[active]"] = convertBoolToString(inv.Input.Active)
		}
		if inv.IsParameterSelected("From") {
			ret["node_ebpf_program[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("FromId") {
			ret["node_ebpf_program[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["node_ebpf_program[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Name") {
			ret["node_ebpf_program[name]"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Node") {
			ret["node_ebpf_program[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("NodeActive") {
			ret["node_ebpf_program[node_active]"] = convertBoolToString(inv.Input.NodeActive)
		}
		if inv.IsParameterSelected("NodeKernelEvidence") {
			ret["node_ebpf_program[node_kernel_evidence]"] = convertInt64ToString(inv.Input.NodeKernelEvidence)
		}
		if inv.IsParameterSelected("Source") {
			ret["node_ebpf_program[source]"] = inv.Input.Source
		}
		if inv.IsParameterSelected("To") {
			ret["node_ebpf_program[to]"] = inv.Input.To
		}
	}

	return nil
}

func (inv *ActionNodeEbpfProgramIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
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
