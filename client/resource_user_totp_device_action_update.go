package client

import (
	"strings"
)

// ActionUserTotpDeviceUpdate is a type for action User.Totp_device#Update
type ActionUserTotpDeviceUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserTotpDeviceUpdate(client *Client) *ActionUserTotpDeviceUpdate {
	return &ActionUserTotpDeviceUpdate{
		Client: client,
	}
}

// ActionUserTotpDeviceUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionUserTotpDeviceUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserTotpDeviceUpdateMetaGlobalInput) SetIncludes(value string) *ActionUserTotpDeviceUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserTotpDeviceUpdateMetaGlobalInput) SetNo(value bool) *ActionUserTotpDeviceUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserTotpDeviceUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserTotpDeviceUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionUserTotpDeviceUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserTotpDeviceUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserTotpDeviceUpdateInput is a type for action input parameters
type ActionUserTotpDeviceUpdateInput struct {
	Enabled bool `json:"enabled"`
	Label string `json:"label"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionUserTotpDeviceUpdateInput) SetEnabled(value bool) *ActionUserTotpDeviceUpdateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}
// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionUserTotpDeviceUpdateInput) SetLabel(value string) *ActionUserTotpDeviceUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserTotpDeviceUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserTotpDeviceUpdateInput) SelectParameters(params ...string) *ActionUserTotpDeviceUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserTotpDeviceUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserTotpDeviceUpdateRequest is a type for the entire action request
type ActionUserTotpDeviceUpdateRequest struct {
	TotpDevice map[string]interface{} `json:"totp_device"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionUserTotpDeviceUpdateOutput is a type for action output parameters
type ActionUserTotpDeviceUpdateOutput struct {
	Confirmed bool `json:"confirmed"`
	CreatedAt string `json:"created_at"`
	Enabled bool `json:"enabled"`
	Id int64 `json:"id"`
	Label string `json:"label"`
	LastUseAt string `json:"last_use_at"`
	UpdatedAt string `json:"updated_at"`
	UseCount int64 `json:"use_count"`
}


// Type for action response, including envelope
type ActionUserTotpDeviceUpdateResponse struct {
	Action *ActionUserTotpDeviceUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		TotpDevice *ActionUserTotpDeviceUpdateOutput `json:"totp_device"`
	}

	// Action output without the namespace
	Output *ActionUserTotpDeviceUpdateOutput
}


// Prepare the action for invocation
func (action *ActionUserTotpDeviceUpdate) Prepare() *ActionUserTotpDeviceUpdateInvocation {
	return &ActionUserTotpDeviceUpdateInvocation{
		Action: action,
		Path: "/v6.0/users/{user_id}/totp_devices/{totp_device_id}",
	}
}

// ActionUserTotpDeviceUpdateInvocation is used to configure action for invocation
type ActionUserTotpDeviceUpdateInvocation struct {
	// Pointer to the action
	Action *ActionUserTotpDeviceUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserTotpDeviceUpdateInput
	// Global meta input parameters
	MetaInput *ActionUserTotpDeviceUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserTotpDeviceUpdateInvocation) SetPathParamInt(param string, value int64) *ActionUserTotpDeviceUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserTotpDeviceUpdateInvocation) SetPathParamString(param string, value string) *ActionUserTotpDeviceUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserTotpDeviceUpdateInvocation) NewInput() *ActionUserTotpDeviceUpdateInput {
	inv.Input = &ActionUserTotpDeviceUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserTotpDeviceUpdateInvocation) SetInput(input *ActionUserTotpDeviceUpdateInput) *ActionUserTotpDeviceUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserTotpDeviceUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserTotpDeviceUpdateInvocation) NewMetaInput() *ActionUserTotpDeviceUpdateMetaGlobalInput {
	inv.MetaInput = &ActionUserTotpDeviceUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserTotpDeviceUpdateInvocation) SetMetaInput(input *ActionUserTotpDeviceUpdateMetaGlobalInput) *ActionUserTotpDeviceUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserTotpDeviceUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserTotpDeviceUpdateInvocation) Call() (*ActionUserTotpDeviceUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionUserTotpDeviceUpdateInvocation) callAsBody() (*ActionUserTotpDeviceUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserTotpDeviceUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.TotpDevice
	}
	return resp, err
}




func (inv *ActionUserTotpDeviceUpdateInvocation) makeAllInputParams() *ActionUserTotpDeviceUpdateRequest {
	return &ActionUserTotpDeviceUpdateRequest{
		TotpDevice: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserTotpDeviceUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Enabled") {
			ret["enabled"] = inv.Input.Enabled
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
	}

	return ret
}

func (inv *ActionUserTotpDeviceUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
