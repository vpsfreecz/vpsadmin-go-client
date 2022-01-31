package client

import (
	"strings"
)

// ActionSessionTokenUpdate is a type for action Session_token#Update
type ActionSessionTokenUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionSessionTokenUpdate(client *Client) *ActionSessionTokenUpdate {
	return &ActionSessionTokenUpdate{
		Client: client,
	}
}

// ActionSessionTokenUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionSessionTokenUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSessionTokenUpdateMetaGlobalInput) SetIncludes(value string) *ActionSessionTokenUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSessionTokenUpdateMetaGlobalInput) SetNo(value bool) *ActionSessionTokenUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSessionTokenUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSessionTokenUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionSessionTokenUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSessionTokenUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSessionTokenUpdateInput is a type for action input parameters
type ActionSessionTokenUpdateInput struct {
	Label string `json:"label"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionSessionTokenUpdateInput) SetLabel(value string) *ActionSessionTokenUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SelectParameters sets parameters from ActionSessionTokenUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSessionTokenUpdateInput) SelectParameters(params ...string) *ActionSessionTokenUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionSessionTokenUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionSessionTokenUpdateInput) UnselectParameters(params ...string) *ActionSessionTokenUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionSessionTokenUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSessionTokenUpdateRequest is a type for the entire action request
type ActionSessionTokenUpdateRequest struct {
	SessionToken map[string]interface{} `json:"session_token"`
	Meta         map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionSessionTokenUpdateResponse struct {
	Action *ActionSessionTokenUpdate `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionSessionTokenUpdate) Prepare() *ActionSessionTokenUpdateInvocation {
	return &ActionSessionTokenUpdateInvocation{
		Action: action,
		Path:   "/v6.0/session_tokens/{session_token_id}",
	}
}

// ActionSessionTokenUpdateInvocation is used to configure action for invocation
type ActionSessionTokenUpdateInvocation struct {
	// Pointer to the action
	Action *ActionSessionTokenUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSessionTokenUpdateInput
	// Global meta input parameters
	MetaInput *ActionSessionTokenUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSessionTokenUpdateInvocation) SetPathParamInt(param string, value int64) *ActionSessionTokenUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSessionTokenUpdateInvocation) SetPathParamString(param string, value string) *ActionSessionTokenUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSessionTokenUpdateInvocation) NewInput() *ActionSessionTokenUpdateInput {
	inv.Input = &ActionSessionTokenUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSessionTokenUpdateInvocation) SetInput(input *ActionSessionTokenUpdateInput) *ActionSessionTokenUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSessionTokenUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionSessionTokenUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSessionTokenUpdateInvocation) NewMetaInput() *ActionSessionTokenUpdateMetaGlobalInput {
	inv.MetaInput = &ActionSessionTokenUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSessionTokenUpdateInvocation) SetMetaInput(input *ActionSessionTokenUpdateMetaGlobalInput) *ActionSessionTokenUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSessionTokenUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSessionTokenUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSessionTokenUpdateInvocation) Call() (*ActionSessionTokenUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionSessionTokenUpdateInvocation) callAsBody() (*ActionSessionTokenUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSessionTokenUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionSessionTokenUpdateInvocation) makeAllInputParams() *ActionSessionTokenUpdateRequest {
	return &ActionSessionTokenUpdateRequest{
		SessionToken: inv.makeInputParams(),
		Meta:         inv.makeMetaInputParams(),
	}
}

func (inv *ActionSessionTokenUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
	}

	return ret
}

func (inv *ActionSessionTokenUpdateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
