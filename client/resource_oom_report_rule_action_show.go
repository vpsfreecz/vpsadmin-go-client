package client

import (
	"strings"
)

// ActionOomReportRuleShow is a type for action Oom_report_rule#Show
type ActionOomReportRuleShow struct {
	// Pointer to client
	Client *Client
}

func NewActionOomReportRuleShow(client *Client) *ActionOomReportRuleShow {
	return &ActionOomReportRuleShow{
		Client: client,
	}
}

// ActionOomReportRuleShowMetaGlobalInput is a type for action global meta input parameters
type ActionOomReportRuleShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOomReportRuleShowMetaGlobalInput) SetIncludes(value string) *ActionOomReportRuleShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOomReportRuleShowMetaGlobalInput) SetNo(value bool) *ActionOomReportRuleShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOomReportRuleShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportRuleShowMetaGlobalInput) SelectParameters(params ...string) *ActionOomReportRuleShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOomReportRuleShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOomReportRuleShowOutput is a type for action output parameters
type ActionOomReportRuleShowOutput struct {
	Action        string               `json:"action"`
	CgroupPattern string               `json:"cgroup_pattern"`
	CreatedAt     string               `json:"created_at"`
	HitCount      int64                `json:"hit_count"`
	Id            int64                `json:"id"`
	Label         string               `json:"label"`
	UpdatedAt     string               `json:"updated_at"`
	Vps           *ActionVpsShowOutput `json:"vps"`
}

// Type for action response, including envelope
type ActionOomReportRuleShowResponse struct {
	Action *ActionOomReportRuleShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OomReportRule *ActionOomReportRuleShowOutput `json:"oom_report_rule"`
	}

	// Action output without the namespace
	Output *ActionOomReportRuleShowOutput
}

// Prepare the action for invocation
func (action *ActionOomReportRuleShow) Prepare() *ActionOomReportRuleShowInvocation {
	return &ActionOomReportRuleShowInvocation{
		Action: action,
		Path:   "/v7.0/oom_report_rules/{oom_report_rule_id}",
	}
}

// ActionOomReportRuleShowInvocation is used to configure action for invocation
type ActionOomReportRuleShowInvocation struct {
	// Pointer to the action
	Action *ActionOomReportRuleShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionOomReportRuleShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOomReportRuleShowInvocation) SetPathParamInt(param string, value int64) *ActionOomReportRuleShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOomReportRuleShowInvocation) SetPathParamString(param string, value string) *ActionOomReportRuleShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOomReportRuleShowInvocation) NewMetaInput() *ActionOomReportRuleShowMetaGlobalInput {
	inv.MetaInput = &ActionOomReportRuleShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOomReportRuleShowInvocation) SetMetaInput(input *ActionOomReportRuleShowMetaGlobalInput) *ActionOomReportRuleShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOomReportRuleShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOomReportRuleShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOomReportRuleShowInvocation) Call() (*ActionOomReportRuleShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOomReportRuleShowInvocation) callAsQuery() (*ActionOomReportRuleShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOomReportRuleShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OomReportRule
	}
	return resp, err
}

func (inv *ActionOomReportRuleShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
