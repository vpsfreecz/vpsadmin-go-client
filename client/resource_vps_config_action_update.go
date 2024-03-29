package client

import (
	"strings"
)

// ActionVpsConfigUpdate is a type for action Vps_config#Update
type ActionVpsConfigUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsConfigUpdate(client *Client) *ActionVpsConfigUpdate {
	return &ActionVpsConfigUpdate{
		Client: client,
	}
}

// ActionVpsConfigUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionVpsConfigUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsConfigUpdateMetaGlobalInput) SetIncludes(value string) *ActionVpsConfigUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsConfigUpdateMetaGlobalInput) SetNo(value bool) *ActionVpsConfigUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsConfigUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsConfigUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionVpsConfigUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsConfigUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsConfigUpdateInput is a type for action input parameters
type ActionVpsConfigUpdateInput struct {
	Config string `json:"config"`
	Label  string `json:"label"`
	Name   string `json:"name"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetConfig sets parameter Config to value and selects it for sending
func (in *ActionVpsConfigUpdateInput) SetConfig(value string) *ActionVpsConfigUpdateInput {
	in.Config = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Config"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionVpsConfigUpdateInput) SetLabel(value string) *ActionVpsConfigUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionVpsConfigUpdateInput) SetName(value string) *ActionVpsConfigUpdateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsConfigUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsConfigUpdateInput) SelectParameters(params ...string) *ActionVpsConfigUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsConfigUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsConfigUpdateInput) UnselectParameters(params ...string) *ActionVpsConfigUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsConfigUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsConfigUpdateRequest is a type for the entire action request
type ActionVpsConfigUpdateRequest struct {
	VpsConfig map[string]interface{} `json:"vps_config"`
	Meta      map[string]interface{} `json:"_meta"`
}

// ActionVpsConfigUpdateMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsConfigUpdateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsConfigUpdateResponse struct {
	Action *ActionVpsConfigUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionVpsConfigUpdateMetaGlobalOutput `json:"_meta"`
	}
}

// Prepare the action for invocation
func (action *ActionVpsConfigUpdate) Prepare() *ActionVpsConfigUpdateInvocation {
	return &ActionVpsConfigUpdateInvocation{
		Action: action,
		Path:   "/v6.0/vps_configs/{vps_config_id}",
	}
}

// ActionVpsConfigUpdateInvocation is used to configure action for invocation
type ActionVpsConfigUpdateInvocation struct {
	// Pointer to the action
	Action *ActionVpsConfigUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsConfigUpdateInput
	// Global meta input parameters
	MetaInput *ActionVpsConfigUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsConfigUpdateInvocation) SetPathParamInt(param string, value int64) *ActionVpsConfigUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsConfigUpdateInvocation) SetPathParamString(param string, value string) *ActionVpsConfigUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsConfigUpdateInvocation) NewInput() *ActionVpsConfigUpdateInput {
	inv.Input = &ActionVpsConfigUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsConfigUpdateInvocation) SetInput(input *ActionVpsConfigUpdateInput) *ActionVpsConfigUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsConfigUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsConfigUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsConfigUpdateInvocation) NewMetaInput() *ActionVpsConfigUpdateMetaGlobalInput {
	inv.MetaInput = &ActionVpsConfigUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsConfigUpdateInvocation) SetMetaInput(input *ActionVpsConfigUpdateMetaGlobalInput) *ActionVpsConfigUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsConfigUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsConfigUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsConfigUpdateInvocation) Call() (*ActionVpsConfigUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionVpsConfigUpdateInvocation) callAsBody() (*ActionVpsConfigUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsConfigUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsConfigUpdateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsConfigUpdateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsConfigUpdateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsConfigUpdateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsConfigUpdateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionVpsConfigUpdateInvocation) makeAllInputParams() *ActionVpsConfigUpdateRequest {
	return &ActionVpsConfigUpdateRequest{
		VpsConfig: inv.makeInputParams(),
		Meta:      inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsConfigUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Config") {
			ret["config"] = inv.Input.Config
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("Name") {
			ret["name"] = inv.Input.Name
		}
	}

	return ret
}

func (inv *ActionVpsConfigUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
