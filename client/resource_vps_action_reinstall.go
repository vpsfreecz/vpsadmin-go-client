package client

import (
	"strings"
)

// ActionVpsReinstall is a type for action Vps#Reinstall
type ActionVpsReinstall struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsReinstall(client *Client) *ActionVpsReinstall {
	return &ActionVpsReinstall{
		Client: client,
	}
}

// ActionVpsReinstallMetaGlobalInput is a type for action global meta input parameters
type ActionVpsReinstallMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsReinstallMetaGlobalInput) SetIncludes(value string) *ActionVpsReinstallMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsReinstallMetaGlobalInput) SetNo(value bool) *ActionVpsReinstallMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsReinstallMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsReinstallMetaGlobalInput) SelectParameters(params ...string) *ActionVpsReinstallMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsReinstallMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsReinstallInput is a type for action input parameters
type ActionVpsReinstallInput struct {
	OsTemplate      int64  `json:"os_template"`
	UserDataContent string `json:"user_data_content"`
	UserDataFormat  string `json:"user_data_format"`
	VpsUserData     int64  `json:"vps_user_data"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetOsTemplate sets parameter OsTemplate to value and selects it for sending
func (in *ActionVpsReinstallInput) SetOsTemplate(value int64) *ActionVpsReinstallInput {
	in.OsTemplate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetOsTemplateNil(false)
	in._selectedParameters["OsTemplate"] = nil
	return in
}

// SetOsTemplateNil sets parameter OsTemplate to nil and selects it for sending
func (in *ActionVpsReinstallInput) SetOsTemplateNil(set bool) *ActionVpsReinstallInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["OsTemplate"] = nil
		in.SelectParameters("OsTemplate")
	} else {
		delete(in._nilParameters, "OsTemplate")
	}
	return in
}

// SetUserDataContent sets parameter UserDataContent to value and selects it for sending
func (in *ActionVpsReinstallInput) SetUserDataContent(value string) *ActionVpsReinstallInput {
	in.UserDataContent = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserDataContent"] = nil
	return in
}

// SetUserDataFormat sets parameter UserDataFormat to value and selects it for sending
func (in *ActionVpsReinstallInput) SetUserDataFormat(value string) *ActionVpsReinstallInput {
	in.UserDataFormat = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserDataFormat"] = nil
	return in
}

// SetVpsUserData sets parameter VpsUserData to value and selects it for sending
func (in *ActionVpsReinstallInput) SetVpsUserData(value int64) *ActionVpsReinstallInput {
	in.VpsUserData = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetVpsUserDataNil(false)
	in._selectedParameters["VpsUserData"] = nil
	return in
}

// SetVpsUserDataNil sets parameter VpsUserData to nil and selects it for sending
func (in *ActionVpsReinstallInput) SetVpsUserDataNil(set bool) *ActionVpsReinstallInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["VpsUserData"] = nil
		in.SelectParameters("VpsUserData")
	} else {
		delete(in._nilParameters, "VpsUserData")
	}
	return in
}

// SelectParameters sets parameters from ActionVpsReinstallInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsReinstallInput) SelectParameters(params ...string) *ActionVpsReinstallInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsReinstallInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsReinstallInput) UnselectParameters(params ...string) *ActionVpsReinstallInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsReinstallInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsReinstallRequest is a type for the entire action request
type ActionVpsReinstallRequest struct {
	Vps  map[string]interface{} `json:"vps"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionVpsReinstallMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsReinstallMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsReinstallResponse struct {
	Action *ActionVpsReinstall `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionVpsReinstallMetaGlobalOutput `json:"_meta"`
	}
}

// Prepare the action for invocation
func (action *ActionVpsReinstall) Prepare() *ActionVpsReinstallInvocation {
	return &ActionVpsReinstallInvocation{
		Action: action,
		Path:   "/v7.0/vpses/{vps_id}/reinstall",
	}
}

// ActionVpsReinstallInvocation is used to configure action for invocation
type ActionVpsReinstallInvocation struct {
	// Pointer to the action
	Action *ActionVpsReinstall

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsReinstallInput
	// Global meta input parameters
	MetaInput *ActionVpsReinstallMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsReinstallInvocation) SetPathParamInt(param string, value int64) *ActionVpsReinstallInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsReinstallInvocation) SetPathParamString(param string, value string) *ActionVpsReinstallInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsReinstallInvocation) NewInput() *ActionVpsReinstallInput {
	inv.Input = &ActionVpsReinstallInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsReinstallInvocation) SetInput(input *ActionVpsReinstallInput) *ActionVpsReinstallInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsReinstallInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsReinstallInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsReinstallInvocation) NewMetaInput() *ActionVpsReinstallMetaGlobalInput {
	inv.MetaInput = &ActionVpsReinstallMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsReinstallInvocation) SetMetaInput(input *ActionVpsReinstallMetaGlobalInput) *ActionVpsReinstallInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsReinstallInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsReinstallInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsReinstallInvocation) Call() (*ActionVpsReinstallResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionVpsReinstallInvocation) callAsBody() (*ActionVpsReinstallResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsReinstallResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsReinstallResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsReinstallResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsReinstallResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsReinstallResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsReinstallResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionVpsReinstallInvocation) makeAllInputParams() *ActionVpsReinstallRequest {
	return &ActionVpsReinstallRequest{
		Vps:  inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsReinstallInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("OsTemplate") {
			if inv.IsParameterNil("OsTemplate") {
				ret["os_template"] = nil
			} else {
				ret["os_template"] = inv.Input.OsTemplate
			}
		}
		if inv.IsParameterSelected("UserDataContent") {
			ret["user_data_content"] = inv.Input.UserDataContent
		}
		if inv.IsParameterSelected("UserDataFormat") {
			ret["user_data_format"] = inv.Input.UserDataFormat
		}
		if inv.IsParameterSelected("VpsUserData") {
			if inv.IsParameterNil("VpsUserData") {
				ret["vps_user_data"] = nil
			} else {
				ret["vps_user_data"] = inv.Input.VpsUserData
			}
		}
	}

	return ret
}

func (inv *ActionVpsReinstallInvocation) makeMetaInputParams() map[string]interface{} {
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
