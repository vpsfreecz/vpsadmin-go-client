package client

import (
	"net/url"
	"strings"
)

// ActionVpsStop is a type for action Vps#Stop
type ActionVpsStop struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsStop(client *Client) *ActionVpsStop {
	return &ActionVpsStop{
		Client: client,
	}
}

// ActionVpsStopMetaGlobalInput is a type for action global meta input parameters
type ActionVpsStopMetaGlobalInput struct {
	No bool "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsStopMetaGlobalInput) SetNo(value bool) *ActionVpsStopMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsStopMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsStopMetaGlobalInput) SelectParameters(params ...string) *ActionVpsStopMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsStopMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsStopInput is a type for action input parameters
type ActionVpsStopInput struct {
	Force bool "json:\"force\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetForce sets parameter Force to value and selects it for sending
func (in *ActionVpsStopInput) SetForce(value bool) *ActionVpsStopInput {
	in.Force = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Force"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsStopInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsStopInput) SelectParameters(params ...string) *ActionVpsStopInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsStopInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsStopInput) UnselectParameters(params ...string) *ActionVpsStopInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsStopInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsStopRequest is a type for the entire action request
type ActionVpsStopRequest struct {
	Vps  map[string]interface{} "json:\"vps\""
	Meta map[string]interface{} "json:\"_meta\""
}

// ActionVpsStopMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsStopMetaGlobalOutput struct {
	ActionStateId int64 "json:\"action_state_id\""
}

// Type for action response, including envelope
type ActionVpsStopResponse struct {
	Action *ActionVpsStop "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionVpsStopMetaGlobalOutput "json:\"_meta\""
	}
}

// Prepare the action for invocation
func (action *ActionVpsStop) Prepare() *ActionVpsStopInvocation {
	return &ActionVpsStopInvocation{
		Action: action,
		Path:   "/v7.0/vpses/{vps_id}/stop",
	}
}

// ActionVpsStopInvocation is used to configure action for invocation
type ActionVpsStopInvocation struct {
	// Pointer to the action
	Action *ActionVpsStop

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsStopInput
	// Global meta input parameters
	MetaInput *ActionVpsStopMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsStopInvocation) SetPathParamInt(param string, value int64) *ActionVpsStopInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsStopInvocation) SetPathParamString(param string, value string) *ActionVpsStopInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsStopInvocation) NewInput() *ActionVpsStopInput {
	inv.Input = &ActionVpsStopInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsStopInvocation) SetInput(input *ActionVpsStopInput) *ActionVpsStopInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsStopInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsStopInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsStopInvocation) NewMetaInput() *ActionVpsStopMetaGlobalInput {
	inv.MetaInput = &ActionVpsStopMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsStopInvocation) SetMetaInput(input *ActionVpsStopMetaGlobalInput) *ActionVpsStopInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsStopInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsStopInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionVpsStopInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsStopInvocation) Call() (*ActionVpsStopResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionVpsStopInvocation) callAsBody() (*ActionVpsStopResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsStopResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsStopResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsStopResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsStopResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsStopResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsStopResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionVpsStopInvocation) makeAllInputParams() *ActionVpsStopRequest {
	return &ActionVpsStopRequest{
		Vps:  inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsStopInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Force") {
			ret["force"] = inv.Input.Force
		}
	}

	return ret
}

func (inv *ActionVpsStopInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
