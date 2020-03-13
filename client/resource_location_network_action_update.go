package client

import (
	"strings"
)

// ActionLocationNetworkUpdate is a type for action Location_network#Update
type ActionLocationNetworkUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionLocationNetworkUpdate(client *Client) *ActionLocationNetworkUpdate {
	return &ActionLocationNetworkUpdate{
		Client: client,
	}
}

// ActionLocationNetworkUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionLocationNetworkUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionLocationNetworkUpdateMetaGlobalInput) SetNo(value bool) *ActionLocationNetworkUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionLocationNetworkUpdateMetaGlobalInput) SetIncludes(value string) *ActionLocationNetworkUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionLocationNetworkUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLocationNetworkUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionLocationNetworkUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionLocationNetworkUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionLocationNetworkUpdateInput is a type for action input parameters
type ActionLocationNetworkUpdateInput struct {
	Priority int64 `json:"priority"`
	Autopick bool `json:"autopick"`
	Userpick bool `json:"userpick"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetPriority sets parameter Priority to value and selects it for sending
func (in *ActionLocationNetworkUpdateInput) SetPriority(value int64) *ActionLocationNetworkUpdateInput {
	in.Priority = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Priority"] = nil
	return in
}
// SetAutopick sets parameter Autopick to value and selects it for sending
func (in *ActionLocationNetworkUpdateInput) SetAutopick(value bool) *ActionLocationNetworkUpdateInput {
	in.Autopick = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Autopick"] = nil
	return in
}
// SetUserpick sets parameter Userpick to value and selects it for sending
func (in *ActionLocationNetworkUpdateInput) SetUserpick(value bool) *ActionLocationNetworkUpdateInput {
	in.Userpick = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Userpick"] = nil
	return in
}

// SelectParameters sets parameters from ActionLocationNetworkUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLocationNetworkUpdateInput) SelectParameters(params ...string) *ActionLocationNetworkUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionLocationNetworkUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionLocationNetworkUpdateRequest is a type for the entire action request
type ActionLocationNetworkUpdateRequest struct {
	LocationNetwork map[string]interface{} `json:"location_network"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionLocationNetworkUpdateOutput is a type for action output parameters
type ActionLocationNetworkUpdateOutput struct {
	Id int64 `json:"id"`
	Location *ActionLocationShowOutput `json:"location"`
	Network *ActionNetworkShowOutput `json:"network"`
	Priority int64 `json:"priority"`
	Autopick bool `json:"autopick"`
	Userpick bool `json:"userpick"`
}


// Type for action response, including envelope
type ActionLocationNetworkUpdateResponse struct {
	Action *ActionLocationNetworkUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		LocationNetwork *ActionLocationNetworkUpdateOutput `json:"location_network"`
	}

	// Action output without the namespace
	Output *ActionLocationNetworkUpdateOutput
}


// Prepare the action for invocation
func (action *ActionLocationNetworkUpdate) Prepare() *ActionLocationNetworkUpdateInvocation {
	return &ActionLocationNetworkUpdateInvocation{
		Action: action,
		Path: "/v6.0/location_networks/{location_network_id}",
	}
}

// ActionLocationNetworkUpdateInvocation is used to configure action for invocation
type ActionLocationNetworkUpdateInvocation struct {
	// Pointer to the action
	Action *ActionLocationNetworkUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionLocationNetworkUpdateInput
	// Global meta input parameters
	MetaInput *ActionLocationNetworkUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionLocationNetworkUpdateInvocation) SetPathParamInt(param string, value int64) *ActionLocationNetworkUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionLocationNetworkUpdateInvocation) SetPathParamString(param string, value string) *ActionLocationNetworkUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionLocationNetworkUpdateInvocation) NewInput() *ActionLocationNetworkUpdateInput {
	inv.Input = &ActionLocationNetworkUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionLocationNetworkUpdateInvocation) SetInput(input *ActionLocationNetworkUpdateInput) *ActionLocationNetworkUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionLocationNetworkUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionLocationNetworkUpdateInvocation) NewMetaInput() *ActionLocationNetworkUpdateMetaGlobalInput {
	inv.MetaInput = &ActionLocationNetworkUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionLocationNetworkUpdateInvocation) SetMetaInput(input *ActionLocationNetworkUpdateMetaGlobalInput) *ActionLocationNetworkUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionLocationNetworkUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionLocationNetworkUpdateInvocation) Call() (*ActionLocationNetworkUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionLocationNetworkUpdateInvocation) callAsBody() (*ActionLocationNetworkUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionLocationNetworkUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.LocationNetwork
	}
	return resp, err
}




func (inv *ActionLocationNetworkUpdateInvocation) makeAllInputParams() *ActionLocationNetworkUpdateRequest {
	return &ActionLocationNetworkUpdateRequest{
		LocationNetwork: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionLocationNetworkUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Priority") {
			ret["priority"] = inv.Input.Priority
		}
		if inv.IsParameterSelected("Autopick") {
			ret["autopick"] = inv.Input.Autopick
		}
		if inv.IsParameterSelected("Userpick") {
			ret["userpick"] = inv.Input.Userpick
		}
	}

	return ret
}

func (inv *ActionLocationNetworkUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
