package client

import ()

// ActionLocationCreate is a type for action Location#Create
type ActionLocationCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionLocationCreate(client *Client) *ActionLocationCreate {
	return &ActionLocationCreate{
		Client: client,
	}
}

// ActionLocationCreateMetaGlobalInput is a type for action global meta input parameters
type ActionLocationCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionLocationCreateMetaGlobalInput) SetIncludes(value string) *ActionLocationCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionLocationCreateMetaGlobalInput) SetNo(value bool) *ActionLocationCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionLocationCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLocationCreateMetaGlobalInput) SelectParameters(params ...string) *ActionLocationCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionLocationCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionLocationCreateInput is a type for action input parameters
type ActionLocationCreateInput struct {
	Description         string `json:"description"`
	Domain              string `json:"domain"`
	Environment         int64  `json:"environment"`
	HasIpv6             bool   `json:"has_ipv6"`
	Label               string `json:"label"`
	RemoteConsoleServer string `json:"remote_console_server"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDescription sets parameter Description to value and selects it for sending
func (in *ActionLocationCreateInput) SetDescription(value string) *ActionLocationCreateInput {
	in.Description = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Description"] = nil
	return in
}

// SetDomain sets parameter Domain to value and selects it for sending
func (in *ActionLocationCreateInput) SetDomain(value string) *ActionLocationCreateInput {
	in.Domain = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Domain"] = nil
	return in
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionLocationCreateInput) SetEnvironment(value int64) *ActionLocationCreateInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEnvironmentNil(false)
	in._selectedParameters["Environment"] = nil
	return in
}

// SetEnvironmentNil sets parameter Environment to nil and selects it for sending
func (in *ActionLocationCreateInput) SetEnvironmentNil(set bool) *ActionLocationCreateInput {
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

// SetHasIpv6 sets parameter HasIpv6 to value and selects it for sending
func (in *ActionLocationCreateInput) SetHasIpv6(value bool) *ActionLocationCreateInput {
	in.HasIpv6 = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HasIpv6"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionLocationCreateInput) SetLabel(value string) *ActionLocationCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SetRemoteConsoleServer sets parameter RemoteConsoleServer to value and selects it for sending
func (in *ActionLocationCreateInput) SetRemoteConsoleServer(value string) *ActionLocationCreateInput {
	in.RemoteConsoleServer = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RemoteConsoleServer"] = nil
	return in
}

// SelectParameters sets parameters from ActionLocationCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLocationCreateInput) SelectParameters(params ...string) *ActionLocationCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionLocationCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionLocationCreateInput) UnselectParameters(params ...string) *ActionLocationCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionLocationCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionLocationCreateRequest is a type for the entire action request
type ActionLocationCreateRequest struct {
	Location map[string]interface{} `json:"location"`
	Meta     map[string]interface{} `json:"_meta"`
}

// ActionLocationCreateOutput is a type for action output parameters
type ActionLocationCreateOutput struct {
	Description         string                       `json:"description"`
	Domain              string                       `json:"domain"`
	Environment         *ActionEnvironmentShowOutput `json:"environment"`
	HasIpv6             bool                         `json:"has_ipv6"`
	Id                  int64                        `json:"id"`
	Label               string                       `json:"label"`
	RemoteConsoleServer string                       `json:"remote_console_server"`
}

// Type for action response, including envelope
type ActionLocationCreateResponse struct {
	Action *ActionLocationCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Location *ActionLocationCreateOutput `json:"location"`
	}

	// Action output without the namespace
	Output *ActionLocationCreateOutput
}

// Prepare the action for invocation
func (action *ActionLocationCreate) Prepare() *ActionLocationCreateInvocation {
	return &ActionLocationCreateInvocation{
		Action: action,
		Path:   "/v7.0/locations",
	}
}

// ActionLocationCreateInvocation is used to configure action for invocation
type ActionLocationCreateInvocation struct {
	// Pointer to the action
	Action *ActionLocationCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionLocationCreateInput
	// Global meta input parameters
	MetaInput *ActionLocationCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionLocationCreateInvocation) NewInput() *ActionLocationCreateInput {
	inv.Input = &ActionLocationCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionLocationCreateInvocation) SetInput(input *ActionLocationCreateInput) *ActionLocationCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionLocationCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionLocationCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionLocationCreateInvocation) NewMetaInput() *ActionLocationCreateMetaGlobalInput {
	inv.MetaInput = &ActionLocationCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionLocationCreateInvocation) SetMetaInput(input *ActionLocationCreateMetaGlobalInput) *ActionLocationCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionLocationCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionLocationCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionLocationCreateInvocation) Call() (*ActionLocationCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionLocationCreateInvocation) callAsBody() (*ActionLocationCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionLocationCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Location
	}
	return resp, err
}

func (inv *ActionLocationCreateInvocation) makeAllInputParams() *ActionLocationCreateRequest {
	return &ActionLocationCreateRequest{
		Location: inv.makeInputParams(),
		Meta:     inv.makeMetaInputParams(),
	}
}

func (inv *ActionLocationCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Description") {
			ret["description"] = inv.Input.Description
		}
		if inv.IsParameterSelected("Domain") {
			ret["domain"] = inv.Input.Domain
		}
		if inv.IsParameterSelected("Environment") {
			if inv.IsParameterNil("Environment") {
				ret["environment"] = nil
			} else {
				ret["environment"] = inv.Input.Environment
			}
		}
		if inv.IsParameterSelected("HasIpv6") {
			ret["has_ipv6"] = inv.Input.HasIpv6
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("RemoteConsoleServer") {
			ret["remote_console_server"] = inv.Input.RemoteConsoleServer
		}
	}

	return ret
}

func (inv *ActionLocationCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
