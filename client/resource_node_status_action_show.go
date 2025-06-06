package client

import (
	"strings"
)

// ActionNodeStatusShow is a type for action Node.Status#Show
type ActionNodeStatusShow struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeStatusShow(client *Client) *ActionNodeStatusShow {
	return &ActionNodeStatusShow{
		Client: client,
	}
}

// ActionNodeStatusShowMetaGlobalInput is a type for action global meta input parameters
type ActionNodeStatusShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeStatusShowMetaGlobalInput) SetIncludes(value string) *ActionNodeStatusShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeStatusShowMetaGlobalInput) SetNo(value bool) *ActionNodeStatusShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeStatusShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeStatusShowMetaGlobalInput) SelectParameters(params ...string) *ActionNodeStatusShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeStatusShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeStatusShowOutput is a type for action output parameters
type ActionNodeStatusShowOutput struct {
	ArcC          int64   `json:"arc_c"`
	ArcCMax       int64   `json:"arc_c_max"`
	ArcHitpercent float64 `json:"arc_hitpercent"`
	ArcSize       int64   `json:"arc_size"`
	CpuGuest      float64 `json:"cpu_guest"`
	CpuIdle       float64 `json:"cpu_idle"`
	CpuIowait     float64 `json:"cpu_iowait"`
	CpuIrq        float64 `json:"cpu_irq"`
	CpuNice       float64 `json:"cpu_nice"`
	CpuSoftirq    float64 `json:"cpu_softirq"`
	CpuSystem     float64 `json:"cpu_system"`
	CpuUser       float64 `json:"cpu_user"`
	Cpus          int64   `json:"cpus"`
	CreatedAt     string  `json:"created_at"`
	Id            int64   `json:"id"`
	Kernel        string  `json:"kernel"`
	Loadavg1      float64 `json:"loadavg1"`
	Loadavg15     float64 `json:"loadavg15"`
	Loadavg5      float64 `json:"loadavg5"`
	ProcessCount  int64   `json:"process_count"`
	TotalMemory   int64   `json:"total_memory"`
	TotalSwap     int64   `json:"total_swap"`
	Uptime        int64   `json:"uptime"`
	UsedMemory    int64   `json:"used_memory"`
	UsedSwap      int64   `json:"used_swap"`
	Version       string  `json:"version"`
}

// Type for action response, including envelope
type ActionNodeStatusShowResponse struct {
	Action *ActionNodeStatusShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Status *ActionNodeStatusShowOutput `json:"status"`
	}

	// Action output without the namespace
	Output *ActionNodeStatusShowOutput
}

// Prepare the action for invocation
func (action *ActionNodeStatusShow) Prepare() *ActionNodeStatusShowInvocation {
	return &ActionNodeStatusShowInvocation{
		Action: action,
		Path:   "/v7.0/nodes/{node_id}/statuses/{status_id}",
	}
}

// ActionNodeStatusShowInvocation is used to configure action for invocation
type ActionNodeStatusShowInvocation struct {
	// Pointer to the action
	Action *ActionNodeStatusShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNodeStatusShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNodeStatusShowInvocation) SetPathParamInt(param string, value int64) *ActionNodeStatusShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNodeStatusShowInvocation) SetPathParamString(param string, value string) *ActionNodeStatusShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeStatusShowInvocation) NewMetaInput() *ActionNodeStatusShowMetaGlobalInput {
	inv.MetaInput = &ActionNodeStatusShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeStatusShowInvocation) SetMetaInput(input *ActionNodeStatusShowMetaGlobalInput) *ActionNodeStatusShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeStatusShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeStatusShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeStatusShowInvocation) Call() (*ActionNodeStatusShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionNodeStatusShowInvocation) callAsQuery() (*ActionNodeStatusShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeStatusShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Status
	}
	return resp, err
}

func (inv *ActionNodeStatusShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
