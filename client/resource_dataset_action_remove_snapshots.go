package client

import (
	"net/url"
	"strings"
)

// ActionDatasetRemoveSnapshots is a type for action Dataset#Remove_snapshots
type ActionDatasetRemoveSnapshots struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetRemoveSnapshots(client *Client) *ActionDatasetRemoveSnapshots {
	return &ActionDatasetRemoveSnapshots{
		Client: client,
	}
}

// ActionDatasetRemoveSnapshotsMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetRemoveSnapshotsMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetRemoveSnapshotsMetaGlobalInput) SetIncludes(value string) *ActionDatasetRemoveSnapshotsMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetRemoveSnapshotsMetaGlobalInput) SetNo(value bool) *ActionDatasetRemoveSnapshotsMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetRemoveSnapshotsMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetRemoveSnapshotsMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetRemoveSnapshotsMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetRemoveSnapshotsMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetRemoveSnapshotsInput is a type for action input parameters
type ActionDatasetRemoveSnapshotsInput struct {
	DisablePlans   bool   "json:\"disable_plans\""
	ManifestDigest string "json:\"manifest_digest\""
	Recursive      bool   "json:\"recursive\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDisablePlans sets parameter DisablePlans to value and selects it for sending
func (in *ActionDatasetRemoveSnapshotsInput) SetDisablePlans(value bool) *ActionDatasetRemoveSnapshotsInput {
	in.DisablePlans = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DisablePlans"] = nil
	return in
}

// SetManifestDigest sets parameter ManifestDigest to value and selects it for sending
func (in *ActionDatasetRemoveSnapshotsInput) SetManifestDigest(value string) *ActionDatasetRemoveSnapshotsInput {
	in.ManifestDigest = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ManifestDigest"] = nil
	return in
}

// SetRecursive sets parameter Recursive to value and selects it for sending
func (in *ActionDatasetRemoveSnapshotsInput) SetRecursive(value bool) *ActionDatasetRemoveSnapshotsInput {
	in.Recursive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Recursive"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetRemoveSnapshotsInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetRemoveSnapshotsInput) SelectParameters(params ...string) *ActionDatasetRemoveSnapshotsInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDatasetRemoveSnapshotsInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDatasetRemoveSnapshotsInput) UnselectParameters(params ...string) *ActionDatasetRemoveSnapshotsInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDatasetRemoveSnapshotsInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetRemoveSnapshotsRequest is a type for the entire action request
type ActionDatasetRemoveSnapshotsRequest struct {
	Dataset map[string]interface{} "json:\"dataset\""
	Meta    map[string]interface{} "json:\"_meta\""
}

// ActionDatasetRemoveSnapshotsMetaGlobalOutput is a type for global output metadata parameters
type ActionDatasetRemoveSnapshotsMetaGlobalOutput struct {
	ActionStateId int64 "json:\"action_state_id\""
}

// Type for action response, including envelope
type ActionDatasetRemoveSnapshotsResponse struct {
	Action *ActionDatasetRemoveSnapshots "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionDatasetRemoveSnapshotsMetaGlobalOutput "json:\"_meta\""
	}
}

// Prepare the action for invocation
func (action *ActionDatasetRemoveSnapshots) Prepare() *ActionDatasetRemoveSnapshotsInvocation {
	return &ActionDatasetRemoveSnapshotsInvocation{
		Action: action,
		Path:   "/v7.0/datasets/{dataset_id}/remove_snapshots",
	}
}

// ActionDatasetRemoveSnapshotsInvocation is used to configure action for invocation
type ActionDatasetRemoveSnapshotsInvocation struct {
	// Pointer to the action
	Action *ActionDatasetRemoveSnapshots

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDatasetRemoveSnapshotsInput
	// Global meta input parameters
	MetaInput *ActionDatasetRemoveSnapshotsMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetRemoveSnapshotsInvocation) SetPathParamInt(param string, value int64) *ActionDatasetRemoveSnapshotsInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetRemoveSnapshotsInvocation) SetPathParamString(param string, value string) *ActionDatasetRemoveSnapshotsInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDatasetRemoveSnapshotsInvocation) NewInput() *ActionDatasetRemoveSnapshotsInput {
	inv.Input = &ActionDatasetRemoveSnapshotsInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDatasetRemoveSnapshotsInvocation) SetInput(input *ActionDatasetRemoveSnapshotsInput) *ActionDatasetRemoveSnapshotsInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDatasetRemoveSnapshotsInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDatasetRemoveSnapshotsInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetRemoveSnapshotsInvocation) NewMetaInput() *ActionDatasetRemoveSnapshotsMetaGlobalInput {
	inv.MetaInput = &ActionDatasetRemoveSnapshotsMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetRemoveSnapshotsInvocation) SetMetaInput(input *ActionDatasetRemoveSnapshotsMetaGlobalInput) *ActionDatasetRemoveSnapshotsInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetRemoveSnapshotsInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDatasetRemoveSnapshotsInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionDatasetRemoveSnapshotsInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.Input != nil {
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetRemoveSnapshotsInvocation) Call() (*ActionDatasetRemoveSnapshotsResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionDatasetRemoveSnapshotsInvocation) callAsBody() (*ActionDatasetRemoveSnapshotsResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDatasetRemoveSnapshotsResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDatasetRemoveSnapshotsResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDatasetRemoveSnapshotsResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDatasetRemoveSnapshotsResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDatasetRemoveSnapshotsResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionDatasetRemoveSnapshotsResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionDatasetRemoveSnapshotsInvocation) makeAllInputParams() *ActionDatasetRemoveSnapshotsRequest {
	return &ActionDatasetRemoveSnapshotsRequest{
		Dataset: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionDatasetRemoveSnapshotsInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("DisablePlans") {
			ret["disable_plans"] = inv.Input.DisablePlans
		}
		if inv.IsParameterSelected("ManifestDigest") {
			ret["manifest_digest"] = inv.Input.ManifestDigest
		}
		if inv.IsParameterSelected("Recursive") {
			ret["recursive"] = inv.Input.Recursive
		}
	}

	return ret
}

func (inv *ActionDatasetRemoveSnapshotsInvocation) makeMetaInputParams() map[string]interface{} {
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
