package client

import ()

// ActionWebauthnAuthenticationBegin is a type for action Webauthn.Authentication#Begin
type ActionWebauthnAuthenticationBegin struct {
	// Pointer to client
	Client *Client
}

func NewActionWebauthnAuthenticationBegin(client *Client) *ActionWebauthnAuthenticationBegin {
	return &ActionWebauthnAuthenticationBegin{
		Client: client,
	}
}

// ActionWebauthnAuthenticationBeginMetaGlobalInput is a type for action global meta input parameters
type ActionWebauthnAuthenticationBeginMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionWebauthnAuthenticationBeginMetaGlobalInput) SetNo(value bool) *ActionWebauthnAuthenticationBeginMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionWebauthnAuthenticationBeginMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionWebauthnAuthenticationBeginMetaGlobalInput) SelectParameters(params ...string) *ActionWebauthnAuthenticationBeginMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionWebauthnAuthenticationBeginMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionWebauthnAuthenticationBeginInput is a type for action input parameters
type ActionWebauthnAuthenticationBeginInput struct {
	AuthToken string `json:"auth_token"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAuthToken sets parameter AuthToken to value and selects it for sending
func (in *ActionWebauthnAuthenticationBeginInput) SetAuthToken(value string) *ActionWebauthnAuthenticationBeginInput {
	in.AuthToken = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AuthToken"] = nil
	return in
}

// SelectParameters sets parameters from ActionWebauthnAuthenticationBeginInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionWebauthnAuthenticationBeginInput) SelectParameters(params ...string) *ActionWebauthnAuthenticationBeginInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionWebauthnAuthenticationBeginInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionWebauthnAuthenticationBeginInput) UnselectParameters(params ...string) *ActionWebauthnAuthenticationBeginInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionWebauthnAuthenticationBeginInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionWebauthnAuthenticationBeginRequest is a type for the entire action request
type ActionWebauthnAuthenticationBeginRequest struct {
	Authentication map[string]interface{} `json:"authentication"`
	Meta           map[string]interface{} `json:"_meta"`
}

// ActionWebauthnAuthenticationBeginOutput is a type for action output parameters
type ActionWebauthnAuthenticationBeginOutput struct {
	ChallengeToken string `json:"challenge_token"`
}

// Type for action response, including envelope
type ActionWebauthnAuthenticationBeginResponse struct {
	Action *ActionWebauthnAuthenticationBegin `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Authentication *ActionWebauthnAuthenticationBeginOutput `json:"authentication"`
	}

	// Action output without the namespace
	Output *ActionWebauthnAuthenticationBeginOutput
}

// Prepare the action for invocation
func (action *ActionWebauthnAuthenticationBegin) Prepare() *ActionWebauthnAuthenticationBeginInvocation {
	return &ActionWebauthnAuthenticationBeginInvocation{
		Action: action,
		Path:   "/v7.0/webauthn/authentication/begin",
	}
}

// ActionWebauthnAuthenticationBeginInvocation is used to configure action for invocation
type ActionWebauthnAuthenticationBeginInvocation struct {
	// Pointer to the action
	Action *ActionWebauthnAuthenticationBegin

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionWebauthnAuthenticationBeginInput
	// Global meta input parameters
	MetaInput *ActionWebauthnAuthenticationBeginMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionWebauthnAuthenticationBeginInvocation) NewInput() *ActionWebauthnAuthenticationBeginInput {
	inv.Input = &ActionWebauthnAuthenticationBeginInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionWebauthnAuthenticationBeginInvocation) SetInput(input *ActionWebauthnAuthenticationBeginInput) *ActionWebauthnAuthenticationBeginInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionWebauthnAuthenticationBeginInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionWebauthnAuthenticationBeginInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionWebauthnAuthenticationBeginInvocation) NewMetaInput() *ActionWebauthnAuthenticationBeginMetaGlobalInput {
	inv.MetaInput = &ActionWebauthnAuthenticationBeginMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionWebauthnAuthenticationBeginInvocation) SetMetaInput(input *ActionWebauthnAuthenticationBeginMetaGlobalInput) *ActionWebauthnAuthenticationBeginInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionWebauthnAuthenticationBeginInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionWebauthnAuthenticationBeginInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionWebauthnAuthenticationBeginInvocation) Call() (*ActionWebauthnAuthenticationBeginResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionWebauthnAuthenticationBeginInvocation) callAsBody() (*ActionWebauthnAuthenticationBeginResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionWebauthnAuthenticationBeginResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Authentication
	}
	return resp, err
}

func (inv *ActionWebauthnAuthenticationBeginInvocation) makeAllInputParams() *ActionWebauthnAuthenticationBeginRequest {
	return &ActionWebauthnAuthenticationBeginRequest{
		Authentication: inv.makeInputParams(),
		Meta:           inv.makeMetaInputParams(),
	}
}

func (inv *ActionWebauthnAuthenticationBeginInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("AuthToken") {
			ret["auth_token"] = inv.Input.AuthToken
		}
	}

	return ret
}

func (inv *ActionWebauthnAuthenticationBeginInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
