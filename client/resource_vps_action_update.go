package client

import (
	"strings"
)

// ActionVpsUpdate is a type for action Vps#Update
type ActionVpsUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsUpdate(client *Client) *ActionVpsUpdate {
	return &ActionVpsUpdate{
		Client: client,
	}
}

// ActionVpsUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionVpsUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsUpdateMetaGlobalInput) SetNo(value bool) *ActionVpsUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsUpdateMetaGlobalInput) SetIncludes(value string) *ActionVpsUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionVpsUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsUpdateInput is a type for action input parameters
type ActionVpsUpdateInput struct {
	User int64 `json:"user"`
	Hostname string `json:"hostname"`
	ManageHostname bool `json:"manage_hostname"`
	OsTemplate int64 `json:"os_template"`
	Info string `json:"info"`
	DnsResolver int64 `json:"dns_resolver"`
	Node int64 `json:"node"`
	Onboot bool `json:"onboot"`
	Onstartall bool `json:"onstartall"`
	Config string `json:"config"`
	CpuLimit int64 `json:"cpu_limit"`
	Memory int64 `json:"memory"`
	Swap int64 `json:"swap"`
	Cpu int64 `json:"cpu"`
	ChangeReason string `json:"change_reason"`
	AdminOverride bool `json:"admin_override"`
	AdminLockType string `json:"admin_lock_type"`
	ObjectState string `json:"object_state"`
	ExpirationDate string `json:"expiration_date"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionVpsUpdateInput) SetUser(value int64) *ActionVpsUpdateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}
// SetHostname sets parameter Hostname to value and selects it for sending
func (in *ActionVpsUpdateInput) SetHostname(value string) *ActionVpsUpdateInput {
	in.Hostname = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Hostname"] = nil
	return in
}
// SetManageHostname sets parameter ManageHostname to value and selects it for sending
func (in *ActionVpsUpdateInput) SetManageHostname(value bool) *ActionVpsUpdateInput {
	in.ManageHostname = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ManageHostname"] = nil
	return in
}
// SetOsTemplate sets parameter OsTemplate to value and selects it for sending
func (in *ActionVpsUpdateInput) SetOsTemplate(value int64) *ActionVpsUpdateInput {
	in.OsTemplate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OsTemplate"] = nil
	return in
}
// SetInfo sets parameter Info to value and selects it for sending
func (in *ActionVpsUpdateInput) SetInfo(value string) *ActionVpsUpdateInput {
	in.Info = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Info"] = nil
	return in
}
// SetDnsResolver sets parameter DnsResolver to value and selects it for sending
func (in *ActionVpsUpdateInput) SetDnsResolver(value int64) *ActionVpsUpdateInput {
	in.DnsResolver = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DnsResolver"] = nil
	return in
}
// SetNode sets parameter Node to value and selects it for sending
func (in *ActionVpsUpdateInput) SetNode(value int64) *ActionVpsUpdateInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}
// SetOnboot sets parameter Onboot to value and selects it for sending
func (in *ActionVpsUpdateInput) SetOnboot(value bool) *ActionVpsUpdateInput {
	in.Onboot = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Onboot"] = nil
	return in
}
// SetOnstartall sets parameter Onstartall to value and selects it for sending
func (in *ActionVpsUpdateInput) SetOnstartall(value bool) *ActionVpsUpdateInput {
	in.Onstartall = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Onstartall"] = nil
	return in
}
// SetConfig sets parameter Config to value and selects it for sending
func (in *ActionVpsUpdateInput) SetConfig(value string) *ActionVpsUpdateInput {
	in.Config = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Config"] = nil
	return in
}
// SetCpuLimit sets parameter CpuLimit to value and selects it for sending
func (in *ActionVpsUpdateInput) SetCpuLimit(value int64) *ActionVpsUpdateInput {
	in.CpuLimit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CpuLimit"] = nil
	return in
}
// SetMemory sets parameter Memory to value and selects it for sending
func (in *ActionVpsUpdateInput) SetMemory(value int64) *ActionVpsUpdateInput {
	in.Memory = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Memory"] = nil
	return in
}
// SetSwap sets parameter Swap to value and selects it for sending
func (in *ActionVpsUpdateInput) SetSwap(value int64) *ActionVpsUpdateInput {
	in.Swap = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Swap"] = nil
	return in
}
// SetCpu sets parameter Cpu to value and selects it for sending
func (in *ActionVpsUpdateInput) SetCpu(value int64) *ActionVpsUpdateInput {
	in.Cpu = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Cpu"] = nil
	return in
}
// SetChangeReason sets parameter ChangeReason to value and selects it for sending
func (in *ActionVpsUpdateInput) SetChangeReason(value string) *ActionVpsUpdateInput {
	in.ChangeReason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ChangeReason"] = nil
	return in
}
// SetAdminOverride sets parameter AdminOverride to value and selects it for sending
func (in *ActionVpsUpdateInput) SetAdminOverride(value bool) *ActionVpsUpdateInput {
	in.AdminOverride = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AdminOverride"] = nil
	return in
}
// SetAdminLockType sets parameter AdminLockType to value and selects it for sending
func (in *ActionVpsUpdateInput) SetAdminLockType(value string) *ActionVpsUpdateInput {
	in.AdminLockType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AdminLockType"] = nil
	return in
}
// SetObjectState sets parameter ObjectState to value and selects it for sending
func (in *ActionVpsUpdateInput) SetObjectState(value string) *ActionVpsUpdateInput {
	in.ObjectState = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ObjectState"] = nil
	return in
}
// SetExpirationDate sets parameter ExpirationDate to value and selects it for sending
func (in *ActionVpsUpdateInput) SetExpirationDate(value string) *ActionVpsUpdateInput {
	in.ExpirationDate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ExpirationDate"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsUpdateInput) SelectParameters(params ...string) *ActionVpsUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsUpdateRequest is a type for the entire action request
type ActionVpsUpdateRequest struct {
	Vps map[string]interface{} `json:"vps"`
	Meta map[string]interface{} `json:"_meta"`
}


// ActionVpsUpdateMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsUpdateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsUpdateResponse struct {
	Action *ActionVpsUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionVpsUpdateMetaGlobalOutput `json:"_meta"`
	}
}


// Prepare the action for invocation
func (action *ActionVpsUpdate) Prepare() *ActionVpsUpdateInvocation {
	return &ActionVpsUpdateInvocation{
		Action: action,
		Path: "/v5.0/vpses/:vps_id",
	}
}

// ActionVpsUpdateInvocation is used to configure action for invocation
type ActionVpsUpdateInvocation struct {
	// Pointer to the action
	Action *ActionVpsUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsUpdateInput
	// Global meta input parameters
	MetaInput *ActionVpsUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsUpdateInvocation) SetPathParamInt(param string, value int64) *ActionVpsUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsUpdateInvocation) SetPathParamString(param string, value string) *ActionVpsUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsUpdateInvocation) NewInput() *ActionVpsUpdateInput {
	inv.Input = &ActionVpsUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsUpdateInvocation) SetInput(input *ActionVpsUpdateInput) *ActionVpsUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsUpdateInvocation) NewMetaInput() *ActionVpsUpdateMetaGlobalInput {
	inv.MetaInput = &ActionVpsUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsUpdateInvocation) SetMetaInput(input *ActionVpsUpdateMetaGlobalInput) *ActionVpsUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsUpdateInvocation) Call() (*ActionVpsUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionVpsUpdateInvocation) callAsBody() (*ActionVpsUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsUpdateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsUpdateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsUpdateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsUpdateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsUpdateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionVpsUpdateInvocation) makeAllInputParams() *ActionVpsUpdateRequest {
	return &ActionVpsUpdateRequest{
		Vps: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("User") {
			ret["user"] = inv.Input.User
		}
		if inv.IsParameterSelected("Hostname") {
			ret["hostname"] = inv.Input.Hostname
		}
		if inv.IsParameterSelected("ManageHostname") {
			ret["manage_hostname"] = inv.Input.ManageHostname
		}
		if inv.IsParameterSelected("OsTemplate") {
			ret["os_template"] = inv.Input.OsTemplate
		}
		if inv.IsParameterSelected("Info") {
			ret["info"] = inv.Input.Info
		}
		if inv.IsParameterSelected("DnsResolver") {
			ret["dns_resolver"] = inv.Input.DnsResolver
		}
		if inv.IsParameterSelected("Node") {
			ret["node"] = inv.Input.Node
		}
		if inv.IsParameterSelected("Onboot") {
			ret["onboot"] = inv.Input.Onboot
		}
		if inv.IsParameterSelected("Onstartall") {
			ret["onstartall"] = inv.Input.Onstartall
		}
		if inv.IsParameterSelected("Config") {
			ret["config"] = inv.Input.Config
		}
		if inv.IsParameterSelected("CpuLimit") {
			ret["cpu_limit"] = inv.Input.CpuLimit
		}
		if inv.IsParameterSelected("Memory") {
			ret["memory"] = inv.Input.Memory
		}
		if inv.IsParameterSelected("Swap") {
			ret["swap"] = inv.Input.Swap
		}
		if inv.IsParameterSelected("Cpu") {
			ret["cpu"] = inv.Input.Cpu
		}
		if inv.IsParameterSelected("ChangeReason") {
			ret["change_reason"] = inv.Input.ChangeReason
		}
		if inv.IsParameterSelected("AdminOverride") {
			ret["admin_override"] = inv.Input.AdminOverride
		}
		if inv.IsParameterSelected("AdminLockType") {
			ret["admin_lock_type"] = inv.Input.AdminLockType
		}
		if inv.IsParameterSelected("ObjectState") {
			ret["object_state"] = inv.Input.ObjectState
		}
		if inv.IsParameterSelected("ExpirationDate") {
			ret["expiration_date"] = inv.Input.ExpirationDate
		}
	}

	return ret
}

func (inv *ActionVpsUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
