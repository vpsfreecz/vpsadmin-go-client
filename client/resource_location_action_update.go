package client

import (
	"strings"
)

// ActionLocationUpdate is a type for action Location#Update
type ActionLocationUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionLocationUpdate(client *Client) *ActionLocationUpdate {
	return &ActionLocationUpdate{
		Client: client,
	}
}

// ActionLocationUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionLocationUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionLocationUpdateMetaGlobalInput) SetIncludes(value string) *ActionLocationUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionLocationUpdateMetaGlobalInput) SetNo(value bool) *ActionLocationUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionLocationUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLocationUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionLocationUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionLocationUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionLocationUpdateInput is a type for action input parameters
type ActionLocationUpdateInput struct {
	Description         string `json:"description"`
	Domain              string `json:"domain"`
	Environment         int64  `json:"environment"`
	HasIpv6             bool   `json:"has_ipv6"`
	Label               string `json:"label"`
	RemoteConsoleServer string `json:"remote_console_server"`
	VpsOnboot           bool   `json:"vps_onboot"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetDescription sets parameter Description to value and selects it for sending
func (in *ActionLocationUpdateInput) SetDescription(value string) *ActionLocationUpdateInput {
	in.Description = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Description"] = nil
	return in
}

// SetDomain sets parameter Domain to value and selects it for sending
func (in *ActionLocationUpdateInput) SetDomain(value string) *ActionLocationUpdateInput {
	in.Domain = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Domain"] = nil
	return in
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionLocationUpdateInput) SetEnvironment(value int64) *ActionLocationUpdateInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Environment"] = nil
	return in
}

// SetHasIpv6 sets parameter HasIpv6 to value and selects it for sending
func (in *ActionLocationUpdateInput) SetHasIpv6(value bool) *ActionLocationUpdateInput {
	in.HasIpv6 = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HasIpv6"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionLocationUpdateInput) SetLabel(value string) *ActionLocationUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SetRemoteConsoleServer sets parameter RemoteConsoleServer to value and selects it for sending
func (in *ActionLocationUpdateInput) SetRemoteConsoleServer(value string) *ActionLocationUpdateInput {
	in.RemoteConsoleServer = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RemoteConsoleServer"] = nil
	return in
}

// SetVpsOnboot sets parameter VpsOnboot to value and selects it for sending
func (in *ActionLocationUpdateInput) SetVpsOnboot(value bool) *ActionLocationUpdateInput {
	in.VpsOnboot = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["VpsOnboot"] = nil
	return in
}

// SelectParameters sets parameters from ActionLocationUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLocationUpdateInput) SelectParameters(params ...string) *ActionLocationUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionLocationUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionLocationUpdateRequest is a type for the entire action request
type ActionLocationUpdateRequest struct {
	Location map[string]interface{} `json:"location"`
	Meta     map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionLocationUpdateResponse struct {
	Action *ActionLocationUpdate `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionLocationUpdate) Prepare() *ActionLocationUpdateInvocation {
	return &ActionLocationUpdateInvocation{
		Action: action,
		Path:   "/v6.0/locations/{location_id}",
	}
}

// ActionLocationUpdateInvocation is used to configure action for invocation
type ActionLocationUpdateInvocation struct {
	// Pointer to the action
	Action *ActionLocationUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionLocationUpdateInput
	// Global meta input parameters
	MetaInput *ActionLocationUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionLocationUpdateInvocation) SetPathParamInt(param string, value int64) *ActionLocationUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionLocationUpdateInvocation) SetPathParamString(param string, value string) *ActionLocationUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionLocationUpdateInvocation) NewInput() *ActionLocationUpdateInput {
	inv.Input = &ActionLocationUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionLocationUpdateInvocation) SetInput(input *ActionLocationUpdateInput) *ActionLocationUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionLocationUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionLocationUpdateInvocation) NewMetaInput() *ActionLocationUpdateMetaGlobalInput {
	inv.MetaInput = &ActionLocationUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionLocationUpdateInvocation) SetMetaInput(input *ActionLocationUpdateMetaGlobalInput) *ActionLocationUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionLocationUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionLocationUpdateInvocation) Call() (*ActionLocationUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionLocationUpdateInvocation) callAsBody() (*ActionLocationUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionLocationUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionLocationUpdateInvocation) makeAllInputParams() *ActionLocationUpdateRequest {
	return &ActionLocationUpdateRequest{
		Location: inv.makeInputParams(),
		Meta:     inv.makeMetaInputParams(),
	}
}

func (inv *ActionLocationUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Description") {
			ret["description"] = inv.Input.Description
		}
		if inv.IsParameterSelected("Domain") {
			ret["domain"] = inv.Input.Domain
		}
		if inv.IsParameterSelected("Environment") {
			ret["environment"] = inv.Input.Environment
		}
		if inv.IsParameterSelected("HasIpv6") {
			ret["has_ipv6"] = inv.Input.HasIpv6
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("RemoteConsoleServer") {
			ret["remote_console_server"] = inv.Input.RemoteConsoleServer
		}
		if inv.IsParameterSelected("VpsOnboot") {
			ret["vps_onboot"] = inv.Input.VpsOnboot
		}
	}

	return ret
}

func (inv *ActionLocationUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
