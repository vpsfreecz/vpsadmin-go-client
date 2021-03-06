package client

import (
	"strings"
)

// ActionNodeShow is a type for action Node#Show
type ActionNodeShow struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeShow(client *Client) *ActionNodeShow {
	return &ActionNodeShow{
		Client: client,
	}
}

// ActionNodeShowMetaGlobalInput is a type for action global meta input parameters
type ActionNodeShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeShowMetaGlobalInput) SetNo(value bool) *ActionNodeShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeShowMetaGlobalInput) SetIncludes(value string) *ActionNodeShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeShowMetaGlobalInput) SelectParameters(params ...string) *ActionNodeShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionNodeShowOutput is a type for action output parameters
type ActionNodeShowOutput struct {
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
	MaintenanceLock string `json:"maintenance_lock"`
	MaintenanceLockReason string `json:"maintenance_lock_reason"`
}


// Type for action response, including envelope
type ActionNodeShowResponse struct {
	Action *ActionNodeShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Node *ActionNodeShowOutput `json:"node"`
	}

	// Action output without the namespace
	Output *ActionNodeShowOutput
}


// Prepare the action for invocation
func (action *ActionNodeShow) Prepare() *ActionNodeShowInvocation {
	return &ActionNodeShowInvocation{
		Action: action,
		Path: "/v6.0/nodes/{node_id}",
	}
}

// ActionNodeShowInvocation is used to configure action for invocation
type ActionNodeShowInvocation struct {
	// Pointer to the action
	Action *ActionNodeShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNodeShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNodeShowInvocation) SetPathParamInt(param string, value int64) *ActionNodeShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNodeShowInvocation) SetPathParamString(param string, value string) *ActionNodeShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeShowInvocation) NewMetaInput() *ActionNodeShowMetaGlobalInput {
	inv.MetaInput = &ActionNodeShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeShowInvocation) SetMetaInput(input *ActionNodeShowMetaGlobalInput) *ActionNodeShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeShowInvocation) Call() (*ActionNodeShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionNodeShowInvocation) callAsQuery() (*ActionNodeShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Node
	}
	return resp, err
}




func (inv *ActionNodeShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

