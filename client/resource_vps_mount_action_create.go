package client

import (
	"strings"
)

// ActionVpsMountCreate is a type for action Vps.Mount#Create
type ActionVpsMountCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsMountCreate(client *Client) *ActionVpsMountCreate {
	return &ActionVpsMountCreate{
		Client: client,
	}
}

// ActionVpsMountCreateMetaGlobalInput is a type for action global meta input parameters
type ActionVpsMountCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsMountCreateMetaGlobalInput) SetIncludes(value string) *ActionVpsMountCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsMountCreateMetaGlobalInput) SetNo(value bool) *ActionVpsMountCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsMountCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsMountCreateMetaGlobalInput) SelectParameters(params ...string) *ActionVpsMountCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsMountCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsMountCreateInput is a type for action input parameters
type ActionVpsMountCreateInput struct {
	Dataset          int64  `json:"dataset"`
	Enabled          bool   `json:"enabled"`
	Mode             string `json:"mode"`
	Mountpoint       string `json:"mountpoint"`
	OnStartFail      string `json:"on_start_fail"`
	UserNamespaceMap int64  `json:"user_namespace_map"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDataset sets parameter Dataset to value and selects it for sending
func (in *ActionVpsMountCreateInput) SetDataset(value int64) *ActionVpsMountCreateInput {
	in.Dataset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetDatasetNil(false)
	in._selectedParameters["Dataset"] = nil
	return in
}

// SetDatasetNil sets parameter Dataset to nil and selects it for sending
func (in *ActionVpsMountCreateInput) SetDatasetNil(set bool) *ActionVpsMountCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Dataset"] = nil
		in.SelectParameters("Dataset")
	} else {
		delete(in._nilParameters, "Dataset")
	}
	return in
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionVpsMountCreateInput) SetEnabled(value bool) *ActionVpsMountCreateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetMode sets parameter Mode to value and selects it for sending
func (in *ActionVpsMountCreateInput) SetMode(value string) *ActionVpsMountCreateInput {
	in.Mode = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Mode"] = nil
	return in
}

// SetMountpoint sets parameter Mountpoint to value and selects it for sending
func (in *ActionVpsMountCreateInput) SetMountpoint(value string) *ActionVpsMountCreateInput {
	in.Mountpoint = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Mountpoint"] = nil
	return in
}

// SetOnStartFail sets parameter OnStartFail to value and selects it for sending
func (in *ActionVpsMountCreateInput) SetOnStartFail(value string) *ActionVpsMountCreateInput {
	in.OnStartFail = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OnStartFail"] = nil
	return in
}

// SetUserNamespaceMap sets parameter UserNamespaceMap to value and selects it for sending
func (in *ActionVpsMountCreateInput) SetUserNamespaceMap(value int64) *ActionVpsMountCreateInput {
	in.UserNamespaceMap = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNamespaceMapNil(false)
	in._selectedParameters["UserNamespaceMap"] = nil
	return in
}

// SetUserNamespaceMapNil sets parameter UserNamespaceMap to nil and selects it for sending
func (in *ActionVpsMountCreateInput) SetUserNamespaceMapNil(set bool) *ActionVpsMountCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["UserNamespaceMap"] = nil
		in.SelectParameters("UserNamespaceMap")
	} else {
		delete(in._nilParameters, "UserNamespaceMap")
	}
	return in
}

// SelectParameters sets parameters from ActionVpsMountCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsMountCreateInput) SelectParameters(params ...string) *ActionVpsMountCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsMountCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsMountCreateInput) UnselectParameters(params ...string) *ActionVpsMountCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsMountCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsMountCreateRequest is a type for the entire action request
type ActionVpsMountCreateRequest struct {
	Mount map[string]interface{} `json:"mount"`
	Meta  map[string]interface{} `json:"_meta"`
}

// ActionVpsMountCreateOutput is a type for action output parameters
type ActionVpsMountCreateOutput struct {
	CurrentState     string                            `json:"current_state"`
	Dataset          *ActionDatasetShowOutput          `json:"dataset"`
	Enabled          bool                              `json:"enabled"`
	ExpirationDate   string                            `json:"expiration_date"`
	Id               int64                             `json:"id"`
	MasterEnabled    bool                              `json:"master_enabled"`
	Mode             string                            `json:"mode"`
	Mountpoint       string                            `json:"mountpoint"`
	OnStartFail      string                            `json:"on_start_fail"`
	UserNamespaceMap *ActionUserNamespaceMapShowOutput `json:"user_namespace_map"`
	Vps              *ActionVpsShowOutput              `json:"vps"`
}

// ActionVpsMountCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsMountCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsMountCreateResponse struct {
	Action *ActionVpsMountCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Mount *ActionVpsMountCreateOutput `json:"mount"`
		// Global output metadata
		Meta *ActionVpsMountCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionVpsMountCreateOutput
}

// Prepare the action for invocation
func (action *ActionVpsMountCreate) Prepare() *ActionVpsMountCreateInvocation {
	return &ActionVpsMountCreateInvocation{
		Action: action,
		Path:   "/v7.0/vpses/{vps_id}/mounts",
	}
}

// ActionVpsMountCreateInvocation is used to configure action for invocation
type ActionVpsMountCreateInvocation struct {
	// Pointer to the action
	Action *ActionVpsMountCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsMountCreateInput
	// Global meta input parameters
	MetaInput *ActionVpsMountCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsMountCreateInvocation) SetPathParamInt(param string, value int64) *ActionVpsMountCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsMountCreateInvocation) SetPathParamString(param string, value string) *ActionVpsMountCreateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsMountCreateInvocation) NewInput() *ActionVpsMountCreateInput {
	inv.Input = &ActionVpsMountCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsMountCreateInvocation) SetInput(input *ActionVpsMountCreateInput) *ActionVpsMountCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsMountCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsMountCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsMountCreateInvocation) NewMetaInput() *ActionVpsMountCreateMetaGlobalInput {
	inv.MetaInput = &ActionVpsMountCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsMountCreateInvocation) SetMetaInput(input *ActionVpsMountCreateMetaGlobalInput) *ActionVpsMountCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsMountCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsMountCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsMountCreateInvocation) Call() (*ActionVpsMountCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionVpsMountCreateInvocation) callAsBody() (*ActionVpsMountCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsMountCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Mount
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsMountCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsMountCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsMountCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsMountCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsMountCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionVpsMountCreateInvocation) makeAllInputParams() *ActionVpsMountCreateRequest {
	return &ActionVpsMountCreateRequest{
		Mount: inv.makeInputParams(),
		Meta:  inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsMountCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Dataset") {
			if inv.IsParameterNil("Dataset") {
				ret["dataset"] = nil
			} else {
				ret["dataset"] = inv.Input.Dataset
			}
		}
		if inv.IsParameterSelected("Enabled") {
			ret["enabled"] = inv.Input.Enabled
		}
		if inv.IsParameterSelected("Mode") {
			ret["mode"] = inv.Input.Mode
		}
		if inv.IsParameterSelected("Mountpoint") {
			ret["mountpoint"] = inv.Input.Mountpoint
		}
		if inv.IsParameterSelected("OnStartFail") {
			ret["on_start_fail"] = inv.Input.OnStartFail
		}
		if inv.IsParameterSelected("UserNamespaceMap") {
			if inv.IsParameterNil("UserNamespaceMap") {
				ret["user_namespace_map"] = nil
			} else {
				ret["user_namespace_map"] = inv.Input.UserNamespaceMap
			}
		}
	}

	return ret
}

func (inv *ActionVpsMountCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
