package client

import (
	"strings"
)

// ActionOomReportTaskIndex is a type for action Oom_report.Task#Index
type ActionOomReportTaskIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionOomReportTaskIndex(client *Client) *ActionOomReportTaskIndex {
	return &ActionOomReportTaskIndex{
		Client: client,
	}
}

// ActionOomReportTaskIndexMetaGlobalInput is a type for action global meta input parameters
type ActionOomReportTaskIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionOomReportTaskIndexMetaGlobalInput) SetCount(value bool) *ActionOomReportTaskIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOomReportTaskIndexMetaGlobalInput) SetIncludes(value string) *ActionOomReportTaskIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOomReportTaskIndexMetaGlobalInput) SetNo(value bool) *ActionOomReportTaskIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOomReportTaskIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportTaskIndexMetaGlobalInput) SelectParameters(params ...string) *ActionOomReportTaskIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOomReportTaskIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOomReportTaskIndexInput is a type for action input parameters
type ActionOomReportTaskIndexInput struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionOomReportTaskIndexInput) SetLimit(value int64) *ActionOomReportTaskIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionOomReportTaskIndexInput) SetOffset(value int64) *ActionOomReportTaskIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SelectParameters sets parameters from ActionOomReportTaskIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportTaskIndexInput) SelectParameters(params ...string) *ActionOomReportTaskIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionOomReportTaskIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOomReportTaskIndexInput) UnselectParameters(params ...string) *ActionOomReportTaskIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionOomReportTaskIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOomReportTaskIndexOutput is a type for action output parameters
type ActionOomReportTaskIndexOutput struct {
	HostPid       int64  `json:"host_pid"`
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	OomScoreAdj   int64  `json:"oom_score_adj"`
	PgtablesBytes int64  `json:"pgtables_bytes"`
	Rss           int64  `json:"rss"`
	Swapents      int64  `json:"swapents"`
	Tgid          int64  `json:"tgid"`
	TotalVm       int64  `json:"total_vm"`
	VpsPid        int64  `json:"vps_pid"`
	VpsUid        int64  `json:"vps_uid"`
}

// Type for action response, including envelope
type ActionOomReportTaskIndexResponse struct {
	Action *ActionOomReportTaskIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Tasks []*ActionOomReportTaskIndexOutput `json:"tasks"`
	}

	// Action output without the namespace
	Output []*ActionOomReportTaskIndexOutput
}

// Prepare the action for invocation
func (action *ActionOomReportTaskIndex) Prepare() *ActionOomReportTaskIndexInvocation {
	return &ActionOomReportTaskIndexInvocation{
		Action: action,
		Path:   "/v6.0/oom_reports/{oom_report_id}/tasks",
	}
}

// ActionOomReportTaskIndexInvocation is used to configure action for invocation
type ActionOomReportTaskIndexInvocation struct {
	// Pointer to the action
	Action *ActionOomReportTaskIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOomReportTaskIndexInput
	// Global meta input parameters
	MetaInput *ActionOomReportTaskIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOomReportTaskIndexInvocation) SetPathParamInt(param string, value int64) *ActionOomReportTaskIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOomReportTaskIndexInvocation) SetPathParamString(param string, value string) *ActionOomReportTaskIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOomReportTaskIndexInvocation) NewInput() *ActionOomReportTaskIndexInput {
	inv.Input = &ActionOomReportTaskIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOomReportTaskIndexInvocation) SetInput(input *ActionOomReportTaskIndexInput) *ActionOomReportTaskIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOomReportTaskIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOomReportTaskIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOomReportTaskIndexInvocation) NewMetaInput() *ActionOomReportTaskIndexMetaGlobalInput {
	inv.MetaInput = &ActionOomReportTaskIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOomReportTaskIndexInvocation) SetMetaInput(input *ActionOomReportTaskIndexMetaGlobalInput) *ActionOomReportTaskIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOomReportTaskIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOomReportTaskIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOomReportTaskIndexInvocation) Call() (*ActionOomReportTaskIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOomReportTaskIndexInvocation) callAsQuery() (*ActionOomReportTaskIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOomReportTaskIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Tasks
	}
	return resp, err
}

func (inv *ActionOomReportTaskIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["task[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["task[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
	}
}

func (inv *ActionOomReportTaskIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
