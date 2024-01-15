package client

import ()

// ActionDebugArrayTop is a type for action Debug#Array_top
type ActionDebugArrayTop struct {
	// Pointer to client
	Client *Client
}

func NewActionDebugArrayTop(client *Client) *ActionDebugArrayTop {
	return &ActionDebugArrayTop{
		Client: client,
	}
}

// ActionDebugArrayTopMetaGlobalInput is a type for action global meta input parameters
type ActionDebugArrayTopMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDebugArrayTopMetaGlobalInput) SetNo(value bool) *ActionDebugArrayTopMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDebugArrayTopMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDebugArrayTopMetaGlobalInput) SelectParameters(params ...string) *ActionDebugArrayTopMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDebugArrayTopMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDebugArrayTopInput is a type for action input parameters
type ActionDebugArrayTopInput struct {
	Limit int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDebugArrayTopInput) SetLimit(value int64) *ActionDebugArrayTopInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionDebugArrayTopInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDebugArrayTopInput) SelectParameters(params ...string) *ActionDebugArrayTopInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDebugArrayTopInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDebugArrayTopInput) UnselectParameters(params ...string) *ActionDebugArrayTopInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDebugArrayTopInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDebugArrayTopOutput is a type for action output parameters
type ActionDebugArrayTopOutput struct {
	Size int64 `json:"size"`
}

// Type for action response, including envelope
type ActionDebugArrayTopResponse struct {
	Action *ActionDebugArrayTop `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Debugs []*ActionDebugArrayTopOutput `json:"debugs"`
	}

	// Action output without the namespace
	Output []*ActionDebugArrayTopOutput
}

// Prepare the action for invocation
func (action *ActionDebugArrayTop) Prepare() *ActionDebugArrayTopInvocation {
	return &ActionDebugArrayTopInvocation{
		Action: action,
		Path:   "/v6.0/debugs/array_top",
	}
}

// ActionDebugArrayTopInvocation is used to configure action for invocation
type ActionDebugArrayTopInvocation struct {
	// Pointer to the action
	Action *ActionDebugArrayTop

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDebugArrayTopInput
	// Global meta input parameters
	MetaInput *ActionDebugArrayTopMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDebugArrayTopInvocation) NewInput() *ActionDebugArrayTopInput {
	inv.Input = &ActionDebugArrayTopInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDebugArrayTopInvocation) SetInput(input *ActionDebugArrayTopInput) *ActionDebugArrayTopInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDebugArrayTopInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDebugArrayTopInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDebugArrayTopInvocation) NewMetaInput() *ActionDebugArrayTopMetaGlobalInput {
	inv.MetaInput = &ActionDebugArrayTopMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDebugArrayTopInvocation) SetMetaInput(input *ActionDebugArrayTopMetaGlobalInput) *ActionDebugArrayTopInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDebugArrayTopInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDebugArrayTopInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDebugArrayTopInvocation) Call() (*ActionDebugArrayTopResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDebugArrayTopInvocation) callAsQuery() (*ActionDebugArrayTopResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDebugArrayTopResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Debugs
	}
	return resp, err
}

func (inv *ActionDebugArrayTopInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["debug[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionDebugArrayTopInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
