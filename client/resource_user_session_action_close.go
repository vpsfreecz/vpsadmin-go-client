package client

import (
	"strings"
)

// ActionUserSessionClose is a type for action User_session#Close
type ActionUserSessionClose struct {
	// Pointer to client
	Client *Client
}

func NewActionUserSessionClose(client *Client) *ActionUserSessionClose {
	return &ActionUserSessionClose{
		Client: client,
	}
}

// ActionUserSessionCloseMetaGlobalInput is a type for action global meta input parameters
type ActionUserSessionCloseMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserSessionCloseMetaGlobalInput) SetIncludes(value string) *ActionUserSessionCloseMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserSessionCloseMetaGlobalInput) SetNo(value bool) *ActionUserSessionCloseMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserSessionCloseMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserSessionCloseMetaGlobalInput) SelectParameters(params ...string) *ActionUserSessionCloseMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserSessionCloseMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserSessionCloseRequest is a type for the entire action request
type ActionUserSessionCloseRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionUserSessionCloseResponse struct {
	Action *ActionUserSessionClose `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionUserSessionClose) Prepare() *ActionUserSessionCloseInvocation {
	return &ActionUserSessionCloseInvocation{
		Action: action,
		Path:   "/v7.0/user_sessions/{user_session_id}",
	}
}

// ActionUserSessionCloseInvocation is used to configure action for invocation
type ActionUserSessionCloseInvocation struct {
	// Pointer to the action
	Action *ActionUserSessionClose

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserSessionCloseMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserSessionCloseInvocation) SetPathParamInt(param string, value int64) *ActionUserSessionCloseInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserSessionCloseInvocation) SetPathParamString(param string, value string) *ActionUserSessionCloseInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserSessionCloseInvocation) NewMetaInput() *ActionUserSessionCloseMetaGlobalInput {
	inv.MetaInput = &ActionUserSessionCloseMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserSessionCloseInvocation) SetMetaInput(input *ActionUserSessionCloseMetaGlobalInput) *ActionUserSessionCloseInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserSessionCloseInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserSessionCloseInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserSessionCloseInvocation) Call() (*ActionUserSessionCloseResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionUserSessionCloseInvocation) callAsBody() (*ActionUserSessionCloseResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserSessionCloseResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionUserSessionCloseInvocation) makeAllInputParams() *ActionUserSessionCloseRequest {
	return &ActionUserSessionCloseRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserSessionCloseInvocation) makeMetaInputParams() map[string]interface{} {
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
