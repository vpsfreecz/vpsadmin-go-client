package client

import (
	"strings"
)

// ActionUserKnownDeviceDelete is a type for action User.Known_device#Delete
type ActionUserKnownDeviceDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionUserKnownDeviceDelete(client *Client) *ActionUserKnownDeviceDelete {
	return &ActionUserKnownDeviceDelete{
		Client: client,
	}
}

// ActionUserKnownDeviceDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionUserKnownDeviceDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserKnownDeviceDeleteMetaGlobalInput) SetIncludes(value string) *ActionUserKnownDeviceDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserKnownDeviceDeleteMetaGlobalInput) SetNo(value bool) *ActionUserKnownDeviceDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserKnownDeviceDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserKnownDeviceDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionUserKnownDeviceDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserKnownDeviceDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserKnownDeviceDeleteRequest is a type for the entire action request
type ActionUserKnownDeviceDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionUserKnownDeviceDeleteResponse struct {
	Action *ActionUserKnownDeviceDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionUserKnownDeviceDelete) Prepare() *ActionUserKnownDeviceDeleteInvocation {
	return &ActionUserKnownDeviceDeleteInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/known_devices/{known_device_id}",
	}
}

// ActionUserKnownDeviceDeleteInvocation is used to configure action for invocation
type ActionUserKnownDeviceDeleteInvocation struct {
	// Pointer to the action
	Action *ActionUserKnownDeviceDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserKnownDeviceDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserKnownDeviceDeleteInvocation) SetPathParamInt(param string, value int64) *ActionUserKnownDeviceDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserKnownDeviceDeleteInvocation) SetPathParamString(param string, value string) *ActionUserKnownDeviceDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserKnownDeviceDeleteInvocation) NewMetaInput() *ActionUserKnownDeviceDeleteMetaGlobalInput {
	inv.MetaInput = &ActionUserKnownDeviceDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserKnownDeviceDeleteInvocation) SetMetaInput(input *ActionUserKnownDeviceDeleteMetaGlobalInput) *ActionUserKnownDeviceDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserKnownDeviceDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserKnownDeviceDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserKnownDeviceDeleteInvocation) Call() (*ActionUserKnownDeviceDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionUserKnownDeviceDeleteInvocation) callAsBody() (*ActionUserKnownDeviceDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserKnownDeviceDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionUserKnownDeviceDeleteInvocation) makeAllInputParams() *ActionUserKnownDeviceDeleteRequest {
	return &ActionUserKnownDeviceDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserKnownDeviceDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
