package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionNewsLogShow is a type for action News_log#Show
type ActionNewsLogShow struct {
	// Pointer to client
	Client *Client
}

func NewActionNewsLogShow(client *Client) *ActionNewsLogShow {
	return &ActionNewsLogShow{
		Client: client,
	}
}

// ActionNewsLogShowMetaGlobalInput is a type for action global meta input parameters
type ActionNewsLogShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNewsLogShowMetaGlobalInput) SetIncludes(value string) *ActionNewsLogShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNewsLogShowMetaGlobalInput) SetNo(value bool) *ActionNewsLogShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNewsLogShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNewsLogShowMetaGlobalInput) SelectParameters(params ...string) *ActionNewsLogShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNewsLogShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNewsLogShowOutput is a type for action output parameters
type ActionNewsLogShowOutput struct {
	CreatedAt   string "json:\"created_at\""
	CsMessage   string "json:\"cs_message\""
	EnMessage   string "json:\"en_message\""
	Id          int64  "json:\"id\""
	Message     string "json:\"message\""
	PublishedAt string "json:\"published_at\""
	UpdatedAt   string "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionNewsLogShowResponse struct {
	Action *ActionNewsLogShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NewsLog *ActionNewsLogShowOutput "json:\"news_log\""
	}

	// Action output without the namespace
	Output *ActionNewsLogShowOutput
}

// Prepare the action for invocation
func (action *ActionNewsLogShow) Prepare() *ActionNewsLogShowInvocation {
	return &ActionNewsLogShowInvocation{
		Action: action,
		Path:   "/v7.0/news_logs/{news_log_id}",
	}
}

// ActionNewsLogShowInvocation is used to configure action for invocation
type ActionNewsLogShowInvocation struct {
	// Pointer to the action
	Action *ActionNewsLogShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNewsLogShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNewsLogShowInvocation) SetPathParamInt(param string, value int64) *ActionNewsLogShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNewsLogShowInvocation) SetPathParamString(param string, value string) *ActionNewsLogShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNewsLogShowInvocation) NewMetaInput() *ActionNewsLogShowMetaGlobalInput {
	inv.MetaInput = &ActionNewsLogShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNewsLogShowInvocation) SetMetaInput(input *ActionNewsLogShowMetaGlobalInput) *ActionNewsLogShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNewsLogShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNewsLogShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNewsLogShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionNewsLogShowInvocation) Call() (*ActionNewsLogShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionNewsLogShowInvocation) CallContext(ctx context.Context) (*ActionNewsLogShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionNewsLogShowResponse{Action: inv.Action}
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
		if resp.Response.NewsLog == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.NewsLog
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionNewsLogShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionNewsLogShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionNewsLogShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
