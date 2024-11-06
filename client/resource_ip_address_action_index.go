package client

import ()

// ActionIpAddressIndex is a type for action Ip_address#Index
type ActionIpAddressIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIpAddressIndex(client *Client) *ActionIpAddressIndex {
	return &ActionIpAddressIndex{
		Client: client,
	}
}

// ActionIpAddressIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIpAddressIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionIpAddressIndexMetaGlobalInput) SetCount(value bool) *ActionIpAddressIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpAddressIndexMetaGlobalInput) SetIncludes(value string) *ActionIpAddressIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpAddressIndexMetaGlobalInput) SetNo(value bool) *ActionIpAddressIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpAddressIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpAddressIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIpAddressIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpAddressIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpAddressIndexInput is a type for action input parameters
type ActionIpAddressIndexInput struct {
	Addr                string `json:"addr"`
	AssignedToInterface bool   `json:"assigned_to_interface"`
	FromId              int64  `json:"from_id"`
	Limit               int64  `json:"limit"`
	Location            int64  `json:"location"`
	Network             int64  `json:"network"`
	NetworkInterface    int64  `json:"network_interface"`
	Order               string `json:"order"`
	Prefix              int64  `json:"prefix"`
	Purpose             string `json:"purpose"`
	Role                string `json:"role"`
	Size                int64  `json:"size"`
	User                int64  `json:"user"`
	Version             int64  `json:"version"`
	Vps                 int64  `json:"vps"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAddr sets parameter Addr to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetAddr(value string) *ActionIpAddressIndexInput {
	in.Addr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Addr"] = nil
	return in
}

// SetAssignedToInterface sets parameter AssignedToInterface to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetAssignedToInterface(value bool) *ActionIpAddressIndexInput {
	in.AssignedToInterface = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AssignedToInterface"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetFromId(value int64) *ActionIpAddressIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetLimit(value int64) *ActionIpAddressIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetLocation(value int64) *ActionIpAddressIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLocationNil(false)
	in._selectedParameters["Location"] = nil
	return in
}

// SetLocationNil sets parameter Location to nil and selects it for sending
func (in *ActionIpAddressIndexInput) SetLocationNil(set bool) *ActionIpAddressIndexInput {
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
func (in *ActionIpAddressIndexInput) SetNetwork(value int64) *ActionIpAddressIndexInput {
	in.Network = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNetworkNil(false)
	in._selectedParameters["Network"] = nil
	return in
}

// SetNetworkNil sets parameter Network to nil and selects it for sending
func (in *ActionIpAddressIndexInput) SetNetworkNil(set bool) *ActionIpAddressIndexInput {
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

// SetNetworkInterface sets parameter NetworkInterface to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetNetworkInterface(value int64) *ActionIpAddressIndexInput {
	in.NetworkInterface = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNetworkInterfaceNil(false)
	in._selectedParameters["NetworkInterface"] = nil
	return in
}

// SetNetworkInterfaceNil sets parameter NetworkInterface to nil and selects it for sending
func (in *ActionIpAddressIndexInput) SetNetworkInterfaceNil(set bool) *ActionIpAddressIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["NetworkInterface"] = nil
		in.SelectParameters("NetworkInterface")
	} else {
		delete(in._nilParameters, "NetworkInterface")
	}
	return in
}

// SetOrder sets parameter Order to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetOrder(value string) *ActionIpAddressIndexInput {
	in.Order = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Order"] = nil
	return in
}

// SetPrefix sets parameter Prefix to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetPrefix(value int64) *ActionIpAddressIndexInput {
	in.Prefix = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Prefix"] = nil
	return in
}

// SetPurpose sets parameter Purpose to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetPurpose(value string) *ActionIpAddressIndexInput {
	in.Purpose = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Purpose"] = nil
	return in
}

// SetRole sets parameter Role to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetRole(value string) *ActionIpAddressIndexInput {
	in.Role = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Role"] = nil
	return in
}

// SetSize sets parameter Size to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetSize(value int64) *ActionIpAddressIndexInput {
	in.Size = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Size"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetUser(value int64) *ActionIpAddressIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionIpAddressIndexInput) SetUserNil(set bool) *ActionIpAddressIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["User"] = nil
		in.SelectParameters("User")
	} else {
		delete(in._nilParameters, "User")
	}
	return in
}

// SetVersion sets parameter Version to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetVersion(value int64) *ActionIpAddressIndexInput {
	in.Version = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Version"] = nil
	return in
}

// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetVps(value int64) *ActionIpAddressIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetVpsNil(false)
	in._selectedParameters["Vps"] = nil
	return in
}

// SetVpsNil sets parameter Vps to nil and selects it for sending
func (in *ActionIpAddressIndexInput) SetVpsNil(set bool) *ActionIpAddressIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Vps"] = nil
		in.SelectParameters("Vps")
	} else {
		delete(in._nilParameters, "Vps")
	}
	return in
}

// SelectParameters sets parameters from ActionIpAddressIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpAddressIndexInput) SelectParameters(params ...string) *ActionIpAddressIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIpAddressIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIpAddressIndexInput) UnselectParameters(params ...string) *ActionIpAddressIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIpAddressIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpAddressIndexOutput is a type for action output parameters
type ActionIpAddressIndexOutput struct {
	Addr               string                            `json:"addr"`
	ChargedEnvironment *ActionEnvironmentShowOutput      `json:"charged_environment"`
	Id                 int64                             `json:"id"`
	Network            *ActionNetworkShowOutput          `json:"network"`
	NetworkInterface   *ActionNetworkInterfaceShowOutput `json:"network_interface"`
	Prefix             int64                             `json:"prefix"`
	RouteVia           *ActionHostIpAddressShowOutput    `json:"route_via"`
	Size               int64                             `json:"size"`
	User               *ActionUserShowOutput             `json:"user"`
}

// Type for action response, including envelope
type ActionIpAddressIndexResponse struct {
	Action *ActionIpAddressIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpAddresses []*ActionIpAddressIndexOutput `json:"ip_addresses"`
	}

	// Action output without the namespace
	Output []*ActionIpAddressIndexOutput
}

// Prepare the action for invocation
func (action *ActionIpAddressIndex) Prepare() *ActionIpAddressIndexInvocation {
	return &ActionIpAddressIndexInvocation{
		Action: action,
		Path:   "/v7.0/ip_addresses",
	}
}

// ActionIpAddressIndexInvocation is used to configure action for invocation
type ActionIpAddressIndexInvocation struct {
	// Pointer to the action
	Action *ActionIpAddressIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpAddressIndexInput
	// Global meta input parameters
	MetaInput *ActionIpAddressIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpAddressIndexInvocation) NewInput() *ActionIpAddressIndexInput {
	inv.Input = &ActionIpAddressIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpAddressIndexInvocation) SetInput(input *ActionIpAddressIndexInput) *ActionIpAddressIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpAddressIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIpAddressIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpAddressIndexInvocation) NewMetaInput() *ActionIpAddressIndexMetaGlobalInput {
	inv.MetaInput = &ActionIpAddressIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpAddressIndexInvocation) SetMetaInput(input *ActionIpAddressIndexMetaGlobalInput) *ActionIpAddressIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpAddressIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpAddressIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIpAddressIndexInvocation) Call() (*ActionIpAddressIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionIpAddressIndexInvocation) callAsQuery() (*ActionIpAddressIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIpAddressIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IpAddresses
	}
	return resp, err
}

func (inv *ActionIpAddressIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Addr") {
			ret["ip_address[addr]"] = inv.Input.Addr
		}
		if inv.IsParameterSelected("AssignedToInterface") {
			ret["ip_address[assigned_to_interface]"] = convertBoolToString(inv.Input.AssignedToInterface)
		}
		if inv.IsParameterSelected("FromId") {
			ret["ip_address[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["ip_address[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Location") {
			ret["ip_address[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("Network") {
			ret["ip_address[network]"] = convertInt64ToString(inv.Input.Network)
		}
		if inv.IsParameterSelected("NetworkInterface") {
			ret["ip_address[network_interface]"] = convertInt64ToString(inv.Input.NetworkInterface)
		}
		if inv.IsParameterSelected("Order") {
			ret["ip_address[order]"] = inv.Input.Order
		}
		if inv.IsParameterSelected("Prefix") {
			ret["ip_address[prefix]"] = convertInt64ToString(inv.Input.Prefix)
		}
		if inv.IsParameterSelected("Purpose") {
			ret["ip_address[purpose]"] = inv.Input.Purpose
		}
		if inv.IsParameterSelected("Role") {
			ret["ip_address[role]"] = inv.Input.Role
		}
		if inv.IsParameterSelected("Size") {
			ret["ip_address[size]"] = convertInt64ToString(inv.Input.Size)
		}
		if inv.IsParameterSelected("User") {
			ret["ip_address[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("Version") {
			ret["ip_address[version]"] = convertInt64ToString(inv.Input.Version)
		}
		if inv.IsParameterSelected("Vps") {
			ret["ip_address[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
	}
}

func (inv *ActionIpAddressIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
