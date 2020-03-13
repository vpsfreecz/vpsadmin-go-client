package client

import (
)

// ActionNodeOverviewList is a type for action Node#Overview_list
type ActionNodeOverviewList struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeOverviewList(client *Client) *ActionNodeOverviewList {
	return &ActionNodeOverviewList{
		Client: client,
	}
}

// ActionNodeOverviewListMetaGlobalInput is a type for action global meta input parameters
type ActionNodeOverviewListMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeOverviewListMetaGlobalInput) SetNo(value bool) *ActionNodeOverviewListMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeOverviewListMetaGlobalInput) SetIncludes(value string) *ActionNodeOverviewListMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeOverviewListMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeOverviewListMetaGlobalInput) SelectParameters(params ...string) *ActionNodeOverviewListMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeOverviewListMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionNodeOverviewListOutput is a type for action output parameters
type ActionNodeOverviewListOutput struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	DomainName string `json:"domain_name"`
	Type string `json:"type"`
	HypervisorType string `json:"hypervisor_type"`
	Location *ActionLocationShowOutput `json:"location"`
	IpAddr string `json:"ip_addr"`
	NetInterface string `json:"net_interface"`
	MaxTx int64 `json:"max_tx"`
	MaxRx int64 `json:"max_rx"`
	Cpus int64 `json:"cpus"`
	TotalMemory int64 `json:"total_memory"`
	TotalSwap int64 `json:"total_swap"`
	MaxVps int64 `json:"max_vps"`
	VePrivate string `json:"ve_private"`
	Status bool `json:"status"`
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
	CpuGuest float64 `json:"cpu_guest"`
	UsedMemory int64 `json:"used_memory"`
	UsedSwap int64 `json:"used_swap"`
	ArcCMax int64 `json:"arc_c_max"`
	ArcC int64 `json:"arc_c"`
	ArcSize int64 `json:"arc_size"`
	ArcHitpercent int64 `json:"arc_hitpercent"`
	Version string `json:"version"`
	Kernel string `json:"kernel"`
	LastReport string `json:"last_report"`
	VpsRunning int64 `json:"vps_running"`
	VpsStopped int64 `json:"vps_stopped"`
	VpsDeleted int64 `json:"vps_deleted"`
	VpsTotal int64 `json:"vps_total"`
	VpsFree int64 `json:"vps_free"`
	VpsMax int64 `json:"vps_max"`
	MaintenanceLock string `json:"maintenance_lock"`
	MaintenanceLockReason string `json:"maintenance_lock_reason"`
}


// Type for action response, including envelope
type ActionNodeOverviewListResponse struct {
	Action *ActionNodeOverviewList `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Nodes []*ActionNodeOverviewListOutput `json:"nodes"`
	}

	// Action output without the namespace
	Output []*ActionNodeOverviewListOutput
}

// Call the action directly without any path or input parameters
func (action *ActionNodeOverviewList) Call() (*ActionNodeOverviewListResponse, error) {
	return action.Prepare().Call()
}

// Prepare the action for invocation
func (action *ActionNodeOverviewList) Prepare() *ActionNodeOverviewListInvocation {
	return &ActionNodeOverviewListInvocation{
		Action: action,
		Path: "/v6.0/nodes/overview_list",
	}
}

// ActionNodeOverviewListInvocation is used to configure action for invocation
type ActionNodeOverviewListInvocation struct {
	// Pointer to the action
	Action *ActionNodeOverviewList

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNodeOverviewListMetaGlobalInput
}


// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeOverviewListInvocation) NewMetaInput() *ActionNodeOverviewListMetaGlobalInput {
	inv.MetaInput = &ActionNodeOverviewListMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeOverviewListInvocation) SetMetaInput(input *ActionNodeOverviewListMetaGlobalInput) *ActionNodeOverviewListInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeOverviewListInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeOverviewListInvocation) Call() (*ActionNodeOverviewListResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionNodeOverviewListInvocation) callAsQuery() (*ActionNodeOverviewListResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeOverviewListResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Nodes
	}
	return resp, err
}




func (inv *ActionNodeOverviewListInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

