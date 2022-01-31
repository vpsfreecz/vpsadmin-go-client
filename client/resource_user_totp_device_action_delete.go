package client

import (
	"strings"
)

// ActionUserTotpDeviceDelete is a type for action User.Totp_device#Delete
type ActionUserTotpDeviceDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionUserTotpDeviceDelete(client *Client) *ActionUserTotpDeviceDelete {
	return &ActionUserTotpDeviceDelete{
		Client: client,
	}
}

// ActionUserTotpDeviceDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionUserTotpDeviceDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserTotpDeviceDeleteMetaGlobalInput) SetIncludes(value string) *ActionUserTotpDeviceDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserTotpDeviceDeleteMetaGlobalInput) SetNo(value bool) *ActionUserTotpDeviceDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserTotpDeviceDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserTotpDeviceDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionUserTotpDeviceDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserTotpDeviceDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserTotpDeviceDeleteRequest is a type for the entire action request
type ActionUserTotpDeviceDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionUserTotpDeviceDeleteResponse struct {
	Action *ActionUserTotpDeviceDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionUserTotpDeviceDelete) Prepare() *ActionUserTotpDeviceDeleteInvocation {
	return &ActionUserTotpDeviceDeleteInvocation{
		Action: action,
		Path:   "/v6.0/users/{user_id}/totp_devices/{totp_device_id}",
	}
}

// ActionUserTotpDeviceDeleteInvocation is used to configure action for invocation
type ActionUserTotpDeviceDeleteInvocation struct {
	// Pointer to the action
	Action *ActionUserTotpDeviceDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserTotpDeviceDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserTotpDeviceDeleteInvocation) SetPathParamInt(param string, value int64) *ActionUserTotpDeviceDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserTotpDeviceDeleteInvocation) SetPathParamString(param string, value string) *ActionUserTotpDeviceDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserTotpDeviceDeleteInvocation) NewMetaInput() *ActionUserTotpDeviceDeleteMetaGlobalInput {
	inv.MetaInput = &ActionUserTotpDeviceDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserTotpDeviceDeleteInvocation) SetMetaInput(input *ActionUserTotpDeviceDeleteMetaGlobalInput) *ActionUserTotpDeviceDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserTotpDeviceDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserTotpDeviceDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserTotpDeviceDeleteInvocation) Call() (*ActionUserTotpDeviceDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionUserTotpDeviceDeleteInvocation) callAsBody() (*ActionUserTotpDeviceDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserTotpDeviceDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionUserTotpDeviceDeleteInvocation) makeAllInputParams() *ActionUserTotpDeviceDeleteRequest {
	return &ActionUserTotpDeviceDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserTotpDeviceDeleteInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
