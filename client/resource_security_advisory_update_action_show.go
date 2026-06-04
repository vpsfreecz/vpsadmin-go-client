package client

import (
	"net/url"
	"strings"
)

// ActionSecurityAdvisoryUpdateShow is a type for action Security_advisory_update#Show
type ActionSecurityAdvisoryUpdateShow struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryUpdateShow(client *Client) *ActionSecurityAdvisoryUpdateShow {
	return &ActionSecurityAdvisoryUpdateShow{
		Client: client,
	}
}

// ActionSecurityAdvisoryUpdateShowMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryUpdateShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateShowMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryUpdateShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateShowMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryUpdateShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryUpdateShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryUpdateShowMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryUpdateShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryUpdateShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryUpdateShowOutput is a type for action output parameters
type ActionSecurityAdvisoryUpdateShowOutput struct {
	CreatedAt        string                            "json:\"created_at\""
	CsMessage        string                            "json:\"cs_message\""
	CsSummary        string                            "json:\"cs_summary\""
	EnMessage        string                            "json:\"en_message\""
	EnSummary        string                            "json:\"en_summary\""
	Id               int64                             "json:\"id\""
	Name             string                            "json:\"name\""
	ReportedBy       *ActionUserShowOutput             "json:\"reported_by\""
	ReporterName     string                            "json:\"reporter_name\""
	SecurityAdvisory *ActionSecurityAdvisoryShowOutput "json:\"security_advisory\""
	State            string                            "json:\"state\""
	UpdatedAt        string                            "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionSecurityAdvisoryUpdateShowResponse struct {
	Action *ActionSecurityAdvisoryUpdateShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SecurityAdvisoryUpdate *ActionSecurityAdvisoryUpdateShowOutput "json:\"security_advisory_update\""
	}

	// Action output without the namespace
	Output *ActionSecurityAdvisoryUpdateShowOutput
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryUpdateShow) Prepare() *ActionSecurityAdvisoryUpdateShowInvocation {
	return &ActionSecurityAdvisoryUpdateShowInvocation{
		Action: action,
		Path:   "/v7.0/security_advisory_updates/{security_advisory_update_id}",
	}
}

// ActionSecurityAdvisoryUpdateShowInvocation is used to configure action for invocation
type ActionSecurityAdvisoryUpdateShowInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryUpdateShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryUpdateShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSecurityAdvisoryUpdateShowInvocation) SetPathParamInt(param string, value int64) *ActionSecurityAdvisoryUpdateShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSecurityAdvisoryUpdateShowInvocation) SetPathParamString(param string, value string) *ActionSecurityAdvisoryUpdateShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryUpdateShowInvocation) NewMetaInput() *ActionSecurityAdvisoryUpdateShowMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryUpdateShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryUpdateShowInvocation) SetMetaInput(input *ActionSecurityAdvisoryUpdateShowMetaGlobalInput) *ActionSecurityAdvisoryUpdateShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryUpdateShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryUpdateShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryUpdateShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSecurityAdvisoryUpdateShowInvocation) Call() (*ActionSecurityAdvisoryUpdateShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionSecurityAdvisoryUpdateShowInvocation) callAsQuery() (*ActionSecurityAdvisoryUpdateShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionSecurityAdvisoryUpdateShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SecurityAdvisoryUpdate
	}
	return resp, err
}

func (inv *ActionSecurityAdvisoryUpdateShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
