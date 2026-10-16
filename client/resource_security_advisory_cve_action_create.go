package client

import (
	"context"
	"net/http"
)

// ActionSecurityAdvisoryCveCreate is a type for action Security_advisory_cve#Create
type ActionSecurityAdvisoryCveCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryCveCreate(client *Client) *ActionSecurityAdvisoryCveCreate {
	return &ActionSecurityAdvisoryCveCreate{
		Client: client,
	}
}

// ActionSecurityAdvisoryCveCreateMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryCveCreateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryCveCreateMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryCveCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryCveCreateMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryCveCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryCveCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryCveCreateMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryCveCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryCveCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryCveCreateInput is a type for action input parameters
type ActionSecurityAdvisoryCveCreateInput struct {
	CveId                   string "json:\"cve_id\""
	ExpectedContentRevision int64  "json:\"expected_content_revision\""
	SecurityAdvisory        int64  "json:\"security_advisory\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCveId sets parameter CveId to value and selects it for sending
func (in *ActionSecurityAdvisoryCveCreateInput) SetCveId(value string) *ActionSecurityAdvisoryCveCreateInput {
	in.CveId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CveId"] = nil
	return in
}

// SetExpectedContentRevision sets parameter ExpectedContentRevision to value and selects it for sending
func (in *ActionSecurityAdvisoryCveCreateInput) SetExpectedContentRevision(value int64) *ActionSecurityAdvisoryCveCreateInput {
	in.ExpectedContentRevision = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ExpectedContentRevision"] = nil
	return in
}

// SetSecurityAdvisory sets parameter SecurityAdvisory to value and selects it for sending
func (in *ActionSecurityAdvisoryCveCreateInput) SetSecurityAdvisory(value int64) *ActionSecurityAdvisoryCveCreateInput {
	in.SecurityAdvisory = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SecurityAdvisory"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryCveCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryCveCreateInput) SelectParameters(params ...string) *ActionSecurityAdvisoryCveCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionSecurityAdvisoryCveCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryCveCreateInput) UnselectParameters(params ...string) *ActionSecurityAdvisoryCveCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionSecurityAdvisoryCveCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryCveCreateRequest is a type for the entire action request
type ActionSecurityAdvisoryCveCreateRequest struct {
	SecurityAdvisoryCve map[string]interface{} "json:\"security_advisory_cve\""
	Meta                map[string]interface{} "json:\"_meta\""
}

// ActionSecurityAdvisoryCveCreateOutput is a type for action output parameters
type ActionSecurityAdvisoryCveCreateOutput struct {
	CveId              string                            "json:\"cve_id\""
	Id                 int64                             "json:\"id\""
	SecurityAdvisory   *ActionSecurityAdvisoryShowOutput "json:\"security_advisory\""
	SecurityAdvisoryId int64                             "json:\"security_advisory_id\""
	Url                string                            "json:\"url\""
}

// Type for action response, including envelope
type ActionSecurityAdvisoryCveCreateResponse struct {
	Action *ActionSecurityAdvisoryCveCreate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SecurityAdvisoryCve *ActionSecurityAdvisoryCveCreateOutput "json:\"security_advisory_cve\""
	}

	// Action output without the namespace
	Output *ActionSecurityAdvisoryCveCreateOutput
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryCveCreate) Prepare() *ActionSecurityAdvisoryCveCreateInvocation {
	return &ActionSecurityAdvisoryCveCreateInvocation{
		Action: action,
		Path:   "/v7.0/security_advisory_cves",
	}
}

// ActionSecurityAdvisoryCveCreateInvocation is used to configure action for invocation
type ActionSecurityAdvisoryCveCreateInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryCveCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSecurityAdvisoryCveCreateInput
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryCveCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSecurityAdvisoryCveCreateInvocation) NewInput() *ActionSecurityAdvisoryCveCreateInput {
	inv.Input = &ActionSecurityAdvisoryCveCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSecurityAdvisoryCveCreateInvocation) SetInput(input *ActionSecurityAdvisoryCveCreateInput) *ActionSecurityAdvisoryCveCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSecurityAdvisoryCveCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryCveCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryCveCreateInvocation) NewMetaInput() *ActionSecurityAdvisoryCveCreateMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryCveCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryCveCreateInvocation) SetMetaInput(input *ActionSecurityAdvisoryCveCreateMetaGlobalInput) *ActionSecurityAdvisoryCveCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryCveCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryCveCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryCveCreateInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.Input != nil {
		if inv.IsParameterSelected("SecurityAdvisory") {
			if !inv.IsParameterNil("SecurityAdvisory") {
				if inv.Input.SecurityAdvisory < 0 {
					verr.Add("security_advisory", t(inv.Action.Client, "validation.invalid_resource_id", nil))
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

// Call invokes the action using a background context.
func (inv *ActionSecurityAdvisoryCveCreateInvocation) Call() (*ActionSecurityAdvisoryCveCreateResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionSecurityAdvisoryCveCreateInvocation) CallContext(ctx context.Context) (*ActionSecurityAdvisoryCveCreateResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionSecurityAdvisoryCveCreateResponse{Action: inv.Action}
	raw, err := inv.Action.Client.doRaw(req)
	if err != nil {
		return resp, err
	}
	if err := decodeTypedResponse(raw.Body, resp); err != nil {
		return resp, err
	}
	if resp.Status {
		if resp.Response == nil {
			return resp, errMalformedResponse
		}
		if resp.Response.SecurityAdvisoryCve == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.SecurityAdvisoryCve
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionSecurityAdvisoryCveCreateInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionSecurityAdvisoryCveCreateInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	input := inv.makeAllInputParams()
	return inv.Action.Client.newBodyRequest(ctx, "POST", inv.Path, input)
}

func (inv *ActionSecurityAdvisoryCveCreateInvocation) makeAllInputParams() *ActionSecurityAdvisoryCveCreateRequest {
	return &ActionSecurityAdvisoryCveCreateRequest{
		SecurityAdvisoryCve: inv.makeInputParams(),
		Meta:                inv.makeMetaInputParams(),
	}
}

func (inv *ActionSecurityAdvisoryCveCreateInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionSecurityAdvisoryCveCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
