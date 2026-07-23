package client

import ()

// ActionNodeSysctlIndex is a type for action Node_sysctl#Index
type ActionNodeSysctlIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeSysctlIndex(client *Client) *ActionNodeSysctlIndex {
	return &ActionNodeSysctlIndex{
		Client: client,
	}
}

// ActionNodeSysctlIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeSysctlIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeSysctlIndexMetaGlobalInput) SetCount(value bool) *ActionNodeSysctlIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeSysctlIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeSysctlIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeSysctlIndexMetaGlobalInput) SetNo(value bool) *ActionNodeSysctlIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeSysctlIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeSysctlIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeSysctlIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeSysctlIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeSysctlIndexInput is a type for action input parameters
type ActionNodeSysctlIndexInput struct {
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

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionNodeSysctlIndexInput) SetFrom(value string) *ActionNodeSysctlIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNodeSysctlIndexInput) SetFromId(value int64) *ActionNodeSysctlIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeSysctlIndexInput) SetLimit(value int64) *ActionNodeSysctlIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionNodeSysctlIndexInput) SetName(value string) *ActionNodeSysctlIndexInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNodeSysctlIndexInput) SetNode(value int64) *ActionNodeSysctlIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeActive sets parameter NodeActive to value and selects it for sending
func (in *ActionNodeSysctlIndexInput) SetNodeActive(value bool) *ActionNodeSysctlIndexInput {
	in.NodeActive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeActive"] = nil
	return in
}

// SetNodeKernelEvidence sets parameter NodeKernelEvidence to value and selects it for sending
func (in *ActionNodeSysctlIndexInput) SetNodeKernelEvidence(value int64) *ActionNodeSysctlIndexInput {
	in.NodeKernelEvidence = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeKernelEvidence"] = nil
	return in
}

// SetSource sets parameter Source to value and selects it for sending
func (in *ActionNodeSysctlIndexInput) SetSource(value string) *ActionNodeSysctlIndexInput {
	in.Source = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Source"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNodeSysctlIndexInput) SetTo(value string) *ActionNodeSysctlIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeSysctlIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeSysctlIndexInput) SelectParameters(params ...string) *ActionNodeSysctlIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeSysctlIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeSysctlIndexInput) UnselectParameters(params ...string) *ActionNodeSysctlIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeSysctlIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeSysctlIndexOutput is a type for action output parameters
type ActionNodeSysctlIndexOutput struct {
	Available          bool                                "json:\"available\""
	ConfiguredValue    string                              "json:\"configured_value\""
	EffectiveValue     string                              "json:\"effective_value\""
	Id                 int64                               "json:\"id\""
	Name               string                              "json:\"name\""
	Node               *ActionNodeShowOutput               "json:\"node\""
	NodeKernelEvidence *ActionNodeKernelEvidenceShowOutput "json:\"node_kernel_evidence\""
	ObservedAt         string                              "json:\"observed_at\""
	Source             string                              "json:\"source\""
	SourceRevision     string                              "json:\"source_revision\""
}

// Type for action response, including envelope
type ActionNodeSysctlIndexResponse struct {
	Action *ActionNodeSysctlIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeSysctls []*ActionNodeSysctlIndexOutput "json:\"node_sysctls\""
	}

	// Action output without the namespace
	Output []*ActionNodeSysctlIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeSysctlIndex) Prepare() *ActionNodeSysctlIndexInvocation {
	return &ActionNodeSysctlIndexInvocation{
		Action: action,
		Path:   "/v7.0/node_sysctls",
	}
}

// ActionNodeSysctlIndexInvocation is used to configure action for invocation
type ActionNodeSysctlIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeSysctlIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeSysctlIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeSysctlIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeSysctlIndexInvocation) NewInput() *ActionNodeSysctlIndexInput {
	inv.Input = &ActionNodeSysctlIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeSysctlIndexInvocation) SetInput(input *ActionNodeSysctlIndexInput) *ActionNodeSysctlIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeSysctlIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeSysctlIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeSysctlIndexInvocation) NewMetaInput() *ActionNodeSysctlIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeSysctlIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeSysctlIndexInvocation) SetMetaInput(input *ActionNodeSysctlIndexMetaGlobalInput) *ActionNodeSysctlIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeSysctlIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeSysctlIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeSysctlIndexInvocation) validate() error {
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
func (inv *ActionNodeSysctlIndexInvocation) Call() (*ActionNodeSysctlIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeSysctlIndexInvocation) callAsQuery() (*ActionNodeSysctlIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeSysctlIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeSysctls
	}
	return resp, err
}

func (inv *ActionNodeSysctlIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("From") {
			ret["node_sysctl[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("FromId") {
			ret["node_sysctl[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["node_sysctl[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Name") {
			ret["node_sysctl[name]"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Node") {
			ret["node_sysctl[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("NodeActive") {
			ret["node_sysctl[node_active]"] = convertBoolToString(inv.Input.NodeActive)
		}
		if inv.IsParameterSelected("NodeKernelEvidence") {
			ret["node_sysctl[node_kernel_evidence]"] = convertInt64ToString(inv.Input.NodeKernelEvidence)
		}
		if inv.IsParameterSelected("Source") {
			ret["node_sysctl[source]"] = inv.Input.Source
		}
		if inv.IsParameterSelected("To") {
			ret["node_sysctl[to]"] = inv.Input.To
		}
	}
}

func (inv *ActionNodeSysctlIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
