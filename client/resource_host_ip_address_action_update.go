package client

import (
	"strings"
)

// ActionHostIpAddressUpdate is a type for action Host_ip_address#Update
type ActionHostIpAddressUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionHostIpAddressUpdate(client *Client) *ActionHostIpAddressUpdate {
	return &ActionHostIpAddressUpdate{
		Client: client,
	}
}

// ActionHostIpAddressUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionHostIpAddressUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionHostIpAddressUpdateMetaGlobalInput) SetIncludes(value string) *ActionHostIpAddressUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionHostIpAddressUpdateMetaGlobalInput) SetNo(value bool) *ActionHostIpAddressUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionHostIpAddressUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionHostIpAddressUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionHostIpAddressUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionHostIpAddressUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionHostIpAddressUpdateInput is a type for action input parameters
type ActionHostIpAddressUpdateInput struct {
	Addr               string `json:"addr"`
	Assigned           bool   `json:"assigned"`
	IpAddress          int64  `json:"ip_address"`
	ReverseRecordValue string `json:"reverse_record_value"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAddr sets parameter Addr to value and selects it for sending
func (in *ActionHostIpAddressUpdateInput) SetAddr(value string) *ActionHostIpAddressUpdateInput {
	in.Addr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Addr"] = nil
	return in
}

// SetAssigned sets parameter Assigned to value and selects it for sending
func (in *ActionHostIpAddressUpdateInput) SetAssigned(value bool) *ActionHostIpAddressUpdateInput {
	in.Assigned = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Assigned"] = nil
	return in
}

// SetIpAddress sets parameter IpAddress to value and selects it for sending
func (in *ActionHostIpAddressUpdateInput) SetIpAddress(value int64) *ActionHostIpAddressUpdateInput {
	in.IpAddress = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetIpAddressNil(false)
	in._selectedParameters["IpAddress"] = nil
	return in
}

// SetIpAddressNil sets parameter IpAddress to nil and selects it for sending
func (in *ActionHostIpAddressUpdateInput) SetIpAddressNil(set bool) *ActionHostIpAddressUpdateInput {
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

// SetReverseRecordValue sets parameter ReverseRecordValue to value and selects it for sending
func (in *ActionHostIpAddressUpdateInput) SetReverseRecordValue(value string) *ActionHostIpAddressUpdateInput {
	in.ReverseRecordValue = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ReverseRecordValue"] = nil
	return in
}

// SelectParameters sets parameters from ActionHostIpAddressUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionHostIpAddressUpdateInput) SelectParameters(params ...string) *ActionHostIpAddressUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionHostIpAddressUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionHostIpAddressUpdateInput) UnselectParameters(params ...string) *ActionHostIpAddressUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionHostIpAddressUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionHostIpAddressUpdateRequest is a type for the entire action request
type ActionHostIpAddressUpdateRequest struct {
	HostIpAddress map[string]interface{} `json:"host_ip_address"`
	Meta          map[string]interface{} `json:"_meta"`
}

// ActionHostIpAddressUpdateOutput is a type for action output parameters
type ActionHostIpAddressUpdateOutput struct {
	Addr               string                     `json:"addr"`
	Assigned           bool                       `json:"assigned"`
	Id                 int64                      `json:"id"`
	IpAddress          *ActionIpAddressShowOutput `json:"ip_address"`
	ReverseRecordValue string                     `json:"reverse_record_value"`
	UserCreated        bool                       `json:"user_created"`
}

// ActionHostIpAddressUpdateMetaGlobalOutput is a type for global output metadata parameters
type ActionHostIpAddressUpdateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionHostIpAddressUpdateResponse struct {
	Action *ActionHostIpAddressUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		HostIpAddress *ActionHostIpAddressUpdateOutput `json:"host_ip_address"`
		// Global output metadata
		Meta *ActionHostIpAddressUpdateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionHostIpAddressUpdateOutput
}

// Prepare the action for invocation
func (action *ActionHostIpAddressUpdate) Prepare() *ActionHostIpAddressUpdateInvocation {
	return &ActionHostIpAddressUpdateInvocation{
		Action: action,
		Path:   "/v7.0/host_ip_addresses/{host_ip_address_id}",
	}
}

// ActionHostIpAddressUpdateInvocation is used to configure action for invocation
type ActionHostIpAddressUpdateInvocation struct {
	// Pointer to the action
	Action *ActionHostIpAddressUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionHostIpAddressUpdateInput
	// Global meta input parameters
	MetaInput *ActionHostIpAddressUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionHostIpAddressUpdateInvocation) SetPathParamInt(param string, value int64) *ActionHostIpAddressUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionHostIpAddressUpdateInvocation) SetPathParamString(param string, value string) *ActionHostIpAddressUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionHostIpAddressUpdateInvocation) NewInput() *ActionHostIpAddressUpdateInput {
	inv.Input = &ActionHostIpAddressUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionHostIpAddressUpdateInvocation) SetInput(input *ActionHostIpAddressUpdateInput) *ActionHostIpAddressUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionHostIpAddressUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionHostIpAddressUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionHostIpAddressUpdateInvocation) NewMetaInput() *ActionHostIpAddressUpdateMetaGlobalInput {
	inv.MetaInput = &ActionHostIpAddressUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionHostIpAddressUpdateInvocation) SetMetaInput(input *ActionHostIpAddressUpdateMetaGlobalInput) *ActionHostIpAddressUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionHostIpAddressUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionHostIpAddressUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionHostIpAddressUpdateInvocation) Call() (*ActionHostIpAddressUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionHostIpAddressUpdateInvocation) callAsBody() (*ActionHostIpAddressUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionHostIpAddressUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.HostIpAddress
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionHostIpAddressUpdateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionHostIpAddressUpdateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionHostIpAddressUpdateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionHostIpAddressUpdateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionHostIpAddressUpdateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionHostIpAddressUpdateInvocation) makeAllInputParams() *ActionHostIpAddressUpdateRequest {
	return &ActionHostIpAddressUpdateRequest{
		HostIpAddress: inv.makeInputParams(),
		Meta:          inv.makeMetaInputParams(),
	}
}

func (inv *ActionHostIpAddressUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Addr") {
			ret["addr"] = inv.Input.Addr
		}
		if inv.IsParameterSelected("Assigned") {
			ret["assigned"] = inv.Input.Assigned
		}
		if inv.IsParameterSelected("IpAddress") {
			if inv.IsParameterNil("IpAddress") {
				ret["ip_address"] = nil
			} else {
				ret["ip_address"] = inv.Input.IpAddress
			}
		}
		if inv.IsParameterSelected("ReverseRecordValue") {
			ret["reverse_record_value"] = inv.Input.ReverseRecordValue
		}
	}

	return ret
}

func (inv *ActionHostIpAddressUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
