package client

import (
	"strings"
)

// ActionUserTotpDeviceConfirm is a type for action User.Totp_device#Confirm
type ActionUserTotpDeviceConfirm struct {
	// Pointer to client
	Client *Client
}

func NewActionUserTotpDeviceConfirm(client *Client) *ActionUserTotpDeviceConfirm {
	return &ActionUserTotpDeviceConfirm{
		Client: client,
	}
}

// ActionUserTotpDeviceConfirmMetaGlobalInput is a type for action global meta input parameters
type ActionUserTotpDeviceConfirmMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserTotpDeviceConfirmMetaGlobalInput) SetNo(value bool) *ActionUserTotpDeviceConfirmMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserTotpDeviceConfirmMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserTotpDeviceConfirmMetaGlobalInput) SelectParameters(params ...string) *ActionUserTotpDeviceConfirmMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserTotpDeviceConfirmMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserTotpDeviceConfirmInput is a type for action input parameters
type ActionUserTotpDeviceConfirmInput struct {
	Code string `json:"code"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCode sets parameter Code to value and selects it for sending
func (in *ActionUserTotpDeviceConfirmInput) SetCode(value string) *ActionUserTotpDeviceConfirmInput {
	in.Code = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Code"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserTotpDeviceConfirmInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserTotpDeviceConfirmInput) SelectParameters(params ...string) *ActionUserTotpDeviceConfirmInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserTotpDeviceConfirmInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserTotpDeviceConfirmInput) UnselectParameters(params ...string) *ActionUserTotpDeviceConfirmInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserTotpDeviceConfirmInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserTotpDeviceConfirmRequest is a type for the entire action request
type ActionUserTotpDeviceConfirmRequest struct {
	TotpDevice map[string]interface{} `json:"totp_device"`
	Meta       map[string]interface{} `json:"_meta"`
}

// ActionUserTotpDeviceConfirmOutput is a type for action output parameters
type ActionUserTotpDeviceConfirmOutput struct {
	RecoveryCode string `json:"recovery_code"`
}

// Type for action response, including envelope
type ActionUserTotpDeviceConfirmResponse struct {
	Action *ActionUserTotpDeviceConfirm `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		TotpDevice *ActionUserTotpDeviceConfirmOutput `json:"totp_device"`
	}

	// Action output without the namespace
	Output *ActionUserTotpDeviceConfirmOutput
}

// Prepare the action for invocation
func (action *ActionUserTotpDeviceConfirm) Prepare() *ActionUserTotpDeviceConfirmInvocation {
	return &ActionUserTotpDeviceConfirmInvocation{
		Action: action,
		Path:   "/v6.0/users/{user_id}/totp_devices/{totp_device_id}/confirm",
	}
}

// ActionUserTotpDeviceConfirmInvocation is used to configure action for invocation
type ActionUserTotpDeviceConfirmInvocation struct {
	// Pointer to the action
	Action *ActionUserTotpDeviceConfirm

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserTotpDeviceConfirmInput
	// Global meta input parameters
	MetaInput *ActionUserTotpDeviceConfirmMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserTotpDeviceConfirmInvocation) SetPathParamInt(param string, value int64) *ActionUserTotpDeviceConfirmInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserTotpDeviceConfirmInvocation) SetPathParamString(param string, value string) *ActionUserTotpDeviceConfirmInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserTotpDeviceConfirmInvocation) NewInput() *ActionUserTotpDeviceConfirmInput {
	inv.Input = &ActionUserTotpDeviceConfirmInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserTotpDeviceConfirmInvocation) SetInput(input *ActionUserTotpDeviceConfirmInput) *ActionUserTotpDeviceConfirmInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserTotpDeviceConfirmInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserTotpDeviceConfirmInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserTotpDeviceConfirmInvocation) NewMetaInput() *ActionUserTotpDeviceConfirmMetaGlobalInput {
	inv.MetaInput = &ActionUserTotpDeviceConfirmMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserTotpDeviceConfirmInvocation) SetMetaInput(input *ActionUserTotpDeviceConfirmMetaGlobalInput) *ActionUserTotpDeviceConfirmInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserTotpDeviceConfirmInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserTotpDeviceConfirmInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserTotpDeviceConfirmInvocation) Call() (*ActionUserTotpDeviceConfirmResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionUserTotpDeviceConfirmInvocation) callAsBody() (*ActionUserTotpDeviceConfirmResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserTotpDeviceConfirmResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.TotpDevice
	}
	return resp, err
}

func (inv *ActionUserTotpDeviceConfirmInvocation) makeAllInputParams() *ActionUserTotpDeviceConfirmRequest {
	return &ActionUserTotpDeviceConfirmRequest{
		TotpDevice: inv.makeInputParams(),
		Meta:       inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserTotpDeviceConfirmInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Code") {
			ret["code"] = inv.Input.Code
		}
	}

	return ret
}

func (inv *ActionUserTotpDeviceConfirmInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
