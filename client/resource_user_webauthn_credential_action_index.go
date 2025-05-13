package client

import (
	"strings"
)

// ActionUserWebauthnCredentialIndex is a type for action User.Webauthn_credential#Index
type ActionUserWebauthnCredentialIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserWebauthnCredentialIndex(client *Client) *ActionUserWebauthnCredentialIndex {
	return &ActionUserWebauthnCredentialIndex{
		Client: client,
	}
}

// ActionUserWebauthnCredentialIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserWebauthnCredentialIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserWebauthnCredentialIndexMetaGlobalInput) SetCount(value bool) *ActionUserWebauthnCredentialIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserWebauthnCredentialIndexMetaGlobalInput) SetIncludes(value string) *ActionUserWebauthnCredentialIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserWebauthnCredentialIndexMetaGlobalInput) SetNo(value bool) *ActionUserWebauthnCredentialIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserWebauthnCredentialIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserWebauthnCredentialIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserWebauthnCredentialIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserWebauthnCredentialIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserWebauthnCredentialIndexInput is a type for action input parameters
type ActionUserWebauthnCredentialIndexInput struct {
	Enabled bool  `json:"enabled"`
	FromId  int64 `json:"from_id"`
	Limit   int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionUserWebauthnCredentialIndexInput) SetEnabled(value bool) *ActionUserWebauthnCredentialIndexInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionUserWebauthnCredentialIndexInput) SetFromId(value int64) *ActionUserWebauthnCredentialIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserWebauthnCredentialIndexInput) SetLimit(value int64) *ActionUserWebauthnCredentialIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserWebauthnCredentialIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserWebauthnCredentialIndexInput) SelectParameters(params ...string) *ActionUserWebauthnCredentialIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserWebauthnCredentialIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserWebauthnCredentialIndexInput) UnselectParameters(params ...string) *ActionUserWebauthnCredentialIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserWebauthnCredentialIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserWebauthnCredentialIndexOutput is a type for action output parameters
type ActionUserWebauthnCredentialIndexOutput struct {
	CreatedAt string `json:"created_at"`
	Enabled   bool   `json:"enabled"`
	Id        int64  `json:"id"`
	Label     string `json:"label"`
	LastUseAt string `json:"last_use_at"`
	UpdatedAt string `json:"updated_at"`
	UseCount  int64  `json:"use_count"`
}

// Type for action response, including envelope
type ActionUserWebauthnCredentialIndexResponse struct {
	Action *ActionUserWebauthnCredentialIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		WebauthnCredentials []*ActionUserWebauthnCredentialIndexOutput `json:"webauthn_credentials"`
	}

	// Action output without the namespace
	Output []*ActionUserWebauthnCredentialIndexOutput
}

// Prepare the action for invocation
func (action *ActionUserWebauthnCredentialIndex) Prepare() *ActionUserWebauthnCredentialIndexInvocation {
	return &ActionUserWebauthnCredentialIndexInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/webauthn_credentials",
	}
}

// ActionUserWebauthnCredentialIndexInvocation is used to configure action for invocation
type ActionUserWebauthnCredentialIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserWebauthnCredentialIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserWebauthnCredentialIndexInput
	// Global meta input parameters
	MetaInput *ActionUserWebauthnCredentialIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserWebauthnCredentialIndexInvocation) SetPathParamInt(param string, value int64) *ActionUserWebauthnCredentialIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserWebauthnCredentialIndexInvocation) SetPathParamString(param string, value string) *ActionUserWebauthnCredentialIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserWebauthnCredentialIndexInvocation) NewInput() *ActionUserWebauthnCredentialIndexInput {
	inv.Input = &ActionUserWebauthnCredentialIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserWebauthnCredentialIndexInvocation) SetInput(input *ActionUserWebauthnCredentialIndexInput) *ActionUserWebauthnCredentialIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserWebauthnCredentialIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserWebauthnCredentialIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserWebauthnCredentialIndexInvocation) NewMetaInput() *ActionUserWebauthnCredentialIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserWebauthnCredentialIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserWebauthnCredentialIndexInvocation) SetMetaInput(input *ActionUserWebauthnCredentialIndexMetaGlobalInput) *ActionUserWebauthnCredentialIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserWebauthnCredentialIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserWebauthnCredentialIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserWebauthnCredentialIndexInvocation) Call() (*ActionUserWebauthnCredentialIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserWebauthnCredentialIndexInvocation) callAsQuery() (*ActionUserWebauthnCredentialIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserWebauthnCredentialIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.WebauthnCredentials
	}
	return resp, err
}

func (inv *ActionUserWebauthnCredentialIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Enabled") {
			ret["webauthn_credential[enabled]"] = convertBoolToString(inv.Input.Enabled)
		}
		if inv.IsParameterSelected("FromId") {
			ret["webauthn_credential[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["webauthn_credential[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionUserWebauthnCredentialIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
