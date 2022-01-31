package client

import (
	"strings"
)

// ActionNodeStatusIndex is a type for action Node.Status#Index
type ActionNodeStatusIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeStatusIndex(client *Client) *ActionNodeStatusIndex {
	return &ActionNodeStatusIndex{
		Client: client,
	}
}

// ActionNodeStatusIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeStatusIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeStatusIndexMetaGlobalInput) SetCount(value bool) *ActionNodeStatusIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeStatusIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeStatusIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeStatusIndexMetaGlobalInput) SetNo(value bool) *ActionNodeStatusIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeStatusIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeStatusIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeStatusIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeStatusIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeStatusIndexInput is a type for action input parameters
type ActionNodeStatusIndexInput struct {
	From   string `json:"from"`
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
	To     string `json:"to"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionNodeStatusIndexInput) SetFrom(value string) *ActionNodeStatusIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeStatusIndexInput) SetLimit(value int64) *ActionNodeStatusIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionNodeStatusIndexInput) SetOffset(value int64) *ActionNodeStatusIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNodeStatusIndexInput) SetTo(value string) *ActionNodeStatusIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeStatusIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeStatusIndexInput) SelectParameters(params ...string) *ActionNodeStatusIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeStatusIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeStatusIndexInput) UnselectParameters(params ...string) *ActionNodeStatusIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeStatusIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeStatusIndexOutput is a type for action output parameters
type ActionNodeStatusIndexOutput struct {
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
	Loadavg       float64 `json:"loadavg"`
	ProcessCount  int64   `json:"process_count"`
	TotalMemory   int64   `json:"total_memory"`
	TotalSwap     int64   `json:"total_swap"`
	Uptime        int64   `json:"uptime"`
	UsedMemory    int64   `json:"used_memory"`
	UsedSwap      int64   `json:"used_swap"`
	Version       string  `json:"version"`
}

// Type for action response, including envelope
type ActionNodeStatusIndexResponse struct {
	Action *ActionNodeStatusIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Statuses []*ActionNodeStatusIndexOutput `json:"statuses"`
	}

	// Action output without the namespace
	Output []*ActionNodeStatusIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeStatusIndex) Prepare() *ActionNodeStatusIndexInvocation {
	return &ActionNodeStatusIndexInvocation{
		Action: action,
		Path:   "/v6.0/nodes/{node_id}/statuses",
	}
}

// ActionNodeStatusIndexInvocation is used to configure action for invocation
type ActionNodeStatusIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeStatusIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeStatusIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeStatusIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNodeStatusIndexInvocation) SetPathParamInt(param string, value int64) *ActionNodeStatusIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNodeStatusIndexInvocation) SetPathParamString(param string, value string) *ActionNodeStatusIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeStatusIndexInvocation) NewInput() *ActionNodeStatusIndexInput {
	inv.Input = &ActionNodeStatusIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeStatusIndexInvocation) SetInput(input *ActionNodeStatusIndexInput) *ActionNodeStatusIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeStatusIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeStatusIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeStatusIndexInvocation) NewMetaInput() *ActionNodeStatusIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeStatusIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeStatusIndexInvocation) SetMetaInput(input *ActionNodeStatusIndexMetaGlobalInput) *ActionNodeStatusIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeStatusIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeStatusIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeStatusIndexInvocation) Call() (*ActionNodeStatusIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionNodeStatusIndexInvocation) callAsQuery() (*ActionNodeStatusIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeStatusIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Statuses
	}
	return resp, err
}

func (inv *ActionNodeStatusIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("From") {
			ret["status[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("Limit") {
			ret["status[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["status[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("To") {
			ret["status[to]"] = inv.Input.To
		}
	}
}

func (inv *ActionNodeStatusIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
