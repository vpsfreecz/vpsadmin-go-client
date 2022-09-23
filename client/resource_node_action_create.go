package client

import ()

// ActionNodeCreate is a type for action Node#Create
type ActionNodeCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeCreate(client *Client) *ActionNodeCreate {
	return &ActionNodeCreate{
		Client: client,
	}
}

// ActionNodeCreateMetaGlobalInput is a type for action global meta input parameters
type ActionNodeCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeCreateMetaGlobalInput) SetIncludes(value string) *ActionNodeCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeCreateMetaGlobalInput) SetNo(value bool) *ActionNodeCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeCreateMetaGlobalInput) SelectParameters(params ...string) *ActionNodeCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeCreateInput is a type for action input parameters
type ActionNodeCreateInput struct {
	Cpus           int64  `json:"cpus"`
	HypervisorType string `json:"hypervisor_type"`
	Id             int64  `json:"id"`
	IpAddr         string `json:"ip_addr"`
	Location       int64  `json:"location"`
	Maintenance    bool   `json:"maintenance"`
	MaxRx          int64  `json:"max_rx"`
	MaxTx          int64  `json:"max_tx"`
	MaxVps         int64  `json:"max_vps"`
	Name           string `json:"name"`
	NetInterface   string `json:"net_interface"`
	TotalMemory    int64  `json:"total_memory"`
	TotalSwap      int64  `json:"total_swap"`
	Type           string `json:"type"`
	VePrivate      string `json:"ve_private"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCpus sets parameter Cpus to value and selects it for sending
func (in *ActionNodeCreateInput) SetCpus(value int64) *ActionNodeCreateInput {
	in.Cpus = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Cpus"] = nil
	return in
}

// SetHypervisorType sets parameter HypervisorType to value and selects it for sending
func (in *ActionNodeCreateInput) SetHypervisorType(value string) *ActionNodeCreateInput {
	in.HypervisorType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HypervisorType"] = nil
	return in
}

// SetId sets parameter Id to value and selects it for sending
func (in *ActionNodeCreateInput) SetId(value int64) *ActionNodeCreateInput {
	in.Id = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Id"] = nil
	return in
}

// SetIpAddr sets parameter IpAddr to value and selects it for sending
func (in *ActionNodeCreateInput) SetIpAddr(value string) *ActionNodeCreateInput {
	in.IpAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IpAddr"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionNodeCreateInput) SetLocation(value int64) *ActionNodeCreateInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLocationNil(false)
	in._selectedParameters["Location"] = nil
	return in
}

// SetLocationNil sets parameter Location to nil and selects it for sending
func (in *ActionNodeCreateInput) SetLocationNil(set bool) *ActionNodeCreateInput {
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

// SetMaintenance sets parameter Maintenance to value and selects it for sending
func (in *ActionNodeCreateInput) SetMaintenance(value bool) *ActionNodeCreateInput {
	in.Maintenance = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Maintenance"] = nil
	return in
}

// SetMaxRx sets parameter MaxRx to value and selects it for sending
func (in *ActionNodeCreateInput) SetMaxRx(value int64) *ActionNodeCreateInput {
	in.MaxRx = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxRx"] = nil
	return in
}

// SetMaxTx sets parameter MaxTx to value and selects it for sending
func (in *ActionNodeCreateInput) SetMaxTx(value int64) *ActionNodeCreateInput {
	in.MaxTx = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxTx"] = nil
	return in
}

// SetMaxVps sets parameter MaxVps to value and selects it for sending
func (in *ActionNodeCreateInput) SetMaxVps(value int64) *ActionNodeCreateInput {
	in.MaxVps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxVps"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionNodeCreateInput) SetName(value string) *ActionNodeCreateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetNetInterface sets parameter NetInterface to value and selects it for sending
func (in *ActionNodeCreateInput) SetNetInterface(value string) *ActionNodeCreateInput {
	in.NetInterface = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NetInterface"] = nil
	return in
}

// SetTotalMemory sets parameter TotalMemory to value and selects it for sending
func (in *ActionNodeCreateInput) SetTotalMemory(value int64) *ActionNodeCreateInput {
	in.TotalMemory = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TotalMemory"] = nil
	return in
}

// SetTotalSwap sets parameter TotalSwap to value and selects it for sending
func (in *ActionNodeCreateInput) SetTotalSwap(value int64) *ActionNodeCreateInput {
	in.TotalSwap = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TotalSwap"] = nil
	return in
}

// SetType sets parameter Type to value and selects it for sending
func (in *ActionNodeCreateInput) SetType(value string) *ActionNodeCreateInput {
	in.Type = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Type"] = nil
	return in
}

// SetVePrivate sets parameter VePrivate to value and selects it for sending
func (in *ActionNodeCreateInput) SetVePrivate(value string) *ActionNodeCreateInput {
	in.VePrivate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["VePrivate"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeCreateInput) SelectParameters(params ...string) *ActionNodeCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeCreateInput) UnselectParameters(params ...string) *ActionNodeCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeCreateRequest is a type for the entire action request
type ActionNodeCreateRequest struct {
	Node map[string]interface{} `json:"node"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionNodeCreateOutput is a type for action output parameters
type ActionNodeCreateOutput struct {
	Active          bool                      `json:"active"`
	ArcC            int64                     `json:"arc_c"`
	ArcCMax         int64                     `json:"arc_c_max"`
	ArcHitpercent   int64                     `json:"arc_hitpercent"`
	ArcSize         int64                     `json:"arc_size"`
	CpuGuest        float64                   `json:"cpu_guest"`
	CpuIdle         float64                   `json:"cpu_idle"`
	CpuIowait       float64                   `json:"cpu_iowait"`
	CpuIrq          float64                   `json:"cpu_irq"`
	CpuNice         float64                   `json:"cpu_nice"`
	CpuSoftirq      float64                   `json:"cpu_softirq"`
	CpuSystem       float64                   `json:"cpu_system"`
	CpuUser         float64                   `json:"cpu_user"`
	Cpus            int64                     `json:"cpus"`
	DomainName      string                    `json:"domain_name"`
	Fqdn            string                    `json:"fqdn"`
	HypervisorType  string                    `json:"hypervisor_type"`
	Id              int64                     `json:"id"`
	IpAddr          string                    `json:"ip_addr"`
	Kernel          string                    `json:"kernel"`
	Loadavg         float64                   `json:"loadavg"`
	Location        *ActionLocationShowOutput `json:"location"`
	MaxRx           int64                     `json:"max_rx"`
	MaxTx           int64                     `json:"max_tx"`
	MaxVps          int64                     `json:"max_vps"`
	Name            string                    `json:"name"`
	NetInterface    string                    `json:"net_interface"`
	PoolCheckedAt   string                    `json:"pool_checked_at"`
	PoolScan        string                    `json:"pool_scan"`
	PoolScanPercent float64                   `json:"pool_scan_percent"`
	PoolState       string                    `json:"pool_state"`
	PoolStatus      bool                      `json:"pool_status"`
	ProcessCount    int64                     `json:"process_count"`
	Status          bool                      `json:"status"`
	TotalMemory     int64                     `json:"total_memory"`
	TotalSwap       int64                     `json:"total_swap"`
	Type            string                    `json:"type"`
	Uptime          int64                     `json:"uptime"`
	UsedMemory      int64                     `json:"used_memory"`
	UsedSwap        int64                     `json:"used_swap"`
	VePrivate       string                    `json:"ve_private"`
	Version         string                    `json:"version"`
}

// ActionNodeCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionNodeCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionNodeCreateResponse struct {
	Action *ActionNodeCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Node *ActionNodeCreateOutput `json:"node"`
		// Global output metadata
		Meta *ActionNodeCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionNodeCreateOutput
}

// Prepare the action for invocation
func (action *ActionNodeCreate) Prepare() *ActionNodeCreateInvocation {
	return &ActionNodeCreateInvocation{
		Action: action,
		Path:   "/v6.0/nodes",
	}
}

// ActionNodeCreateInvocation is used to configure action for invocation
type ActionNodeCreateInvocation struct {
	// Pointer to the action
	Action *ActionNodeCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeCreateInput
	// Global meta input parameters
	MetaInput *ActionNodeCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeCreateInvocation) NewInput() *ActionNodeCreateInput {
	inv.Input = &ActionNodeCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeCreateInvocation) SetInput(input *ActionNodeCreateInput) *ActionNodeCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeCreateInvocation) NewMetaInput() *ActionNodeCreateMetaGlobalInput {
	inv.MetaInput = &ActionNodeCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeCreateInvocation) SetMetaInput(input *ActionNodeCreateMetaGlobalInput) *ActionNodeCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeCreateInvocation) Call() (*ActionNodeCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionNodeCreateInvocation) callAsBody() (*ActionNodeCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNodeCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Node
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionNodeCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionNodeCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionNodeCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionNodeCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionNodeCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionNodeCreateInvocation) makeAllInputParams() *ActionNodeCreateRequest {
	return &ActionNodeCreateRequest{
		Node: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionNodeCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Cpus") {
			ret["cpus"] = inv.Input.Cpus
		}
		if inv.IsParameterSelected("HypervisorType") {
			ret["hypervisor_type"] = inv.Input.HypervisorType
		}
		if inv.IsParameterSelected("Id") {
			ret["id"] = inv.Input.Id
		}
		if inv.IsParameterSelected("IpAddr") {
			ret["ip_addr"] = inv.Input.IpAddr
		}
		if inv.IsParameterSelected("Location") {
			if inv.IsParameterNil("Location") {
				ret["location"] = nil
			} else {
				ret["location"] = inv.Input.Location
			}
		}
		if inv.IsParameterSelected("Maintenance") {
			ret["maintenance"] = inv.Input.Maintenance
		}
		if inv.IsParameterSelected("MaxRx") {
			ret["max_rx"] = inv.Input.MaxRx
		}
		if inv.IsParameterSelected("MaxTx") {
			ret["max_tx"] = inv.Input.MaxTx
		}
		if inv.IsParameterSelected("MaxVps") {
			ret["max_vps"] = inv.Input.MaxVps
		}
		if inv.IsParameterSelected("Name") {
			ret["name"] = inv.Input.Name
		}
		if inv.IsParameterSelected("NetInterface") {
			ret["net_interface"] = inv.Input.NetInterface
		}
		if inv.IsParameterSelected("TotalMemory") {
			ret["total_memory"] = inv.Input.TotalMemory
		}
		if inv.IsParameterSelected("TotalSwap") {
			ret["total_swap"] = inv.Input.TotalSwap
		}
		if inv.IsParameterSelected("Type") {
			ret["type"] = inv.Input.Type
		}
		if inv.IsParameterSelected("VePrivate") {
			ret["ve_private"] = inv.Input.VePrivate
		}
	}

	return ret
}

func (inv *ActionNodeCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
