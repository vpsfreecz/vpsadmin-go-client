package client

import (
	"strings"
)

// ActionOomReportStatShow is a type for action Oom_report.Stat#Show
type ActionOomReportStatShow struct {
	// Pointer to client
	Client *Client
}

func NewActionOomReportStatShow(client *Client) *ActionOomReportStatShow {
	return &ActionOomReportStatShow{
		Client: client,
	}
}

// ActionOomReportStatShowMetaGlobalInput is a type for action global meta input parameters
type ActionOomReportStatShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOomReportStatShowMetaGlobalInput) SetIncludes(value string) *ActionOomReportStatShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOomReportStatShowMetaGlobalInput) SetNo(value bool) *ActionOomReportStatShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOomReportStatShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportStatShowMetaGlobalInput) SelectParameters(params ...string) *ActionOomReportStatShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOomReportStatShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOomReportStatShowOutput is a type for action output parameters
type ActionOomReportStatShowOutput struct {
	Id        int64  `json:"id"`
	Parameter string `json:"parameter"`
	Value     int64  `json:"value"`
}

// Type for action response, including envelope
type ActionOomReportStatShowResponse struct {
	Action *ActionOomReportStatShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Stats []*ActionOomReportStatShowOutput `json:"stats"`
	}

	// Action output without the namespace
	Output []*ActionOomReportStatShowOutput
}

// Prepare the action for invocation
func (action *ActionOomReportStatShow) Prepare() *ActionOomReportStatShowInvocation {
	return &ActionOomReportStatShowInvocation{
		Action: action,
		Path:   "/v6.0/oom_reports/{oom_report_id}/stats/{stat_id}",
	}
}

// ActionOomReportStatShowInvocation is used to configure action for invocation
type ActionOomReportStatShowInvocation struct {
	// Pointer to the action
	Action *ActionOomReportStatShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionOomReportStatShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOomReportStatShowInvocation) SetPathParamInt(param string, value int64) *ActionOomReportStatShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOomReportStatShowInvocation) SetPathParamString(param string, value string) *ActionOomReportStatShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOomReportStatShowInvocation) NewMetaInput() *ActionOomReportStatShowMetaGlobalInput {
	inv.MetaInput = &ActionOomReportStatShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOomReportStatShowInvocation) SetMetaInput(input *ActionOomReportStatShowMetaGlobalInput) *ActionOomReportStatShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOomReportStatShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOomReportStatShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOomReportStatShowInvocation) Call() (*ActionOomReportStatShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOomReportStatShowInvocation) callAsQuery() (*ActionOomReportStatShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOomReportStatShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Stats
	}
	return resp, err
}

func (inv *ActionOomReportStatShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
