package client

import (
	"strings"
)

// ActionVpsStatusShow is a type for action Vps.Status#Show
type ActionVpsStatusShow struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsStatusShow(client *Client) *ActionVpsStatusShow {
	return &ActionVpsStatusShow{
		Client: client,
	}
}

// ActionVpsStatusShowMetaGlobalInput is a type for action global meta input parameters
type ActionVpsStatusShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsStatusShowMetaGlobalInput) SetIncludes(value string) *ActionVpsStatusShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsStatusShowMetaGlobalInput) SetNo(value bool) *ActionVpsStatusShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsStatusShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsStatusShowMetaGlobalInput) SelectParameters(params ...string) *ActionVpsStatusShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsStatusShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionVpsStatusShowOutput is a type for action output parameters
type ActionVpsStatusShowOutput struct {
	CpuIdle float64 `json:"cpu_idle"`
	CpuIowait float64 `json:"cpu_iowait"`
	CpuIrq float64 `json:"cpu_irq"`
	CpuNice float64 `json:"cpu_nice"`
	CpuSoftirq float64 `json:"cpu_softirq"`
	CpuSystem float64 `json:"cpu_system"`
	CpuUser float64 `json:"cpu_user"`
	Cpus int64 `json:"cpus"`
	CreatedAt string `json:"created_at"`
	Id int64 `json:"id"`
	InRescueMode bool `json:"in_rescue_mode"`
	IsRunning bool `json:"is_running"`
	Loadavg float64 `json:"loadavg"`
	ProcessCount int64 `json:"process_count"`
	Status bool `json:"status"`
	TotalMemory int64 `json:"total_memory"`
	TotalSwap int64 `json:"total_swap"`
	Uptime int64 `json:"uptime"`
	UsedMemory int64 `json:"used_memory"`
	UsedSwap int64 `json:"used_swap"`
}


// Type for action response, including envelope
type ActionVpsStatusShowResponse struct {
	Action *ActionVpsStatusShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Status *ActionVpsStatusShowOutput `json:"status"`
	}

	// Action output without the namespace
	Output *ActionVpsStatusShowOutput
}


// Prepare the action for invocation
func (action *ActionVpsStatusShow) Prepare() *ActionVpsStatusShowInvocation {
	return &ActionVpsStatusShowInvocation{
		Action: action,
		Path: "/v6.0/vpses/{vps_id}/statuses/{status_id}",
	}
}

// ActionVpsStatusShowInvocation is used to configure action for invocation
type ActionVpsStatusShowInvocation struct {
	// Pointer to the action
	Action *ActionVpsStatusShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionVpsStatusShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsStatusShowInvocation) SetPathParamInt(param string, value int64) *ActionVpsStatusShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsStatusShowInvocation) SetPathParamString(param string, value string) *ActionVpsStatusShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsStatusShowInvocation) NewMetaInput() *ActionVpsStatusShowMetaGlobalInput {
	inv.MetaInput = &ActionVpsStatusShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsStatusShowInvocation) SetMetaInput(input *ActionVpsStatusShowMetaGlobalInput) *ActionVpsStatusShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsStatusShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsStatusShowInvocation) Call() (*ActionVpsStatusShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsStatusShowInvocation) callAsQuery() (*ActionVpsStatusShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsStatusShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Status
	}
	return resp, err
}




func (inv *ActionVpsStatusShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}

