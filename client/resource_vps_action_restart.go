package client

import (
	"strings"
)

// ActionVpsRestart is a type for action Vps#Restart
type ActionVpsRestart struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsRestart(client *Client) *ActionVpsRestart {
	return &ActionVpsRestart{
		Client: client,
	}
}

// ActionVpsRestartMetaGlobalInput is a type for action global meta input parameters
type ActionVpsRestartMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsRestartMetaGlobalInput) SetIncludes(value string) *ActionVpsRestartMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsRestartMetaGlobalInput) SetNo(value bool) *ActionVpsRestartMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsRestartMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsRestartMetaGlobalInput) SelectParameters(params ...string) *ActionVpsRestartMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsRestartMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsRestartRequest is a type for the entire action request
type ActionVpsRestartRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// ActionVpsRestartMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsRestartMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsRestartResponse struct {
	Action *ActionVpsRestart `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionVpsRestartMetaGlobalOutput `json:"_meta"`
	}
}

// Prepare the action for invocation
func (action *ActionVpsRestart) Prepare() *ActionVpsRestartInvocation {
	return &ActionVpsRestartInvocation{
		Action: action,
		Path:   "/v6.0/vpses/{vps_id}/restart",
	}
}

// ActionVpsRestartInvocation is used to configure action for invocation
type ActionVpsRestartInvocation struct {
	// Pointer to the action
	Action *ActionVpsRestart

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionVpsRestartMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsRestartInvocation) SetPathParamInt(param string, value int64) *ActionVpsRestartInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsRestartInvocation) SetPathParamString(param string, value string) *ActionVpsRestartInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsRestartInvocation) NewMetaInput() *ActionVpsRestartMetaGlobalInput {
	inv.MetaInput = &ActionVpsRestartMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsRestartInvocation) SetMetaInput(input *ActionVpsRestartMetaGlobalInput) *ActionVpsRestartInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsRestartInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsRestartInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsRestartInvocation) Call() (*ActionVpsRestartResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionVpsRestartInvocation) callAsBody() (*ActionVpsRestartResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsRestartResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsRestartResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsRestartResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsRestartResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsRestartResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsRestartResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionVpsRestartInvocation) makeAllInputParams() *ActionVpsRestartRequest {
	return &ActionVpsRestartRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsRestartInvocation) makeMetaInputParams() map[string]interface{} {
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
