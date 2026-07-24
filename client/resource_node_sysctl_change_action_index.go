package client

import ()

// ActionNodeSysctlChangeIndex is a type for action Node_sysctl_change#Index
type ActionNodeSysctlChangeIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeSysctlChangeIndex(client *Client) *ActionNodeSysctlChangeIndex {
	return &ActionNodeSysctlChangeIndex{
		Client: client,
	}
}

// ActionNodeSysctlChangeIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeSysctlChangeIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeSysctlChangeIndexMetaGlobalInput) SetCount(value bool) *ActionNodeSysctlChangeIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeSysctlChangeIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeSysctlChangeIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeSysctlChangeIndexMetaGlobalInput) SetNo(value bool) *ActionNodeSysctlChangeIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeSysctlChangeIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeSysctlChangeIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeSysctlChangeIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeSysctlChangeIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeSysctlChangeIndexInput is a type for action input parameters
type ActionNodeSysctlChangeIndexInput struct {
	From            string "json:\"from\""
	Limit           int64  "json:\"limit\""
	Name            string "json:\"name\""
	Node            int64  "json:\"node\""
	NodeActive      bool   "json:\"node_active\""
	NodeKernelEvent int64  "json:\"node_kernel_event\""
	Offset          int64  "json:\"offset\""
	To              string "json:\"to\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionNodeSysctlChangeIndexInput) SetFrom(value string) *ActionNodeSysctlChangeIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeSysctlChangeIndexInput) SetLimit(value int64) *ActionNodeSysctlChangeIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionNodeSysctlChangeIndexInput) SetName(value string) *ActionNodeSysctlChangeIndexInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNodeSysctlChangeIndexInput) SetNode(value int64) *ActionNodeSysctlChangeIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeActive sets parameter NodeActive to value and selects it for sending
func (in *ActionNodeSysctlChangeIndexInput) SetNodeActive(value bool) *ActionNodeSysctlChangeIndexInput {
	in.NodeActive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeActive"] = nil
	return in
}

// SetNodeKernelEvent sets parameter NodeKernelEvent to value and selects it for sending
func (in *ActionNodeSysctlChangeIndexInput) SetNodeKernelEvent(value int64) *ActionNodeSysctlChangeIndexInput {
	in.NodeKernelEvent = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeKernelEvent"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionNodeSysctlChangeIndexInput) SetOffset(value int64) *ActionNodeSysctlChangeIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNodeSysctlChangeIndexInput) SetTo(value string) *ActionNodeSysctlChangeIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeSysctlChangeIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeSysctlChangeIndexInput) SelectParameters(params ...string) *ActionNodeSysctlChangeIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeSysctlChangeIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeSysctlChangeIndexInput) UnselectParameters(params ...string) *ActionNodeSysctlChangeIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeSysctlChangeIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeSysctlChangeIndexOutput is a type for action output parameters
type ActionNodeSysctlChangeIndexOutput struct {
	AfterAvailable        bool                                "json:\"after_available\""
	AfterConfiguredValue  string                              "json:\"after_configured_value\""
	AfterEffectiveValue   string                              "json:\"after_effective_value\""
	BeforeAvailable       bool                                "json:\"before_available\""
	BeforeConfiguredValue string                              "json:\"before_configured_value\""
	BeforeEffectiveValue  string                              "json:\"before_effective_value\""
	Id                    int64                               "json:\"id\""
	Name                  string                              "json:\"name\""
	Node                  *ActionNodeShowOutput               "json:\"node\""
	NodeKernelEvent       *ActionNodeKernelEventShowOutput    "json:\"node_kernel_event\""
	NodeKernelEvidence    *ActionNodeKernelEvidenceShowOutput "json:\"node_kernel_evidence\""
	ObservedAfter         string                              "json:\"observed_after\""
	ObservedBefore        string                              "json:\"observed_before\""
	SourceRevision        string                              "json:\"source_revision\""
}

// Type for action response, including envelope
type ActionNodeSysctlChangeIndexResponse struct {
	Action *ActionNodeSysctlChangeIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeSysctlChanges []*ActionNodeSysctlChangeIndexOutput "json:\"node_sysctl_changes\""
	}

	// Action output without the namespace
	Output []*ActionNodeSysctlChangeIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeSysctlChangeIndex) Prepare() *ActionNodeSysctlChangeIndexInvocation {
	return &ActionNodeSysctlChangeIndexInvocation{
		Action: action,
		Path:   "/v7.0/node_sysctl_changes",
	}
}

// ActionNodeSysctlChangeIndexInvocation is used to configure action for invocation
type ActionNodeSysctlChangeIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeSysctlChangeIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeSysctlChangeIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeSysctlChangeIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeSysctlChangeIndexInvocation) NewInput() *ActionNodeSysctlChangeIndexInput {
	inv.Input = &ActionNodeSysctlChangeIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeSysctlChangeIndexInvocation) SetInput(input *ActionNodeSysctlChangeIndexInput) *ActionNodeSysctlChangeIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeSysctlChangeIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeSysctlChangeIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeSysctlChangeIndexInvocation) NewMetaInput() *ActionNodeSysctlChangeIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeSysctlChangeIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeSysctlChangeIndexInvocation) SetMetaInput(input *ActionNodeSysctlChangeIndexMetaGlobalInput) *ActionNodeSysctlChangeIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeSysctlChangeIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeSysctlChangeIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeSysctlChangeIndexInvocation) validate() error {
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
		if inv.IsParameterSelected("NodeKernelEvent") {
			if !inv.IsParameterNil("NodeKernelEvent") {
				if inv.Input.NodeKernelEvent < 0 {
					verr.Add("node_kernel_event", "not a valid resource id")
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
func (inv *ActionNodeSysctlChangeIndexInvocation) Call() (*ActionNodeSysctlChangeIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeSysctlChangeIndexInvocation) callAsQuery() (*ActionNodeSysctlChangeIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionNodeSysctlChangeIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeSysctlChanges
	}
	return resp, err
}

func (inv *ActionNodeSysctlChangeIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("From") {
			ret["node_sysctl_change[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("Limit") {
			ret["node_sysctl_change[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Name") {
			ret["node_sysctl_change[name]"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Node") {
			ret["node_sysctl_change[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("NodeActive") {
			ret["node_sysctl_change[node_active]"] = convertBoolToString(inv.Input.NodeActive)
		}
		if inv.IsParameterSelected("NodeKernelEvent") {
			ret["node_sysctl_change[node_kernel_event]"] = convertInt64ToString(inv.Input.NodeKernelEvent)
		}
		if inv.IsParameterSelected("Offset") {
			ret["node_sysctl_change[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("To") {
			ret["node_sysctl_change[to]"] = inv.Input.To
		}
	}

	return nil
}

func (inv *ActionNodeSysctlChangeIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
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
