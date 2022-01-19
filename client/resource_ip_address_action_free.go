package client

import (
	"strings"
)

// ActionIpAddressFree is a type for action Ip_address#Free
type ActionIpAddressFree struct {
	// Pointer to client
	Client *Client
}

func NewActionIpAddressFree(client *Client) *ActionIpAddressFree {
	return &ActionIpAddressFree{
		Client: client,
	}
}

// ActionIpAddressFreeMetaGlobalInput is a type for action global meta input parameters
type ActionIpAddressFreeMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpAddressFreeMetaGlobalInput) SetIncludes(value string) *ActionIpAddressFreeMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpAddressFreeMetaGlobalInput) SetNo(value bool) *ActionIpAddressFreeMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpAddressFreeMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpAddressFreeMetaGlobalInput) SelectParameters(params ...string) *ActionIpAddressFreeMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpAddressFreeMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionIpAddressFreeRequest is a type for the entire action request
type ActionIpAddressFreeRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// ActionIpAddressFreeOutput is a type for action output parameters
type ActionIpAddressFreeOutput struct {
	Addr string `json:"addr"`
	ChargedEnvironment *ActionEnvironmentShowOutput `json:"charged_environment"`
	ClassId int64 `json:"class_id"`
	Id int64 `json:"id"`
	MaxRx int64 `json:"max_rx"`
	MaxTx int64 `json:"max_tx"`
	Network *ActionNetworkShowOutput `json:"network"`
	NetworkInterface *ActionNetworkInterfaceShowOutput `json:"network_interface"`
	Prefix int64 `json:"prefix"`
	RouteVia *ActionHostIpAddressShowOutput `json:"route_via"`
	Size int64 `json:"size"`
	User *ActionUserShowOutput `json:"user"`
}

// ActionIpAddressFreeMetaGlobalOutput is a type for global output metadata parameters
type ActionIpAddressFreeMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionIpAddressFreeResponse struct {
	Action *ActionIpAddressFree `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpAddress *ActionIpAddressFreeOutput `json:"ip_address"`
		// Global output metadata
		Meta *ActionIpAddressFreeMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionIpAddressFreeOutput
}


// Prepare the action for invocation
func (action *ActionIpAddressFree) Prepare() *ActionIpAddressFreeInvocation {
	return &ActionIpAddressFreeInvocation{
		Action: action,
		Path: "/v6.0/ip_addresses/{ip_address_id}/free",
	}
}

// ActionIpAddressFreeInvocation is used to configure action for invocation
type ActionIpAddressFreeInvocation struct {
	// Pointer to the action
	Action *ActionIpAddressFree

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionIpAddressFreeMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpAddressFreeInvocation) SetPathParamInt(param string, value int64) *ActionIpAddressFreeInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpAddressFreeInvocation) SetPathParamString(param string, value string) *ActionIpAddressFreeInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpAddressFreeInvocation) NewMetaInput() *ActionIpAddressFreeMetaGlobalInput {
	inv.MetaInput = &ActionIpAddressFreeMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpAddressFreeInvocation) SetMetaInput(input *ActionIpAddressFreeMetaGlobalInput) *ActionIpAddressFreeInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpAddressFreeInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIpAddressFreeInvocation) Call() (*ActionIpAddressFreeResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionIpAddressFreeInvocation) callAsBody() (*ActionIpAddressFreeResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionIpAddressFreeResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IpAddress
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionIpAddressFreeResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionIpAddressFreeResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionIpAddressFreeResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionIpAddressFreeResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionIpAddressFreeResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionIpAddressFreeInvocation) makeAllInputParams() *ActionIpAddressFreeRequest {
	return &ActionIpAddressFreeRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionIpAddressFreeInvocation) makeMetaInputParams() map[string]interface{} {
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
