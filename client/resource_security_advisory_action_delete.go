package client

import (
	"net/url"
	"strings"
)

// ActionSecurityAdvisoryDelete is a type for action Security_advisory#Delete
type ActionSecurityAdvisoryDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryDelete(client *Client) *ActionSecurityAdvisoryDelete {
	return &ActionSecurityAdvisoryDelete{
		Client: client,
	}
}

// ActionSecurityAdvisoryDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryDeleteMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryDeleteMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryDeleteMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryDeleteInput is a type for action input parameters
type ActionSecurityAdvisoryDeleteInput struct {
	ExpectedContentRevision int64  "json:\"expected_content_revision\""
	ExpectedState           string "json:\"expected_state\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetExpectedContentRevision sets parameter ExpectedContentRevision to value and selects it for sending
func (in *ActionSecurityAdvisoryDeleteInput) SetExpectedContentRevision(value int64) *ActionSecurityAdvisoryDeleteInput {
	in.ExpectedContentRevision = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ExpectedContentRevision"] = nil
	return in
}

// SetExpectedState sets parameter ExpectedState to value and selects it for sending
func (in *ActionSecurityAdvisoryDeleteInput) SetExpectedState(value string) *ActionSecurityAdvisoryDeleteInput {
	in.ExpectedState = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ExpectedState"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryDeleteInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryDeleteInput) SelectParameters(params ...string) *ActionSecurityAdvisoryDeleteInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionSecurityAdvisoryDeleteInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryDeleteInput) UnselectParameters(params ...string) *ActionSecurityAdvisoryDeleteInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionSecurityAdvisoryDeleteInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryDeleteRequest is a type for the entire action request
type ActionSecurityAdvisoryDeleteRequest struct {
	SecurityAdvisory map[string]interface{} "json:\"security_advisory\""
	Meta             map[string]interface{} "json:\"_meta\""
}

// Type for action response, including envelope
type ActionSecurityAdvisoryDeleteResponse struct {
	Action *ActionSecurityAdvisoryDelete "json:\"-\""
	*Envelope
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryDelete) Prepare() *ActionSecurityAdvisoryDeleteInvocation {
	return &ActionSecurityAdvisoryDeleteInvocation{
		Action: action,
		Path:   "/v7.0/security_advisories/{security_advisory_id}",
	}
}

// ActionSecurityAdvisoryDeleteInvocation is used to configure action for invocation
type ActionSecurityAdvisoryDeleteInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryDelete

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSecurityAdvisoryDeleteInput
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSecurityAdvisoryDeleteInvocation) SetPathParamInt(param string, value int64) *ActionSecurityAdvisoryDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSecurityAdvisoryDeleteInvocation) SetPathParamString(param string, value string) *ActionSecurityAdvisoryDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSecurityAdvisoryDeleteInvocation) NewInput() *ActionSecurityAdvisoryDeleteInput {
	inv.Input = &ActionSecurityAdvisoryDeleteInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSecurityAdvisoryDeleteInvocation) SetInput(input *ActionSecurityAdvisoryDeleteInput) *ActionSecurityAdvisoryDeleteInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSecurityAdvisoryDeleteInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryDeleteInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryDeleteInvocation) NewMetaInput() *ActionSecurityAdvisoryDeleteMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryDeleteInvocation) SetMetaInput(input *ActionSecurityAdvisoryDeleteMetaGlobalInput) *ActionSecurityAdvisoryDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryDeleteInvocation) validate() error {
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
func (inv *ActionSecurityAdvisoryDeleteInvocation) Call() (*ActionSecurityAdvisoryDeleteResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionSecurityAdvisoryDeleteInvocation) callAsBody() (*ActionSecurityAdvisoryDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSecurityAdvisoryDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionSecurityAdvisoryDeleteInvocation) makeAllInputParams() *ActionSecurityAdvisoryDeleteRequest {
	return &ActionSecurityAdvisoryDeleteRequest{
		SecurityAdvisory: inv.makeInputParams(),
		Meta:             inv.makeMetaInputParams(),
	}
}

func (inv *ActionSecurityAdvisoryDeleteInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionSecurityAdvisoryDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
