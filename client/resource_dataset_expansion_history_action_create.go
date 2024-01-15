package client

import (
	"strings"
)

// ActionDatasetExpansionHistoryCreate is a type for action Dataset_expansion.History#Create
type ActionDatasetExpansionHistoryCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetExpansionHistoryCreate(client *Client) *ActionDatasetExpansionHistoryCreate {
	return &ActionDatasetExpansionHistoryCreate{
		Client: client,
	}
}

// ActionDatasetExpansionHistoryCreateMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetExpansionHistoryCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetExpansionHistoryCreateMetaGlobalInput) SetIncludes(value string) *ActionDatasetExpansionHistoryCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetExpansionHistoryCreateMetaGlobalInput) SetNo(value bool) *ActionDatasetExpansionHistoryCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetExpansionHistoryCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetExpansionHistoryCreateMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetExpansionHistoryCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetExpansionHistoryCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetExpansionHistoryCreateInput is a type for action input parameters
type ActionDatasetExpansionHistoryCreateInput struct {
	AddedSpace int64 `json:"added_space"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAddedSpace sets parameter AddedSpace to value and selects it for sending
func (in *ActionDatasetExpansionHistoryCreateInput) SetAddedSpace(value int64) *ActionDatasetExpansionHistoryCreateInput {
	in.AddedSpace = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AddedSpace"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetExpansionHistoryCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetExpansionHistoryCreateInput) SelectParameters(params ...string) *ActionDatasetExpansionHistoryCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDatasetExpansionHistoryCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDatasetExpansionHistoryCreateInput) UnselectParameters(params ...string) *ActionDatasetExpansionHistoryCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDatasetExpansionHistoryCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetExpansionHistoryCreateRequest is a type for the entire action request
type ActionDatasetExpansionHistoryCreateRequest struct {
	History map[string]interface{} `json:"history"`
	Meta    map[string]interface{} `json:"_meta"`
}

// ActionDatasetExpansionHistoryCreateOutput is a type for action output parameters
type ActionDatasetExpansionHistoryCreateOutput struct {
	AddedSpace       int64                 `json:"added_space"`
	Admin            *ActionUserShowOutput `json:"admin"`
	CreatedAt        string                `json:"created_at"`
	Id               int64                 `json:"id"`
	NewRefquota      int64                 `json:"new_refquota"`
	OriginalRefquota int64                 `json:"original_refquota"`
}

// ActionDatasetExpansionHistoryCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionDatasetExpansionHistoryCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionDatasetExpansionHistoryCreateResponse struct {
	Action *ActionDatasetExpansionHistoryCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		History *ActionDatasetExpansionHistoryCreateOutput `json:"history"`
		// Global output metadata
		Meta *ActionDatasetExpansionHistoryCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionDatasetExpansionHistoryCreateOutput
}

// Prepare the action for invocation
func (action *ActionDatasetExpansionHistoryCreate) Prepare() *ActionDatasetExpansionHistoryCreateInvocation {
	return &ActionDatasetExpansionHistoryCreateInvocation{
		Action: action,
		Path:   "/v6.0/dataset_expansions/{dataset_expansion_id}/history",
	}
}

// ActionDatasetExpansionHistoryCreateInvocation is used to configure action for invocation
type ActionDatasetExpansionHistoryCreateInvocation struct {
	// Pointer to the action
	Action *ActionDatasetExpansionHistoryCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDatasetExpansionHistoryCreateInput
	// Global meta input parameters
	MetaInput *ActionDatasetExpansionHistoryCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetExpansionHistoryCreateInvocation) SetPathParamInt(param string, value int64) *ActionDatasetExpansionHistoryCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetExpansionHistoryCreateInvocation) SetPathParamString(param string, value string) *ActionDatasetExpansionHistoryCreateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDatasetExpansionHistoryCreateInvocation) NewInput() *ActionDatasetExpansionHistoryCreateInput {
	inv.Input = &ActionDatasetExpansionHistoryCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDatasetExpansionHistoryCreateInvocation) SetInput(input *ActionDatasetExpansionHistoryCreateInput) *ActionDatasetExpansionHistoryCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDatasetExpansionHistoryCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDatasetExpansionHistoryCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetExpansionHistoryCreateInvocation) NewMetaInput() *ActionDatasetExpansionHistoryCreateMetaGlobalInput {
	inv.MetaInput = &ActionDatasetExpansionHistoryCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetExpansionHistoryCreateInvocation) SetMetaInput(input *ActionDatasetExpansionHistoryCreateMetaGlobalInput) *ActionDatasetExpansionHistoryCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetExpansionHistoryCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDatasetExpansionHistoryCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetExpansionHistoryCreateInvocation) Call() (*ActionDatasetExpansionHistoryCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDatasetExpansionHistoryCreateInvocation) callAsBody() (*ActionDatasetExpansionHistoryCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDatasetExpansionHistoryCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.History
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDatasetExpansionHistoryCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDatasetExpansionHistoryCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDatasetExpansionHistoryCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDatasetExpansionHistoryCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionDatasetExpansionHistoryCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionDatasetExpansionHistoryCreateInvocation) makeAllInputParams() *ActionDatasetExpansionHistoryCreateRequest {
	return &ActionDatasetExpansionHistoryCreateRequest{
		History: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionDatasetExpansionHistoryCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("AddedSpace") {
			ret["added_space"] = inv.Input.AddedSpace
		}
	}

	return ret
}

func (inv *ActionDatasetExpansionHistoryCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
