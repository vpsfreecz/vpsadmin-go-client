package client

import (
	"strings"
)

// ActionSessionTokenDelete is a type for action Session_token#Delete
type ActionSessionTokenDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionSessionTokenDelete(client *Client) *ActionSessionTokenDelete {
	return &ActionSessionTokenDelete{
		Client: client,
	}
}

// ActionSessionTokenDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionSessionTokenDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSessionTokenDeleteMetaGlobalInput) SetIncludes(value string) *ActionSessionTokenDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSessionTokenDeleteMetaGlobalInput) SetNo(value bool) *ActionSessionTokenDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSessionTokenDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSessionTokenDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionSessionTokenDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSessionTokenDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSessionTokenDeleteRequest is a type for the entire action request
type ActionSessionTokenDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionSessionTokenDeleteResponse struct {
	Action *ActionSessionTokenDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionSessionTokenDelete) Prepare() *ActionSessionTokenDeleteInvocation {
	return &ActionSessionTokenDeleteInvocation{
		Action: action,
		Path:   "/v6.0/session_tokens/{session_token_id}",
	}
}

// ActionSessionTokenDeleteInvocation is used to configure action for invocation
type ActionSessionTokenDeleteInvocation struct {
	// Pointer to the action
	Action *ActionSessionTokenDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionSessionTokenDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSessionTokenDeleteInvocation) SetPathParamInt(param string, value int64) *ActionSessionTokenDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSessionTokenDeleteInvocation) SetPathParamString(param string, value string) *ActionSessionTokenDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSessionTokenDeleteInvocation) NewMetaInput() *ActionSessionTokenDeleteMetaGlobalInput {
	inv.MetaInput = &ActionSessionTokenDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSessionTokenDeleteInvocation) SetMetaInput(input *ActionSessionTokenDeleteMetaGlobalInput) *ActionSessionTokenDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSessionTokenDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSessionTokenDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSessionTokenDeleteInvocation) Call() (*ActionSessionTokenDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionSessionTokenDeleteInvocation) callAsBody() (*ActionSessionTokenDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSessionTokenDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionSessionTokenDeleteInvocation) makeAllInputParams() *ActionSessionTokenDeleteRequest {
	return &ActionSessionTokenDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionSessionTokenDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
