package client

import (
	"strings"
)

// ActionVpsMountUpdate is a type for action Vps.Mount#Update
type ActionVpsMountUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsMountUpdate(client *Client) *ActionVpsMountUpdate {
	return &ActionVpsMountUpdate{
		Client: client,
	}
}

// ActionVpsMountUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionVpsMountUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsMountUpdateMetaGlobalInput) SetNo(value bool) *ActionVpsMountUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsMountUpdateMetaGlobalInput) SetIncludes(value string) *ActionVpsMountUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsMountUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsMountUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionVpsMountUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsMountUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsMountUpdateInput is a type for action input parameters
type ActionVpsMountUpdateInput struct {
	OnStartFail string `json:"on_start_fail"`
	Enabled bool `json:"enabled"`
	MasterEnabled bool `json:"master_enabled"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOnStartFail sets parameter OnStartFail to value and selects it for sending
func (in *ActionVpsMountUpdateInput) SetOnStartFail(value string) *ActionVpsMountUpdateInput {
	in.OnStartFail = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OnStartFail"] = nil
	return in
}
// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionVpsMountUpdateInput) SetEnabled(value bool) *ActionVpsMountUpdateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}
// SetMasterEnabled sets parameter MasterEnabled to value and selects it for sending
func (in *ActionVpsMountUpdateInput) SetMasterEnabled(value bool) *ActionVpsMountUpdateInput {
	in.MasterEnabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MasterEnabled"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsMountUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsMountUpdateInput) SelectParameters(params ...string) *ActionVpsMountUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsMountUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsMountUpdateRequest is a type for the entire action request
type ActionVpsMountUpdateRequest struct {
	Mount map[string]interface{} `json:"mount"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionVpsMountUpdateOutput is a type for action output parameters
type ActionVpsMountUpdateOutput struct {
	Id int64 `json:"id"`
	Vps *ActionVpsShowOutput `json:"vps"`
	Dataset *ActionDatasetShowOutput `json:"dataset"`
	Snapshot *ActionDatasetSnapshotShowOutput `json:"snapshot"`
	UserNamespaceMap *ActionUserNamespaceMapShowOutput `json:"user_namespace_map"`
	Mountpoint string `json:"mountpoint"`
	Mode string `json:"mode"`
	OnStartFail string `json:"on_start_fail"`
	ExpirationDate string `json:"expiration_date"`
	Enabled bool `json:"enabled"`
	MasterEnabled bool `json:"master_enabled"`
	CurrentState string `json:"current_state"`
}

// ActionVpsMountUpdateMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsMountUpdateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsMountUpdateResponse struct {
	Action *ActionVpsMountUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Mount *ActionVpsMountUpdateOutput `json:"mount"`
		// Global output metadata
		Meta *ActionVpsMountUpdateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionVpsMountUpdateOutput
}


// Prepare the action for invocation
func (action *ActionVpsMountUpdate) Prepare() *ActionVpsMountUpdateInvocation {
	return &ActionVpsMountUpdateInvocation{
		Action: action,
		Path: "/v5.0/vpses/{vps_id}/mounts/{mount_id}",
	}
}

// ActionVpsMountUpdateInvocation is used to configure action for invocation
type ActionVpsMountUpdateInvocation struct {
	// Pointer to the action
	Action *ActionVpsMountUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsMountUpdateInput
	// Global meta input parameters
	MetaInput *ActionVpsMountUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsMountUpdateInvocation) SetPathParamInt(param string, value int64) *ActionVpsMountUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsMountUpdateInvocation) SetPathParamString(param string, value string) *ActionVpsMountUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsMountUpdateInvocation) NewInput() *ActionVpsMountUpdateInput {
	inv.Input = &ActionVpsMountUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsMountUpdateInvocation) SetInput(input *ActionVpsMountUpdateInput) *ActionVpsMountUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsMountUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsMountUpdateInvocation) NewMetaInput() *ActionVpsMountUpdateMetaGlobalInput {
	inv.MetaInput = &ActionVpsMountUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsMountUpdateInvocation) SetMetaInput(input *ActionVpsMountUpdateMetaGlobalInput) *ActionVpsMountUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsMountUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsMountUpdateInvocation) Call() (*ActionVpsMountUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionVpsMountUpdateInvocation) callAsBody() (*ActionVpsMountUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsMountUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Mount
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsMountUpdateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsMountUpdateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsMountUpdateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsMountUpdateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsMountUpdateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionVpsMountUpdateInvocation) makeAllInputParams() *ActionVpsMountUpdateRequest {
	return &ActionVpsMountUpdateRequest{
		Mount: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsMountUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("OnStartFail") {
			ret["on_start_fail"] = inv.Input.OnStartFail
		}
		if inv.IsParameterSelected("Enabled") {
			ret["enabled"] = inv.Input.Enabled
		}
		if inv.IsParameterSelected("MasterEnabled") {
			ret["master_enabled"] = inv.Input.MasterEnabled
		}
	}

	return ret
}

func (inv *ActionVpsMountUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
