package client

import ()

// ActionDebugListObjectCounts is a type for action Debug#List_object_counts
type ActionDebugListObjectCounts struct {
	// Pointer to client
	Client *Client
}

func NewActionDebugListObjectCounts(client *Client) *ActionDebugListObjectCounts {
	return &ActionDebugListObjectCounts{
		Client: client,
	}
}

// ActionDebugListObjectCountsMetaGlobalInput is a type for action global meta input parameters
type ActionDebugListObjectCountsMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDebugListObjectCountsMetaGlobalInput) SetNo(value bool) *ActionDebugListObjectCountsMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDebugListObjectCountsMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDebugListObjectCountsMetaGlobalInput) SelectParameters(params ...string) *ActionDebugListObjectCountsMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDebugListObjectCountsMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDebugListObjectCountsOutput is a type for action output parameters
type ActionDebugListObjectCountsOutput struct {
	Count  int64  `json:"count"`
	Object string `json:"object"`
}

// Type for action response, including envelope
type ActionDebugListObjectCountsResponse struct {
	Action *ActionDebugListObjectCounts `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Debugs []*ActionDebugListObjectCountsOutput `json:"debugs"`
	}

	// Action output without the namespace
	Output []*ActionDebugListObjectCountsOutput
}

// Call the action directly without any path or input parameters
func (action *ActionDebugListObjectCounts) Call() (*ActionDebugListObjectCountsResponse, error) {
	return action.Prepare().Call()
}

// Prepare the action for invocation
func (action *ActionDebugListObjectCounts) Prepare() *ActionDebugListObjectCountsInvocation {
	return &ActionDebugListObjectCountsInvocation{
		Action: action,
		Path:   "/v7.0/debugs/list_object_counts",
	}
}

// ActionDebugListObjectCountsInvocation is used to configure action for invocation
type ActionDebugListObjectCountsInvocation struct {
	// Pointer to the action
	Action *ActionDebugListObjectCounts

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDebugListObjectCountsMetaGlobalInput
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDebugListObjectCountsInvocation) NewMetaInput() *ActionDebugListObjectCountsMetaGlobalInput {
	inv.MetaInput = &ActionDebugListObjectCountsMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDebugListObjectCountsInvocation) SetMetaInput(input *ActionDebugListObjectCountsMetaGlobalInput) *ActionDebugListObjectCountsInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDebugListObjectCountsInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDebugListObjectCountsInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDebugListObjectCountsInvocation) Call() (*ActionDebugListObjectCountsResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDebugListObjectCountsInvocation) callAsQuery() (*ActionDebugListObjectCountsResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDebugListObjectCountsResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Debugs
	}
	return resp, err
}

func (inv *ActionDebugListObjectCountsInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
