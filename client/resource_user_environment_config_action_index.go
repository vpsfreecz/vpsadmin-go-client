package client

import (
	"strings"
)

// ActionUserEnvironmentConfigIndex is a type for action User.Environment_config#Index
type ActionUserEnvironmentConfigIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserEnvironmentConfigIndex(client *Client) *ActionUserEnvironmentConfigIndex {
	return &ActionUserEnvironmentConfigIndex{
		Client: client,
	}
}

// ActionUserEnvironmentConfigIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserEnvironmentConfigIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserEnvironmentConfigIndexMetaGlobalInput) SetCount(value bool) *ActionUserEnvironmentConfigIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserEnvironmentConfigIndexMetaGlobalInput) SetIncludes(value string) *ActionUserEnvironmentConfigIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserEnvironmentConfigIndexMetaGlobalInput) SetNo(value bool) *ActionUserEnvironmentConfigIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserEnvironmentConfigIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserEnvironmentConfigIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserEnvironmentConfigIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserEnvironmentConfigIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserEnvironmentConfigIndexInput is a type for action input parameters
type ActionUserEnvironmentConfigIndexInput struct {
	Environment int64 `json:"environment"`
	FromId      int64 `json:"from_id"`
	Limit       int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionUserEnvironmentConfigIndexInput) SetEnvironment(value int64) *ActionUserEnvironmentConfigIndexInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEnvironmentNil(false)
	in._selectedParameters["Environment"] = nil
	return in
}

// SetEnvironmentNil sets parameter Environment to nil and selects it for sending
func (in *ActionUserEnvironmentConfigIndexInput) SetEnvironmentNil(set bool) *ActionUserEnvironmentConfigIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Environment"] = nil
		in.SelectParameters("Environment")
	} else {
		delete(in._nilParameters, "Environment")
	}
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionUserEnvironmentConfigIndexInput) SetFromId(value int64) *ActionUserEnvironmentConfigIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserEnvironmentConfigIndexInput) SetLimit(value int64) *ActionUserEnvironmentConfigIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserEnvironmentConfigIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserEnvironmentConfigIndexInput) SelectParameters(params ...string) *ActionUserEnvironmentConfigIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserEnvironmentConfigIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserEnvironmentConfigIndexInput) UnselectParameters(params ...string) *ActionUserEnvironmentConfigIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserEnvironmentConfigIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserEnvironmentConfigIndexOutput is a type for action output parameters
type ActionUserEnvironmentConfigIndexOutput struct {
	CanCreateVps  bool                         `json:"can_create_vps"`
	CanDestroyVps bool                         `json:"can_destroy_vps"`
	Default       bool                         `json:"default"`
	Environment   *ActionEnvironmentShowOutput `json:"environment"`
	Id            int64                        `json:"id"`
	MaxVpsCount   int64                        `json:"max_vps_count"`
	VpsLifetime   int64                        `json:"vps_lifetime"`
}

// Type for action response, including envelope
type ActionUserEnvironmentConfigIndexResponse struct {
	Action *ActionUserEnvironmentConfigIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		EnvironmentConfigs []*ActionUserEnvironmentConfigIndexOutput `json:"environment_configs"`
	}

	// Action output without the namespace
	Output []*ActionUserEnvironmentConfigIndexOutput
}

// Prepare the action for invocation
func (action *ActionUserEnvironmentConfigIndex) Prepare() *ActionUserEnvironmentConfigIndexInvocation {
	return &ActionUserEnvironmentConfigIndexInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/environment_configs",
	}
}

// ActionUserEnvironmentConfigIndexInvocation is used to configure action for invocation
type ActionUserEnvironmentConfigIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserEnvironmentConfigIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserEnvironmentConfigIndexInput
	// Global meta input parameters
	MetaInput *ActionUserEnvironmentConfigIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserEnvironmentConfigIndexInvocation) SetPathParamInt(param string, value int64) *ActionUserEnvironmentConfigIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserEnvironmentConfigIndexInvocation) SetPathParamString(param string, value string) *ActionUserEnvironmentConfigIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserEnvironmentConfigIndexInvocation) NewInput() *ActionUserEnvironmentConfigIndexInput {
	inv.Input = &ActionUserEnvironmentConfigIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserEnvironmentConfigIndexInvocation) SetInput(input *ActionUserEnvironmentConfigIndexInput) *ActionUserEnvironmentConfigIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserEnvironmentConfigIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserEnvironmentConfigIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserEnvironmentConfigIndexInvocation) NewMetaInput() *ActionUserEnvironmentConfigIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserEnvironmentConfigIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserEnvironmentConfigIndexInvocation) SetMetaInput(input *ActionUserEnvironmentConfigIndexMetaGlobalInput) *ActionUserEnvironmentConfigIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserEnvironmentConfigIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserEnvironmentConfigIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserEnvironmentConfigIndexInvocation) Call() (*ActionUserEnvironmentConfigIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserEnvironmentConfigIndexInvocation) callAsQuery() (*ActionUserEnvironmentConfigIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserEnvironmentConfigIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.EnvironmentConfigs
	}
	return resp, err
}

func (inv *ActionUserEnvironmentConfigIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			ret["environment_config[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("FromId") {
			ret["environment_config[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["environment_config[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionUserEnvironmentConfigIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
