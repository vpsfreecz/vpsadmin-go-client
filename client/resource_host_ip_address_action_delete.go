package client

import (
	"strings"
)

// ActionHostIpAddressDelete is a type for action Host_ip_address#Delete
type ActionHostIpAddressDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionHostIpAddressDelete(client *Client) *ActionHostIpAddressDelete {
	return &ActionHostIpAddressDelete{
		Client: client,
	}
}

// ActionHostIpAddressDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionHostIpAddressDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionHostIpAddressDeleteMetaGlobalInput) SetIncludes(value string) *ActionHostIpAddressDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionHostIpAddressDeleteMetaGlobalInput) SetNo(value bool) *ActionHostIpAddressDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionHostIpAddressDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionHostIpAddressDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionHostIpAddressDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionHostIpAddressDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionHostIpAddressDeleteRequest is a type for the entire action request
type ActionHostIpAddressDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// ActionHostIpAddressDeleteMetaGlobalOutput is a type for global output metadata parameters
type ActionHostIpAddressDeleteMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionHostIpAddressDeleteResponse struct {
	Action *ActionHostIpAddressDelete `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionHostIpAddressDeleteMetaGlobalOutput `json:"_meta"`
	}
}

// Prepare the action for invocation
func (action *ActionHostIpAddressDelete) Prepare() *ActionHostIpAddressDeleteInvocation {
	return &ActionHostIpAddressDeleteInvocation{
		Action: action,
		Path:   "/v7.0/host_ip_addresses/{host_ip_address_id}",
	}
}

// ActionHostIpAddressDeleteInvocation is used to configure action for invocation
type ActionHostIpAddressDeleteInvocation struct {
	// Pointer to the action
	Action *ActionHostIpAddressDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionHostIpAddressDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionHostIpAddressDeleteInvocation) SetPathParamInt(param string, value int64) *ActionHostIpAddressDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionHostIpAddressDeleteInvocation) SetPathParamString(param string, value string) *ActionHostIpAddressDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionHostIpAddressDeleteInvocation) NewMetaInput() *ActionHostIpAddressDeleteMetaGlobalInput {
	inv.MetaInput = &ActionHostIpAddressDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionHostIpAddressDeleteInvocation) SetMetaInput(input *ActionHostIpAddressDeleteMetaGlobalInput) *ActionHostIpAddressDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionHostIpAddressDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionHostIpAddressDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionHostIpAddressDeleteInvocation) Call() (*ActionHostIpAddressDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionHostIpAddressDeleteInvocation) callAsBody() (*ActionHostIpAddressDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionHostIpAddressDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionHostIpAddressDeleteResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionHostIpAddressDeleteResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionHostIpAddressDeleteResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionHostIpAddressDeleteResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionHostIpAddressDeleteResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionHostIpAddressDeleteInvocation) makeAllInputParams() *ActionHostIpAddressDeleteRequest {
	return &ActionHostIpAddressDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionHostIpAddressDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
