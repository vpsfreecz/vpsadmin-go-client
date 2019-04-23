package client

import (
	"strings"
)

// ActionVpsMountDelete is a type for action Vps.Mount#Delete
type ActionVpsMountDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsMountDelete(client *Client) *ActionVpsMountDelete {
	return &ActionVpsMountDelete{
		Client: client,
	}
}

// ActionVpsMountDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionVpsMountDeleteMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsMountDeleteMetaGlobalInput) SetNo(value bool) *ActionVpsMountDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsMountDeleteMetaGlobalInput) SetIncludes(value string) *ActionVpsMountDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsMountDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsMountDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionVpsMountDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsMountDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionVpsMountDeleteRequest is a type for the entire action request
type ActionVpsMountDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}


// ActionVpsMountDeleteMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsMountDeleteMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsMountDeleteResponse struct {
	Action *ActionVpsMountDelete `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionVpsMountDeleteMetaGlobalOutput `json:"_meta"`
	}
}


// Prepare the action for invocation
func (action *ActionVpsMountDelete) Prepare() *ActionVpsMountDeleteInvocation {
	return &ActionVpsMountDeleteInvocation{
		Action: action,
		Path: "/v5.0/vpses/:vps_id/mounts/:mount_id",
	}
}

// ActionVpsMountDeleteInvocation is used to configure action for invocation
type ActionVpsMountDeleteInvocation struct {
	// Pointer to the action
	Action *ActionVpsMountDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionVpsMountDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsMountDeleteInvocation) SetPathParamInt(param string, value int64) *ActionVpsMountDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsMountDeleteInvocation) SetPathParamString(param string, value string) *ActionVpsMountDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsMountDeleteInvocation) SetMetaInput(input *ActionVpsMountDeleteMetaGlobalInput) *ActionVpsMountDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsMountDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsMountDeleteInvocation) Call() (*ActionVpsMountDeleteResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionVpsMountDeleteInvocation) callAsBody() (*ActionVpsMountDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsMountDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsMountDeleteResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsMountDeleteResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsMountDeleteResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsMountDeleteResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsMountDeleteResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionVpsMountDeleteInvocation) makeAllInputParams() *ActionVpsMountDeleteRequest {
	return &ActionVpsMountDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionVpsMountDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
