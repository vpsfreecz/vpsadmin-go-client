package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionPasswordChangeLogShow is a type for action Password_change_log#Show
type ActionPasswordChangeLogShow struct {
	// Pointer to client
	Client *Client
}

func NewActionPasswordChangeLogShow(client *Client) *ActionPasswordChangeLogShow {
	return &ActionPasswordChangeLogShow{
		Client: client,
	}
}

// ActionPasswordChangeLogShowMetaGlobalInput is a type for action global meta input parameters
type ActionPasswordChangeLogShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionPasswordChangeLogShowMetaGlobalInput) SetIncludes(value string) *ActionPasswordChangeLogShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionPasswordChangeLogShowMetaGlobalInput) SetNo(value bool) *ActionPasswordChangeLogShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionPasswordChangeLogShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionPasswordChangeLogShowMetaGlobalInput) SelectParameters(params ...string) *ActionPasswordChangeLogShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionPasswordChangeLogShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionPasswordChangeLogShowOutput is a type for action output parameters
type ActionPasswordChangeLogShowOutput struct {
	ClientIpAddr           string                       "json:\"client_ip_addr\""
	ClientIpPtr            string                       "json:\"client_ip_ptr\""
	CreatedAt              string                       "json:\"created_at\""
	Id                     int64                        "json:\"id\""
	Source                 string                       "json:\"source\""
	User                   *ActionUserShowOutput        "json:\"user\""
	UserAgent              string                       "json:\"user_agent\""
	UserSession            *ActionUserSessionShowOutput "json:\"user_session\""
	UserSessionId          int64                        "json:\"user_session_id\""
	UserSessionOwnedByUser bool                         "json:\"user_session_owned_by_user\""
}

// Type for action response, including envelope
type ActionPasswordChangeLogShowResponse struct {
	Action *ActionPasswordChangeLogShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		PasswordChangeLog *ActionPasswordChangeLogShowOutput "json:\"password_change_log\""
	}

	// Action output without the namespace
	Output *ActionPasswordChangeLogShowOutput
}

// Prepare the action for invocation
func (action *ActionPasswordChangeLogShow) Prepare() *ActionPasswordChangeLogShowInvocation {
	return &ActionPasswordChangeLogShowInvocation{
		Action: action,
		Path:   "/v7.0/password_change_logs/{password_change_log_id}",
	}
}

// ActionPasswordChangeLogShowInvocation is used to configure action for invocation
type ActionPasswordChangeLogShowInvocation struct {
	// Pointer to the action
	Action *ActionPasswordChangeLogShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionPasswordChangeLogShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionPasswordChangeLogShowInvocation) SetPathParamInt(param string, value int64) *ActionPasswordChangeLogShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionPasswordChangeLogShowInvocation) SetPathParamString(param string, value string) *ActionPasswordChangeLogShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionPasswordChangeLogShowInvocation) NewMetaInput() *ActionPasswordChangeLogShowMetaGlobalInput {
	inv.MetaInput = &ActionPasswordChangeLogShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionPasswordChangeLogShowInvocation) SetMetaInput(input *ActionPasswordChangeLogShowMetaGlobalInput) *ActionPasswordChangeLogShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionPasswordChangeLogShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionPasswordChangeLogShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionPasswordChangeLogShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionPasswordChangeLogShowInvocation) Call() (*ActionPasswordChangeLogShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionPasswordChangeLogShowInvocation) CallContext(ctx context.Context) (*ActionPasswordChangeLogShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionPasswordChangeLogShowResponse{Action: inv.Action}
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
		if resp.Response.PasswordChangeLog == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.PasswordChangeLog
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionPasswordChangeLogShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionPasswordChangeLogShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionPasswordChangeLogShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
