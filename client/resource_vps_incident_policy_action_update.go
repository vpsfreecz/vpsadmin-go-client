package client

import (
	"net/url"
	"strings"
)

// ActionVpsIncidentPolicyUpdate is a type for action Vps.Incident_policy#Update
type ActionVpsIncidentPolicyUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsIncidentPolicyUpdate(client *Client) *ActionVpsIncidentPolicyUpdate {
	return &ActionVpsIncidentPolicyUpdate{
		Client: client,
	}
}

// ActionVpsIncidentPolicyUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionVpsIncidentPolicyUpdateMetaGlobalInput struct {
	No bool "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsIncidentPolicyUpdateMetaGlobalInput) SetNo(value bool) *ActionVpsIncidentPolicyUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsIncidentPolicyUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsIncidentPolicyUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionVpsIncidentPolicyUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsIncidentPolicyUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsIncidentPolicyUpdateInput is a type for action input parameters
type ActionVpsIncidentPolicyUpdateInput struct {
	StopLimitOverride    int64 "json:\"stop_limit_override\""
	WarningLimitOverride int64 "json:\"warning_limit_override\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetStopLimitOverride sets parameter StopLimitOverride to value and selects it for sending
func (in *ActionVpsIncidentPolicyUpdateInput) SetStopLimitOverride(value int64) *ActionVpsIncidentPolicyUpdateInput {
	in.StopLimitOverride = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetStopLimitOverrideNil(false)
	in._selectedParameters["StopLimitOverride"] = nil
	return in
}

// SetStopLimitOverrideNil sets parameter StopLimitOverride to nil and selects it for sending
func (in *ActionVpsIncidentPolicyUpdateInput) SetStopLimitOverrideNil(set bool) *ActionVpsIncidentPolicyUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["StopLimitOverride"] = nil
		in.SelectParameters("StopLimitOverride")
	} else {
		delete(in._nilParameters, "StopLimitOverride")
	}
	return in
}

// SetWarningLimitOverride sets parameter WarningLimitOverride to value and selects it for sending
func (in *ActionVpsIncidentPolicyUpdateInput) SetWarningLimitOverride(value int64) *ActionVpsIncidentPolicyUpdateInput {
	in.WarningLimitOverride = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetWarningLimitOverrideNil(false)
	in._selectedParameters["WarningLimitOverride"] = nil
	return in
}

// SetWarningLimitOverrideNil sets parameter WarningLimitOverride to nil and selects it for sending
func (in *ActionVpsIncidentPolicyUpdateInput) SetWarningLimitOverrideNil(set bool) *ActionVpsIncidentPolicyUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["WarningLimitOverride"] = nil
		in.SelectParameters("WarningLimitOverride")
	} else {
		delete(in._nilParameters, "WarningLimitOverride")
	}
	return in
}

// SelectParameters sets parameters from ActionVpsIncidentPolicyUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsIncidentPolicyUpdateInput) SelectParameters(params ...string) *ActionVpsIncidentPolicyUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsIncidentPolicyUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsIncidentPolicyUpdateInput) UnselectParameters(params ...string) *ActionVpsIncidentPolicyUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsIncidentPolicyUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsIncidentPolicyUpdateRequest is a type for the entire action request
type ActionVpsIncidentPolicyUpdateRequest struct {
	IncidentPolicy map[string]interface{} "json:\"incident_policy\""
	Meta           map[string]interface{} "json:\"_meta\""
}

// ActionVpsIncidentPolicyUpdateOutput is a type for action output parameters
type ActionVpsIncidentPolicyUpdateOutput struct {
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
type ActionVpsIncidentPolicyUpdateResponse struct {
	Action *ActionVpsIncidentPolicyUpdate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IncidentPolicy *ActionVpsIncidentPolicyUpdateOutput "json:\"incident_policy\""
	}

	// Action output without the namespace
	Output *ActionVpsIncidentPolicyUpdateOutput
}

// Prepare the action for invocation
func (action *ActionVpsIncidentPolicyUpdate) Prepare() *ActionVpsIncidentPolicyUpdateInvocation {
	return &ActionVpsIncidentPolicyUpdateInvocation{
		Action: action,
		Path:   "/v7.0/vpses/{vps_id}/incident_policies/{key}",
	}
}

// ActionVpsIncidentPolicyUpdateInvocation is used to configure action for invocation
type ActionVpsIncidentPolicyUpdateInvocation struct {
	// Pointer to the action
	Action *ActionVpsIncidentPolicyUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsIncidentPolicyUpdateInput
	// Global meta input parameters
	MetaInput *ActionVpsIncidentPolicyUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsIncidentPolicyUpdateInvocation) SetPathParamInt(param string, value int64) *ActionVpsIncidentPolicyUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsIncidentPolicyUpdateInvocation) SetPathParamString(param string, value string) *ActionVpsIncidentPolicyUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsIncidentPolicyUpdateInvocation) NewInput() *ActionVpsIncidentPolicyUpdateInput {
	inv.Input = &ActionVpsIncidentPolicyUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsIncidentPolicyUpdateInvocation) SetInput(input *ActionVpsIncidentPolicyUpdateInput) *ActionVpsIncidentPolicyUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsIncidentPolicyUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsIncidentPolicyUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsIncidentPolicyUpdateInvocation) NewMetaInput() *ActionVpsIncidentPolicyUpdateMetaGlobalInput {
	inv.MetaInput = &ActionVpsIncidentPolicyUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsIncidentPolicyUpdateInvocation) SetMetaInput(input *ActionVpsIncidentPolicyUpdateMetaGlobalInput) *ActionVpsIncidentPolicyUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsIncidentPolicyUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsIncidentPolicyUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionVpsIncidentPolicyUpdateInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.Input != nil {
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsIncidentPolicyUpdateInvocation) Call() (*ActionVpsIncidentPolicyUpdateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionVpsIncidentPolicyUpdateInvocation) callAsBody() (*ActionVpsIncidentPolicyUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsIncidentPolicyUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IncidentPolicy
	}
	return resp, err
}

func (inv *ActionVpsIncidentPolicyUpdateInvocation) makeAllInputParams() *ActionVpsIncidentPolicyUpdateRequest {
	return &ActionVpsIncidentPolicyUpdateRequest{
		IncidentPolicy: inv.makeInputParams(),
		Meta:           inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsIncidentPolicyUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("StopLimitOverride") {
			if inv.IsParameterNil("StopLimitOverride") {
				ret["stop_limit_override"] = nil
			} else {
				ret["stop_limit_override"] = inv.Input.StopLimitOverride
			}
		}
		if inv.IsParameterSelected("WarningLimitOverride") {
			if inv.IsParameterNil("WarningLimitOverride") {
				ret["warning_limit_override"] = nil
			} else {
				ret["warning_limit_override"] = inv.Input.WarningLimitOverride
			}
		}
	}

	return ret
}

func (inv *ActionVpsIncidentPolicyUpdateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
