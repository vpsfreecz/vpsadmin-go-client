package client

import (
	"strings"
)

// ActionVpsConsoleTokenDelete is a type for action Vps.Console_token#Delete
type ActionVpsConsoleTokenDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsConsoleTokenDelete(client *Client) *ActionVpsConsoleTokenDelete {
	return &ActionVpsConsoleTokenDelete{
		Client: client,
	}
}

// ActionVpsConsoleTokenDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionVpsConsoleTokenDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsConsoleTokenDeleteMetaGlobalInput) SetIncludes(value string) *ActionVpsConsoleTokenDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsConsoleTokenDeleteMetaGlobalInput) SetNo(value bool) *ActionVpsConsoleTokenDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsConsoleTokenDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsConsoleTokenDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionVpsConsoleTokenDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsConsoleTokenDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsConsoleTokenDeleteRequest is a type for the entire action request
type ActionVpsConsoleTokenDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionVpsConsoleTokenDeleteResponse struct {
	Action *ActionVpsConsoleTokenDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionVpsConsoleTokenDelete) Prepare() *ActionVpsConsoleTokenDeleteInvocation {
	return &ActionVpsConsoleTokenDeleteInvocation{
		Action: action,
		Path:   "/v6.0/vpses/{vps_id}/console_token",
	}
}

// ActionVpsConsoleTokenDeleteInvocation is used to configure action for invocation
type ActionVpsConsoleTokenDeleteInvocation struct {
	// Pointer to the action
	Action *ActionVpsConsoleTokenDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionVpsConsoleTokenDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsConsoleTokenDeleteInvocation) SetPathParamInt(param string, value int64) *ActionVpsConsoleTokenDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsConsoleTokenDeleteInvocation) SetPathParamString(param string, value string) *ActionVpsConsoleTokenDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsConsoleTokenDeleteInvocation) NewMetaInput() *ActionVpsConsoleTokenDeleteMetaGlobalInput {
	inv.MetaInput = &ActionVpsConsoleTokenDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsConsoleTokenDeleteInvocation) SetMetaInput(input *ActionVpsConsoleTokenDeleteMetaGlobalInput) *ActionVpsConsoleTokenDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsConsoleTokenDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsConsoleTokenDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsConsoleTokenDeleteInvocation) Call() (*ActionVpsConsoleTokenDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionVpsConsoleTokenDeleteInvocation) callAsBody() (*ActionVpsConsoleTokenDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsConsoleTokenDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionVpsConsoleTokenDeleteInvocation) makeAllInputParams() *ActionVpsConsoleTokenDeleteRequest {
	return &ActionVpsConsoleTokenDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsConsoleTokenDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
