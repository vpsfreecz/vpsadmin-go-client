package client

import ()

// ActionWebauthnAuthenticationFinish is a type for action Webauthn.Authentication#Finish
type ActionWebauthnAuthenticationFinish struct {
	// Pointer to client
	Client *Client
}

func NewActionWebauthnAuthenticationFinish(client *Client) *ActionWebauthnAuthenticationFinish {
	return &ActionWebauthnAuthenticationFinish{
		Client: client,
	}
}

// ActionWebauthnAuthenticationFinishMetaGlobalInput is a type for action global meta input parameters
type ActionWebauthnAuthenticationFinishMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionWebauthnAuthenticationFinishMetaGlobalInput) SetNo(value bool) *ActionWebauthnAuthenticationFinishMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionWebauthnAuthenticationFinishMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionWebauthnAuthenticationFinishMetaGlobalInput) SelectParameters(params ...string) *ActionWebauthnAuthenticationFinishMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionWebauthnAuthenticationFinishMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionWebauthnAuthenticationFinishInput is a type for action input parameters
type ActionWebauthnAuthenticationFinishInput struct {
	AuthToken      string `json:"auth_token"`
	ChallengeToken string `json:"challenge_token"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAuthToken sets parameter AuthToken to value and selects it for sending
func (in *ActionWebauthnAuthenticationFinishInput) SetAuthToken(value string) *ActionWebauthnAuthenticationFinishInput {
	in.AuthToken = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AuthToken"] = nil
	return in
}

// SetChallengeToken sets parameter ChallengeToken to value and selects it for sending
func (in *ActionWebauthnAuthenticationFinishInput) SetChallengeToken(value string) *ActionWebauthnAuthenticationFinishInput {
	in.ChallengeToken = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ChallengeToken"] = nil
	return in
}

// SelectParameters sets parameters from ActionWebauthnAuthenticationFinishInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionWebauthnAuthenticationFinishInput) SelectParameters(params ...string) *ActionWebauthnAuthenticationFinishInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionWebauthnAuthenticationFinishInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionWebauthnAuthenticationFinishInput) UnselectParameters(params ...string) *ActionWebauthnAuthenticationFinishInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionWebauthnAuthenticationFinishInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionWebauthnAuthenticationFinishRequest is a type for the entire action request
type ActionWebauthnAuthenticationFinishRequest struct {
	Authentication map[string]interface{} `json:"authentication"`
	Meta           map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionWebauthnAuthenticationFinishResponse struct {
	Action *ActionWebauthnAuthenticationFinish `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionWebauthnAuthenticationFinish) Prepare() *ActionWebauthnAuthenticationFinishInvocation {
	return &ActionWebauthnAuthenticationFinishInvocation{
		Action: action,
		Path:   "/v7.0/webauthn/authentication/finish",
	}
}

// ActionWebauthnAuthenticationFinishInvocation is used to configure action for invocation
type ActionWebauthnAuthenticationFinishInvocation struct {
	// Pointer to the action
	Action *ActionWebauthnAuthenticationFinish

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionWebauthnAuthenticationFinishInput
	// Global meta input parameters
	MetaInput *ActionWebauthnAuthenticationFinishMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionWebauthnAuthenticationFinishInvocation) NewInput() *ActionWebauthnAuthenticationFinishInput {
	inv.Input = &ActionWebauthnAuthenticationFinishInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionWebauthnAuthenticationFinishInvocation) SetInput(input *ActionWebauthnAuthenticationFinishInput) *ActionWebauthnAuthenticationFinishInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionWebauthnAuthenticationFinishInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionWebauthnAuthenticationFinishInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionWebauthnAuthenticationFinishInvocation) NewMetaInput() *ActionWebauthnAuthenticationFinishMetaGlobalInput {
	inv.MetaInput = &ActionWebauthnAuthenticationFinishMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionWebauthnAuthenticationFinishInvocation) SetMetaInput(input *ActionWebauthnAuthenticationFinishMetaGlobalInput) *ActionWebauthnAuthenticationFinishInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionWebauthnAuthenticationFinishInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionWebauthnAuthenticationFinishInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionWebauthnAuthenticationFinishInvocation) Call() (*ActionWebauthnAuthenticationFinishResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionWebauthnAuthenticationFinishInvocation) callAsBody() (*ActionWebauthnAuthenticationFinishResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionWebauthnAuthenticationFinishResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionWebauthnAuthenticationFinishInvocation) makeAllInputParams() *ActionWebauthnAuthenticationFinishRequest {
	return &ActionWebauthnAuthenticationFinishRequest{
		Authentication: inv.makeInputParams(),
		Meta:           inv.makeMetaInputParams(),
	}
}

func (inv *ActionWebauthnAuthenticationFinishInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("AuthToken") {
			ret["auth_token"] = inv.Input.AuthToken
		}
		if inv.IsParameterSelected("ChallengeToken") {
			ret["challenge_token"] = inv.Input.ChallengeToken
		}
	}

	return ret
}

func (inv *ActionWebauthnAuthenticationFinishInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
