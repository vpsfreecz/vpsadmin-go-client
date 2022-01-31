package client

import ()

// AuthTokenActionTokenRequest is a type for action Token#Request
type AuthTokenActionTokenRequest struct {
	// Pointer to client
	Client *Client
}

func NewAuthTokenActionTokenRequest(client *Client) *AuthTokenActionTokenRequest {
	return &AuthTokenActionTokenRequest{
		Client: client,
	}
}

// AuthTokenActionTokenRequestMetaGlobalInput is a type for action global meta input parameters
type AuthTokenActionTokenRequestMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *AuthTokenActionTokenRequestMetaGlobalInput) SetNo(value bool) *AuthTokenActionTokenRequestMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from AuthTokenActionTokenRequestMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *AuthTokenActionTokenRequestMetaGlobalInput) SelectParameters(params ...string) *AuthTokenActionTokenRequestMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *AuthTokenActionTokenRequestMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// AuthTokenActionTokenRequestInput is a type for action input parameters
type AuthTokenActionTokenRequestInput struct {
	Interval int64  `json:"interval"`
	Lifetime string `json:"lifetime"`
	Password string `json:"password"`
	User     string `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetInterval sets parameter Interval to value and selects it for sending
func (in *AuthTokenActionTokenRequestInput) SetInterval(value int64) *AuthTokenActionTokenRequestInput {
	in.Interval = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Interval"] = nil
	return in
}

// SetLifetime sets parameter Lifetime to value and selects it for sending
func (in *AuthTokenActionTokenRequestInput) SetLifetime(value string) *AuthTokenActionTokenRequestInput {
	in.Lifetime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Lifetime"] = nil
	return in
}

// SetPassword sets parameter Password to value and selects it for sending
func (in *AuthTokenActionTokenRequestInput) SetPassword(value string) *AuthTokenActionTokenRequestInput {
	in.Password = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Password"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *AuthTokenActionTokenRequestInput) SetUser(value string) *AuthTokenActionTokenRequestInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from AuthTokenActionTokenRequestInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *AuthTokenActionTokenRequestInput) SelectParameters(params ...string) *AuthTokenActionTokenRequestInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from AuthTokenActionTokenRequestInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *AuthTokenActionTokenRequestInput) UnselectParameters(params ...string) *AuthTokenActionTokenRequestInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *AuthTokenActionTokenRequestInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// AuthTokenActionTokenRequestRequest is a type for the entire action request
type AuthTokenActionTokenRequestRequest struct {
	Token map[string]interface{} `json:"token"`
	Meta  map[string]interface{} `json:"_meta"`
}

// AuthTokenActionTokenRequestOutput is a type for action output parameters
type AuthTokenActionTokenRequestOutput struct {
	Complete   bool   `json:"complete"`
	NextAction string `json:"next_action"`
	Token      string `json:"token"`
	ValidTo    string `json:"valid_to"`
}

// Type for action response, including envelope
type AuthTokenActionTokenRequestResponse struct {
	Action *AuthTokenActionTokenRequest `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Token *AuthTokenActionTokenRequestOutput `json:"token"`
	}

	// Action output without the namespace
	Output *AuthTokenActionTokenRequestOutput
}

// Prepare the action for invocation
func (action *AuthTokenActionTokenRequest) Prepare() *AuthTokenActionTokenRequestInvocation {
	return &AuthTokenActionTokenRequestInvocation{
		Action: action,
		Path:   "/_auth/token/tokens",
	}
}

// AuthTokenActionTokenRequestInvocation is used to configure action for invocation
type AuthTokenActionTokenRequestInvocation struct {
	// Pointer to the action
	Action *AuthTokenActionTokenRequest

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *AuthTokenActionTokenRequestInput
	// Global meta input parameters
	MetaInput *AuthTokenActionTokenRequestMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *AuthTokenActionTokenRequestInvocation) NewInput() *AuthTokenActionTokenRequestInput {
	inv.Input = &AuthTokenActionTokenRequestInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *AuthTokenActionTokenRequestInvocation) SetInput(input *AuthTokenActionTokenRequestInput) *AuthTokenActionTokenRequestInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *AuthTokenActionTokenRequestInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *AuthTokenActionTokenRequestInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *AuthTokenActionTokenRequestInvocation) NewMetaInput() *AuthTokenActionTokenRequestMetaGlobalInput {
	inv.MetaInput = &AuthTokenActionTokenRequestMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *AuthTokenActionTokenRequestInvocation) SetMetaInput(input *AuthTokenActionTokenRequestMetaGlobalInput) *AuthTokenActionTokenRequestInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *AuthTokenActionTokenRequestInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *AuthTokenActionTokenRequestInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *AuthTokenActionTokenRequestInvocation) Call() (*AuthTokenActionTokenRequestResponse, error) {
	return inv.callAsBody()
}

func (inv *AuthTokenActionTokenRequestInvocation) callAsBody() (*AuthTokenActionTokenRequestResponse, error) {
	input := inv.makeAllInputParams()
	resp := &AuthTokenActionTokenRequestResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Token
	}
	return resp, err
}

func (inv *AuthTokenActionTokenRequestInvocation) makeAllInputParams() *AuthTokenActionTokenRequestRequest {
	return &AuthTokenActionTokenRequestRequest{
		Token: inv.makeInputParams(),
		Meta:  inv.makeMetaInputParams(),
	}
}

func (inv *AuthTokenActionTokenRequestInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Interval") {
			ret["interval"] = inv.Input.Interval
		}
		if inv.IsParameterSelected("Lifetime") {
			ret["lifetime"] = inv.Input.Lifetime
		}
		if inv.IsParameterSelected("Password") {
			ret["password"] = inv.Input.Password
		}
		if inv.IsParameterSelected("User") {
			ret["user"] = inv.Input.User
		}
	}

	return ret
}

func (inv *AuthTokenActionTokenRequestInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
