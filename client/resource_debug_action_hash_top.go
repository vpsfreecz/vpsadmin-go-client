package client

import ()

// ActionDebugHashTop is a type for action Debug#Hash_top
type ActionDebugHashTop struct {
	// Pointer to client
	Client *Client
}

func NewActionDebugHashTop(client *Client) *ActionDebugHashTop {
	return &ActionDebugHashTop{
		Client: client,
	}
}

// ActionDebugHashTopMetaGlobalInput is a type for action global meta input parameters
type ActionDebugHashTopMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDebugHashTopMetaGlobalInput) SetNo(value bool) *ActionDebugHashTopMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDebugHashTopMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDebugHashTopMetaGlobalInput) SelectParameters(params ...string) *ActionDebugHashTopMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDebugHashTopMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDebugHashTopInput is a type for action input parameters
type ActionDebugHashTopInput struct {
	Limit int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDebugHashTopInput) SetLimit(value int64) *ActionDebugHashTopInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionDebugHashTopInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDebugHashTopInput) SelectParameters(params ...string) *ActionDebugHashTopInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDebugHashTopInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDebugHashTopInput) UnselectParameters(params ...string) *ActionDebugHashTopInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDebugHashTopInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDebugHashTopOutput is a type for action output parameters
type ActionDebugHashTopOutput struct {
	Size int64 `json:"size"`
}

// Type for action response, including envelope
type ActionDebugHashTopResponse struct {
	Action *ActionDebugHashTop `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Debugs []*ActionDebugHashTopOutput `json:"debugs"`
	}

	// Action output without the namespace
	Output []*ActionDebugHashTopOutput
}

// Prepare the action for invocation
func (action *ActionDebugHashTop) Prepare() *ActionDebugHashTopInvocation {
	return &ActionDebugHashTopInvocation{
		Action: action,
		Path:   "/v6.0/debugs/hash_top",
	}
}

// ActionDebugHashTopInvocation is used to configure action for invocation
type ActionDebugHashTopInvocation struct {
	// Pointer to the action
	Action *ActionDebugHashTop

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDebugHashTopInput
	// Global meta input parameters
	MetaInput *ActionDebugHashTopMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDebugHashTopInvocation) NewInput() *ActionDebugHashTopInput {
	inv.Input = &ActionDebugHashTopInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDebugHashTopInvocation) SetInput(input *ActionDebugHashTopInput) *ActionDebugHashTopInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDebugHashTopInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDebugHashTopInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDebugHashTopInvocation) NewMetaInput() *ActionDebugHashTopMetaGlobalInput {
	inv.MetaInput = &ActionDebugHashTopMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDebugHashTopInvocation) SetMetaInput(input *ActionDebugHashTopMetaGlobalInput) *ActionDebugHashTopInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDebugHashTopInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDebugHashTopInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDebugHashTopInvocation) Call() (*ActionDebugHashTopResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDebugHashTopInvocation) callAsQuery() (*ActionDebugHashTopResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDebugHashTopResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Debugs
	}
	return resp, err
}

func (inv *ActionDebugHashTopInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["debug[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionDebugHashTopInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
