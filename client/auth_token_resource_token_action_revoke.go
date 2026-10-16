package client

import (
	"context"
	"net/http"
)

// AuthTokenActionTokenRevoke is a type for action Token#Revoke
type AuthTokenActionTokenRevoke struct {
	// Pointer to client
	Client *Client
}

func NewAuthTokenActionTokenRevoke(client *Client) *AuthTokenActionTokenRevoke {
	return &AuthTokenActionTokenRevoke{
		Client: client,
	}
}

// AuthTokenActionTokenRevokeMetaGlobalInput is a type for action global meta input parameters
type AuthTokenActionTokenRevokeMetaGlobalInput struct {
	No bool "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *AuthTokenActionTokenRevokeMetaGlobalInput) SetNo(value bool) *AuthTokenActionTokenRevokeMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from AuthTokenActionTokenRevokeMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *AuthTokenActionTokenRevokeMetaGlobalInput) SelectParameters(params ...string) *AuthTokenActionTokenRevokeMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *AuthTokenActionTokenRevokeMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// AuthTokenActionTokenRevokeRequest is a type for the entire action request
type AuthTokenActionTokenRevokeRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// Type for action response, including envelope
type AuthTokenActionTokenRevokeResponse struct {
	Action *AuthTokenActionTokenRevoke "json:\"-\""
	*Envelope
}

// Call the action directly without any path or input parameters
func (action *AuthTokenActionTokenRevoke) Call() (*AuthTokenActionTokenRevokeResponse, error) {
	return action.CallContext(context.Background())
}

// CallContext calls the action with cancellation and deadline support.
func (action *AuthTokenActionTokenRevoke) CallContext(ctx context.Context) (*AuthTokenActionTokenRevokeResponse, error) {
	return action.Prepare().CallContext(ctx)
}

// CallRaw calls the action without decoding or classifying its response.
func (action *AuthTokenActionTokenRevoke) CallRaw(ctx context.Context) (*RawResponse, error) {
	return action.Prepare().CallRaw(ctx)
}

// Prepare the action for invocation
func (action *AuthTokenActionTokenRevoke) Prepare() *AuthTokenActionTokenRevokeInvocation {
	return &AuthTokenActionTokenRevokeInvocation{
		Action: action,
		Path:   "/_auth/token/tokens/revoke",
	}
}

// AuthTokenActionTokenRevokeInvocation is used to configure action for invocation
type AuthTokenActionTokenRevokeInvocation struct {
	// Pointer to the action
	Action *AuthTokenActionTokenRevoke

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *AuthTokenActionTokenRevokeMetaGlobalInput
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *AuthTokenActionTokenRevokeInvocation) NewMetaInput() *AuthTokenActionTokenRevokeMetaGlobalInput {
	inv.MetaInput = &AuthTokenActionTokenRevokeMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *AuthTokenActionTokenRevokeInvocation) SetMetaInput(input *AuthTokenActionTokenRevokeMetaGlobalInput) *AuthTokenActionTokenRevokeInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *AuthTokenActionTokenRevokeInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *AuthTokenActionTokenRevokeInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *AuthTokenActionTokenRevokeInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *AuthTokenActionTokenRevokeInvocation) Call() (*AuthTokenActionTokenRevokeResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *AuthTokenActionTokenRevokeInvocation) CallContext(ctx context.Context) (*AuthTokenActionTokenRevokeResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &AuthTokenActionTokenRevokeResponse{Action: inv.Action}
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
func (inv *AuthTokenActionTokenRevokeInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *AuthTokenActionTokenRevokeInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	input := inv.makeAllInputParams()
	return inv.Action.Client.newBodyRequest(ctx, "POST", inv.Path, input)
}

func (inv *AuthTokenActionTokenRevokeInvocation) makeAllInputParams() *AuthTokenActionTokenRevokeRequest {
	return &AuthTokenActionTokenRevokeRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *AuthTokenActionTokenRevokeInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
