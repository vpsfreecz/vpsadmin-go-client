package client

import (
	"strings"
)

// ActionVpsBoot is a type for action Vps#Boot
type ActionVpsBoot struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsBoot(client *Client) *ActionVpsBoot {
	return &ActionVpsBoot{
		Client: client,
	}
}

// ActionVpsBootMetaGlobalInput is a type for action global meta input parameters
type ActionVpsBootMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsBootMetaGlobalInput) SetIncludes(value string) *ActionVpsBootMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsBootMetaGlobalInput) SetNo(value bool) *ActionVpsBootMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsBootMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsBootMetaGlobalInput) SelectParameters(params ...string) *ActionVpsBootMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsBootMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsBootInput is a type for action input parameters
type ActionVpsBootInput struct {
	MountRootDataset string `json:"mount_root_dataset"`
	OsTemplate       int64  `json:"os_template"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetMountRootDataset sets parameter MountRootDataset to value and selects it for sending
func (in *ActionVpsBootInput) SetMountRootDataset(value string) *ActionVpsBootInput {
	in.MountRootDataset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MountRootDataset"] = nil
	return in
}

// SetOsTemplate sets parameter OsTemplate to value and selects it for sending
func (in *ActionVpsBootInput) SetOsTemplate(value int64) *ActionVpsBootInput {
	in.OsTemplate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OsTemplate"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsBootInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsBootInput) SelectParameters(params ...string) *ActionVpsBootInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsBootInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsBootRequest is a type for the entire action request
type ActionVpsBootRequest struct {
	Vps  map[string]interface{} `json:"vps"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionVpsBootMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsBootMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsBootResponse struct {
	Action *ActionVpsBoot `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionVpsBootMetaGlobalOutput `json:"_meta"`
	}
}

// Prepare the action for invocation
func (action *ActionVpsBoot) Prepare() *ActionVpsBootInvocation {
	return &ActionVpsBootInvocation{
		Action: action,
		Path:   "/v6.0/vpses/{vps_id}/boot",
	}
}

// ActionVpsBootInvocation is used to configure action for invocation
type ActionVpsBootInvocation struct {
	// Pointer to the action
	Action *ActionVpsBoot

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsBootInput
	// Global meta input parameters
	MetaInput *ActionVpsBootMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsBootInvocation) SetPathParamInt(param string, value int64) *ActionVpsBootInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsBootInvocation) SetPathParamString(param string, value string) *ActionVpsBootInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsBootInvocation) NewInput() *ActionVpsBootInput {
	inv.Input = &ActionVpsBootInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsBootInvocation) SetInput(input *ActionVpsBootInput) *ActionVpsBootInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsBootInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsBootInvocation) NewMetaInput() *ActionVpsBootMetaGlobalInput {
	inv.MetaInput = &ActionVpsBootMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsBootInvocation) SetMetaInput(input *ActionVpsBootMetaGlobalInput) *ActionVpsBootInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsBootInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsBootInvocation) Call() (*ActionVpsBootResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionVpsBootInvocation) callAsBody() (*ActionVpsBootResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsBootResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsBootResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsBootResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsBootResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsBootResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsBootResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionVpsBootInvocation) makeAllInputParams() *ActionVpsBootRequest {
	return &ActionVpsBootRequest{
		Vps:  inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsBootInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("MountRootDataset") {
			ret["mount_root_dataset"] = inv.Input.MountRootDataset
		}
		if inv.IsParameterSelected("OsTemplate") {
			ret["os_template"] = inv.Input.OsTemplate
		}
	}

	return ret
}

func (inv *ActionVpsBootInvocation) makeMetaInputParams() map[string]interface{} {
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
