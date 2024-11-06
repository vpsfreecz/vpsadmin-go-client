package client

import (
	"strings"
)

// ActionUserDelete is a type for action User#Delete
type ActionUserDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionUserDelete(client *Client) *ActionUserDelete {
	return &ActionUserDelete{
		Client: client,
	}
}

// ActionUserDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionUserDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserDeleteMetaGlobalInput) SetIncludes(value string) *ActionUserDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserDeleteMetaGlobalInput) SetNo(value bool) *ActionUserDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionUserDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserDeleteInput is a type for action input parameters
type ActionUserDeleteInput struct {
	ChangeReason    string `json:"change_reason"`
	ExpirationDate  string `json:"expiration_date"`
	ObjectState     string `json:"object_state"`
	RemindAfterDate string `json:"remind_after_date"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetChangeReason sets parameter ChangeReason to value and selects it for sending
func (in *ActionUserDeleteInput) SetChangeReason(value string) *ActionUserDeleteInput {
	in.ChangeReason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ChangeReason"] = nil
	return in
}

// SetExpirationDate sets parameter ExpirationDate to value and selects it for sending
func (in *ActionUserDeleteInput) SetExpirationDate(value string) *ActionUserDeleteInput {
	in.ExpirationDate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ExpirationDate"] = nil
	return in
}

// SetObjectState sets parameter ObjectState to value and selects it for sending
func (in *ActionUserDeleteInput) SetObjectState(value string) *ActionUserDeleteInput {
	in.ObjectState = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ObjectState"] = nil
	return in
}

// SetRemindAfterDate sets parameter RemindAfterDate to value and selects it for sending
func (in *ActionUserDeleteInput) SetRemindAfterDate(value string) *ActionUserDeleteInput {
	in.RemindAfterDate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RemindAfterDate"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserDeleteInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserDeleteInput) SelectParameters(params ...string) *ActionUserDeleteInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserDeleteInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserDeleteInput) UnselectParameters(params ...string) *ActionUserDeleteInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserDeleteInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserDeleteRequest is a type for the entire action request
type ActionUserDeleteRequest struct {
	User map[string]interface{} `json:"user"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionUserDeleteMetaGlobalOutput is a type for global output metadata parameters
type ActionUserDeleteMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionUserDeleteResponse struct {
	Action *ActionUserDelete `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionUserDeleteMetaGlobalOutput `json:"_meta"`
	}
}

// Prepare the action for invocation
func (action *ActionUserDelete) Prepare() *ActionUserDeleteInvocation {
	return &ActionUserDeleteInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}",
	}
}

// ActionUserDeleteInvocation is used to configure action for invocation
type ActionUserDeleteInvocation struct {
	// Pointer to the action
	Action *ActionUserDelete

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserDeleteInput
	// Global meta input parameters
	MetaInput *ActionUserDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserDeleteInvocation) SetPathParamInt(param string, value int64) *ActionUserDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserDeleteInvocation) SetPathParamString(param string, value string) *ActionUserDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserDeleteInvocation) NewInput() *ActionUserDeleteInput {
	inv.Input = &ActionUserDeleteInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserDeleteInvocation) SetInput(input *ActionUserDeleteInput) *ActionUserDeleteInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserDeleteInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserDeleteInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserDeleteInvocation) NewMetaInput() *ActionUserDeleteMetaGlobalInput {
	inv.MetaInput = &ActionUserDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserDeleteInvocation) SetMetaInput(input *ActionUserDeleteMetaGlobalInput) *ActionUserDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserDeleteInvocation) Call() (*ActionUserDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionUserDeleteInvocation) callAsBody() (*ActionUserDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionUserDeleteResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionUserDeleteResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionUserDeleteResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionUserDeleteResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionUserDeleteResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionUserDeleteInvocation) makeAllInputParams() *ActionUserDeleteRequest {
	return &ActionUserDeleteRequest{
		User: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserDeleteInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("ChangeReason") {
			ret["change_reason"] = inv.Input.ChangeReason
		}
		if inv.IsParameterSelected("ExpirationDate") {
			ret["expiration_date"] = inv.Input.ExpirationDate
		}
		if inv.IsParameterSelected("ObjectState") {
			ret["object_state"] = inv.Input.ObjectState
		}
		if inv.IsParameterSelected("RemindAfterDate") {
			ret["remind_after_date"] = inv.Input.RemindAfterDate
		}
	}

	return ret
}

func (inv *ActionUserDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
