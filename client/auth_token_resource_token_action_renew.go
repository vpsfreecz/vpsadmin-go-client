package client

import ()

// AuthTokenActionTokenRenew is a type for action Token#Renew
type AuthTokenActionTokenRenew struct {
	// Pointer to client
	Client *Client
}

func NewAuthTokenActionTokenRenew(client *Client) *AuthTokenActionTokenRenew {
	return &AuthTokenActionTokenRenew{
		Client: client,
	}
}

// AuthTokenActionTokenRenewMetaGlobalInput is a type for action global meta input parameters
type AuthTokenActionTokenRenewMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *AuthTokenActionTokenRenewMetaGlobalInput) SetNo(value bool) *AuthTokenActionTokenRenewMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from AuthTokenActionTokenRenewMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *AuthTokenActionTokenRenewMetaGlobalInput) SelectParameters(params ...string) *AuthTokenActionTokenRenewMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *AuthTokenActionTokenRenewMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// AuthTokenActionTokenRenewRequest is a type for the entire action request
type AuthTokenActionTokenRenewRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// AuthTokenActionTokenRenewOutput is a type for action output parameters
type AuthTokenActionTokenRenewOutput struct {
	ValidTo string `json:"valid_to"`
}

// Type for action response, including envelope
type AuthTokenActionTokenRenewResponse struct {
	Action *AuthTokenActionTokenRenew `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Token *AuthTokenActionTokenRenewOutput `json:"token"`
	}

	// Action output without the namespace
	Output *AuthTokenActionTokenRenewOutput
}

// Call the action directly without any path or input parameters
func (action *AuthTokenActionTokenRenew) Call() (*AuthTokenActionTokenRenewResponse, error) {
	return action.Prepare().Call()
}

// Prepare the action for invocation
func (action *AuthTokenActionTokenRenew) Prepare() *AuthTokenActionTokenRenewInvocation {
	return &AuthTokenActionTokenRenewInvocation{
		Action: action,
		Path:   "/_auth/token/tokens/renew",
	}
}

// AuthTokenActionTokenRenewInvocation is used to configure action for invocation
type AuthTokenActionTokenRenewInvocation struct {
	// Pointer to the action
	Action *AuthTokenActionTokenRenew

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *AuthTokenActionTokenRenewMetaGlobalInput
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *AuthTokenActionTokenRenewInvocation) NewMetaInput() *AuthTokenActionTokenRenewMetaGlobalInput {
	inv.MetaInput = &AuthTokenActionTokenRenewMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *AuthTokenActionTokenRenewInvocation) SetMetaInput(input *AuthTokenActionTokenRenewMetaGlobalInput) *AuthTokenActionTokenRenewInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *AuthTokenActionTokenRenewInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *AuthTokenActionTokenRenewInvocation) Call() (*AuthTokenActionTokenRenewResponse, error) {
	return inv.callAsBody()
}

func (inv *AuthTokenActionTokenRenewInvocation) callAsBody() (*AuthTokenActionTokenRenewResponse, error) {
	input := inv.makeAllInputParams()
	resp := &AuthTokenActionTokenRenewResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Token
	}
	return resp, err
}

func (inv *AuthTokenActionTokenRenewInvocation) makeAllInputParams() *AuthTokenActionTokenRenewRequest {
	return &AuthTokenActionTokenRenewRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *AuthTokenActionTokenRenewInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
