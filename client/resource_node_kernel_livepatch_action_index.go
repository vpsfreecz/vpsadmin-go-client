package client

import ()

// ActionNodeKernelLivepatchIndex is a type for action Node_kernel_livepatch#Index
type ActionNodeKernelLivepatchIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeKernelLivepatchIndex(client *Client) *ActionNodeKernelLivepatchIndex {
	return &ActionNodeKernelLivepatchIndex{
		Client: client,
	}
}

// ActionNodeKernelLivepatchIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeKernelLivepatchIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeKernelLivepatchIndexMetaGlobalInput) SetCount(value bool) *ActionNodeKernelLivepatchIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeKernelLivepatchIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeKernelLivepatchIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeKernelLivepatchIndexMetaGlobalInput) SetNo(value bool) *ActionNodeKernelLivepatchIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelLivepatchIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelLivepatchIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeKernelLivepatchIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeKernelLivepatchIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelLivepatchIndexInput is a type for action input parameters
type ActionNodeKernelLivepatchIndexInput struct {
	From               string "json:\"from\""
	FromId             int64  "json:\"from_id\""
	Limit              int64  "json:\"limit\""
	LivepatchId        string "json:\"livepatch_id\""
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
func (in *ActionNodeKernelLivepatchIndexInput) SetFrom(value string) *ActionNodeKernelLivepatchIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNodeKernelLivepatchIndexInput) SetFromId(value int64) *ActionNodeKernelLivepatchIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeKernelLivepatchIndexInput) SetLimit(value int64) *ActionNodeKernelLivepatchIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetLivepatchId sets parameter LivepatchId to value and selects it for sending
func (in *ActionNodeKernelLivepatchIndexInput) SetLivepatchId(value string) *ActionNodeKernelLivepatchIndexInput {
	in.LivepatchId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["LivepatchId"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNodeKernelLivepatchIndexInput) SetNode(value int64) *ActionNodeKernelLivepatchIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeActive sets parameter NodeActive to value and selects it for sending
func (in *ActionNodeKernelLivepatchIndexInput) SetNodeActive(value bool) *ActionNodeKernelLivepatchIndexInput {
	in.NodeActive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeActive"] = nil
	return in
}

// SetNodeKernelEvidence sets parameter NodeKernelEvidence to value and selects it for sending
func (in *ActionNodeKernelLivepatchIndexInput) SetNodeKernelEvidence(value int64) *ActionNodeKernelLivepatchIndexInput {
	in.NodeKernelEvidence = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeKernelEvidence"] = nil
	return in
}

// SetSource sets parameter Source to value and selects it for sending
func (in *ActionNodeKernelLivepatchIndexInput) SetSource(value string) *ActionNodeKernelLivepatchIndexInput {
	in.Source = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Source"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNodeKernelLivepatchIndexInput) SetTo(value string) *ActionNodeKernelLivepatchIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelLivepatchIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelLivepatchIndexInput) SelectParameters(params ...string) *ActionNodeKernelLivepatchIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeKernelLivepatchIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeKernelLivepatchIndexInput) UnselectParameters(params ...string) *ActionNodeKernelLivepatchIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeKernelLivepatchIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelLivepatchIndexOutput is a type for action output parameters
type ActionNodeKernelLivepatchIndexOutput struct {
	AppliedAt          string                              "json:\"applied_at\""
	Enabled            bool                                "json:\"enabled\""
	Id                 int64                               "json:\"id\""
	KernelVersion      string                              "json:\"kernel_version\""
	LivepatchId        string                              "json:\"livepatch_id\""
	Loaded             bool                                "json:\"loaded\""
	Node               *ActionNodeShowOutput               "json:\"node\""
	NodeKernelEvidence *ActionNodeKernelEvidenceShowOutput "json:\"node_kernel_evidence\""
	ObservedAt         string                              "json:\"observed_at\""
	PatchVersion       string                              "json:\"patch_version\""
	Source             string                              "json:\"source\""
	SourceRevision     string                              "json:\"source_revision\""
	Transition         bool                                "json:\"transition\""
	VerifiedAt         string                              "json:\"verified_at\""
}

// Type for action response, including envelope
type ActionNodeKernelLivepatchIndexResponse struct {
	Action *ActionNodeKernelLivepatchIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeKernelLivepatches []*ActionNodeKernelLivepatchIndexOutput "json:\"node_kernel_livepatches\""
	}

	// Action output without the namespace
	Output []*ActionNodeKernelLivepatchIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeKernelLivepatchIndex) Prepare() *ActionNodeKernelLivepatchIndexInvocation {
	return &ActionNodeKernelLivepatchIndexInvocation{
		Action: action,
		Path:   "/v7.0/node_kernel_livepatches",
	}
}

// ActionNodeKernelLivepatchIndexInvocation is used to configure action for invocation
type ActionNodeKernelLivepatchIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeKernelLivepatchIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeKernelLivepatchIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeKernelLivepatchIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeKernelLivepatchIndexInvocation) NewInput() *ActionNodeKernelLivepatchIndexInput {
	inv.Input = &ActionNodeKernelLivepatchIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeKernelLivepatchIndexInvocation) SetInput(input *ActionNodeKernelLivepatchIndexInput) *ActionNodeKernelLivepatchIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeKernelLivepatchIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeKernelLivepatchIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeKernelLivepatchIndexInvocation) NewMetaInput() *ActionNodeKernelLivepatchIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeKernelLivepatchIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeKernelLivepatchIndexInvocation) SetMetaInput(input *ActionNodeKernelLivepatchIndexMetaGlobalInput) *ActionNodeKernelLivepatchIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeKernelLivepatchIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeKernelLivepatchIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeKernelLivepatchIndexInvocation) validate() error {
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
func (inv *ActionNodeKernelLivepatchIndexInvocation) Call() (*ActionNodeKernelLivepatchIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeKernelLivepatchIndexInvocation) callAsQuery() (*ActionNodeKernelLivepatchIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeKernelLivepatchIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeKernelLivepatches
	}
	return resp, err
}

func (inv *ActionNodeKernelLivepatchIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("From") {
			ret["node_kernel_livepatch[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("FromId") {
			ret["node_kernel_livepatch[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["node_kernel_livepatch[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("LivepatchId") {
			ret["node_kernel_livepatch[livepatch_id]"] = inv.Input.LivepatchId
		}
		if inv.IsParameterSelected("Node") {
			ret["node_kernel_livepatch[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("NodeActive") {
			ret["node_kernel_livepatch[node_active]"] = convertBoolToString(inv.Input.NodeActive)
		}
		if inv.IsParameterSelected("NodeKernelEvidence") {
			ret["node_kernel_livepatch[node_kernel_evidence]"] = convertInt64ToString(inv.Input.NodeKernelEvidence)
		}
		if inv.IsParameterSelected("Source") {
			ret["node_kernel_livepatch[source]"] = inv.Input.Source
		}
		if inv.IsParameterSelected("To") {
			ret["node_kernel_livepatch[to]"] = inv.Input.To
		}
	}
}

func (inv *ActionNodeKernelLivepatchIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
