package client

import ()

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
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsIndexMetaGlobalInput) SetNo(value bool) *ActionVpsIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
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
	Environment      int64  `json:"environment"`
	HostnameAny      string `json:"hostname_any"`
	HostnameExact    string `json:"hostname_exact"`
	Limit            int64  `json:"limit"`
	Location         int64  `json:"location"`
	Node             int64  `json:"node"`
	ObjectState      string `json:"object_state"`
	Offset           int64  `json:"offset"`
	OsTemplate       int64  `json:"os_template"`
	User             int64  `json:"user"`
	UserNamespaceMap int64  `json:"user_namespace_map"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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

// SetHostnameAny sets parameter HostnameAny to value and selects it for sending
func (in *ActionVpsIndexInput) SetHostnameAny(value string) *ActionVpsIndexInput {
	in.HostnameAny = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HostnameAny"] = nil
	return in
}

// SetHostnameExact sets parameter HostnameExact to value and selects it for sending
func (in *ActionVpsIndexInput) SetHostnameExact(value string) *ActionVpsIndexInput {
	in.HostnameExact = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HostnameExact"] = nil
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

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionVpsIndexInput) SetLocation(value int64) *ActionVpsIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
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

// SetObjectState sets parameter ObjectState to value and selects it for sending
func (in *ActionVpsIndexInput) SetObjectState(value string) *ActionVpsIndexInput {
	in.ObjectState = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ObjectState"] = nil
	return in
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

// SetOsTemplate sets parameter OsTemplate to value and selects it for sending
func (in *ActionVpsIndexInput) SetOsTemplate(value int64) *ActionVpsIndexInput {
	in.OsTemplate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OsTemplate"] = nil
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

// SetUserNamespaceMap sets parameter UserNamespaceMap to value and selects it for sending
func (in *ActionVpsIndexInput) SetUserNamespaceMap(value int64) *ActionVpsIndexInput {
	in.UserNamespaceMap = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserNamespaceMap"] = nil
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
	Config                string                       `json:"config"`
	Cpu                   int64                        `json:"cpu"`
	CpuIdle               float64                      `json:"cpu_idle"`
	CpuIowait             float64                      `json:"cpu_iowait"`
	CpuIrq                float64                      `json:"cpu_irq"`
	CpuLimit              int64                        `json:"cpu_limit"`
	CpuNice               float64                      `json:"cpu_nice"`
	CpuSoftirq            float64                      `json:"cpu_softirq"`
	CpuSystem             float64                      `json:"cpu_system"`
	CpuUser               float64                      `json:"cpu_user"`
	CreatedAt             string                       `json:"created_at"`
	Dataset               *ActionDatasetShowOutput     `json:"dataset"`
	Diskspace             int64                        `json:"diskspace"`
	DnsResolver           *ActionDnsResolverShowOutput `json:"dns_resolver"`
	ExpirationDate        string                       `json:"expiration_date"`
	Hostname              string                       `json:"hostname"`
	Id                    int64                        `json:"id"`
	InRescueMode          bool                         `json:"in_rescue_mode"`
	Info                  string                       `json:"info"`
	IsRunning             bool                         `json:"is_running"`
	Loadavg               float64                      `json:"loadavg"`
	MaintenanceLock       string                       `json:"maintenance_lock"`
	MaintenanceLockReason string                       `json:"maintenance_lock_reason"`
	ManageHostname        bool                         `json:"manage_hostname"`
	Memory                int64                        `json:"memory"`
	Node                  *ActionNodeShowOutput        `json:"node"`
	ObjectState           string                       `json:"object_state"`
	Onboot                bool                         `json:"onboot"`
	Onstartall            bool                         `json:"onstartall"`
	OsTemplate            *ActionOsTemplateShowOutput  `json:"os_template"`
	ProcessCount          int64                        `json:"process_count"`
	Swap                  int64                        `json:"swap"`
	Uptime                int64                        `json:"uptime"`
	UsedDiskspace         int64                        `json:"used_diskspace"`
	UsedMemory            int64                        `json:"used_memory"`
	UsedSwap              int64                        `json:"used_swap"`
	User                  *ActionUserShowOutput        `json:"user"`
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
		Path:   "/v6.0/vpses",
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

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsIndexInvocation) NewInput() *ActionVpsIndexInput {
	inv.Input = &ActionVpsIndexInput{}
	return inv.Input
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

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsIndexInvocation) NewMetaInput() *ActionVpsIndexMetaGlobalInput {
	inv.MetaInput = &ActionVpsIndexMetaGlobalInput{}
	return inv.MetaInput
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
		if inv.IsParameterSelected("Environment") {
			ret["vps[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("HostnameAny") {
			ret["vps[hostname_any]"] = inv.Input.HostnameAny
		}
		if inv.IsParameterSelected("HostnameExact") {
			ret["vps[hostname_exact]"] = inv.Input.HostnameExact
		}
		if inv.IsParameterSelected("Limit") {
			ret["vps[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Location") {
			ret["vps[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("Node") {
			ret["vps[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("ObjectState") {
			ret["vps[object_state]"] = inv.Input.ObjectState
		}
		if inv.IsParameterSelected("Offset") {
			ret["vps[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("OsTemplate") {
			ret["vps[os_template]"] = convertInt64ToString(inv.Input.OsTemplate)
		}
		if inv.IsParameterSelected("User") {
			ret["vps[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("UserNamespaceMap") {
			ret["vps[user_namespace_map]"] = convertInt64ToString(inv.Input.UserNamespaceMap)
		}
	}
}

func (inv *ActionVpsIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
