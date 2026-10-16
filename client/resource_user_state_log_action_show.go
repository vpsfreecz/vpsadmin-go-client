package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionUserStateLogShow is a type for action User.State_log#Show
type ActionUserStateLogShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserStateLogShow(client *Client) *ActionUserStateLogShow {
	return &ActionUserStateLogShow{
		Client: client,
	}
}

// ActionUserStateLogShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserStateLogShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserStateLogShowMetaGlobalInput) SetIncludes(value string) *ActionUserStateLogShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserStateLogShowMetaGlobalInput) SetNo(value bool) *ActionUserStateLogShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserStateLogShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserStateLogShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserStateLogShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserStateLogShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserStateLogShowOutput is a type for action output parameters
type ActionUserStateLogShowOutput struct {
	ChangedAt   string                "json:\"changed_at\""
	Expiration  string                "json:\"expiration\""
	Id          int64                 "json:\"id\""
	Reason      string                "json:\"reason\""
	RemindAfter string                "json:\"remind_after\""
	State       string                "json:\"state\""
	User        *ActionUserShowOutput "json:\"user\""
}

// Type for action response, including envelope
type ActionUserStateLogShowResponse struct {
	Action *ActionUserStateLogShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		StateLog *ActionUserStateLogShowOutput "json:\"state_log\""
	}

	// Action output without the namespace
	Output *ActionUserStateLogShowOutput
}

// Prepare the action for invocation
func (action *ActionUserStateLogShow) Prepare() *ActionUserStateLogShowInvocation {
	return &ActionUserStateLogShowInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/state_logs/{state_log_id}",
	}
}

// ActionUserStateLogShowInvocation is used to configure action for invocation
type ActionUserStateLogShowInvocation struct {
	// Pointer to the action
	Action *ActionUserStateLogShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserStateLogShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserStateLogShowInvocation) SetPathParamInt(param string, value int64) *ActionUserStateLogShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserStateLogShowInvocation) SetPathParamString(param string, value string) *ActionUserStateLogShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserStateLogShowInvocation) NewMetaInput() *ActionUserStateLogShowMetaGlobalInput {
	inv.MetaInput = &ActionUserStateLogShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserStateLogShowInvocation) SetMetaInput(input *ActionUserStateLogShowMetaGlobalInput) *ActionUserStateLogShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserStateLogShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserStateLogShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionUserStateLogShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionUserStateLogShowInvocation) Call() (*ActionUserStateLogShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionUserStateLogShowInvocation) CallContext(ctx context.Context) (*ActionUserStateLogShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionUserStateLogShowResponse{Action: inv.Action}
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
		if resp.Response.StateLog == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.StateLog
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionUserStateLogShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionUserStateLogShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionUserStateLogShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
