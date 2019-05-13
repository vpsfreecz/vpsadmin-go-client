package client

import (
	"strings"
)

// ActionUserTotpConfirm is a type for action User#Totp_confirm
type ActionUserTotpConfirm struct {
	// Pointer to client
	Client *Client
}

func NewActionUserTotpConfirm(client *Client) *ActionUserTotpConfirm {
	return &ActionUserTotpConfirm{
		Client: client,
	}
}

// ActionUserTotpConfirmMetaGlobalInput is a type for action global meta input parameters
type ActionUserTotpConfirmMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserTotpConfirmMetaGlobalInput) SetNo(value bool) *ActionUserTotpConfirmMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserTotpConfirmMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserTotpConfirmMetaGlobalInput) SelectParameters(params ...string) *ActionUserTotpConfirmMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserTotpConfirmMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserTotpConfirmInput is a type for action input parameters
type ActionUserTotpConfirmInput struct {
	Code string `json:"code"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCode sets parameter Code to value and selects it for sending
func (in *ActionUserTotpConfirmInput) SetCode(value string) *ActionUserTotpConfirmInput {
	in.Code = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Code"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserTotpConfirmInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserTotpConfirmInput) SelectParameters(params ...string) *ActionUserTotpConfirmInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserTotpConfirmInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserTotpConfirmRequest is a type for the entire action request
type ActionUserTotpConfirmRequest struct {
	User map[string]interface{} `json:"user"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionUserTotpConfirmOutput is a type for action output parameters
type ActionUserTotpConfirmOutput struct {
	RecoveryCode string `json:"recovery_code"`
}


// Type for action response, including envelope
type ActionUserTotpConfirmResponse struct {
	Action *ActionUserTotpConfirm `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		User *ActionUserTotpConfirmOutput `json:"user"`
	}

	// Action output without the namespace
	Output *ActionUserTotpConfirmOutput
}


// Prepare the action for invocation
func (action *ActionUserTotpConfirm) Prepare() *ActionUserTotpConfirmInvocation {
	return &ActionUserTotpConfirmInvocation{
		Action: action,
		Path: "/v5.0/users/totp_confirm/{user_id}",
	}
}

// ActionUserTotpConfirmInvocation is used to configure action for invocation
type ActionUserTotpConfirmInvocation struct {
	// Pointer to the action
	Action *ActionUserTotpConfirm

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserTotpConfirmInput
	// Global meta input parameters
	MetaInput *ActionUserTotpConfirmMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserTotpConfirmInvocation) SetPathParamInt(param string, value int64) *ActionUserTotpConfirmInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserTotpConfirmInvocation) SetPathParamString(param string, value string) *ActionUserTotpConfirmInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserTotpConfirmInvocation) NewInput() *ActionUserTotpConfirmInput {
	inv.Input = &ActionUserTotpConfirmInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserTotpConfirmInvocation) SetInput(input *ActionUserTotpConfirmInput) *ActionUserTotpConfirmInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserTotpConfirmInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserTotpConfirmInvocation) NewMetaInput() *ActionUserTotpConfirmMetaGlobalInput {
	inv.MetaInput = &ActionUserTotpConfirmMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserTotpConfirmInvocation) SetMetaInput(input *ActionUserTotpConfirmMetaGlobalInput) *ActionUserTotpConfirmInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserTotpConfirmInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserTotpConfirmInvocation) Call() (*ActionUserTotpConfirmResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionUserTotpConfirmInvocation) callAsBody() (*ActionUserTotpConfirmResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserTotpConfirmResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.User
	}
	return resp, err
}




func (inv *ActionUserTotpConfirmInvocation) makeAllInputParams() *ActionUserTotpConfirmRequest {
	return &ActionUserTotpConfirmRequest{
		User: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserTotpConfirmInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Code") {
			ret["code"] = inv.Input.Code
		}
	}

	return ret
}

func (inv *ActionUserTotpConfirmInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
