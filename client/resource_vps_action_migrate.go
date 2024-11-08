package client

import (
	"strings"
)

// ActionVpsMigrate is a type for action Vps#Migrate
type ActionVpsMigrate struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsMigrate(client *Client) *ActionVpsMigrate {
	return &ActionVpsMigrate{
		Client: client,
	}
}

// ActionVpsMigrateMetaGlobalInput is a type for action global meta input parameters
type ActionVpsMigrateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsMigrateMetaGlobalInput) SetIncludes(value string) *ActionVpsMigrateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsMigrateMetaGlobalInput) SetNo(value bool) *ActionVpsMigrateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsMigrateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsMigrateMetaGlobalInput) SelectParameters(params ...string) *ActionVpsMigrateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsMigrateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsMigrateInput is a type for action input parameters
type ActionVpsMigrateInput struct {
	CleanupData         bool   `json:"cleanup_data"`
	FinishMinutes       int64  `json:"finish_minutes"`
	FinishWeekday       int64  `json:"finish_weekday"`
	MaintenanceWindow   bool   `json:"maintenance_window"`
	NoStart             bool   `json:"no_start"`
	Node                int64  `json:"node"`
	Reason              string `json:"reason"`
	ReplaceIpAddresses  bool   `json:"replace_ip_addresses"`
	SendMail            bool   `json:"send_mail"`
	SkipStart           bool   `json:"skip_start"`
	Swap                string `json:"swap"`
	TransferIpAddresses bool   `json:"transfer_ip_addresses"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCleanupData sets parameter CleanupData to value and selects it for sending
func (in *ActionVpsMigrateInput) SetCleanupData(value bool) *ActionVpsMigrateInput {
	in.CleanupData = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CleanupData"] = nil
	return in
}

// SetFinishMinutes sets parameter FinishMinutes to value and selects it for sending
func (in *ActionVpsMigrateInput) SetFinishMinutes(value int64) *ActionVpsMigrateInput {
	in.FinishMinutes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FinishMinutes"] = nil
	return in
}

// SetFinishWeekday sets parameter FinishWeekday to value and selects it for sending
func (in *ActionVpsMigrateInput) SetFinishWeekday(value int64) *ActionVpsMigrateInput {
	in.FinishWeekday = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FinishWeekday"] = nil
	return in
}

// SetMaintenanceWindow sets parameter MaintenanceWindow to value and selects it for sending
func (in *ActionVpsMigrateInput) SetMaintenanceWindow(value bool) *ActionVpsMigrateInput {
	in.MaintenanceWindow = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaintenanceWindow"] = nil
	return in
}

// SetNoStart sets parameter NoStart to value and selects it for sending
func (in *ActionVpsMigrateInput) SetNoStart(value bool) *ActionVpsMigrateInput {
	in.NoStart = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NoStart"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionVpsMigrateInput) SetNode(value int64) *ActionVpsMigrateInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNodeNil(false)
	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeNil sets parameter Node to nil and selects it for sending
func (in *ActionVpsMigrateInput) SetNodeNil(set bool) *ActionVpsMigrateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Node"] = nil
		in.SelectParameters("Node")
	} else {
		delete(in._nilParameters, "Node")
	}
	return in
}

// SetReason sets parameter Reason to value and selects it for sending
func (in *ActionVpsMigrateInput) SetReason(value string) *ActionVpsMigrateInput {
	in.Reason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Reason"] = nil
	return in
}

// SetReplaceIpAddresses sets parameter ReplaceIpAddresses to value and selects it for sending
func (in *ActionVpsMigrateInput) SetReplaceIpAddresses(value bool) *ActionVpsMigrateInput {
	in.ReplaceIpAddresses = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ReplaceIpAddresses"] = nil
	return in
}

// SetSendMail sets parameter SendMail to value and selects it for sending
func (in *ActionVpsMigrateInput) SetSendMail(value bool) *ActionVpsMigrateInput {
	in.SendMail = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SendMail"] = nil
	return in
}

// SetSkipStart sets parameter SkipStart to value and selects it for sending
func (in *ActionVpsMigrateInput) SetSkipStart(value bool) *ActionVpsMigrateInput {
	in.SkipStart = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SkipStart"] = nil
	return in
}

// SetSwap sets parameter Swap to value and selects it for sending
func (in *ActionVpsMigrateInput) SetSwap(value string) *ActionVpsMigrateInput {
	in.Swap = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Swap"] = nil
	return in
}

// SetTransferIpAddresses sets parameter TransferIpAddresses to value and selects it for sending
func (in *ActionVpsMigrateInput) SetTransferIpAddresses(value bool) *ActionVpsMigrateInput {
	in.TransferIpAddresses = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TransferIpAddresses"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsMigrateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsMigrateInput) SelectParameters(params ...string) *ActionVpsMigrateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsMigrateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsMigrateInput) UnselectParameters(params ...string) *ActionVpsMigrateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsMigrateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsMigrateRequest is a type for the entire action request
type ActionVpsMigrateRequest struct {
	Vps  map[string]interface{} `json:"vps"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionVpsMigrateMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsMigrateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsMigrateResponse struct {
	Action *ActionVpsMigrate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionVpsMigrateMetaGlobalOutput `json:"_meta"`
	}
}

// Prepare the action for invocation
func (action *ActionVpsMigrate) Prepare() *ActionVpsMigrateInvocation {
	return &ActionVpsMigrateInvocation{
		Action: action,
		Path:   "/v7.0/vpses/{vps_id}/migrate",
	}
}

// ActionVpsMigrateInvocation is used to configure action for invocation
type ActionVpsMigrateInvocation struct {
	// Pointer to the action
	Action *ActionVpsMigrate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsMigrateInput
	// Global meta input parameters
	MetaInput *ActionVpsMigrateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsMigrateInvocation) SetPathParamInt(param string, value int64) *ActionVpsMigrateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsMigrateInvocation) SetPathParamString(param string, value string) *ActionVpsMigrateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsMigrateInvocation) NewInput() *ActionVpsMigrateInput {
	inv.Input = &ActionVpsMigrateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsMigrateInvocation) SetInput(input *ActionVpsMigrateInput) *ActionVpsMigrateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsMigrateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsMigrateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsMigrateInvocation) NewMetaInput() *ActionVpsMigrateMetaGlobalInput {
	inv.MetaInput = &ActionVpsMigrateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsMigrateInvocation) SetMetaInput(input *ActionVpsMigrateMetaGlobalInput) *ActionVpsMigrateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsMigrateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsMigrateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsMigrateInvocation) Call() (*ActionVpsMigrateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionVpsMigrateInvocation) callAsBody() (*ActionVpsMigrateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsMigrateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsMigrateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsMigrateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsMigrateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsMigrateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsMigrateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionVpsMigrateInvocation) makeAllInputParams() *ActionVpsMigrateRequest {
	return &ActionVpsMigrateRequest{
		Vps:  inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsMigrateInvocation) makeInputParams() map[string]interface{} {
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
		if inv.IsParameterSelected("MaintenanceWindow") {
			ret["maintenance_window"] = inv.Input.MaintenanceWindow
		}
		if inv.IsParameterSelected("NoStart") {
			ret["no_start"] = inv.Input.NoStart
		}
		if inv.IsParameterSelected("Node") {
			if inv.IsParameterNil("Node") {
				ret["node"] = nil
			} else {
				ret["node"] = inv.Input.Node
			}
		}
		if inv.IsParameterSelected("Reason") {
			ret["reason"] = inv.Input.Reason
		}
		if inv.IsParameterSelected("ReplaceIpAddresses") {
			ret["replace_ip_addresses"] = inv.Input.ReplaceIpAddresses
		}
		if inv.IsParameterSelected("SendMail") {
			ret["send_mail"] = inv.Input.SendMail
		}
		if inv.IsParameterSelected("SkipStart") {
			ret["skip_start"] = inv.Input.SkipStart
		}
		if inv.IsParameterSelected("Swap") {
			ret["swap"] = inv.Input.Swap
		}
		if inv.IsParameterSelected("TransferIpAddresses") {
			ret["transfer_ip_addresses"] = inv.Input.TransferIpAddresses
		}
	}

	return ret
}

func (inv *ActionVpsMigrateInvocation) makeMetaInputParams() map[string]interface{} {
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
