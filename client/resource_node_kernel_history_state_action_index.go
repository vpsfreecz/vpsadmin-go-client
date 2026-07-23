package client

import ()

// ActionNodeKernelHistoryStateIndex is a type for action Node_kernel_history_state#Index
type ActionNodeKernelHistoryStateIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeKernelHistoryStateIndex(client *Client) *ActionNodeKernelHistoryStateIndex {
	return &ActionNodeKernelHistoryStateIndex{
		Client: client,
	}
}

// ActionNodeKernelHistoryStateIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeKernelHistoryStateIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeKernelHistoryStateIndexMetaGlobalInput) SetCount(value bool) *ActionNodeKernelHistoryStateIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeKernelHistoryStateIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeKernelHistoryStateIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeKernelHistoryStateIndexMetaGlobalInput) SetNo(value bool) *ActionNodeKernelHistoryStateIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelHistoryStateIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelHistoryStateIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeKernelHistoryStateIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeKernelHistoryStateIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelHistoryStateIndexInput is a type for action input parameters
type ActionNodeKernelHistoryStateIndexInput struct {
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

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionNodeKernelHistoryStateIndexInput) SetFrom(value string) *ActionNodeKernelHistoryStateIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNodeKernelHistoryStateIndexInput) SetFromId(value int64) *ActionNodeKernelHistoryStateIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeKernelHistoryStateIndexInput) SetLimit(value int64) *ActionNodeKernelHistoryStateIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNodeKernelHistoryStateIndexInput) SetNode(value int64) *ActionNodeKernelHistoryStateIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeActive sets parameter NodeActive to value and selects it for sending
func (in *ActionNodeKernelHistoryStateIndexInput) SetNodeActive(value bool) *ActionNodeKernelHistoryStateIndexInput {
	in.NodeActive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeActive"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNodeKernelHistoryStateIndexInput) SetTo(value string) *ActionNodeKernelHistoryStateIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelHistoryStateIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelHistoryStateIndexInput) SelectParameters(params ...string) *ActionNodeKernelHistoryStateIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeKernelHistoryStateIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeKernelHistoryStateIndexInput) UnselectParameters(params ...string) *ActionNodeKernelHistoryStateIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeKernelHistoryStateIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelHistoryStateIndexOutput is a type for action output parameters
type ActionNodeKernelHistoryStateIndexOutput struct {
	CompletedAt     string                "json:\"completed_at\""
	FromStatusId    int64                 "json:\"from_status_id\""
	Id              int64                 "json:\"id\""
	Node            *ActionNodeShowOutput "json:\"node\""
	ObservedThrough string                "json:\"observed_through\""
	StartedAt       string                "json:\"started_at\""
	ThroughStatusId int64                 "json:\"through_status_id\""
}

// Type for action response, including envelope
type ActionNodeKernelHistoryStateIndexResponse struct {
	Action *ActionNodeKernelHistoryStateIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeKernelHistoryStates []*ActionNodeKernelHistoryStateIndexOutput "json:\"node_kernel_history_states\""
	}

	// Action output without the namespace
	Output []*ActionNodeKernelHistoryStateIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeKernelHistoryStateIndex) Prepare() *ActionNodeKernelHistoryStateIndexInvocation {
	return &ActionNodeKernelHistoryStateIndexInvocation{
		Action: action,
		Path:   "/v7.0/node_kernel_history_states",
	}
}

// ActionNodeKernelHistoryStateIndexInvocation is used to configure action for invocation
type ActionNodeKernelHistoryStateIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeKernelHistoryStateIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeKernelHistoryStateIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeKernelHistoryStateIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeKernelHistoryStateIndexInvocation) NewInput() *ActionNodeKernelHistoryStateIndexInput {
	inv.Input = &ActionNodeKernelHistoryStateIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeKernelHistoryStateIndexInvocation) SetInput(input *ActionNodeKernelHistoryStateIndexInput) *ActionNodeKernelHistoryStateIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeKernelHistoryStateIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeKernelHistoryStateIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeKernelHistoryStateIndexInvocation) NewMetaInput() *ActionNodeKernelHistoryStateIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeKernelHistoryStateIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeKernelHistoryStateIndexInvocation) SetMetaInput(input *ActionNodeKernelHistoryStateIndexMetaGlobalInput) *ActionNodeKernelHistoryStateIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeKernelHistoryStateIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeKernelHistoryStateIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeKernelHistoryStateIndexInvocation) validate() error {
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
func (inv *ActionNodeKernelHistoryStateIndexInvocation) Call() (*ActionNodeKernelHistoryStateIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeKernelHistoryStateIndexInvocation) callAsQuery() (*ActionNodeKernelHistoryStateIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeKernelHistoryStateIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeKernelHistoryStates
	}
	return resp, err
}

func (inv *ActionNodeKernelHistoryStateIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("From") {
			ret["node_kernel_history_state[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("FromId") {
			ret["node_kernel_history_state[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["node_kernel_history_state[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Node") {
			ret["node_kernel_history_state[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("NodeActive") {
			ret["node_kernel_history_state[node_active]"] = convertBoolToString(inv.Input.NodeActive)
		}
		if inv.IsParameterSelected("To") {
			ret["node_kernel_history_state[to]"] = inv.Input.To
		}
	}
}

func (inv *ActionNodeKernelHistoryStateIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
