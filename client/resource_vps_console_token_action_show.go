package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionVpsConsoleTokenShow is a type for action Vps.Console_token#Show
type ActionVpsConsoleTokenShow struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsConsoleTokenShow(client *Client) *ActionVpsConsoleTokenShow {
	return &ActionVpsConsoleTokenShow{
		Client: client,
	}
}

// ActionVpsConsoleTokenShowMetaGlobalInput is a type for action global meta input parameters
type ActionVpsConsoleTokenShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsConsoleTokenShowMetaGlobalInput) SetIncludes(value string) *ActionVpsConsoleTokenShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsConsoleTokenShowMetaGlobalInput) SetNo(value bool) *ActionVpsConsoleTokenShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsConsoleTokenShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsConsoleTokenShowMetaGlobalInput) SelectParameters(params ...string) *ActionVpsConsoleTokenShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsConsoleTokenShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsConsoleTokenShowOutput is a type for action output parameters
type ActionVpsConsoleTokenShowOutput struct {
	Expiration string "json:\"expiration\""
	Token      string "json:\"token\""
}

// Type for action response, including envelope
type ActionVpsConsoleTokenShowResponse struct {
	Action *ActionVpsConsoleTokenShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ConsoleToken *ActionVpsConsoleTokenShowOutput "json:\"console_token\""
	}

	// Action output without the namespace
	Output *ActionVpsConsoleTokenShowOutput
}

// Prepare the action for invocation
func (action *ActionVpsConsoleTokenShow) Prepare() *ActionVpsConsoleTokenShowInvocation {
	return &ActionVpsConsoleTokenShowInvocation{
		Action: action,
		Path:   "/v7.0/vpses/{vps_id}/console_token",
	}
}

// ActionVpsConsoleTokenShowInvocation is used to configure action for invocation
type ActionVpsConsoleTokenShowInvocation struct {
	// Pointer to the action
	Action *ActionVpsConsoleTokenShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionVpsConsoleTokenShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsConsoleTokenShowInvocation) SetPathParamInt(param string, value int64) *ActionVpsConsoleTokenShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsConsoleTokenShowInvocation) SetPathParamString(param string, value string) *ActionVpsConsoleTokenShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsConsoleTokenShowInvocation) NewMetaInput() *ActionVpsConsoleTokenShowMetaGlobalInput {
	inv.MetaInput = &ActionVpsConsoleTokenShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsConsoleTokenShowInvocation) SetMetaInput(input *ActionVpsConsoleTokenShowMetaGlobalInput) *ActionVpsConsoleTokenShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsConsoleTokenShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsConsoleTokenShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionVpsConsoleTokenShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionVpsConsoleTokenShowInvocation) Call() (*ActionVpsConsoleTokenShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionVpsConsoleTokenShowInvocation) CallContext(ctx context.Context) (*ActionVpsConsoleTokenShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionVpsConsoleTokenShowResponse{Action: inv.Action}
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
		if resp.Response.ConsoleToken == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.ConsoleToken
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionVpsConsoleTokenShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionVpsConsoleTokenShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionVpsConsoleTokenShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
