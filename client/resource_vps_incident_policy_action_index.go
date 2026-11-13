package client

import (
	"net/url"
	"strings"
)

// ActionVpsIncidentPolicyIndex is a type for action Vps.Incident_policy#Index
type ActionVpsIncidentPolicyIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsIncidentPolicyIndex(client *Client) *ActionVpsIncidentPolicyIndex {
	return &ActionVpsIncidentPolicyIndex{
		Client: client,
	}
}

// ActionVpsIncidentPolicyIndexMetaGlobalInput is a type for action global meta input parameters
type ActionVpsIncidentPolicyIndexMetaGlobalInput struct {
	No bool "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsIncidentPolicyIndexMetaGlobalInput) SetNo(value bool) *ActionVpsIncidentPolicyIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsIncidentPolicyIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsIncidentPolicyIndexMetaGlobalInput) SelectParameters(params ...string) *ActionVpsIncidentPolicyIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsIncidentPolicyIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsIncidentPolicyIndexOutput is a type for action output parameters
type ActionVpsIncidentPolicyIndexOutput struct {
	ConfigurationValid   bool   "json:\"configuration_valid\""
	Count30d             int64  "json:\"count_30d\""
	Count7d              int64  "json:\"count_7d\""
	Description          string "json:\"description\""
	EffectiveMode        string "json:\"effective_mode\""
	EvaluatedAt          string "json:\"evaluated_at\""
	EvaluationReportId   int64  "json:\"evaluation_report_id\""
	EvaluationState      string "json:\"evaluation_state\""
	GlobalMode           string "json:\"global_mode\""
	Key                  string "json:\"key\""
	Label                string "json:\"label\""
	LastIncidentAt       string "json:\"last_incident_at\""
	LastWarningAt        string "json:\"last_warning_at\""
	SettingsCategory     string "json:\"settings_category\""
	Status               string "json:\"status\""
	StopEnforced         bool   "json:\"stop_enforced\""
	StopLimit            int64  "json:\"stop_limit\""
	StopLimitOverride    int64  "json:\"stop_limit_override\""
	StopPeriodDays       int64  "json:\"stop_period_days\""
	WarningLimit         int64  "json:\"warning_limit\""
	WarningLimitOverride int64  "json:\"warning_limit_override\""
	WarningPeriodDays    int64  "json:\"warning_period_days\""
}

// Type for action response, including envelope
type ActionVpsIncidentPolicyIndexResponse struct {
	Action *ActionVpsIncidentPolicyIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IncidentPolicies []*ActionVpsIncidentPolicyIndexOutput "json:\"incident_policies\""
	}

	// Action output without the namespace
	Output []*ActionVpsIncidentPolicyIndexOutput
}

// Prepare the action for invocation
func (action *ActionVpsIncidentPolicyIndex) Prepare() *ActionVpsIncidentPolicyIndexInvocation {
	return &ActionVpsIncidentPolicyIndexInvocation{
		Action: action,
		Path:   "/v7.0/vpses/{vps_id}/incident_policies",
	}
}

// ActionVpsIncidentPolicyIndexInvocation is used to configure action for invocation
type ActionVpsIncidentPolicyIndexInvocation struct {
	// Pointer to the action
	Action *ActionVpsIncidentPolicyIndex

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionVpsIncidentPolicyIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsIncidentPolicyIndexInvocation) SetPathParamInt(param string, value int64) *ActionVpsIncidentPolicyIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsIncidentPolicyIndexInvocation) SetPathParamString(param string, value string) *ActionVpsIncidentPolicyIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsIncidentPolicyIndexInvocation) NewMetaInput() *ActionVpsIncidentPolicyIndexMetaGlobalInput {
	inv.MetaInput = &ActionVpsIncidentPolicyIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsIncidentPolicyIndexInvocation) SetMetaInput(input *ActionVpsIncidentPolicyIndexMetaGlobalInput) *ActionVpsIncidentPolicyIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsIncidentPolicyIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsIncidentPolicyIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionVpsIncidentPolicyIndexInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsIncidentPolicyIndexInvocation) Call() (*ActionVpsIncidentPolicyIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionVpsIncidentPolicyIndexInvocation) callAsQuery() (*ActionVpsIncidentPolicyIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsIncidentPolicyIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IncidentPolicies
	}
	return resp, err
}

func (inv *ActionVpsIncidentPolicyIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
