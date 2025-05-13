package client

import ()

// ActionWebauthnRegistrationFinish is a type for action Webauthn.Registration#Finish
type ActionWebauthnRegistrationFinish struct {
	// Pointer to client
	Client *Client
}

func NewActionWebauthnRegistrationFinish(client *Client) *ActionWebauthnRegistrationFinish {
	return &ActionWebauthnRegistrationFinish{
		Client: client,
	}
}

// ActionWebauthnRegistrationFinishMetaGlobalInput is a type for action global meta input parameters
type ActionWebauthnRegistrationFinishMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionWebauthnRegistrationFinishMetaGlobalInput) SetNo(value bool) *ActionWebauthnRegistrationFinishMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionWebauthnRegistrationFinishMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionWebauthnRegistrationFinishMetaGlobalInput) SelectParameters(params ...string) *ActionWebauthnRegistrationFinishMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionWebauthnRegistrationFinishMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionWebauthnRegistrationFinishInput is a type for action input parameters
type ActionWebauthnRegistrationFinishInput struct {
	ChallengeToken string `json:"challenge_token"`
	Label          string `json:"label"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetChallengeToken sets parameter ChallengeToken to value and selects it for sending
func (in *ActionWebauthnRegistrationFinishInput) SetChallengeToken(value string) *ActionWebauthnRegistrationFinishInput {
	in.ChallengeToken = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ChallengeToken"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionWebauthnRegistrationFinishInput) SetLabel(value string) *ActionWebauthnRegistrationFinishInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SelectParameters sets parameters from ActionWebauthnRegistrationFinishInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionWebauthnRegistrationFinishInput) SelectParameters(params ...string) *ActionWebauthnRegistrationFinishInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionWebauthnRegistrationFinishInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionWebauthnRegistrationFinishInput) UnselectParameters(params ...string) *ActionWebauthnRegistrationFinishInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionWebauthnRegistrationFinishInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionWebauthnRegistrationFinishRequest is a type for the entire action request
type ActionWebauthnRegistrationFinishRequest struct {
	Registration map[string]interface{} `json:"registration"`
	Meta         map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionWebauthnRegistrationFinishResponse struct {
	Action *ActionWebauthnRegistrationFinish `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionWebauthnRegistrationFinish) Prepare() *ActionWebauthnRegistrationFinishInvocation {
	return &ActionWebauthnRegistrationFinishInvocation{
		Action: action,
		Path:   "/v7.0/webauthn/registration/finish",
	}
}

// ActionWebauthnRegistrationFinishInvocation is used to configure action for invocation
type ActionWebauthnRegistrationFinishInvocation struct {
	// Pointer to the action
	Action *ActionWebauthnRegistrationFinish

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionWebauthnRegistrationFinishInput
	// Global meta input parameters
	MetaInput *ActionWebauthnRegistrationFinishMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionWebauthnRegistrationFinishInvocation) NewInput() *ActionWebauthnRegistrationFinishInput {
	inv.Input = &ActionWebauthnRegistrationFinishInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionWebauthnRegistrationFinishInvocation) SetInput(input *ActionWebauthnRegistrationFinishInput) *ActionWebauthnRegistrationFinishInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionWebauthnRegistrationFinishInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionWebauthnRegistrationFinishInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionWebauthnRegistrationFinishInvocation) NewMetaInput() *ActionWebauthnRegistrationFinishMetaGlobalInput {
	inv.MetaInput = &ActionWebauthnRegistrationFinishMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionWebauthnRegistrationFinishInvocation) SetMetaInput(input *ActionWebauthnRegistrationFinishMetaGlobalInput) *ActionWebauthnRegistrationFinishInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionWebauthnRegistrationFinishInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionWebauthnRegistrationFinishInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionWebauthnRegistrationFinishInvocation) Call() (*ActionWebauthnRegistrationFinishResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionWebauthnRegistrationFinishInvocation) callAsBody() (*ActionWebauthnRegistrationFinishResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionWebauthnRegistrationFinishResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionWebauthnRegistrationFinishInvocation) makeAllInputParams() *ActionWebauthnRegistrationFinishRequest {
	return &ActionWebauthnRegistrationFinishRequest{
		Registration: inv.makeInputParams(),
		Meta:         inv.makeMetaInputParams(),
	}
}

func (inv *ActionWebauthnRegistrationFinishInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("ChallengeToken") {
			ret["challenge_token"] = inv.Input.ChallengeToken
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
	}

	return ret
}

func (inv *ActionWebauthnRegistrationFinishInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
