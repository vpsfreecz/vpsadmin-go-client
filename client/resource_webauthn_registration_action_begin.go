package client

import (
	"context"
	"net/http"
)

// ActionWebauthnRegistrationBegin is a type for action Webauthn.Registration#Begin
type ActionWebauthnRegistrationBegin struct {
	// Pointer to client
	Client *Client
}

func NewActionWebauthnRegistrationBegin(client *Client) *ActionWebauthnRegistrationBegin {
	return &ActionWebauthnRegistrationBegin{
		Client: client,
	}
}

// ActionWebauthnRegistrationBeginMetaGlobalInput is a type for action global meta input parameters
type ActionWebauthnRegistrationBeginMetaGlobalInput struct {
	No bool "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionWebauthnRegistrationBeginMetaGlobalInput) SetNo(value bool) *ActionWebauthnRegistrationBeginMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionWebauthnRegistrationBeginMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionWebauthnRegistrationBeginMetaGlobalInput) SelectParameters(params ...string) *ActionWebauthnRegistrationBeginMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionWebauthnRegistrationBeginMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionWebauthnRegistrationBeginRequest is a type for the entire action request
type ActionWebauthnRegistrationBeginRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// ActionWebauthnRegistrationBeginOutput is a type for action output parameters
type ActionWebauthnRegistrationBeginOutput struct {
	ChallengeToken string "json:\"challenge_token\""
}

// Type for action response, including envelope
type ActionWebauthnRegistrationBeginResponse struct {
	Action *ActionWebauthnRegistrationBegin "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Registration *ActionWebauthnRegistrationBeginOutput "json:\"registration\""
	}

	// Action output without the namespace
	Output *ActionWebauthnRegistrationBeginOutput
}

// Call the action directly without any path or input parameters
func (action *ActionWebauthnRegistrationBegin) Call() (*ActionWebauthnRegistrationBeginResponse, error) {
	return action.CallContext(context.Background())
}

// CallContext calls the action with cancellation and deadline support.
func (action *ActionWebauthnRegistrationBegin) CallContext(ctx context.Context) (*ActionWebauthnRegistrationBeginResponse, error) {
	return action.Prepare().CallContext(ctx)
}

// CallRaw calls the action without decoding or classifying its response.
func (action *ActionWebauthnRegistrationBegin) CallRaw(ctx context.Context) (*RawResponse, error) {
	return action.Prepare().CallRaw(ctx)
}

// Prepare the action for invocation
func (action *ActionWebauthnRegistrationBegin) Prepare() *ActionWebauthnRegistrationBeginInvocation {
	return &ActionWebauthnRegistrationBeginInvocation{
		Action: action,
		Path:   "/v7.0/webauthn/registration/begin",
	}
}

// ActionWebauthnRegistrationBeginInvocation is used to configure action for invocation
type ActionWebauthnRegistrationBeginInvocation struct {
	// Pointer to the action
	Action *ActionWebauthnRegistrationBegin

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionWebauthnRegistrationBeginMetaGlobalInput
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionWebauthnRegistrationBeginInvocation) NewMetaInput() *ActionWebauthnRegistrationBeginMetaGlobalInput {
	inv.MetaInput = &ActionWebauthnRegistrationBeginMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionWebauthnRegistrationBeginInvocation) SetMetaInput(input *ActionWebauthnRegistrationBeginMetaGlobalInput) *ActionWebauthnRegistrationBeginInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionWebauthnRegistrationBeginInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionWebauthnRegistrationBeginInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionWebauthnRegistrationBeginInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionWebauthnRegistrationBeginInvocation) Call() (*ActionWebauthnRegistrationBeginResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionWebauthnRegistrationBeginInvocation) CallContext(ctx context.Context) (*ActionWebauthnRegistrationBeginResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionWebauthnRegistrationBeginResponse{Action: inv.Action}
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
		if resp.Response.Registration == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.Registration
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionWebauthnRegistrationBeginInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionWebauthnRegistrationBeginInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	input := inv.makeAllInputParams()
	return inv.Action.Client.newBodyRequest(ctx, "POST", inv.Path, input)
}

func (inv *ActionWebauthnRegistrationBeginInvocation) makeAllInputParams() *ActionWebauthnRegistrationBeginRequest {
	return &ActionWebauthnRegistrationBeginRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionWebauthnRegistrationBeginInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
