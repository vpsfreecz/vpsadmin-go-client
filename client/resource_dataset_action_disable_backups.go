package client

import (
	"net/url"
	"strings"
)

// ActionDatasetDisableBackups is a type for action Dataset#Disable_backups
type ActionDatasetDisableBackups struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetDisableBackups(client *Client) *ActionDatasetDisableBackups {
	return &ActionDatasetDisableBackups{
		Client: client,
	}
}

// ActionDatasetDisableBackupsMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetDisableBackupsMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetDisableBackupsMetaGlobalInput) SetIncludes(value string) *ActionDatasetDisableBackupsMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetDisableBackupsMetaGlobalInput) SetNo(value bool) *ActionDatasetDisableBackupsMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetDisableBackupsMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetDisableBackupsMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetDisableBackupsMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetDisableBackupsMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetDisableBackupsInput is a type for action input parameters
type ActionDatasetDisableBackupsInput struct {
	BackupFirst bool "json:\"backup_first\""
	Recursive   bool "json:\"recursive\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetBackupFirst sets parameter BackupFirst to value and selects it for sending
func (in *ActionDatasetDisableBackupsInput) SetBackupFirst(value bool) *ActionDatasetDisableBackupsInput {
	in.BackupFirst = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["BackupFirst"] = nil
	return in
}

// SetRecursive sets parameter Recursive to value and selects it for sending
func (in *ActionDatasetDisableBackupsInput) SetRecursive(value bool) *ActionDatasetDisableBackupsInput {
	in.Recursive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Recursive"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetDisableBackupsInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetDisableBackupsInput) SelectParameters(params ...string) *ActionDatasetDisableBackupsInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDatasetDisableBackupsInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDatasetDisableBackupsInput) UnselectParameters(params ...string) *ActionDatasetDisableBackupsInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDatasetDisableBackupsInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetDisableBackupsRequest is a type for the entire action request
type ActionDatasetDisableBackupsRequest struct {
	Dataset map[string]interface{} "json:\"dataset\""
	Meta    map[string]interface{} "json:\"_meta\""
}

// ActionDatasetDisableBackupsMetaGlobalOutput is a type for global output metadata parameters
type ActionDatasetDisableBackupsMetaGlobalOutput struct {
	ActionStateId int64 "json:\"action_state_id\""
}

// Type for action response, including envelope
type ActionDatasetDisableBackupsResponse struct {
	Action *ActionDatasetDisableBackups "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionDatasetDisableBackupsMetaGlobalOutput "json:\"_meta\""
	}
}

// Prepare the action for invocation
func (action *ActionDatasetDisableBackups) Prepare() *ActionDatasetDisableBackupsInvocation {
	return &ActionDatasetDisableBackupsInvocation{
		Action: action,
		Path:   "/v7.0/datasets/{dataset_id}/disable_backups",
	}
}

// ActionDatasetDisableBackupsInvocation is used to configure action for invocation
type ActionDatasetDisableBackupsInvocation struct {
	// Pointer to the action
	Action *ActionDatasetDisableBackups

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDatasetDisableBackupsInput
	// Global meta input parameters
	MetaInput *ActionDatasetDisableBackupsMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetDisableBackupsInvocation) SetPathParamInt(param string, value int64) *ActionDatasetDisableBackupsInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetDisableBackupsInvocation) SetPathParamString(param string, value string) *ActionDatasetDisableBackupsInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDatasetDisableBackupsInvocation) NewInput() *ActionDatasetDisableBackupsInput {
	inv.Input = &ActionDatasetDisableBackupsInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDatasetDisableBackupsInvocation) SetInput(input *ActionDatasetDisableBackupsInput) *ActionDatasetDisableBackupsInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDatasetDisableBackupsInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDatasetDisableBackupsInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetDisableBackupsInvocation) NewMetaInput() *ActionDatasetDisableBackupsMetaGlobalInput {
	inv.MetaInput = &ActionDatasetDisableBackupsMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetDisableBackupsInvocation) SetMetaInput(input *ActionDatasetDisableBackupsMetaGlobalInput) *ActionDatasetDisableBackupsInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetDisableBackupsInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDatasetDisableBackupsInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionDatasetDisableBackupsInvocation) validate() error {
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
func (inv *ActionDatasetDisableBackupsInvocation) Call() (*ActionDatasetDisableBackupsResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionDatasetDisableBackupsInvocation) callAsBody() (*ActionDatasetDisableBackupsResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDatasetDisableBackupsResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDatasetDisableBackupsResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDatasetDisableBackupsResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDatasetDisableBackupsResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDatasetDisableBackupsResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionDatasetDisableBackupsResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionDatasetDisableBackupsInvocation) makeAllInputParams() *ActionDatasetDisableBackupsRequest {
	return &ActionDatasetDisableBackupsRequest{
		Dataset: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionDatasetDisableBackupsInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("BackupFirst") {
			ret["backup_first"] = inv.Input.BackupFirst
		}
		if inv.IsParameterSelected("Recursive") {
			ret["recursive"] = inv.Input.Recursive
		}
	}

	return ret
}

func (inv *ActionDatasetDisableBackupsInvocation) makeMetaInputParams() map[string]interface{} {
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
