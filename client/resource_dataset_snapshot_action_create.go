package client

import (
	"strings"
)

// ActionDatasetSnapshotCreate is a type for action Dataset.Snapshot#Create
type ActionDatasetSnapshotCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetSnapshotCreate(client *Client) *ActionDatasetSnapshotCreate {
	return &ActionDatasetSnapshotCreate{
		Client: client,
	}
}

// ActionDatasetSnapshotCreateMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetSnapshotCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetSnapshotCreateMetaGlobalInput) SetIncludes(value string) *ActionDatasetSnapshotCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetSnapshotCreateMetaGlobalInput) SetNo(value bool) *ActionDatasetSnapshotCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetSnapshotCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetSnapshotCreateMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetSnapshotCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetSnapshotCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetSnapshotCreateInput is a type for action input parameters
type ActionDatasetSnapshotCreateInput struct {
	Label string `json:"label"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionDatasetSnapshotCreateInput) SetLabel(value string) *ActionDatasetSnapshotCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetSnapshotCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetSnapshotCreateInput) SelectParameters(params ...string) *ActionDatasetSnapshotCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDatasetSnapshotCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDatasetSnapshotCreateInput) UnselectParameters(params ...string) *ActionDatasetSnapshotCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDatasetSnapshotCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetSnapshotCreateRequest is a type for the entire action request
type ActionDatasetSnapshotCreateRequest struct {
	Snapshot map[string]interface{} `json:"snapshot"`
	Meta     map[string]interface{} `json:"_meta"`
}

// ActionDatasetSnapshotCreateOutput is a type for action output parameters
type ActionDatasetSnapshotCreateOutput struct {
	CreatedAt string                    `json:"created_at"`
	Dataset   *ActionDatasetShowOutput  `json:"dataset"`
	Export    *ActionExportShowOutput   `json:"export"`
	HistoryId int64                     `json:"history_id"`
	Id        int64                     `json:"id"`
	Label     string                    `json:"label"`
	Mount     *ActionVpsMountShowOutput `json:"mount"`
	Name      string                    `json:"name"`
}

// ActionDatasetSnapshotCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionDatasetSnapshotCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionDatasetSnapshotCreateResponse struct {
	Action *ActionDatasetSnapshotCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Snapshot *ActionDatasetSnapshotCreateOutput `json:"snapshot"`
		// Global output metadata
		Meta *ActionDatasetSnapshotCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionDatasetSnapshotCreateOutput
}

// Prepare the action for invocation
func (action *ActionDatasetSnapshotCreate) Prepare() *ActionDatasetSnapshotCreateInvocation {
	return &ActionDatasetSnapshotCreateInvocation{
		Action: action,
		Path:   "/v6.0/datasets/{dataset_id}/snapshots",
	}
}

// ActionDatasetSnapshotCreateInvocation is used to configure action for invocation
type ActionDatasetSnapshotCreateInvocation struct {
	// Pointer to the action
	Action *ActionDatasetSnapshotCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDatasetSnapshotCreateInput
	// Global meta input parameters
	MetaInput *ActionDatasetSnapshotCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetSnapshotCreateInvocation) SetPathParamInt(param string, value int64) *ActionDatasetSnapshotCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetSnapshotCreateInvocation) SetPathParamString(param string, value string) *ActionDatasetSnapshotCreateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDatasetSnapshotCreateInvocation) NewInput() *ActionDatasetSnapshotCreateInput {
	inv.Input = &ActionDatasetSnapshotCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDatasetSnapshotCreateInvocation) SetInput(input *ActionDatasetSnapshotCreateInput) *ActionDatasetSnapshotCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDatasetSnapshotCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDatasetSnapshotCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetSnapshotCreateInvocation) NewMetaInput() *ActionDatasetSnapshotCreateMetaGlobalInput {
	inv.MetaInput = &ActionDatasetSnapshotCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetSnapshotCreateInvocation) SetMetaInput(input *ActionDatasetSnapshotCreateMetaGlobalInput) *ActionDatasetSnapshotCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetSnapshotCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDatasetSnapshotCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetSnapshotCreateInvocation) Call() (*ActionDatasetSnapshotCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDatasetSnapshotCreateInvocation) callAsBody() (*ActionDatasetSnapshotCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDatasetSnapshotCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Snapshot
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDatasetSnapshotCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDatasetSnapshotCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDatasetSnapshotCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDatasetSnapshotCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionDatasetSnapshotCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionDatasetSnapshotCreateInvocation) makeAllInputParams() *ActionDatasetSnapshotCreateRequest {
	return &ActionDatasetSnapshotCreateRequest{
		Snapshot: inv.makeInputParams(),
		Meta:     inv.makeMetaInputParams(),
	}
}

func (inv *ActionDatasetSnapshotCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
	}

	return ret
}

func (inv *ActionDatasetSnapshotCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
