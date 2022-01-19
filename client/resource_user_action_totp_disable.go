package client

import (
	"strings"
)

// ActionUserTotpDisable is a type for action User#Totp_disable
type ActionUserTotpDisable struct {
	// Pointer to client
	Client *Client
}

func NewActionUserTotpDisable(client *Client) *ActionUserTotpDisable {
	return &ActionUserTotpDisable{
		Client: client,
	}
}

// ActionUserTotpDisableMetaGlobalInput is a type for action global meta input parameters
type ActionUserTotpDisableMetaGlobalInput struct {
	No       bool   `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserTotpDisableMetaGlobalInput) SetNo(value bool) *ActionUserTotpDisableMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserTotpDisableMetaGlobalInput) SetIncludes(value string) *ActionUserTotpDisableMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserTotpDisableMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserTotpDisableMetaGlobalInput) SelectParameters(params ...string) *ActionUserTotpDisableMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserTotpDisableMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserTotpDisableRequest is a type for the entire action request
type ActionUserTotpDisableRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionUserTotpDisableResponse struct {
	Action *ActionUserTotpDisable `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionUserTotpDisable) Prepare() *ActionUserTotpDisableInvocation {
	return &ActionUserTotpDisableInvocation{
		Action: action,
		Path:   "/v5.0/users/totp_disable/{user_id}",
	}
}

// ActionUserTotpDisableInvocation is used to configure action for invocation
type ActionUserTotpDisableInvocation struct {
	// Pointer to the action
	Action *ActionUserTotpDisable

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserTotpDisableMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserTotpDisableInvocation) SetPathParamInt(param string, value int64) *ActionUserTotpDisableInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserTotpDisableInvocation) SetPathParamString(param string, value string) *ActionUserTotpDisableInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserTotpDisableInvocation) NewMetaInput() *ActionUserTotpDisableMetaGlobalInput {
	inv.MetaInput = &ActionUserTotpDisableMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserTotpDisableInvocation) SetMetaInput(input *ActionUserTotpDisableMetaGlobalInput) *ActionUserTotpDisableInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserTotpDisableInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserTotpDisableInvocation) Call() (*ActionUserTotpDisableResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionUserTotpDisableInvocation) callAsBody() (*ActionUserTotpDisableResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserTotpDisableResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionUserTotpDisableInvocation) makeAllInputParams() *ActionUserTotpDisableRequest {
	return &ActionUserTotpDisableRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserTotpDisableInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
