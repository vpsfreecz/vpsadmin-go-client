package client

import (
	"strings"
)

// ActionVpsConfigReplace is a type for action Vps.Config#Replace
type ActionVpsConfigReplace struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsConfigReplace(client *Client) *ActionVpsConfigReplace {
	return &ActionVpsConfigReplace{
		Client: client,
	}
}

// ActionVpsConfigReplaceMetaGlobalInput is a type for action global meta input parameters
type ActionVpsConfigReplaceMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsConfigReplaceMetaGlobalInput) SetIncludes(value string) *ActionVpsConfigReplaceMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsConfigReplaceMetaGlobalInput) SetNo(value bool) *ActionVpsConfigReplaceMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsConfigReplaceMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsConfigReplaceMetaGlobalInput) SelectParameters(params ...string) *ActionVpsConfigReplaceMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsConfigReplaceMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsConfigReplaceInput is a type for action input parameters
type ActionVpsConfigReplaceInput struct {
	VpsConfig int64 `json:"vps_config"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetVpsConfig sets parameter VpsConfig to value and selects it for sending
func (in *ActionVpsConfigReplaceInput) SetVpsConfig(value int64) *ActionVpsConfigReplaceInput {
	in.VpsConfig = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["VpsConfig"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsConfigReplaceInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsConfigReplaceInput) SelectParameters(params ...string) *ActionVpsConfigReplaceInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsConfigReplaceInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsConfigReplaceRequest is a type for the entire action request
type ActionVpsConfigReplaceRequest struct {
	Configs map[string]interface{} `json:"configs"`
	Meta    map[string]interface{} `json:"_meta"`
}

// ActionVpsConfigReplaceMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsConfigReplaceMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsConfigReplaceResponse struct {
	Action *ActionVpsConfigReplace `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionVpsConfigReplaceMetaGlobalOutput `json:"_meta"`
	}
}

// Prepare the action for invocation
func (action *ActionVpsConfigReplace) Prepare() *ActionVpsConfigReplaceInvocation {
	return &ActionVpsConfigReplaceInvocation{
		Action: action,
		Path:   "/v6.0/vpses/{vps_id}/configs/replace",
	}
}

// ActionVpsConfigReplaceInvocation is used to configure action for invocation
type ActionVpsConfigReplaceInvocation struct {
	// Pointer to the action
	Action *ActionVpsConfigReplace

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsConfigReplaceInput
	// Global meta input parameters
	MetaInput *ActionVpsConfigReplaceMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsConfigReplaceInvocation) SetPathParamInt(param string, value int64) *ActionVpsConfigReplaceInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsConfigReplaceInvocation) SetPathParamString(param string, value string) *ActionVpsConfigReplaceInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsConfigReplaceInvocation) NewInput() *ActionVpsConfigReplaceInput {
	inv.Input = &ActionVpsConfigReplaceInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsConfigReplaceInvocation) SetInput(input *ActionVpsConfigReplaceInput) *ActionVpsConfigReplaceInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsConfigReplaceInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsConfigReplaceInvocation) NewMetaInput() *ActionVpsConfigReplaceMetaGlobalInput {
	inv.MetaInput = &ActionVpsConfigReplaceMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsConfigReplaceInvocation) SetMetaInput(input *ActionVpsConfigReplaceMetaGlobalInput) *ActionVpsConfigReplaceInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsConfigReplaceInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsConfigReplaceInvocation) Call() (*ActionVpsConfigReplaceResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionVpsConfigReplaceInvocation) callAsBody() (*ActionVpsConfigReplaceResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsConfigReplaceResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsConfigReplaceResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsConfigReplaceResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsConfigReplaceResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsConfigReplaceResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsConfigReplaceResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionVpsConfigReplaceInvocation) makeAllInputParams() *ActionVpsConfigReplaceRequest {
	return &ActionVpsConfigReplaceRequest{
		Configs: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsConfigReplaceInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("VpsConfig") {
			ret["vps_config"] = inv.Input.VpsConfig
		}
	}

	return ret
}

func (inv *ActionVpsConfigReplaceInvocation) makeMetaInputParams() map[string]interface{} {
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
