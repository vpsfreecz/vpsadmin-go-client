package client

import ()

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
	No bool `json:"no"`
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
	Meta map[string]interface{} `json:"_meta"`
}

// ActionWebauthnRegistrationBeginOutput is a type for action output parameters
type ActionWebauthnRegistrationBeginOutput struct {
	ChallengeToken string `json:"challenge_token"`
}

// Type for action response, including envelope
type ActionWebauthnRegistrationBeginResponse struct {
	Action *ActionWebauthnRegistrationBegin `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Registration *ActionWebauthnRegistrationBeginOutput `json:"registration"`
	}

	// Action output without the namespace
	Output *ActionWebauthnRegistrationBeginOutput
}

// Call the action directly without any path or input parameters
func (action *ActionWebauthnRegistrationBegin) Call() (*ActionWebauthnRegistrationBeginResponse, error) {
	return action.Prepare().Call()
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

// Call() invokes the action and returns a response from the API server
func (inv *ActionWebauthnRegistrationBeginInvocation) Call() (*ActionWebauthnRegistrationBeginResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionWebauthnRegistrationBeginInvocation) callAsBody() (*ActionWebauthnRegistrationBeginResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionWebauthnRegistrationBeginResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Registration
	}
	return resp, err
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
