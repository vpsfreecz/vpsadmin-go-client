package client

import (
	"strings"
)

// ActionUserWebauthnCredentialShow is a type for action User.Webauthn_credential#Show
type ActionUserWebauthnCredentialShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserWebauthnCredentialShow(client *Client) *ActionUserWebauthnCredentialShow {
	return &ActionUserWebauthnCredentialShow{
		Client: client,
	}
}

// ActionUserWebauthnCredentialShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserWebauthnCredentialShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserWebauthnCredentialShowMetaGlobalInput) SetIncludes(value string) *ActionUserWebauthnCredentialShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserWebauthnCredentialShowMetaGlobalInput) SetNo(value bool) *ActionUserWebauthnCredentialShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserWebauthnCredentialShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserWebauthnCredentialShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserWebauthnCredentialShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserWebauthnCredentialShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserWebauthnCredentialShowOutput is a type for action output parameters
type ActionUserWebauthnCredentialShowOutput struct {
	CreatedAt string `json:"created_at"`
	Enabled   bool   `json:"enabled"`
	Id        int64  `json:"id"`
	Label     string `json:"label"`
	LastUseAt string `json:"last_use_at"`
	UpdatedAt string `json:"updated_at"`
	UseCount  int64  `json:"use_count"`
}

// Type for action response, including envelope
type ActionUserWebauthnCredentialShowResponse struct {
	Action *ActionUserWebauthnCredentialShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		WebauthnCredential *ActionUserWebauthnCredentialShowOutput `json:"webauthn_credential"`
	}

	// Action output without the namespace
	Output *ActionUserWebauthnCredentialShowOutput
}

// Prepare the action for invocation
func (action *ActionUserWebauthnCredentialShow) Prepare() *ActionUserWebauthnCredentialShowInvocation {
	return &ActionUserWebauthnCredentialShowInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/webauthn_credentials/{webauthn_credential_id}",
	}
}

// ActionUserWebauthnCredentialShowInvocation is used to configure action for invocation
type ActionUserWebauthnCredentialShowInvocation struct {
	// Pointer to the action
	Action *ActionUserWebauthnCredentialShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserWebauthnCredentialShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserWebauthnCredentialShowInvocation) SetPathParamInt(param string, value int64) *ActionUserWebauthnCredentialShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserWebauthnCredentialShowInvocation) SetPathParamString(param string, value string) *ActionUserWebauthnCredentialShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserWebauthnCredentialShowInvocation) NewMetaInput() *ActionUserWebauthnCredentialShowMetaGlobalInput {
	inv.MetaInput = &ActionUserWebauthnCredentialShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserWebauthnCredentialShowInvocation) SetMetaInput(input *ActionUserWebauthnCredentialShowMetaGlobalInput) *ActionUserWebauthnCredentialShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserWebauthnCredentialShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserWebauthnCredentialShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserWebauthnCredentialShowInvocation) Call() (*ActionUserWebauthnCredentialShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserWebauthnCredentialShowInvocation) callAsQuery() (*ActionUserWebauthnCredentialShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserWebauthnCredentialShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.WebauthnCredential
	}
	return resp, err
}

func (inv *ActionUserWebauthnCredentialShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
