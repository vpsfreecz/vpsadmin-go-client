package client

import (
	"net/url"
	"strings"
)

// ActionDatasetMigrate is a type for action Dataset#Migrate
type ActionDatasetMigrate struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetMigrate(client *Client) *ActionDatasetMigrate {
	return &ActionDatasetMigrate{
		Client: client,
	}
}

// ActionDatasetMigrateMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetMigrateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetMigrateMetaGlobalInput) SetIncludes(value string) *ActionDatasetMigrateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetMigrateMetaGlobalInput) SetNo(value bool) *ActionDatasetMigrateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetMigrateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetMigrateMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetMigrateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetMigrateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetMigrateInput is a type for action input parameters
type ActionDatasetMigrateInput struct {
	CleanupData               bool   "json:\"cleanup_data\""
	FinishMinutes             int64  "json:\"finish_minutes\""
	FinishWeekday             int64  "json:\"finish_weekday\""
	MaintenanceWindowVps      int64  "json:\"maintenance_window_vps\""
	OptionalMaintenanceWindow bool   "json:\"optional_maintenance_window\""
	Pool                      int64  "json:\"pool\""
	Reason                    string "json:\"reason\""
	RestartVps                bool   "json:\"restart_vps\""
	Rsync                     bool   "json:\"rsync\""
	SendMail                  bool   "json:\"send_mail\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCleanupData sets parameter CleanupData to value and selects it for sending
func (in *ActionDatasetMigrateInput) SetCleanupData(value bool) *ActionDatasetMigrateInput {
	in.CleanupData = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CleanupData"] = nil
	return in
}

// SetFinishMinutes sets parameter FinishMinutes to value and selects it for sending
func (in *ActionDatasetMigrateInput) SetFinishMinutes(value int64) *ActionDatasetMigrateInput {
	in.FinishMinutes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FinishMinutes"] = nil
	return in
}

// SetFinishWeekday sets parameter FinishWeekday to value and selects it for sending
func (in *ActionDatasetMigrateInput) SetFinishWeekday(value int64) *ActionDatasetMigrateInput {
	in.FinishWeekday = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FinishWeekday"] = nil
	return in
}

// SetMaintenanceWindowVps sets parameter MaintenanceWindowVps to value and selects it for sending
func (in *ActionDatasetMigrateInput) SetMaintenanceWindowVps(value int64) *ActionDatasetMigrateInput {
	in.MaintenanceWindowVps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaintenanceWindowVps"] = nil
	return in
}

// SetOptionalMaintenanceWindow sets parameter OptionalMaintenanceWindow to value and selects it for sending
func (in *ActionDatasetMigrateInput) SetOptionalMaintenanceWindow(value bool) *ActionDatasetMigrateInput {
	in.OptionalMaintenanceWindow = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OptionalMaintenanceWindow"] = nil
	return in
}

// SetPool sets parameter Pool to value and selects it for sending
func (in *ActionDatasetMigrateInput) SetPool(value int64) *ActionDatasetMigrateInput {
	in.Pool = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Pool"] = nil
	return in
}

// SetReason sets parameter Reason to value and selects it for sending
func (in *ActionDatasetMigrateInput) SetReason(value string) *ActionDatasetMigrateInput {
	in.Reason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Reason"] = nil
	return in
}

// SetRestartVps sets parameter RestartVps to value and selects it for sending
func (in *ActionDatasetMigrateInput) SetRestartVps(value bool) *ActionDatasetMigrateInput {
	in.RestartVps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RestartVps"] = nil
	return in
}

// SetRsync sets parameter Rsync to value and selects it for sending
func (in *ActionDatasetMigrateInput) SetRsync(value bool) *ActionDatasetMigrateInput {
	in.Rsync = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Rsync"] = nil
	return in
}

// SetSendMail sets parameter SendMail to value and selects it for sending
func (in *ActionDatasetMigrateInput) SetSendMail(value bool) *ActionDatasetMigrateInput {
	in.SendMail = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SendMail"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetMigrateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetMigrateInput) SelectParameters(params ...string) *ActionDatasetMigrateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDatasetMigrateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDatasetMigrateInput) UnselectParameters(params ...string) *ActionDatasetMigrateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDatasetMigrateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetMigrateRequest is a type for the entire action request
type ActionDatasetMigrateRequest struct {
	Dataset map[string]interface{} "json:\"dataset\""
	Meta    map[string]interface{} "json:\"_meta\""
}

// ActionDatasetMigrateMetaGlobalOutput is a type for global output metadata parameters
type ActionDatasetMigrateMetaGlobalOutput struct {
	ActionStateId int64 "json:\"action_state_id\""
}

// Type for action response, including envelope
type ActionDatasetMigrateResponse struct {
	Action *ActionDatasetMigrate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionDatasetMigrateMetaGlobalOutput "json:\"_meta\""
	}
}

// Prepare the action for invocation
func (action *ActionDatasetMigrate) Prepare() *ActionDatasetMigrateInvocation {
	return &ActionDatasetMigrateInvocation{
		Action: action,
		Path:   "/v7.0/datasets/{dataset_id}/migrate",
	}
}

// ActionDatasetMigrateInvocation is used to configure action for invocation
type ActionDatasetMigrateInvocation struct {
	// Pointer to the action
	Action *ActionDatasetMigrate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDatasetMigrateInput
	// Global meta input parameters
	MetaInput *ActionDatasetMigrateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetMigrateInvocation) SetPathParamInt(param string, value int64) *ActionDatasetMigrateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetMigrateInvocation) SetPathParamString(param string, value string) *ActionDatasetMigrateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDatasetMigrateInvocation) NewInput() *ActionDatasetMigrateInput {
	inv.Input = &ActionDatasetMigrateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDatasetMigrateInvocation) SetInput(input *ActionDatasetMigrateInput) *ActionDatasetMigrateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDatasetMigrateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDatasetMigrateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetMigrateInvocation) NewMetaInput() *ActionDatasetMigrateMetaGlobalInput {
	inv.MetaInput = &ActionDatasetMigrateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetMigrateInvocation) SetMetaInput(input *ActionDatasetMigrateMetaGlobalInput) *ActionDatasetMigrateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetMigrateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDatasetMigrateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionDatasetMigrateInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("MaintenanceWindowVps") {
			if !inv.IsParameterNil("MaintenanceWindowVps") {
				if inv.Input.MaintenanceWindowVps < 0 {
					verr.Add("maintenance_window_vps", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("Pool") {
			if !inv.IsParameterNil("Pool") {
				if inv.Input.Pool < 0 {
					verr.Add("pool", "not a valid resource id")
				}
			}
		}
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetMigrateInvocation) Call() (*ActionDatasetMigrateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionDatasetMigrateInvocation) callAsBody() (*ActionDatasetMigrateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDatasetMigrateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDatasetMigrateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDatasetMigrateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDatasetMigrateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDatasetMigrateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionDatasetMigrateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionDatasetMigrateInvocation) makeAllInputParams() *ActionDatasetMigrateRequest {
	return &ActionDatasetMigrateRequest{
		Dataset: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionDatasetMigrateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("CleanupData") {
			ret["cleanup_data"] = inv.Input.CleanupData
		}
		if inv.IsParameterSelected("FinishMinutes") {
			ret["finish_minutes"] = inv.Input.FinishMinutes
		}
		if inv.IsParameterSelected("FinishWeekday") {
			ret["finish_weekday"] = inv.Input.FinishWeekday
		}
		if inv.IsParameterSelected("MaintenanceWindowVps") {
			ret["maintenance_window_vps"] = inv.Input.MaintenanceWindowVps
		}
		if inv.IsParameterSelected("OptionalMaintenanceWindow") {
			ret["optional_maintenance_window"] = inv.Input.OptionalMaintenanceWindow
		}
		if inv.IsParameterSelected("Pool") {
			ret["pool"] = inv.Input.Pool
		}
		if inv.IsParameterSelected("Reason") {
			ret["reason"] = inv.Input.Reason
		}
		if inv.IsParameterSelected("RestartVps") {
			ret["restart_vps"] = inv.Input.RestartVps
		}
		if inv.IsParameterSelected("Rsync") {
			ret["rsync"] = inv.Input.Rsync
		}
		if inv.IsParameterSelected("SendMail") {
			ret["send_mail"] = inv.Input.SendMail
		}
	}

	return ret
}

func (inv *ActionDatasetMigrateInvocation) makeMetaInputParams() map[string]interface{} {
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
