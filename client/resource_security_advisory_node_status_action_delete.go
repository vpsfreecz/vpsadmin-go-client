package client

import (
	"net/url"
	"strings"
)

// ActionSecurityAdvisoryNodeStatusDelete is a type for action Security_advisory.Node_status#Delete
type ActionSecurityAdvisoryNodeStatusDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryNodeStatusDelete(client *Client) *ActionSecurityAdvisoryNodeStatusDelete {
	return &ActionSecurityAdvisoryNodeStatusDelete{
		Client: client,
	}
}

// ActionSecurityAdvisoryNodeStatusDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryNodeStatusDeleteMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusDeleteMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryNodeStatusDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusDeleteMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryNodeStatusDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryNodeStatusDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryNodeStatusDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryNodeStatusDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryNodeStatusDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryNodeStatusDeleteInput is a type for action input parameters
type ActionSecurityAdvisoryNodeStatusDeleteInput struct {
	ExpectedContentRevision int64 "json:\"expected_content_revision\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetExpectedContentRevision sets parameter ExpectedContentRevision to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusDeleteInput) SetExpectedContentRevision(value int64) *ActionSecurityAdvisoryNodeStatusDeleteInput {
	in.ExpectedContentRevision = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ExpectedContentRevision"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryNodeStatusDeleteInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryNodeStatusDeleteInput) SelectParameters(params ...string) *ActionSecurityAdvisoryNodeStatusDeleteInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionSecurityAdvisoryNodeStatusDeleteInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryNodeStatusDeleteInput) UnselectParameters(params ...string) *ActionSecurityAdvisoryNodeStatusDeleteInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionSecurityAdvisoryNodeStatusDeleteInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryNodeStatusDeleteRequest is a type for the entire action request
type ActionSecurityAdvisoryNodeStatusDeleteRequest struct {
	NodeStatus map[string]interface{} "json:\"node_status\""
	Meta       map[string]interface{} "json:\"_meta\""
}

// Type for action response, including envelope
type ActionSecurityAdvisoryNodeStatusDeleteResponse struct {
	Action *ActionSecurityAdvisoryNodeStatusDelete "json:\"-\""
	*Envelope
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryNodeStatusDelete) Prepare() *ActionSecurityAdvisoryNodeStatusDeleteInvocation {
	return &ActionSecurityAdvisoryNodeStatusDeleteInvocation{
		Action: action,
		Path:   "/v7.0/security_advisories/{security_advisory_id}/node_statuses/{node_status_id}",
	}
}

// ActionSecurityAdvisoryNodeStatusDeleteInvocation is used to configure action for invocation
type ActionSecurityAdvisoryNodeStatusDeleteInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryNodeStatusDelete

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSecurityAdvisoryNodeStatusDeleteInput
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryNodeStatusDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSecurityAdvisoryNodeStatusDeleteInvocation) SetPathParamInt(param string, value int64) *ActionSecurityAdvisoryNodeStatusDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSecurityAdvisoryNodeStatusDeleteInvocation) SetPathParamString(param string, value string) *ActionSecurityAdvisoryNodeStatusDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSecurityAdvisoryNodeStatusDeleteInvocation) NewInput() *ActionSecurityAdvisoryNodeStatusDeleteInput {
	inv.Input = &ActionSecurityAdvisoryNodeStatusDeleteInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSecurityAdvisoryNodeStatusDeleteInvocation) SetInput(input *ActionSecurityAdvisoryNodeStatusDeleteInput) *ActionSecurityAdvisoryNodeStatusDeleteInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSecurityAdvisoryNodeStatusDeleteInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryNodeStatusDeleteInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryNodeStatusDeleteInvocation) NewMetaInput() *ActionSecurityAdvisoryNodeStatusDeleteMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryNodeStatusDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryNodeStatusDeleteInvocation) SetMetaInput(input *ActionSecurityAdvisoryNodeStatusDeleteMetaGlobalInput) *ActionSecurityAdvisoryNodeStatusDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryNodeStatusDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryNodeStatusDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryNodeStatusDeleteInvocation) validate() error {
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
func (inv *ActionSecurityAdvisoryNodeStatusDeleteInvocation) Call() (*ActionSecurityAdvisoryNodeStatusDeleteResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionSecurityAdvisoryNodeStatusDeleteInvocation) callAsBody() (*ActionSecurityAdvisoryNodeStatusDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSecurityAdvisoryNodeStatusDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionSecurityAdvisoryNodeStatusDeleteInvocation) makeAllInputParams() *ActionSecurityAdvisoryNodeStatusDeleteRequest {
	return &ActionSecurityAdvisoryNodeStatusDeleteRequest{
		NodeStatus: inv.makeInputParams(),
		Meta:       inv.makeMetaInputParams(),
	}
}

func (inv *ActionSecurityAdvisoryNodeStatusDeleteInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("ExpectedContentRevision") {
			ret["expected_content_revision"] = inv.Input.ExpectedContentRevision
		}
	}

	return ret
}

func (inv *ActionSecurityAdvisoryNodeStatusDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
