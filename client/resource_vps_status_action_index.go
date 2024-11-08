package client

import (
	"strings"
)

// ActionVpsStatusIndex is a type for action Vps.Status#Index
type ActionVpsStatusIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsStatusIndex(client *Client) *ActionVpsStatusIndex {
	return &ActionVpsStatusIndex{
		Client: client,
	}
}

// ActionVpsStatusIndexMetaGlobalInput is a type for action global meta input parameters
type ActionVpsStatusIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionVpsStatusIndexMetaGlobalInput) SetCount(value bool) *ActionVpsStatusIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsStatusIndexMetaGlobalInput) SetIncludes(value string) *ActionVpsStatusIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsStatusIndexMetaGlobalInput) SetNo(value bool) *ActionVpsStatusIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsStatusIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsStatusIndexMetaGlobalInput) SelectParameters(params ...string) *ActionVpsStatusIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsStatusIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsStatusIndexInput is a type for action input parameters
type ActionVpsStatusIndexInput struct {
	From      string `json:"from"`
	FromId    int64  `json:"from_id"`
	IsRunning bool   `json:"is_running"`
	Limit     int64  `json:"limit"`
	Status    bool   `json:"status"`
	To        string `json:"to"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionVpsStatusIndexInput) SetFrom(value string) *ActionVpsStatusIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionVpsStatusIndexInput) SetFromId(value int64) *ActionVpsStatusIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetIsRunning sets parameter IsRunning to value and selects it for sending
func (in *ActionVpsStatusIndexInput) SetIsRunning(value bool) *ActionVpsStatusIndexInput {
	in.IsRunning = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IsRunning"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionVpsStatusIndexInput) SetLimit(value int64) *ActionVpsStatusIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetStatus sets parameter Status to value and selects it for sending
func (in *ActionVpsStatusIndexInput) SetStatus(value bool) *ActionVpsStatusIndexInput {
	in.Status = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Status"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionVpsStatusIndexInput) SetTo(value string) *ActionVpsStatusIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsStatusIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsStatusIndexInput) SelectParameters(params ...string) *ActionVpsStatusIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsStatusIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsStatusIndexInput) UnselectParameters(params ...string) *ActionVpsStatusIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsStatusIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsStatusIndexOutput is a type for action output parameters
type ActionVpsStatusIndexOutput struct {
	CpuIdle      float64 `json:"cpu_idle"`
	CpuIowait    float64 `json:"cpu_iowait"`
	CpuIrq       float64 `json:"cpu_irq"`
	CpuNice      float64 `json:"cpu_nice"`
	CpuSoftirq   float64 `json:"cpu_softirq"`
	CpuSystem    float64 `json:"cpu_system"`
	CpuUser      float64 `json:"cpu_user"`
	Cpus         int64   `json:"cpus"`
	CreatedAt    string  `json:"created_at"`
	Id           int64   `json:"id"`
	InRescueMode bool    `json:"in_rescue_mode"`
	IsRunning    bool    `json:"is_running"`
	Loadavg1     float64 `json:"loadavg1"`
	Loadavg15    float64 `json:"loadavg15"`
	Loadavg5     float64 `json:"loadavg5"`
	ProcessCount int64   `json:"process_count"`
	Status       bool    `json:"status"`
	TotalMemory  int64   `json:"total_memory"`
	TotalSwap    int64   `json:"total_swap"`
	Uptime       int64   `json:"uptime"`
	UsedMemory   int64   `json:"used_memory"`
	UsedSwap     int64   `json:"used_swap"`
}

// Type for action response, including envelope
type ActionVpsStatusIndexResponse struct {
	Action *ActionVpsStatusIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Statuses []*ActionVpsStatusIndexOutput `json:"statuses"`
	}

	// Action output without the namespace
	Output []*ActionVpsStatusIndexOutput
}

// Prepare the action for invocation
func (action *ActionVpsStatusIndex) Prepare() *ActionVpsStatusIndexInvocation {
	return &ActionVpsStatusIndexInvocation{
		Action: action,
		Path:   "/v7.0/vpses/{vps_id}/statuses",
	}
}

// ActionVpsStatusIndexInvocation is used to configure action for invocation
type ActionVpsStatusIndexInvocation struct {
	// Pointer to the action
	Action *ActionVpsStatusIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsStatusIndexInput
	// Global meta input parameters
	MetaInput *ActionVpsStatusIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsStatusIndexInvocation) SetPathParamInt(param string, value int64) *ActionVpsStatusIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsStatusIndexInvocation) SetPathParamString(param string, value string) *ActionVpsStatusIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsStatusIndexInvocation) NewInput() *ActionVpsStatusIndexInput {
	inv.Input = &ActionVpsStatusIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsStatusIndexInvocation) SetInput(input *ActionVpsStatusIndexInput) *ActionVpsStatusIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsStatusIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsStatusIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsStatusIndexInvocation) NewMetaInput() *ActionVpsStatusIndexMetaGlobalInput {
	inv.MetaInput = &ActionVpsStatusIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsStatusIndexInvocation) SetMetaInput(input *ActionVpsStatusIndexMetaGlobalInput) *ActionVpsStatusIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsStatusIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsStatusIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsStatusIndexInvocation) Call() (*ActionVpsStatusIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsStatusIndexInvocation) callAsQuery() (*ActionVpsStatusIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsStatusIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Statuses
	}
	return resp, err
}

func (inv *ActionVpsStatusIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("From") {
			ret["status[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("FromId") {
			ret["status[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("IsRunning") {
			ret["status[is_running]"] = convertBoolToString(inv.Input.IsRunning)
		}
		if inv.IsParameterSelected("Limit") {
			ret["status[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Status") {
			ret["status[status]"] = convertBoolToString(inv.Input.Status)
		}
		if inv.IsParameterSelected("To") {
			ret["status[to]"] = inv.Input.To
		}
	}
}

func (inv *ActionVpsStatusIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
