package client

import ()

// AuthTokenActionTokenTotp is a type for action Token#Totp
type AuthTokenActionTokenTotp struct {
	// Pointer to client
	Client *Client
}

func NewAuthTokenActionTokenTotp(client *Client) *AuthTokenActionTokenTotp {
	return &AuthTokenActionTokenTotp{
		Client: client,
	}
}

// AuthTokenActionTokenTotpMetaGlobalInput is a type for action global meta input parameters
type AuthTokenActionTokenTotpMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *AuthTokenActionTokenTotpMetaGlobalInput) SetNo(value bool) *AuthTokenActionTokenTotpMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from AuthTokenActionTokenTotpMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *AuthTokenActionTokenTotpMetaGlobalInput) SelectParameters(params ...string) *AuthTokenActionTokenTotpMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *AuthTokenActionTokenTotpMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// AuthTokenActionTokenTotpInput is a type for action input parameters
type AuthTokenActionTokenTotpInput struct {
	Code  string `json:"code"`
	Token string `json:"token"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCode sets parameter Code to value and selects it for sending
func (in *AuthTokenActionTokenTotpInput) SetCode(value string) *AuthTokenActionTokenTotpInput {
	in.Code = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Code"] = nil
	return in
}

// SetToken sets parameter Token to value and selects it for sending
func (in *AuthTokenActionTokenTotpInput) SetToken(value string) *AuthTokenActionTokenTotpInput {
	in.Token = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Token"] = nil
	return in
}

// SelectParameters sets parameters from AuthTokenActionTokenTotpInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *AuthTokenActionTokenTotpInput) SelectParameters(params ...string) *AuthTokenActionTokenTotpInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from AuthTokenActionTokenTotpInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *AuthTokenActionTokenTotpInput) UnselectParameters(params ...string) *AuthTokenActionTokenTotpInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *AuthTokenActionTokenTotpInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// AuthTokenActionTokenTotpRequest is a type for the entire action request
type AuthTokenActionTokenTotpRequest struct {
	Token map[string]interface{} `json:"token"`
	Meta  map[string]interface{} `json:"_meta"`
}

// AuthTokenActionTokenTotpOutput is a type for action output parameters
type AuthTokenActionTokenTotpOutput struct {
	Complete   bool   `json:"complete"`
	NextAction string `json:"next_action"`
	Token      string `json:"token"`
	ValidTo    string `json:"valid_to"`
}

// Type for action response, including envelope
type AuthTokenActionTokenTotpResponse struct {
	Action *AuthTokenActionTokenTotp `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Token *AuthTokenActionTokenTotpOutput `json:"token"`
	}

	// Action output without the namespace
	Output *AuthTokenActionTokenTotpOutput
}

// Prepare the action for invocation
func (action *AuthTokenActionTokenTotp) Prepare() *AuthTokenActionTokenTotpInvocation {
	return &AuthTokenActionTokenTotpInvocation{
		Action: action,
		Path:   "/_auth/token/tokens/totp",
	}
}

// AuthTokenActionTokenTotpInvocation is used to configure action for invocation
type AuthTokenActionTokenTotpInvocation struct {
	// Pointer to the action
	Action *AuthTokenActionTokenTotp

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *AuthTokenActionTokenTotpInput
	// Global meta input parameters
	MetaInput *AuthTokenActionTokenTotpMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *AuthTokenActionTokenTotpInvocation) NewInput() *AuthTokenActionTokenTotpInput {
	inv.Input = &AuthTokenActionTokenTotpInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *AuthTokenActionTokenTotpInvocation) SetInput(input *AuthTokenActionTokenTotpInput) *AuthTokenActionTokenTotpInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *AuthTokenActionTokenTotpInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *AuthTokenActionTokenTotpInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *AuthTokenActionTokenTotpInvocation) NewMetaInput() *AuthTokenActionTokenTotpMetaGlobalInput {
	inv.MetaInput = &AuthTokenActionTokenTotpMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *AuthTokenActionTokenTotpInvocation) SetMetaInput(input *AuthTokenActionTokenTotpMetaGlobalInput) *AuthTokenActionTokenTotpInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *AuthTokenActionTokenTotpInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *AuthTokenActionTokenTotpInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *AuthTokenActionTokenTotpInvocation) Call() (*AuthTokenActionTokenTotpResponse, error) {
	return inv.callAsBody()
}

func (inv *AuthTokenActionTokenTotpInvocation) callAsBody() (*AuthTokenActionTokenTotpResponse, error) {
	input := inv.makeAllInputParams()
	resp := &AuthTokenActionTokenTotpResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Token
	}
	return resp, err
}

func (inv *AuthTokenActionTokenTotpInvocation) makeAllInputParams() *AuthTokenActionTokenTotpRequest {
	return &AuthTokenActionTokenTotpRequest{
		Token: inv.makeInputParams(),
		Meta:  inv.makeMetaInputParams(),
	}
}

func (inv *AuthTokenActionTokenTotpInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Code") {
			ret["code"] = inv.Input.Code
		}
		if inv.IsParameterSelected("Token") {
			ret["token"] = inv.Input.Token
		}
	}

	return ret
}

func (inv *AuthTokenActionTokenTotpInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
