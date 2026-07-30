package client

import (
	"net/url"
	"strings"
)

// ActionWebuiUserSettingSet is a type for action Webui_user_setting#Set
type ActionWebuiUserSettingSet struct {
	// Pointer to client
	Client *Client
}

func NewActionWebuiUserSettingSet(client *Client) *ActionWebuiUserSettingSet {
	return &ActionWebuiUserSettingSet{
		Client: client,
	}
}

// ActionWebuiUserSettingSetMetaGlobalInput is a type for action global meta input parameters
type ActionWebuiUserSettingSetMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionWebuiUserSettingSetMetaGlobalInput) SetIncludes(value string) *ActionWebuiUserSettingSetMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionWebuiUserSettingSetMetaGlobalInput) SetNo(value bool) *ActionWebuiUserSettingSetMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionWebuiUserSettingSetMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionWebuiUserSettingSetMetaGlobalInput) SelectParameters(params ...string) *ActionWebuiUserSettingSetMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionWebuiUserSettingSetMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionWebuiUserSettingSetInput is a type for action input parameters
type ActionWebuiUserSettingSetInput struct {
	Value interface{} "json:\"value\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetValue sets parameter Value to value and selects it for sending
func (in *ActionWebuiUserSettingSetInput) SetValue(value interface{}) *ActionWebuiUserSettingSetInput {
	in.Value = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Value"] = nil
	return in
}

// SelectParameters sets parameters from ActionWebuiUserSettingSetInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionWebuiUserSettingSetInput) SelectParameters(params ...string) *ActionWebuiUserSettingSetInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionWebuiUserSettingSetInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionWebuiUserSettingSetInput) UnselectParameters(params ...string) *ActionWebuiUserSettingSetInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionWebuiUserSettingSetInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionWebuiUserSettingSetRequest is a type for the entire action request
type ActionWebuiUserSettingSetRequest struct {
	WebuiUserSetting map[string]interface{} "json:\"webui_user_setting\""
	Meta             map[string]interface{} "json:\"_meta\""
}

// ActionWebuiUserSettingSetOutput is a type for action output parameters
type ActionWebuiUserSettingSetOutput struct {
	CreatedAt string                "json:\"created_at\""
	Id        int64                 "json:\"id\""
	Key       string                "json:\"key\""
	Namespace string                "json:\"namespace\""
	UpdatedAt string                "json:\"updated_at\""
	User      *ActionUserShowOutput "json:\"user\""
	Value     interface{}           "json:\"value\""
}

// Type for action response, including envelope
type ActionWebuiUserSettingSetResponse struct {
	Action *ActionWebuiUserSettingSet "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		WebuiUserSetting *ActionWebuiUserSettingSetOutput "json:\"webui_user_setting\""
	}

	// Action output without the namespace
	Output *ActionWebuiUserSettingSetOutput
}

// Prepare the action for invocation
func (action *ActionWebuiUserSettingSet) Prepare() *ActionWebuiUserSettingSetInvocation {
	return &ActionWebuiUserSettingSetInvocation{
		Action: action,
		Path:   "/v7.0/webui_user_settings/{namespace}/{key}",
	}
}

// ActionWebuiUserSettingSetInvocation is used to configure action for invocation
type ActionWebuiUserSettingSetInvocation struct {
	// Pointer to the action
	Action *ActionWebuiUserSettingSet

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionWebuiUserSettingSetInput
	// Global meta input parameters
	MetaInput *ActionWebuiUserSettingSetMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionWebuiUserSettingSetInvocation) SetPathParamInt(param string, value int64) *ActionWebuiUserSettingSetInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionWebuiUserSettingSetInvocation) SetPathParamString(param string, value string) *ActionWebuiUserSettingSetInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionWebuiUserSettingSetInvocation) NewInput() *ActionWebuiUserSettingSetInput {
	inv.Input = &ActionWebuiUserSettingSetInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionWebuiUserSettingSetInvocation) SetInput(input *ActionWebuiUserSettingSetInput) *ActionWebuiUserSettingSetInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionWebuiUserSettingSetInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionWebuiUserSettingSetInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionWebuiUserSettingSetInvocation) NewMetaInput() *ActionWebuiUserSettingSetMetaGlobalInput {
	inv.MetaInput = &ActionWebuiUserSettingSetMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionWebuiUserSettingSetInvocation) SetMetaInput(input *ActionWebuiUserSettingSetMetaGlobalInput) *ActionWebuiUserSettingSetInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionWebuiUserSettingSetInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionWebuiUserSettingSetInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionWebuiUserSettingSetInvocation) validate() error {
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
func (inv *ActionWebuiUserSettingSetInvocation) Call() (*ActionWebuiUserSettingSetResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionWebuiUserSettingSetInvocation) callAsBody() (*ActionWebuiUserSettingSetResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionWebuiUserSettingSetResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.WebuiUserSetting
	}
	return resp, err
}

func (inv *ActionWebuiUserSettingSetInvocation) makeAllInputParams() *ActionWebuiUserSettingSetRequest {
	return &ActionWebuiUserSettingSetRequest{
		WebuiUserSetting: inv.makeInputParams(),
		Meta:             inv.makeMetaInputParams(),
	}
}

func (inv *ActionWebuiUserSettingSetInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Value") {
			ret["value"] = inv.Input.Value
		}
	}

	return ret
}

func (inv *ActionWebuiUserSettingSetInvocation) makeMetaInputParams() map[string]interface{} {
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
