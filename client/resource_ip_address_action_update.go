package client

import (
	"strings"
)

// ActionIpAddressUpdate is a type for action Ip_address#Update
type ActionIpAddressUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionIpAddressUpdate(client *Client) *ActionIpAddressUpdate {
	return &ActionIpAddressUpdate{
		Client: client,
	}
}

// ActionIpAddressUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionIpAddressUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpAddressUpdateMetaGlobalInput) SetIncludes(value string) *ActionIpAddressUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpAddressUpdateMetaGlobalInput) SetNo(value bool) *ActionIpAddressUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpAddressUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpAddressUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionIpAddressUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpAddressUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpAddressUpdateInput is a type for action input parameters
type ActionIpAddressUpdateInput struct {
	Environment int64 `json:"environment"`
	MaxRx       int64 `json:"max_rx"`
	MaxTx       int64 `json:"max_tx"`
	User        int64 `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionIpAddressUpdateInput) SetEnvironment(value int64) *ActionIpAddressUpdateInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEnvironmentNil(false)
	in._selectedParameters["Environment"] = nil
	return in
}

// SetEnvironmentNil sets parameter Environment to nil and selects it for sending
func (in *ActionIpAddressUpdateInput) SetEnvironmentNil(set bool) *ActionIpAddressUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Environment"] = nil
		in.SelectParameters("Environment")
	} else {
		delete(in._nilParameters, "Environment")
	}
	return in
}

// SetMaxRx sets parameter MaxRx to value and selects it for sending
func (in *ActionIpAddressUpdateInput) SetMaxRx(value int64) *ActionIpAddressUpdateInput {
	in.MaxRx = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxRx"] = nil
	return in
}

// SetMaxTx sets parameter MaxTx to value and selects it for sending
func (in *ActionIpAddressUpdateInput) SetMaxTx(value int64) *ActionIpAddressUpdateInput {
	in.MaxTx = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxTx"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionIpAddressUpdateInput) SetUser(value int64) *ActionIpAddressUpdateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionIpAddressUpdateInput) SetUserNil(set bool) *ActionIpAddressUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["User"] = nil
		in.SelectParameters("User")
	} else {
		delete(in._nilParameters, "User")
	}
	return in
}

// SelectParameters sets parameters from ActionIpAddressUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpAddressUpdateInput) SelectParameters(params ...string) *ActionIpAddressUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIpAddressUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIpAddressUpdateInput) UnselectParameters(params ...string) *ActionIpAddressUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIpAddressUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpAddressUpdateRequest is a type for the entire action request
type ActionIpAddressUpdateRequest struct {
	IpAddress map[string]interface{} `json:"ip_address"`
	Meta      map[string]interface{} `json:"_meta"`
}

// ActionIpAddressUpdateOutput is a type for action output parameters
type ActionIpAddressUpdateOutput struct {
	Addr               string                            `json:"addr"`
	ChargedEnvironment *ActionEnvironmentShowOutput      `json:"charged_environment"`
	ClassId            int64                             `json:"class_id"`
	Id                 int64                             `json:"id"`
	MaxRx              int64                             `json:"max_rx"`
	MaxTx              int64                             `json:"max_tx"`
	Network            *ActionNetworkShowOutput          `json:"network"`
	NetworkInterface   *ActionNetworkInterfaceShowOutput `json:"network_interface"`
	Prefix             int64                             `json:"prefix"`
	RouteVia           *ActionHostIpAddressShowOutput    `json:"route_via"`
	Size               int64                             `json:"size"`
	User               *ActionUserShowOutput             `json:"user"`
}

// ActionIpAddressUpdateMetaGlobalOutput is a type for global output metadata parameters
type ActionIpAddressUpdateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionIpAddressUpdateResponse struct {
	Action *ActionIpAddressUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpAddress *ActionIpAddressUpdateOutput `json:"ip_address"`
		// Global output metadata
		Meta *ActionIpAddressUpdateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionIpAddressUpdateOutput
}

// Prepare the action for invocation
func (action *ActionIpAddressUpdate) Prepare() *ActionIpAddressUpdateInvocation {
	return &ActionIpAddressUpdateInvocation{
		Action: action,
		Path:   "/v6.0/ip_addresses/{ip_address_id}",
	}
}

// ActionIpAddressUpdateInvocation is used to configure action for invocation
type ActionIpAddressUpdateInvocation struct {
	// Pointer to the action
	Action *ActionIpAddressUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpAddressUpdateInput
	// Global meta input parameters
	MetaInput *ActionIpAddressUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpAddressUpdateInvocation) SetPathParamInt(param string, value int64) *ActionIpAddressUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpAddressUpdateInvocation) SetPathParamString(param string, value string) *ActionIpAddressUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpAddressUpdateInvocation) NewInput() *ActionIpAddressUpdateInput {
	inv.Input = &ActionIpAddressUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpAddressUpdateInvocation) SetInput(input *ActionIpAddressUpdateInput) *ActionIpAddressUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpAddressUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIpAddressUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpAddressUpdateInvocation) NewMetaInput() *ActionIpAddressUpdateMetaGlobalInput {
	inv.MetaInput = &ActionIpAddressUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpAddressUpdateInvocation) SetMetaInput(input *ActionIpAddressUpdateMetaGlobalInput) *ActionIpAddressUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpAddressUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpAddressUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIpAddressUpdateInvocation) Call() (*ActionIpAddressUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionIpAddressUpdateInvocation) callAsBody() (*ActionIpAddressUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionIpAddressUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IpAddress
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionIpAddressUpdateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionIpAddressUpdateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionIpAddressUpdateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionIpAddressUpdateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionIpAddressUpdateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionIpAddressUpdateInvocation) makeAllInputParams() *ActionIpAddressUpdateRequest {
	return &ActionIpAddressUpdateRequest{
		IpAddress: inv.makeInputParams(),
		Meta:      inv.makeMetaInputParams(),
	}
}

func (inv *ActionIpAddressUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			if inv.IsParameterNil("Environment") {
				ret["environment"] = nil
			} else {
				ret["environment"] = inv.Input.Environment
			}
		}
		if inv.IsParameterSelected("MaxRx") {
			ret["max_rx"] = inv.Input.MaxRx
		}
		if inv.IsParameterSelected("MaxTx") {
			ret["max_tx"] = inv.Input.MaxTx
		}
		if inv.IsParameterSelected("User") {
			if inv.IsParameterNil("User") {
				ret["user"] = nil
			} else {
				ret["user"] = inv.Input.User
			}
		}
	}

	return ret
}

func (inv *ActionIpAddressUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
