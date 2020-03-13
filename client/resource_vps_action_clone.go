package client

import (
	"strings"
)

// ActionVpsClone is a type for action Vps#Clone
type ActionVpsClone struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsClone(client *Client) *ActionVpsClone {
	return &ActionVpsClone{
		Client: client,
	}
}

// ActionVpsCloneMetaGlobalInput is a type for action global meta input parameters
type ActionVpsCloneMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsCloneMetaGlobalInput) SetNo(value bool) *ActionVpsCloneMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsCloneMetaGlobalInput) SetIncludes(value string) *ActionVpsCloneMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsCloneMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsCloneMetaGlobalInput) SelectParameters(params ...string) *ActionVpsCloneMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsCloneMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsCloneInput is a type for action input parameters
type ActionVpsCloneInput struct {
	Environment int64 `json:"environment"`
	Location int64 `json:"location"`
	Node int64 `json:"node"`
	User int64 `json:"user"`
	KeepPlatform bool `json:"keep_platform"`
	Subdatasets bool `json:"subdatasets"`
	DatasetPlans bool `json:"dataset_plans"`
	Configs bool `json:"configs"`
	Resources bool `json:"resources"`
	Features bool `json:"features"`
	Hostname string `json:"hostname"`
	Stop bool `json:"stop"`
	KeepSnapshots bool `json:"keep_snapshots"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionVpsCloneInput) SetEnvironment(value int64) *ActionVpsCloneInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Environment"] = nil
	return in
}
// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionVpsCloneInput) SetLocation(value int64) *ActionVpsCloneInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}
// SetNode sets parameter Node to value and selects it for sending
func (in *ActionVpsCloneInput) SetNode(value int64) *ActionVpsCloneInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}
// SetUser sets parameter User to value and selects it for sending
func (in *ActionVpsCloneInput) SetUser(value int64) *ActionVpsCloneInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}
// SetKeepPlatform sets parameter KeepPlatform to value and selects it for sending
func (in *ActionVpsCloneInput) SetKeepPlatform(value bool) *ActionVpsCloneInput {
	in.KeepPlatform = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["KeepPlatform"] = nil
	return in
}
// SetSubdatasets sets parameter Subdatasets to value and selects it for sending
func (in *ActionVpsCloneInput) SetSubdatasets(value bool) *ActionVpsCloneInput {
	in.Subdatasets = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Subdatasets"] = nil
	return in
}
// SetDatasetPlans sets parameter DatasetPlans to value and selects it for sending
func (in *ActionVpsCloneInput) SetDatasetPlans(value bool) *ActionVpsCloneInput {
	in.DatasetPlans = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DatasetPlans"] = nil
	return in
}
// SetConfigs sets parameter Configs to value and selects it for sending
func (in *ActionVpsCloneInput) SetConfigs(value bool) *ActionVpsCloneInput {
	in.Configs = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Configs"] = nil
	return in
}
// SetResources sets parameter Resources to value and selects it for sending
func (in *ActionVpsCloneInput) SetResources(value bool) *ActionVpsCloneInput {
	in.Resources = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Resources"] = nil
	return in
}
// SetFeatures sets parameter Features to value and selects it for sending
func (in *ActionVpsCloneInput) SetFeatures(value bool) *ActionVpsCloneInput {
	in.Features = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Features"] = nil
	return in
}
// SetHostname sets parameter Hostname to value and selects it for sending
func (in *ActionVpsCloneInput) SetHostname(value string) *ActionVpsCloneInput {
	in.Hostname = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Hostname"] = nil
	return in
}
// SetStop sets parameter Stop to value and selects it for sending
func (in *ActionVpsCloneInput) SetStop(value bool) *ActionVpsCloneInput {
	in.Stop = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Stop"] = nil
	return in
}
// SetKeepSnapshots sets parameter KeepSnapshots to value and selects it for sending
func (in *ActionVpsCloneInput) SetKeepSnapshots(value bool) *ActionVpsCloneInput {
	in.KeepSnapshots = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["KeepSnapshots"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsCloneInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsCloneInput) SelectParameters(params ...string) *ActionVpsCloneInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsCloneInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsCloneRequest is a type for the entire action request
type ActionVpsCloneRequest struct {
	Vps map[string]interface{} `json:"vps"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionVpsCloneOutput is a type for action output parameters
type ActionVpsCloneOutput struct {
	Id int64 `json:"id"`
	User *ActionUserShowOutput `json:"user"`
	Hostname string `json:"hostname"`
	ManageHostname bool `json:"manage_hostname"`
	OsTemplate *ActionOsTemplateShowOutput `json:"os_template"`
	Info string `json:"info"`
	DnsResolver *ActionDnsResolverShowOutput `json:"dns_resolver"`
	Node *ActionNodeShowOutput `json:"node"`
	Onboot bool `json:"onboot"`
	Onstartall bool `json:"onstartall"`
	Config string `json:"config"`
	CpuLimit int64 `json:"cpu_limit"`
	Dataset *ActionDatasetShowOutput `json:"dataset"`
	CreatedAt string `json:"created_at"`
	Memory int64 `json:"memory"`
	Swap int64 `json:"swap"`
	Cpu int64 `json:"cpu"`
	IsRunning bool `json:"is_running"`
	Uptime int64 `json:"uptime"`
	Loadavg float64 `json:"loadavg"`
	ProcessCount int64 `json:"process_count"`
	CpuUser float64 `json:"cpu_user"`
	CpuNice float64 `json:"cpu_nice"`
	CpuSystem float64 `json:"cpu_system"`
	CpuIdle float64 `json:"cpu_idle"`
	CpuIowait float64 `json:"cpu_iowait"`
	CpuIrq float64 `json:"cpu_irq"`
	CpuSoftirq float64 `json:"cpu_softirq"`
	UsedMemory int64 `json:"used_memory"`
	UsedSwap int64 `json:"used_swap"`
	UsedDiskspace int64 `json:"used_diskspace"`
}

// ActionVpsCloneMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsCloneMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsCloneResponse struct {
	Action *ActionVpsClone `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Vps *ActionVpsCloneOutput `json:"vps"`
		// Global output metadata
		Meta *ActionVpsCloneMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionVpsCloneOutput
}


// Prepare the action for invocation
func (action *ActionVpsClone) Prepare() *ActionVpsCloneInvocation {
	return &ActionVpsCloneInvocation{
		Action: action,
		Path: "/v6.0/vpses/{vps_id}/clone",
	}
}

// ActionVpsCloneInvocation is used to configure action for invocation
type ActionVpsCloneInvocation struct {
	// Pointer to the action
	Action *ActionVpsClone

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsCloneInput
	// Global meta input parameters
	MetaInput *ActionVpsCloneMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsCloneInvocation) SetPathParamInt(param string, value int64) *ActionVpsCloneInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsCloneInvocation) SetPathParamString(param string, value string) *ActionVpsCloneInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsCloneInvocation) NewInput() *ActionVpsCloneInput {
	inv.Input = &ActionVpsCloneInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsCloneInvocation) SetInput(input *ActionVpsCloneInput) *ActionVpsCloneInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsCloneInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsCloneInvocation) NewMetaInput() *ActionVpsCloneMetaGlobalInput {
	inv.MetaInput = &ActionVpsCloneMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsCloneInvocation) SetMetaInput(input *ActionVpsCloneMetaGlobalInput) *ActionVpsCloneInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsCloneInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsCloneInvocation) Call() (*ActionVpsCloneResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionVpsCloneInvocation) callAsBody() (*ActionVpsCloneResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsCloneResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Vps
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsCloneResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsCloneResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsCloneResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsCloneResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsCloneResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionVpsCloneInvocation) makeAllInputParams() *ActionVpsCloneRequest {
	return &ActionVpsCloneRequest{
		Vps: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsCloneInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			ret["environment"] = inv.Input.Environment
		}
		if inv.IsParameterSelected("Location") {
			ret["location"] = inv.Input.Location
		}
		if inv.IsParameterSelected("Node") {
			ret["node"] = inv.Input.Node
		}
		if inv.IsParameterSelected("User") {
			ret["user"] = inv.Input.User
		}
		if inv.IsParameterSelected("KeepPlatform") {
			ret["keep_platform"] = inv.Input.KeepPlatform
		}
		if inv.IsParameterSelected("Subdatasets") {
			ret["subdatasets"] = inv.Input.Subdatasets
		}
		if inv.IsParameterSelected("DatasetPlans") {
			ret["dataset_plans"] = inv.Input.DatasetPlans
		}
		if inv.IsParameterSelected("Configs") {
			ret["configs"] = inv.Input.Configs
		}
		if inv.IsParameterSelected("Resources") {
			ret["resources"] = inv.Input.Resources
		}
		if inv.IsParameterSelected("Features") {
			ret["features"] = inv.Input.Features
		}
		if inv.IsParameterSelected("Hostname") {
			ret["hostname"] = inv.Input.Hostname
		}
		if inv.IsParameterSelected("Stop") {
			ret["stop"] = inv.Input.Stop
		}
		if inv.IsParameterSelected("KeepSnapshots") {
			ret["keep_snapshots"] = inv.Input.KeepSnapshots
		}
	}

	return ret
}

func (inv *ActionVpsCloneInvocation) makeMetaInputParams() map[string]interface{} {
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
