package client

import ()

// ActionIncidentPolicyEvaluationIndex is a type for action Incident_policy_evaluation#Index
type ActionIncidentPolicyEvaluationIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIncidentPolicyEvaluationIndex(client *Client) *ActionIncidentPolicyEvaluationIndex {
	return &ActionIncidentPolicyEvaluationIndex{
		Client: client,
	}
}

// ActionIncidentPolicyEvaluationIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIncidentPolicyEvaluationIndexMetaGlobalInput struct {
	No bool "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIncidentPolicyEvaluationIndexMetaGlobalInput) SetNo(value bool) *ActionIncidentPolicyEvaluationIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIncidentPolicyEvaluationIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncidentPolicyEvaluationIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIncidentPolicyEvaluationIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIncidentPolicyEvaluationIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncidentPolicyEvaluationIndexInput is a type for action input parameters
type ActionIncidentPolicyEvaluationIndexInput struct {
	ActionKind  string "json:\"action_kind\""
	ActionState string "json:\"action_state\""
	AsOf        string "json:\"as_of\""
	Decision    string "json:\"decision\""
	Direction   string "json:\"direction\""
	From        string "json:\"from\""
	Limit       int64  "json:\"limit\""
	MailState   string "json:\"mail_state\""
	Offset      int64  "json:\"offset\""
	PolicyKey   string "json:\"policy_key\""
	Scope       string "json:\"scope\""
	To          string "json:\"to\""
	UserId      int64  "json:\"user_id\""
	VpsId       int64  "json:\"vps_id\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetActionKind sets parameter ActionKind to value and selects it for sending
func (in *ActionIncidentPolicyEvaluationIndexInput) SetActionKind(value string) *ActionIncidentPolicyEvaluationIndexInput {
	in.ActionKind = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ActionKind"] = nil
	return in
}

// SetActionState sets parameter ActionState to value and selects it for sending
func (in *ActionIncidentPolicyEvaluationIndexInput) SetActionState(value string) *ActionIncidentPolicyEvaluationIndexInput {
	in.ActionState = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ActionState"] = nil
	return in
}

// SetAsOf sets parameter AsOf to value and selects it for sending
func (in *ActionIncidentPolicyEvaluationIndexInput) SetAsOf(value string) *ActionIncidentPolicyEvaluationIndexInput {
	in.AsOf = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AsOf"] = nil
	return in
}

// SetDecision sets parameter Decision to value and selects it for sending
func (in *ActionIncidentPolicyEvaluationIndexInput) SetDecision(value string) *ActionIncidentPolicyEvaluationIndexInput {
	in.Decision = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Decision"] = nil
	return in
}

// SetDirection sets parameter Direction to value and selects it for sending
func (in *ActionIncidentPolicyEvaluationIndexInput) SetDirection(value string) *ActionIncidentPolicyEvaluationIndexInput {
	in.Direction = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Direction"] = nil
	return in
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionIncidentPolicyEvaluationIndexInput) SetFrom(value string) *ActionIncidentPolicyEvaluationIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIncidentPolicyEvaluationIndexInput) SetLimit(value int64) *ActionIncidentPolicyEvaluationIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetMailState sets parameter MailState to value and selects it for sending
func (in *ActionIncidentPolicyEvaluationIndexInput) SetMailState(value string) *ActionIncidentPolicyEvaluationIndexInput {
	in.MailState = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MailState"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionIncidentPolicyEvaluationIndexInput) SetOffset(value int64) *ActionIncidentPolicyEvaluationIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetPolicyKey sets parameter PolicyKey to value and selects it for sending
func (in *ActionIncidentPolicyEvaluationIndexInput) SetPolicyKey(value string) *ActionIncidentPolicyEvaluationIndexInput {
	in.PolicyKey = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["PolicyKey"] = nil
	return in
}

// SetScope sets parameter Scope to value and selects it for sending
func (in *ActionIncidentPolicyEvaluationIndexInput) SetScope(value string) *ActionIncidentPolicyEvaluationIndexInput {
	in.Scope = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Scope"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionIncidentPolicyEvaluationIndexInput) SetTo(value string) *ActionIncidentPolicyEvaluationIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SetUserId sets parameter UserId to value and selects it for sending
func (in *ActionIncidentPolicyEvaluationIndexInput) SetUserId(value int64) *ActionIncidentPolicyEvaluationIndexInput {
	in.UserId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserId"] = nil
	return in
}

// SetVpsId sets parameter VpsId to value and selects it for sending
func (in *ActionIncidentPolicyEvaluationIndexInput) SetVpsId(value int64) *ActionIncidentPolicyEvaluationIndexInput {
	in.VpsId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["VpsId"] = nil
	return in
}

// SelectParameters sets parameters from ActionIncidentPolicyEvaluationIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncidentPolicyEvaluationIndexInput) SelectParameters(params ...string) *ActionIncidentPolicyEvaluationIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIncidentPolicyEvaluationIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIncidentPolicyEvaluationIndexInput) UnselectParameters(params ...string) *ActionIncidentPolicyEvaluationIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIncidentPolicyEvaluationIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncidentPolicyEvaluationIndexOutput is a type for action output parameters
type ActionIncidentPolicyEvaluationIndexOutput struct {
	ActionCreatedAt      string "json:\"action_created_at\""
	ActionId             int64  "json:\"action_id\""
	ActionKind           string "json:\"action_kind\""
	ActionState          string "json:\"action_state\""
	ActionUpdatedAt      string "json:\"action_updated_at\""
	Count30d             int64  "json:\"count_30d\""
	Count7d              int64  "json:\"count_7d\""
	Decision             string "json:\"decision\""
	DetectedAt           string "json:\"detected_at\""
	EffectiveMode        string "json:\"effective_mode\""
	EvaluatedAt          string "json:\"evaluated_at\""
	Hostname             string "json:\"hostname\""
	Id                   int64  "json:\"id\""
	IncidentReportId     int64  "json:\"incident_report_id\""
	Label                string "json:\"label\""
	MailFinishedAt       string "json:\"mail_finished_at\""
	MailQueuedAt         string "json:\"mail_queued_at\""
	MailState            string "json:\"mail_state\""
	Matched              bool   "json:\"matched\""
	NoticeReportId       int64  "json:\"notice_report_id\""
	PhysicalStopActionId int64  "json:\"physical_stop_action_id\""
	PolicyKey            string "json:\"policy_key\""
	State                string "json:\"state\""
	StopCompletedAt      string "json:\"stop_completed_at\""
	StopEnforced         bool   "json:\"stop_enforced\""
	StopLimit            int64  "json:\"stop_limit\""
	StopPeriodDays       int64  "json:\"stop_period_days\""
	ThresholdStatus      string "json:\"threshold_status\""
	TransactionChainId   int64  "json:\"transaction_chain_id\""
	UserId               int64  "json:\"user_id\""
	UserLogin            string "json:\"user_login\""
	VpsId                int64  "json:\"vps_id\""
	WarningLimit         int64  "json:\"warning_limit\""
	WarningPeriodDays    int64  "json:\"warning_period_days\""
}

// Type for action response, including envelope
type ActionIncidentPolicyEvaluationIndexResponse struct {
	Action *ActionIncidentPolicyEvaluationIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IncidentPolicyEvaluations []*ActionIncidentPolicyEvaluationIndexOutput "json:\"incident_policy_evaluations\""
	}

	// Action output without the namespace
	Output []*ActionIncidentPolicyEvaluationIndexOutput
}

// Prepare the action for invocation
func (action *ActionIncidentPolicyEvaluationIndex) Prepare() *ActionIncidentPolicyEvaluationIndexInvocation {
	return &ActionIncidentPolicyEvaluationIndexInvocation{
		Action: action,
		Path:   "/v7.0/incident_policy_evaluations",
	}
}

// ActionIncidentPolicyEvaluationIndexInvocation is used to configure action for invocation
type ActionIncidentPolicyEvaluationIndexInvocation struct {
	// Pointer to the action
	Action *ActionIncidentPolicyEvaluationIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIncidentPolicyEvaluationIndexInput
	// Global meta input parameters
	MetaInput *ActionIncidentPolicyEvaluationIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIncidentPolicyEvaluationIndexInvocation) NewInput() *ActionIncidentPolicyEvaluationIndexInput {
	inv.Input = &ActionIncidentPolicyEvaluationIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIncidentPolicyEvaluationIndexInvocation) SetInput(input *ActionIncidentPolicyEvaluationIndexInput) *ActionIncidentPolicyEvaluationIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIncidentPolicyEvaluationIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIncidentPolicyEvaluationIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIncidentPolicyEvaluationIndexInvocation) NewMetaInput() *ActionIncidentPolicyEvaluationIndexMetaGlobalInput {
	inv.MetaInput = &ActionIncidentPolicyEvaluationIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIncidentPolicyEvaluationIndexInvocation) SetMetaInput(input *ActionIncidentPolicyEvaluationIndexMetaGlobalInput) *ActionIncidentPolicyEvaluationIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIncidentPolicyEvaluationIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIncidentPolicyEvaluationIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIncidentPolicyEvaluationIndexInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.Input != nil {
		if inv.IsParameterSelected("AsOf") {
			if !inv.IsParameterNil("AsOf") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.AsOf)
				if !ok {
					verr.Add("as_of", t(inv.Action.Client, "validation.invalid_datetime", nil))
				} else {
					inv.Input.AsOf = normalized
				}
			}
		}
		if inv.IsParameterSelected("From") {
			if !inv.IsParameterNil("From") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.From)
				if !ok {
					verr.Add("from", t(inv.Action.Client, "validation.invalid_datetime", nil))
				} else {
					inv.Input.From = normalized
				}
			}
		}
		if inv.IsParameterSelected("To") {
			if !inv.IsParameterNil("To") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.To)
				if !ok {
					verr.Add("to", t(inv.Action.Client, "validation.invalid_datetime", nil))
				} else {
					inv.Input.To = normalized
				}
			}
		}
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIncidentPolicyEvaluationIndexInvocation) Call() (*ActionIncidentPolicyEvaluationIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionIncidentPolicyEvaluationIndexInvocation) callAsQuery() (*ActionIncidentPolicyEvaluationIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIncidentPolicyEvaluationIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IncidentPolicyEvaluations
	}
	return resp, err
}

func (inv *ActionIncidentPolicyEvaluationIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("ActionKind") {
			ret["incident_policy_evaluation[action_kind]"] = inv.Input.ActionKind
		}
		if inv.IsParameterSelected("ActionState") {
			ret["incident_policy_evaluation[action_state]"] = inv.Input.ActionState
		}
		if inv.IsParameterSelected("AsOf") {
			ret["incident_policy_evaluation[as_of]"] = inv.Input.AsOf
		}
		if inv.IsParameterSelected("Decision") {
			ret["incident_policy_evaluation[decision]"] = inv.Input.Decision
		}
		if inv.IsParameterSelected("Direction") {
			ret["incident_policy_evaluation[direction]"] = inv.Input.Direction
		}
		if inv.IsParameterSelected("From") {
			ret["incident_policy_evaluation[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("Limit") {
			ret["incident_policy_evaluation[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("MailState") {
			ret["incident_policy_evaluation[mail_state]"] = inv.Input.MailState
		}
		if inv.IsParameterSelected("Offset") {
			ret["incident_policy_evaluation[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("PolicyKey") {
			ret["incident_policy_evaluation[policy_key]"] = inv.Input.PolicyKey
		}
		if inv.IsParameterSelected("Scope") {
			ret["incident_policy_evaluation[scope]"] = inv.Input.Scope
		}
		if inv.IsParameterSelected("To") {
			ret["incident_policy_evaluation[to]"] = inv.Input.To
		}
		if inv.IsParameterSelected("UserId") {
			ret["incident_policy_evaluation[user_id]"] = convertInt64ToString(inv.Input.UserId)
		}
		if inv.IsParameterSelected("VpsId") {
			ret["incident_policy_evaluation[vps_id]"] = convertInt64ToString(inv.Input.VpsId)
		}
	}
}

func (inv *ActionIncidentPolicyEvaluationIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
