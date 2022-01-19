package client

import (
	"strings"
)

// ActionHostIpAddressFree is a type for action Host_ip_address#Free
type ActionHostIpAddressFree struct {
	// Pointer to client
	Client *Client
}

func NewActionHostIpAddressFree(client *Client) *ActionHostIpAddressFree {
	return &ActionHostIpAddressFree{
		Client: client,
	}
}

// ActionHostIpAddressFreeMetaGlobalInput is a type for action global meta input parameters
type ActionHostIpAddressFreeMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionHostIpAddressFreeMetaGlobalInput) SetIncludes(value string) *ActionHostIpAddressFreeMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionHostIpAddressFreeMetaGlobalInput) SetNo(value bool) *ActionHostIpAddressFreeMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionHostIpAddressFreeMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionHostIpAddressFreeMetaGlobalInput) SelectParameters(params ...string) *ActionHostIpAddressFreeMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionHostIpAddressFreeMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionHostIpAddressFreeRequest is a type for the entire action request
type ActionHostIpAddressFreeRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// ActionHostIpAddressFreeOutput is a type for action output parameters
type ActionHostIpAddressFreeOutput struct {
	Addr string `json:"addr"`
	Assigned bool `json:"assigned"`
	Id int64 `json:"id"`
	IpAddress *ActionIpAddressShowOutput `json:"ip_address"`
}

// ActionHostIpAddressFreeMetaGlobalOutput is a type for global output metadata parameters
type ActionHostIpAddressFreeMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionHostIpAddressFreeResponse struct {
	Action *ActionHostIpAddressFree `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		HostIpAddress *ActionHostIpAddressFreeOutput `json:"host_ip_address"`
		// Global output metadata
		Meta *ActionHostIpAddressFreeMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionHostIpAddressFreeOutput
}


// Prepare the action for invocation
func (action *ActionHostIpAddressFree) Prepare() *ActionHostIpAddressFreeInvocation {
	return &ActionHostIpAddressFreeInvocation{
		Action: action,
		Path: "/v6.0/host_ip_addresses/{host_ip_address_id}/free",
	}
}

// ActionHostIpAddressFreeInvocation is used to configure action for invocation
type ActionHostIpAddressFreeInvocation struct {
	// Pointer to the action
	Action *ActionHostIpAddressFree

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionHostIpAddressFreeMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionHostIpAddressFreeInvocation) SetPathParamInt(param string, value int64) *ActionHostIpAddressFreeInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionHostIpAddressFreeInvocation) SetPathParamString(param string, value string) *ActionHostIpAddressFreeInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionHostIpAddressFreeInvocation) NewMetaInput() *ActionHostIpAddressFreeMetaGlobalInput {
	inv.MetaInput = &ActionHostIpAddressFreeMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionHostIpAddressFreeInvocation) SetMetaInput(input *ActionHostIpAddressFreeMetaGlobalInput) *ActionHostIpAddressFreeInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionHostIpAddressFreeInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionHostIpAddressFreeInvocation) Call() (*ActionHostIpAddressFreeResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionHostIpAddressFreeInvocation) callAsBody() (*ActionHostIpAddressFreeResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionHostIpAddressFreeResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.HostIpAddress
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionHostIpAddressFreeResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionHostIpAddressFreeResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionHostIpAddressFreeResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionHostIpAddressFreeResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionHostIpAddressFreeResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionHostIpAddressFreeInvocation) makeAllInputParams() *ActionHostIpAddressFreeRequest {
	return &ActionHostIpAddressFreeRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionHostIpAddressFreeInvocation) makeMetaInputParams() map[string]interface{} {
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
