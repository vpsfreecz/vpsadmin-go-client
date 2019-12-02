package client

import (
)

// ActionNetworkCreate is a type for action Network#Create
type ActionNetworkCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionNetworkCreate(client *Client) *ActionNetworkCreate {
	return &ActionNetworkCreate{
		Client: client,
	}
}

// ActionNetworkCreateMetaGlobalInput is a type for action global meta input parameters
type ActionNetworkCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNetworkCreateMetaGlobalInput) SetNo(value bool) *ActionNetworkCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNetworkCreateMetaGlobalInput) SetIncludes(value string) *ActionNetworkCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkCreateMetaGlobalInput) SelectParameters(params ...string) *ActionNetworkCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNetworkCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNetworkCreateInput is a type for action input parameters
type ActionNetworkCreateInput struct {
	Label string `json:"label"`
	Location int64 `json:"location"`
	IpVersion int64 `json:"ip_version"`
	Address string `json:"address"`
	Prefix int64 `json:"prefix"`
	Role string `json:"role"`
	Managed bool `json:"managed"`
	SplitAccess string `json:"split_access"`
	SplitPrefix int64 `json:"split_prefix"`
	Autopick bool `json:"autopick"`
	Purpose string `json:"purpose"`
	AddIpAddresses bool `json:"add_ip_addresses"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionNetworkCreateInput) SetLabel(value string) *ActionNetworkCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}
// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionNetworkCreateInput) SetLocation(value int64) *ActionNetworkCreateInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}
// SetIpVersion sets parameter IpVersion to value and selects it for sending
func (in *ActionNetworkCreateInput) SetIpVersion(value int64) *ActionNetworkCreateInput {
	in.IpVersion = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IpVersion"] = nil
	return in
}
// SetAddress sets parameter Address to value and selects it for sending
func (in *ActionNetworkCreateInput) SetAddress(value string) *ActionNetworkCreateInput {
	in.Address = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Address"] = nil
	return in
}
// SetPrefix sets parameter Prefix to value and selects it for sending
func (in *ActionNetworkCreateInput) SetPrefix(value int64) *ActionNetworkCreateInput {
	in.Prefix = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Prefix"] = nil
	return in
}
// SetRole sets parameter Role to value and selects it for sending
func (in *ActionNetworkCreateInput) SetRole(value string) *ActionNetworkCreateInput {
	in.Role = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Role"] = nil
	return in
}
// SetManaged sets parameter Managed to value and selects it for sending
func (in *ActionNetworkCreateInput) SetManaged(value bool) *ActionNetworkCreateInput {
	in.Managed = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Managed"] = nil
	return in
}
// SetSplitAccess sets parameter SplitAccess to value and selects it for sending
func (in *ActionNetworkCreateInput) SetSplitAccess(value string) *ActionNetworkCreateInput {
	in.SplitAccess = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SplitAccess"] = nil
	return in
}
// SetSplitPrefix sets parameter SplitPrefix to value and selects it for sending
func (in *ActionNetworkCreateInput) SetSplitPrefix(value int64) *ActionNetworkCreateInput {
	in.SplitPrefix = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SplitPrefix"] = nil
	return in
}
// SetAutopick sets parameter Autopick to value and selects it for sending
func (in *ActionNetworkCreateInput) SetAutopick(value bool) *ActionNetworkCreateInput {
	in.Autopick = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Autopick"] = nil
	return in
}
// SetPurpose sets parameter Purpose to value and selects it for sending
func (in *ActionNetworkCreateInput) SetPurpose(value string) *ActionNetworkCreateInput {
	in.Purpose = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Purpose"] = nil
	return in
}
// SetAddIpAddresses sets parameter AddIpAddresses to value and selects it for sending
func (in *ActionNetworkCreateInput) SetAddIpAddresses(value bool) *ActionNetworkCreateInput {
	in.AddIpAddresses = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AddIpAddresses"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkCreateInput) SelectParameters(params ...string) *ActionNetworkCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNetworkCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNetworkCreateRequest is a type for the entire action request
type ActionNetworkCreateRequest struct {
	Network map[string]interface{} `json:"network"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionNetworkCreateOutput is a type for action output parameters
type ActionNetworkCreateOutput struct {
	Id int64 `json:"id"`
	Label string `json:"label"`
	Location *ActionLocationShowOutput `json:"location"`
	IpVersion int64 `json:"ip_version"`
	Address string `json:"address"`
	Prefix int64 `json:"prefix"`
	Role string `json:"role"`
	Managed bool `json:"managed"`
	SplitAccess string `json:"split_access"`
	SplitPrefix int64 `json:"split_prefix"`
	Autopick bool `json:"autopick"`
	Purpose string `json:"purpose"`
	Size int64 `json:"size"`
	Used int64 `json:"used"`
	Assigned int64 `json:"assigned"`
	Owned int64 `json:"owned"`
}

// ActionNetworkCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionNetworkCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionNetworkCreateResponse struct {
	Action *ActionNetworkCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Network *ActionNetworkCreateOutput `json:"network"`
		// Global output metadata
		Meta *ActionNetworkCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionNetworkCreateOutput
}


// Prepare the action for invocation
func (action *ActionNetworkCreate) Prepare() *ActionNetworkCreateInvocation {
	return &ActionNetworkCreateInvocation{
		Action: action,
		Path: "/v5.0/networks",
	}
}

// ActionNetworkCreateInvocation is used to configure action for invocation
type ActionNetworkCreateInvocation struct {
	// Pointer to the action
	Action *ActionNetworkCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNetworkCreateInput
	// Global meta input parameters
	MetaInput *ActionNetworkCreateMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNetworkCreateInvocation) NewInput() *ActionNetworkCreateInput {
	inv.Input = &ActionNetworkCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNetworkCreateInvocation) SetInput(input *ActionNetworkCreateInput) *ActionNetworkCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNetworkCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNetworkCreateInvocation) NewMetaInput() *ActionNetworkCreateMetaGlobalInput {
	inv.MetaInput = &ActionNetworkCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNetworkCreateInvocation) SetMetaInput(input *ActionNetworkCreateMetaGlobalInput) *ActionNetworkCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNetworkCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNetworkCreateInvocation) Call() (*ActionNetworkCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionNetworkCreateInvocation) callAsBody() (*ActionNetworkCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNetworkCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Network
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionNetworkCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionNetworkCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionNetworkCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionNetworkCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionNetworkCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionNetworkCreateInvocation) makeAllInputParams() *ActionNetworkCreateRequest {
	return &ActionNetworkCreateRequest{
		Network: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionNetworkCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("Location") {
			ret["location"] = inv.Input.Location
		}
		if inv.IsParameterSelected("IpVersion") {
			ret["ip_version"] = inv.Input.IpVersion
		}
		if inv.IsParameterSelected("Address") {
			ret["address"] = inv.Input.Address
		}
		if inv.IsParameterSelected("Prefix") {
			ret["prefix"] = inv.Input.Prefix
		}
		if inv.IsParameterSelected("Role") {
			ret["role"] = inv.Input.Role
		}
		if inv.IsParameterSelected("Managed") {
			ret["managed"] = inv.Input.Managed
		}
		if inv.IsParameterSelected("SplitAccess") {
			ret["split_access"] = inv.Input.SplitAccess
		}
		if inv.IsParameterSelected("SplitPrefix") {
			ret["split_prefix"] = inv.Input.SplitPrefix
		}
		if inv.IsParameterSelected("Autopick") {
			ret["autopick"] = inv.Input.Autopick
		}
		if inv.IsParameterSelected("Purpose") {
			ret["purpose"] = inv.Input.Purpose
		}
		if inv.IsParameterSelected("AddIpAddresses") {
			ret["add_ip_addresses"] = inv.Input.AddIpAddresses
		}
	}

	return ret
}

func (inv *ActionNetworkCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
