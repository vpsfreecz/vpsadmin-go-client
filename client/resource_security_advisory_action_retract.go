package client

import (
	"net/url"
	"strings"
)

// ActionSecurityAdvisoryRetract is a type for action Security_advisory#Retract
type ActionSecurityAdvisoryRetract struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryRetract(client *Client) *ActionSecurityAdvisoryRetract {
	return &ActionSecurityAdvisoryRetract{
		Client: client,
	}
}

// ActionSecurityAdvisoryRetractMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryRetractMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryRetractMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryRetractMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryRetractMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryRetractMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryRetractMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryRetractMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryRetractMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryRetractMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryRetractInput is a type for action input parameters
type ActionSecurityAdvisoryRetractInput struct {
	ExpectedContentRevision int64  "json:\"expected_content_revision\""
	ExpectedState           string "json:\"expected_state\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetExpectedContentRevision sets parameter ExpectedContentRevision to value and selects it for sending
func (in *ActionSecurityAdvisoryRetractInput) SetExpectedContentRevision(value int64) *ActionSecurityAdvisoryRetractInput {
	in.ExpectedContentRevision = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ExpectedContentRevision"] = nil
	return in
}

// SetExpectedState sets parameter ExpectedState to value and selects it for sending
func (in *ActionSecurityAdvisoryRetractInput) SetExpectedState(value string) *ActionSecurityAdvisoryRetractInput {
	in.ExpectedState = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ExpectedState"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryRetractInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryRetractInput) SelectParameters(params ...string) *ActionSecurityAdvisoryRetractInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionSecurityAdvisoryRetractInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryRetractInput) UnselectParameters(params ...string) *ActionSecurityAdvisoryRetractInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionSecurityAdvisoryRetractInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryRetractRequest is a type for the entire action request
type ActionSecurityAdvisoryRetractRequest struct {
	SecurityAdvisory map[string]interface{} "json:\"security_advisory\""
	Meta             map[string]interface{} "json:\"_meta\""
}

// ActionSecurityAdvisoryRetractOutput is a type for action output parameters
type ActionSecurityAdvisoryRetractOutput struct {
	Affected          bool                  "json:\"affected\""
	AffectedNodeCount int64                 "json:\"affected_node_count\""
	AffectedUserCount int64                 "json:\"affected_user_count\""
	AffectedVpsCount  int64                 "json:\"affected_vps_count\""
	ContentRevision   int64                 "json:\"content_revision\""
	CreatedAt         string                "json:\"created_at\""
	CreatedBy         *ActionUserShowOutput "json:\"created_by\""
	CsDescription     string                "json:\"cs_description\""
	CsResponse        string                "json:\"cs_response\""
	CsSummary         string                "json:\"cs_summary\""
	EnDescription     string                "json:\"en_description\""
	EnResponse        string                "json:\"en_response\""
	EnSummary         string                "json:\"en_summary\""
	ExternalId        string                "json:\"external_id\""
	Id                int64                 "json:\"id\""
	Name              string                "json:\"name\""
	PublishedAt       string                "json:\"published_at\""
	PublishedBy       *ActionUserShowOutput "json:\"published_by\""
	RetractedAt       string                "json:\"retracted_at\""
	State             string                "json:\"state\""
	UpdatedAt         string                "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionSecurityAdvisoryRetractResponse struct {
	Action *ActionSecurityAdvisoryRetract "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SecurityAdvisory *ActionSecurityAdvisoryRetractOutput "json:\"security_advisory\""
	}

	// Action output without the namespace
	Output *ActionSecurityAdvisoryRetractOutput
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryRetract) Prepare() *ActionSecurityAdvisoryRetractInvocation {
	return &ActionSecurityAdvisoryRetractInvocation{
		Action: action,
		Path:   "/v7.0/security_advisories/{security_advisory_id}/retract",
	}
}

// ActionSecurityAdvisoryRetractInvocation is used to configure action for invocation
type ActionSecurityAdvisoryRetractInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryRetract

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSecurityAdvisoryRetractInput
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryRetractMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSecurityAdvisoryRetractInvocation) SetPathParamInt(param string, value int64) *ActionSecurityAdvisoryRetractInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSecurityAdvisoryRetractInvocation) SetPathParamString(param string, value string) *ActionSecurityAdvisoryRetractInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSecurityAdvisoryRetractInvocation) NewInput() *ActionSecurityAdvisoryRetractInput {
	inv.Input = &ActionSecurityAdvisoryRetractInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSecurityAdvisoryRetractInvocation) SetInput(input *ActionSecurityAdvisoryRetractInput) *ActionSecurityAdvisoryRetractInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSecurityAdvisoryRetractInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryRetractInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryRetractInvocation) NewMetaInput() *ActionSecurityAdvisoryRetractMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryRetractMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryRetractInvocation) SetMetaInput(input *ActionSecurityAdvisoryRetractMetaGlobalInput) *ActionSecurityAdvisoryRetractInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryRetractInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryRetractInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryRetractInvocation) validate() error {
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
func (inv *ActionSecurityAdvisoryRetractInvocation) Call() (*ActionSecurityAdvisoryRetractResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionSecurityAdvisoryRetractInvocation) callAsBody() (*ActionSecurityAdvisoryRetractResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSecurityAdvisoryRetractResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SecurityAdvisory
	}
	return resp, err
}

func (inv *ActionSecurityAdvisoryRetractInvocation) makeAllInputParams() *ActionSecurityAdvisoryRetractRequest {
	return &ActionSecurityAdvisoryRetractRequest{
		SecurityAdvisory: inv.makeInputParams(),
		Meta:             inv.makeMetaInputParams(),
	}
}

func (inv *ActionSecurityAdvisoryRetractInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("ExpectedContentRevision") {
			ret["expected_content_revision"] = inv.Input.ExpectedContentRevision
		}
		if inv.IsParameterSelected("ExpectedState") {
			ret["expected_state"] = inv.Input.ExpectedState
		}
	}

	return ret
}

func (inv *ActionSecurityAdvisoryRetractInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
