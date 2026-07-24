package client

import ()

// ActionNodeKernelHistoryGapIndex is a type for action Node_kernel_history_gap#Index
type ActionNodeKernelHistoryGapIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeKernelHistoryGapIndex(client *Client) *ActionNodeKernelHistoryGapIndex {
	return &ActionNodeKernelHistoryGapIndex{
		Client: client,
	}
}

// ActionNodeKernelHistoryGapIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeKernelHistoryGapIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeKernelHistoryGapIndexMetaGlobalInput) SetCount(value bool) *ActionNodeKernelHistoryGapIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeKernelHistoryGapIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeKernelHistoryGapIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeKernelHistoryGapIndexMetaGlobalInput) SetNo(value bool) *ActionNodeKernelHistoryGapIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelHistoryGapIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelHistoryGapIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeKernelHistoryGapIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeKernelHistoryGapIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelHistoryGapIndexInput is a type for action input parameters
type ActionNodeKernelHistoryGapIndexInput struct {
	From                   string "json:\"from\""
	FromId                 int64  "json:\"from_id\""
	Limit                  int64  "json:\"limit\""
	Node                   int64  "json:\"node\""
	NodeActive             bool   "json:\"node_active\""
	NodeKernelHistoryState int64  "json:\"node_kernel_history_state\""
	To                     string "json:\"to\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionNodeKernelHistoryGapIndexInput) SetFrom(value string) *ActionNodeKernelHistoryGapIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNodeKernelHistoryGapIndexInput) SetFromId(value int64) *ActionNodeKernelHistoryGapIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeKernelHistoryGapIndexInput) SetLimit(value int64) *ActionNodeKernelHistoryGapIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNodeKernelHistoryGapIndexInput) SetNode(value int64) *ActionNodeKernelHistoryGapIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeActive sets parameter NodeActive to value and selects it for sending
func (in *ActionNodeKernelHistoryGapIndexInput) SetNodeActive(value bool) *ActionNodeKernelHistoryGapIndexInput {
	in.NodeActive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeActive"] = nil
	return in
}

// SetNodeKernelHistoryState sets parameter NodeKernelHistoryState to value and selects it for sending
func (in *ActionNodeKernelHistoryGapIndexInput) SetNodeKernelHistoryState(value int64) *ActionNodeKernelHistoryGapIndexInput {
	in.NodeKernelHistoryState = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeKernelHistoryState"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNodeKernelHistoryGapIndexInput) SetTo(value string) *ActionNodeKernelHistoryGapIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelHistoryGapIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelHistoryGapIndexInput) SelectParameters(params ...string) *ActionNodeKernelHistoryGapIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeKernelHistoryGapIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeKernelHistoryGapIndexInput) UnselectParameters(params ...string) *ActionNodeKernelHistoryGapIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeKernelHistoryGapIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelHistoryGapIndexOutput is a type for action output parameters
type ActionNodeKernelHistoryGapIndexOutput struct {
	From                   string                                  "json:\"from\""
	Id                     int64                                   "json:\"id\""
	Node                   *ActionNodeShowOutput                   "json:\"node\""
	NodeKernelHistoryState *ActionNodeKernelHistoryStateShowOutput "json:\"node_kernel_history_state\""
	Reason                 string                                  "json:\"reason\""
	To                     string                                  "json:\"to\""
}

// Type for action response, including envelope
type ActionNodeKernelHistoryGapIndexResponse struct {
	Action *ActionNodeKernelHistoryGapIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeKernelHistoryGaps []*ActionNodeKernelHistoryGapIndexOutput "json:\"node_kernel_history_gaps\""
	}

	// Action output without the namespace
	Output []*ActionNodeKernelHistoryGapIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeKernelHistoryGapIndex) Prepare() *ActionNodeKernelHistoryGapIndexInvocation {
	return &ActionNodeKernelHistoryGapIndexInvocation{
		Action: action,
		Path:   "/v7.0/node_kernel_history_gaps",
	}
}

// ActionNodeKernelHistoryGapIndexInvocation is used to configure action for invocation
type ActionNodeKernelHistoryGapIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeKernelHistoryGapIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeKernelHistoryGapIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeKernelHistoryGapIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeKernelHistoryGapIndexInvocation) NewInput() *ActionNodeKernelHistoryGapIndexInput {
	inv.Input = &ActionNodeKernelHistoryGapIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeKernelHistoryGapIndexInvocation) SetInput(input *ActionNodeKernelHistoryGapIndexInput) *ActionNodeKernelHistoryGapIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeKernelHistoryGapIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeKernelHistoryGapIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeKernelHistoryGapIndexInvocation) NewMetaInput() *ActionNodeKernelHistoryGapIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeKernelHistoryGapIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeKernelHistoryGapIndexInvocation) SetMetaInput(input *ActionNodeKernelHistoryGapIndexMetaGlobalInput) *ActionNodeKernelHistoryGapIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeKernelHistoryGapIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeKernelHistoryGapIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeKernelHistoryGapIndexInvocation) validate() error {
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
		if inv.IsParameterSelected("NodeKernelHistoryState") {
			if !inv.IsParameterNil("NodeKernelHistoryState") {
				if inv.Input.NodeKernelHistoryState < 0 {
					verr.Add("node_kernel_history_state", "not a valid resource id")
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
func (inv *ActionNodeKernelHistoryGapIndexInvocation) Call() (*ActionNodeKernelHistoryGapIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeKernelHistoryGapIndexInvocation) callAsQuery() (*ActionNodeKernelHistoryGapIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionNodeKernelHistoryGapIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeKernelHistoryGaps
	}
	return resp, err
}

func (inv *ActionNodeKernelHistoryGapIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("From") {
			ret["node_kernel_history_gap[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("FromId") {
			ret["node_kernel_history_gap[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["node_kernel_history_gap[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Node") {
			ret["node_kernel_history_gap[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("NodeActive") {
			ret["node_kernel_history_gap[node_active]"] = convertBoolToString(inv.Input.NodeActive)
		}
		if inv.IsParameterSelected("NodeKernelHistoryState") {
			ret["node_kernel_history_gap[node_kernel_history_state]"] = convertInt64ToString(inv.Input.NodeKernelHistoryState)
		}
		if inv.IsParameterSelected("To") {
			ret["node_kernel_history_gap[to]"] = inv.Input.To
		}
	}

	return nil
}

func (inv *ActionNodeKernelHistoryGapIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
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
