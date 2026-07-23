package client

import ()

// ActionNodeKernelEvidenceIndex is a type for action Node_kernel_evidence#Index
type ActionNodeKernelEvidenceIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeKernelEvidenceIndex(client *Client) *ActionNodeKernelEvidenceIndex {
	return &ActionNodeKernelEvidenceIndex{
		Client: client,
	}
}

// ActionNodeKernelEvidenceIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeKernelEvidenceIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeKernelEvidenceIndexMetaGlobalInput) SetCount(value bool) *ActionNodeKernelEvidenceIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeKernelEvidenceIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeKernelEvidenceIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeKernelEvidenceIndexMetaGlobalInput) SetNo(value bool) *ActionNodeKernelEvidenceIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelEvidenceIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelEvidenceIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeKernelEvidenceIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeKernelEvidenceIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelEvidenceIndexInput is a type for action input parameters
type ActionNodeKernelEvidenceIndexInput struct {
	Limit      int64 "json:\"limit\""
	Node       int64 "json:\"node\""
	NodeActive bool  "json:\"node_active\""
	Offset     int64 "json:\"offset\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeKernelEvidenceIndexInput) SetLimit(value int64) *ActionNodeKernelEvidenceIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNodeKernelEvidenceIndexInput) SetNode(value int64) *ActionNodeKernelEvidenceIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeActive sets parameter NodeActive to value and selects it for sending
func (in *ActionNodeKernelEvidenceIndexInput) SetNodeActive(value bool) *ActionNodeKernelEvidenceIndexInput {
	in.NodeActive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeActive"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionNodeKernelEvidenceIndexInput) SetOffset(value int64) *ActionNodeKernelEvidenceIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelEvidenceIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelEvidenceIndexInput) SelectParameters(params ...string) *ActionNodeKernelEvidenceIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeKernelEvidenceIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeKernelEvidenceIndexInput) UnselectParameters(params ...string) *ActionNodeKernelEvidenceIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeKernelEvidenceIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelEvidenceIndexOutput is a type for action output parameters
type ActionNodeKernelEvidenceIndexOutput struct {
	BootId                string                "json:\"boot_id\""
	BootedAt              string                "json:\"booted_at\""
	BootedRelease         string                "json:\"booted_release\""
	BootedSystem          string                "json:\"booted_system\""
	CurrentSystem         string                "json:\"current_system\""
	EvidenceRevision      string                "json:\"evidence_revision\""
	Id                    int64                 "json:\"id\""
	KernelCommandLine     string                "json:\"kernel_command_line\""
	KernelConfigAvailable bool                  "json:\"kernel_config_available\""
	KernelConfigDigest    string                "json:\"kernel_config_digest\""
	KernelSourceRevision  string                "json:\"kernel_source_revision\""
	Node                  *ActionNodeShowOutput "json:\"node\""
	ObservedAt            string                "json:\"observed_at\""
	ReceivedAt            string                "json:\"received_at\""
	ReportSchemaVersion   int64                 "json:\"report_schema_version\""
	ReportedRelease       string                "json:\"reported_release\""
	SnapshotRevision      string                "json:\"snapshot_revision\""
}

// Type for action response, including envelope
type ActionNodeKernelEvidenceIndexResponse struct {
	Action *ActionNodeKernelEvidenceIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeKernelEvidences []*ActionNodeKernelEvidenceIndexOutput "json:\"node_kernel_evidences\""
	}

	// Action output without the namespace
	Output []*ActionNodeKernelEvidenceIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeKernelEvidenceIndex) Prepare() *ActionNodeKernelEvidenceIndexInvocation {
	return &ActionNodeKernelEvidenceIndexInvocation{
		Action: action,
		Path:   "/v7.0/node_kernel_evidences",
	}
}

// ActionNodeKernelEvidenceIndexInvocation is used to configure action for invocation
type ActionNodeKernelEvidenceIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeKernelEvidenceIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeKernelEvidenceIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeKernelEvidenceIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeKernelEvidenceIndexInvocation) NewInput() *ActionNodeKernelEvidenceIndexInput {
	inv.Input = &ActionNodeKernelEvidenceIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeKernelEvidenceIndexInvocation) SetInput(input *ActionNodeKernelEvidenceIndexInput) *ActionNodeKernelEvidenceIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeKernelEvidenceIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeKernelEvidenceIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeKernelEvidenceIndexInvocation) NewMetaInput() *ActionNodeKernelEvidenceIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeKernelEvidenceIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeKernelEvidenceIndexInvocation) SetMetaInput(input *ActionNodeKernelEvidenceIndexMetaGlobalInput) *ActionNodeKernelEvidenceIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeKernelEvidenceIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeKernelEvidenceIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeKernelEvidenceIndexInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("Node") {
			if !inv.IsParameterNil("Node") {
				if inv.Input.Node < 0 {
					verr.Add("node", "not a valid resource id")
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
func (inv *ActionNodeKernelEvidenceIndexInvocation) Call() (*ActionNodeKernelEvidenceIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeKernelEvidenceIndexInvocation) callAsQuery() (*ActionNodeKernelEvidenceIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeKernelEvidenceIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeKernelEvidences
	}
	return resp, err
}

func (inv *ActionNodeKernelEvidenceIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["node_kernel_evidence[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Node") {
			ret["node_kernel_evidence[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("NodeActive") {
			ret["node_kernel_evidence[node_active]"] = convertBoolToString(inv.Input.NodeActive)
		}
		if inv.IsParameterSelected("Offset") {
			ret["node_kernel_evidence[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
	}
}

func (inv *ActionNodeKernelEvidenceIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
