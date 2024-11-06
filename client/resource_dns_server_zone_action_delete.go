package client

import (
	"strings"
)

// ActionDnsServerZoneDelete is a type for action Dns_server_zone#Delete
type ActionDnsServerZoneDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsServerZoneDelete(client *Client) *ActionDnsServerZoneDelete {
	return &ActionDnsServerZoneDelete{
		Client: client,
	}
}

// ActionDnsServerZoneDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionDnsServerZoneDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsServerZoneDeleteMetaGlobalInput) SetIncludes(value string) *ActionDnsServerZoneDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsServerZoneDeleteMetaGlobalInput) SetNo(value bool) *ActionDnsServerZoneDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsServerZoneDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsServerZoneDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionDnsServerZoneDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsServerZoneDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsServerZoneDeleteRequest is a type for the entire action request
type ActionDnsServerZoneDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// ActionDnsServerZoneDeleteMetaGlobalOutput is a type for global output metadata parameters
type ActionDnsServerZoneDeleteMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionDnsServerZoneDeleteResponse struct {
	Action *ActionDnsServerZoneDelete `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionDnsServerZoneDeleteMetaGlobalOutput `json:"_meta"`
	}
}

// Prepare the action for invocation
func (action *ActionDnsServerZoneDelete) Prepare() *ActionDnsServerZoneDeleteInvocation {
	return &ActionDnsServerZoneDeleteInvocation{
		Action: action,
		Path:   "/v7.0/dns_server_zones/{dns_server_zone_id}",
	}
}

// ActionDnsServerZoneDeleteInvocation is used to configure action for invocation
type ActionDnsServerZoneDeleteInvocation struct {
	// Pointer to the action
	Action *ActionDnsServerZoneDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDnsServerZoneDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDnsServerZoneDeleteInvocation) SetPathParamInt(param string, value int64) *ActionDnsServerZoneDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDnsServerZoneDeleteInvocation) SetPathParamString(param string, value string) *ActionDnsServerZoneDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsServerZoneDeleteInvocation) NewMetaInput() *ActionDnsServerZoneDeleteMetaGlobalInput {
	inv.MetaInput = &ActionDnsServerZoneDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsServerZoneDeleteInvocation) SetMetaInput(input *ActionDnsServerZoneDeleteMetaGlobalInput) *ActionDnsServerZoneDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsServerZoneDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsServerZoneDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsServerZoneDeleteInvocation) Call() (*ActionDnsServerZoneDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDnsServerZoneDeleteInvocation) callAsBody() (*ActionDnsServerZoneDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDnsServerZoneDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDnsServerZoneDeleteResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDnsServerZoneDeleteResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDnsServerZoneDeleteResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDnsServerZoneDeleteResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionDnsServerZoneDeleteResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionDnsServerZoneDeleteInvocation) makeAllInputParams() *ActionDnsServerZoneDeleteRequest {
	return &ActionDnsServerZoneDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionDnsServerZoneDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
