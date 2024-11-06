package client

import ()

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
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
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

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsCreateMetaGlobalInput) SetNo(value bool) *ActionVpsCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
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
	AddressLocation         int64  `json:"address_location"`
	AllowAdminModifications bool   `json:"allow_admin_modifications"`
	AutostartEnable         bool   `json:"autostart_enable"`
	AutostartPriority       int64  `json:"autostart_priority"`
	CgroupVersion           string `json:"cgroup_version"`
	Config                  string `json:"config"`
	Cpu                     int64  `json:"cpu"`
	CpuLimit                int64  `json:"cpu_limit"`
	Diskspace               int64  `json:"diskspace"`
	DnsResolver             int64  `json:"dns_resolver"`
	Environment             int64  `json:"environment"`
	Hostname                string `json:"hostname"`
	Info                    string `json:"info"`
	Ipv4                    int64  `json:"ipv4"`
	Ipv4Private             int64  `json:"ipv4_private"`
	Ipv6                    int64  `json:"ipv6"`
	Location                int64  `json:"location"`
	Memory                  int64  `json:"memory"`
	Node                    int64  `json:"node"`
	Onstartall              bool   `json:"onstartall"`
	OsTemplate              int64  `json:"os_template"`
	Start                   bool   `json:"start"`
	StartMenuTimeout        int64  `json:"start_menu_timeout"`
	Swap                    int64  `json:"swap"`
	User                    int64  `json:"user"`
	UserNamespaceMap        int64  `json:"user_namespace_map"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAddressLocation sets parameter AddressLocation to value and selects it for sending
func (in *ActionVpsCreateInput) SetAddressLocation(value int64) *ActionVpsCreateInput {
	in.AddressLocation = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetAddressLocationNil(false)
	in._selectedParameters["AddressLocation"] = nil
	return in
}

// SetAddressLocationNil sets parameter AddressLocation to nil and selects it for sending
func (in *ActionVpsCreateInput) SetAddressLocationNil(set bool) *ActionVpsCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["AddressLocation"] = nil
		in.SelectParameters("AddressLocation")
	} else {
		delete(in._nilParameters, "AddressLocation")
	}
	return in
}

// SetAllowAdminModifications sets parameter AllowAdminModifications to value and selects it for sending
func (in *ActionVpsCreateInput) SetAllowAdminModifications(value bool) *ActionVpsCreateInput {
	in.AllowAdminModifications = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AllowAdminModifications"] = nil
	return in
}

// SetAutostartEnable sets parameter AutostartEnable to value and selects it for sending
func (in *ActionVpsCreateInput) SetAutostartEnable(value bool) *ActionVpsCreateInput {
	in.AutostartEnable = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AutostartEnable"] = nil
	return in
}

// SetAutostartPriority sets parameter AutostartPriority to value and selects it for sending
func (in *ActionVpsCreateInput) SetAutostartPriority(value int64) *ActionVpsCreateInput {
	in.AutostartPriority = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AutostartPriority"] = nil
	return in
}

// SetCgroupVersion sets parameter CgroupVersion to value and selects it for sending
func (in *ActionVpsCreateInput) SetCgroupVersion(value string) *ActionVpsCreateInput {
	in.CgroupVersion = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CgroupVersion"] = nil
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

// SetCpu sets parameter Cpu to value and selects it for sending
func (in *ActionVpsCreateInput) SetCpu(value int64) *ActionVpsCreateInput {
	in.Cpu = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Cpu"] = nil
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

// SetDiskspace sets parameter Diskspace to value and selects it for sending
func (in *ActionVpsCreateInput) SetDiskspace(value int64) *ActionVpsCreateInput {
	in.Diskspace = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Diskspace"] = nil
	return in
}

// SetDnsResolver sets parameter DnsResolver to value and selects it for sending
func (in *ActionVpsCreateInput) SetDnsResolver(value int64) *ActionVpsCreateInput {
	in.DnsResolver = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetDnsResolverNil(false)
	in._selectedParameters["DnsResolver"] = nil
	return in
}

// SetDnsResolverNil sets parameter DnsResolver to nil and selects it for sending
func (in *ActionVpsCreateInput) SetDnsResolverNil(set bool) *ActionVpsCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["DnsResolver"] = nil
		in.SelectParameters("DnsResolver")
	} else {
		delete(in._nilParameters, "DnsResolver")
	}
	return in
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionVpsCreateInput) SetEnvironment(value int64) *ActionVpsCreateInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEnvironmentNil(false)
	in._selectedParameters["Environment"] = nil
	return in
}

// SetEnvironmentNil sets parameter Environment to nil and selects it for sending
func (in *ActionVpsCreateInput) SetEnvironmentNil(set bool) *ActionVpsCreateInput {
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

// SetHostname sets parameter Hostname to value and selects it for sending
func (in *ActionVpsCreateInput) SetHostname(value string) *ActionVpsCreateInput {
	in.Hostname = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Hostname"] = nil
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

// SetIpv4 sets parameter Ipv4 to value and selects it for sending
func (in *ActionVpsCreateInput) SetIpv4(value int64) *ActionVpsCreateInput {
	in.Ipv4 = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Ipv4"] = nil
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

// SetIpv6 sets parameter Ipv6 to value and selects it for sending
func (in *ActionVpsCreateInput) SetIpv6(value int64) *ActionVpsCreateInput {
	in.Ipv6 = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Ipv6"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionVpsCreateInput) SetLocation(value int64) *ActionVpsCreateInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLocationNil(false)
	in._selectedParameters["Location"] = nil
	return in
}

// SetLocationNil sets parameter Location to nil and selects it for sending
func (in *ActionVpsCreateInput) SetLocationNil(set bool) *ActionVpsCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Location"] = nil
		in.SelectParameters("Location")
	} else {
		delete(in._nilParameters, "Location")
	}
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

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionVpsCreateInput) SetNode(value int64) *ActionVpsCreateInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNodeNil(false)
	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeNil sets parameter Node to nil and selects it for sending
func (in *ActionVpsCreateInput) SetNodeNil(set bool) *ActionVpsCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Node"] = nil
		in.SelectParameters("Node")
	} else {
		delete(in._nilParameters, "Node")
	}
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

// SetOsTemplate sets parameter OsTemplate to value and selects it for sending
func (in *ActionVpsCreateInput) SetOsTemplate(value int64) *ActionVpsCreateInput {
	in.OsTemplate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetOsTemplateNil(false)
	in._selectedParameters["OsTemplate"] = nil
	return in
}

// SetOsTemplateNil sets parameter OsTemplate to nil and selects it for sending
func (in *ActionVpsCreateInput) SetOsTemplateNil(set bool) *ActionVpsCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["OsTemplate"] = nil
		in.SelectParameters("OsTemplate")
	} else {
		delete(in._nilParameters, "OsTemplate")
	}
	return in
}

// SetStart sets parameter Start to value and selects it for sending
func (in *ActionVpsCreateInput) SetStart(value bool) *ActionVpsCreateInput {
	in.Start = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Start"] = nil
	return in
}

// SetStartMenuTimeout sets parameter StartMenuTimeout to value and selects it for sending
func (in *ActionVpsCreateInput) SetStartMenuTimeout(value int64) *ActionVpsCreateInput {
	in.StartMenuTimeout = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["StartMenuTimeout"] = nil
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

// SetUser sets parameter User to value and selects it for sending
func (in *ActionVpsCreateInput) SetUser(value int64) *ActionVpsCreateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionVpsCreateInput) SetUserNil(set bool) *ActionVpsCreateInput {
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

// SetUserNamespaceMap sets parameter UserNamespaceMap to value and selects it for sending
func (in *ActionVpsCreateInput) SetUserNamespaceMap(value int64) *ActionVpsCreateInput {
	in.UserNamespaceMap = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNamespaceMapNil(false)
	in._selectedParameters["UserNamespaceMap"] = nil
	return in
}

// SetUserNamespaceMapNil sets parameter UserNamespaceMap to nil and selects it for sending
func (in *ActionVpsCreateInput) SetUserNamespaceMapNil(set bool) *ActionVpsCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["UserNamespaceMap"] = nil
		in.SelectParameters("UserNamespaceMap")
	} else {
		delete(in._nilParameters, "UserNamespaceMap")
	}
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

// UnselectParameters unsets parameters from ActionVpsCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsCreateInput) UnselectParameters(params ...string) *ActionVpsCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
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
	Vps  map[string]interface{} `json:"vps"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionVpsCreateOutput is a type for action output parameters
type ActionVpsCreateOutput struct {
	AllowAdminModifications bool                              `json:"allow_admin_modifications"`
	AutostartEnable         bool                              `json:"autostart_enable"`
	AutostartPriority       int64                             `json:"autostart_priority"`
	CgroupVersion           string                            `json:"cgroup_version"`
	Config                  string                            `json:"config"`
	Cpu                     int64                             `json:"cpu"`
	CpuIdle                 float64                           `json:"cpu_idle"`
	CpuIowait               float64                           `json:"cpu_iowait"`
	CpuIrq                  float64                           `json:"cpu_irq"`
	CpuLimit                int64                             `json:"cpu_limit"`
	CpuNice                 float64                           `json:"cpu_nice"`
	CpuSoftirq              float64                           `json:"cpu_softirq"`
	CpuSystem               float64                           `json:"cpu_system"`
	CpuUser                 float64                           `json:"cpu_user"`
	CreatedAt               string                            `json:"created_at"`
	Dataset                 *ActionDatasetShowOutput          `json:"dataset"`
	Diskspace               int64                             `json:"diskspace"`
	DnsResolver             *ActionDnsResolverShowOutput      `json:"dns_resolver"`
	ExpirationDate          string                            `json:"expiration_date"`
	Hostname                string                            `json:"hostname"`
	Id                      int64                             `json:"id"`
	InRescueMode            bool                              `json:"in_rescue_mode"`
	Info                    string                            `json:"info"`
	IsRunning               bool                              `json:"is_running"`
	Loadavg1                float64                           `json:"loadavg1"`
	Loadavg15               float64                           `json:"loadavg15"`
	Loadavg5                float64                           `json:"loadavg5"`
	ManageHostname          bool                              `json:"manage_hostname"`
	Memory                  int64                             `json:"memory"`
	Node                    *ActionNodeShowOutput             `json:"node"`
	ObjectState             string                            `json:"object_state"`
	Onstartall              bool                              `json:"onstartall"`
	OsTemplate              *ActionOsTemplateShowOutput       `json:"os_template"`
	Pool                    *ActionPoolShowOutput             `json:"pool"`
	ProcessCount            int64                             `json:"process_count"`
	RemindAfterDate         string                            `json:"remind_after_date"`
	StartMenuTimeout        int64                             `json:"start_menu_timeout"`
	Swap                    int64                             `json:"swap"`
	Uptime                  int64                             `json:"uptime"`
	UsedDiskspace           int64                             `json:"used_diskspace"`
	UsedMemory              int64                             `json:"used_memory"`
	UsedSwap                int64                             `json:"used_swap"`
	User                    *ActionUserShowOutput             `json:"user"`
	UserNamespaceMap        *ActionUserNamespaceMapShowOutput `json:"user_namespace_map"`
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
		Path:   "/v7.0/vpses",
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

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
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

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
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
func (resp *ActionVpsCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionVpsCreateInvocation) makeAllInputParams() *ActionVpsCreateRequest {
	return &ActionVpsCreateRequest{
		Vps:  inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("AddressLocation") {
			if inv.IsParameterNil("AddressLocation") {
				ret["address_location"] = nil
			} else {
				ret["address_location"] = inv.Input.AddressLocation
			}
		}
		if inv.IsParameterSelected("AllowAdminModifications") {
			ret["allow_admin_modifications"] = inv.Input.AllowAdminModifications
		}
		if inv.IsParameterSelected("AutostartEnable") {
			ret["autostart_enable"] = inv.Input.AutostartEnable
		}
		if inv.IsParameterSelected("AutostartPriority") {
			ret["autostart_priority"] = inv.Input.AutostartPriority
		}
		if inv.IsParameterSelected("CgroupVersion") {
			ret["cgroup_version"] = inv.Input.CgroupVersion
		}
		if inv.IsParameterSelected("Config") {
			ret["config"] = inv.Input.Config
		}
		if inv.IsParameterSelected("Cpu") {
			ret["cpu"] = inv.Input.Cpu
		}
		if inv.IsParameterSelected("CpuLimit") {
			ret["cpu_limit"] = inv.Input.CpuLimit
		}
		if inv.IsParameterSelected("Diskspace") {
			ret["diskspace"] = inv.Input.Diskspace
		}
		if inv.IsParameterSelected("DnsResolver") {
			if inv.IsParameterNil("DnsResolver") {
				ret["dns_resolver"] = nil
			} else {
				ret["dns_resolver"] = inv.Input.DnsResolver
			}
		}
		if inv.IsParameterSelected("Environment") {
			if inv.IsParameterNil("Environment") {
				ret["environment"] = nil
			} else {
				ret["environment"] = inv.Input.Environment
			}
		}
		if inv.IsParameterSelected("Hostname") {
			ret["hostname"] = inv.Input.Hostname
		}
		if inv.IsParameterSelected("Info") {
			ret["info"] = inv.Input.Info
		}
		if inv.IsParameterSelected("Ipv4") {
			ret["ipv4"] = inv.Input.Ipv4
		}
		if inv.IsParameterSelected("Ipv4Private") {
			ret["ipv4_private"] = inv.Input.Ipv4Private
		}
		if inv.IsParameterSelected("Ipv6") {
			ret["ipv6"] = inv.Input.Ipv6
		}
		if inv.IsParameterSelected("Location") {
			if inv.IsParameterNil("Location") {
				ret["location"] = nil
			} else {
				ret["location"] = inv.Input.Location
			}
		}
		if inv.IsParameterSelected("Memory") {
			ret["memory"] = inv.Input.Memory
		}
		if inv.IsParameterSelected("Node") {
			if inv.IsParameterNil("Node") {
				ret["node"] = nil
			} else {
				ret["node"] = inv.Input.Node
			}
		}
		if inv.IsParameterSelected("Onstartall") {
			ret["onstartall"] = inv.Input.Onstartall
		}
		if inv.IsParameterSelected("OsTemplate") {
			if inv.IsParameterNil("OsTemplate") {
				ret["os_template"] = nil
			} else {
				ret["os_template"] = inv.Input.OsTemplate
			}
		}
		if inv.IsParameterSelected("Start") {
			ret["start"] = inv.Input.Start
		}
		if inv.IsParameterSelected("StartMenuTimeout") {
			ret["start_menu_timeout"] = inv.Input.StartMenuTimeout
		}
		if inv.IsParameterSelected("Swap") {
			ret["swap"] = inv.Input.Swap
		}
		if inv.IsParameterSelected("User") {
			if inv.IsParameterNil("User") {
				ret["user"] = nil
			} else {
				ret["user"] = inv.Input.User
			}
		}
		if inv.IsParameterSelected("UserNamespaceMap") {
			if inv.IsParameterNil("UserNamespaceMap") {
				ret["user_namespace_map"] = nil
			} else {
				ret["user_namespace_map"] = inv.Input.UserNamespaceMap
			}
		}
	}

	return ret
}

func (inv *ActionVpsCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
