package client

import (
)

// ActionVpsCreate is a type for action Vps#Create
type ActionVpsCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsCreate(client *Client) *ActionVpsCreate {
	return &ActionVpsCreate{
		Client: client,
	}
}

// ActionVpsCreateMetaGlobalInput is a type for action global meta input parameters
type ActionVpsCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsCreateMetaGlobalInput) SetNo(value bool) *ActionVpsCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsCreateMetaGlobalInput) SetIncludes(value string) *ActionVpsCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsCreateMetaGlobalInput) SelectParameters(params ...string) *ActionVpsCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsCreateInput is a type for action input parameters
type ActionVpsCreateInput struct {
	Environment int64 `json:"environment"`
	Location int64 `json:"location"`
	User int64 `json:"user"`
	Hostname string `json:"hostname"`
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
	Diskspace int64 `json:"diskspace"`
	Ipv4 int64 `json:"ipv4"`
	Ipv6 int64 `json:"ipv6"`
	Ipv4Private int64 `json:"ipv4_private"`
	UserNamespaceMap int64 `json:"user_namespace_map"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionVpsCreateInput) SetEnvironment(value int64) *ActionVpsCreateInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Environment"] = nil
	return in
}
// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionVpsCreateInput) SetLocation(value int64) *ActionVpsCreateInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}
// SetUser sets parameter User to value and selects it for sending
func (in *ActionVpsCreateInput) SetUser(value int64) *ActionVpsCreateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}
// SetHostname sets parameter Hostname to value and selects it for sending
func (in *ActionVpsCreateInput) SetHostname(value string) *ActionVpsCreateInput {
	in.Hostname = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Hostname"] = nil
	return in
}
// SetOsTemplate sets parameter OsTemplate to value and selects it for sending
func (in *ActionVpsCreateInput) SetOsTemplate(value int64) *ActionVpsCreateInput {
	in.OsTemplate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OsTemplate"] = nil
	return in
}
// SetInfo sets parameter Info to value and selects it for sending
func (in *ActionVpsCreateInput) SetInfo(value string) *ActionVpsCreateInput {
	in.Info = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Info"] = nil
	return in
}
// SetDnsResolver sets parameter DnsResolver to value and selects it for sending
func (in *ActionVpsCreateInput) SetDnsResolver(value int64) *ActionVpsCreateInput {
	in.DnsResolver = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DnsResolver"] = nil
	return in
}
// SetNode sets parameter Node to value and selects it for sending
func (in *ActionVpsCreateInput) SetNode(value int64) *ActionVpsCreateInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}
// SetOnboot sets parameter Onboot to value and selects it for sending
func (in *ActionVpsCreateInput) SetOnboot(value bool) *ActionVpsCreateInput {
	in.Onboot = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Onboot"] = nil
	return in
}
// SetOnstartall sets parameter Onstartall to value and selects it for sending
func (in *ActionVpsCreateInput) SetOnstartall(value bool) *ActionVpsCreateInput {
	in.Onstartall = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Onstartall"] = nil
	return in
}
// SetConfig sets parameter Config to value and selects it for sending
func (in *ActionVpsCreateInput) SetConfig(value string) *ActionVpsCreateInput {
	in.Config = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Config"] = nil
	return in
}
// SetCpuLimit sets parameter CpuLimit to value and selects it for sending
func (in *ActionVpsCreateInput) SetCpuLimit(value int64) *ActionVpsCreateInput {
	in.CpuLimit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CpuLimit"] = nil
	return in
}
// SetMemory sets parameter Memory to value and selects it for sending
func (in *ActionVpsCreateInput) SetMemory(value int64) *ActionVpsCreateInput {
	in.Memory = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Memory"] = nil
	return in
}
// SetSwap sets parameter Swap to value and selects it for sending
func (in *ActionVpsCreateInput) SetSwap(value int64) *ActionVpsCreateInput {
	in.Swap = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Swap"] = nil
	return in
}
// SetCpu sets parameter Cpu to value and selects it for sending
func (in *ActionVpsCreateInput) SetCpu(value int64) *ActionVpsCreateInput {
	in.Cpu = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Cpu"] = nil
	return in
}
// SetDiskspace sets parameter Diskspace to value and selects it for sending
func (in *ActionVpsCreateInput) SetDiskspace(value int64) *ActionVpsCreateInput {
	in.Diskspace = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Diskspace"] = nil
	return in
}
// SetIpv4 sets parameter Ipv4 to value and selects it for sending
func (in *ActionVpsCreateInput) SetIpv4(value int64) *ActionVpsCreateInput {
	in.Ipv4 = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Ipv4"] = nil
	return in
}
// SetIpv6 sets parameter Ipv6 to value and selects it for sending
func (in *ActionVpsCreateInput) SetIpv6(value int64) *ActionVpsCreateInput {
	in.Ipv6 = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Ipv6"] = nil
	return in
}
// SetIpv4Private sets parameter Ipv4Private to value and selects it for sending
func (in *ActionVpsCreateInput) SetIpv4Private(value int64) *ActionVpsCreateInput {
	in.Ipv4Private = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Ipv4Private"] = nil
	return in
}
// SetUserNamespaceMap sets parameter UserNamespaceMap to value and selects it for sending
func (in *ActionVpsCreateInput) SetUserNamespaceMap(value int64) *ActionVpsCreateInput {
	in.UserNamespaceMap = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserNamespaceMap"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsCreateInput) SelectParameters(params ...string) *ActionVpsCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsCreateRequest is a type for the entire action request
type ActionVpsCreateRequest struct {
	Vps map[string]interface{} `json:"vps"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionVpsCreateOutput is a type for action output parameters
type ActionVpsCreateOutput struct {
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
	ObjectState string `json:"object_state"`
	ExpirationDate string `json:"expiration_date"`
}

// ActionVpsCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsCreateResponse struct {
	Action *ActionVpsCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Vps *ActionVpsCreateOutput `json:"vps"`
		// Global output metadata
		Meta *ActionVpsCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionVpsCreateOutput
}


// Prepare the action for invocation
func (action *ActionVpsCreate) Prepare() *ActionVpsCreateInvocation {
	return &ActionVpsCreateInvocation{
		Action: action,
		Path: "/v6.0/vpses",
	}
}

// ActionVpsCreateInvocation is used to configure action for invocation
type ActionVpsCreateInvocation struct {
	// Pointer to the action
	Action *ActionVpsCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsCreateInput
	// Global meta input parameters
	MetaInput *ActionVpsCreateMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsCreateInvocation) NewInput() *ActionVpsCreateInput {
	inv.Input = &ActionVpsCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsCreateInvocation) SetInput(input *ActionVpsCreateInput) *ActionVpsCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsCreateInvocation) NewMetaInput() *ActionVpsCreateMetaGlobalInput {
	inv.MetaInput = &ActionVpsCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsCreateInvocation) SetMetaInput(input *ActionVpsCreateMetaGlobalInput) *ActionVpsCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsCreateInvocation) Call() (*ActionVpsCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionVpsCreateInvocation) callAsBody() (*ActionVpsCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Vps
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionVpsCreateInvocation) makeAllInputParams() *ActionVpsCreateRequest {
	return &ActionVpsCreateRequest{
		Vps: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			ret["environment"] = inv.Input.Environment
		}
		if inv.IsParameterSelected("Location") {
			ret["location"] = inv.Input.Location
		}
		if inv.IsParameterSelected("User") {
			ret["user"] = inv.Input.User
		}
		if inv.IsParameterSelected("Hostname") {
			ret["hostname"] = inv.Input.Hostname
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
		if inv.IsParameterSelected("Diskspace") {
			ret["diskspace"] = inv.Input.Diskspace
		}
		if inv.IsParameterSelected("Ipv4") {
			ret["ipv4"] = inv.Input.Ipv4
		}
		if inv.IsParameterSelected("Ipv6") {
			ret["ipv6"] = inv.Input.Ipv6
		}
		if inv.IsParameterSelected("Ipv4Private") {
			ret["ipv4_private"] = inv.Input.Ipv4Private
		}
		if inv.IsParameterSelected("UserNamespaceMap") {
			ret["user_namespace_map"] = inv.Input.UserNamespaceMap
		}
	}

	return ret
}

func (inv *ActionVpsCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
