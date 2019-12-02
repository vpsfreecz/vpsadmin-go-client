package client

import (
	"strings"
)

// ActionUserTotpDeviceCreate is a type for action User.Totp_device#Create
type ActionUserTotpDeviceCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserTotpDeviceCreate(client *Client) *ActionUserTotpDeviceCreate {
	return &ActionUserTotpDeviceCreate{
		Client: client,
	}
}

// ActionUserTotpDeviceCreateMetaGlobalInput is a type for action global meta input parameters
type ActionUserTotpDeviceCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserTotpDeviceCreateMetaGlobalInput) SetNo(value bool) *ActionUserTotpDeviceCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserTotpDeviceCreateMetaGlobalInput) SetIncludes(value string) *ActionUserTotpDeviceCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserTotpDeviceCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserTotpDeviceCreateMetaGlobalInput) SelectParameters(params ...string) *ActionUserTotpDeviceCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserTotpDeviceCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserTotpDeviceCreateInput is a type for action input parameters
type ActionUserTotpDeviceCreateInput struct {
	Label string `json:"label"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionUserTotpDeviceCreateInput) SetLabel(value string) *ActionUserTotpDeviceCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserTotpDeviceCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserTotpDeviceCreateInput) SelectParameters(params ...string) *ActionUserTotpDeviceCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserTotpDeviceCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserTotpDeviceCreateRequest is a type for the entire action request
type ActionUserTotpDeviceCreateRequest struct {
	TotpDevice map[string]interface{} `json:"totp_device"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionUserTotpDeviceCreateOutput is a type for action output parameters
type ActionUserTotpDeviceCreateOutput struct {
	Id int64 `json:"id"`
	Label string `json:"label"`
	Confirmed bool `json:"confirmed"`
	Enabled bool `json:"enabled"`
	LastUseAt string `json:"last_use_at"`
	UseCount int64 `json:"use_count"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Secret string `json:"secret"`
	ProvisioningUri string `json:"provisioning_uri"`
}


// Type for action response, including envelope
type ActionUserTotpDeviceCreateResponse struct {
	Action *ActionUserTotpDeviceCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		TotpDevice *ActionUserTotpDeviceCreateOutput `json:"totp_device"`
	}

	// Action output without the namespace
	Output *ActionUserTotpDeviceCreateOutput
}


// Prepare the action for invocation
func (action *ActionUserTotpDeviceCreate) Prepare() *ActionUserTotpDeviceCreateInvocation {
	return &ActionUserTotpDeviceCreateInvocation{
		Action: action,
		Path: "/v5.0/users/{user_id}/totp_devices",
	}
}

// ActionUserTotpDeviceCreateInvocation is used to configure action for invocation
type ActionUserTotpDeviceCreateInvocation struct {
	// Pointer to the action
	Action *ActionUserTotpDeviceCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserTotpDeviceCreateInput
	// Global meta input parameters
	MetaInput *ActionUserTotpDeviceCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserTotpDeviceCreateInvocation) SetPathParamInt(param string, value int64) *ActionUserTotpDeviceCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserTotpDeviceCreateInvocation) SetPathParamString(param string, value string) *ActionUserTotpDeviceCreateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserTotpDeviceCreateInvocation) NewInput() *ActionUserTotpDeviceCreateInput {
	inv.Input = &ActionUserTotpDeviceCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserTotpDeviceCreateInvocation) SetInput(input *ActionUserTotpDeviceCreateInput) *ActionUserTotpDeviceCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserTotpDeviceCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserTotpDeviceCreateInvocation) NewMetaInput() *ActionUserTotpDeviceCreateMetaGlobalInput {
	inv.MetaInput = &ActionUserTotpDeviceCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserTotpDeviceCreateInvocation) SetMetaInput(input *ActionUserTotpDeviceCreateMetaGlobalInput) *ActionUserTotpDeviceCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserTotpDeviceCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserTotpDeviceCreateInvocation) Call() (*ActionUserTotpDeviceCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionUserTotpDeviceCreateInvocation) callAsBody() (*ActionUserTotpDeviceCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserTotpDeviceCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.TotpDevice
	}
	return resp, err
}




func (inv *ActionUserTotpDeviceCreateInvocation) makeAllInputParams() *ActionUserTotpDeviceCreateRequest {
	return &ActionUserTotpDeviceCreateRequest{
		TotpDevice: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserTotpDeviceCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
	}

	return ret
}

func (inv *ActionUserTotpDeviceCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
