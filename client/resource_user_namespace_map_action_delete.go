package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionUserNamespaceMapDelete is a type for action User_namespace_map#Delete
type ActionUserNamespaceMapDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionUserNamespaceMapDelete(client *Client) *ActionUserNamespaceMapDelete {
	return &ActionUserNamespaceMapDelete{
		Client: client,
	}
}

// ActionUserNamespaceMapDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionUserNamespaceMapDeleteMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserNamespaceMapDeleteMetaGlobalInput) SetIncludes(value string) *ActionUserNamespaceMapDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserNamespaceMapDeleteMetaGlobalInput) SetNo(value bool) *ActionUserNamespaceMapDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNamespaceMapDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceMapDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionUserNamespaceMapDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNamespaceMapDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNamespaceMapDeleteRequest is a type for the entire action request
type ActionUserNamespaceMapDeleteRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// Type for action response, including envelope
type ActionUserNamespaceMapDeleteResponse struct {
	Action *ActionUserNamespaceMapDelete "json:\"-\""
	*Envelope
}

// Prepare the action for invocation
func (action *ActionUserNamespaceMapDelete) Prepare() *ActionUserNamespaceMapDeleteInvocation {
	return &ActionUserNamespaceMapDeleteInvocation{
		Action: action,
		Path:   "/v7.0/user_namespace_maps/{user_namespace_map_id}",
	}
}

// ActionUserNamespaceMapDeleteInvocation is used to configure action for invocation
type ActionUserNamespaceMapDeleteInvocation struct {
	// Pointer to the action
	Action *ActionUserNamespaceMapDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserNamespaceMapDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserNamespaceMapDeleteInvocation) SetPathParamInt(param string, value int64) *ActionUserNamespaceMapDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserNamespaceMapDeleteInvocation) SetPathParamString(param string, value string) *ActionUserNamespaceMapDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserNamespaceMapDeleteInvocation) NewMetaInput() *ActionUserNamespaceMapDeleteMetaGlobalInput {
	inv.MetaInput = &ActionUserNamespaceMapDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserNamespaceMapDeleteInvocation) SetMetaInput(input *ActionUserNamespaceMapDeleteMetaGlobalInput) *ActionUserNamespaceMapDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserNamespaceMapDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserNamespaceMapDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionUserNamespaceMapDeleteInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionUserNamespaceMapDeleteInvocation) Call() (*ActionUserNamespaceMapDeleteResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionUserNamespaceMapDeleteInvocation) CallContext(ctx context.Context) (*ActionUserNamespaceMapDeleteResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionUserNamespaceMapDeleteResponse{Action: inv.Action}
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
func (inv *ActionUserNamespaceMapDeleteInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionUserNamespaceMapDeleteInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	input := inv.makeAllInputParams()
	return inv.Action.Client.newBodyRequest(ctx, "DELETE", inv.Path, input)
}

func (inv *ActionUserNamespaceMapDeleteInvocation) makeAllInputParams() *ActionUserNamespaceMapDeleteRequest {
	return &ActionUserNamespaceMapDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserNamespaceMapDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
