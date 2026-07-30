package client

import (
	"net/url"
	"strings"
)

// ActionSystemConfigUpdate is a type for action System_config#Update
type ActionSystemConfigUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionSystemConfigUpdate(client *Client) *ActionSystemConfigUpdate {
	return &ActionSystemConfigUpdate{
		Client: client,
	}
}

// ActionSystemConfigUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionSystemConfigUpdateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSystemConfigUpdateMetaGlobalInput) SetIncludes(value string) *ActionSystemConfigUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSystemConfigUpdateMetaGlobalInput) SetNo(value bool) *ActionSystemConfigUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSystemConfigUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSystemConfigUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionSystemConfigUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSystemConfigUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSystemConfigUpdateInput is a type for action input parameters
type ActionSystemConfigUpdateInput struct {
	Value interface{} "json:\"value\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetValue sets parameter Value to value and selects it for sending
func (in *ActionSystemConfigUpdateInput) SetValue(value interface{}) *ActionSystemConfigUpdateInput {
	in.Value = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Value"] = nil
	return in
}

// SelectParameters sets parameters from ActionSystemConfigUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSystemConfigUpdateInput) SelectParameters(params ...string) *ActionSystemConfigUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionSystemConfigUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionSystemConfigUpdateInput) UnselectParameters(params ...string) *ActionSystemConfigUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionSystemConfigUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSystemConfigUpdateRequest is a type for the entire action request
type ActionSystemConfigUpdateRequest struct {
	SystemConfig map[string]interface{} "json:\"system_config\""
	Meta         map[string]interface{} "json:\"_meta\""
}

// ActionSystemConfigUpdateOutput is a type for action output parameters
type ActionSystemConfigUpdateOutput struct {
	Category       string      "json:\"category\""
	Description    string      "json:\"description\""
	Label          string      "json:\"label\""
	Localized      bool        "json:\"localized\""
	LocalizedValue interface{} "json:\"localized_value\""
	MinUserLevel   int64       "json:\"min_user_level\""
	Name           string      "json:\"name\""
	Type           string      "json:\"type\""
	Value          interface{} "json:\"value\""
}

// Type for action response, including envelope
type ActionSystemConfigUpdateResponse struct {
	Action *ActionSystemConfigUpdate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SystemConfig *ActionSystemConfigUpdateOutput "json:\"system_config\""
	}

	// Action output without the namespace
	Output *ActionSystemConfigUpdateOutput
}

// Prepare the action for invocation
func (action *ActionSystemConfigUpdate) Prepare() *ActionSystemConfigUpdateInvocation {
	return &ActionSystemConfigUpdateInvocation{
		Action: action,
		Path:   "/v7.0/system_configs/{category}/{name}",
	}
}

// ActionSystemConfigUpdateInvocation is used to configure action for invocation
type ActionSystemConfigUpdateInvocation struct {
	// Pointer to the action
	Action *ActionSystemConfigUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSystemConfigUpdateInput
	// Global meta input parameters
	MetaInput *ActionSystemConfigUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSystemConfigUpdateInvocation) SetPathParamInt(param string, value int64) *ActionSystemConfigUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSystemConfigUpdateInvocation) SetPathParamString(param string, value string) *ActionSystemConfigUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSystemConfigUpdateInvocation) NewInput() *ActionSystemConfigUpdateInput {
	inv.Input = &ActionSystemConfigUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSystemConfigUpdateInvocation) SetInput(input *ActionSystemConfigUpdateInput) *ActionSystemConfigUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSystemConfigUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionSystemConfigUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSystemConfigUpdateInvocation) NewMetaInput() *ActionSystemConfigUpdateMetaGlobalInput {
	inv.MetaInput = &ActionSystemConfigUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSystemConfigUpdateInvocation) SetMetaInput(input *ActionSystemConfigUpdateMetaGlobalInput) *ActionSystemConfigUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSystemConfigUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSystemConfigUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSystemConfigUpdateInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSystemConfigUpdateInvocation) Call() (*ActionSystemConfigUpdateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionSystemConfigUpdateInvocation) callAsBody() (*ActionSystemConfigUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSystemConfigUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SystemConfig
	}
	return resp, err
}

func (inv *ActionSystemConfigUpdateInvocation) makeAllInputParams() *ActionSystemConfigUpdateRequest {
	return &ActionSystemConfigUpdateRequest{
		SystemConfig: inv.makeInputParams(),
		Meta:         inv.makeMetaInputParams(),
	}
}

func (inv *ActionSystemConfigUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Value") {
			ret["value"] = inv.Input.Value
		}
	}

	return ret
}

func (inv *ActionSystemConfigUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
