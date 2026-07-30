package client

import (
	"net/url"
	"strings"
)

// ActionSecurityAdvisoryCveUpdate is a type for action Security_advisory_cve#Update
type ActionSecurityAdvisoryCveUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryCveUpdate(client *Client) *ActionSecurityAdvisoryCveUpdate {
	return &ActionSecurityAdvisoryCveUpdate{
		Client: client,
	}
}

// ActionSecurityAdvisoryCveUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryCveUpdateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryCveUpdateMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryCveUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryCveUpdateMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryCveUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryCveUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryCveUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryCveUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryCveUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryCveUpdateInput is a type for action input parameters
type ActionSecurityAdvisoryCveUpdateInput struct {
	CveId                   string "json:\"cve_id\""
	ExpectedContentRevision int64  "json:\"expected_content_revision\""
	SecurityAdvisory        int64  "json:\"security_advisory\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCveId sets parameter CveId to value and selects it for sending
func (in *ActionSecurityAdvisoryCveUpdateInput) SetCveId(value string) *ActionSecurityAdvisoryCveUpdateInput {
	in.CveId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CveId"] = nil
	return in
}

// SetExpectedContentRevision sets parameter ExpectedContentRevision to value and selects it for sending
func (in *ActionSecurityAdvisoryCveUpdateInput) SetExpectedContentRevision(value int64) *ActionSecurityAdvisoryCveUpdateInput {
	in.ExpectedContentRevision = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ExpectedContentRevision"] = nil
	return in
}

// SetSecurityAdvisory sets parameter SecurityAdvisory to value and selects it for sending
func (in *ActionSecurityAdvisoryCveUpdateInput) SetSecurityAdvisory(value int64) *ActionSecurityAdvisoryCveUpdateInput {
	in.SecurityAdvisory = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SecurityAdvisory"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryCveUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryCveUpdateInput) SelectParameters(params ...string) *ActionSecurityAdvisoryCveUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionSecurityAdvisoryCveUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryCveUpdateInput) UnselectParameters(params ...string) *ActionSecurityAdvisoryCveUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionSecurityAdvisoryCveUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryCveUpdateRequest is a type for the entire action request
type ActionSecurityAdvisoryCveUpdateRequest struct {
	SecurityAdvisoryCve map[string]interface{} "json:\"security_advisory_cve\""
	Meta                map[string]interface{} "json:\"_meta\""
}

// ActionSecurityAdvisoryCveUpdateOutput is a type for action output parameters
type ActionSecurityAdvisoryCveUpdateOutput struct {
	CveId              string                            "json:\"cve_id\""
	Id                 int64                             "json:\"id\""
	SecurityAdvisory   *ActionSecurityAdvisoryShowOutput "json:\"security_advisory\""
	SecurityAdvisoryId int64                             "json:\"security_advisory_id\""
	Url                string                            "json:\"url\""
}

// Type for action response, including envelope
type ActionSecurityAdvisoryCveUpdateResponse struct {
	Action *ActionSecurityAdvisoryCveUpdate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SecurityAdvisoryCve *ActionSecurityAdvisoryCveUpdateOutput "json:\"security_advisory_cve\""
	}

	// Action output without the namespace
	Output *ActionSecurityAdvisoryCveUpdateOutput
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryCveUpdate) Prepare() *ActionSecurityAdvisoryCveUpdateInvocation {
	return &ActionSecurityAdvisoryCveUpdateInvocation{
		Action: action,
		Path:   "/v7.0/security_advisory_cves/{security_advisory_cve_id}",
	}
}

// ActionSecurityAdvisoryCveUpdateInvocation is used to configure action for invocation
type ActionSecurityAdvisoryCveUpdateInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryCveUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSecurityAdvisoryCveUpdateInput
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryCveUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSecurityAdvisoryCveUpdateInvocation) SetPathParamInt(param string, value int64) *ActionSecurityAdvisoryCveUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSecurityAdvisoryCveUpdateInvocation) SetPathParamString(param string, value string) *ActionSecurityAdvisoryCveUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSecurityAdvisoryCveUpdateInvocation) NewInput() *ActionSecurityAdvisoryCveUpdateInput {
	inv.Input = &ActionSecurityAdvisoryCveUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSecurityAdvisoryCveUpdateInvocation) SetInput(input *ActionSecurityAdvisoryCveUpdateInput) *ActionSecurityAdvisoryCveUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSecurityAdvisoryCveUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryCveUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryCveUpdateInvocation) NewMetaInput() *ActionSecurityAdvisoryCveUpdateMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryCveUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryCveUpdateInvocation) SetMetaInput(input *ActionSecurityAdvisoryCveUpdateMetaGlobalInput) *ActionSecurityAdvisoryCveUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryCveUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryCveUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryCveUpdateInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("SecurityAdvisory") {
			if !inv.IsParameterNil("SecurityAdvisory") {
				if inv.Input.SecurityAdvisory < 0 {
					verr.Add("security_advisory", "not a valid resource id")
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
func (inv *ActionSecurityAdvisoryCveUpdateInvocation) Call() (*ActionSecurityAdvisoryCveUpdateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionSecurityAdvisoryCveUpdateInvocation) callAsBody() (*ActionSecurityAdvisoryCveUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSecurityAdvisoryCveUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SecurityAdvisoryCve
	}
	return resp, err
}

func (inv *ActionSecurityAdvisoryCveUpdateInvocation) makeAllInputParams() *ActionSecurityAdvisoryCveUpdateRequest {
	return &ActionSecurityAdvisoryCveUpdateRequest{
		SecurityAdvisoryCve: inv.makeInputParams(),
		Meta:                inv.makeMetaInputParams(),
	}
}

func (inv *ActionSecurityAdvisoryCveUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("CveId") {
			ret["cve_id"] = inv.Input.CveId
		}
		if inv.IsParameterSelected("ExpectedContentRevision") {
			ret["expected_content_revision"] = inv.Input.ExpectedContentRevision
		}
		if inv.IsParameterSelected("SecurityAdvisory") {
			ret["security_advisory"] = inv.Input.SecurityAdvisory
		}
	}

	return ret
}

func (inv *ActionSecurityAdvisoryCveUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
