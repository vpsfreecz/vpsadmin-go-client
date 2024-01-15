package client

import (
	"strings"
)

// ActionVpsShow is a type for action Vps#Show
type ActionVpsShow struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsShow(client *Client) *ActionVpsShow {
	return &ActionVpsShow{
		Client: client,
	}
}

// ActionVpsShowMetaGlobalInput is a type for action global meta input parameters
type ActionVpsShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsShowMetaGlobalInput) SetIncludes(value string) *ActionVpsShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsShowMetaGlobalInput) SetNo(value bool) *ActionVpsShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsShowMetaGlobalInput) SelectParameters(params ...string) *ActionVpsShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsShowOutput is a type for action output parameters
type ActionVpsShowOutput struct {
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
	Loadavg                 float64                           `json:"loadavg"`
	MaintenanceLock         string                            `json:"maintenance_lock"`
	MaintenanceLockReason   string                            `json:"maintenance_lock_reason"`
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

// Type for action response, including envelope
type ActionVpsShowResponse struct {
	Action *ActionVpsShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Vps *ActionVpsShowOutput `json:"vps"`
	}

	// Action output without the namespace
	Output *ActionVpsShowOutput
}

// Prepare the action for invocation
func (action *ActionVpsShow) Prepare() *ActionVpsShowInvocation {
	return &ActionVpsShowInvocation{
		Action: action,
		Path:   "/v6.0/vpses/{vps_id}",
	}
}

// ActionVpsShowInvocation is used to configure action for invocation
type ActionVpsShowInvocation struct {
	// Pointer to the action
	Action *ActionVpsShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionVpsShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsShowInvocation) SetPathParamInt(param string, value int64) *ActionVpsShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsShowInvocation) SetPathParamString(param string, value string) *ActionVpsShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsShowInvocation) NewMetaInput() *ActionVpsShowMetaGlobalInput {
	inv.MetaInput = &ActionVpsShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsShowInvocation) SetMetaInput(input *ActionVpsShowMetaGlobalInput) *ActionVpsShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsShowInvocation) Call() (*ActionVpsShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsShowInvocation) callAsQuery() (*ActionVpsShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Vps
	}
	return resp, err
}

func (inv *ActionVpsShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
