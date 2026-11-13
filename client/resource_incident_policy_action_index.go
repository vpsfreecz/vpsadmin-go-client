package client

import ()

// ActionIncidentPolicyIndex is a type for action Incident_policy#Index
type ActionIncidentPolicyIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIncidentPolicyIndex(client *Client) *ActionIncidentPolicyIndex {
	return &ActionIncidentPolicyIndex{
		Client: client,
	}
}

// ActionIncidentPolicyIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIncidentPolicyIndexMetaGlobalInput struct {
	No bool "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIncidentPolicyIndexMetaGlobalInput) SetNo(value bool) *ActionIncidentPolicyIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIncidentPolicyIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncidentPolicyIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIncidentPolicyIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIncidentPolicyIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncidentPolicyIndexOutput is a type for action output parameters
type ActionIncidentPolicyIndexOutput struct {
	Description       string "json:\"description\""
	Key               string "json:\"key\""
	Label             string "json:\"label\""
	SettingsCategory  string "json:\"settings_category\""
	StopPeriodDays    int64  "json:\"stop_period_days\""
	WarningPeriodDays int64  "json:\"warning_period_days\""
}

// Type for action response, including envelope
type ActionIncidentPolicyIndexResponse struct {
	Action *ActionIncidentPolicyIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IncidentPolicies []*ActionIncidentPolicyIndexOutput "json:\"incident_policies\""
	}

	// Action output without the namespace
	Output []*ActionIncidentPolicyIndexOutput
}

// Call the action directly without any path or input parameters
func (action *ActionIncidentPolicyIndex) Call() (*ActionIncidentPolicyIndexResponse, error) {
	return action.Prepare().Call()
}

// Prepare the action for invocation
func (action *ActionIncidentPolicyIndex) Prepare() *ActionIncidentPolicyIndexInvocation {
	return &ActionIncidentPolicyIndexInvocation{
		Action: action,
		Path:   "/v7.0/incident_policies",
	}
}

// ActionIncidentPolicyIndexInvocation is used to configure action for invocation
type ActionIncidentPolicyIndexInvocation struct {
	// Pointer to the action
	Action *ActionIncidentPolicyIndex

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionIncidentPolicyIndexMetaGlobalInput
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIncidentPolicyIndexInvocation) NewMetaInput() *ActionIncidentPolicyIndexMetaGlobalInput {
	inv.MetaInput = &ActionIncidentPolicyIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIncidentPolicyIndexInvocation) SetMetaInput(input *ActionIncidentPolicyIndexMetaGlobalInput) *ActionIncidentPolicyIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIncidentPolicyIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIncidentPolicyIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIncidentPolicyIndexInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIncidentPolicyIndexInvocation) Call() (*ActionIncidentPolicyIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionIncidentPolicyIndexInvocation) callAsQuery() (*ActionIncidentPolicyIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIncidentPolicyIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IncidentPolicies
	}
	return resp, err
}

func (inv *ActionIncidentPolicyIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
