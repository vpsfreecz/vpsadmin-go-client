package client

import (
	"strings"
)

// ActionAuthTokenUpdate is a type for action Auth_token#Update
type ActionAuthTokenUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionAuthTokenUpdate(client *Client) *ActionAuthTokenUpdate {
	return &ActionAuthTokenUpdate{
		Client: client,
	}
}

// ActionAuthTokenUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionAuthTokenUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionAuthTokenUpdateMetaGlobalInput) SetNo(value bool) *ActionAuthTokenUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionAuthTokenUpdateMetaGlobalInput) SetIncludes(value string) *ActionAuthTokenUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionAuthTokenUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionAuthTokenUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionAuthTokenUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionAuthTokenUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionAuthTokenUpdateInput is a type for action input parameters
type ActionAuthTokenUpdateInput struct {
	Label string `json:"label"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionAuthTokenUpdateInput) SetLabel(value string) *ActionAuthTokenUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SelectParameters sets parameters from ActionAuthTokenUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionAuthTokenUpdateInput) SelectParameters(params ...string) *ActionAuthTokenUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionAuthTokenUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionAuthTokenUpdateRequest is a type for the entire action request
type ActionAuthTokenUpdateRequest struct {
	AuthToken map[string]interface{} `json:"auth_token"`
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionAuthTokenUpdateResponse struct {
	Action *ActionAuthTokenUpdate `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionAuthTokenUpdate) Prepare() *ActionAuthTokenUpdateInvocation {
	return &ActionAuthTokenUpdateInvocation{
		Action: action,
		Path: "/v5.0/auth_tokens/:auth_token_id",
	}
}

// ActionAuthTokenUpdateInvocation is used to configure action for invocation
type ActionAuthTokenUpdateInvocation struct {
	// Pointer to the action
	Action *ActionAuthTokenUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionAuthTokenUpdateInput
	// Global meta input parameters
	MetaInput *ActionAuthTokenUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionAuthTokenUpdateInvocation) SetPathParamInt(param string, value int64) *ActionAuthTokenUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionAuthTokenUpdateInvocation) SetPathParamString(param string, value string) *ActionAuthTokenUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionAuthTokenUpdateInvocation) NewInput() *ActionAuthTokenUpdateInput {
	inv.Input = &ActionAuthTokenUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionAuthTokenUpdateInvocation) SetInput(input *ActionAuthTokenUpdateInput) *ActionAuthTokenUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionAuthTokenUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionAuthTokenUpdateInvocation) NewMetaInput() *ActionAuthTokenUpdateMetaGlobalInput {
	inv.MetaInput = &ActionAuthTokenUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionAuthTokenUpdateInvocation) SetMetaInput(input *ActionAuthTokenUpdateMetaGlobalInput) *ActionAuthTokenUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionAuthTokenUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionAuthTokenUpdateInvocation) Call() (*ActionAuthTokenUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionAuthTokenUpdateInvocation) callAsBody() (*ActionAuthTokenUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionAuthTokenUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionAuthTokenUpdateInvocation) makeAllInputParams() *ActionAuthTokenUpdateRequest {
	return &ActionAuthTokenUpdateRequest{
		AuthToken: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionAuthTokenUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
	}

	return ret
}

func (inv *ActionAuthTokenUpdateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
