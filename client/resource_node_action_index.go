package client

import ()

// ActionNodeIndex is a type for action Node#Index
type ActionNodeIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeIndex(client *Client) *ActionNodeIndex {
	return &ActionNodeIndex{
		Client: client,
	}
}

// ActionNodeIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeIndexMetaGlobalInput) SetCount(value bool) *ActionNodeIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeIndexMetaGlobalInput) SetNo(value bool) *ActionNodeIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeIndexInput is a type for action input parameters
type ActionNodeIndexInput struct {
	Environment    int64  `json:"environment"`
	HypervisorType string `json:"hypervisor_type"`
	Limit          int64  `json:"limit"`
	Location       int64  `json:"location"`
	Offset         int64  `json:"offset"`
	State          string `json:"state"`
	Type           string `json:"type"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionNodeIndexInput) SetEnvironment(value int64) *ActionNodeIndexInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEnvironmentNil(false)
	in._selectedParameters["Environment"] = nil
	return in
}

// SetEnvironmentNil sets parameter Environment to nil and selects it for sending
func (in *ActionNodeIndexInput) SetEnvironmentNil(set bool) *ActionNodeIndexInput {
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

// SetHypervisorType sets parameter HypervisorType to value and selects it for sending
func (in *ActionNodeIndexInput) SetHypervisorType(value string) *ActionNodeIndexInput {
	in.HypervisorType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HypervisorType"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeIndexInput) SetLimit(value int64) *ActionNodeIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionNodeIndexInput) SetLocation(value int64) *ActionNodeIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLocationNil(false)
	in._selectedParameters["Location"] = nil
	return in
}

// SetLocationNil sets parameter Location to nil and selects it for sending
func (in *ActionNodeIndexInput) SetLocationNil(set bool) *ActionNodeIndexInput {
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

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionNodeIndexInput) SetOffset(value int64) *ActionNodeIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetState sets parameter State to value and selects it for sending
func (in *ActionNodeIndexInput) SetState(value string) *ActionNodeIndexInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["State"] = nil
	return in
}

// SetType sets parameter Type to value and selects it for sending
func (in *ActionNodeIndexInput) SetType(value string) *ActionNodeIndexInput {
	in.Type = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Type"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeIndexInput) SelectParameters(params ...string) *ActionNodeIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeIndexInput) UnselectParameters(params ...string) *ActionNodeIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeIndexOutput is a type for action output parameters
type ActionNodeIndexOutput struct {
	Active                bool                      `json:"active"`
	ArcC                  int64                     `json:"arc_c"`
	ArcCMax               int64                     `json:"arc_c_max"`
	ArcHitpercent         int64                     `json:"arc_hitpercent"`
	ArcSize               int64                     `json:"arc_size"`
	CgroupVersion         string                    `json:"cgroup_version"`
	CpuGuest              float64                   `json:"cpu_guest"`
	CpuIdle               float64                   `json:"cpu_idle"`
	CpuIowait             float64                   `json:"cpu_iowait"`
	CpuIrq                float64                   `json:"cpu_irq"`
	CpuNice               float64                   `json:"cpu_nice"`
	CpuSoftirq            float64                   `json:"cpu_softirq"`
	CpuSystem             float64                   `json:"cpu_system"`
	CpuUser               float64                   `json:"cpu_user"`
	Cpus                  int64                     `json:"cpus"`
	DomainName            string                    `json:"domain_name"`
	Fqdn                  string                    `json:"fqdn"`
	HypervisorType        string                    `json:"hypervisor_type"`
	Id                    int64                     `json:"id"`
	IpAddr                string                    `json:"ip_addr"`
	Kernel                string                    `json:"kernel"`
	Loadavg               float64                   `json:"loadavg"`
	Location              *ActionLocationShowOutput `json:"location"`
	MaintenanceLock       string                    `json:"maintenance_lock"`
	MaintenanceLockReason string                    `json:"maintenance_lock_reason"`
	MaxRx                 int64                     `json:"max_rx"`
	MaxTx                 int64                     `json:"max_tx"`
	MaxVps                int64                     `json:"max_vps"`
	Name                  string                    `json:"name"`
	PoolCheckedAt         string                    `json:"pool_checked_at"`
	PoolScan              string                    `json:"pool_scan"`
	PoolScanPercent       float64                   `json:"pool_scan_percent"`
	PoolState             string                    `json:"pool_state"`
	PoolStatus            bool                      `json:"pool_status"`
	ProcessCount          int64                     `json:"process_count"`
	Status                bool                      `json:"status"`
	TotalMemory           int64                     `json:"total_memory"`
	TotalSwap             int64                     `json:"total_swap"`
	Type                  string                    `json:"type"`
	Uptime                int64                     `json:"uptime"`
	UsedMemory            int64                     `json:"used_memory"`
	UsedSwap              int64                     `json:"used_swap"`
	Version               string                    `json:"version"`
}

// Type for action response, including envelope
type ActionNodeIndexResponse struct {
	Action *ActionNodeIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Nodes []*ActionNodeIndexOutput `json:"nodes"`
	}

	// Action output without the namespace
	Output []*ActionNodeIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeIndex) Prepare() *ActionNodeIndexInvocation {
	return &ActionNodeIndexInvocation{
		Action: action,
		Path:   "/v6.0/nodes",
	}
}

// ActionNodeIndexInvocation is used to configure action for invocation
type ActionNodeIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeIndexInvocation) NewInput() *ActionNodeIndexInput {
	inv.Input = &ActionNodeIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeIndexInvocation) SetInput(input *ActionNodeIndexInput) *ActionNodeIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeIndexInvocation) NewMetaInput() *ActionNodeIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeIndexInvocation) SetMetaInput(input *ActionNodeIndexMetaGlobalInput) *ActionNodeIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeIndexInvocation) Call() (*ActionNodeIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionNodeIndexInvocation) callAsQuery() (*ActionNodeIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Nodes
	}
	return resp, err
}

func (inv *ActionNodeIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			ret["node[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("HypervisorType") {
			ret["node[hypervisor_type]"] = inv.Input.HypervisorType
		}
		if inv.IsParameterSelected("Limit") {
			ret["node[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Location") {
			ret["node[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("Offset") {
			ret["node[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("State") {
			ret["node[state]"] = inv.Input.State
		}
		if inv.IsParameterSelected("Type") {
			ret["node[type]"] = inv.Input.Type
		}
	}
}

func (inv *ActionNodeIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
