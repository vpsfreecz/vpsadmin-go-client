package client

import (
	"strings"
)

// ActionOauth2ClientDelete is a type for action Oauth2_client#Delete
type ActionOauth2ClientDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionOauth2ClientDelete(client *Client) *ActionOauth2ClientDelete {
	return &ActionOauth2ClientDelete{
		Client: client,
	}
}

// ActionOauth2ClientDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionOauth2ClientDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOauth2ClientDeleteMetaGlobalInput) SetIncludes(value string) *ActionOauth2ClientDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOauth2ClientDeleteMetaGlobalInput) SetNo(value bool) *ActionOauth2ClientDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOauth2ClientDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOauth2ClientDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionOauth2ClientDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOauth2ClientDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOauth2ClientDeleteRequest is a type for the entire action request
type ActionOauth2ClientDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionOauth2ClientDeleteResponse struct {
	Action *ActionOauth2ClientDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionOauth2ClientDelete) Prepare() *ActionOauth2ClientDeleteInvocation {
	return &ActionOauth2ClientDeleteInvocation{
		Action: action,
		Path:   "/v7.0/oauth2_clients/{oauth2_client_id}",
	}
}

// ActionOauth2ClientDeleteInvocation is used to configure action for invocation
type ActionOauth2ClientDeleteInvocation struct {
	// Pointer to the action
	Action *ActionOauth2ClientDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionOauth2ClientDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOauth2ClientDeleteInvocation) SetPathParamInt(param string, value int64) *ActionOauth2ClientDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOauth2ClientDeleteInvocation) SetPathParamString(param string, value string) *ActionOauth2ClientDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOauth2ClientDeleteInvocation) NewMetaInput() *ActionOauth2ClientDeleteMetaGlobalInput {
	inv.MetaInput = &ActionOauth2ClientDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOauth2ClientDeleteInvocation) SetMetaInput(input *ActionOauth2ClientDeleteMetaGlobalInput) *ActionOauth2ClientDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOauth2ClientDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOauth2ClientDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOauth2ClientDeleteInvocation) Call() (*ActionOauth2ClientDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionOauth2ClientDeleteInvocation) callAsBody() (*ActionOauth2ClientDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOauth2ClientDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionOauth2ClientDeleteInvocation) makeAllInputParams() *ActionOauth2ClientDeleteRequest {
	return &ActionOauth2ClientDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionOauth2ClientDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
