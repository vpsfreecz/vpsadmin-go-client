package client

import (
	"strings"
)

// ActionUserWebauthnCredentialUpdate is a type for action User.Webauthn_credential#Update
type ActionUserWebauthnCredentialUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserWebauthnCredentialUpdate(client *Client) *ActionUserWebauthnCredentialUpdate {
	return &ActionUserWebauthnCredentialUpdate{
		Client: client,
	}
}

// ActionUserWebauthnCredentialUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionUserWebauthnCredentialUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserWebauthnCredentialUpdateMetaGlobalInput) SetIncludes(value string) *ActionUserWebauthnCredentialUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserWebauthnCredentialUpdateMetaGlobalInput) SetNo(value bool) *ActionUserWebauthnCredentialUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserWebauthnCredentialUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserWebauthnCredentialUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionUserWebauthnCredentialUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserWebauthnCredentialUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserWebauthnCredentialUpdateInput is a type for action input parameters
type ActionUserWebauthnCredentialUpdateInput struct {
	Enabled bool   `json:"enabled"`
	Label   string `json:"label"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionUserWebauthnCredentialUpdateInput) SetEnabled(value bool) *ActionUserWebauthnCredentialUpdateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionUserWebauthnCredentialUpdateInput) SetLabel(value string) *ActionUserWebauthnCredentialUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserWebauthnCredentialUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserWebauthnCredentialUpdateInput) SelectParameters(params ...string) *ActionUserWebauthnCredentialUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserWebauthnCredentialUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserWebauthnCredentialUpdateInput) UnselectParameters(params ...string) *ActionUserWebauthnCredentialUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserWebauthnCredentialUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserWebauthnCredentialUpdateRequest is a type for the entire action request
type ActionUserWebauthnCredentialUpdateRequest struct {
	WebauthnCredential map[string]interface{} `json:"webauthn_credential"`
	Meta               map[string]interface{} `json:"_meta"`
}

// ActionUserWebauthnCredentialUpdateOutput is a type for action output parameters
type ActionUserWebauthnCredentialUpdateOutput struct {
	CreatedAt string `json:"created_at"`
	Enabled   bool   `json:"enabled"`
	Id        int64  `json:"id"`
	Label     string `json:"label"`
	LastUseAt string `json:"last_use_at"`
	UpdatedAt string `json:"updated_at"`
	UseCount  int64  `json:"use_count"`
}

// Type for action response, including envelope
type ActionUserWebauthnCredentialUpdateResponse struct {
	Action *ActionUserWebauthnCredentialUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		WebauthnCredential *ActionUserWebauthnCredentialUpdateOutput `json:"webauthn_credential"`
	}

	// Action output without the namespace
	Output *ActionUserWebauthnCredentialUpdateOutput
}

// Prepare the action for invocation
func (action *ActionUserWebauthnCredentialUpdate) Prepare() *ActionUserWebauthnCredentialUpdateInvocation {
	return &ActionUserWebauthnCredentialUpdateInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/webauthn_credentials/{webauthn_credential_id}",
	}
}

// ActionUserWebauthnCredentialUpdateInvocation is used to configure action for invocation
type ActionUserWebauthnCredentialUpdateInvocation struct {
	// Pointer to the action
	Action *ActionUserWebauthnCredentialUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserWebauthnCredentialUpdateInput
	// Global meta input parameters
	MetaInput *ActionUserWebauthnCredentialUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserWebauthnCredentialUpdateInvocation) SetPathParamInt(param string, value int64) *ActionUserWebauthnCredentialUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserWebauthnCredentialUpdateInvocation) SetPathParamString(param string, value string) *ActionUserWebauthnCredentialUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserWebauthnCredentialUpdateInvocation) NewInput() *ActionUserWebauthnCredentialUpdateInput {
	inv.Input = &ActionUserWebauthnCredentialUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserWebauthnCredentialUpdateInvocation) SetInput(input *ActionUserWebauthnCredentialUpdateInput) *ActionUserWebauthnCredentialUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserWebauthnCredentialUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserWebauthnCredentialUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserWebauthnCredentialUpdateInvocation) NewMetaInput() *ActionUserWebauthnCredentialUpdateMetaGlobalInput {
	inv.MetaInput = &ActionUserWebauthnCredentialUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserWebauthnCredentialUpdateInvocation) SetMetaInput(input *ActionUserWebauthnCredentialUpdateMetaGlobalInput) *ActionUserWebauthnCredentialUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserWebauthnCredentialUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserWebauthnCredentialUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserWebauthnCredentialUpdateInvocation) Call() (*ActionUserWebauthnCredentialUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionUserWebauthnCredentialUpdateInvocation) callAsBody() (*ActionUserWebauthnCredentialUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserWebauthnCredentialUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.WebauthnCredential
	}
	return resp, err
}

func (inv *ActionUserWebauthnCredentialUpdateInvocation) makeAllInputParams() *ActionUserWebauthnCredentialUpdateRequest {
	return &ActionUserWebauthnCredentialUpdateRequest{
		WebauthnCredential: inv.makeInputParams(),
		Meta:               inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserWebauthnCredentialUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Enabled") {
			ret["enabled"] = inv.Input.Enabled
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
	}

	return ret
}

func (inv *ActionUserWebauthnCredentialUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
