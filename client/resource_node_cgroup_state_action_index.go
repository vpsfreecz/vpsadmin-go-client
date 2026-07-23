package client

import ()

// ActionNodeCgroupStateIndex is a type for action Node_cgroup_state#Index
type ActionNodeCgroupStateIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeCgroupStateIndex(client *Client) *ActionNodeCgroupStateIndex {
	return &ActionNodeCgroupStateIndex{
		Client: client,
	}
}

// ActionNodeCgroupStateIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeCgroupStateIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeCgroupStateIndexMetaGlobalInput) SetCount(value bool) *ActionNodeCgroupStateIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeCgroupStateIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeCgroupStateIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeCgroupStateIndexMetaGlobalInput) SetNo(value bool) *ActionNodeCgroupStateIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeCgroupStateIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeCgroupStateIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeCgroupStateIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeCgroupStateIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeCgroupStateIndexInput is a type for action input parameters
type ActionNodeCgroupStateIndexInput struct {
	Current    bool   "json:\"current\""
	From       string "json:\"from\""
	FromId     int64  "json:\"from_id\""
	Limit      int64  "json:\"limit\""
	Node       int64  "json:\"node\""
	NodeActive bool   "json:\"node_active\""
	To         string "json:\"to\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCurrent sets parameter Current to value and selects it for sending
func (in *ActionNodeCgroupStateIndexInput) SetCurrent(value bool) *ActionNodeCgroupStateIndexInput {
	in.Current = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Current"] = nil
	return in
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionNodeCgroupStateIndexInput) SetFrom(value string) *ActionNodeCgroupStateIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNodeCgroupStateIndexInput) SetFromId(value int64) *ActionNodeCgroupStateIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeCgroupStateIndexInput) SetLimit(value int64) *ActionNodeCgroupStateIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNodeCgroupStateIndexInput) SetNode(value int64) *ActionNodeCgroupStateIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeActive sets parameter NodeActive to value and selects it for sending
func (in *ActionNodeCgroupStateIndexInput) SetNodeActive(value bool) *ActionNodeCgroupStateIndexInput {
	in.NodeActive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeActive"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNodeCgroupStateIndexInput) SetTo(value string) *ActionNodeCgroupStateIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeCgroupStateIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeCgroupStateIndexInput) SelectParameters(params ...string) *ActionNodeCgroupStateIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeCgroupStateIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeCgroupStateIndexInput) UnselectParameters(params ...string) *ActionNodeCgroupStateIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeCgroupStateIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeCgroupStateIndexOutput is a type for action output parameters
type ActionNodeCgroupStateIndexOutput struct {
	CgroupVersion   string                "json:\"cgroup_version\""
	Current         bool                  "json:\"current\""
	FirstObservedAt string                "json:\"first_observed_at\""
	Id              int64                 "json:\"id\""
	LastObservedAt  string                "json:\"last_observed_at\""
	Node            *ActionNodeShowOutput "json:\"node\""
}

// Type for action response, including envelope
type ActionNodeCgroupStateIndexResponse struct {
	Action *ActionNodeCgroupStateIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeCgroupStates []*ActionNodeCgroupStateIndexOutput "json:\"node_cgroup_states\""
	}

	// Action output without the namespace
	Output []*ActionNodeCgroupStateIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeCgroupStateIndex) Prepare() *ActionNodeCgroupStateIndexInvocation {
	return &ActionNodeCgroupStateIndexInvocation{
		Action: action,
		Path:   "/v7.0/node_cgroup_states",
	}
}

// ActionNodeCgroupStateIndexInvocation is used to configure action for invocation
type ActionNodeCgroupStateIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeCgroupStateIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeCgroupStateIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeCgroupStateIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeCgroupStateIndexInvocation) NewInput() *ActionNodeCgroupStateIndexInput {
	inv.Input = &ActionNodeCgroupStateIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeCgroupStateIndexInvocation) SetInput(input *ActionNodeCgroupStateIndexInput) *ActionNodeCgroupStateIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeCgroupStateIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeCgroupStateIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeCgroupStateIndexInvocation) NewMetaInput() *ActionNodeCgroupStateIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeCgroupStateIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeCgroupStateIndexInvocation) SetMetaInput(input *ActionNodeCgroupStateIndexMetaGlobalInput) *ActionNodeCgroupStateIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeCgroupStateIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeCgroupStateIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeCgroupStateIndexInvocation) validate() error {
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
func (inv *ActionNodeCgroupStateIndexInvocation) Call() (*ActionNodeCgroupStateIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeCgroupStateIndexInvocation) callAsQuery() (*ActionNodeCgroupStateIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeCgroupStateIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeCgroupStates
	}
	return resp, err
}

func (inv *ActionNodeCgroupStateIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Current") {
			ret["node_cgroup_state[current]"] = convertBoolToString(inv.Input.Current)
		}
		if inv.IsParameterSelected("From") {
			ret["node_cgroup_state[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("FromId") {
			ret["node_cgroup_state[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["node_cgroup_state[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Node") {
			ret["node_cgroup_state[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("NodeActive") {
			ret["node_cgroup_state[node_active]"] = convertBoolToString(inv.Input.NodeActive)
		}
		if inv.IsParameterSelected("To") {
			ret["node_cgroup_state[to]"] = inv.Input.To
		}
	}
}

func (inv *ActionNodeCgroupStateIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
