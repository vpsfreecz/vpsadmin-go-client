package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionSystemConfigShow is a type for action System_config#Show
type ActionSystemConfigShow struct {
	// Pointer to client
	Client *Client
}

func NewActionSystemConfigShow(client *Client) *ActionSystemConfigShow {
	return &ActionSystemConfigShow{
		Client: client,
	}
}

// ActionSystemConfigShowMetaGlobalInput is a type for action global meta input parameters
type ActionSystemConfigShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSystemConfigShowMetaGlobalInput) SetIncludes(value string) *ActionSystemConfigShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSystemConfigShowMetaGlobalInput) SetNo(value bool) *ActionSystemConfigShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSystemConfigShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSystemConfigShowMetaGlobalInput) SelectParameters(params ...string) *ActionSystemConfigShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSystemConfigShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSystemConfigShowOutput is a type for action output parameters
type ActionSystemConfigShowOutput struct {
	Category     string "json:\"category\""
	Description  string "json:\"description\""
	Label        string "json:\"label\""
	Localized    bool   "json:\"localized\""
	MinUserLevel int64  "json:\"min_user_level\""
	Name         string "json:\"name\""
	Type         string "json:\"type\""
}

// Type for action response, including envelope
type ActionSystemConfigShowResponse struct {
	Action *ActionSystemConfigShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SystemConfig *ActionSystemConfigShowOutput "json:\"system_config\""
	}

	// Action output without the namespace
	Output *ActionSystemConfigShowOutput
}

// Prepare the action for invocation
func (action *ActionSystemConfigShow) Prepare() *ActionSystemConfigShowInvocation {
	return &ActionSystemConfigShowInvocation{
		Action: action,
		Path:   "/v7.0/system_configs/{category}/{name}",
	}
}

// ActionSystemConfigShowInvocation is used to configure action for invocation
type ActionSystemConfigShowInvocation struct {
	// Pointer to the action
	Action *ActionSystemConfigShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionSystemConfigShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSystemConfigShowInvocation) SetPathParamInt(param string, value int64) *ActionSystemConfigShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSystemConfigShowInvocation) SetPathParamString(param string, value string) *ActionSystemConfigShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSystemConfigShowInvocation) NewMetaInput() *ActionSystemConfigShowMetaGlobalInput {
	inv.MetaInput = &ActionSystemConfigShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSystemConfigShowInvocation) SetMetaInput(input *ActionSystemConfigShowMetaGlobalInput) *ActionSystemConfigShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSystemConfigShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSystemConfigShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSystemConfigShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionSystemConfigShowInvocation) Call() (*ActionSystemConfigShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionSystemConfigShowInvocation) CallContext(ctx context.Context) (*ActionSystemConfigShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionSystemConfigShowResponse{Action: inv.Action}
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
		if resp.Response.SystemConfig == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.SystemConfig
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionSystemConfigShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionSystemConfigShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionSystemConfigShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
