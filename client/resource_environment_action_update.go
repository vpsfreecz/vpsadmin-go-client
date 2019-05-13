package client

import (
	"strings"
)

// ActionEnvironmentUpdate is a type for action Environment#Update
type ActionEnvironmentUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionEnvironmentUpdate(client *Client) *ActionEnvironmentUpdate {
	return &ActionEnvironmentUpdate{
		Client: client,
	}
}

// ActionEnvironmentUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionEnvironmentUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEnvironmentUpdateMetaGlobalInput) SetNo(value bool) *ActionEnvironmentUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEnvironmentUpdateMetaGlobalInput) SetIncludes(value string) *ActionEnvironmentUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionEnvironmentUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEnvironmentUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionEnvironmentUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEnvironmentUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEnvironmentUpdateInput is a type for action input parameters
type ActionEnvironmentUpdateInput struct {
	Label string `json:"label"`
	Domain string `json:"domain"`
	CanCreateVps bool `json:"can_create_vps"`
	CanDestroyVps bool `json:"can_destroy_vps"`
	VpsLifetime int64 `json:"vps_lifetime"`
	MaxVpsCount int64 `json:"max_vps_count"`
	UserIpOwnership bool `json:"user_ip_ownership"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionEnvironmentUpdateInput) SetLabel(value string) *ActionEnvironmentUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}
// SetDomain sets parameter Domain to value and selects it for sending
func (in *ActionEnvironmentUpdateInput) SetDomain(value string) *ActionEnvironmentUpdateInput {
	in.Domain = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Domain"] = nil
	return in
}
// SetCanCreateVps sets parameter CanCreateVps to value and selects it for sending
func (in *ActionEnvironmentUpdateInput) SetCanCreateVps(value bool) *ActionEnvironmentUpdateInput {
	in.CanCreateVps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CanCreateVps"] = nil
	return in
}
// SetCanDestroyVps sets parameter CanDestroyVps to value and selects it for sending
func (in *ActionEnvironmentUpdateInput) SetCanDestroyVps(value bool) *ActionEnvironmentUpdateInput {
	in.CanDestroyVps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CanDestroyVps"] = nil
	return in
}
// SetVpsLifetime sets parameter VpsLifetime to value and selects it for sending
func (in *ActionEnvironmentUpdateInput) SetVpsLifetime(value int64) *ActionEnvironmentUpdateInput {
	in.VpsLifetime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["VpsLifetime"] = nil
	return in
}
// SetMaxVpsCount sets parameter MaxVpsCount to value and selects it for sending
func (in *ActionEnvironmentUpdateInput) SetMaxVpsCount(value int64) *ActionEnvironmentUpdateInput {
	in.MaxVpsCount = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxVpsCount"] = nil
	return in
}
// SetUserIpOwnership sets parameter UserIpOwnership to value and selects it for sending
func (in *ActionEnvironmentUpdateInput) SetUserIpOwnership(value bool) *ActionEnvironmentUpdateInput {
	in.UserIpOwnership = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserIpOwnership"] = nil
	return in
}

// SelectParameters sets parameters from ActionEnvironmentUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEnvironmentUpdateInput) SelectParameters(params ...string) *ActionEnvironmentUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEnvironmentUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEnvironmentUpdateRequest is a type for the entire action request
type ActionEnvironmentUpdateRequest struct {
	Environment map[string]interface{} `json:"environment"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionEnvironmentUpdateOutput is a type for action output parameters
type ActionEnvironmentUpdateOutput struct {
	Id int64 `json:"id"`
	Label string `json:"label"`
	Domain string `json:"domain"`
	CanCreateVps bool `json:"can_create_vps"`
	CanDestroyVps bool `json:"can_destroy_vps"`
	VpsLifetime int64 `json:"vps_lifetime"`
	MaxVpsCount int64 `json:"max_vps_count"`
	UserIpOwnership bool `json:"user_ip_ownership"`
}


// Type for action response, including envelope
type ActionEnvironmentUpdateResponse struct {
	Action *ActionEnvironmentUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Environment *ActionEnvironmentUpdateOutput `json:"environment"`
	}

	// Action output without the namespace
	Output *ActionEnvironmentUpdateOutput
}


// Prepare the action for invocation
func (action *ActionEnvironmentUpdate) Prepare() *ActionEnvironmentUpdateInvocation {
	return &ActionEnvironmentUpdateInvocation{
		Action: action,
		Path: "/v5.0/environments/{environment_id}",
	}
}

// ActionEnvironmentUpdateInvocation is used to configure action for invocation
type ActionEnvironmentUpdateInvocation struct {
	// Pointer to the action
	Action *ActionEnvironmentUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEnvironmentUpdateInput
	// Global meta input parameters
	MetaInput *ActionEnvironmentUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEnvironmentUpdateInvocation) SetPathParamInt(param string, value int64) *ActionEnvironmentUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEnvironmentUpdateInvocation) SetPathParamString(param string, value string) *ActionEnvironmentUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEnvironmentUpdateInvocation) NewInput() *ActionEnvironmentUpdateInput {
	inv.Input = &ActionEnvironmentUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEnvironmentUpdateInvocation) SetInput(input *ActionEnvironmentUpdateInput) *ActionEnvironmentUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEnvironmentUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEnvironmentUpdateInvocation) NewMetaInput() *ActionEnvironmentUpdateMetaGlobalInput {
	inv.MetaInput = &ActionEnvironmentUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEnvironmentUpdateInvocation) SetMetaInput(input *ActionEnvironmentUpdateMetaGlobalInput) *ActionEnvironmentUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEnvironmentUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEnvironmentUpdateInvocation) Call() (*ActionEnvironmentUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionEnvironmentUpdateInvocation) callAsBody() (*ActionEnvironmentUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionEnvironmentUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Environment
	}
	return resp, err
}




func (inv *ActionEnvironmentUpdateInvocation) makeAllInputParams() *ActionEnvironmentUpdateRequest {
	return &ActionEnvironmentUpdateRequest{
		Environment: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionEnvironmentUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("Domain") {
			ret["domain"] = inv.Input.Domain
		}
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
		if inv.IsParameterSelected("UserIpOwnership") {
			ret["user_ip_ownership"] = inv.Input.UserIpOwnership
		}
	}

	return ret
}

func (inv *ActionEnvironmentUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
