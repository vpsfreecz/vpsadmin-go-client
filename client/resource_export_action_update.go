package client

import (
	"strings"
)

// ActionExportUpdate is a type for action Export#Update
type ActionExportUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionExportUpdate(client *Client) *ActionExportUpdate {
	return &ActionExportUpdate{
		Client: client,
	}
}

// ActionExportUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionExportUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionExportUpdateMetaGlobalInput) SetIncludes(value string) *ActionExportUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionExportUpdateMetaGlobalInput) SetNo(value bool) *ActionExportUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionExportUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionExportUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionExportUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionExportUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionExportUpdateInput is a type for action input parameters
type ActionExportUpdateInput struct {
	AllVps bool `json:"all_vps"`
	Enabled bool `json:"enabled"`
	RootSquash bool `json:"root_squash"`
	Rw bool `json:"rw"`
	SubtreeCheck bool `json:"subtree_check"`
	Sync bool `json:"sync"`
	Threads int64 `json:"threads"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetAllVps sets parameter AllVps to value and selects it for sending
func (in *ActionExportUpdateInput) SetAllVps(value bool) *ActionExportUpdateInput {
	in.AllVps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AllVps"] = nil
	return in
}
// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionExportUpdateInput) SetEnabled(value bool) *ActionExportUpdateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}
// SetRootSquash sets parameter RootSquash to value and selects it for sending
func (in *ActionExportUpdateInput) SetRootSquash(value bool) *ActionExportUpdateInput {
	in.RootSquash = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RootSquash"] = nil
	return in
}
// SetRw sets parameter Rw to value and selects it for sending
func (in *ActionExportUpdateInput) SetRw(value bool) *ActionExportUpdateInput {
	in.Rw = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Rw"] = nil
	return in
}
// SetSubtreeCheck sets parameter SubtreeCheck to value and selects it for sending
func (in *ActionExportUpdateInput) SetSubtreeCheck(value bool) *ActionExportUpdateInput {
	in.SubtreeCheck = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SubtreeCheck"] = nil
	return in
}
// SetSync sets parameter Sync to value and selects it for sending
func (in *ActionExportUpdateInput) SetSync(value bool) *ActionExportUpdateInput {
	in.Sync = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Sync"] = nil
	return in
}
// SetThreads sets parameter Threads to value and selects it for sending
func (in *ActionExportUpdateInput) SetThreads(value int64) *ActionExportUpdateInput {
	in.Threads = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Threads"] = nil
	return in
}

// SelectParameters sets parameters from ActionExportUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionExportUpdateInput) SelectParameters(params ...string) *ActionExportUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionExportUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionExportUpdateRequest is a type for the entire action request
type ActionExportUpdateRequest struct {
	Export map[string]interface{} `json:"export"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionExportUpdateOutput is a type for action output parameters
type ActionExportUpdateOutput struct {
	AllVps bool `json:"all_vps"`
	CreatedAt string `json:"created_at"`
	Dataset *ActionDatasetShowOutput `json:"dataset"`
	Enabled bool `json:"enabled"`
	ExpirationDate string `json:"expiration_date"`
	HostIpAddress *ActionHostIpAddressShowOutput `json:"host_ip_address"`
	Id int64 `json:"id"`
	IpAddress *ActionIpAddressShowOutput `json:"ip_address"`
	Path string `json:"path"`
	RootSquash bool `json:"root_squash"`
	Rw bool `json:"rw"`
	Snapshot *ActionDatasetSnapshotShowOutput `json:"snapshot"`
	SubtreeCheck bool `json:"subtree_check"`
	Sync bool `json:"sync"`
	Threads int64 `json:"threads"`
	UpdatedAt string `json:"updated_at"`
	User *ActionUserShowOutput `json:"user"`
}

// ActionExportUpdateMetaGlobalOutput is a type for global output metadata parameters
type ActionExportUpdateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionExportUpdateResponse struct {
	Action *ActionExportUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Export *ActionExportUpdateOutput `json:"export"`
		// Global output metadata
		Meta *ActionExportUpdateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionExportUpdateOutput
}


// Prepare the action for invocation
func (action *ActionExportUpdate) Prepare() *ActionExportUpdateInvocation {
	return &ActionExportUpdateInvocation{
		Action: action,
		Path: "/v6.0/exports/{export_id}",
	}
}

// ActionExportUpdateInvocation is used to configure action for invocation
type ActionExportUpdateInvocation struct {
	// Pointer to the action
	Action *ActionExportUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionExportUpdateInput
	// Global meta input parameters
	MetaInput *ActionExportUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionExportUpdateInvocation) SetPathParamInt(param string, value int64) *ActionExportUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionExportUpdateInvocation) SetPathParamString(param string, value string) *ActionExportUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionExportUpdateInvocation) NewInput() *ActionExportUpdateInput {
	inv.Input = &ActionExportUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionExportUpdateInvocation) SetInput(input *ActionExportUpdateInput) *ActionExportUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionExportUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionExportUpdateInvocation) NewMetaInput() *ActionExportUpdateMetaGlobalInput {
	inv.MetaInput = &ActionExportUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionExportUpdateInvocation) SetMetaInput(input *ActionExportUpdateMetaGlobalInput) *ActionExportUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionExportUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionExportUpdateInvocation) Call() (*ActionExportUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionExportUpdateInvocation) callAsBody() (*ActionExportUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionExportUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Export
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionExportUpdateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionExportUpdateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionExportUpdateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionExportUpdateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionExportUpdateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionExportUpdateInvocation) makeAllInputParams() *ActionExportUpdateRequest {
	return &ActionExportUpdateRequest{
		Export: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionExportUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("AllVps") {
			ret["all_vps"] = inv.Input.AllVps
		}
		if inv.IsParameterSelected("Enabled") {
			ret["enabled"] = inv.Input.Enabled
		}
		if inv.IsParameterSelected("RootSquash") {
			ret["root_squash"] = inv.Input.RootSquash
		}
		if inv.IsParameterSelected("Rw") {
			ret["rw"] = inv.Input.Rw
		}
		if inv.IsParameterSelected("SubtreeCheck") {
			ret["subtree_check"] = inv.Input.SubtreeCheck
		}
		if inv.IsParameterSelected("Sync") {
			ret["sync"] = inv.Input.Sync
		}
		if inv.IsParameterSelected("Threads") {
			ret["threads"] = inv.Input.Threads
		}
	}

	return ret
}

func (inv *ActionExportUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
