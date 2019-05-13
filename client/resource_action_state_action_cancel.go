package client

import (
	"strings"
)

// ActionActionStateCancel is a type for action Action_state#Cancel
type ActionActionStateCancel struct {
	// Pointer to client
	Client *Client
}

func NewActionActionStateCancel(client *Client) *ActionActionStateCancel {
	return &ActionActionStateCancel{
		Client: client,
	}
}

// ActionActionStateCancelMetaGlobalInput is a type for action global meta input parameters
type ActionActionStateCancelMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionActionStateCancelMetaGlobalInput) SetNo(value bool) *ActionActionStateCancelMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionActionStateCancelMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionActionStateCancelMetaGlobalInput) SelectParameters(params ...string) *ActionActionStateCancelMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionActionStateCancelMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionActionStateCancelRequest is a type for the entire action request
type ActionActionStateCancelRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}


// ActionActionStateCancelMetaGlobalOutput is a type for global output metadata parameters
type ActionActionStateCancelMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionActionStateCancelResponse struct {
	Action *ActionActionStateCancel `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionActionStateCancelMetaGlobalOutput `json:"_meta"`
	}
}


// Prepare the action for invocation
func (action *ActionActionStateCancel) Prepare() *ActionActionStateCancelInvocation {
	return &ActionActionStateCancelInvocation{
		Action: action,
		Path: "/v5.0/action_states/{action_state_id}/cancel",
	}
}

// ActionActionStateCancelInvocation is used to configure action for invocation
type ActionActionStateCancelInvocation struct {
	// Pointer to the action
	Action *ActionActionStateCancel

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionActionStateCancelMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionActionStateCancelInvocation) SetPathParamInt(param string, value int64) *ActionActionStateCancelInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionActionStateCancelInvocation) SetPathParamString(param string, value string) *ActionActionStateCancelInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionActionStateCancelInvocation) NewMetaInput() *ActionActionStateCancelMetaGlobalInput {
	inv.MetaInput = &ActionActionStateCancelMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionActionStateCancelInvocation) SetMetaInput(input *ActionActionStateCancelMetaGlobalInput) *ActionActionStateCancelInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionActionStateCancelInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionActionStateCancelInvocation) Call() (*ActionActionStateCancelResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionActionStateCancelInvocation) callAsBody() (*ActionActionStateCancelResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionActionStateCancelResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionActionStateCancelResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionActionStateCancelResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionActionStateCancelResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionActionStateCancelResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionActionStateCancelResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionActionStateCancelInvocation) makeAllInputParams() *ActionActionStateCancelRequest {
	return &ActionActionStateCancelRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionActionStateCancelInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
