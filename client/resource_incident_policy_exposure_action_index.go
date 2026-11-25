package client

import ()

// ActionIncidentPolicyExposureIndex is a type for action Incident_policy_exposure#Index
type ActionIncidentPolicyExposureIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIncidentPolicyExposureIndex(client *Client) *ActionIncidentPolicyExposureIndex {
	return &ActionIncidentPolicyExposureIndex{
		Client: client,
	}
}

// ActionIncidentPolicyExposureIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIncidentPolicyExposureIndexMetaGlobalInput struct {
	No bool "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIncidentPolicyExposureIndexMetaGlobalInput) SetNo(value bool) *ActionIncidentPolicyExposureIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIncidentPolicyExposureIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncidentPolicyExposureIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIncidentPolicyExposureIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIncidentPolicyExposureIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncidentPolicyExposureIndexInput is a type for action input parameters
type ActionIncidentPolicyExposureIndexInput struct {
	AsOf          string "json:\"as_of\""
	Direction     string "json:\"direction\""
	EffectiveMode string "json:\"effective_mode\""
	Limit         int64  "json:\"limit\""
	Node          int64  "json:\"node\""
	Offset        int64  "json:\"offset\""
	PolicyKey     string "json:\"policy_key\""
	Sort          string "json:\"sort\""
	Status        string "json:\"status\""
	StopEnforced  bool   "json:\"stop_enforced\""
	User          int64  "json:\"user\""
	Vps           int64  "json:\"vps\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAsOf sets parameter AsOf to value and selects it for sending
func (in *ActionIncidentPolicyExposureIndexInput) SetAsOf(value string) *ActionIncidentPolicyExposureIndexInput {
	in.AsOf = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AsOf"] = nil
	return in
}

// SetDirection sets parameter Direction to value and selects it for sending
func (in *ActionIncidentPolicyExposureIndexInput) SetDirection(value string) *ActionIncidentPolicyExposureIndexInput {
	in.Direction = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Direction"] = nil
	return in
}

// SetEffectiveMode sets parameter EffectiveMode to value and selects it for sending
func (in *ActionIncidentPolicyExposureIndexInput) SetEffectiveMode(value string) *ActionIncidentPolicyExposureIndexInput {
	in.EffectiveMode = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EffectiveMode"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIncidentPolicyExposureIndexInput) SetLimit(value int64) *ActionIncidentPolicyExposureIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionIncidentPolicyExposureIndexInput) SetNode(value int64) *ActionIncidentPolicyExposureIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionIncidentPolicyExposureIndexInput) SetOffset(value int64) *ActionIncidentPolicyExposureIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetPolicyKey sets parameter PolicyKey to value and selects it for sending
func (in *ActionIncidentPolicyExposureIndexInput) SetPolicyKey(value string) *ActionIncidentPolicyExposureIndexInput {
	in.PolicyKey = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["PolicyKey"] = nil
	return in
}

// SetSort sets parameter Sort to value and selects it for sending
func (in *ActionIncidentPolicyExposureIndexInput) SetSort(value string) *ActionIncidentPolicyExposureIndexInput {
	in.Sort = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Sort"] = nil
	return in
}

// SetStatus sets parameter Status to value and selects it for sending
func (in *ActionIncidentPolicyExposureIndexInput) SetStatus(value string) *ActionIncidentPolicyExposureIndexInput {
	in.Status = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Status"] = nil
	return in
}

// SetStopEnforced sets parameter StopEnforced to value and selects it for sending
func (in *ActionIncidentPolicyExposureIndexInput) SetStopEnforced(value bool) *ActionIncidentPolicyExposureIndexInput {
	in.StopEnforced = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["StopEnforced"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionIncidentPolicyExposureIndexInput) SetUser(value int64) *ActionIncidentPolicyExposureIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionIncidentPolicyExposureIndexInput) SetVps(value int64) *ActionIncidentPolicyExposureIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Vps"] = nil
	return in
}

// SelectParameters sets parameters from ActionIncidentPolicyExposureIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncidentPolicyExposureIndexInput) SelectParameters(params ...string) *ActionIncidentPolicyExposureIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIncidentPolicyExposureIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIncidentPolicyExposureIndexInput) UnselectParameters(params ...string) *ActionIncidentPolicyExposureIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIncidentPolicyExposureIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncidentPolicyExposureIndexOutput is a type for action output parameters
type ActionIncidentPolicyExposureIndexOutput struct {
	AccountingSince         string                "json:\"accounting_since\""
	ConfigurationValid      bool                  "json:\"configuration_valid\""
	Count30d                int64                 "json:\"count_30d\""
	Count7d                 int64                 "json:\"count_7d\""
	EffectiveMode           string                "json:\"effective_mode\""
	GlobalMode              string                "json:\"global_mode\""
	Id                      string                "json:\"id\""
	Label                   string                "json:\"label\""
	LastIncidentAt          string                "json:\"last_incident_at\""
	LastStopAt              string                "json:\"last_stop_at\""
	LastWarningAt           string                "json:\"last_warning_at\""
	LatestActionChainId     int64                 "json:\"latest_action_chain_id\""
	LatestActionId          int64                 "json:\"latest_action_id\""
	LatestActionKind        string                "json:\"latest_action_kind\""
	LatestActionState       string                "json:\"latest_action_state\""
	LatestMailState         string                "json:\"latest_mail_state\""
	LatestNoticeReportId    int64                 "json:\"latest_notice_report_id\""
	LatestStopCompletedAt   string                "json:\"latest_stop_completed_at\""
	NextStopAllowedAt       string                "json:\"next_stop_allowed_at\""
	NextWarningAllowedAt    string                "json:\"next_warning_allowed_at\""
	Node                    *ActionNodeShowOutput "json:\"node\""
	OutstandingStopActionId int64                 "json:\"outstanding_stop_action_id\""
	OutstandingStopChainId  int64                 "json:\"outstanding_stop_chain_id\""
	OutstandingStopState    string                "json:\"outstanding_stop_state\""
	OwnerActive             bool                  "json:\"owner_active\""
	PolicyKey               string                "json:\"policy_key\""
	Status                  string                "json:\"status\""
	StopEnforced            bool                  "json:\"stop_enforced\""
	StopLimit               int64                 "json:\"stop_limit\""
	StopPeriodDays          int64                 "json:\"stop_period_days\""
	StopRemaining           int64                 "json:\"stop_remaining\""
	TransferPending         bool                  "json:\"transfer_pending\""
	User                    *ActionUserShowOutput "json:\"user\""
	Vps                     *ActionVpsShowOutput  "json:\"vps\""
	WarningLimit            int64                 "json:\"warning_limit\""
	WarningPeriodDays       int64                 "json:\"warning_period_days\""
	WarningPermitted        bool                  "json:\"warning_permitted\""
	WarningRemaining        int64                 "json:\"warning_remaining\""
}

// Type for action response, including envelope
type ActionIncidentPolicyExposureIndexResponse struct {
	Action *ActionIncidentPolicyExposureIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IncidentPolicyExposures []*ActionIncidentPolicyExposureIndexOutput "json:\"incident_policy_exposures\""
	}

	// Action output without the namespace
	Output []*ActionIncidentPolicyExposureIndexOutput
}

// Prepare the action for invocation
func (action *ActionIncidentPolicyExposureIndex) Prepare() *ActionIncidentPolicyExposureIndexInvocation {
	return &ActionIncidentPolicyExposureIndexInvocation{
		Action: action,
		Path:   "/v7.0/incident_policy_exposures",
	}
}

// ActionIncidentPolicyExposureIndexInvocation is used to configure action for invocation
type ActionIncidentPolicyExposureIndexInvocation struct {
	// Pointer to the action
	Action *ActionIncidentPolicyExposureIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIncidentPolicyExposureIndexInput
	// Global meta input parameters
	MetaInput *ActionIncidentPolicyExposureIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIncidentPolicyExposureIndexInvocation) NewInput() *ActionIncidentPolicyExposureIndexInput {
	inv.Input = &ActionIncidentPolicyExposureIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIncidentPolicyExposureIndexInvocation) SetInput(input *ActionIncidentPolicyExposureIndexInput) *ActionIncidentPolicyExposureIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIncidentPolicyExposureIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIncidentPolicyExposureIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIncidentPolicyExposureIndexInvocation) NewMetaInput() *ActionIncidentPolicyExposureIndexMetaGlobalInput {
	inv.MetaInput = &ActionIncidentPolicyExposureIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIncidentPolicyExposureIndexInvocation) SetMetaInput(input *ActionIncidentPolicyExposureIndexMetaGlobalInput) *ActionIncidentPolicyExposureIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIncidentPolicyExposureIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIncidentPolicyExposureIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIncidentPolicyExposureIndexInvocation) validate() error {
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
		if inv.IsParameterSelected("Node") {
			if !inv.IsParameterNil("Node") {
				if inv.Input.Node < 0 {
					verr.Add("node", t(inv.Action.Client, "validation.invalid_resource_id", nil))
				}
			}
		}
		if inv.IsParameterSelected("User") {
			if !inv.IsParameterNil("User") {
				if inv.Input.User < 0 {
					verr.Add("user", t(inv.Action.Client, "validation.invalid_resource_id", nil))
				}
			}
		}
		if inv.IsParameterSelected("Vps") {
			if !inv.IsParameterNil("Vps") {
				if inv.Input.Vps < 0 {
					verr.Add("vps", t(inv.Action.Client, "validation.invalid_resource_id", nil))
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
func (inv *ActionIncidentPolicyExposureIndexInvocation) Call() (*ActionIncidentPolicyExposureIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionIncidentPolicyExposureIndexInvocation) callAsQuery() (*ActionIncidentPolicyExposureIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIncidentPolicyExposureIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IncidentPolicyExposures
	}
	return resp, err
}

func (inv *ActionIncidentPolicyExposureIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("AsOf") {
			ret["incident_policy_exposure[as_of]"] = inv.Input.AsOf
		}
		if inv.IsParameterSelected("Direction") {
			ret["incident_policy_exposure[direction]"] = inv.Input.Direction
		}
		if inv.IsParameterSelected("EffectiveMode") {
			ret["incident_policy_exposure[effective_mode]"] = inv.Input.EffectiveMode
		}
		if inv.IsParameterSelected("Limit") {
			ret["incident_policy_exposure[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Node") {
			ret["incident_policy_exposure[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("Offset") {
			ret["incident_policy_exposure[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("PolicyKey") {
			ret["incident_policy_exposure[policy_key]"] = inv.Input.PolicyKey
		}
		if inv.IsParameterSelected("Sort") {
			ret["incident_policy_exposure[sort]"] = inv.Input.Sort
		}
		if inv.IsParameterSelected("Status") {
			ret["incident_policy_exposure[status]"] = inv.Input.Status
		}
		if inv.IsParameterSelected("StopEnforced") {
			ret["incident_policy_exposure[stop_enforced]"] = convertBoolToString(inv.Input.StopEnforced)
		}
		if inv.IsParameterSelected("User") {
			ret["incident_policy_exposure[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("Vps") {
			ret["incident_policy_exposure[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
	}
}

func (inv *ActionIncidentPolicyExposureIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
