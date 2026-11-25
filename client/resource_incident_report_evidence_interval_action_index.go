package client

import (
	"net/url"
	"strings"
)

// ActionIncidentReportEvidenceIntervalIndex is a type for action Incident_report.Evidence_interval#Index
type ActionIncidentReportEvidenceIntervalIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIncidentReportEvidenceIntervalIndex(client *Client) *ActionIncidentReportEvidenceIntervalIndex {
	return &ActionIncidentReportEvidenceIntervalIndex{
		Client: client,
	}
}

// ActionIncidentReportEvidenceIntervalIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIncidentReportEvidenceIntervalIndexMetaGlobalInput struct {
	No bool "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIncidentReportEvidenceIntervalIndexMetaGlobalInput) SetNo(value bool) *ActionIncidentReportEvidenceIntervalIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIncidentReportEvidenceIntervalIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncidentReportEvidenceIntervalIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIncidentReportEvidenceIntervalIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIncidentReportEvidenceIntervalIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncidentReportEvidenceIntervalIndexInput is a type for action input parameters
type ActionIncidentReportEvidenceIntervalIndexInput struct {
	Limit  int64 "json:\"limit\""
	Offset int64 "json:\"offset\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIncidentReportEvidenceIntervalIndexInput) SetLimit(value int64) *ActionIncidentReportEvidenceIntervalIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionIncidentReportEvidenceIntervalIndexInput) SetOffset(value int64) *ActionIncidentReportEvidenceIntervalIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SelectParameters sets parameters from ActionIncidentReportEvidenceIntervalIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncidentReportEvidenceIntervalIndexInput) SelectParameters(params ...string) *ActionIncidentReportEvidenceIntervalIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIncidentReportEvidenceIntervalIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIncidentReportEvidenceIntervalIndexInput) UnselectParameters(params ...string) *ActionIncidentReportEvidenceIntervalIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIncidentReportEvidenceIntervalIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncidentReportEvidenceIntervalIndexOutput is a type for action output parameters
type ActionIncidentReportEvidenceIntervalIndexOutput struct {
	CreditUs         int64  "json:\"credit_us\""
	FromAt           string "json:\"from_at\""
	Id               int64  "json:\"id\""
	Measurement      string "json:\"measurement\""
	PeakIsLowerBound bool   "json:\"peak_is_lower_bound\""
	PeakUnit         string "json:\"peak_unit\""
	PeakValue        int64  "json:\"peak_value\""
	ToAt             string "json:\"to_at\""
}

// Type for action response, including envelope
type ActionIncidentReportEvidenceIntervalIndexResponse struct {
	Action *ActionIncidentReportEvidenceIntervalIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		EvidenceIntervals []*ActionIncidentReportEvidenceIntervalIndexOutput "json:\"evidence_intervals\""
	}

	// Action output without the namespace
	Output []*ActionIncidentReportEvidenceIntervalIndexOutput
}

// Prepare the action for invocation
func (action *ActionIncidentReportEvidenceIntervalIndex) Prepare() *ActionIncidentReportEvidenceIntervalIndexInvocation {
	return &ActionIncidentReportEvidenceIntervalIndexInvocation{
		Action: action,
		Path:   "/v7.0/incident_reports/{incident_report_id}/evidence_intervals",
	}
}

// ActionIncidentReportEvidenceIntervalIndexInvocation is used to configure action for invocation
type ActionIncidentReportEvidenceIntervalIndexInvocation struct {
	// Pointer to the action
	Action *ActionIncidentReportEvidenceIntervalIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIncidentReportEvidenceIntervalIndexInput
	// Global meta input parameters
	MetaInput *ActionIncidentReportEvidenceIntervalIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIncidentReportEvidenceIntervalIndexInvocation) SetPathParamInt(param string, value int64) *ActionIncidentReportEvidenceIntervalIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIncidentReportEvidenceIntervalIndexInvocation) SetPathParamString(param string, value string) *ActionIncidentReportEvidenceIntervalIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIncidentReportEvidenceIntervalIndexInvocation) NewInput() *ActionIncidentReportEvidenceIntervalIndexInput {
	inv.Input = &ActionIncidentReportEvidenceIntervalIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIncidentReportEvidenceIntervalIndexInvocation) SetInput(input *ActionIncidentReportEvidenceIntervalIndexInput) *ActionIncidentReportEvidenceIntervalIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIncidentReportEvidenceIntervalIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIncidentReportEvidenceIntervalIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIncidentReportEvidenceIntervalIndexInvocation) NewMetaInput() *ActionIncidentReportEvidenceIntervalIndexMetaGlobalInput {
	inv.MetaInput = &ActionIncidentReportEvidenceIntervalIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIncidentReportEvidenceIntervalIndexInvocation) SetMetaInput(input *ActionIncidentReportEvidenceIntervalIndexMetaGlobalInput) *ActionIncidentReportEvidenceIntervalIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIncidentReportEvidenceIntervalIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIncidentReportEvidenceIntervalIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIncidentReportEvidenceIntervalIndexInvocation) validate() error {
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
func (inv *ActionIncidentReportEvidenceIntervalIndexInvocation) Call() (*ActionIncidentReportEvidenceIntervalIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionIncidentReportEvidenceIntervalIndexInvocation) callAsQuery() (*ActionIncidentReportEvidenceIntervalIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIncidentReportEvidenceIntervalIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.EvidenceIntervals
	}
	return resp, err
}

func (inv *ActionIncidentReportEvidenceIntervalIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["evidence_interval[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["evidence_interval[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
	}
}

func (inv *ActionIncidentReportEvidenceIntervalIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
