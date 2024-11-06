package client

import ()

// ActionHostIpAddressCreate is a type for action Host_ip_address#Create
type ActionHostIpAddressCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionHostIpAddressCreate(client *Client) *ActionHostIpAddressCreate {
	return &ActionHostIpAddressCreate{
		Client: client,
	}
}

// ActionHostIpAddressCreateMetaGlobalInput is a type for action global meta input parameters
type ActionHostIpAddressCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionHostIpAddressCreateMetaGlobalInput) SetIncludes(value string) *ActionHostIpAddressCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionHostIpAddressCreateMetaGlobalInput) SetNo(value bool) *ActionHostIpAddressCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionHostIpAddressCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionHostIpAddressCreateMetaGlobalInput) SelectParameters(params ...string) *ActionHostIpAddressCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionHostIpAddressCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionHostIpAddressCreateInput is a type for action input parameters
type ActionHostIpAddressCreateInput struct {
	Addr      string `json:"addr"`
	Assigned  bool   `json:"assigned"`
	IpAddress int64  `json:"ip_address"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAddr sets parameter Addr to value and selects it for sending
func (in *ActionHostIpAddressCreateInput) SetAddr(value string) *ActionHostIpAddressCreateInput {
	in.Addr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Addr"] = nil
	return in
}

// SetAssigned sets parameter Assigned to value and selects it for sending
func (in *ActionHostIpAddressCreateInput) SetAssigned(value bool) *ActionHostIpAddressCreateInput {
	in.Assigned = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Assigned"] = nil
	return in
}

// SetIpAddress sets parameter IpAddress to value and selects it for sending
func (in *ActionHostIpAddressCreateInput) SetIpAddress(value int64) *ActionHostIpAddressCreateInput {
	in.IpAddress = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetIpAddressNil(false)
	in._selectedParameters["IpAddress"] = nil
	return in
}

// SetIpAddressNil sets parameter IpAddress to nil and selects it for sending
func (in *ActionHostIpAddressCreateInput) SetIpAddressNil(set bool) *ActionHostIpAddressCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["IpAddress"] = nil
		in.SelectParameters("IpAddress")
	} else {
		delete(in._nilParameters, "IpAddress")
	}
	return in
}

// SelectParameters sets parameters from ActionHostIpAddressCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionHostIpAddressCreateInput) SelectParameters(params ...string) *ActionHostIpAddressCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionHostIpAddressCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionHostIpAddressCreateInput) UnselectParameters(params ...string) *ActionHostIpAddressCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionHostIpAddressCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionHostIpAddressCreateRequest is a type for the entire action request
type ActionHostIpAddressCreateRequest struct {
	HostIpAddress map[string]interface{} `json:"host_ip_address"`
	Meta          map[string]interface{} `json:"_meta"`
}

// ActionHostIpAddressCreateOutput is a type for action output parameters
type ActionHostIpAddressCreateOutput struct {
	Addr               string                     `json:"addr"`
	Assigned           bool                       `json:"assigned"`
	Id                 int64                      `json:"id"`
	IpAddress          *ActionIpAddressShowOutput `json:"ip_address"`
	ReverseRecordValue string                     `json:"reverse_record_value"`
	UserCreated        bool                       `json:"user_created"`
}

// Type for action response, including envelope
type ActionHostIpAddressCreateResponse struct {
	Action *ActionHostIpAddressCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		HostIpAddress *ActionHostIpAddressCreateOutput `json:"host_ip_address"`
	}

	// Action output without the namespace
	Output *ActionHostIpAddressCreateOutput
}

// Prepare the action for invocation
func (action *ActionHostIpAddressCreate) Prepare() *ActionHostIpAddressCreateInvocation {
	return &ActionHostIpAddressCreateInvocation{
		Action: action,
		Path:   "/v7.0/host_ip_addresses",
	}
}

// ActionHostIpAddressCreateInvocation is used to configure action for invocation
type ActionHostIpAddressCreateInvocation struct {
	// Pointer to the action
	Action *ActionHostIpAddressCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionHostIpAddressCreateInput
	// Global meta input parameters
	MetaInput *ActionHostIpAddressCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionHostIpAddressCreateInvocation) NewInput() *ActionHostIpAddressCreateInput {
	inv.Input = &ActionHostIpAddressCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionHostIpAddressCreateInvocation) SetInput(input *ActionHostIpAddressCreateInput) *ActionHostIpAddressCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionHostIpAddressCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionHostIpAddressCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionHostIpAddressCreateInvocation) NewMetaInput() *ActionHostIpAddressCreateMetaGlobalInput {
	inv.MetaInput = &ActionHostIpAddressCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionHostIpAddressCreateInvocation) SetMetaInput(input *ActionHostIpAddressCreateMetaGlobalInput) *ActionHostIpAddressCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionHostIpAddressCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionHostIpAddressCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionHostIpAddressCreateInvocation) Call() (*ActionHostIpAddressCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionHostIpAddressCreateInvocation) callAsBody() (*ActionHostIpAddressCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionHostIpAddressCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.HostIpAddress
	}
	return resp, err
}

func (inv *ActionHostIpAddressCreateInvocation) makeAllInputParams() *ActionHostIpAddressCreateRequest {
	return &ActionHostIpAddressCreateRequest{
		HostIpAddress: inv.makeInputParams(),
		Meta:          inv.makeMetaInputParams(),
	}
}

func (inv *ActionHostIpAddressCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Addr") {
			ret["addr"] = inv.Input.Addr
		}
		if inv.IsParameterSelected("Assigned") {
			ret["assigned"] = inv.Input.Assigned
		}
		if inv.IsParameterSelected("IpAddress") {
			if inv.IsParameterNil("IpAddress") {
				ret["ip_address"] = nil
			} else {
				ret["ip_address"] = inv.Input.IpAddress
			}
		}
	}

	return ret
}

func (inv *ActionHostIpAddressCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
