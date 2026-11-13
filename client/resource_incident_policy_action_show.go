package client

import (
	"net/url"
	"strings"
)

// ActionIncidentPolicyShow is a type for action Incident_policy#Show
type ActionIncidentPolicyShow struct {
	// Pointer to client
	Client *Client
}

func NewActionIncidentPolicyShow(client *Client) *ActionIncidentPolicyShow {
	return &ActionIncidentPolicyShow{
		Client: client,
	}
}

// ActionIncidentPolicyShowMetaGlobalInput is a type for action global meta input parameters
type ActionIncidentPolicyShowMetaGlobalInput struct {
	No bool "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIncidentPolicyShowMetaGlobalInput) SetNo(value bool) *ActionIncidentPolicyShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIncidentPolicyShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncidentPolicyShowMetaGlobalInput) SelectParameters(params ...string) *ActionIncidentPolicyShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIncidentPolicyShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncidentPolicyShowOutput is a type for action output parameters
type ActionIncidentPolicyShowOutput struct {
	Description       string "json:\"description\""
	Key               string "json:\"key\""
	Label             string "json:\"label\""
	SettingsCategory  string "json:\"settings_category\""
	StopPeriodDays    int64  "json:\"stop_period_days\""
	WarningPeriodDays int64  "json:\"warning_period_days\""
}

// Type for action response, including envelope
type ActionIncidentPolicyShowResponse struct {
	Action *ActionIncidentPolicyShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IncidentPolicy *ActionIncidentPolicyShowOutput "json:\"incident_policy\""
	}

	// Action output without the namespace
	Output *ActionIncidentPolicyShowOutput
}

// Prepare the action for invocation
func (action *ActionIncidentPolicyShow) Prepare() *ActionIncidentPolicyShowInvocation {
	return &ActionIncidentPolicyShowInvocation{
		Action: action,
		Path:   "/v7.0/incident_policies/{key}",
	}
}

// ActionIncidentPolicyShowInvocation is used to configure action for invocation
type ActionIncidentPolicyShowInvocation struct {
	// Pointer to the action
	Action *ActionIncidentPolicyShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionIncidentPolicyShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIncidentPolicyShowInvocation) SetPathParamInt(param string, value int64) *ActionIncidentPolicyShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIncidentPolicyShowInvocation) SetPathParamString(param string, value string) *ActionIncidentPolicyShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIncidentPolicyShowInvocation) NewMetaInput() *ActionIncidentPolicyShowMetaGlobalInput {
	inv.MetaInput = &ActionIncidentPolicyShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIncidentPolicyShowInvocation) SetMetaInput(input *ActionIncidentPolicyShowMetaGlobalInput) *ActionIncidentPolicyShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIncidentPolicyShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIncidentPolicyShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIncidentPolicyShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIncidentPolicyShowInvocation) Call() (*ActionIncidentPolicyShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionIncidentPolicyShowInvocation) callAsQuery() (*ActionIncidentPolicyShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIncidentPolicyShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IncidentPolicy
	}
	return resp, err
}

func (inv *ActionIncidentPolicyShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
