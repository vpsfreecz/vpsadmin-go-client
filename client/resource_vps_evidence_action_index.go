package client

import ()

// ActionVpsEvidenceIndex is a type for action Vps_evidence#Index
type ActionVpsEvidenceIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsEvidenceIndex(client *Client) *ActionVpsEvidenceIndex {
	return &ActionVpsEvidenceIndex{
		Client: client,
	}
}

// ActionVpsEvidenceIndexMetaGlobalInput is a type for action global meta input parameters
type ActionVpsEvidenceIndexMetaGlobalInput struct {
	No bool "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsEvidenceIndexMetaGlobalInput) SetNo(value bool) *ActionVpsEvidenceIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsEvidenceIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsEvidenceIndexMetaGlobalInput) SelectParameters(params ...string) *ActionVpsEvidenceIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsEvidenceIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsEvidenceIndexInput is a type for action input parameters
type ActionVpsEvidenceIndexInput struct {
	Environment int64  "json:\"environment\""
	Limit       int64  "json:\"limit\""
	Location    int64  "json:\"location\""
	Node        int64  "json:\"node\""
	Offset      int64  "json:\"offset\""
	Signal      string "json:\"signal\""
	User        int64  "json:\"user\""
	Vps         int64  "json:\"vps\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionVpsEvidenceIndexInput) SetEnvironment(value int64) *ActionVpsEvidenceIndexInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Environment"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionVpsEvidenceIndexInput) SetLimit(value int64) *ActionVpsEvidenceIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionVpsEvidenceIndexInput) SetLocation(value int64) *ActionVpsEvidenceIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionVpsEvidenceIndexInput) SetNode(value int64) *ActionVpsEvidenceIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionVpsEvidenceIndexInput) SetOffset(value int64) *ActionVpsEvidenceIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetSignal sets parameter Signal to value and selects it for sending
func (in *ActionVpsEvidenceIndexInput) SetSignal(value string) *ActionVpsEvidenceIndexInput {
	in.Signal = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Signal"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionVpsEvidenceIndexInput) SetUser(value int64) *ActionVpsEvidenceIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionVpsEvidenceIndexInput) SetVps(value int64) *ActionVpsEvidenceIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Vps"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsEvidenceIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsEvidenceIndexInput) SelectParameters(params ...string) *ActionVpsEvidenceIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsEvidenceIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsEvidenceIndexInput) UnselectParameters(params ...string) *ActionVpsEvidenceIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsEvidenceIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsEvidenceIndexOutput is a type for action output parameters
type ActionVpsEvidenceIndexOutput struct {
	ConditionMeasurement  string                "json:\"condition_measurement\""
	ConditionOperator     string                "json:\"condition_operator\""
	ConditionUnit         string                "json:\"condition_unit\""
	ConditionValue        int64                 "json:\"condition_value\""
	ConfigurationValid    bool                  "json:\"configuration_valid\""
	CreditUs              int64                 "json:\"credit_us\""
	Hostname              string                "json:\"hostname\""
	Id                    string                "json:\"id\""
	Label                 string                "json:\"label\""
	LastIncidentAt        string                "json:\"last_incident_at\""
	LastObservedAt        string                "json:\"last_observed_at\""
	Measurement           string                "json:\"measurement\""
	Node                  *ActionNodeShowOutput "json:\"node\""
	ObservationAgeSeconds float64               "json:\"observation_age_seconds\""
	ObservationState      string                "json:\"observation_state\""
	PeakIsLowerBound      bool                  "json:\"peak_is_lower_bound\""
	PeakUnit              string                "json:\"peak_unit\""
	PeakValue             int64                 "json:\"peak_value\""
	Qualification         string                "json:\"qualification\""
	RepeatSeconds         int64                 "json:\"repeat_seconds\""
	ReportDue             bool                  "json:\"report_due\""
	ReportEligible        bool                  "json:\"report_eligible\""
	ReportingEnabled      bool                  "json:\"reporting_enabled\""
	Representation        string                "json:\"representation\""
	Signal                string                "json:\"signal\""
	ThresholdSeconds      int64                 "json:\"threshold_seconds\""
	User                  *ActionUserShowOutput "json:\"user\""
	Vps                   *ActionVpsShowOutput  "json:\"vps\""
	WindowComplete        bool                  "json:\"window_complete\""
	WindowFrom            string                "json:\"window_from\""
	WindowSeconds         int64                 "json:\"window_seconds\""
	WindowTo              string                "json:\"window_to\""
}

// Type for action response, including envelope
type ActionVpsEvidenceIndexResponse struct {
	Action *ActionVpsEvidenceIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		VpsEvidence []*ActionVpsEvidenceIndexOutput "json:\"vps_evidence\""
	}

	// Action output without the namespace
	Output []*ActionVpsEvidenceIndexOutput
}

// Prepare the action for invocation
func (action *ActionVpsEvidenceIndex) Prepare() *ActionVpsEvidenceIndexInvocation {
	return &ActionVpsEvidenceIndexInvocation{
		Action: action,
		Path:   "/v7.0/vps_evidence",
	}
}

// ActionVpsEvidenceIndexInvocation is used to configure action for invocation
type ActionVpsEvidenceIndexInvocation struct {
	// Pointer to the action
	Action *ActionVpsEvidenceIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsEvidenceIndexInput
	// Global meta input parameters
	MetaInput *ActionVpsEvidenceIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsEvidenceIndexInvocation) NewInput() *ActionVpsEvidenceIndexInput {
	inv.Input = &ActionVpsEvidenceIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsEvidenceIndexInvocation) SetInput(input *ActionVpsEvidenceIndexInput) *ActionVpsEvidenceIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsEvidenceIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsEvidenceIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsEvidenceIndexInvocation) NewMetaInput() *ActionVpsEvidenceIndexMetaGlobalInput {
	inv.MetaInput = &ActionVpsEvidenceIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsEvidenceIndexInvocation) SetMetaInput(input *ActionVpsEvidenceIndexMetaGlobalInput) *ActionVpsEvidenceIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsEvidenceIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsEvidenceIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionVpsEvidenceIndexInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			if !inv.IsParameterNil("Environment") {
				if inv.Input.Environment < 0 {
					verr.Add("environment", t(inv.Action.Client, "validation.invalid_resource_id", nil))
				}
			}
		}
		if inv.IsParameterSelected("Location") {
			if !inv.IsParameterNil("Location") {
				if inv.Input.Location < 0 {
					verr.Add("location", t(inv.Action.Client, "validation.invalid_resource_id", nil))
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
func (inv *ActionVpsEvidenceIndexInvocation) Call() (*ActionVpsEvidenceIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionVpsEvidenceIndexInvocation) callAsQuery() (*ActionVpsEvidenceIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsEvidenceIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.VpsEvidence
	}
	return resp, err
}

func (inv *ActionVpsEvidenceIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			ret["vps_evidence[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("Limit") {
			ret["vps_evidence[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Location") {
			ret["vps_evidence[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("Node") {
			ret["vps_evidence[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("Offset") {
			ret["vps_evidence[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Signal") {
			ret["vps_evidence[signal]"] = inv.Input.Signal
		}
		if inv.IsParameterSelected("User") {
			ret["vps_evidence[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("Vps") {
			ret["vps_evidence[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
	}
}

func (inv *ActionVpsEvidenceIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
