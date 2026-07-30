package client

import ()

// ActionNodeSystemStateIndex is a type for action Node_system_state#Index
type ActionNodeSystemStateIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeSystemStateIndex(client *Client) *ActionNodeSystemStateIndex {
	return &ActionNodeSystemStateIndex{
		Client: client,
	}
}

// ActionNodeSystemStateIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeSystemStateIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeSystemStateIndexMetaGlobalInput) SetCount(value bool) *ActionNodeSystemStateIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeSystemStateIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeSystemStateIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeSystemStateIndexMetaGlobalInput) SetNo(value bool) *ActionNodeSystemStateIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeSystemStateIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeSystemStateIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeSystemStateIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeSystemStateIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeSystemStateIndexInput is a type for action input parameters
type ActionNodeSystemStateIndexInput struct {
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
func (in *ActionNodeSystemStateIndexInput) SetCurrent(value bool) *ActionNodeSystemStateIndexInput {
	in.Current = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Current"] = nil
	return in
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionNodeSystemStateIndexInput) SetFrom(value string) *ActionNodeSystemStateIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNodeSystemStateIndexInput) SetFromId(value int64) *ActionNodeSystemStateIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeSystemStateIndexInput) SetLimit(value int64) *ActionNodeSystemStateIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNodeSystemStateIndexInput) SetNode(value int64) *ActionNodeSystemStateIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeActive sets parameter NodeActive to value and selects it for sending
func (in *ActionNodeSystemStateIndexInput) SetNodeActive(value bool) *ActionNodeSystemStateIndexInput {
	in.NodeActive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeActive"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNodeSystemStateIndexInput) SetTo(value string) *ActionNodeSystemStateIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeSystemStateIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeSystemStateIndexInput) SelectParameters(params ...string) *ActionNodeSystemStateIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeSystemStateIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeSystemStateIndexInput) UnselectParameters(params ...string) *ActionNodeSystemStateIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeSystemStateIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeSystemStateIndexOutput is a type for action output parameters
type ActionNodeSystemStateIndexOutput struct {
	CgroupVersion   string                "json:\"cgroup_version\""
	Cpus            int64                 "json:\"cpus\""
	Current         bool                  "json:\"current\""
	FirstObservedAt string                "json:\"first_observed_at\""
	Id              int64                 "json:\"id\""
	LastObservedAt  string                "json:\"last_observed_at\""
	Node            *ActionNodeShowOutput "json:\"node\""
	TotalMemory     int64                 "json:\"total_memory\""
	TotalSwap       int64                 "json:\"total_swap\""
}

// Type for action response, including envelope
type ActionNodeSystemStateIndexResponse struct {
	Action *ActionNodeSystemStateIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeSystemStates []*ActionNodeSystemStateIndexOutput "json:\"node_system_states\""
	}

	// Action output without the namespace
	Output []*ActionNodeSystemStateIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeSystemStateIndex) Prepare() *ActionNodeSystemStateIndexInvocation {
	return &ActionNodeSystemStateIndexInvocation{
		Action: action,
		Path:   "/v7.0/node_system_states",
	}
}

// ActionNodeSystemStateIndexInvocation is used to configure action for invocation
type ActionNodeSystemStateIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeSystemStateIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeSystemStateIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeSystemStateIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeSystemStateIndexInvocation) NewInput() *ActionNodeSystemStateIndexInput {
	inv.Input = &ActionNodeSystemStateIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeSystemStateIndexInvocation) SetInput(input *ActionNodeSystemStateIndexInput) *ActionNodeSystemStateIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeSystemStateIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeSystemStateIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeSystemStateIndexInvocation) NewMetaInput() *ActionNodeSystemStateIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeSystemStateIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeSystemStateIndexInvocation) SetMetaInput(input *ActionNodeSystemStateIndexMetaGlobalInput) *ActionNodeSystemStateIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeSystemStateIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeSystemStateIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeSystemStateIndexInvocation) validate() error {
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
func (inv *ActionNodeSystemStateIndexInvocation) Call() (*ActionNodeSystemStateIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeSystemStateIndexInvocation) callAsQuery() (*ActionNodeSystemStateIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionNodeSystemStateIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeSystemStates
	}
	return resp, err
}

func (inv *ActionNodeSystemStateIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("Current") {
			ret["node_system_state[current]"] = convertBoolToString(inv.Input.Current)
		}
		if inv.IsParameterSelected("From") {
			ret["node_system_state[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("FromId") {
			ret["node_system_state[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["node_system_state[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Node") {
			ret["node_system_state[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("NodeActive") {
			ret["node_system_state[node_active]"] = convertBoolToString(inv.Input.NodeActive)
		}
		if inv.IsParameterSelected("To") {
			ret["node_system_state[to]"] = inv.Input.To
		}
	}

	return nil
}

func (inv *ActionNodeSystemStateIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
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
