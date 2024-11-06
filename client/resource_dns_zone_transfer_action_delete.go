package client

import (
	"strings"
)

// ActionDnsZoneTransferDelete is a type for action Dns_zone_transfer#Delete
type ActionDnsZoneTransferDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsZoneTransferDelete(client *Client) *ActionDnsZoneTransferDelete {
	return &ActionDnsZoneTransferDelete{
		Client: client,
	}
}

// ActionDnsZoneTransferDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionDnsZoneTransferDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsZoneTransferDeleteMetaGlobalInput) SetIncludes(value string) *ActionDnsZoneTransferDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsZoneTransferDeleteMetaGlobalInput) SetNo(value bool) *ActionDnsZoneTransferDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsZoneTransferDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsZoneTransferDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionDnsZoneTransferDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsZoneTransferDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsZoneTransferDeleteRequest is a type for the entire action request
type ActionDnsZoneTransferDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// ActionDnsZoneTransferDeleteMetaGlobalOutput is a type for global output metadata parameters
type ActionDnsZoneTransferDeleteMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionDnsZoneTransferDeleteResponse struct {
	Action *ActionDnsZoneTransferDelete `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionDnsZoneTransferDeleteMetaGlobalOutput `json:"_meta"`
	}
}

// Prepare the action for invocation
func (action *ActionDnsZoneTransferDelete) Prepare() *ActionDnsZoneTransferDeleteInvocation {
	return &ActionDnsZoneTransferDeleteInvocation{
		Action: action,
		Path:   "/v7.0/dns_zone_transfers/{dns_zone_transfer_id}",
	}
}

// ActionDnsZoneTransferDeleteInvocation is used to configure action for invocation
type ActionDnsZoneTransferDeleteInvocation struct {
	// Pointer to the action
	Action *ActionDnsZoneTransferDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDnsZoneTransferDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDnsZoneTransferDeleteInvocation) SetPathParamInt(param string, value int64) *ActionDnsZoneTransferDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDnsZoneTransferDeleteInvocation) SetPathParamString(param string, value string) *ActionDnsZoneTransferDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsZoneTransferDeleteInvocation) NewMetaInput() *ActionDnsZoneTransferDeleteMetaGlobalInput {
	inv.MetaInput = &ActionDnsZoneTransferDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsZoneTransferDeleteInvocation) SetMetaInput(input *ActionDnsZoneTransferDeleteMetaGlobalInput) *ActionDnsZoneTransferDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsZoneTransferDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsZoneTransferDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsZoneTransferDeleteInvocation) Call() (*ActionDnsZoneTransferDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDnsZoneTransferDeleteInvocation) callAsBody() (*ActionDnsZoneTransferDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDnsZoneTransferDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDnsZoneTransferDeleteResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDnsZoneTransferDeleteResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDnsZoneTransferDeleteResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDnsZoneTransferDeleteResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionDnsZoneTransferDeleteResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionDnsZoneTransferDeleteInvocation) makeAllInputParams() *ActionDnsZoneTransferDeleteRequest {
	return &ActionDnsZoneTransferDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionDnsZoneTransferDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
