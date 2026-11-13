package client

import (
	"net/url"
	"strings"
)

// ActionIncidentReportPolicyEvaluationIndex is a type for action Incident_report.Policy_evaluation#Index
type ActionIncidentReportPolicyEvaluationIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIncidentReportPolicyEvaluationIndex(client *Client) *ActionIncidentReportPolicyEvaluationIndex {
	return &ActionIncidentReportPolicyEvaluationIndex{
		Client: client,
	}
}

// ActionIncidentReportPolicyEvaluationIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIncidentReportPolicyEvaluationIndexMetaGlobalInput struct {
	No bool "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIncidentReportPolicyEvaluationIndexMetaGlobalInput) SetNo(value bool) *ActionIncidentReportPolicyEvaluationIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIncidentReportPolicyEvaluationIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncidentReportPolicyEvaluationIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIncidentReportPolicyEvaluationIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIncidentReportPolicyEvaluationIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncidentReportPolicyEvaluationIndexOutput is a type for action output parameters
type ActionIncidentReportPolicyEvaluationIndexOutput struct {
	Count30d          int64  "json:\"count_30d\""
	Count7d           int64  "json:\"count_7d\""
	Decision          string "json:\"decision\""
	EffectiveMode     string "json:\"effective_mode\""
	ErrorClass        string "json:\"error_class\""
	ErrorStage        string "json:\"error_stage\""
	EvaluatedAt       string "json:\"evaluated_at\""
	Label             string "json:\"label\""
	Matched           bool   "json:\"matched\""
	PolicyKey         string "json:\"policy_key\""
	State             string "json:\"state\""
	StopEnforced      bool   "json:\"stop_enforced\""
	StopLimit         int64  "json:\"stop_limit\""
	StopPeriodDays    int64  "json:\"stop_period_days\""
	WarningLimit      int64  "json:\"warning_limit\""
	WarningPeriodDays int64  "json:\"warning_period_days\""
}

// Type for action response, including envelope
type ActionIncidentReportPolicyEvaluationIndexResponse struct {
	Action *ActionIncidentReportPolicyEvaluationIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		PolicyEvaluations []*ActionIncidentReportPolicyEvaluationIndexOutput "json:\"policy_evaluations\""
	}

	// Action output without the namespace
	Output []*ActionIncidentReportPolicyEvaluationIndexOutput
}

// Prepare the action for invocation
func (action *ActionIncidentReportPolicyEvaluationIndex) Prepare() *ActionIncidentReportPolicyEvaluationIndexInvocation {
	return &ActionIncidentReportPolicyEvaluationIndexInvocation{
		Action: action,
		Path:   "/v7.0/incident_reports/{incident_report_id}/policy_evaluations",
	}
}

// ActionIncidentReportPolicyEvaluationIndexInvocation is used to configure action for invocation
type ActionIncidentReportPolicyEvaluationIndexInvocation struct {
	// Pointer to the action
	Action *ActionIncidentReportPolicyEvaluationIndex

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionIncidentReportPolicyEvaluationIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIncidentReportPolicyEvaluationIndexInvocation) SetPathParamInt(param string, value int64) *ActionIncidentReportPolicyEvaluationIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIncidentReportPolicyEvaluationIndexInvocation) SetPathParamString(param string, value string) *ActionIncidentReportPolicyEvaluationIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIncidentReportPolicyEvaluationIndexInvocation) NewMetaInput() *ActionIncidentReportPolicyEvaluationIndexMetaGlobalInput {
	inv.MetaInput = &ActionIncidentReportPolicyEvaluationIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIncidentReportPolicyEvaluationIndexInvocation) SetMetaInput(input *ActionIncidentReportPolicyEvaluationIndexMetaGlobalInput) *ActionIncidentReportPolicyEvaluationIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIncidentReportPolicyEvaluationIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIncidentReportPolicyEvaluationIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIncidentReportPolicyEvaluationIndexInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIncidentReportPolicyEvaluationIndexInvocation) Call() (*ActionIncidentReportPolicyEvaluationIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionIncidentReportPolicyEvaluationIndexInvocation) callAsQuery() (*ActionIncidentReportPolicyEvaluationIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIncidentReportPolicyEvaluationIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.PolicyEvaluations
	}
	return resp, err
}

func (inv *ActionIncidentReportPolicyEvaluationIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
