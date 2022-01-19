package client

import ()

// ActionApiServerUnlockTransactionSigningKey is a type for action Api_server#Unlock_transaction_signing_key
type ActionApiServerUnlockTransactionSigningKey struct {
	// Pointer to client
	Client *Client
}

func NewActionApiServerUnlockTransactionSigningKey(client *Client) *ActionApiServerUnlockTransactionSigningKey {
	return &ActionApiServerUnlockTransactionSigningKey{
		Client: client,
	}
}

// ActionApiServerUnlockTransactionSigningKeyMetaGlobalInput is a type for action global meta input parameters
type ActionApiServerUnlockTransactionSigningKeyMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionApiServerUnlockTransactionSigningKeyMetaGlobalInput) SetNo(value bool) *ActionApiServerUnlockTransactionSigningKeyMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionApiServerUnlockTransactionSigningKeyMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionApiServerUnlockTransactionSigningKeyMetaGlobalInput) SelectParameters(params ...string) *ActionApiServerUnlockTransactionSigningKeyMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionApiServerUnlockTransactionSigningKeyMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionApiServerUnlockTransactionSigningKeyInput is a type for action input parameters
type ActionApiServerUnlockTransactionSigningKeyInput struct {
	Passphrase string `json:"passphrase"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetPassphrase sets parameter Passphrase to value and selects it for sending
func (in *ActionApiServerUnlockTransactionSigningKeyInput) SetPassphrase(value string) *ActionApiServerUnlockTransactionSigningKeyInput {
	in.Passphrase = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Passphrase"] = nil
	return in
}

// SelectParameters sets parameters from ActionApiServerUnlockTransactionSigningKeyInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionApiServerUnlockTransactionSigningKeyInput) SelectParameters(params ...string) *ActionApiServerUnlockTransactionSigningKeyInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionApiServerUnlockTransactionSigningKeyInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// Type for action response, including envelope
type ActionApiServerUnlockTransactionSigningKeyResponse struct {
	Action *ActionApiServerUnlockTransactionSigningKey `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionApiServerUnlockTransactionSigningKey) Prepare() *ActionApiServerUnlockTransactionSigningKeyInvocation {
	return &ActionApiServerUnlockTransactionSigningKeyInvocation{
		Action: action,
		Path:   "/v6.0/api_servers/unlock_transaction_signing_key",
	}
}

// ActionApiServerUnlockTransactionSigningKeyInvocation is used to configure action for invocation
type ActionApiServerUnlockTransactionSigningKeyInvocation struct {
	// Pointer to the action
	Action *ActionApiServerUnlockTransactionSigningKey

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionApiServerUnlockTransactionSigningKeyInput
	// Global meta input parameters
	MetaInput *ActionApiServerUnlockTransactionSigningKeyMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionApiServerUnlockTransactionSigningKeyInvocation) NewInput() *ActionApiServerUnlockTransactionSigningKeyInput {
	inv.Input = &ActionApiServerUnlockTransactionSigningKeyInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionApiServerUnlockTransactionSigningKeyInvocation) SetInput(input *ActionApiServerUnlockTransactionSigningKeyInput) *ActionApiServerUnlockTransactionSigningKeyInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionApiServerUnlockTransactionSigningKeyInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionApiServerUnlockTransactionSigningKeyInvocation) NewMetaInput() *ActionApiServerUnlockTransactionSigningKeyMetaGlobalInput {
	inv.MetaInput = &ActionApiServerUnlockTransactionSigningKeyMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionApiServerUnlockTransactionSigningKeyInvocation) SetMetaInput(input *ActionApiServerUnlockTransactionSigningKeyMetaGlobalInput) *ActionApiServerUnlockTransactionSigningKeyInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionApiServerUnlockTransactionSigningKeyInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionApiServerUnlockTransactionSigningKeyInvocation) Call() (*ActionApiServerUnlockTransactionSigningKeyResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionApiServerUnlockTransactionSigningKeyInvocation) callAsQuery() (*ActionApiServerUnlockTransactionSigningKeyResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionApiServerUnlockTransactionSigningKeyResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	return resp, err
}

func (inv *ActionApiServerUnlockTransactionSigningKeyInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Passphrase") {
			ret["api_server[passphrase]"] = inv.Input.Passphrase
		}
	}
}

func (inv *ActionApiServerUnlockTransactionSigningKeyInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
