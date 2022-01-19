package client

import (
	"strings"
)

// ActionDatasetSnapshotDelete is a type for action Dataset.Snapshot#Delete
type ActionDatasetSnapshotDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetSnapshotDelete(client *Client) *ActionDatasetSnapshotDelete {
	return &ActionDatasetSnapshotDelete{
		Client: client,
	}
}

// ActionDatasetSnapshotDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetSnapshotDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetSnapshotDeleteMetaGlobalInput) SetIncludes(value string) *ActionDatasetSnapshotDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetSnapshotDeleteMetaGlobalInput) SetNo(value bool) *ActionDatasetSnapshotDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetSnapshotDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetSnapshotDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetSnapshotDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetSnapshotDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetSnapshotDeleteRequest is a type for the entire action request
type ActionDatasetSnapshotDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// ActionDatasetSnapshotDeleteMetaGlobalOutput is a type for global output metadata parameters
type ActionDatasetSnapshotDeleteMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionDatasetSnapshotDeleteResponse struct {
	Action *ActionDatasetSnapshotDelete `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionDatasetSnapshotDeleteMetaGlobalOutput `json:"_meta"`
	}
}

// Prepare the action for invocation
func (action *ActionDatasetSnapshotDelete) Prepare() *ActionDatasetSnapshotDeleteInvocation {
	return &ActionDatasetSnapshotDeleteInvocation{
		Action: action,
		Path:   "/v6.0/datasets/{dataset_id}/snapshots/{snapshot_id}",
	}
}

// ActionDatasetSnapshotDeleteInvocation is used to configure action for invocation
type ActionDatasetSnapshotDeleteInvocation struct {
	// Pointer to the action
	Action *ActionDatasetSnapshotDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDatasetSnapshotDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetSnapshotDeleteInvocation) SetPathParamInt(param string, value int64) *ActionDatasetSnapshotDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetSnapshotDeleteInvocation) SetPathParamString(param string, value string) *ActionDatasetSnapshotDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetSnapshotDeleteInvocation) NewMetaInput() *ActionDatasetSnapshotDeleteMetaGlobalInput {
	inv.MetaInput = &ActionDatasetSnapshotDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetSnapshotDeleteInvocation) SetMetaInput(input *ActionDatasetSnapshotDeleteMetaGlobalInput) *ActionDatasetSnapshotDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetSnapshotDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetSnapshotDeleteInvocation) Call() (*ActionDatasetSnapshotDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDatasetSnapshotDeleteInvocation) callAsBody() (*ActionDatasetSnapshotDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDatasetSnapshotDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDatasetSnapshotDeleteResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDatasetSnapshotDeleteResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDatasetSnapshotDeleteResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDatasetSnapshotDeleteResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionDatasetSnapshotDeleteResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionDatasetSnapshotDeleteInvocation) makeAllInputParams() *ActionDatasetSnapshotDeleteRequest {
	return &ActionDatasetSnapshotDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionDatasetSnapshotDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
