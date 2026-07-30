package client

import ()

// ActionSecurityAdvisoryCreate is a type for action Security_advisory#Create
type ActionSecurityAdvisoryCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryCreate(client *Client) *ActionSecurityAdvisoryCreate {
	return &ActionSecurityAdvisoryCreate{
		Client: client,
	}
}

// ActionSecurityAdvisoryCreateMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryCreateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryCreateMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryCreateMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryCreateMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryCreateInput is a type for action input parameters
type ActionSecurityAdvisoryCreateInput struct {
	CsDescription string "json:\"cs_description\""
	CsResponse    string "json:\"cs_response\""
	CsSummary     string "json:\"cs_summary\""
	Cve           string "json:\"cve\""
	EnDescription string "json:\"en_description\""
	EnResponse    string "json:\"en_response\""
	EnSummary     string "json:\"en_summary\""
	ExternalId    string "json:\"external_id\""
	Name          string "json:\"name\""
	PublishedAt   string "json:\"published_at\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCsDescription sets parameter CsDescription to value and selects it for sending
func (in *ActionSecurityAdvisoryCreateInput) SetCsDescription(value string) *ActionSecurityAdvisoryCreateInput {
	in.CsDescription = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CsDescription"] = nil
	return in
}

// SetCsResponse sets parameter CsResponse to value and selects it for sending
func (in *ActionSecurityAdvisoryCreateInput) SetCsResponse(value string) *ActionSecurityAdvisoryCreateInput {
	in.CsResponse = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CsResponse"] = nil
	return in
}

// SetCsSummary sets parameter CsSummary to value and selects it for sending
func (in *ActionSecurityAdvisoryCreateInput) SetCsSummary(value string) *ActionSecurityAdvisoryCreateInput {
	in.CsSummary = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CsSummary"] = nil
	return in
}

// SetCve sets parameter Cve to value and selects it for sending
func (in *ActionSecurityAdvisoryCreateInput) SetCve(value string) *ActionSecurityAdvisoryCreateInput {
	in.Cve = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetCveNil(false)
	in._selectedParameters["Cve"] = nil
	return in
}

// SetCveNil sets parameter Cve to nil and selects it for sending
func (in *ActionSecurityAdvisoryCreateInput) SetCveNil(set bool) *ActionSecurityAdvisoryCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Cve"] = nil
		in.SelectParameters("Cve")
	} else {
		delete(in._nilParameters, "Cve")
	}
	return in
}

// SetEnDescription sets parameter EnDescription to value and selects it for sending
func (in *ActionSecurityAdvisoryCreateInput) SetEnDescription(value string) *ActionSecurityAdvisoryCreateInput {
	in.EnDescription = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnDescription"] = nil
	return in
}

// SetEnResponse sets parameter EnResponse to value and selects it for sending
func (in *ActionSecurityAdvisoryCreateInput) SetEnResponse(value string) *ActionSecurityAdvisoryCreateInput {
	in.EnResponse = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnResponse"] = nil
	return in
}

// SetEnSummary sets parameter EnSummary to value and selects it for sending
func (in *ActionSecurityAdvisoryCreateInput) SetEnSummary(value string) *ActionSecurityAdvisoryCreateInput {
	in.EnSummary = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnSummary"] = nil
	return in
}

// SetExternalId sets parameter ExternalId to value and selects it for sending
func (in *ActionSecurityAdvisoryCreateInput) SetExternalId(value string) *ActionSecurityAdvisoryCreateInput {
	in.ExternalId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetExternalIdNil(false)
	in._selectedParameters["ExternalId"] = nil
	return in
}

// SetExternalIdNil sets parameter ExternalId to nil and selects it for sending
func (in *ActionSecurityAdvisoryCreateInput) SetExternalIdNil(set bool) *ActionSecurityAdvisoryCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["ExternalId"] = nil
		in.SelectParameters("ExternalId")
	} else {
		delete(in._nilParameters, "ExternalId")
	}
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionSecurityAdvisoryCreateInput) SetName(value string) *ActionSecurityAdvisoryCreateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNameNil(false)
	in._selectedParameters["Name"] = nil
	return in
}

// SetNameNil sets parameter Name to nil and selects it for sending
func (in *ActionSecurityAdvisoryCreateInput) SetNameNil(set bool) *ActionSecurityAdvisoryCreateInput {
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
func (in *ActionSecurityAdvisoryCreateInput) SetPublishedAt(value string) *ActionSecurityAdvisoryCreateInput {
	in.PublishedAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetPublishedAtNil(false)
	in._selectedParameters["PublishedAt"] = nil
	return in
}

// SetPublishedAtNil sets parameter PublishedAt to nil and selects it for sending
func (in *ActionSecurityAdvisoryCreateInput) SetPublishedAtNil(set bool) *ActionSecurityAdvisoryCreateInput {
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

// SelectParameters sets parameters from ActionSecurityAdvisoryCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryCreateInput) SelectParameters(params ...string) *ActionSecurityAdvisoryCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionSecurityAdvisoryCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryCreateInput) UnselectParameters(params ...string) *ActionSecurityAdvisoryCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionSecurityAdvisoryCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryCreateRequest is a type for the entire action request
type ActionSecurityAdvisoryCreateRequest struct {
	SecurityAdvisory map[string]interface{} "json:\"security_advisory\""
	Meta             map[string]interface{} "json:\"_meta\""
}

// ActionSecurityAdvisoryCreateOutput is a type for action output parameters
type ActionSecurityAdvisoryCreateOutput struct {
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
type ActionSecurityAdvisoryCreateResponse struct {
	Action *ActionSecurityAdvisoryCreate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SecurityAdvisory *ActionSecurityAdvisoryCreateOutput "json:\"security_advisory\""
	}

	// Action output without the namespace
	Output *ActionSecurityAdvisoryCreateOutput
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryCreate) Prepare() *ActionSecurityAdvisoryCreateInvocation {
	return &ActionSecurityAdvisoryCreateInvocation{
		Action: action,
		Path:   "/v7.0/security_advisories",
	}
}

// ActionSecurityAdvisoryCreateInvocation is used to configure action for invocation
type ActionSecurityAdvisoryCreateInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSecurityAdvisoryCreateInput
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSecurityAdvisoryCreateInvocation) NewInput() *ActionSecurityAdvisoryCreateInput {
	inv.Input = &ActionSecurityAdvisoryCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSecurityAdvisoryCreateInvocation) SetInput(input *ActionSecurityAdvisoryCreateInput) *ActionSecurityAdvisoryCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSecurityAdvisoryCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryCreateInvocation) NewMetaInput() *ActionSecurityAdvisoryCreateMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryCreateInvocation) SetMetaInput(input *ActionSecurityAdvisoryCreateMetaGlobalInput) *ActionSecurityAdvisoryCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryCreateInvocation) validate() error {
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
func (inv *ActionSecurityAdvisoryCreateInvocation) Call() (*ActionSecurityAdvisoryCreateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionSecurityAdvisoryCreateInvocation) callAsBody() (*ActionSecurityAdvisoryCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSecurityAdvisoryCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SecurityAdvisory
	}
	return resp, err
}

func (inv *ActionSecurityAdvisoryCreateInvocation) makeAllInputParams() *ActionSecurityAdvisoryCreateRequest {
	return &ActionSecurityAdvisoryCreateRequest{
		SecurityAdvisory: inv.makeInputParams(),
		Meta:             inv.makeMetaInputParams(),
	}
}

func (inv *ActionSecurityAdvisoryCreateInvocation) makeInputParams() map[string]interface{} {
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
		if inv.IsParameterSelected("Cve") {
			if inv.IsParameterNil("Cve") {
				ret["cve"] = nil
			} else {
				ret["cve"] = inv.Input.Cve
			}
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
		if inv.IsParameterSelected("ExternalId") {
			if inv.IsParameterNil("ExternalId") {
				ret["external_id"] = nil
			} else {
				ret["external_id"] = inv.Input.ExternalId
			}
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

func (inv *ActionSecurityAdvisoryCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
