package client

import ()

// ActionNodeKernelLivepatchPatchIndex is a type for action Node_kernel_livepatch_patch#Index
type ActionNodeKernelLivepatchPatchIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeKernelLivepatchPatchIndex(client *Client) *ActionNodeKernelLivepatchPatchIndex {
	return &ActionNodeKernelLivepatchPatchIndex{
		Client: client,
	}
}

// ActionNodeKernelLivepatchPatchIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeKernelLivepatchPatchIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeKernelLivepatchPatchIndexMetaGlobalInput) SetCount(value bool) *ActionNodeKernelLivepatchPatchIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeKernelLivepatchPatchIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeKernelLivepatchPatchIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeKernelLivepatchPatchIndexMetaGlobalInput) SetNo(value bool) *ActionNodeKernelLivepatchPatchIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelLivepatchPatchIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelLivepatchPatchIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeKernelLivepatchPatchIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeKernelLivepatchPatchIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelLivepatchPatchIndexInput is a type for action input parameters
type ActionNodeKernelLivepatchPatchIndexInput struct {
	From                string "json:\"from\""
	FromId              int64  "json:\"from_id\""
	Limit               int64  "json:\"limit\""
	Name                string "json:\"name\""
	Node                int64  "json:\"node\""
	NodeActive          bool   "json:\"node_active\""
	NodeKernelEvidence  int64  "json:\"node_kernel_evidence\""
	NodeKernelLivepatch int64  "json:\"node_kernel_livepatch\""
	Source              string "json:\"source\""
	To                  string "json:\"to\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionNodeKernelLivepatchPatchIndexInput) SetFrom(value string) *ActionNodeKernelLivepatchPatchIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNodeKernelLivepatchPatchIndexInput) SetFromId(value int64) *ActionNodeKernelLivepatchPatchIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeKernelLivepatchPatchIndexInput) SetLimit(value int64) *ActionNodeKernelLivepatchPatchIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionNodeKernelLivepatchPatchIndexInput) SetName(value string) *ActionNodeKernelLivepatchPatchIndexInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNodeKernelLivepatchPatchIndexInput) SetNode(value int64) *ActionNodeKernelLivepatchPatchIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeActive sets parameter NodeActive to value and selects it for sending
func (in *ActionNodeKernelLivepatchPatchIndexInput) SetNodeActive(value bool) *ActionNodeKernelLivepatchPatchIndexInput {
	in.NodeActive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeActive"] = nil
	return in
}

// SetNodeKernelEvidence sets parameter NodeKernelEvidence to value and selects it for sending
func (in *ActionNodeKernelLivepatchPatchIndexInput) SetNodeKernelEvidence(value int64) *ActionNodeKernelLivepatchPatchIndexInput {
	in.NodeKernelEvidence = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeKernelEvidence"] = nil
	return in
}

// SetNodeKernelLivepatch sets parameter NodeKernelLivepatch to value and selects it for sending
func (in *ActionNodeKernelLivepatchPatchIndexInput) SetNodeKernelLivepatch(value int64) *ActionNodeKernelLivepatchPatchIndexInput {
	in.NodeKernelLivepatch = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeKernelLivepatch"] = nil
	return in
}

// SetSource sets parameter Source to value and selects it for sending
func (in *ActionNodeKernelLivepatchPatchIndexInput) SetSource(value string) *ActionNodeKernelLivepatchPatchIndexInput {
	in.Source = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Source"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNodeKernelLivepatchPatchIndexInput) SetTo(value string) *ActionNodeKernelLivepatchPatchIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelLivepatchPatchIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelLivepatchPatchIndexInput) SelectParameters(params ...string) *ActionNodeKernelLivepatchPatchIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeKernelLivepatchPatchIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeKernelLivepatchPatchIndexInput) UnselectParameters(params ...string) *ActionNodeKernelLivepatchPatchIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeKernelLivepatchPatchIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelLivepatchPatchIndexOutput is a type for action output parameters
type ActionNodeKernelLivepatchPatchIndexOutput struct {
	Id                  int64                                "json:\"id\""
	Name                string                               "json:\"name\""
	Node                *ActionNodeShowOutput                "json:\"node\""
	NodeKernelEvidence  *ActionNodeKernelEvidenceShowOutput  "json:\"node_kernel_evidence\""
	NodeKernelLivepatch *ActionNodeKernelLivepatchShowOutput "json:\"node_kernel_livepatch\""
	ObservedAt          string                               "json:\"observed_at\""
	Source              string                               "json:\"source\""
	SourceRevision      string                               "json:\"source_revision\""
	Version             string                               "json:\"version\""
}

// Type for action response, including envelope
type ActionNodeKernelLivepatchPatchIndexResponse struct {
	Action *ActionNodeKernelLivepatchPatchIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeKernelLivepatchPatches []*ActionNodeKernelLivepatchPatchIndexOutput "json:\"node_kernel_livepatch_patches\""
	}

	// Action output without the namespace
	Output []*ActionNodeKernelLivepatchPatchIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeKernelLivepatchPatchIndex) Prepare() *ActionNodeKernelLivepatchPatchIndexInvocation {
	return &ActionNodeKernelLivepatchPatchIndexInvocation{
		Action: action,
		Path:   "/v7.0/node_kernel_livepatch_patches",
	}
}

// ActionNodeKernelLivepatchPatchIndexInvocation is used to configure action for invocation
type ActionNodeKernelLivepatchPatchIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeKernelLivepatchPatchIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeKernelLivepatchPatchIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeKernelLivepatchPatchIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeKernelLivepatchPatchIndexInvocation) NewInput() *ActionNodeKernelLivepatchPatchIndexInput {
	inv.Input = &ActionNodeKernelLivepatchPatchIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeKernelLivepatchPatchIndexInvocation) SetInput(input *ActionNodeKernelLivepatchPatchIndexInput) *ActionNodeKernelLivepatchPatchIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeKernelLivepatchPatchIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeKernelLivepatchPatchIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeKernelLivepatchPatchIndexInvocation) NewMetaInput() *ActionNodeKernelLivepatchPatchIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeKernelLivepatchPatchIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeKernelLivepatchPatchIndexInvocation) SetMetaInput(input *ActionNodeKernelLivepatchPatchIndexMetaGlobalInput) *ActionNodeKernelLivepatchPatchIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeKernelLivepatchPatchIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeKernelLivepatchPatchIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeKernelLivepatchPatchIndexInvocation) validate() error {
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
		if inv.IsParameterSelected("NodeKernelLivepatch") {
			if !inv.IsParameterNil("NodeKernelLivepatch") {
				if inv.Input.NodeKernelLivepatch < 0 {
					verr.Add("node_kernel_livepatch", "not a valid resource id")
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
func (inv *ActionNodeKernelLivepatchPatchIndexInvocation) Call() (*ActionNodeKernelLivepatchPatchIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeKernelLivepatchPatchIndexInvocation) callAsQuery() (*ActionNodeKernelLivepatchPatchIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeKernelLivepatchPatchIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeKernelLivepatchPatches
	}
	return resp, err
}

func (inv *ActionNodeKernelLivepatchPatchIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("From") {
			ret["node_kernel_livepatch_patch[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("FromId") {
			ret["node_kernel_livepatch_patch[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["node_kernel_livepatch_patch[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Name") {
			ret["node_kernel_livepatch_patch[name]"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Node") {
			ret["node_kernel_livepatch_patch[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("NodeActive") {
			ret["node_kernel_livepatch_patch[node_active]"] = convertBoolToString(inv.Input.NodeActive)
		}
		if inv.IsParameterSelected("NodeKernelEvidence") {
			ret["node_kernel_livepatch_patch[node_kernel_evidence]"] = convertInt64ToString(inv.Input.NodeKernelEvidence)
		}
		if inv.IsParameterSelected("NodeKernelLivepatch") {
			ret["node_kernel_livepatch_patch[node_kernel_livepatch]"] = convertInt64ToString(inv.Input.NodeKernelLivepatch)
		}
		if inv.IsParameterSelected("Source") {
			ret["node_kernel_livepatch_patch[source]"] = inv.Input.Source
		}
		if inv.IsParameterSelected("To") {
			ret["node_kernel_livepatch_patch[to]"] = inv.Input.To
		}
	}
}

func (inv *ActionNodeKernelLivepatchPatchIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
