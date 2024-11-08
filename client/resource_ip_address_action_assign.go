package client

import (
	"strings"
)

// ActionIpAddressAssign is a type for action Ip_address#Assign
type ActionIpAddressAssign struct {
	// Pointer to client
	Client *Client
}

func NewActionIpAddressAssign(client *Client) *ActionIpAddressAssign {
	return &ActionIpAddressAssign{
		Client: client,
	}
}

// ActionIpAddressAssignMetaGlobalInput is a type for action global meta input parameters
type ActionIpAddressAssignMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpAddressAssignMetaGlobalInput) SetIncludes(value string) *ActionIpAddressAssignMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpAddressAssignMetaGlobalInput) SetNo(value bool) *ActionIpAddressAssignMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpAddressAssignMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpAddressAssignMetaGlobalInput) SelectParameters(params ...string) *ActionIpAddressAssignMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpAddressAssignMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpAddressAssignInput is a type for action input parameters
type ActionIpAddressAssignInput struct {
	NetworkInterface int64 `json:"network_interface"`
	RouteVia         int64 `json:"route_via"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNetworkInterface sets parameter NetworkInterface to value and selects it for sending
func (in *ActionIpAddressAssignInput) SetNetworkInterface(value int64) *ActionIpAddressAssignInput {
	in.NetworkInterface = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNetworkInterfaceNil(false)
	in._selectedParameters["NetworkInterface"] = nil
	return in
}

// SetNetworkInterfaceNil sets parameter NetworkInterface to nil and selects it for sending
func (in *ActionIpAddressAssignInput) SetNetworkInterfaceNil(set bool) *ActionIpAddressAssignInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["NetworkInterface"] = nil
		in.SelectParameters("NetworkInterface")
	} else {
		delete(in._nilParameters, "NetworkInterface")
	}
	return in
}

// SetRouteVia sets parameter RouteVia to value and selects it for sending
func (in *ActionIpAddressAssignInput) SetRouteVia(value int64) *ActionIpAddressAssignInput {
	in.RouteVia = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetRouteViaNil(false)
	in._selectedParameters["RouteVia"] = nil
	return in
}

// SetRouteViaNil sets parameter RouteVia to nil and selects it for sending
func (in *ActionIpAddressAssignInput) SetRouteViaNil(set bool) *ActionIpAddressAssignInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["RouteVia"] = nil
		in.SelectParameters("RouteVia")
	} else {
		delete(in._nilParameters, "RouteVia")
	}
	return in
}

// SelectParameters sets parameters from ActionIpAddressAssignInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpAddressAssignInput) SelectParameters(params ...string) *ActionIpAddressAssignInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIpAddressAssignInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIpAddressAssignInput) UnselectParameters(params ...string) *ActionIpAddressAssignInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIpAddressAssignInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpAddressAssignRequest is a type for the entire action request
type ActionIpAddressAssignRequest struct {
	IpAddress map[string]interface{} `json:"ip_address"`
	Meta      map[string]interface{} `json:"_meta"`
}

// ActionIpAddressAssignOutput is a type for action output parameters
type ActionIpAddressAssignOutput struct {
	Addr               string                            `json:"addr"`
	ChargedEnvironment *ActionEnvironmentShowOutput      `json:"charged_environment"`
	Id                 int64                             `json:"id"`
	Network            *ActionNetworkShowOutput          `json:"network"`
	NetworkInterface   *ActionNetworkInterfaceShowOutput `json:"network_interface"`
	Prefix             int64                             `json:"prefix"`
	RouteVia           *ActionHostIpAddressShowOutput    `json:"route_via"`
	Size               int64                             `json:"size"`
	User               *ActionUserShowOutput             `json:"user"`
}

// ActionIpAddressAssignMetaGlobalOutput is a type for global output metadata parameters
type ActionIpAddressAssignMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionIpAddressAssignResponse struct {
	Action *ActionIpAddressAssign `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpAddress *ActionIpAddressAssignOutput `json:"ip_address"`
		// Global output metadata
		Meta *ActionIpAddressAssignMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionIpAddressAssignOutput
}

// Prepare the action for invocation
func (action *ActionIpAddressAssign) Prepare() *ActionIpAddressAssignInvocation {
	return &ActionIpAddressAssignInvocation{
		Action: action,
		Path:   "/v7.0/ip_addresses/{ip_address_id}/assign",
	}
}

// ActionIpAddressAssignInvocation is used to configure action for invocation
type ActionIpAddressAssignInvocation struct {
	// Pointer to the action
	Action *ActionIpAddressAssign

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpAddressAssignInput
	// Global meta input parameters
	MetaInput *ActionIpAddressAssignMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpAddressAssignInvocation) SetPathParamInt(param string, value int64) *ActionIpAddressAssignInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpAddressAssignInvocation) SetPathParamString(param string, value string) *ActionIpAddressAssignInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpAddressAssignInvocation) NewInput() *ActionIpAddressAssignInput {
	inv.Input = &ActionIpAddressAssignInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpAddressAssignInvocation) SetInput(input *ActionIpAddressAssignInput) *ActionIpAddressAssignInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpAddressAssignInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIpAddressAssignInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpAddressAssignInvocation) NewMetaInput() *ActionIpAddressAssignMetaGlobalInput {
	inv.MetaInput = &ActionIpAddressAssignMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpAddressAssignInvocation) SetMetaInput(input *ActionIpAddressAssignMetaGlobalInput) *ActionIpAddressAssignInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpAddressAssignInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpAddressAssignInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIpAddressAssignInvocation) Call() (*ActionIpAddressAssignResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionIpAddressAssignInvocation) callAsBody() (*ActionIpAddressAssignResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionIpAddressAssignResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IpAddress
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionIpAddressAssignResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionIpAddressAssignResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionIpAddressAssignResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionIpAddressAssignResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)
	input.SetUpdateIn(updateIn)

	pollResp, err := req.Call()

	if err != nil {
		return pollResp, err
	} else if pollResp.Output.Finished {
		return pollResp, nil
	}

	if callback(pollResp.Output) == StopWatching {
		return pollResp, nil
	}

	for {
		req = resp.Action.Client.ActionState.Poll.Prepare()
		req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
		req.SetInput(&ActionActionStatePollInput{
			Timeout:  timeout,
			UpdateIn: updateIn,
			Status:   pollResp.Output.Status,
			Current:  pollResp.Output.Current,
			Total:    pollResp.Output.Total,
		})
		pollResp, err = req.Call()

		if err != nil {
			return pollResp, err
		} else if pollResp.Output.Finished {
			return pollResp, nil
		}

		if callback(pollResp.Output) == StopWatching {
			return pollResp, nil
		}
	}
}

// CancelOperation cancels the current blocking operation
func (resp *ActionIpAddressAssignResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionIpAddressAssignInvocation) makeAllInputParams() *ActionIpAddressAssignRequest {
	return &ActionIpAddressAssignRequest{
		IpAddress: inv.makeInputParams(),
		Meta:      inv.makeMetaInputParams(),
	}
}

func (inv *ActionIpAddressAssignInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("NetworkInterface") {
			if inv.IsParameterNil("NetworkInterface") {
				ret["network_interface"] = nil
			} else {
				ret["network_interface"] = inv.Input.NetworkInterface
			}
		}
		if inv.IsParameterSelected("RouteVia") {
			if inv.IsParameterNil("RouteVia") {
				ret["route_via"] = nil
			} else {
				ret["route_via"] = inv.Input.RouteVia
			}
		}
	}

	return ret
}

func (inv *ActionIpAddressAssignInvocation) makeMetaInputParams() map[string]interface{} {
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
