package client

import (
	"strings"
)

// ActionDatasetSnapshotRollback is a type for action Dataset.Snapshot#Rollback
type ActionDatasetSnapshotRollback struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetSnapshotRollback(client *Client) *ActionDatasetSnapshotRollback {
	return &ActionDatasetSnapshotRollback{
		Client: client,
	}
}

// ActionDatasetSnapshotRollbackMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetSnapshotRollbackMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetSnapshotRollbackMetaGlobalInput) SetIncludes(value string) *ActionDatasetSnapshotRollbackMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetSnapshotRollbackMetaGlobalInput) SetNo(value bool) *ActionDatasetSnapshotRollbackMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetSnapshotRollbackMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetSnapshotRollbackMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetSnapshotRollbackMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetSnapshotRollbackMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetSnapshotRollbackRequest is a type for the entire action request
type ActionDatasetSnapshotRollbackRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// ActionDatasetSnapshotRollbackMetaGlobalOutput is a type for global output metadata parameters
type ActionDatasetSnapshotRollbackMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionDatasetSnapshotRollbackResponse struct {
	Action *ActionDatasetSnapshotRollback `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionDatasetSnapshotRollbackMetaGlobalOutput `json:"_meta"`
	}
}

// Prepare the action for invocation
func (action *ActionDatasetSnapshotRollback) Prepare() *ActionDatasetSnapshotRollbackInvocation {
	return &ActionDatasetSnapshotRollbackInvocation{
		Action: action,
		Path:   "/v6.0/datasets/{dataset_id}/snapshots/{snapshot_id}/rollback",
	}
}

// ActionDatasetSnapshotRollbackInvocation is used to configure action for invocation
type ActionDatasetSnapshotRollbackInvocation struct {
	// Pointer to the action
	Action *ActionDatasetSnapshotRollback

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDatasetSnapshotRollbackMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetSnapshotRollbackInvocation) SetPathParamInt(param string, value int64) *ActionDatasetSnapshotRollbackInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetSnapshotRollbackInvocation) SetPathParamString(param string, value string) *ActionDatasetSnapshotRollbackInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetSnapshotRollbackInvocation) NewMetaInput() *ActionDatasetSnapshotRollbackMetaGlobalInput {
	inv.MetaInput = &ActionDatasetSnapshotRollbackMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetSnapshotRollbackInvocation) SetMetaInput(input *ActionDatasetSnapshotRollbackMetaGlobalInput) *ActionDatasetSnapshotRollbackInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetSnapshotRollbackInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetSnapshotRollbackInvocation) Call() (*ActionDatasetSnapshotRollbackResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDatasetSnapshotRollbackInvocation) callAsBody() (*ActionDatasetSnapshotRollbackResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDatasetSnapshotRollbackResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDatasetSnapshotRollbackResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDatasetSnapshotRollbackResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDatasetSnapshotRollbackResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDatasetSnapshotRollbackResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionDatasetSnapshotRollbackResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionDatasetSnapshotRollbackInvocation) makeAllInputParams() *ActionDatasetSnapshotRollbackRequest {
	return &ActionDatasetSnapshotRollbackRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionDatasetSnapshotRollbackInvocation) makeMetaInputParams() map[string]interface{} {
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
