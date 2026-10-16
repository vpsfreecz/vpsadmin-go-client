package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionUserWebauthnCredentialShow is a type for action User.Webauthn_credential#Show
type ActionUserWebauthnCredentialShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserWebauthnCredentialShow(client *Client) *ActionUserWebauthnCredentialShow {
	return &ActionUserWebauthnCredentialShow{
		Client: client,
	}
}

// ActionUserWebauthnCredentialShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserWebauthnCredentialShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserWebauthnCredentialShowMetaGlobalInput) SetIncludes(value string) *ActionUserWebauthnCredentialShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserWebauthnCredentialShowMetaGlobalInput) SetNo(value bool) *ActionUserWebauthnCredentialShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserWebauthnCredentialShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserWebauthnCredentialShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserWebauthnCredentialShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserWebauthnCredentialShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserWebauthnCredentialShowOutput is a type for action output parameters
type ActionUserWebauthnCredentialShowOutput struct {
	CreatedAt string "json:\"created_at\""
	Enabled   bool   "json:\"enabled\""
	Id        int64  "json:\"id\""
	Label     string "json:\"label\""
	LastUseAt string "json:\"last_use_at\""
	UpdatedAt string "json:\"updated_at\""
	UseCount  int64  "json:\"use_count\""
}

// Type for action response, including envelope
type ActionUserWebauthnCredentialShowResponse struct {
	Action *ActionUserWebauthnCredentialShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		WebauthnCredential *ActionUserWebauthnCredentialShowOutput "json:\"webauthn_credential\""
	}

	// Action output without the namespace
	Output *ActionUserWebauthnCredentialShowOutput
}

// Prepare the action for invocation
func (action *ActionUserWebauthnCredentialShow) Prepare() *ActionUserWebauthnCredentialShowInvocation {
	return &ActionUserWebauthnCredentialShowInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/webauthn_credentials/{webauthn_credential_id}",
	}
}

// ActionUserWebauthnCredentialShowInvocation is used to configure action for invocation
type ActionUserWebauthnCredentialShowInvocation struct {
	// Pointer to the action
	Action *ActionUserWebauthnCredentialShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserWebauthnCredentialShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserWebauthnCredentialShowInvocation) SetPathParamInt(param string, value int64) *ActionUserWebauthnCredentialShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserWebauthnCredentialShowInvocation) SetPathParamString(param string, value string) *ActionUserWebauthnCredentialShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserWebauthnCredentialShowInvocation) NewMetaInput() *ActionUserWebauthnCredentialShowMetaGlobalInput {
	inv.MetaInput = &ActionUserWebauthnCredentialShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserWebauthnCredentialShowInvocation) SetMetaInput(input *ActionUserWebauthnCredentialShowMetaGlobalInput) *ActionUserWebauthnCredentialShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserWebauthnCredentialShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserWebauthnCredentialShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionUserWebauthnCredentialShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionUserWebauthnCredentialShowInvocation) Call() (*ActionUserWebauthnCredentialShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionUserWebauthnCredentialShowInvocation) CallContext(ctx context.Context) (*ActionUserWebauthnCredentialShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionUserWebauthnCredentialShowResponse{Action: inv.Action}
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
		if resp.Response.WebauthnCredential == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.WebauthnCredential
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionUserWebauthnCredentialShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionUserWebauthnCredentialShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionUserWebauthnCredentialShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
