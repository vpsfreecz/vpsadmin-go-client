package client

import (
	"strings"
)

// ActionVpsDelete is a type for action Vps#Delete
type ActionVpsDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsDelete(client *Client) *ActionVpsDelete {
	return &ActionVpsDelete{
		Client: client,
	}
}

// ActionVpsDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionVpsDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsDeleteMetaGlobalInput) SetIncludes(value string) *ActionVpsDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsDeleteMetaGlobalInput) SetNo(value bool) *ActionVpsDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionVpsDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsDeleteInput is a type for action input parameters
type ActionVpsDeleteInput struct {
	ChangeReason    string `json:"change_reason"`
	ExpirationDate  string `json:"expiration_date"`
	Lazy            bool   `json:"lazy"`
	ObjectState     string `json:"object_state"`
	RemindAfterDate string `json:"remind_after_date"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetChangeReason sets parameter ChangeReason to value and selects it for sending
func (in *ActionVpsDeleteInput) SetChangeReason(value string) *ActionVpsDeleteInput {
	in.ChangeReason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ChangeReason"] = nil
	return in
}

// SetExpirationDate sets parameter ExpirationDate to value and selects it for sending
func (in *ActionVpsDeleteInput) SetExpirationDate(value string) *ActionVpsDeleteInput {
	in.ExpirationDate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ExpirationDate"] = nil
	return in
}

// SetLazy sets parameter Lazy to value and selects it for sending
func (in *ActionVpsDeleteInput) SetLazy(value bool) *ActionVpsDeleteInput {
	in.Lazy = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Lazy"] = nil
	return in
}

// SetObjectState sets parameter ObjectState to value and selects it for sending
func (in *ActionVpsDeleteInput) SetObjectState(value string) *ActionVpsDeleteInput {
	in.ObjectState = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ObjectState"] = nil
	return in
}

// SetRemindAfterDate sets parameter RemindAfterDate to value and selects it for sending
func (in *ActionVpsDeleteInput) SetRemindAfterDate(value string) *ActionVpsDeleteInput {
	in.RemindAfterDate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RemindAfterDate"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsDeleteInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsDeleteInput) SelectParameters(params ...string) *ActionVpsDeleteInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsDeleteInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsDeleteInput) UnselectParameters(params ...string) *ActionVpsDeleteInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsDeleteInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsDeleteRequest is a type for the entire action request
type ActionVpsDeleteRequest struct {
	Vps  map[string]interface{} `json:"vps"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionVpsDeleteMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsDeleteMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsDeleteResponse struct {
	Action *ActionVpsDelete `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionVpsDeleteMetaGlobalOutput `json:"_meta"`
	}
}

// Prepare the action for invocation
func (action *ActionVpsDelete) Prepare() *ActionVpsDeleteInvocation {
	return &ActionVpsDeleteInvocation{
		Action: action,
		Path:   "/v7.0/vpses/{vps_id}",
	}
}

// ActionVpsDeleteInvocation is used to configure action for invocation
type ActionVpsDeleteInvocation struct {
	// Pointer to the action
	Action *ActionVpsDelete

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsDeleteInput
	// Global meta input parameters
	MetaInput *ActionVpsDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsDeleteInvocation) SetPathParamInt(param string, value int64) *ActionVpsDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsDeleteInvocation) SetPathParamString(param string, value string) *ActionVpsDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsDeleteInvocation) NewInput() *ActionVpsDeleteInput {
	inv.Input = &ActionVpsDeleteInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsDeleteInvocation) SetInput(input *ActionVpsDeleteInput) *ActionVpsDeleteInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsDeleteInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsDeleteInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsDeleteInvocation) NewMetaInput() *ActionVpsDeleteMetaGlobalInput {
	inv.MetaInput = &ActionVpsDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsDeleteInvocation) SetMetaInput(input *ActionVpsDeleteMetaGlobalInput) *ActionVpsDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsDeleteInvocation) Call() (*ActionVpsDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionVpsDeleteInvocation) callAsBody() (*ActionVpsDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsDeleteResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsDeleteResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsDeleteResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsDeleteResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsDeleteResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionVpsDeleteInvocation) makeAllInputParams() *ActionVpsDeleteRequest {
	return &ActionVpsDeleteRequest{
		Vps:  inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsDeleteInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("ChangeReason") {
			ret["change_reason"] = inv.Input.ChangeReason
		}
		if inv.IsParameterSelected("ExpirationDate") {
			ret["expiration_date"] = inv.Input.ExpirationDate
		}
		if inv.IsParameterSelected("Lazy") {
			ret["lazy"] = inv.Input.Lazy
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

func (inv *ActionVpsDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
