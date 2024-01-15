package client

import (
	"strings"
)

// ActionOomReportShow is a type for action Oom_report#Show
type ActionOomReportShow struct {
	// Pointer to client
	Client *Client
}

func NewActionOomReportShow(client *Client) *ActionOomReportShow {
	return &ActionOomReportShow{
		Client: client,
	}
}

// ActionOomReportShowMetaGlobalInput is a type for action global meta input parameters
type ActionOomReportShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOomReportShowMetaGlobalInput) SetIncludes(value string) *ActionOomReportShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOomReportShowMetaGlobalInput) SetNo(value bool) *ActionOomReportShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOomReportShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportShowMetaGlobalInput) SelectParameters(params ...string) *ActionOomReportShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOomReportShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOomReportShowOutput is a type for action output parameters
type ActionOomReportShowOutput struct {
	Cgroup        string               `json:"cgroup"`
	Count         int64                `json:"count"`
	CreatedAt     string               `json:"created_at"`
	Id            int64                `json:"id"`
	InvokedByName string               `json:"invoked_by_name"`
	InvokedByPid  int64                `json:"invoked_by_pid"`
	KilledName    string               `json:"killed_name"`
	KilledPid     int64                `json:"killed_pid"`
	ReportedAt    string               `json:"reported_at"`
	Vps           *ActionVpsShowOutput `json:"vps"`
}

// Type for action response, including envelope
type ActionOomReportShowResponse struct {
	Action *ActionOomReportShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OomReport *ActionOomReportShowOutput `json:"oom_report"`
	}

	// Action output without the namespace
	Output *ActionOomReportShowOutput
}

// Prepare the action for invocation
func (action *ActionOomReportShow) Prepare() *ActionOomReportShowInvocation {
	return &ActionOomReportShowInvocation{
		Action: action,
		Path:   "/v6.0/oom_reports/{oom_report_id}",
	}
}

// ActionOomReportShowInvocation is used to configure action for invocation
type ActionOomReportShowInvocation struct {
	// Pointer to the action
	Action *ActionOomReportShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionOomReportShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOomReportShowInvocation) SetPathParamInt(param string, value int64) *ActionOomReportShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOomReportShowInvocation) SetPathParamString(param string, value string) *ActionOomReportShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOomReportShowInvocation) NewMetaInput() *ActionOomReportShowMetaGlobalInput {
	inv.MetaInput = &ActionOomReportShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOomReportShowInvocation) SetMetaInput(input *ActionOomReportShowMetaGlobalInput) *ActionOomReportShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOomReportShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOomReportShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOomReportShowInvocation) Call() (*ActionOomReportShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOomReportShowInvocation) callAsQuery() (*ActionOomReportShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOomReportShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OomReport
	}
	return resp, err
}

func (inv *ActionOomReportShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
