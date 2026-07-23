package client

import ()

// ActionNodeSoftwareChangeIndex is a type for action Node_software_change#Index
type ActionNodeSoftwareChangeIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeSoftwareChangeIndex(client *Client) *ActionNodeSoftwareChangeIndex {
	return &ActionNodeSoftwareChangeIndex{
		Client: client,
	}
}

// ActionNodeSoftwareChangeIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeSoftwareChangeIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeSoftwareChangeIndexMetaGlobalInput) SetCount(value bool) *ActionNodeSoftwareChangeIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeSoftwareChangeIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeSoftwareChangeIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeSoftwareChangeIndexMetaGlobalInput) SetNo(value bool) *ActionNodeSoftwareChangeIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeSoftwareChangeIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeSoftwareChangeIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeSoftwareChangeIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeSoftwareChangeIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeSoftwareChangeIndexInput is a type for action input parameters
type ActionNodeSoftwareChangeIndexInput struct {
	Component       string "json:\"component\""
	From            string "json:\"from\""
	Generation      string "json:\"generation\""
	Limit           int64  "json:\"limit\""
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

// SetComponent sets parameter Component to value and selects it for sending
func (in *ActionNodeSoftwareChangeIndexInput) SetComponent(value string) *ActionNodeSoftwareChangeIndexInput {
	in.Component = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Component"] = nil
	return in
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionNodeSoftwareChangeIndexInput) SetFrom(value string) *ActionNodeSoftwareChangeIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetGeneration sets parameter Generation to value and selects it for sending
func (in *ActionNodeSoftwareChangeIndexInput) SetGeneration(value string) *ActionNodeSoftwareChangeIndexInput {
	in.Generation = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Generation"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeSoftwareChangeIndexInput) SetLimit(value int64) *ActionNodeSoftwareChangeIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNodeSoftwareChangeIndexInput) SetNode(value int64) *ActionNodeSoftwareChangeIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeActive sets parameter NodeActive to value and selects it for sending
func (in *ActionNodeSoftwareChangeIndexInput) SetNodeActive(value bool) *ActionNodeSoftwareChangeIndexInput {
	in.NodeActive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeActive"] = nil
	return in
}

// SetNodeKernelEvent sets parameter NodeKernelEvent to value and selects it for sending
func (in *ActionNodeSoftwareChangeIndexInput) SetNodeKernelEvent(value int64) *ActionNodeSoftwareChangeIndexInput {
	in.NodeKernelEvent = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeKernelEvent"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionNodeSoftwareChangeIndexInput) SetOffset(value int64) *ActionNodeSoftwareChangeIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNodeSoftwareChangeIndexInput) SetTo(value string) *ActionNodeSoftwareChangeIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeSoftwareChangeIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeSoftwareChangeIndexInput) SelectParameters(params ...string) *ActionNodeSoftwareChangeIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeSoftwareChangeIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeSoftwareChangeIndexInput) UnselectParameters(params ...string) *ActionNodeSoftwareChangeIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeSoftwareChangeIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeSoftwareChangeIndexOutput is a type for action output parameters
type ActionNodeSoftwareChangeIndexOutput struct {
	AfterRevision        string                              "json:\"after_revision\""
	AfterRevisionDirty   bool                                "json:\"after_revision_dirty\""
	AfterRevisionSource  string                              "json:\"after_revision_source\""
	AfterVersion         string                              "json:\"after_version\""
	AfterVersionSource   string                              "json:\"after_version_source\""
	BeforeRevision       string                              "json:\"before_revision\""
	BeforeRevisionDirty  bool                                "json:\"before_revision_dirty\""
	BeforeRevisionSource string                              "json:\"before_revision_source\""
	BeforeVersion        string                              "json:\"before_version\""
	BeforeVersionSource  string                              "json:\"before_version_source\""
	Component            string                              "json:\"component\""
	Generation           string                              "json:\"generation\""
	Id                   int64                               "json:\"id\""
	Node                 *ActionNodeShowOutput               "json:\"node\""
	NodeKernelEvent      *ActionNodeKernelEventShowOutput    "json:\"node_kernel_event\""
	NodeKernelEvidence   *ActionNodeKernelEvidenceShowOutput "json:\"node_kernel_evidence\""
	ObservedAfter        string                              "json:\"observed_after\""
	ObservedBefore       string                              "json:\"observed_before\""
	SourceRevision       string                              "json:\"source_revision\""
}

// Type for action response, including envelope
type ActionNodeSoftwareChangeIndexResponse struct {
	Action *ActionNodeSoftwareChangeIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeSoftwareChanges []*ActionNodeSoftwareChangeIndexOutput "json:\"node_software_changes\""
	}

	// Action output without the namespace
	Output []*ActionNodeSoftwareChangeIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeSoftwareChangeIndex) Prepare() *ActionNodeSoftwareChangeIndexInvocation {
	return &ActionNodeSoftwareChangeIndexInvocation{
		Action: action,
		Path:   "/v7.0/node_software_changes",
	}
}

// ActionNodeSoftwareChangeIndexInvocation is used to configure action for invocation
type ActionNodeSoftwareChangeIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeSoftwareChangeIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeSoftwareChangeIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeSoftwareChangeIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeSoftwareChangeIndexInvocation) NewInput() *ActionNodeSoftwareChangeIndexInput {
	inv.Input = &ActionNodeSoftwareChangeIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeSoftwareChangeIndexInvocation) SetInput(input *ActionNodeSoftwareChangeIndexInput) *ActionNodeSoftwareChangeIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeSoftwareChangeIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeSoftwareChangeIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeSoftwareChangeIndexInvocation) NewMetaInput() *ActionNodeSoftwareChangeIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeSoftwareChangeIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeSoftwareChangeIndexInvocation) SetMetaInput(input *ActionNodeSoftwareChangeIndexMetaGlobalInput) *ActionNodeSoftwareChangeIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeSoftwareChangeIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeSoftwareChangeIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeSoftwareChangeIndexInvocation) validate() error {
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
func (inv *ActionNodeSoftwareChangeIndexInvocation) Call() (*ActionNodeSoftwareChangeIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeSoftwareChangeIndexInvocation) callAsQuery() (*ActionNodeSoftwareChangeIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeSoftwareChangeIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeSoftwareChanges
	}
	return resp, err
}

func (inv *ActionNodeSoftwareChangeIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Component") {
			ret["node_software_change[component]"] = inv.Input.Component
		}
		if inv.IsParameterSelected("From") {
			ret["node_software_change[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("Generation") {
			ret["node_software_change[generation]"] = inv.Input.Generation
		}
		if inv.IsParameterSelected("Limit") {
			ret["node_software_change[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Node") {
			ret["node_software_change[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("NodeActive") {
			ret["node_software_change[node_active]"] = convertBoolToString(inv.Input.NodeActive)
		}
		if inv.IsParameterSelected("NodeKernelEvent") {
			ret["node_software_change[node_kernel_event]"] = convertInt64ToString(inv.Input.NodeKernelEvent)
		}
		if inv.IsParameterSelected("Offset") {
			ret["node_software_change[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("To") {
			ret["node_software_change[to]"] = inv.Input.To
		}
	}
}

func (inv *ActionNodeSoftwareChangeIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
