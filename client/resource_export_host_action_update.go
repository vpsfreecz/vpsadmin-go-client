package client

import (
	"strings"
)

// ActionExportHostUpdate is a type for action Export.Host#Update
type ActionExportHostUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionExportHostUpdate(client *Client) *ActionExportHostUpdate {
	return &ActionExportHostUpdate{
		Client: client,
	}
}

// ActionExportHostUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionExportHostUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionExportHostUpdateMetaGlobalInput) SetIncludes(value string) *ActionExportHostUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionExportHostUpdateMetaGlobalInput) SetNo(value bool) *ActionExportHostUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionExportHostUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionExportHostUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionExportHostUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionExportHostUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionExportHostUpdateInput is a type for action input parameters
type ActionExportHostUpdateInput struct {
	RootSquash   bool `json:"root_squash"`
	Rw           bool `json:"rw"`
	SubtreeCheck bool `json:"subtree_check"`
	Sync         bool `json:"sync"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetRootSquash sets parameter RootSquash to value and selects it for sending
func (in *ActionExportHostUpdateInput) SetRootSquash(value bool) *ActionExportHostUpdateInput {
	in.RootSquash = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RootSquash"] = nil
	return in
}

// SetRw sets parameter Rw to value and selects it for sending
func (in *ActionExportHostUpdateInput) SetRw(value bool) *ActionExportHostUpdateInput {
	in.Rw = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Rw"] = nil
	return in
}

// SetSubtreeCheck sets parameter SubtreeCheck to value and selects it for sending
func (in *ActionExportHostUpdateInput) SetSubtreeCheck(value bool) *ActionExportHostUpdateInput {
	in.SubtreeCheck = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SubtreeCheck"] = nil
	return in
}

// SetSync sets parameter Sync to value and selects it for sending
func (in *ActionExportHostUpdateInput) SetSync(value bool) *ActionExportHostUpdateInput {
	in.Sync = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Sync"] = nil
	return in
}

// SelectParameters sets parameters from ActionExportHostUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionExportHostUpdateInput) SelectParameters(params ...string) *ActionExportHostUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionExportHostUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionExportHostUpdateInput) UnselectParameters(params ...string) *ActionExportHostUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionExportHostUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionExportHostUpdateRequest is a type for the entire action request
type ActionExportHostUpdateRequest struct {
	Host map[string]interface{} `json:"host"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionExportHostUpdateOutput is a type for action output parameters
type ActionExportHostUpdateOutput struct {
	Id           int64                      `json:"id"`
	IpAddress    *ActionIpAddressShowOutput `json:"ip_address"`
	RootSquash   bool                       `json:"root_squash"`
	Rw           bool                       `json:"rw"`
	SubtreeCheck bool                       `json:"subtree_check"`
	Sync         bool                       `json:"sync"`
}

// ActionExportHostUpdateMetaGlobalOutput is a type for global output metadata parameters
type ActionExportHostUpdateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionExportHostUpdateResponse struct {
	Action *ActionExportHostUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Host *ActionExportHostUpdateOutput `json:"host"`
		// Global output metadata
		Meta *ActionExportHostUpdateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionExportHostUpdateOutput
}

// Prepare the action for invocation
func (action *ActionExportHostUpdate) Prepare() *ActionExportHostUpdateInvocation {
	return &ActionExportHostUpdateInvocation{
		Action: action,
		Path:   "/v7.0/exports/{export_id}/hosts/{host_id}",
	}
}

// ActionExportHostUpdateInvocation is used to configure action for invocation
type ActionExportHostUpdateInvocation struct {
	// Pointer to the action
	Action *ActionExportHostUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionExportHostUpdateInput
	// Global meta input parameters
	MetaInput *ActionExportHostUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionExportHostUpdateInvocation) SetPathParamInt(param string, value int64) *ActionExportHostUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionExportHostUpdateInvocation) SetPathParamString(param string, value string) *ActionExportHostUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionExportHostUpdateInvocation) NewInput() *ActionExportHostUpdateInput {
	inv.Input = &ActionExportHostUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionExportHostUpdateInvocation) SetInput(input *ActionExportHostUpdateInput) *ActionExportHostUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionExportHostUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionExportHostUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionExportHostUpdateInvocation) NewMetaInput() *ActionExportHostUpdateMetaGlobalInput {
	inv.MetaInput = &ActionExportHostUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionExportHostUpdateInvocation) SetMetaInput(input *ActionExportHostUpdateMetaGlobalInput) *ActionExportHostUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionExportHostUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionExportHostUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionExportHostUpdateInvocation) Call() (*ActionExportHostUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionExportHostUpdateInvocation) callAsBody() (*ActionExportHostUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionExportHostUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Host
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionExportHostUpdateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionExportHostUpdateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionExportHostUpdateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionExportHostUpdateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionExportHostUpdateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionExportHostUpdateInvocation) makeAllInputParams() *ActionExportHostUpdateRequest {
	return &ActionExportHostUpdateRequest{
		Host: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionExportHostUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
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
	}

	return ret
}

func (inv *ActionExportHostUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
