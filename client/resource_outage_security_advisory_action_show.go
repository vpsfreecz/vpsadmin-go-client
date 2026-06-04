package client

import (
	"net/url"
	"strings"
)

// ActionOutageSecurityAdvisoryShow is a type for action Outage_security_advisory#Show
type ActionOutageSecurityAdvisoryShow struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageSecurityAdvisoryShow(client *Client) *ActionOutageSecurityAdvisoryShow {
	return &ActionOutageSecurityAdvisoryShow{
		Client: client,
	}
}

// ActionOutageSecurityAdvisoryShowMetaGlobalInput is a type for action global meta input parameters
type ActionOutageSecurityAdvisoryShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageSecurityAdvisoryShowMetaGlobalInput) SetIncludes(value string) *ActionOutageSecurityAdvisoryShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageSecurityAdvisoryShowMetaGlobalInput) SetNo(value bool) *ActionOutageSecurityAdvisoryShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageSecurityAdvisoryShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageSecurityAdvisoryShowMetaGlobalInput) SelectParameters(params ...string) *ActionOutageSecurityAdvisoryShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageSecurityAdvisoryShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageSecurityAdvisoryShowOutput is a type for action output parameters
type ActionOutageSecurityAdvisoryShowOutput struct {
	Id                 int64                             "json:\"id\""
	Outage             *ActionOutageShowOutput           "json:\"outage\""
	OutageId           int64                             "json:\"outage_id\""
	SecurityAdvisory   *ActionSecurityAdvisoryShowOutput "json:\"security_advisory\""
	SecurityAdvisoryId int64                             "json:\"security_advisory_id\""
}

// Type for action response, including envelope
type ActionOutageSecurityAdvisoryShowResponse struct {
	Action *ActionOutageSecurityAdvisoryShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OutageSecurityAdvisory *ActionOutageSecurityAdvisoryShowOutput "json:\"outage_security_advisory\""
	}

	// Action output without the namespace
	Output *ActionOutageSecurityAdvisoryShowOutput
}

// Prepare the action for invocation
func (action *ActionOutageSecurityAdvisoryShow) Prepare() *ActionOutageSecurityAdvisoryShowInvocation {
	return &ActionOutageSecurityAdvisoryShowInvocation{
		Action: action,
		Path:   "/v7.0/outage_security_advisories/{outage_security_advisory_id}",
	}
}

// ActionOutageSecurityAdvisoryShowInvocation is used to configure action for invocation
type ActionOutageSecurityAdvisoryShowInvocation struct {
	// Pointer to the action
	Action *ActionOutageSecurityAdvisoryShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionOutageSecurityAdvisoryShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOutageSecurityAdvisoryShowInvocation) SetPathParamInt(param string, value int64) *ActionOutageSecurityAdvisoryShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOutageSecurityAdvisoryShowInvocation) SetPathParamString(param string, value string) *ActionOutageSecurityAdvisoryShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOutageSecurityAdvisoryShowInvocation) NewMetaInput() *ActionOutageSecurityAdvisoryShowMetaGlobalInput {
	inv.MetaInput = &ActionOutageSecurityAdvisoryShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageSecurityAdvisoryShowInvocation) SetMetaInput(input *ActionOutageSecurityAdvisoryShowMetaGlobalInput) *ActionOutageSecurityAdvisoryShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageSecurityAdvisoryShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOutageSecurityAdvisoryShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionOutageSecurityAdvisoryShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOutageSecurityAdvisoryShowInvocation) Call() (*ActionOutageSecurityAdvisoryShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionOutageSecurityAdvisoryShowInvocation) callAsQuery() (*ActionOutageSecurityAdvisoryShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOutageSecurityAdvisoryShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OutageSecurityAdvisory
	}
	return resp, err
}

func (inv *ActionOutageSecurityAdvisoryShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
