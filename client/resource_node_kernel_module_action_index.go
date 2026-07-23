package client

import ()

// ActionNodeKernelModuleIndex is a type for action Node_kernel_module#Index
type ActionNodeKernelModuleIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeKernelModuleIndex(client *Client) *ActionNodeKernelModuleIndex {
	return &ActionNodeKernelModuleIndex{
		Client: client,
	}
}

// ActionNodeKernelModuleIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeKernelModuleIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeKernelModuleIndexMetaGlobalInput) SetCount(value bool) *ActionNodeKernelModuleIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeKernelModuleIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeKernelModuleIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeKernelModuleIndexMetaGlobalInput) SetNo(value bool) *ActionNodeKernelModuleIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelModuleIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelModuleIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeKernelModuleIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeKernelModuleIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelModuleIndexInput is a type for action input parameters
type ActionNodeKernelModuleIndexInput struct {
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
func (in *ActionNodeKernelModuleIndexInput) SetFrom(value string) *ActionNodeKernelModuleIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNodeKernelModuleIndexInput) SetFromId(value int64) *ActionNodeKernelModuleIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeKernelModuleIndexInput) SetLimit(value int64) *ActionNodeKernelModuleIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionNodeKernelModuleIndexInput) SetName(value string) *ActionNodeKernelModuleIndexInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNodeKernelModuleIndexInput) SetNode(value int64) *ActionNodeKernelModuleIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeActive sets parameter NodeActive to value and selects it for sending
func (in *ActionNodeKernelModuleIndexInput) SetNodeActive(value bool) *ActionNodeKernelModuleIndexInput {
	in.NodeActive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeActive"] = nil
	return in
}

// SetNodeKernelEvidence sets parameter NodeKernelEvidence to value and selects it for sending
func (in *ActionNodeKernelModuleIndexInput) SetNodeKernelEvidence(value int64) *ActionNodeKernelModuleIndexInput {
	in.NodeKernelEvidence = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeKernelEvidence"] = nil
	return in
}

// SetSource sets parameter Source to value and selects it for sending
func (in *ActionNodeKernelModuleIndexInput) SetSource(value string) *ActionNodeKernelModuleIndexInput {
	in.Source = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Source"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNodeKernelModuleIndexInput) SetTo(value string) *ActionNodeKernelModuleIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelModuleIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelModuleIndexInput) SelectParameters(params ...string) *ActionNodeKernelModuleIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeKernelModuleIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeKernelModuleIndexInput) UnselectParameters(params ...string) *ActionNodeKernelModuleIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeKernelModuleIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelModuleIndexOutput is a type for action output parameters
type ActionNodeKernelModuleIndexOutput struct {
	Id                 int64                               "json:\"id\""
	Name               string                              "json:\"name\""
	Node               *ActionNodeShowOutput               "json:\"node\""
	NodeKernelEvidence *ActionNodeKernelEvidenceShowOutput "json:\"node_kernel_evidence\""
	ObservedAt         string                              "json:\"observed_at\""
	Source             string                              "json:\"source\""
	SourceRevision     string                              "json:\"source_revision\""
}

// Type for action response, including envelope
type ActionNodeKernelModuleIndexResponse struct {
	Action *ActionNodeKernelModuleIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeKernelModules []*ActionNodeKernelModuleIndexOutput "json:\"node_kernel_modules\""
	}

	// Action output without the namespace
	Output []*ActionNodeKernelModuleIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeKernelModuleIndex) Prepare() *ActionNodeKernelModuleIndexInvocation {
	return &ActionNodeKernelModuleIndexInvocation{
		Action: action,
		Path:   "/v7.0/node_kernel_modules",
	}
}

// ActionNodeKernelModuleIndexInvocation is used to configure action for invocation
type ActionNodeKernelModuleIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeKernelModuleIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeKernelModuleIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeKernelModuleIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeKernelModuleIndexInvocation) NewInput() *ActionNodeKernelModuleIndexInput {
	inv.Input = &ActionNodeKernelModuleIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeKernelModuleIndexInvocation) SetInput(input *ActionNodeKernelModuleIndexInput) *ActionNodeKernelModuleIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeKernelModuleIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeKernelModuleIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeKernelModuleIndexInvocation) NewMetaInput() *ActionNodeKernelModuleIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeKernelModuleIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeKernelModuleIndexInvocation) SetMetaInput(input *ActionNodeKernelModuleIndexMetaGlobalInput) *ActionNodeKernelModuleIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeKernelModuleIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeKernelModuleIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeKernelModuleIndexInvocation) validate() error {
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
func (inv *ActionNodeKernelModuleIndexInvocation) Call() (*ActionNodeKernelModuleIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeKernelModuleIndexInvocation) callAsQuery() (*ActionNodeKernelModuleIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeKernelModuleIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeKernelModules
	}
	return resp, err
}

func (inv *ActionNodeKernelModuleIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("From") {
			ret["node_kernel_module[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("FromId") {
			ret["node_kernel_module[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["node_kernel_module[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Name") {
			ret["node_kernel_module[name]"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Node") {
			ret["node_kernel_module[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("NodeActive") {
			ret["node_kernel_module[node_active]"] = convertBoolToString(inv.Input.NodeActive)
		}
		if inv.IsParameterSelected("NodeKernelEvidence") {
			ret["node_kernel_module[node_kernel_evidence]"] = convertInt64ToString(inv.Input.NodeKernelEvidence)
		}
		if inv.IsParameterSelected("Source") {
			ret["node_kernel_module[source]"] = inv.Input.Source
		}
		if inv.IsParameterSelected("To") {
			ret["node_kernel_module[to]"] = inv.Input.To
		}
	}
}

func (inv *ActionNodeKernelModuleIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
