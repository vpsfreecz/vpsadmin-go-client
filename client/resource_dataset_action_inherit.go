package client

import (
	"strings"
)

// ActionDatasetInherit is a type for action Dataset#Inherit
type ActionDatasetInherit struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetInherit(client *Client) *ActionDatasetInherit {
	return &ActionDatasetInherit{
		Client: client,
	}
}

// ActionDatasetInheritMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetInheritMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetInheritMetaGlobalInput) SetIncludes(value string) *ActionDatasetInheritMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetInheritMetaGlobalInput) SetNo(value bool) *ActionDatasetInheritMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetInheritMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetInheritMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetInheritMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetInheritMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetInheritInput is a type for action input parameters
type ActionDatasetInheritInput struct {
	Property string `json:"property"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetProperty sets parameter Property to value and selects it for sending
func (in *ActionDatasetInheritInput) SetProperty(value string) *ActionDatasetInheritInput {
	in.Property = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Property"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetInheritInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetInheritInput) SelectParameters(params ...string) *ActionDatasetInheritInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDatasetInheritInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDatasetInheritInput) UnselectParameters(params ...string) *ActionDatasetInheritInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDatasetInheritInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetInheritRequest is a type for the entire action request
type ActionDatasetInheritRequest struct {
	Dataset map[string]interface{} `json:"dataset"`
	Meta    map[string]interface{} `json:"_meta"`
}

// ActionDatasetInheritMetaGlobalOutput is a type for global output metadata parameters
type ActionDatasetInheritMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionDatasetInheritResponse struct {
	Action *ActionDatasetInherit `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionDatasetInheritMetaGlobalOutput `json:"_meta"`
	}
}

// Prepare the action for invocation
func (action *ActionDatasetInherit) Prepare() *ActionDatasetInheritInvocation {
	return &ActionDatasetInheritInvocation{
		Action: action,
		Path:   "/v6.0/datasets/{dataset_id}/inherit",
	}
}

// ActionDatasetInheritInvocation is used to configure action for invocation
type ActionDatasetInheritInvocation struct {
	// Pointer to the action
	Action *ActionDatasetInherit

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDatasetInheritInput
	// Global meta input parameters
	MetaInput *ActionDatasetInheritMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetInheritInvocation) SetPathParamInt(param string, value int64) *ActionDatasetInheritInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetInheritInvocation) SetPathParamString(param string, value string) *ActionDatasetInheritInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDatasetInheritInvocation) NewInput() *ActionDatasetInheritInput {
	inv.Input = &ActionDatasetInheritInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDatasetInheritInvocation) SetInput(input *ActionDatasetInheritInput) *ActionDatasetInheritInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDatasetInheritInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDatasetInheritInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetInheritInvocation) NewMetaInput() *ActionDatasetInheritMetaGlobalInput {
	inv.MetaInput = &ActionDatasetInheritMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetInheritInvocation) SetMetaInput(input *ActionDatasetInheritMetaGlobalInput) *ActionDatasetInheritInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetInheritInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDatasetInheritInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetInheritInvocation) Call() (*ActionDatasetInheritResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDatasetInheritInvocation) callAsBody() (*ActionDatasetInheritResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDatasetInheritResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDatasetInheritResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDatasetInheritResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDatasetInheritResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDatasetInheritResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionDatasetInheritResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionDatasetInheritInvocation) makeAllInputParams() *ActionDatasetInheritRequest {
	return &ActionDatasetInheritRequest{
		Dataset: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionDatasetInheritInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Property") {
			ret["property"] = inv.Input.Property
		}
	}

	return ret
}

func (inv *ActionDatasetInheritInvocation) makeMetaInputParams() map[string]interface{} {
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
