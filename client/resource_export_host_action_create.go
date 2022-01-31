package client

import (
	"strings"
)

// ActionExportHostCreate is a type for action Export.Host#Create
type ActionExportHostCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionExportHostCreate(client *Client) *ActionExportHostCreate {
	return &ActionExportHostCreate{
		Client: client,
	}
}

// ActionExportHostCreateMetaGlobalInput is a type for action global meta input parameters
type ActionExportHostCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionExportHostCreateMetaGlobalInput) SetIncludes(value string) *ActionExportHostCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionExportHostCreateMetaGlobalInput) SetNo(value bool) *ActionExportHostCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionExportHostCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionExportHostCreateMetaGlobalInput) SelectParameters(params ...string) *ActionExportHostCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionExportHostCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionExportHostCreateInput is a type for action input parameters
type ActionExportHostCreateInput struct {
	IpAddress    int64 `json:"ip_address"`
	RootSquash   bool  `json:"root_squash"`
	Rw           bool  `json:"rw"`
	SubtreeCheck bool  `json:"subtree_check"`
	Sync         bool  `json:"sync"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIpAddress sets parameter IpAddress to value and selects it for sending
func (in *ActionExportHostCreateInput) SetIpAddress(value int64) *ActionExportHostCreateInput {
	in.IpAddress = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetIpAddressNil(false)
	in._selectedParameters["IpAddress"] = nil
	return in
}

// SetIpAddressNil sets parameter IpAddress to nil and selects it for sending
func (in *ActionExportHostCreateInput) SetIpAddressNil(set bool) *ActionExportHostCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["IpAddress"] = nil
		in.SelectParameters("IpAddress")
	} else {
		delete(in._nilParameters, "IpAddress")
	}
	return in
}

// SetRootSquash sets parameter RootSquash to value and selects it for sending
func (in *ActionExportHostCreateInput) SetRootSquash(value bool) *ActionExportHostCreateInput {
	in.RootSquash = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RootSquash"] = nil
	return in
}

// SetRw sets parameter Rw to value and selects it for sending
func (in *ActionExportHostCreateInput) SetRw(value bool) *ActionExportHostCreateInput {
	in.Rw = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Rw"] = nil
	return in
}

// SetSubtreeCheck sets parameter SubtreeCheck to value and selects it for sending
func (in *ActionExportHostCreateInput) SetSubtreeCheck(value bool) *ActionExportHostCreateInput {
	in.SubtreeCheck = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SubtreeCheck"] = nil
	return in
}

// SetSync sets parameter Sync to value and selects it for sending
func (in *ActionExportHostCreateInput) SetSync(value bool) *ActionExportHostCreateInput {
	in.Sync = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Sync"] = nil
	return in
}

// SelectParameters sets parameters from ActionExportHostCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionExportHostCreateInput) SelectParameters(params ...string) *ActionExportHostCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionExportHostCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionExportHostCreateInput) UnselectParameters(params ...string) *ActionExportHostCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionExportHostCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionExportHostCreateRequest is a type for the entire action request
type ActionExportHostCreateRequest struct {
	Host map[string]interface{} `json:"host"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionExportHostCreateOutput is a type for action output parameters
type ActionExportHostCreateOutput struct {
	Id           int64                      `json:"id"`
	IpAddress    *ActionIpAddressShowOutput `json:"ip_address"`
	RootSquash   bool                       `json:"root_squash"`
	Rw           bool                       `json:"rw"`
	SubtreeCheck bool                       `json:"subtree_check"`
	Sync         bool                       `json:"sync"`
}

// ActionExportHostCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionExportHostCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionExportHostCreateResponse struct {
	Action *ActionExportHostCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Host *ActionExportHostCreateOutput `json:"host"`
		// Global output metadata
		Meta *ActionExportHostCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionExportHostCreateOutput
}

// Prepare the action for invocation
func (action *ActionExportHostCreate) Prepare() *ActionExportHostCreateInvocation {
	return &ActionExportHostCreateInvocation{
		Action: action,
		Path:   "/v6.0/exports/{export_id}/hosts",
	}
}

// ActionExportHostCreateInvocation is used to configure action for invocation
type ActionExportHostCreateInvocation struct {
	// Pointer to the action
	Action *ActionExportHostCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionExportHostCreateInput
	// Global meta input parameters
	MetaInput *ActionExportHostCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionExportHostCreateInvocation) SetPathParamInt(param string, value int64) *ActionExportHostCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionExportHostCreateInvocation) SetPathParamString(param string, value string) *ActionExportHostCreateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionExportHostCreateInvocation) NewInput() *ActionExportHostCreateInput {
	inv.Input = &ActionExportHostCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionExportHostCreateInvocation) SetInput(input *ActionExportHostCreateInput) *ActionExportHostCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionExportHostCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionExportHostCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionExportHostCreateInvocation) NewMetaInput() *ActionExportHostCreateMetaGlobalInput {
	inv.MetaInput = &ActionExportHostCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionExportHostCreateInvocation) SetMetaInput(input *ActionExportHostCreateMetaGlobalInput) *ActionExportHostCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionExportHostCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionExportHostCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionExportHostCreateInvocation) Call() (*ActionExportHostCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionExportHostCreateInvocation) callAsBody() (*ActionExportHostCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionExportHostCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Host
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionExportHostCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionExportHostCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionExportHostCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionExportHostCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionExportHostCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionExportHostCreateInvocation) makeAllInputParams() *ActionExportHostCreateRequest {
	return &ActionExportHostCreateRequest{
		Host: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionExportHostCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("IpAddress") {
			if inv.IsParameterNil("IpAddress") {
				ret["ip_address"] = nil
			} else {
				ret["ip_address"] = inv.Input.IpAddress
			}
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
	}

	return ret
}

func (inv *ActionExportHostCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
