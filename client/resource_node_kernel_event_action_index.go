package client

import ()

// ActionNodeKernelEventIndex is a type for action Node_kernel_event#Index
type ActionNodeKernelEventIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeKernelEventIndex(client *Client) *ActionNodeKernelEventIndex {
	return &ActionNodeKernelEventIndex{
		Client: client,
	}
}

// ActionNodeKernelEventIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeKernelEventIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeKernelEventIndexMetaGlobalInput) SetCount(value bool) *ActionNodeKernelEventIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeKernelEventIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeKernelEventIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeKernelEventIndexMetaGlobalInput) SetNo(value bool) *ActionNodeKernelEventIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelEventIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelEventIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeKernelEventIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeKernelEventIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelEventIndexInput is a type for action input parameters
type ActionNodeKernelEventIndexInput struct {
	Confidence  string "json:\"confidence\""
	Current     bool   "json:\"current\""
	EventSource string "json:\"event_source\""
	EventType   string "json:\"event_type\""
	From        string "json:\"from\""
	FromId      int64  "json:\"from_id\""
	Limit       int64  "json:\"limit\""
	Node        int64  "json:\"node\""
	NodeActive  bool   "json:\"node_active\""
	To          string "json:\"to\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetConfidence sets parameter Confidence to value and selects it for sending
func (in *ActionNodeKernelEventIndexInput) SetConfidence(value string) *ActionNodeKernelEventIndexInput {
	in.Confidence = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Confidence"] = nil
	return in
}

// SetCurrent sets parameter Current to value and selects it for sending
func (in *ActionNodeKernelEventIndexInput) SetCurrent(value bool) *ActionNodeKernelEventIndexInput {
	in.Current = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Current"] = nil
	return in
}

// SetEventSource sets parameter EventSource to value and selects it for sending
func (in *ActionNodeKernelEventIndexInput) SetEventSource(value string) *ActionNodeKernelEventIndexInput {
	in.EventSource = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EventSource"] = nil
	return in
}

// SetEventType sets parameter EventType to value and selects it for sending
func (in *ActionNodeKernelEventIndexInput) SetEventType(value string) *ActionNodeKernelEventIndexInput {
	in.EventType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EventType"] = nil
	return in
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionNodeKernelEventIndexInput) SetFrom(value string) *ActionNodeKernelEventIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNodeKernelEventIndexInput) SetFromId(value int64) *ActionNodeKernelEventIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeKernelEventIndexInput) SetLimit(value int64) *ActionNodeKernelEventIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNodeKernelEventIndexInput) SetNode(value int64) *ActionNodeKernelEventIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeActive sets parameter NodeActive to value and selects it for sending
func (in *ActionNodeKernelEventIndexInput) SetNodeActive(value bool) *ActionNodeKernelEventIndexInput {
	in.NodeActive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeActive"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNodeKernelEventIndexInput) SetTo(value string) *ActionNodeKernelEventIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelEventIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelEventIndexInput) SelectParameters(params ...string) *ActionNodeKernelEventIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeKernelEventIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeKernelEventIndexInput) UnselectParameters(params ...string) *ActionNodeKernelEventIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeKernelEventIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelEventIndexOutput is a type for action output parameters
type ActionNodeKernelEventIndexOutput struct {
	BootId                string                              "json:\"boot_id\""
	BootedAt              string                              "json:\"booted_at\""
	BootedRelease         string                              "json:\"booted_release\""
	BootedSystem          string                              "json:\"booted_system\""
	Confidence            string                              "json:\"confidence\""
	Current               bool                                "json:\"current\""
	CurrentSystem         string                              "json:\"current_system\""
	EffectiveAt           string                              "json:\"effective_at\""
	EventType             string                              "json:\"event_type\""
	EvidenceRevision      string                              "json:\"evidence_revision\""
	Id                    int64                               "json:\"id\""
	KernelCommandLine     string                              "json:\"kernel_command_line\""
	KernelConfigAvailable bool                                "json:\"kernel_config_available\""
	KernelConfigDigest    string                              "json:\"kernel_config_digest\""
	KernelSourceRevision  string                              "json:\"kernel_source_revision\""
	Node                  *ActionNodeShowOutput               "json:\"node\""
	NodeKernelEvidence    *ActionNodeKernelEvidenceShowOutput "json:\"node_kernel_evidence\""
	ObservedAfter         string                              "json:\"observed_after\""
	ObservedBefore        string                              "json:\"observed_before\""
	ReportSchemaVersion   int64                               "json:\"report_schema_version\""
	ReportedRelease       string                              "json:\"reported_release\""
	SnapshotRevision      string                              "json:\"snapshot_revision\""
	Source                string                              "json:\"source\""
	SourceStatusId        int64                               "json:\"source_status_id\""
}

// Type for action response, including envelope
type ActionNodeKernelEventIndexResponse struct {
	Action *ActionNodeKernelEventIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeKernelEvents []*ActionNodeKernelEventIndexOutput "json:\"node_kernel_events\""
	}

	// Action output without the namespace
	Output []*ActionNodeKernelEventIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeKernelEventIndex) Prepare() *ActionNodeKernelEventIndexInvocation {
	return &ActionNodeKernelEventIndexInvocation{
		Action: action,
		Path:   "/v7.0/node_kernel_events",
	}
}

// ActionNodeKernelEventIndexInvocation is used to configure action for invocation
type ActionNodeKernelEventIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeKernelEventIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeKernelEventIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeKernelEventIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeKernelEventIndexInvocation) NewInput() *ActionNodeKernelEventIndexInput {
	inv.Input = &ActionNodeKernelEventIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeKernelEventIndexInvocation) SetInput(input *ActionNodeKernelEventIndexInput) *ActionNodeKernelEventIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeKernelEventIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeKernelEventIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeKernelEventIndexInvocation) NewMetaInput() *ActionNodeKernelEventIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeKernelEventIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeKernelEventIndexInvocation) SetMetaInput(input *ActionNodeKernelEventIndexMetaGlobalInput) *ActionNodeKernelEventIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeKernelEventIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeKernelEventIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeKernelEventIndexInvocation) validate() error {
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
func (inv *ActionNodeKernelEventIndexInvocation) Call() (*ActionNodeKernelEventIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeKernelEventIndexInvocation) callAsQuery() (*ActionNodeKernelEventIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionNodeKernelEventIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeKernelEvents
	}
	return resp, err
}

func (inv *ActionNodeKernelEventIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("Confidence") {
			ret["node_kernel_event[confidence]"] = inv.Input.Confidence
		}
		if inv.IsParameterSelected("Current") {
			ret["node_kernel_event[current]"] = convertBoolToString(inv.Input.Current)
		}
		if inv.IsParameterSelected("EventSource") {
			ret["node_kernel_event[event_source]"] = inv.Input.EventSource
		}
		if inv.IsParameterSelected("EventType") {
			ret["node_kernel_event[event_type]"] = inv.Input.EventType
		}
		if inv.IsParameterSelected("From") {
			ret["node_kernel_event[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("FromId") {
			ret["node_kernel_event[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["node_kernel_event[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Node") {
			ret["node_kernel_event[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("NodeActive") {
			ret["node_kernel_event[node_active]"] = convertBoolToString(inv.Input.NodeActive)
		}
		if inv.IsParameterSelected("To") {
			ret["node_kernel_event[to]"] = inv.Input.To
		}
	}

	return nil
}

func (inv *ActionNodeKernelEventIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
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
