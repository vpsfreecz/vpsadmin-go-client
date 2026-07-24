package client

import (
	"net/url"
	"strings"
)

// ActionSecurityAdvisoryShow is a type for action Security_advisory#Show
type ActionSecurityAdvisoryShow struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryShow(client *Client) *ActionSecurityAdvisoryShow {
	return &ActionSecurityAdvisoryShow{
		Client: client,
	}
}

// ActionSecurityAdvisoryShowMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryShowMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryShowMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryShowMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryShowOutput is a type for action output parameters
type ActionSecurityAdvisoryShowOutput struct {
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
type ActionSecurityAdvisoryShowResponse struct {
	Action *ActionSecurityAdvisoryShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SecurityAdvisory *ActionSecurityAdvisoryShowOutput "json:\"security_advisory\""
	}

	// Action output without the namespace
	Output *ActionSecurityAdvisoryShowOutput
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryShow) Prepare() *ActionSecurityAdvisoryShowInvocation {
	return &ActionSecurityAdvisoryShowInvocation{
		Action: action,
		Path:   "/v7.0/security_advisories/{security_advisory_id}",
	}
}

// ActionSecurityAdvisoryShowInvocation is used to configure action for invocation
type ActionSecurityAdvisoryShowInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSecurityAdvisoryShowInvocation) SetPathParamInt(param string, value int64) *ActionSecurityAdvisoryShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSecurityAdvisoryShowInvocation) SetPathParamString(param string, value string) *ActionSecurityAdvisoryShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryShowInvocation) NewMetaInput() *ActionSecurityAdvisoryShowMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryShowInvocation) SetMetaInput(input *ActionSecurityAdvisoryShowMetaGlobalInput) *ActionSecurityAdvisoryShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSecurityAdvisoryShowInvocation) Call() (*ActionSecurityAdvisoryShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionSecurityAdvisoryShowInvocation) callAsQuery() (*ActionSecurityAdvisoryShowResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionSecurityAdvisoryShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SecurityAdvisory
	}
	return resp, err
}

func (inv *ActionSecurityAdvisoryShowInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			queryValue, err := convertCustomToString(inv.MetaInput.Includes)
			if err != nil {
				return err
			}
			ret["_meta[includes]"] = queryValue
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}

	return nil
}
