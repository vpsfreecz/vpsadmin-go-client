package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionNewsLogDelete is a type for action News_log#Delete
type ActionNewsLogDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionNewsLogDelete(client *Client) *ActionNewsLogDelete {
	return &ActionNewsLogDelete{
		Client: client,
	}
}

// ActionNewsLogDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionNewsLogDeleteMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNewsLogDeleteMetaGlobalInput) SetIncludes(value string) *ActionNewsLogDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNewsLogDeleteMetaGlobalInput) SetNo(value bool) *ActionNewsLogDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNewsLogDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNewsLogDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionNewsLogDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNewsLogDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNewsLogDeleteRequest is a type for the entire action request
type ActionNewsLogDeleteRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// Type for action response, including envelope
type ActionNewsLogDeleteResponse struct {
	Action *ActionNewsLogDelete "json:\"-\""
	*Envelope
}

// Prepare the action for invocation
func (action *ActionNewsLogDelete) Prepare() *ActionNewsLogDeleteInvocation {
	return &ActionNewsLogDeleteInvocation{
		Action: action,
		Path:   "/v7.0/news_logs/{news_log_id}",
	}
}

// ActionNewsLogDeleteInvocation is used to configure action for invocation
type ActionNewsLogDeleteInvocation struct {
	// Pointer to the action
	Action *ActionNewsLogDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNewsLogDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNewsLogDeleteInvocation) SetPathParamInt(param string, value int64) *ActionNewsLogDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNewsLogDeleteInvocation) SetPathParamString(param string, value string) *ActionNewsLogDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNewsLogDeleteInvocation) NewMetaInput() *ActionNewsLogDeleteMetaGlobalInput {
	inv.MetaInput = &ActionNewsLogDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNewsLogDeleteInvocation) SetMetaInput(input *ActionNewsLogDeleteMetaGlobalInput) *ActionNewsLogDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNewsLogDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNewsLogDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNewsLogDeleteInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionNewsLogDeleteInvocation) Call() (*ActionNewsLogDeleteResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionNewsLogDeleteInvocation) CallContext(ctx context.Context) (*ActionNewsLogDeleteResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionNewsLogDeleteResponse{Action: inv.Action}
	raw, err := inv.Action.Client.doRaw(req)
	if err != nil {
		return resp, err
	}
	if err := decodeTypedResponse(raw.Body, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionNewsLogDeleteInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionNewsLogDeleteInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	input := inv.makeAllInputParams()
	return inv.Action.Client.newBodyRequest(ctx, "DELETE", inv.Path, input)
}

func (inv *ActionNewsLogDeleteInvocation) makeAllInputParams() *ActionNewsLogDeleteRequest {
	return &ActionNewsLogDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionNewsLogDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
