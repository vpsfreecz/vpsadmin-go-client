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
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionLocationUpdateMetaGlobalInput) SetIncludes(value string) *ActionLocationUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
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
	Label string `json:"label"`
	HasIpv6 bool `json:"has_ipv6"`
	VpsOnboot bool `json:"vps_onboot"`
	RemoteConsoleServer string `json:"remote_console_server"`
	Domain string `json:"domain"`
	Environment int64 `json:"environment"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
// SetHasIpv6 sets parameter HasIpv6 to value and selects it for sending
func (in *ActionLocationUpdateInput) SetHasIpv6(value bool) *ActionLocationUpdateInput {
	in.HasIpv6 = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HasIpv6"] = nil
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
// SetRemoteConsoleServer sets parameter RemoteConsoleServer to value and selects it for sending
func (in *ActionLocationUpdateInput) SetRemoteConsoleServer(value string) *ActionLocationUpdateInput {
	in.RemoteConsoleServer = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RemoteConsoleServer"] = nil
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
	Meta map[string]interface{} `json:"_meta"`
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
		Path: "/v5.0/locations/:location_id",
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
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
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
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionLocationUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("HasIpv6") {
			ret["has_ipv6"] = inv.Input.HasIpv6
		}
		if inv.IsParameterSelected("VpsOnboot") {
			ret["vps_onboot"] = inv.Input.VpsOnboot
		}
		if inv.IsParameterSelected("RemoteConsoleServer") {
			ret["remote_console_server"] = inv.Input.RemoteConsoleServer
		}
		if inv.IsParameterSelected("Domain") {
			ret["domain"] = inv.Input.Domain
		}
		if inv.IsParameterSelected("Environment") {
			ret["environment"] = inv.Input.Environment
		}
	}

	return ret
}

func (inv *ActionLocationUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
