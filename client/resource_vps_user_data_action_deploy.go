package client

import (
	"strings"
)

// ActionVpsUserDataDeploy is a type for action Vps_user_data#Deploy
type ActionVpsUserDataDeploy struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsUserDataDeploy(client *Client) *ActionVpsUserDataDeploy {
	return &ActionVpsUserDataDeploy{
		Client: client,
	}
}

// ActionVpsUserDataDeployMetaGlobalInput is a type for action global meta input parameters
type ActionVpsUserDataDeployMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsUserDataDeployMetaGlobalInput) SetIncludes(value string) *ActionVpsUserDataDeployMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsUserDataDeployMetaGlobalInput) SetNo(value bool) *ActionVpsUserDataDeployMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsUserDataDeployMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsUserDataDeployMetaGlobalInput) SelectParameters(params ...string) *ActionVpsUserDataDeployMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsUserDataDeployMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsUserDataDeployInput is a type for action input parameters
type ActionVpsUserDataDeployInput struct {
	Vps int64 `json:"vps"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionVpsUserDataDeployInput) SetVps(value int64) *ActionVpsUserDataDeployInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetVpsNil(false)
	in._selectedParameters["Vps"] = nil
	return in
}

// SetVpsNil sets parameter Vps to nil and selects it for sending
func (in *ActionVpsUserDataDeployInput) SetVpsNil(set bool) *ActionVpsUserDataDeployInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Vps"] = nil
		in.SelectParameters("Vps")
	} else {
		delete(in._nilParameters, "Vps")
	}
	return in
}

// SelectParameters sets parameters from ActionVpsUserDataDeployInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsUserDataDeployInput) SelectParameters(params ...string) *ActionVpsUserDataDeployInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsUserDataDeployInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsUserDataDeployInput) UnselectParameters(params ...string) *ActionVpsUserDataDeployInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsUserDataDeployInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsUserDataDeployRequest is a type for the entire action request
type ActionVpsUserDataDeployRequest struct {
	VpsUserData map[string]interface{} `json:"vps_user_data"`
	Meta        map[string]interface{} `json:"_meta"`
}

// ActionVpsUserDataDeployMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsUserDataDeployMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsUserDataDeployResponse struct {
	Action *ActionVpsUserDataDeploy `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionVpsUserDataDeployMetaGlobalOutput `json:"_meta"`
	}
}

// Prepare the action for invocation
func (action *ActionVpsUserDataDeploy) Prepare() *ActionVpsUserDataDeployInvocation {
	return &ActionVpsUserDataDeployInvocation{
		Action: action,
		Path:   "/v7.0/vps_user_data/{vps_user_data_id}/deploy",
	}
}

// ActionVpsUserDataDeployInvocation is used to configure action for invocation
type ActionVpsUserDataDeployInvocation struct {
	// Pointer to the action
	Action *ActionVpsUserDataDeploy

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsUserDataDeployInput
	// Global meta input parameters
	MetaInput *ActionVpsUserDataDeployMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsUserDataDeployInvocation) SetPathParamInt(param string, value int64) *ActionVpsUserDataDeployInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsUserDataDeployInvocation) SetPathParamString(param string, value string) *ActionVpsUserDataDeployInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsUserDataDeployInvocation) NewInput() *ActionVpsUserDataDeployInput {
	inv.Input = &ActionVpsUserDataDeployInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsUserDataDeployInvocation) SetInput(input *ActionVpsUserDataDeployInput) *ActionVpsUserDataDeployInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsUserDataDeployInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsUserDataDeployInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsUserDataDeployInvocation) NewMetaInput() *ActionVpsUserDataDeployMetaGlobalInput {
	inv.MetaInput = &ActionVpsUserDataDeployMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsUserDataDeployInvocation) SetMetaInput(input *ActionVpsUserDataDeployMetaGlobalInput) *ActionVpsUserDataDeployInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsUserDataDeployInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsUserDataDeployInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsUserDataDeployInvocation) Call() (*ActionVpsUserDataDeployResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionVpsUserDataDeployInvocation) callAsBody() (*ActionVpsUserDataDeployResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsUserDataDeployResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsUserDataDeployResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsUserDataDeployResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsUserDataDeployResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsUserDataDeployResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsUserDataDeployResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionVpsUserDataDeployInvocation) makeAllInputParams() *ActionVpsUserDataDeployRequest {
	return &ActionVpsUserDataDeployRequest{
		VpsUserData: inv.makeInputParams(),
		Meta:        inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsUserDataDeployInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Vps") {
			if inv.IsParameterNil("Vps") {
				ret["vps"] = nil
			} else {
				ret["vps"] = inv.Input.Vps
			}
		}
	}

	return ret
}

func (inv *ActionVpsUserDataDeployInvocation) makeMetaInputParams() map[string]interface{} {
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
