package client

import (
	"net/url"
	"strings"
)

// ActionSecurityAdvisoryUpdate is a type for action Security_advisory#Update
type ActionSecurityAdvisoryUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryUpdate(client *Client) *ActionSecurityAdvisoryUpdate {
	return &ActionSecurityAdvisoryUpdate{
		Client: client,
	}
}

// ActionSecurityAdvisoryUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryUpdateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryUpdateInput is a type for action input parameters
type ActionSecurityAdvisoryUpdateInput struct {
	CsDescription           string "json:\"cs_description\""
	CsResponse              string "json:\"cs_response\""
	CsSummary               string "json:\"cs_summary\""
	EnDescription           string "json:\"en_description\""
	EnResponse              string "json:\"en_response\""
	EnSummary               string "json:\"en_summary\""
	ExpectedContentRevision int64  "json:\"expected_content_revision\""
	Name                    string "json:\"name\""
	PublishedAt             string "json:\"published_at\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCsDescription sets parameter CsDescription to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateInput) SetCsDescription(value string) *ActionSecurityAdvisoryUpdateInput {
	in.CsDescription = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CsDescription"] = nil
	return in
}

// SetCsResponse sets parameter CsResponse to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateInput) SetCsResponse(value string) *ActionSecurityAdvisoryUpdateInput {
	in.CsResponse = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CsResponse"] = nil
	return in
}

// SetCsSummary sets parameter CsSummary to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateInput) SetCsSummary(value string) *ActionSecurityAdvisoryUpdateInput {
	in.CsSummary = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CsSummary"] = nil
	return in
}

// SetEnDescription sets parameter EnDescription to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateInput) SetEnDescription(value string) *ActionSecurityAdvisoryUpdateInput {
	in.EnDescription = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnDescription"] = nil
	return in
}

// SetEnResponse sets parameter EnResponse to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateInput) SetEnResponse(value string) *ActionSecurityAdvisoryUpdateInput {
	in.EnResponse = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnResponse"] = nil
	return in
}

// SetEnSummary sets parameter EnSummary to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateInput) SetEnSummary(value string) *ActionSecurityAdvisoryUpdateInput {
	in.EnSummary = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnSummary"] = nil
	return in
}

// SetExpectedContentRevision sets parameter ExpectedContentRevision to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateInput) SetExpectedContentRevision(value int64) *ActionSecurityAdvisoryUpdateInput {
	in.ExpectedContentRevision = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ExpectedContentRevision"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateInput) SetName(value string) *ActionSecurityAdvisoryUpdateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNameNil(false)
	in._selectedParameters["Name"] = nil
	return in
}

// SetNameNil sets parameter Name to nil and selects it for sending
func (in *ActionSecurityAdvisoryUpdateInput) SetNameNil(set bool) *ActionSecurityAdvisoryUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Name"] = nil
		in.SelectParameters("Name")
	} else {
		delete(in._nilParameters, "Name")
	}
	return in
}

// SetPublishedAt sets parameter PublishedAt to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateInput) SetPublishedAt(value string) *ActionSecurityAdvisoryUpdateInput {
	in.PublishedAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetPublishedAtNil(false)
	in._selectedParameters["PublishedAt"] = nil
	return in
}

// SetPublishedAtNil sets parameter PublishedAt to nil and selects it for sending
func (in *ActionSecurityAdvisoryUpdateInput) SetPublishedAtNil(set bool) *ActionSecurityAdvisoryUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["PublishedAt"] = nil
		in.SelectParameters("PublishedAt")
	} else {
		delete(in._nilParameters, "PublishedAt")
	}
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryUpdateInput) SelectParameters(params ...string) *ActionSecurityAdvisoryUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionSecurityAdvisoryUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryUpdateInput) UnselectParameters(params ...string) *ActionSecurityAdvisoryUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionSecurityAdvisoryUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryUpdateRequest is a type for the entire action request
type ActionSecurityAdvisoryUpdateRequest struct {
	SecurityAdvisory map[string]interface{} "json:\"security_advisory\""
	Meta             map[string]interface{} "json:\"_meta\""
}

// ActionSecurityAdvisoryUpdateOutput is a type for action output parameters
type ActionSecurityAdvisoryUpdateOutput struct {
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
type ActionSecurityAdvisoryUpdateResponse struct {
	Action *ActionSecurityAdvisoryUpdate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SecurityAdvisory *ActionSecurityAdvisoryUpdateOutput "json:\"security_advisory\""
	}

	// Action output without the namespace
	Output *ActionSecurityAdvisoryUpdateOutput
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryUpdate) Prepare() *ActionSecurityAdvisoryUpdateInvocation {
	return &ActionSecurityAdvisoryUpdateInvocation{
		Action: action,
		Path:   "/v7.0/security_advisories/{security_advisory_id}",
	}
}

// ActionSecurityAdvisoryUpdateInvocation is used to configure action for invocation
type ActionSecurityAdvisoryUpdateInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSecurityAdvisoryUpdateInput
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSecurityAdvisoryUpdateInvocation) SetPathParamInt(param string, value int64) *ActionSecurityAdvisoryUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSecurityAdvisoryUpdateInvocation) SetPathParamString(param string, value string) *ActionSecurityAdvisoryUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSecurityAdvisoryUpdateInvocation) NewInput() *ActionSecurityAdvisoryUpdateInput {
	inv.Input = &ActionSecurityAdvisoryUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSecurityAdvisoryUpdateInvocation) SetInput(input *ActionSecurityAdvisoryUpdateInput) *ActionSecurityAdvisoryUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSecurityAdvisoryUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryUpdateInvocation) NewMetaInput() *ActionSecurityAdvisoryUpdateMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryUpdateInvocation) SetMetaInput(input *ActionSecurityAdvisoryUpdateMetaGlobalInput) *ActionSecurityAdvisoryUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryUpdateInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("PublishedAt") {
			if !inv.IsParameterNil("PublishedAt") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.PublishedAt)
				if !ok {
					verr.Add("published_at", "not a valid datetime")
				} else {
					inv.Input.PublishedAt = normalized
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
func (inv *ActionSecurityAdvisoryUpdateInvocation) Call() (*ActionSecurityAdvisoryUpdateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionSecurityAdvisoryUpdateInvocation) callAsBody() (*ActionSecurityAdvisoryUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSecurityAdvisoryUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SecurityAdvisory
	}
	return resp, err
}

func (inv *ActionSecurityAdvisoryUpdateInvocation) makeAllInputParams() *ActionSecurityAdvisoryUpdateRequest {
	return &ActionSecurityAdvisoryUpdateRequest{
		SecurityAdvisory: inv.makeInputParams(),
		Meta:             inv.makeMetaInputParams(),
	}
}

func (inv *ActionSecurityAdvisoryUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("CsDescription") {
			ret["cs_description"] = inv.Input.CsDescription
		}
		if inv.IsParameterSelected("CsResponse") {
			ret["cs_response"] = inv.Input.CsResponse
		}
		if inv.IsParameterSelected("CsSummary") {
			ret["cs_summary"] = inv.Input.CsSummary
		}
		if inv.IsParameterSelected("EnDescription") {
			ret["en_description"] = inv.Input.EnDescription
		}
		if inv.IsParameterSelected("EnResponse") {
			ret["en_response"] = inv.Input.EnResponse
		}
		if inv.IsParameterSelected("EnSummary") {
			ret["en_summary"] = inv.Input.EnSummary
		}
		if inv.IsParameterSelected("ExpectedContentRevision") {
			ret["expected_content_revision"] = inv.Input.ExpectedContentRevision
		}
		if inv.IsParameterSelected("Name") {
			if inv.IsParameterNil("Name") {
				ret["name"] = nil
			} else {
				ret["name"] = inv.Input.Name
			}
		}
		if inv.IsParameterSelected("PublishedAt") {
			if inv.IsParameterNil("PublishedAt") {
				ret["published_at"] = nil
			} else {
				ret["published_at"] = inv.Input.PublishedAt
			}
		}
	}

	return ret
}

func (inv *ActionSecurityAdvisoryUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
