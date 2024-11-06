package client

import ()

// ActionDnsServerCreate is a type for action Dns_server#Create
type ActionDnsServerCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsServerCreate(client *Client) *ActionDnsServerCreate {
	return &ActionDnsServerCreate{
		Client: client,
	}
}

// ActionDnsServerCreateMetaGlobalInput is a type for action global meta input parameters
type ActionDnsServerCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsServerCreateMetaGlobalInput) SetIncludes(value string) *ActionDnsServerCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsServerCreateMetaGlobalInput) SetNo(value bool) *ActionDnsServerCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsServerCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsServerCreateMetaGlobalInput) SelectParameters(params ...string) *ActionDnsServerCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsServerCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsServerCreateInput is a type for action input parameters
type ActionDnsServerCreateInput struct {
	EnableUserDnsZones bool   `json:"enable_user_dns_zones"`
	Hidden             bool   `json:"hidden"`
	Ipv4Addr           string `json:"ipv4_addr"`
	Ipv6Addr           string `json:"ipv6_addr"`
	Name               string `json:"name"`
	Node               int64  `json:"node"`
	UserDnsZoneType    string `json:"user_dns_zone_type"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnableUserDnsZones sets parameter EnableUserDnsZones to value and selects it for sending
func (in *ActionDnsServerCreateInput) SetEnableUserDnsZones(value bool) *ActionDnsServerCreateInput {
	in.EnableUserDnsZones = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnableUserDnsZones"] = nil
	return in
}

// SetHidden sets parameter Hidden to value and selects it for sending
func (in *ActionDnsServerCreateInput) SetHidden(value bool) *ActionDnsServerCreateInput {
	in.Hidden = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Hidden"] = nil
	return in
}

// SetIpv4Addr sets parameter Ipv4Addr to value and selects it for sending
func (in *ActionDnsServerCreateInput) SetIpv4Addr(value string) *ActionDnsServerCreateInput {
	in.Ipv4Addr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Ipv4Addr"] = nil
	return in
}

// SetIpv6Addr sets parameter Ipv6Addr to value and selects it for sending
func (in *ActionDnsServerCreateInput) SetIpv6Addr(value string) *ActionDnsServerCreateInput {
	in.Ipv6Addr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Ipv6Addr"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionDnsServerCreateInput) SetName(value string) *ActionDnsServerCreateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionDnsServerCreateInput) SetNode(value int64) *ActionDnsServerCreateInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNodeNil(false)
	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeNil sets parameter Node to nil and selects it for sending
func (in *ActionDnsServerCreateInput) SetNodeNil(set bool) *ActionDnsServerCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Node"] = nil
		in.SelectParameters("Node")
	} else {
		delete(in._nilParameters, "Node")
	}
	return in
}

// SetUserDnsZoneType sets parameter UserDnsZoneType to value and selects it for sending
func (in *ActionDnsServerCreateInput) SetUserDnsZoneType(value string) *ActionDnsServerCreateInput {
	in.UserDnsZoneType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserDnsZoneType"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsServerCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsServerCreateInput) SelectParameters(params ...string) *ActionDnsServerCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnsServerCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnsServerCreateInput) UnselectParameters(params ...string) *ActionDnsServerCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnsServerCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsServerCreateRequest is a type for the entire action request
type ActionDnsServerCreateRequest struct {
	DnsServer map[string]interface{} `json:"dns_server"`
	Meta      map[string]interface{} `json:"_meta"`
}

// ActionDnsServerCreateOutput is a type for action output parameters
type ActionDnsServerCreateOutput struct {
	CreatedAt          string                `json:"created_at"`
	EnableUserDnsZones bool                  `json:"enable_user_dns_zones"`
	Hidden             bool                  `json:"hidden"`
	Id                 int64                 `json:"id"`
	Ipv4Addr           string                `json:"ipv4_addr"`
	Ipv6Addr           string                `json:"ipv6_addr"`
	Name               string                `json:"name"`
	Node               *ActionNodeShowOutput `json:"node"`
	UpdatedAt          string                `json:"updated_at"`
	UserDnsZoneType    string                `json:"user_dns_zone_type"`
}

// Type for action response, including envelope
type ActionDnsServerCreateResponse struct {
	Action *ActionDnsServerCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsServer *ActionDnsServerCreateOutput `json:"dns_server"`
	}

	// Action output without the namespace
	Output *ActionDnsServerCreateOutput
}

// Prepare the action for invocation
func (action *ActionDnsServerCreate) Prepare() *ActionDnsServerCreateInvocation {
	return &ActionDnsServerCreateInvocation{
		Action: action,
		Path:   "/v7.0/dns_servers",
	}
}

// ActionDnsServerCreateInvocation is used to configure action for invocation
type ActionDnsServerCreateInvocation struct {
	// Pointer to the action
	Action *ActionDnsServerCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsServerCreateInput
	// Global meta input parameters
	MetaInput *ActionDnsServerCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsServerCreateInvocation) NewInput() *ActionDnsServerCreateInput {
	inv.Input = &ActionDnsServerCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsServerCreateInvocation) SetInput(input *ActionDnsServerCreateInput) *ActionDnsServerCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsServerCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnsServerCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsServerCreateInvocation) NewMetaInput() *ActionDnsServerCreateMetaGlobalInput {
	inv.MetaInput = &ActionDnsServerCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsServerCreateInvocation) SetMetaInput(input *ActionDnsServerCreateMetaGlobalInput) *ActionDnsServerCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsServerCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsServerCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsServerCreateInvocation) Call() (*ActionDnsServerCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDnsServerCreateInvocation) callAsBody() (*ActionDnsServerCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDnsServerCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsServer
	}
	return resp, err
}

func (inv *ActionDnsServerCreateInvocation) makeAllInputParams() *ActionDnsServerCreateRequest {
	return &ActionDnsServerCreateRequest{
		DnsServer: inv.makeInputParams(),
		Meta:      inv.makeMetaInputParams(),
	}
}

func (inv *ActionDnsServerCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("EnableUserDnsZones") {
			ret["enable_user_dns_zones"] = inv.Input.EnableUserDnsZones
		}
		if inv.IsParameterSelected("Hidden") {
			ret["hidden"] = inv.Input.Hidden
		}
		if inv.IsParameterSelected("Ipv4Addr") {
			ret["ipv4_addr"] = inv.Input.Ipv4Addr
		}
		if inv.IsParameterSelected("Ipv6Addr") {
			ret["ipv6_addr"] = inv.Input.Ipv6Addr
		}
		if inv.IsParameterSelected("Name") {
			ret["name"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Node") {
			if inv.IsParameterNil("Node") {
				ret["node"] = nil
			} else {
				ret["node"] = inv.Input.Node
			}
		}
		if inv.IsParameterSelected("UserDnsZoneType") {
			ret["user_dns_zone_type"] = inv.Input.UserDnsZoneType
		}
	}

	return ret
}

func (inv *ActionDnsServerCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
