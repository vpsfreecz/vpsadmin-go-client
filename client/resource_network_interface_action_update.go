package client

import (
	"strings"
)

// ActionNetworkInterfaceUpdate is a type for action Network_interface#Update
type ActionNetworkInterfaceUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionNetworkInterfaceUpdate(client *Client) *ActionNetworkInterfaceUpdate {
	return &ActionNetworkInterfaceUpdate{
		Client: client,
	}
}

// ActionNetworkInterfaceUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionNetworkInterfaceUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNetworkInterfaceUpdateMetaGlobalInput) SetIncludes(value string) *ActionNetworkInterfaceUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNetworkInterfaceUpdateMetaGlobalInput) SetNo(value bool) *ActionNetworkInterfaceUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkInterfaceUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkInterfaceUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionNetworkInterfaceUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNetworkInterfaceUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNetworkInterfaceUpdateInput is a type for action input parameters
type ActionNetworkInterfaceUpdateInput struct {
	Enable bool   `json:"enable"`
	MaxRx  int64  `json:"max_rx"`
	MaxTx  int64  `json:"max_tx"`
	Name   string `json:"name"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnable sets parameter Enable to value and selects it for sending
func (in *ActionNetworkInterfaceUpdateInput) SetEnable(value bool) *ActionNetworkInterfaceUpdateInput {
	in.Enable = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enable"] = nil
	return in
}

// SetMaxRx sets parameter MaxRx to value and selects it for sending
func (in *ActionNetworkInterfaceUpdateInput) SetMaxRx(value int64) *ActionNetworkInterfaceUpdateInput {
	in.MaxRx = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxRx"] = nil
	return in
}

// SetMaxTx sets parameter MaxTx to value and selects it for sending
func (in *ActionNetworkInterfaceUpdateInput) SetMaxTx(value int64) *ActionNetworkInterfaceUpdateInput {
	in.MaxTx = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxTx"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionNetworkInterfaceUpdateInput) SetName(value string) *ActionNetworkInterfaceUpdateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkInterfaceUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkInterfaceUpdateInput) SelectParameters(params ...string) *ActionNetworkInterfaceUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNetworkInterfaceUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNetworkInterfaceUpdateInput) UnselectParameters(params ...string) *ActionNetworkInterfaceUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNetworkInterfaceUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNetworkInterfaceUpdateRequest is a type for the entire action request
type ActionNetworkInterfaceUpdateRequest struct {
	NetworkInterface map[string]interface{} `json:"network_interface"`
	Meta             map[string]interface{} `json:"_meta"`
}

// ActionNetworkInterfaceUpdateOutput is a type for action output parameters
type ActionNetworkInterfaceUpdateOutput struct {
	Enable bool                 `json:"enable"`
	Id     int64                `json:"id"`
	Mac    string               `json:"mac"`
	MaxRx  int64                `json:"max_rx"`
	MaxTx  int64                `json:"max_tx"`
	Name   string               `json:"name"`
	Type   string               `json:"type"`
	Vps    *ActionVpsShowOutput `json:"vps"`
}

// ActionNetworkInterfaceUpdateMetaGlobalOutput is a type for global output metadata parameters
type ActionNetworkInterfaceUpdateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionNetworkInterfaceUpdateResponse struct {
	Action *ActionNetworkInterfaceUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NetworkInterface *ActionNetworkInterfaceUpdateOutput `json:"network_interface"`
		// Global output metadata
		Meta *ActionNetworkInterfaceUpdateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionNetworkInterfaceUpdateOutput
}

// Prepare the action for invocation
func (action *ActionNetworkInterfaceUpdate) Prepare() *ActionNetworkInterfaceUpdateInvocation {
	return &ActionNetworkInterfaceUpdateInvocation{
		Action: action,
		Path:   "/v7.0/network_interfaces/{network_interface_id}",
	}
}

// ActionNetworkInterfaceUpdateInvocation is used to configure action for invocation
type ActionNetworkInterfaceUpdateInvocation struct {
	// Pointer to the action
	Action *ActionNetworkInterfaceUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNetworkInterfaceUpdateInput
	// Global meta input parameters
	MetaInput *ActionNetworkInterfaceUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNetworkInterfaceUpdateInvocation) SetPathParamInt(param string, value int64) *ActionNetworkInterfaceUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNetworkInterfaceUpdateInvocation) SetPathParamString(param string, value string) *ActionNetworkInterfaceUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNetworkInterfaceUpdateInvocation) NewInput() *ActionNetworkInterfaceUpdateInput {
	inv.Input = &ActionNetworkInterfaceUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNetworkInterfaceUpdateInvocation) SetInput(input *ActionNetworkInterfaceUpdateInput) *ActionNetworkInterfaceUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNetworkInterfaceUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNetworkInterfaceUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNetworkInterfaceUpdateInvocation) NewMetaInput() *ActionNetworkInterfaceUpdateMetaGlobalInput {
	inv.MetaInput = &ActionNetworkInterfaceUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNetworkInterfaceUpdateInvocation) SetMetaInput(input *ActionNetworkInterfaceUpdateMetaGlobalInput) *ActionNetworkInterfaceUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNetworkInterfaceUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNetworkInterfaceUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNetworkInterfaceUpdateInvocation) Call() (*ActionNetworkInterfaceUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionNetworkInterfaceUpdateInvocation) callAsBody() (*ActionNetworkInterfaceUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNetworkInterfaceUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NetworkInterface
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionNetworkInterfaceUpdateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionNetworkInterfaceUpdateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionNetworkInterfaceUpdateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionNetworkInterfaceUpdateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionNetworkInterfaceUpdateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionNetworkInterfaceUpdateInvocation) makeAllInputParams() *ActionNetworkInterfaceUpdateRequest {
	return &ActionNetworkInterfaceUpdateRequest{
		NetworkInterface: inv.makeInputParams(),
		Meta:             inv.makeMetaInputParams(),
	}
}

func (inv *ActionNetworkInterfaceUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Enable") {
			ret["enable"] = inv.Input.Enable
		}
		if inv.IsParameterSelected("MaxRx") {
			ret["max_rx"] = inv.Input.MaxRx
		}
		if inv.IsParameterSelected("MaxTx") {
			ret["max_tx"] = inv.Input.MaxTx
		}
		if inv.IsParameterSelected("Name") {
			ret["name"] = inv.Input.Name
		}
	}

	return ret
}

func (inv *ActionNetworkInterfaceUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
