package client

import ()

// ActionNodeKernelEvidenceErrorIndex is a type for action Node_kernel_evidence_error#Index
type ActionNodeKernelEvidenceErrorIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeKernelEvidenceErrorIndex(client *Client) *ActionNodeKernelEvidenceErrorIndex {
	return &ActionNodeKernelEvidenceErrorIndex{
		Client: client,
	}
}

// ActionNodeKernelEvidenceErrorIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeKernelEvidenceErrorIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeKernelEvidenceErrorIndexMetaGlobalInput) SetCount(value bool) *ActionNodeKernelEvidenceErrorIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeKernelEvidenceErrorIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeKernelEvidenceErrorIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeKernelEvidenceErrorIndexMetaGlobalInput) SetNo(value bool) *ActionNodeKernelEvidenceErrorIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelEvidenceErrorIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelEvidenceErrorIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeKernelEvidenceErrorIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeKernelEvidenceErrorIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelEvidenceErrorIndexInput is a type for action input parameters
type ActionNodeKernelEvidenceErrorIndexInput struct {
	Component          string "json:\"component\""
	From               string "json:\"from\""
	FromId             int64  "json:\"from_id\""
	Limit              int64  "json:\"limit\""
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

// SetComponent sets parameter Component to value and selects it for sending
func (in *ActionNodeKernelEvidenceErrorIndexInput) SetComponent(value string) *ActionNodeKernelEvidenceErrorIndexInput {
	in.Component = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Component"] = nil
	return in
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionNodeKernelEvidenceErrorIndexInput) SetFrom(value string) *ActionNodeKernelEvidenceErrorIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNodeKernelEvidenceErrorIndexInput) SetFromId(value int64) *ActionNodeKernelEvidenceErrorIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeKernelEvidenceErrorIndexInput) SetLimit(value int64) *ActionNodeKernelEvidenceErrorIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNodeKernelEvidenceErrorIndexInput) SetNode(value int64) *ActionNodeKernelEvidenceErrorIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeActive sets parameter NodeActive to value and selects it for sending
func (in *ActionNodeKernelEvidenceErrorIndexInput) SetNodeActive(value bool) *ActionNodeKernelEvidenceErrorIndexInput {
	in.NodeActive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeActive"] = nil
	return in
}

// SetNodeKernelEvidence sets parameter NodeKernelEvidence to value and selects it for sending
func (in *ActionNodeKernelEvidenceErrorIndexInput) SetNodeKernelEvidence(value int64) *ActionNodeKernelEvidenceErrorIndexInput {
	in.NodeKernelEvidence = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeKernelEvidence"] = nil
	return in
}

// SetSource sets parameter Source to value and selects it for sending
func (in *ActionNodeKernelEvidenceErrorIndexInput) SetSource(value string) *ActionNodeKernelEvidenceErrorIndexInput {
	in.Source = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Source"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNodeKernelEvidenceErrorIndexInput) SetTo(value string) *ActionNodeKernelEvidenceErrorIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelEvidenceErrorIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelEvidenceErrorIndexInput) SelectParameters(params ...string) *ActionNodeKernelEvidenceErrorIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeKernelEvidenceErrorIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeKernelEvidenceErrorIndexInput) UnselectParameters(params ...string) *ActionNodeKernelEvidenceErrorIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeKernelEvidenceErrorIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelEvidenceErrorIndexOutput is a type for action output parameters
type ActionNodeKernelEvidenceErrorIndexOutput struct {
	Component          string                              "json:\"component\""
	Id                 int64                               "json:\"id\""
	Node               *ActionNodeShowOutput               "json:\"node\""
	NodeKernelEvidence *ActionNodeKernelEvidenceShowOutput "json:\"node_kernel_evidence\""
	ObservedAt         string                              "json:\"observed_at\""
	Reason             string                              "json:\"reason\""
	Source             string                              "json:\"source\""
	SourceRevision     string                              "json:\"source_revision\""
}

// Type for action response, including envelope
type ActionNodeKernelEvidenceErrorIndexResponse struct {
	Action *ActionNodeKernelEvidenceErrorIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeKernelEvidenceErrors []*ActionNodeKernelEvidenceErrorIndexOutput "json:\"node_kernel_evidence_errors\""
	}

	// Action output without the namespace
	Output []*ActionNodeKernelEvidenceErrorIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeKernelEvidenceErrorIndex) Prepare() *ActionNodeKernelEvidenceErrorIndexInvocation {
	return &ActionNodeKernelEvidenceErrorIndexInvocation{
		Action: action,
		Path:   "/v7.0/node_kernel_evidence_errors",
	}
}

// ActionNodeKernelEvidenceErrorIndexInvocation is used to configure action for invocation
type ActionNodeKernelEvidenceErrorIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeKernelEvidenceErrorIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeKernelEvidenceErrorIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeKernelEvidenceErrorIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeKernelEvidenceErrorIndexInvocation) NewInput() *ActionNodeKernelEvidenceErrorIndexInput {
	inv.Input = &ActionNodeKernelEvidenceErrorIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeKernelEvidenceErrorIndexInvocation) SetInput(input *ActionNodeKernelEvidenceErrorIndexInput) *ActionNodeKernelEvidenceErrorIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeKernelEvidenceErrorIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeKernelEvidenceErrorIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeKernelEvidenceErrorIndexInvocation) NewMetaInput() *ActionNodeKernelEvidenceErrorIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeKernelEvidenceErrorIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeKernelEvidenceErrorIndexInvocation) SetMetaInput(input *ActionNodeKernelEvidenceErrorIndexMetaGlobalInput) *ActionNodeKernelEvidenceErrorIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeKernelEvidenceErrorIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeKernelEvidenceErrorIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeKernelEvidenceErrorIndexInvocation) validate() error {
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
func (inv *ActionNodeKernelEvidenceErrorIndexInvocation) Call() (*ActionNodeKernelEvidenceErrorIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeKernelEvidenceErrorIndexInvocation) callAsQuery() (*ActionNodeKernelEvidenceErrorIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeKernelEvidenceErrorIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeKernelEvidenceErrors
	}
	return resp, err
}

func (inv *ActionNodeKernelEvidenceErrorIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Component") {
			ret["node_kernel_evidence_error[component]"] = inv.Input.Component
		}
		if inv.IsParameterSelected("From") {
			ret["node_kernel_evidence_error[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("FromId") {
			ret["node_kernel_evidence_error[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["node_kernel_evidence_error[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Node") {
			ret["node_kernel_evidence_error[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("NodeActive") {
			ret["node_kernel_evidence_error[node_active]"] = convertBoolToString(inv.Input.NodeActive)
		}
		if inv.IsParameterSelected("NodeKernelEvidence") {
			ret["node_kernel_evidence_error[node_kernel_evidence]"] = convertInt64ToString(inv.Input.NodeKernelEvidence)
		}
		if inv.IsParameterSelected("Source") {
			ret["node_kernel_evidence_error[source]"] = inv.Input.Source
		}
		if inv.IsParameterSelected("To") {
			ret["node_kernel_evidence_error[to]"] = inv.Input.To
		}
	}
}

func (inv *ActionNodeKernelEvidenceErrorIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
