package client

import (
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
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type AuthTokenActionTokenRevokeResponse struct {
	Action *AuthTokenActionTokenRevoke `json:"-"`
	*Envelope
}

// Call the action directly without any path or input parameters
func (action *AuthTokenActionTokenRevoke) Call() (*AuthTokenActionTokenRevokeResponse, error) {
	return action.Prepare().Call()
}

// Prepare the action for invocation
func (action *AuthTokenActionTokenRevoke) Prepare() *AuthTokenActionTokenRevokeInvocation {
	return &AuthTokenActionTokenRevokeInvocation{
		Action: action,
		Path: "/_auth/token/tokens/revoke",
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

// Call() invokes the action and returns a response from the API server
func (inv *AuthTokenActionTokenRevokeInvocation) Call() (*AuthTokenActionTokenRevokeResponse, error) {
	return inv.callAsBody()
}


func (inv *AuthTokenActionTokenRevokeInvocation) callAsBody() (*AuthTokenActionTokenRevokeResponse, error) {
	input := inv.makeAllInputParams()
	resp := &AuthTokenActionTokenRevokeResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
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
