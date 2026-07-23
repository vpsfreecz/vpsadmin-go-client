package client

import ()

// ActionNodeSoftwareDeploymentIndex is a type for action Node_software_deployment#Index
type ActionNodeSoftwareDeploymentIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeSoftwareDeploymentIndex(client *Client) *ActionNodeSoftwareDeploymentIndex {
	return &ActionNodeSoftwareDeploymentIndex{
		Client: client,
	}
}

// ActionNodeSoftwareDeploymentIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeSoftwareDeploymentIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeSoftwareDeploymentIndexMetaGlobalInput) SetCount(value bool) *ActionNodeSoftwareDeploymentIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeSoftwareDeploymentIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeSoftwareDeploymentIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeSoftwareDeploymentIndexMetaGlobalInput) SetNo(value bool) *ActionNodeSoftwareDeploymentIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeSoftwareDeploymentIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeSoftwareDeploymentIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeSoftwareDeploymentIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeSoftwareDeploymentIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeSoftwareDeploymentIndexInput is a type for action input parameters
type ActionNodeSoftwareDeploymentIndexInput struct {
	Component  string "json:\"component\""
	From       string "json:\"from\""
	Generation string "json:\"generation\""
	Limit      int64  "json:\"limit\""
	Node       int64  "json:\"node\""
	NodeActive bool   "json:\"node_active\""
	Offset     int64  "json:\"offset\""
	To         string "json:\"to\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetComponent sets parameter Component to value and selects it for sending
func (in *ActionNodeSoftwareDeploymentIndexInput) SetComponent(value string) *ActionNodeSoftwareDeploymentIndexInput {
	in.Component = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Component"] = nil
	return in
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionNodeSoftwareDeploymentIndexInput) SetFrom(value string) *ActionNodeSoftwareDeploymentIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetGeneration sets parameter Generation to value and selects it for sending
func (in *ActionNodeSoftwareDeploymentIndexInput) SetGeneration(value string) *ActionNodeSoftwareDeploymentIndexInput {
	in.Generation = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Generation"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeSoftwareDeploymentIndexInput) SetLimit(value int64) *ActionNodeSoftwareDeploymentIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNodeSoftwareDeploymentIndexInput) SetNode(value int64) *ActionNodeSoftwareDeploymentIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeActive sets parameter NodeActive to value and selects it for sending
func (in *ActionNodeSoftwareDeploymentIndexInput) SetNodeActive(value bool) *ActionNodeSoftwareDeploymentIndexInput {
	in.NodeActive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeActive"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionNodeSoftwareDeploymentIndexInput) SetOffset(value int64) *ActionNodeSoftwareDeploymentIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNodeSoftwareDeploymentIndexInput) SetTo(value string) *ActionNodeSoftwareDeploymentIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeSoftwareDeploymentIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeSoftwareDeploymentIndexInput) SelectParameters(params ...string) *ActionNodeSoftwareDeploymentIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeSoftwareDeploymentIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeSoftwareDeploymentIndexInput) UnselectParameters(params ...string) *ActionNodeSoftwareDeploymentIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeSoftwareDeploymentIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeSoftwareDeploymentIndexOutput is a type for action output parameters
type ActionNodeSoftwareDeploymentIndexOutput struct {
	BootedSystem       string                              "json:\"booted_system\""
	ChangeCount        int64                               "json:\"change_count\""
	CurrentSystem      string                              "json:\"current_system\""
	EffectiveAt        string                              "json:\"effective_at\""
	EventType          string                              "json:\"event_type\""
	Id                 int64                               "json:\"id\""
	Node               *ActionNodeShowOutput               "json:\"node\""
	NodeKernelEvidence *ActionNodeKernelEvidenceShowOutput "json:\"node_kernel_evidence\""
	ObservedAfter      string                              "json:\"observed_after\""
	ObservedBefore     string                              "json:\"observed_before\""
	SourceRevision     string                              "json:\"source_revision\""
}

// Type for action response, including envelope
type ActionNodeSoftwareDeploymentIndexResponse struct {
	Action *ActionNodeSoftwareDeploymentIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeSoftwareDeployments []*ActionNodeSoftwareDeploymentIndexOutput "json:\"node_software_deployments\""
	}

	// Action output without the namespace
	Output []*ActionNodeSoftwareDeploymentIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeSoftwareDeploymentIndex) Prepare() *ActionNodeSoftwareDeploymentIndexInvocation {
	return &ActionNodeSoftwareDeploymentIndexInvocation{
		Action: action,
		Path:   "/v7.0/node_software_deployments",
	}
}

// ActionNodeSoftwareDeploymentIndexInvocation is used to configure action for invocation
type ActionNodeSoftwareDeploymentIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeSoftwareDeploymentIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeSoftwareDeploymentIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeSoftwareDeploymentIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeSoftwareDeploymentIndexInvocation) NewInput() *ActionNodeSoftwareDeploymentIndexInput {
	inv.Input = &ActionNodeSoftwareDeploymentIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeSoftwareDeploymentIndexInvocation) SetInput(input *ActionNodeSoftwareDeploymentIndexInput) *ActionNodeSoftwareDeploymentIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeSoftwareDeploymentIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeSoftwareDeploymentIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeSoftwareDeploymentIndexInvocation) NewMetaInput() *ActionNodeSoftwareDeploymentIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeSoftwareDeploymentIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeSoftwareDeploymentIndexInvocation) SetMetaInput(input *ActionNodeSoftwareDeploymentIndexMetaGlobalInput) *ActionNodeSoftwareDeploymentIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeSoftwareDeploymentIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeSoftwareDeploymentIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeSoftwareDeploymentIndexInvocation) validate() error {
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
func (inv *ActionNodeSoftwareDeploymentIndexInvocation) Call() (*ActionNodeSoftwareDeploymentIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeSoftwareDeploymentIndexInvocation) callAsQuery() (*ActionNodeSoftwareDeploymentIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeSoftwareDeploymentIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeSoftwareDeployments
	}
	return resp, err
}

func (inv *ActionNodeSoftwareDeploymentIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Component") {
			ret["node_software_deployment[component]"] = inv.Input.Component
		}
		if inv.IsParameterSelected("From") {
			ret["node_software_deployment[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("Generation") {
			ret["node_software_deployment[generation]"] = inv.Input.Generation
		}
		if inv.IsParameterSelected("Limit") {
			ret["node_software_deployment[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Node") {
			ret["node_software_deployment[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("NodeActive") {
			ret["node_software_deployment[node_active]"] = convertBoolToString(inv.Input.NodeActive)
		}
		if inv.IsParameterSelected("Offset") {
			ret["node_software_deployment[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("To") {
			ret["node_software_deployment[to]"] = inv.Input.To
		}
	}
}

func (inv *ActionNodeSoftwareDeploymentIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
