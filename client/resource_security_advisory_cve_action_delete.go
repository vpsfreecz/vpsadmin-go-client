package client

import (
	"net/url"
	"strings"
)

// ActionSecurityAdvisoryCveDelete is a type for action Security_advisory_cve#Delete
type ActionSecurityAdvisoryCveDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryCveDelete(client *Client) *ActionSecurityAdvisoryCveDelete {
	return &ActionSecurityAdvisoryCveDelete{
		Client: client,
	}
}

// ActionSecurityAdvisoryCveDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryCveDeleteMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryCveDeleteMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryCveDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryCveDeleteMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryCveDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryCveDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryCveDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryCveDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryCveDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryCveDeleteInput is a type for action input parameters
type ActionSecurityAdvisoryCveDeleteInput struct {
	ExpectedContentRevision int64 "json:\"expected_content_revision\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetExpectedContentRevision sets parameter ExpectedContentRevision to value and selects it for sending
func (in *ActionSecurityAdvisoryCveDeleteInput) SetExpectedContentRevision(value int64) *ActionSecurityAdvisoryCveDeleteInput {
	in.ExpectedContentRevision = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ExpectedContentRevision"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryCveDeleteInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryCveDeleteInput) SelectParameters(params ...string) *ActionSecurityAdvisoryCveDeleteInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionSecurityAdvisoryCveDeleteInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryCveDeleteInput) UnselectParameters(params ...string) *ActionSecurityAdvisoryCveDeleteInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionSecurityAdvisoryCveDeleteInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryCveDeleteRequest is a type for the entire action request
type ActionSecurityAdvisoryCveDeleteRequest struct {
	SecurityAdvisoryCve map[string]interface{} "json:\"security_advisory_cve\""
	Meta                map[string]interface{} "json:\"_meta\""
}

// Type for action response, including envelope
type ActionSecurityAdvisoryCveDeleteResponse struct {
	Action *ActionSecurityAdvisoryCveDelete "json:\"-\""
	*Envelope
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryCveDelete) Prepare() *ActionSecurityAdvisoryCveDeleteInvocation {
	return &ActionSecurityAdvisoryCveDeleteInvocation{
		Action: action,
		Path:   "/v7.0/security_advisory_cves/{security_advisory_cve_id}",
	}
}

// ActionSecurityAdvisoryCveDeleteInvocation is used to configure action for invocation
type ActionSecurityAdvisoryCveDeleteInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryCveDelete

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSecurityAdvisoryCveDeleteInput
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryCveDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSecurityAdvisoryCveDeleteInvocation) SetPathParamInt(param string, value int64) *ActionSecurityAdvisoryCveDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSecurityAdvisoryCveDeleteInvocation) SetPathParamString(param string, value string) *ActionSecurityAdvisoryCveDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSecurityAdvisoryCveDeleteInvocation) NewInput() *ActionSecurityAdvisoryCveDeleteInput {
	inv.Input = &ActionSecurityAdvisoryCveDeleteInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSecurityAdvisoryCveDeleteInvocation) SetInput(input *ActionSecurityAdvisoryCveDeleteInput) *ActionSecurityAdvisoryCveDeleteInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSecurityAdvisoryCveDeleteInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryCveDeleteInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryCveDeleteInvocation) NewMetaInput() *ActionSecurityAdvisoryCveDeleteMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryCveDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryCveDeleteInvocation) SetMetaInput(input *ActionSecurityAdvisoryCveDeleteMetaGlobalInput) *ActionSecurityAdvisoryCveDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryCveDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryCveDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryCveDeleteInvocation) validate() error {
	verr := NewValidationError()
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
func (inv *ActionSecurityAdvisoryCveDeleteInvocation) Call() (*ActionSecurityAdvisoryCveDeleteResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionSecurityAdvisoryCveDeleteInvocation) callAsBody() (*ActionSecurityAdvisoryCveDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSecurityAdvisoryCveDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionSecurityAdvisoryCveDeleteInvocation) makeAllInputParams() *ActionSecurityAdvisoryCveDeleteRequest {
	return &ActionSecurityAdvisoryCveDeleteRequest{
		SecurityAdvisoryCve: inv.makeInputParams(),
		Meta:                inv.makeMetaInputParams(),
	}
}

func (inv *ActionSecurityAdvisoryCveDeleteInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("ExpectedContentRevision") {
			ret["expected_content_revision"] = inv.Input.ExpectedContentRevision
		}
	}

	return ret
}

func (inv *ActionSecurityAdvisoryCveDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
