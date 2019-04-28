package client

import (
	"strings"
)

// ActionSnapshotDownloadDelete is a type for action Snapshot_download#Delete
type ActionSnapshotDownloadDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionSnapshotDownloadDelete(client *Client) *ActionSnapshotDownloadDelete {
	return &ActionSnapshotDownloadDelete{
		Client: client,
	}
}

// ActionSnapshotDownloadDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionSnapshotDownloadDeleteMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSnapshotDownloadDeleteMetaGlobalInput) SetNo(value bool) *ActionSnapshotDownloadDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSnapshotDownloadDeleteMetaGlobalInput) SetIncludes(value string) *ActionSnapshotDownloadDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionSnapshotDownloadDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSnapshotDownloadDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionSnapshotDownloadDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSnapshotDownloadDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionSnapshotDownloadDeleteRequest is a type for the entire action request
type ActionSnapshotDownloadDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}


// ActionSnapshotDownloadDeleteMetaGlobalOutput is a type for global output metadata parameters
type ActionSnapshotDownloadDeleteMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionSnapshotDownloadDeleteResponse struct {
	Action *ActionSnapshotDownloadDelete `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionSnapshotDownloadDeleteMetaGlobalOutput `json:"_meta"`
	}
}


// Prepare the action for invocation
func (action *ActionSnapshotDownloadDelete) Prepare() *ActionSnapshotDownloadDeleteInvocation {
	return &ActionSnapshotDownloadDeleteInvocation{
		Action: action,
		Path: "/v5.0/snapshot_downloads/:snapshot_download_id",
	}
}

// ActionSnapshotDownloadDeleteInvocation is used to configure action for invocation
type ActionSnapshotDownloadDeleteInvocation struct {
	// Pointer to the action
	Action *ActionSnapshotDownloadDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionSnapshotDownloadDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSnapshotDownloadDeleteInvocation) SetPathParamInt(param string, value int64) *ActionSnapshotDownloadDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSnapshotDownloadDeleteInvocation) SetPathParamString(param string, value string) *ActionSnapshotDownloadDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSnapshotDownloadDeleteInvocation) NewMetaInput() *ActionSnapshotDownloadDeleteMetaGlobalInput {
	inv.MetaInput = &ActionSnapshotDownloadDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSnapshotDownloadDeleteInvocation) SetMetaInput(input *ActionSnapshotDownloadDeleteMetaGlobalInput) *ActionSnapshotDownloadDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSnapshotDownloadDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSnapshotDownloadDeleteInvocation) Call() (*ActionSnapshotDownloadDeleteResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionSnapshotDownloadDeleteInvocation) callAsBody() (*ActionSnapshotDownloadDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSnapshotDownloadDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionSnapshotDownloadDeleteResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionSnapshotDownloadDeleteResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionSnapshotDownloadDeleteResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionSnapshotDownloadDeleteResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
			Timeout: timeout,
			UpdateIn: updateIn,
			Status: pollResp.Output.Status,
			Current: pollResp.Output.Current,
			Total: pollResp.Output.Total,
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
func (resp *ActionSnapshotDownloadDeleteResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionSnapshotDownloadDeleteInvocation) makeAllInputParams() *ActionSnapshotDownloadDeleteRequest {
	return &ActionSnapshotDownloadDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionSnapshotDownloadDeleteInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
