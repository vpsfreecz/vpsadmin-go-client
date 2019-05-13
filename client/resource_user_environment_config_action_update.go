package client

import (
	"strings"
)

// ActionUserEnvironmentConfigUpdate is a type for action User.Environment_config#Update
type ActionUserEnvironmentConfigUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserEnvironmentConfigUpdate(client *Client) *ActionUserEnvironmentConfigUpdate {
	return &ActionUserEnvironmentConfigUpdate{
		Client: client,
	}
}

// ActionUserEnvironmentConfigUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionUserEnvironmentConfigUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserEnvironmentConfigUpdateMetaGlobalInput) SetNo(value bool) *ActionUserEnvironmentConfigUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserEnvironmentConfigUpdateMetaGlobalInput) SetIncludes(value string) *ActionUserEnvironmentConfigUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserEnvironmentConfigUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserEnvironmentConfigUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionUserEnvironmentConfigUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserEnvironmentConfigUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserEnvironmentConfigUpdateInput is a type for action input parameters
type ActionUserEnvironmentConfigUpdateInput struct {
	CanCreateVps bool `json:"can_create_vps"`
	CanDestroyVps bool `json:"can_destroy_vps"`
	VpsLifetime int64 `json:"vps_lifetime"`
	MaxVpsCount int64 `json:"max_vps_count"`
	Default bool `json:"default"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCanCreateVps sets parameter CanCreateVps to value and selects it for sending
func (in *ActionUserEnvironmentConfigUpdateInput) SetCanCreateVps(value bool) *ActionUserEnvironmentConfigUpdateInput {
	in.CanCreateVps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CanCreateVps"] = nil
	return in
}
// SetCanDestroyVps sets parameter CanDestroyVps to value and selects it for sending
func (in *ActionUserEnvironmentConfigUpdateInput) SetCanDestroyVps(value bool) *ActionUserEnvironmentConfigUpdateInput {
	in.CanDestroyVps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CanDestroyVps"] = nil
	return in
}
// SetVpsLifetime sets parameter VpsLifetime to value and selects it for sending
func (in *ActionUserEnvironmentConfigUpdateInput) SetVpsLifetime(value int64) *ActionUserEnvironmentConfigUpdateInput {
	in.VpsLifetime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["VpsLifetime"] = nil
	return in
}
// SetMaxVpsCount sets parameter MaxVpsCount to value and selects it for sending
func (in *ActionUserEnvironmentConfigUpdateInput) SetMaxVpsCount(value int64) *ActionUserEnvironmentConfigUpdateInput {
	in.MaxVpsCount = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxVpsCount"] = nil
	return in
}
// SetDefault sets parameter Default to value and selects it for sending
func (in *ActionUserEnvironmentConfigUpdateInput) SetDefault(value bool) *ActionUserEnvironmentConfigUpdateInput {
	in.Default = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Default"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserEnvironmentConfigUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserEnvironmentConfigUpdateInput) SelectParameters(params ...string) *ActionUserEnvironmentConfigUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserEnvironmentConfigUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserEnvironmentConfigUpdateRequest is a type for the entire action request
type ActionUserEnvironmentConfigUpdateRequest struct {
	EnvironmentConfig map[string]interface{} `json:"environment_config"`
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionUserEnvironmentConfigUpdateResponse struct {
	Action *ActionUserEnvironmentConfigUpdate `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionUserEnvironmentConfigUpdate) Prepare() *ActionUserEnvironmentConfigUpdateInvocation {
	return &ActionUserEnvironmentConfigUpdateInvocation{
		Action: action,
		Path: "/v5.0/users/{user_id}/environment_configs/{environment_config_id}",
	}
}

// ActionUserEnvironmentConfigUpdateInvocation is used to configure action for invocation
type ActionUserEnvironmentConfigUpdateInvocation struct {
	// Pointer to the action
	Action *ActionUserEnvironmentConfigUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserEnvironmentConfigUpdateInput
	// Global meta input parameters
	MetaInput *ActionUserEnvironmentConfigUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserEnvironmentConfigUpdateInvocation) SetPathParamInt(param string, value int64) *ActionUserEnvironmentConfigUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserEnvironmentConfigUpdateInvocation) SetPathParamString(param string, value string) *ActionUserEnvironmentConfigUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserEnvironmentConfigUpdateInvocation) NewInput() *ActionUserEnvironmentConfigUpdateInput {
	inv.Input = &ActionUserEnvironmentConfigUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserEnvironmentConfigUpdateInvocation) SetInput(input *ActionUserEnvironmentConfigUpdateInput) *ActionUserEnvironmentConfigUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserEnvironmentConfigUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserEnvironmentConfigUpdateInvocation) NewMetaInput() *ActionUserEnvironmentConfigUpdateMetaGlobalInput {
	inv.MetaInput = &ActionUserEnvironmentConfigUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserEnvironmentConfigUpdateInvocation) SetMetaInput(input *ActionUserEnvironmentConfigUpdateMetaGlobalInput) *ActionUserEnvironmentConfigUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserEnvironmentConfigUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserEnvironmentConfigUpdateInvocation) Call() (*ActionUserEnvironmentConfigUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionUserEnvironmentConfigUpdateInvocation) callAsBody() (*ActionUserEnvironmentConfigUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserEnvironmentConfigUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionUserEnvironmentConfigUpdateInvocation) makeAllInputParams() *ActionUserEnvironmentConfigUpdateRequest {
	return &ActionUserEnvironmentConfigUpdateRequest{
		EnvironmentConfig: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserEnvironmentConfigUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("CanCreateVps") {
			ret["can_create_vps"] = inv.Input.CanCreateVps
		}
		if inv.IsParameterSelected("CanDestroyVps") {
			ret["can_destroy_vps"] = inv.Input.CanDestroyVps
		}
		if inv.IsParameterSelected("VpsLifetime") {
			ret["vps_lifetime"] = inv.Input.VpsLifetime
		}
		if inv.IsParameterSelected("MaxVpsCount") {
			ret["max_vps_count"] = inv.Input.MaxVpsCount
		}
		if inv.IsParameterSelected("Default") {
			ret["default"] = inv.Input.Default
		}
	}

	return ret
}

func (inv *ActionUserEnvironmentConfigUpdateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
