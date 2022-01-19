package client

import (
	"strings"
)

// ActionOomReportTaskShow is a type for action Oom_report.Task#Show
type ActionOomReportTaskShow struct {
	// Pointer to client
	Client *Client
}

func NewActionOomReportTaskShow(client *Client) *ActionOomReportTaskShow {
	return &ActionOomReportTaskShow{
		Client: client,
	}
}

// ActionOomReportTaskShowMetaGlobalInput is a type for action global meta input parameters
type ActionOomReportTaskShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOomReportTaskShowMetaGlobalInput) SetIncludes(value string) *ActionOomReportTaskShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionOomReportTaskShowMetaGlobalInput) SetNo(value bool) *ActionOomReportTaskShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOomReportTaskShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportTaskShowMetaGlobalInput) SelectParameters(params ...string) *ActionOomReportTaskShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOomReportTaskShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionOomReportTaskShowOutput is a type for action output parameters
type ActionOomReportTaskShowOutput struct {
	HostPid int64 `json:"host_pid"`
	Id int64 `json:"id"`
	Name string `json:"name"`
	OomScoreAdj int64 `json:"oom_score_adj"`
	PgtablesBytes int64 `json:"pgtables_bytes"`
	Rss int64 `json:"rss"`
	Swapents int64 `json:"swapents"`
	Tgid int64 `json:"tgid"`
	TotalVm int64 `json:"total_vm"`
	VpsPid int64 `json:"vps_pid"`
	VpsUid int64 `json:"vps_uid"`
}


// Type for action response, including envelope
type ActionOomReportTaskShowResponse struct {
	Action *ActionOomReportTaskShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Tasks []*ActionOomReportTaskShowOutput `json:"tasks"`
	}

	// Action output without the namespace
	Output []*ActionOomReportTaskShowOutput
}


// Prepare the action for invocation
func (action *ActionOomReportTaskShow) Prepare() *ActionOomReportTaskShowInvocation {
	return &ActionOomReportTaskShowInvocation{
		Action: action,
		Path: "/v6.0/oom_reports/{oom_report_id}/tasks/{task_id}",
	}
}

// ActionOomReportTaskShowInvocation is used to configure action for invocation
type ActionOomReportTaskShowInvocation struct {
	// Pointer to the action
	Action *ActionOomReportTaskShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionOomReportTaskShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOomReportTaskShowInvocation) SetPathParamInt(param string, value int64) *ActionOomReportTaskShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOomReportTaskShowInvocation) SetPathParamString(param string, value string) *ActionOomReportTaskShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOomReportTaskShowInvocation) NewMetaInput() *ActionOomReportTaskShowMetaGlobalInput {
	inv.MetaInput = &ActionOomReportTaskShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOomReportTaskShowInvocation) SetMetaInput(input *ActionOomReportTaskShowMetaGlobalInput) *ActionOomReportTaskShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOomReportTaskShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOomReportTaskShowInvocation) Call() (*ActionOomReportTaskShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOomReportTaskShowInvocation) callAsQuery() (*ActionOomReportTaskShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOomReportTaskShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Tasks
	}
	return resp, err
}




func (inv *ActionOomReportTaskShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}

