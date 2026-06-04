package client

import ()

// ActionIpAddressCreate is a type for action Ip_address#Create
type ActionIpAddressCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionIpAddressCreate(client *Client) *ActionIpAddressCreate {
	return &ActionIpAddressCreate{
		Client: client,
	}
}

// ActionIpAddressCreateMetaGlobalInput is a type for action global meta input parameters
type ActionIpAddressCreateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpAddressCreateMetaGlobalInput) SetIncludes(value string) *ActionIpAddressCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpAddressCreateMetaGlobalInput) SetNo(value bool) *ActionIpAddressCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpAddressCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpAddressCreateMetaGlobalInput) SelectParameters(params ...string) *ActionIpAddressCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpAddressCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpAddressCreateInput is a type for action input parameters
type ActionIpAddressCreateInput struct {
	Addr             string "json:\"addr\""
	Location         int64  "json:\"location\""
	Network          int64  "json:\"network\""
	NetworkInterface int64  "json:\"network_interface\""
	Prefix           int64  "json:\"prefix\""
	RouteVia         int64  "json:\"route_via\""
	Size             int64  "json:\"size\""
	User             int64  "json:\"user\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAddr sets parameter Addr to value and selects it for sending
func (in *ActionIpAddressCreateInput) SetAddr(value string) *ActionIpAddressCreateInput {
	in.Addr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Addr"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionIpAddressCreateInput) SetLocation(value int64) *ActionIpAddressCreateInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}

// SetNetwork sets parameter Network to value and selects it for sending
func (in *ActionIpAddressCreateInput) SetNetwork(value int64) *ActionIpAddressCreateInput {
	in.Network = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Network"] = nil
	return in
}

// SetNetworkInterface sets parameter NetworkInterface to value and selects it for sending
func (in *ActionIpAddressCreateInput) SetNetworkInterface(value int64) *ActionIpAddressCreateInput {
	in.NetworkInterface = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NetworkInterface"] = nil
	return in
}

// SetPrefix sets parameter Prefix to value and selects it for sending
func (in *ActionIpAddressCreateInput) SetPrefix(value int64) *ActionIpAddressCreateInput {
	in.Prefix = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Prefix"] = nil
	return in
}

// SetRouteVia sets parameter RouteVia to value and selects it for sending
func (in *ActionIpAddressCreateInput) SetRouteVia(value int64) *ActionIpAddressCreateInput {
	in.RouteVia = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RouteVia"] = nil
	return in
}

// SetSize sets parameter Size to value and selects it for sending
func (in *ActionIpAddressCreateInput) SetSize(value int64) *ActionIpAddressCreateInput {
	in.Size = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Size"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionIpAddressCreateInput) SetUser(value int64) *ActionIpAddressCreateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionIpAddressCreateInput) SetUserNil(set bool) *ActionIpAddressCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["User"] = nil
		in.SelectParameters("User")
	} else {
		delete(in._nilParameters, "User")
	}
	return in
}

// SelectParameters sets parameters from ActionIpAddressCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpAddressCreateInput) SelectParameters(params ...string) *ActionIpAddressCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIpAddressCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIpAddressCreateInput) UnselectParameters(params ...string) *ActionIpAddressCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIpAddressCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpAddressCreateRequest is a type for the entire action request
type ActionIpAddressCreateRequest struct {
	IpAddress map[string]interface{} "json:\"ip_address\""
	Meta      map[string]interface{} "json:\"_meta\""
}

// ActionIpAddressCreateOutput is a type for action output parameters
type ActionIpAddressCreateOutput struct {
	Addr               string                            "json:\"addr\""
	ChargedEnvironment *ActionEnvironmentShowOutput      "json:\"charged_environment\""
	Id                 int64                             "json:\"id\""
	Network            *ActionNetworkShowOutput          "json:\"network\""
	NetworkInterface   *ActionNetworkInterfaceShowOutput "json:\"network_interface\""
	Prefix             int64                             "json:\"prefix\""
	RouteVia           *ActionHostIpAddressShowOutput    "json:\"route_via\""
	Size               int64                             "json:\"size\""
	User               *ActionUserShowOutput             "json:\"user\""
}

// Type for action response, including envelope
type ActionIpAddressCreateResponse struct {
	Action *ActionIpAddressCreate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpAddress *ActionIpAddressCreateOutput "json:\"ip_address\""
	}

	// Action output without the namespace
	Output *ActionIpAddressCreateOutput
}

// Prepare the action for invocation
func (action *ActionIpAddressCreate) Prepare() *ActionIpAddressCreateInvocation {
	return &ActionIpAddressCreateInvocation{
		Action: action,
		Path:   "/v7.0/ip_addresses",
	}
}

// ActionIpAddressCreateInvocation is used to configure action for invocation
type ActionIpAddressCreateInvocation struct {
	// Pointer to the action
	Action *ActionIpAddressCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpAddressCreateInput
	// Global meta input parameters
	MetaInput *ActionIpAddressCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpAddressCreateInvocation) NewInput() *ActionIpAddressCreateInput {
	inv.Input = &ActionIpAddressCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpAddressCreateInvocation) SetInput(input *ActionIpAddressCreateInput) *ActionIpAddressCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpAddressCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIpAddressCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpAddressCreateInvocation) NewMetaInput() *ActionIpAddressCreateMetaGlobalInput {
	inv.MetaInput = &ActionIpAddressCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpAddressCreateInvocation) SetMetaInput(input *ActionIpAddressCreateMetaGlobalInput) *ActionIpAddressCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpAddressCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpAddressCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIpAddressCreateInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("Location") {
			if !inv.IsParameterNil("Location") {
				if inv.Input.Location < 0 {
					verr.Add("location", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("Network") {
			if !inv.IsParameterNil("Network") {
				if inv.Input.Network < 0 {
					verr.Add("network", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("NetworkInterface") {
			if !inv.IsParameterNil("NetworkInterface") {
				if inv.Input.NetworkInterface < 0 {
					verr.Add("network_interface", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("RouteVia") {
			if !inv.IsParameterNil("RouteVia") {
				if inv.Input.RouteVia < 0 {
					verr.Add("route_via", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("User") {
			if !inv.IsParameterNil("User") {
				if inv.Input.User < 0 {
					verr.Add("user", "not a valid resource id")
				}
			}
		}
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIpAddressCreateInvocation) Call() (*ActionIpAddressCreateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionIpAddressCreateInvocation) callAsBody() (*ActionIpAddressCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionIpAddressCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IpAddress
	}
	return resp, err
}

func (inv *ActionIpAddressCreateInvocation) makeAllInputParams() *ActionIpAddressCreateRequest {
	return &ActionIpAddressCreateRequest{
		IpAddress: inv.makeInputParams(),
		Meta:      inv.makeMetaInputParams(),
	}
}

func (inv *ActionIpAddressCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Addr") {
			ret["addr"] = inv.Input.Addr
		}
		if inv.IsParameterSelected("Location") {
			ret["location"] = inv.Input.Location
		}
		if inv.IsParameterSelected("Network") {
			ret["network"] = inv.Input.Network
		}
		if inv.IsParameterSelected("NetworkInterface") {
			ret["network_interface"] = inv.Input.NetworkInterface
		}
		if inv.IsParameterSelected("Prefix") {
			ret["prefix"] = inv.Input.Prefix
		}
		if inv.IsParameterSelected("RouteVia") {
			ret["route_via"] = inv.Input.RouteVia
		}
		if inv.IsParameterSelected("Size") {
			ret["size"] = inv.Input.Size
		}
		if inv.IsParameterSelected("User") {
			if inv.IsParameterNil("User") {
				ret["user"] = nil
			} else {
				ret["user"] = inv.Input.User
			}
		}
	}

	return ret
}

func (inv *ActionIpAddressCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
