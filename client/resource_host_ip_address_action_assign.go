package client

import (
	"strings"
)

// ActionHostIpAddressAssign is a type for action Host_ip_address#Assign
type ActionHostIpAddressAssign struct {
	// Pointer to client
	Client *Client
}

func NewActionHostIpAddressAssign(client *Client) *ActionHostIpAddressAssign {
	return &ActionHostIpAddressAssign{
		Client: client,
	}
}

// ActionHostIpAddressAssignMetaGlobalInput is a type for action global meta input parameters
type ActionHostIpAddressAssignMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionHostIpAddressAssignMetaGlobalInput) SetNo(value bool) *ActionHostIpAddressAssignMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionHostIpAddressAssignMetaGlobalInput) SetIncludes(value string) *ActionHostIpAddressAssignMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionHostIpAddressAssignMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionHostIpAddressAssignMetaGlobalInput) SelectParameters(params ...string) *ActionHostIpAddressAssignMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionHostIpAddressAssignMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionHostIpAddressAssignRequest is a type for the entire action request
type ActionHostIpAddressAssignRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// ActionHostIpAddressAssignOutput is a type for action output parameters
type ActionHostIpAddressAssignOutput struct {
	Id int64 `json:"id"`
	IpAddress *ActionIpAddressShowOutput `json:"ip_address"`
	Addr string `json:"addr"`
	Assigned bool `json:"assigned"`
}

// ActionHostIpAddressAssignMetaGlobalOutput is a type for global output metadata parameters
type ActionHostIpAddressAssignMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionHostIpAddressAssignResponse struct {
	Action *ActionHostIpAddressAssign `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		HostIpAddress *ActionHostIpAddressAssignOutput `json:"host_ip_address"`
		// Global output metadata
		Meta *ActionHostIpAddressAssignMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionHostIpAddressAssignOutput
}


// Prepare the action for invocation
func (action *ActionHostIpAddressAssign) Prepare() *ActionHostIpAddressAssignInvocation {
	return &ActionHostIpAddressAssignInvocation{
		Action: action,
		Path: "/v5.0/host_ip_addresses/:host_ip_address_id/assign",
	}
}

// ActionHostIpAddressAssignInvocation is used to configure action for invocation
type ActionHostIpAddressAssignInvocation struct {
	// Pointer to the action
	Action *ActionHostIpAddressAssign

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionHostIpAddressAssignMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionHostIpAddressAssignInvocation) SetPathParamInt(param string, value int64) *ActionHostIpAddressAssignInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionHostIpAddressAssignInvocation) SetPathParamString(param string, value string) *ActionHostIpAddressAssignInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionHostIpAddressAssignInvocation) NewMetaInput() *ActionHostIpAddressAssignMetaGlobalInput {
	inv.MetaInput = &ActionHostIpAddressAssignMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionHostIpAddressAssignInvocation) SetMetaInput(input *ActionHostIpAddressAssignMetaGlobalInput) *ActionHostIpAddressAssignInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionHostIpAddressAssignInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionHostIpAddressAssignInvocation) Call() (*ActionHostIpAddressAssignResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionHostIpAddressAssignInvocation) callAsBody() (*ActionHostIpAddressAssignResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionHostIpAddressAssignResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.HostIpAddress
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionHostIpAddressAssignResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionHostIpAddressAssignResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionHostIpAddressAssignResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionHostIpAddressAssignResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionHostIpAddressAssignResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionHostIpAddressAssignInvocation) makeAllInputParams() *ActionHostIpAddressAssignRequest {
	return &ActionHostIpAddressAssignRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionHostIpAddressAssignInvocation) makeMetaInputParams() map[string]interface{} {
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
