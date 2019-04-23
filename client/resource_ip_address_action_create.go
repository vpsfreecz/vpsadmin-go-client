package client

import (
)

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
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpAddressCreateMetaGlobalInput) SetIncludes(value string) *ActionIpAddressCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
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
	NetworkInterface int64 `json:"network_interface"`
	Network int64 `json:"network"`
	User int64 `json:"user"`
	Addr string `json:"addr"`
	Prefix int64 `json:"prefix"`
	Size int64 `json:"size"`
	RouteVia int64 `json:"route_via"`
	MaxTx int64 `json:"max_tx"`
	MaxRx int64 `json:"max_rx"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
// SetNetwork sets parameter Network to value and selects it for sending
func (in *ActionIpAddressCreateInput) SetNetwork(value int64) *ActionIpAddressCreateInput {
	in.Network = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Network"] = nil
	return in
}
// SetUser sets parameter User to value and selects it for sending
func (in *ActionIpAddressCreateInput) SetUser(value int64) *ActionIpAddressCreateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
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
// SetPrefix sets parameter Prefix to value and selects it for sending
func (in *ActionIpAddressCreateInput) SetPrefix(value int64) *ActionIpAddressCreateInput {
	in.Prefix = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Prefix"] = nil
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
// SetRouteVia sets parameter RouteVia to value and selects it for sending
func (in *ActionIpAddressCreateInput) SetRouteVia(value int64) *ActionIpAddressCreateInput {
	in.RouteVia = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RouteVia"] = nil
	return in
}
// SetMaxTx sets parameter MaxTx to value and selects it for sending
func (in *ActionIpAddressCreateInput) SetMaxTx(value int64) *ActionIpAddressCreateInput {
	in.MaxTx = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxTx"] = nil
	return in
}
// SetMaxRx sets parameter MaxRx to value and selects it for sending
func (in *ActionIpAddressCreateInput) SetMaxRx(value int64) *ActionIpAddressCreateInput {
	in.MaxRx = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxRx"] = nil
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

func (in *ActionIpAddressCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpAddressCreateRequest is a type for the entire action request
type ActionIpAddressCreateRequest struct {
	IpAddress map[string]interface{} `json:"ip_address"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionIpAddressCreateOutput is a type for action output parameters
type ActionIpAddressCreateOutput struct {
	Id int64 `json:"id"`
	NetworkInterface *ActionNetworkInterfaceShowOutput `json:"network_interface"`
	Network *ActionNetworkShowOutput `json:"network"`
	User *ActionUserShowOutput `json:"user"`
	Addr string `json:"addr"`
	Prefix int64 `json:"prefix"`
	Size int64 `json:"size"`
	RouteVia *ActionHostIpAddressShowOutput `json:"route_via"`
	MaxTx int64 `json:"max_tx"`
	MaxRx int64 `json:"max_rx"`
	ClassId int64 `json:"class_id"`
}


// Type for action response, including envelope
type ActionIpAddressCreateResponse struct {
	Action *ActionIpAddressCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpAddress *ActionIpAddressCreateOutput `json:"ip_address"`
	}

	// Action output without the namespace
	Output *ActionIpAddressCreateOutput
}


// Prepare the action for invocation
func (action *ActionIpAddressCreate) Prepare() *ActionIpAddressCreateInvocation {
	return &ActionIpAddressCreateInvocation{
		Action: action,
		Path: "/v5.0/ip_addresses",
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

// Call() invokes the action and returns a response from the API server
func (inv *ActionIpAddressCreateInvocation) Call() (*ActionIpAddressCreateResponse, error) {
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
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionIpAddressCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("NetworkInterface") {
			ret["network_interface"] = inv.Input.NetworkInterface
		}
		if inv.IsParameterSelected("Network") {
			ret["network"] = inv.Input.Network
		}
		if inv.IsParameterSelected("User") {
			ret["user"] = inv.Input.User
		}
		if inv.IsParameterSelected("Addr") {
			ret["addr"] = inv.Input.Addr
		}
		if inv.IsParameterSelected("Prefix") {
			ret["prefix"] = inv.Input.Prefix
		}
		if inv.IsParameterSelected("Size") {
			ret["size"] = inv.Input.Size
		}
		if inv.IsParameterSelected("RouteVia") {
			ret["route_via"] = inv.Input.RouteVia
		}
		if inv.IsParameterSelected("MaxTx") {
			ret["max_tx"] = inv.Input.MaxTx
		}
		if inv.IsParameterSelected("MaxRx") {
			ret["max_rx"] = inv.Input.MaxRx
		}
	}

	return ret
}

func (inv *ActionIpAddressCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
