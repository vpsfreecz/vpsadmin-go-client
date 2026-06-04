package client

import ()

// AuthTokenActionTokenResetPassword is a type for action Token#Reset_password
type AuthTokenActionTokenResetPassword struct {
	// Pointer to client
	Client *Client
}

func NewAuthTokenActionTokenResetPassword(client *Client) *AuthTokenActionTokenResetPassword {
	return &AuthTokenActionTokenResetPassword{
		Client: client,
	}
}

// AuthTokenActionTokenResetPasswordMetaGlobalInput is a type for action global meta input parameters
type AuthTokenActionTokenResetPasswordMetaGlobalInput struct {
	No bool "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *AuthTokenActionTokenResetPasswordMetaGlobalInput) SetNo(value bool) *AuthTokenActionTokenResetPasswordMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from AuthTokenActionTokenResetPasswordMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *AuthTokenActionTokenResetPasswordMetaGlobalInput) SelectParameters(params ...string) *AuthTokenActionTokenResetPasswordMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *AuthTokenActionTokenResetPasswordMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// AuthTokenActionTokenResetPasswordInput is a type for action input parameters
type AuthTokenActionTokenResetPasswordInput struct {
	NewPassword1 string "json:\"new_password1\""
	NewPassword2 string "json:\"new_password2\""
	Token        string "json:\"token\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNewPassword1 sets parameter NewPassword1 to value and selects it for sending
func (in *AuthTokenActionTokenResetPasswordInput) SetNewPassword1(value string) *AuthTokenActionTokenResetPasswordInput {
	in.NewPassword1 = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NewPassword1"] = nil
	return in
}

// SetNewPassword2 sets parameter NewPassword2 to value and selects it for sending
func (in *AuthTokenActionTokenResetPasswordInput) SetNewPassword2(value string) *AuthTokenActionTokenResetPasswordInput {
	in.NewPassword2 = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NewPassword2"] = nil
	return in
}

// SetToken sets parameter Token to value and selects it for sending
func (in *AuthTokenActionTokenResetPasswordInput) SetToken(value string) *AuthTokenActionTokenResetPasswordInput {
	in.Token = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Token"] = nil
	return in
}

// SelectParameters sets parameters from AuthTokenActionTokenResetPasswordInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *AuthTokenActionTokenResetPasswordInput) SelectParameters(params ...string) *AuthTokenActionTokenResetPasswordInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from AuthTokenActionTokenResetPasswordInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *AuthTokenActionTokenResetPasswordInput) UnselectParameters(params ...string) *AuthTokenActionTokenResetPasswordInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *AuthTokenActionTokenResetPasswordInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// AuthTokenActionTokenResetPasswordRequest is a type for the entire action request
type AuthTokenActionTokenResetPasswordRequest struct {
	Token map[string]interface{} "json:\"token\""
	Meta  map[string]interface{} "json:\"_meta\""
}

// AuthTokenActionTokenResetPasswordOutput is a type for action output parameters
type AuthTokenActionTokenResetPasswordOutput struct {
	Complete   bool   "json:\"complete\""
	NextAction string "json:\"next_action\""
	Token      string "json:\"token\""
	ValidTo    string "json:\"valid_to\""
}

// Type for action response, including envelope
type AuthTokenActionTokenResetPasswordResponse struct {
	Action *AuthTokenActionTokenResetPassword "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Token *AuthTokenActionTokenResetPasswordOutput "json:\"token\""
	}

	// Action output without the namespace
	Output *AuthTokenActionTokenResetPasswordOutput
}

// Prepare the action for invocation
func (action *AuthTokenActionTokenResetPassword) Prepare() *AuthTokenActionTokenResetPasswordInvocation {
	return &AuthTokenActionTokenResetPasswordInvocation{
		Action: action,
		Path:   "/_auth/token/tokens/reset_password",
	}
}

// AuthTokenActionTokenResetPasswordInvocation is used to configure action for invocation
type AuthTokenActionTokenResetPasswordInvocation struct {
	// Pointer to the action
	Action *AuthTokenActionTokenResetPassword

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *AuthTokenActionTokenResetPasswordInput
	// Global meta input parameters
	MetaInput *AuthTokenActionTokenResetPasswordMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *AuthTokenActionTokenResetPasswordInvocation) NewInput() *AuthTokenActionTokenResetPasswordInput {
	inv.Input = &AuthTokenActionTokenResetPasswordInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *AuthTokenActionTokenResetPasswordInvocation) SetInput(input *AuthTokenActionTokenResetPasswordInput) *AuthTokenActionTokenResetPasswordInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *AuthTokenActionTokenResetPasswordInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *AuthTokenActionTokenResetPasswordInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *AuthTokenActionTokenResetPasswordInvocation) NewMetaInput() *AuthTokenActionTokenResetPasswordMetaGlobalInput {
	inv.MetaInput = &AuthTokenActionTokenResetPasswordMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *AuthTokenActionTokenResetPasswordInvocation) SetMetaInput(input *AuthTokenActionTokenResetPasswordMetaGlobalInput) *AuthTokenActionTokenResetPasswordInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *AuthTokenActionTokenResetPasswordInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *AuthTokenActionTokenResetPasswordInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *AuthTokenActionTokenResetPasswordInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *AuthTokenActionTokenResetPasswordInvocation) Call() (*AuthTokenActionTokenResetPasswordResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *AuthTokenActionTokenResetPasswordInvocation) callAsBody() (*AuthTokenActionTokenResetPasswordResponse, error) {
	input := inv.makeAllInputParams()
	resp := &AuthTokenActionTokenResetPasswordResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Token
	}
	return resp, err
}

func (inv *AuthTokenActionTokenResetPasswordInvocation) makeAllInputParams() *AuthTokenActionTokenResetPasswordRequest {
	return &AuthTokenActionTokenResetPasswordRequest{
		Token: inv.makeInputParams(),
		Meta:  inv.makeMetaInputParams(),
	}
}

func (inv *AuthTokenActionTokenResetPasswordInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("NewPassword1") {
			ret["new_password1"] = inv.Input.NewPassword1
		}
		if inv.IsParameterSelected("NewPassword2") {
			ret["new_password2"] = inv.Input.NewPassword2
		}
		if inv.IsParameterSelected("Token") {
			ret["token"] = inv.Input.Token
		}
	}

	return ret
}

func (inv *AuthTokenActionTokenResetPasswordInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
