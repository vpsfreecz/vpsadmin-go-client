package client

import ()

// ActionLocationNetworkCreate is a type for action Location_network#Create
type ActionLocationNetworkCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionLocationNetworkCreate(client *Client) *ActionLocationNetworkCreate {
	return &ActionLocationNetworkCreate{
		Client: client,
	}
}

// ActionLocationNetworkCreateMetaGlobalInput is a type for action global meta input parameters
type ActionLocationNetworkCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionLocationNetworkCreateMetaGlobalInput) SetIncludes(value string) *ActionLocationNetworkCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionLocationNetworkCreateMetaGlobalInput) SetNo(value bool) *ActionLocationNetworkCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionLocationNetworkCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLocationNetworkCreateMetaGlobalInput) SelectParameters(params ...string) *ActionLocationNetworkCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionLocationNetworkCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionLocationNetworkCreateInput is a type for action input parameters
type ActionLocationNetworkCreateInput struct {
	Autopick bool  `json:"autopick"`
	Location int64 `json:"location"`
	Network  int64 `json:"network"`
	Primary  bool  `json:"primary"`
	Priority int64 `json:"priority"`
	Userpick bool  `json:"userpick"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAutopick sets parameter Autopick to value and selects it for sending
func (in *ActionLocationNetworkCreateInput) SetAutopick(value bool) *ActionLocationNetworkCreateInput {
	in.Autopick = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Autopick"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionLocationNetworkCreateInput) SetLocation(value int64) *ActionLocationNetworkCreateInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLocationNil(false)
	in._selectedParameters["Location"] = nil
	return in
}

// SetLocationNil sets parameter Location to nil and selects it for sending
func (in *ActionLocationNetworkCreateInput) SetLocationNil(set bool) *ActionLocationNetworkCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Location"] = nil
		in.SelectParameters("Location")
	} else {
		delete(in._nilParameters, "Location")
	}
	return in
}

// SetNetwork sets parameter Network to value and selects it for sending
func (in *ActionLocationNetworkCreateInput) SetNetwork(value int64) *ActionLocationNetworkCreateInput {
	in.Network = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNetworkNil(false)
	in._selectedParameters["Network"] = nil
	return in
}

// SetNetworkNil sets parameter Network to nil and selects it for sending
func (in *ActionLocationNetworkCreateInput) SetNetworkNil(set bool) *ActionLocationNetworkCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Network"] = nil
		in.SelectParameters("Network")
	} else {
		delete(in._nilParameters, "Network")
	}
	return in
}

// SetPrimary sets parameter Primary to value and selects it for sending
func (in *ActionLocationNetworkCreateInput) SetPrimary(value bool) *ActionLocationNetworkCreateInput {
	in.Primary = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Primary"] = nil
	return in
}

// SetPriority sets parameter Priority to value and selects it for sending
func (in *ActionLocationNetworkCreateInput) SetPriority(value int64) *ActionLocationNetworkCreateInput {
	in.Priority = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Priority"] = nil
	return in
}

// SetUserpick sets parameter Userpick to value and selects it for sending
func (in *ActionLocationNetworkCreateInput) SetUserpick(value bool) *ActionLocationNetworkCreateInput {
	in.Userpick = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Userpick"] = nil
	return in
}

// SelectParameters sets parameters from ActionLocationNetworkCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLocationNetworkCreateInput) SelectParameters(params ...string) *ActionLocationNetworkCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionLocationNetworkCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionLocationNetworkCreateInput) UnselectParameters(params ...string) *ActionLocationNetworkCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionLocationNetworkCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionLocationNetworkCreateRequest is a type for the entire action request
type ActionLocationNetworkCreateRequest struct {
	LocationNetwork map[string]interface{} `json:"location_network"`
	Meta            map[string]interface{} `json:"_meta"`
}

// ActionLocationNetworkCreateOutput is a type for action output parameters
type ActionLocationNetworkCreateOutput struct {
	Autopick bool                      `json:"autopick"`
	Id       int64                     `json:"id"`
	Location *ActionLocationShowOutput `json:"location"`
	Network  *ActionNetworkShowOutput  `json:"network"`
	Primary  bool                      `json:"primary"`
	Priority int64                     `json:"priority"`
	Userpick bool                      `json:"userpick"`
}

// Type for action response, including envelope
type ActionLocationNetworkCreateResponse struct {
	Action *ActionLocationNetworkCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		LocationNetwork *ActionLocationNetworkCreateOutput `json:"location_network"`
	}

	// Action output without the namespace
	Output *ActionLocationNetworkCreateOutput
}

// Prepare the action for invocation
func (action *ActionLocationNetworkCreate) Prepare() *ActionLocationNetworkCreateInvocation {
	return &ActionLocationNetworkCreateInvocation{
		Action: action,
		Path:   "/v7.0/location_networks",
	}
}

// ActionLocationNetworkCreateInvocation is used to configure action for invocation
type ActionLocationNetworkCreateInvocation struct {
	// Pointer to the action
	Action *ActionLocationNetworkCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionLocationNetworkCreateInput
	// Global meta input parameters
	MetaInput *ActionLocationNetworkCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionLocationNetworkCreateInvocation) NewInput() *ActionLocationNetworkCreateInput {
	inv.Input = &ActionLocationNetworkCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionLocationNetworkCreateInvocation) SetInput(input *ActionLocationNetworkCreateInput) *ActionLocationNetworkCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionLocationNetworkCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionLocationNetworkCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionLocationNetworkCreateInvocation) NewMetaInput() *ActionLocationNetworkCreateMetaGlobalInput {
	inv.MetaInput = &ActionLocationNetworkCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionLocationNetworkCreateInvocation) SetMetaInput(input *ActionLocationNetworkCreateMetaGlobalInput) *ActionLocationNetworkCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionLocationNetworkCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionLocationNetworkCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionLocationNetworkCreateInvocation) Call() (*ActionLocationNetworkCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionLocationNetworkCreateInvocation) callAsBody() (*ActionLocationNetworkCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionLocationNetworkCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.LocationNetwork
	}
	return resp, err
}

func (inv *ActionLocationNetworkCreateInvocation) makeAllInputParams() *ActionLocationNetworkCreateRequest {
	return &ActionLocationNetworkCreateRequest{
		LocationNetwork: inv.makeInputParams(),
		Meta:            inv.makeMetaInputParams(),
	}
}

func (inv *ActionLocationNetworkCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Autopick") {
			ret["autopick"] = inv.Input.Autopick
		}
		if inv.IsParameterSelected("Location") {
			if inv.IsParameterNil("Location") {
				ret["location"] = nil
			} else {
				ret["location"] = inv.Input.Location
			}
		}
		if inv.IsParameterSelected("Network") {
			if inv.IsParameterNil("Network") {
				ret["network"] = nil
			} else {
				ret["network"] = inv.Input.Network
			}
		}
		if inv.IsParameterSelected("Primary") {
			ret["primary"] = inv.Input.Primary
		}
		if inv.IsParameterSelected("Priority") {
			ret["priority"] = inv.Input.Priority
		}
		if inv.IsParameterSelected("Userpick") {
			ret["userpick"] = inv.Input.Userpick
		}
	}

	return ret
}

func (inv *ActionLocationNetworkCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
