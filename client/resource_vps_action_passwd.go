package client

import (
	"strings"
)

// ActionVpsPasswd is a type for action Vps#Passwd
type ActionVpsPasswd struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsPasswd(client *Client) *ActionVpsPasswd {
	return &ActionVpsPasswd{
		Client: client,
	}
}

// ActionVpsPasswdMetaGlobalInput is a type for action global meta input parameters
type ActionVpsPasswdMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsPasswdMetaGlobalInput) SetNo(value bool) *ActionVpsPasswdMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsPasswdMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsPasswdMetaGlobalInput) SelectParameters(params ...string) *ActionVpsPasswdMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsPasswdMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsPasswdInput is a type for action input parameters
type ActionVpsPasswdInput struct {
	Type string `json:"type"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetType sets parameter Type to value and selects it for sending
func (in *ActionVpsPasswdInput) SetType(value string) *ActionVpsPasswdInput {
	in.Type = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Type"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsPasswdInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsPasswdInput) SelectParameters(params ...string) *ActionVpsPasswdInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsPasswdInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsPasswdInput) UnselectParameters(params ...string) *ActionVpsPasswdInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsPasswdInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsPasswdRequest is a type for the entire action request
type ActionVpsPasswdRequest struct {
	Vps  map[string]interface{} `json:"vps"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionVpsPasswdOutput is a type for action output parameters
type ActionVpsPasswdOutput struct {
	Password string `json:"password"`
}

// ActionVpsPasswdMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsPasswdMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsPasswdResponse struct {
	Action *ActionVpsPasswd `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Vps *ActionVpsPasswdOutput `json:"vps"`
		// Global output metadata
		Meta *ActionVpsPasswdMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionVpsPasswdOutput
}

// Prepare the action for invocation
func (action *ActionVpsPasswd) Prepare() *ActionVpsPasswdInvocation {
	return &ActionVpsPasswdInvocation{
		Action: action,
		Path:   "/v6.0/vpses/{vps_id}/passwd",
	}
}

// ActionVpsPasswdInvocation is used to configure action for invocation
type ActionVpsPasswdInvocation struct {
	// Pointer to the action
	Action *ActionVpsPasswd

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsPasswdInput
	// Global meta input parameters
	MetaInput *ActionVpsPasswdMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsPasswdInvocation) SetPathParamInt(param string, value int64) *ActionVpsPasswdInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsPasswdInvocation) SetPathParamString(param string, value string) *ActionVpsPasswdInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsPasswdInvocation) NewInput() *ActionVpsPasswdInput {
	inv.Input = &ActionVpsPasswdInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsPasswdInvocation) SetInput(input *ActionVpsPasswdInput) *ActionVpsPasswdInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsPasswdInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsPasswdInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsPasswdInvocation) NewMetaInput() *ActionVpsPasswdMetaGlobalInput {
	inv.MetaInput = &ActionVpsPasswdMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsPasswdInvocation) SetMetaInput(input *ActionVpsPasswdMetaGlobalInput) *ActionVpsPasswdInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsPasswdInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsPasswdInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsPasswdInvocation) Call() (*ActionVpsPasswdResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionVpsPasswdInvocation) callAsBody() (*ActionVpsPasswdResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsPasswdResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Vps
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsPasswdResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsPasswdResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsPasswdResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsPasswdResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsPasswdResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionVpsPasswdInvocation) makeAllInputParams() *ActionVpsPasswdRequest {
	return &ActionVpsPasswdRequest{
		Vps:  inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsPasswdInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Type") {
			ret["type"] = inv.Input.Type
		}
	}

	return ret
}

func (inv *ActionVpsPasswdInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
