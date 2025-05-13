package client

import (
	"strings"
)

// ActionUserWebauthnCredentialDelete is a type for action User.Webauthn_credential#Delete
type ActionUserWebauthnCredentialDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionUserWebauthnCredentialDelete(client *Client) *ActionUserWebauthnCredentialDelete {
	return &ActionUserWebauthnCredentialDelete{
		Client: client,
	}
}

// ActionUserWebauthnCredentialDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionUserWebauthnCredentialDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserWebauthnCredentialDeleteMetaGlobalInput) SetIncludes(value string) *ActionUserWebauthnCredentialDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserWebauthnCredentialDeleteMetaGlobalInput) SetNo(value bool) *ActionUserWebauthnCredentialDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserWebauthnCredentialDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserWebauthnCredentialDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionUserWebauthnCredentialDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserWebauthnCredentialDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserWebauthnCredentialDeleteRequest is a type for the entire action request
type ActionUserWebauthnCredentialDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionUserWebauthnCredentialDeleteResponse struct {
	Action *ActionUserWebauthnCredentialDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionUserWebauthnCredentialDelete) Prepare() *ActionUserWebauthnCredentialDeleteInvocation {
	return &ActionUserWebauthnCredentialDeleteInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/webauthn_credentials/{webauthn_credential_id}",
	}
}

// ActionUserWebauthnCredentialDeleteInvocation is used to configure action for invocation
type ActionUserWebauthnCredentialDeleteInvocation struct {
	// Pointer to the action
	Action *ActionUserWebauthnCredentialDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserWebauthnCredentialDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserWebauthnCredentialDeleteInvocation) SetPathParamInt(param string, value int64) *ActionUserWebauthnCredentialDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserWebauthnCredentialDeleteInvocation) SetPathParamString(param string, value string) *ActionUserWebauthnCredentialDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserWebauthnCredentialDeleteInvocation) NewMetaInput() *ActionUserWebauthnCredentialDeleteMetaGlobalInput {
	inv.MetaInput = &ActionUserWebauthnCredentialDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserWebauthnCredentialDeleteInvocation) SetMetaInput(input *ActionUserWebauthnCredentialDeleteMetaGlobalInput) *ActionUserWebauthnCredentialDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserWebauthnCredentialDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserWebauthnCredentialDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserWebauthnCredentialDeleteInvocation) Call() (*ActionUserWebauthnCredentialDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionUserWebauthnCredentialDeleteInvocation) callAsBody() (*ActionUserWebauthnCredentialDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserWebauthnCredentialDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionUserWebauthnCredentialDeleteInvocation) makeAllInputParams() *ActionUserWebauthnCredentialDeleteRequest {
	return &ActionUserWebauthnCredentialDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserWebauthnCredentialDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
