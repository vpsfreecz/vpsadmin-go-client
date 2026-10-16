package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionSecurityAdvisoryCveShow is a type for action Security_advisory_cve#Show
type ActionSecurityAdvisoryCveShow struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryCveShow(client *Client) *ActionSecurityAdvisoryCveShow {
	return &ActionSecurityAdvisoryCveShow{
		Client: client,
	}
}

// ActionSecurityAdvisoryCveShowMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryCveShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryCveShowMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryCveShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryCveShowMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryCveShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryCveShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryCveShowMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryCveShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryCveShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryCveShowOutput is a type for action output parameters
type ActionSecurityAdvisoryCveShowOutput struct {
	CveId              string                            "json:\"cve_id\""
	Id                 int64                             "json:\"id\""
	SecurityAdvisory   *ActionSecurityAdvisoryShowOutput "json:\"security_advisory\""
	SecurityAdvisoryId int64                             "json:\"security_advisory_id\""
	Url                string                            "json:\"url\""
}

// Type for action response, including envelope
type ActionSecurityAdvisoryCveShowResponse struct {
	Action *ActionSecurityAdvisoryCveShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SecurityAdvisoryCve *ActionSecurityAdvisoryCveShowOutput "json:\"security_advisory_cve\""
	}

	// Action output without the namespace
	Output *ActionSecurityAdvisoryCveShowOutput
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryCveShow) Prepare() *ActionSecurityAdvisoryCveShowInvocation {
	return &ActionSecurityAdvisoryCveShowInvocation{
		Action: action,
		Path:   "/v7.0/security_advisory_cves/{security_advisory_cve_id}",
	}
}

// ActionSecurityAdvisoryCveShowInvocation is used to configure action for invocation
type ActionSecurityAdvisoryCveShowInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryCveShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryCveShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSecurityAdvisoryCveShowInvocation) SetPathParamInt(param string, value int64) *ActionSecurityAdvisoryCveShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSecurityAdvisoryCveShowInvocation) SetPathParamString(param string, value string) *ActionSecurityAdvisoryCveShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryCveShowInvocation) NewMetaInput() *ActionSecurityAdvisoryCveShowMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryCveShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryCveShowInvocation) SetMetaInput(input *ActionSecurityAdvisoryCveShowMetaGlobalInput) *ActionSecurityAdvisoryCveShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryCveShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryCveShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryCveShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionSecurityAdvisoryCveShowInvocation) Call() (*ActionSecurityAdvisoryCveShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionSecurityAdvisoryCveShowInvocation) CallContext(ctx context.Context) (*ActionSecurityAdvisoryCveShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionSecurityAdvisoryCveShowResponse{Action: inv.Action}
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
func (inv *ActionSecurityAdvisoryCveShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionSecurityAdvisoryCveShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionSecurityAdvisoryCveShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
