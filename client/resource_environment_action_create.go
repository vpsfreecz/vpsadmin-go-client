package client

import (
)

// ActionEnvironmentCreate is a type for action Environment#Create
type ActionEnvironmentCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionEnvironmentCreate(client *Client) *ActionEnvironmentCreate {
	return &ActionEnvironmentCreate{
		Client: client,
	}
}

// ActionEnvironmentCreateMetaGlobalInput is a type for action global meta input parameters
type ActionEnvironmentCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEnvironmentCreateMetaGlobalInput) SetNo(value bool) *ActionEnvironmentCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEnvironmentCreateMetaGlobalInput) SetIncludes(value string) *ActionEnvironmentCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionEnvironmentCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEnvironmentCreateMetaGlobalInput) SelectParameters(params ...string) *ActionEnvironmentCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEnvironmentCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEnvironmentCreateInput is a type for action input parameters
type ActionEnvironmentCreateInput struct {
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
func (in *ActionEnvironmentCreateInput) SetLabel(value string) *ActionEnvironmentCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}
// SetDomain sets parameter Domain to value and selects it for sending
func (in *ActionEnvironmentCreateInput) SetDomain(value string) *ActionEnvironmentCreateInput {
	in.Domain = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Domain"] = nil
	return in
}
// SetCanCreateVps sets parameter CanCreateVps to value and selects it for sending
func (in *ActionEnvironmentCreateInput) SetCanCreateVps(value bool) *ActionEnvironmentCreateInput {
	in.CanCreateVps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CanCreateVps"] = nil
	return in
}
// SetCanDestroyVps sets parameter CanDestroyVps to value and selects it for sending
func (in *ActionEnvironmentCreateInput) SetCanDestroyVps(value bool) *ActionEnvironmentCreateInput {
	in.CanDestroyVps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CanDestroyVps"] = nil
	return in
}
// SetVpsLifetime sets parameter VpsLifetime to value and selects it for sending
func (in *ActionEnvironmentCreateInput) SetVpsLifetime(value int64) *ActionEnvironmentCreateInput {
	in.VpsLifetime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["VpsLifetime"] = nil
	return in
}
// SetMaxVpsCount sets parameter MaxVpsCount to value and selects it for sending
func (in *ActionEnvironmentCreateInput) SetMaxVpsCount(value int64) *ActionEnvironmentCreateInput {
	in.MaxVpsCount = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxVpsCount"] = nil
	return in
}
// SetUserIpOwnership sets parameter UserIpOwnership to value and selects it for sending
func (in *ActionEnvironmentCreateInput) SetUserIpOwnership(value bool) *ActionEnvironmentCreateInput {
	in.UserIpOwnership = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserIpOwnership"] = nil
	return in
}

// SelectParameters sets parameters from ActionEnvironmentCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEnvironmentCreateInput) SelectParameters(params ...string) *ActionEnvironmentCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEnvironmentCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEnvironmentCreateRequest is a type for the entire action request
type ActionEnvironmentCreateRequest struct {
	Environment map[string]interface{} `json:"environment"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionEnvironmentCreateOutput is a type for action output parameters
type ActionEnvironmentCreateOutput struct {
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
type ActionEnvironmentCreateResponse struct {
	Action *ActionEnvironmentCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Environment *ActionEnvironmentCreateOutput `json:"environment"`
	}

	// Action output without the namespace
	Output *ActionEnvironmentCreateOutput
}


// Prepare the action for invocation
func (action *ActionEnvironmentCreate) Prepare() *ActionEnvironmentCreateInvocation {
	return &ActionEnvironmentCreateInvocation{
		Action: action,
		Path: "/v5.0/environments",
	}
}

// ActionEnvironmentCreateInvocation is used to configure action for invocation
type ActionEnvironmentCreateInvocation struct {
	// Pointer to the action
	Action *ActionEnvironmentCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEnvironmentCreateInput
	// Global meta input parameters
	MetaInput *ActionEnvironmentCreateMetaGlobalInput
}


// SetInput provides input parameters to send to the API
func (inv *ActionEnvironmentCreateInvocation) SetInput(input *ActionEnvironmentCreateInput) *ActionEnvironmentCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEnvironmentCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEnvironmentCreateInvocation) SetMetaInput(input *ActionEnvironmentCreateMetaGlobalInput) *ActionEnvironmentCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEnvironmentCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEnvironmentCreateInvocation) Call() (*ActionEnvironmentCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionEnvironmentCreateInvocation) callAsBody() (*ActionEnvironmentCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionEnvironmentCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Environment
	}
	return resp, err
}




func (inv *ActionEnvironmentCreateInvocation) makeAllInputParams() *ActionEnvironmentCreateRequest {
	return &ActionEnvironmentCreateRequest{
		Environment: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionEnvironmentCreateInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionEnvironmentCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
