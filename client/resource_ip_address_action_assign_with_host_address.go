package client

import (
	"strings"
)

// ActionIpAddressAssignWithHostAddress is a type for action Ip_address#Assign_with_host_address
type ActionIpAddressAssignWithHostAddress struct {
	// Pointer to client
	Client *Client
}

func NewActionIpAddressAssignWithHostAddress(client *Client) *ActionIpAddressAssignWithHostAddress {
	return &ActionIpAddressAssignWithHostAddress{
		Client: client,
	}
}

// ActionIpAddressAssignWithHostAddressMetaGlobalInput is a type for action global meta input parameters
type ActionIpAddressAssignWithHostAddressMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpAddressAssignWithHostAddressMetaGlobalInput) SetIncludes(value string) *ActionIpAddressAssignWithHostAddressMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpAddressAssignWithHostAddressMetaGlobalInput) SetNo(value bool) *ActionIpAddressAssignWithHostAddressMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpAddressAssignWithHostAddressMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpAddressAssignWithHostAddressMetaGlobalInput) SelectParameters(params ...string) *ActionIpAddressAssignWithHostAddressMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpAddressAssignWithHostAddressMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpAddressAssignWithHostAddressInput is a type for action input parameters
type ActionIpAddressAssignWithHostAddressInput struct {
	HostIpAddress    int64 `json:"host_ip_address"`
	NetworkInterface int64 `json:"network_interface"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetHostIpAddress sets parameter HostIpAddress to value and selects it for sending
func (in *ActionIpAddressAssignWithHostAddressInput) SetHostIpAddress(value int64) *ActionIpAddressAssignWithHostAddressInput {
	in.HostIpAddress = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetHostIpAddressNil(false)
	in._selectedParameters["HostIpAddress"] = nil
	return in
}

// SetHostIpAddressNil sets parameter HostIpAddress to nil and selects it for sending
func (in *ActionIpAddressAssignWithHostAddressInput) SetHostIpAddressNil(set bool) *ActionIpAddressAssignWithHostAddressInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["HostIpAddress"] = nil
		in.SelectParameters("HostIpAddress")
	} else {
		delete(in._nilParameters, "HostIpAddress")
	}
	return in
}

// SetNetworkInterface sets parameter NetworkInterface to value and selects it for sending
func (in *ActionIpAddressAssignWithHostAddressInput) SetNetworkInterface(value int64) *ActionIpAddressAssignWithHostAddressInput {
	in.NetworkInterface = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNetworkInterfaceNil(false)
	in._selectedParameters["NetworkInterface"] = nil
	return in
}

// SetNetworkInterfaceNil sets parameter NetworkInterface to nil and selects it for sending
func (in *ActionIpAddressAssignWithHostAddressInput) SetNetworkInterfaceNil(set bool) *ActionIpAddressAssignWithHostAddressInput {
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

// SelectParameters sets parameters from ActionIpAddressAssignWithHostAddressInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpAddressAssignWithHostAddressInput) SelectParameters(params ...string) *ActionIpAddressAssignWithHostAddressInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIpAddressAssignWithHostAddressInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIpAddressAssignWithHostAddressInput) UnselectParameters(params ...string) *ActionIpAddressAssignWithHostAddressInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIpAddressAssignWithHostAddressInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpAddressAssignWithHostAddressRequest is a type for the entire action request
type ActionIpAddressAssignWithHostAddressRequest struct {
	IpAddress map[string]interface{} `json:"ip_address"`
	Meta      map[string]interface{} `json:"_meta"`
}

// ActionIpAddressAssignWithHostAddressOutput is a type for action output parameters
type ActionIpAddressAssignWithHostAddressOutput struct {
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

// ActionIpAddressAssignWithHostAddressMetaGlobalOutput is a type for global output metadata parameters
type ActionIpAddressAssignWithHostAddressMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionIpAddressAssignWithHostAddressResponse struct {
	Action *ActionIpAddressAssignWithHostAddress `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpAddress *ActionIpAddressAssignWithHostAddressOutput `json:"ip_address"`
		// Global output metadata
		Meta *ActionIpAddressAssignWithHostAddressMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionIpAddressAssignWithHostAddressOutput
}

// Prepare the action for invocation
func (action *ActionIpAddressAssignWithHostAddress) Prepare() *ActionIpAddressAssignWithHostAddressInvocation {
	return &ActionIpAddressAssignWithHostAddressInvocation{
		Action: action,
		Path:   "/v7.0/ip_addresses/{ip_address_id}/assign_with_host_address",
	}
}

// ActionIpAddressAssignWithHostAddressInvocation is used to configure action for invocation
type ActionIpAddressAssignWithHostAddressInvocation struct {
	// Pointer to the action
	Action *ActionIpAddressAssignWithHostAddress

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpAddressAssignWithHostAddressInput
	// Global meta input parameters
	MetaInput *ActionIpAddressAssignWithHostAddressMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpAddressAssignWithHostAddressInvocation) SetPathParamInt(param string, value int64) *ActionIpAddressAssignWithHostAddressInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpAddressAssignWithHostAddressInvocation) SetPathParamString(param string, value string) *ActionIpAddressAssignWithHostAddressInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpAddressAssignWithHostAddressInvocation) NewInput() *ActionIpAddressAssignWithHostAddressInput {
	inv.Input = &ActionIpAddressAssignWithHostAddressInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpAddressAssignWithHostAddressInvocation) SetInput(input *ActionIpAddressAssignWithHostAddressInput) *ActionIpAddressAssignWithHostAddressInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpAddressAssignWithHostAddressInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIpAddressAssignWithHostAddressInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpAddressAssignWithHostAddressInvocation) NewMetaInput() *ActionIpAddressAssignWithHostAddressMetaGlobalInput {
	inv.MetaInput = &ActionIpAddressAssignWithHostAddressMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpAddressAssignWithHostAddressInvocation) SetMetaInput(input *ActionIpAddressAssignWithHostAddressMetaGlobalInput) *ActionIpAddressAssignWithHostAddressInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpAddressAssignWithHostAddressInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpAddressAssignWithHostAddressInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIpAddressAssignWithHostAddressInvocation) Call() (*ActionIpAddressAssignWithHostAddressResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionIpAddressAssignWithHostAddressInvocation) callAsBody() (*ActionIpAddressAssignWithHostAddressResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionIpAddressAssignWithHostAddressResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IpAddress
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionIpAddressAssignWithHostAddressResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionIpAddressAssignWithHostAddressResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionIpAddressAssignWithHostAddressResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionIpAddressAssignWithHostAddressResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionIpAddressAssignWithHostAddressResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionIpAddressAssignWithHostAddressInvocation) makeAllInputParams() *ActionIpAddressAssignWithHostAddressRequest {
	return &ActionIpAddressAssignWithHostAddressRequest{
		IpAddress: inv.makeInputParams(),
		Meta:      inv.makeMetaInputParams(),
	}
}

func (inv *ActionIpAddressAssignWithHostAddressInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("HostIpAddress") {
			if inv.IsParameterNil("HostIpAddress") {
				ret["host_ip_address"] = nil
			} else {
				ret["host_ip_address"] = inv.Input.HostIpAddress
			}
		}
		if inv.IsParameterSelected("NetworkInterface") {
			if inv.IsParameterNil("NetworkInterface") {
				ret["network_interface"] = nil
			} else {
				ret["network_interface"] = inv.Input.NetworkInterface
			}
		}
	}

	return ret
}

func (inv *ActionIpAddressAssignWithHostAddressInvocation) makeMetaInputParams() map[string]interface{} {
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
