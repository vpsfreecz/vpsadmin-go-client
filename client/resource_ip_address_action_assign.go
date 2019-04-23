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
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpAddressAssignMetaGlobalInput) SetIncludes(value string) *ActionIpAddressAssignMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
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
	RouteVia int64 `json:"route_via"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNetworkInterface sets parameter NetworkInterface to value and selects it for sending
func (in *ActionIpAddressAssignInput) SetNetworkInterface(value int64) *ActionIpAddressAssignInput {
	in.NetworkInterface = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NetworkInterface"] = nil
	return in
}
// SetRouteVia sets parameter RouteVia to value and selects it for sending
func (in *ActionIpAddressAssignInput) SetRouteVia(value int64) *ActionIpAddressAssignInput {
	in.RouteVia = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RouteVia"] = nil
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

func (in *ActionIpAddressAssignInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpAddressAssignRequest is a type for the entire action request
type ActionIpAddressAssignRequest struct {
	IpAddress map[string]interface{} `json:"ip_address"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionIpAddressAssignOutput is a type for action output parameters
type ActionIpAddressAssignOutput struct {
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
		Path: "/v5.0/ip_addresses/:ip_address_id/assign",
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
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
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
	req.SetInput(&ActionActionStatePollInput{
		Timeout: timeout,
	})
	req.Input.SelectParameters("Timeout")
	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionIpAddressAssignResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	req.SetInput(&ActionActionStatePollInput{
		Timeout: timeout,
		UpdateIn: updateIn,
	})
	req.Input.SelectParameters("Timeout", "UpdateIn")
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
			Timeout: timeout,
			UpdateIn: updateIn,
			Status: pollResp.Output.Status,
			Current: pollResp.Output.Current,
			Total: pollResp.Output.Total,
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
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionIpAddressAssignInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("NetworkInterface") {
			ret["network_interface"] = inv.Input.NetworkInterface
		}
		if inv.IsParameterSelected("RouteVia") {
			ret["route_via"] = inv.Input.RouteVia
		}
	}

	return ret
}

func (inv *ActionIpAddressAssignInvocation) makeMetaInputParams() map[string]interface{} {
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
