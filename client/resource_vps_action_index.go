package client

import (
)

// ActionVpsIndex is a type for action Vps#Index
type ActionVpsIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsIndex(client *Client) *ActionVpsIndex {
	return &ActionVpsIndex{
		Client: client,
	}
}

// ActionVpsIndexMetaGlobalInput is a type for action global meta input parameters
type ActionVpsIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsIndexMetaGlobalInput) SetNo(value bool) *ActionVpsIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionVpsIndexMetaGlobalInput) SetCount(value bool) *ActionVpsIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsIndexMetaGlobalInput) SetIncludes(value string) *ActionVpsIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsIndexMetaGlobalInput) SelectParameters(params ...string) *ActionVpsIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsIndexInput is a type for action input parameters
type ActionVpsIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	User int64 `json:"user"`
	Node int64 `json:"node"`
	Location int64 `json:"location"`
	Environment int64 `json:"environment"`
	UserNamespaceMap int64 `json:"user_namespace_map"`
	OsTemplate int64 `json:"os_template"`
	ObjectState string `json:"object_state"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionVpsIndexInput) SetOffset(value int64) *ActionVpsIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionVpsIndexInput) SetLimit(value int64) *ActionVpsIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetUser sets parameter User to value and selects it for sending
func (in *ActionVpsIndexInput) SetUser(value int64) *ActionVpsIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}
// SetNode sets parameter Node to value and selects it for sending
func (in *ActionVpsIndexInput) SetNode(value int64) *ActionVpsIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}
// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionVpsIndexInput) SetLocation(value int64) *ActionVpsIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}
// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionVpsIndexInput) SetEnvironment(value int64) *ActionVpsIndexInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Environment"] = nil
	return in
}
// SetUserNamespaceMap sets parameter UserNamespaceMap to value and selects it for sending
func (in *ActionVpsIndexInput) SetUserNamespaceMap(value int64) *ActionVpsIndexInput {
	in.UserNamespaceMap = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserNamespaceMap"] = nil
	return in
}
// SetOsTemplate sets parameter OsTemplate to value and selects it for sending
func (in *ActionVpsIndexInput) SetOsTemplate(value int64) *ActionVpsIndexInput {
	in.OsTemplate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OsTemplate"] = nil
	return in
}
// SetObjectState sets parameter ObjectState to value and selects it for sending
func (in *ActionVpsIndexInput) SetObjectState(value string) *ActionVpsIndexInput {
	in.ObjectState = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ObjectState"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsIndexInput) SelectParameters(params ...string) *ActionVpsIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionVpsIndexOutput is a type for action output parameters
type ActionVpsIndexOutput struct {
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
	MaintenanceLock string `json:"maintenance_lock"`
	MaintenanceLockReason string `json:"maintenance_lock_reason"`
	ObjectState string `json:"object_state"`
	ExpirationDate string `json:"expiration_date"`
}


// Type for action response, including envelope
type ActionVpsIndexResponse struct {
	Action *ActionVpsIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Vpses []*ActionVpsIndexOutput `json:"vpses"`
	}

	// Action output without the namespace
	Output []*ActionVpsIndexOutput
}


// Prepare the action for invocation
func (action *ActionVpsIndex) Prepare() *ActionVpsIndexInvocation {
	return &ActionVpsIndexInvocation{
		Action: action,
		Path: "/v5.0/vpses",
	}
}

// ActionVpsIndexInvocation is used to configure action for invocation
type ActionVpsIndexInvocation struct {
	// Pointer to the action
	Action *ActionVpsIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsIndexInput
	// Global meta input parameters
	MetaInput *ActionVpsIndexMetaGlobalInput
}


// SetInput provides input parameters to send to the API
func (inv *ActionVpsIndexInvocation) SetInput(input *ActionVpsIndexInput) *ActionVpsIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsIndexInvocation) SetMetaInput(input *ActionVpsIndexMetaGlobalInput) *ActionVpsIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsIndexInvocation) Call() (*ActionVpsIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsIndexInvocation) callAsQuery() (*ActionVpsIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Vpses
	}
	return resp, err
}



func (inv *ActionVpsIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["vps[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["vps[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("User") {
			ret["vps[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("Node") {
			ret["vps[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("Location") {
			ret["vps[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("Environment") {
			ret["vps[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("UserNamespaceMap") {
			ret["vps[user_namespace_map]"] = convertInt64ToString(inv.Input.UserNamespaceMap)
		}
		if inv.IsParameterSelected("OsTemplate") {
			ret["vps[os_template]"] = convertInt64ToString(inv.Input.OsTemplate)
		}
		if inv.IsParameterSelected("ObjectState") {
			ret["vps[object_state]"] = inv.Input.ObjectState
		}
	}
}

func (inv *ActionVpsIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

