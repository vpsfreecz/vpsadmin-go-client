package client

import ()

// ActionHostIpAddressIndex is a type for action Host_ip_address#Index
type ActionHostIpAddressIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionHostIpAddressIndex(client *Client) *ActionHostIpAddressIndex {
	return &ActionHostIpAddressIndex{
		Client: client,
	}
}

// ActionHostIpAddressIndexMetaGlobalInput is a type for action global meta input parameters
type ActionHostIpAddressIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionHostIpAddressIndexMetaGlobalInput) SetCount(value bool) *ActionHostIpAddressIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionHostIpAddressIndexMetaGlobalInput) SetIncludes(value string) *ActionHostIpAddressIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionHostIpAddressIndexMetaGlobalInput) SetNo(value bool) *ActionHostIpAddressIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionHostIpAddressIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionHostIpAddressIndexMetaGlobalInput) SelectParameters(params ...string) *ActionHostIpAddressIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionHostIpAddressIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionHostIpAddressIndexInput is a type for action input parameters
type ActionHostIpAddressIndexInput struct {
	Addr             string `json:"addr"`
	Assigned         bool   `json:"assigned"`
	IpAddress        int64  `json:"ip_address"`
	Limit            int64  `json:"limit"`
	Location         int64  `json:"location"`
	Network          int64  `json:"network"`
	NetworkInterface int64  `json:"network_interface"`
	Offset           int64  `json:"offset"`
	Order            string `json:"order"`
	Prefix           int64  `json:"prefix"`
	Purpose          string `json:"purpose"`
	Role             string `json:"role"`
	Size             int64  `json:"size"`
	User             int64  `json:"user"`
	Version          int64  `json:"version"`
	Vps              int64  `json:"vps"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAddr sets parameter Addr to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetAddr(value string) *ActionHostIpAddressIndexInput {
	in.Addr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Addr"] = nil
	return in
}

// SetAssigned sets parameter Assigned to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetAssigned(value bool) *ActionHostIpAddressIndexInput {
	in.Assigned = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Assigned"] = nil
	return in
}

// SetIpAddress sets parameter IpAddress to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetIpAddress(value int64) *ActionHostIpAddressIndexInput {
	in.IpAddress = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetIpAddressNil(false)
	in._selectedParameters["IpAddress"] = nil
	return in
}

// SetIpAddressNil sets parameter IpAddress to nil and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetIpAddressNil(set bool) *ActionHostIpAddressIndexInput {
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

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetLimit(value int64) *ActionHostIpAddressIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetLocation(value int64) *ActionHostIpAddressIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLocationNil(false)
	in._selectedParameters["Location"] = nil
	return in
}

// SetLocationNil sets parameter Location to nil and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetLocationNil(set bool) *ActionHostIpAddressIndexInput {
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
func (in *ActionHostIpAddressIndexInput) SetNetwork(value int64) *ActionHostIpAddressIndexInput {
	in.Network = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNetworkNil(false)
	in._selectedParameters["Network"] = nil
	return in
}

// SetNetworkNil sets parameter Network to nil and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetNetworkNil(set bool) *ActionHostIpAddressIndexInput {
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
func (in *ActionHostIpAddressIndexInput) SetNetworkInterface(value int64) *ActionHostIpAddressIndexInput {
	in.NetworkInterface = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNetworkInterfaceNil(false)
	in._selectedParameters["NetworkInterface"] = nil
	return in
}

// SetNetworkInterfaceNil sets parameter NetworkInterface to nil and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetNetworkInterfaceNil(set bool) *ActionHostIpAddressIndexInput {
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

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetOffset(value int64) *ActionHostIpAddressIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetOrder sets parameter Order to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetOrder(value string) *ActionHostIpAddressIndexInput {
	in.Order = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Order"] = nil
	return in
}

// SetPrefix sets parameter Prefix to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetPrefix(value int64) *ActionHostIpAddressIndexInput {
	in.Prefix = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Prefix"] = nil
	return in
}

// SetPurpose sets parameter Purpose to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetPurpose(value string) *ActionHostIpAddressIndexInput {
	in.Purpose = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Purpose"] = nil
	return in
}

// SetRole sets parameter Role to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetRole(value string) *ActionHostIpAddressIndexInput {
	in.Role = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Role"] = nil
	return in
}

// SetSize sets parameter Size to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetSize(value int64) *ActionHostIpAddressIndexInput {
	in.Size = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Size"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetUser(value int64) *ActionHostIpAddressIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetUserNil(set bool) *ActionHostIpAddressIndexInput {
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
func (in *ActionHostIpAddressIndexInput) SetVersion(value int64) *ActionHostIpAddressIndexInput {
	in.Version = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Version"] = nil
	return in
}

// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetVps(value int64) *ActionHostIpAddressIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetVpsNil(false)
	in._selectedParameters["Vps"] = nil
	return in
}

// SetVpsNil sets parameter Vps to nil and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetVpsNil(set bool) *ActionHostIpAddressIndexInput {
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

// SelectParameters sets parameters from ActionHostIpAddressIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionHostIpAddressIndexInput) SelectParameters(params ...string) *ActionHostIpAddressIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionHostIpAddressIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionHostIpAddressIndexInput) UnselectParameters(params ...string) *ActionHostIpAddressIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionHostIpAddressIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionHostIpAddressIndexOutput is a type for action output parameters
type ActionHostIpAddressIndexOutput struct {
	Addr      string                     `json:"addr"`
	Assigned  bool                       `json:"assigned"`
	Id        int64                      `json:"id"`
	IpAddress *ActionIpAddressShowOutput `json:"ip_address"`
}

// Type for action response, including envelope
type ActionHostIpAddressIndexResponse struct {
	Action *ActionHostIpAddressIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		HostIpAddresses []*ActionHostIpAddressIndexOutput `json:"host_ip_addresses"`
	}

	// Action output without the namespace
	Output []*ActionHostIpAddressIndexOutput
}

// Prepare the action for invocation
func (action *ActionHostIpAddressIndex) Prepare() *ActionHostIpAddressIndexInvocation {
	return &ActionHostIpAddressIndexInvocation{
		Action: action,
		Path:   "/v6.0/host_ip_addresses",
	}
}

// ActionHostIpAddressIndexInvocation is used to configure action for invocation
type ActionHostIpAddressIndexInvocation struct {
	// Pointer to the action
	Action *ActionHostIpAddressIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionHostIpAddressIndexInput
	// Global meta input parameters
	MetaInput *ActionHostIpAddressIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionHostIpAddressIndexInvocation) NewInput() *ActionHostIpAddressIndexInput {
	inv.Input = &ActionHostIpAddressIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionHostIpAddressIndexInvocation) SetInput(input *ActionHostIpAddressIndexInput) *ActionHostIpAddressIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionHostIpAddressIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionHostIpAddressIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionHostIpAddressIndexInvocation) NewMetaInput() *ActionHostIpAddressIndexMetaGlobalInput {
	inv.MetaInput = &ActionHostIpAddressIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionHostIpAddressIndexInvocation) SetMetaInput(input *ActionHostIpAddressIndexMetaGlobalInput) *ActionHostIpAddressIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionHostIpAddressIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionHostIpAddressIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionHostIpAddressIndexInvocation) Call() (*ActionHostIpAddressIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionHostIpAddressIndexInvocation) callAsQuery() (*ActionHostIpAddressIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionHostIpAddressIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.HostIpAddresses
	}
	return resp, err
}

func (inv *ActionHostIpAddressIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Addr") {
			ret["host_ip_address[addr]"] = inv.Input.Addr
		}
		if inv.IsParameterSelected("Assigned") {
			ret["host_ip_address[assigned]"] = convertBoolToString(inv.Input.Assigned)
		}
		if inv.IsParameterSelected("IpAddress") {
			ret["host_ip_address[ip_address]"] = convertInt64ToString(inv.Input.IpAddress)
		}
		if inv.IsParameterSelected("Limit") {
			ret["host_ip_address[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Location") {
			ret["host_ip_address[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("Network") {
			ret["host_ip_address[network]"] = convertInt64ToString(inv.Input.Network)
		}
		if inv.IsParameterSelected("NetworkInterface") {
			ret["host_ip_address[network_interface]"] = convertInt64ToString(inv.Input.NetworkInterface)
		}
		if inv.IsParameterSelected("Offset") {
			ret["host_ip_address[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Order") {
			ret["host_ip_address[order]"] = inv.Input.Order
		}
		if inv.IsParameterSelected("Prefix") {
			ret["host_ip_address[prefix]"] = convertInt64ToString(inv.Input.Prefix)
		}
		if inv.IsParameterSelected("Purpose") {
			ret["host_ip_address[purpose]"] = inv.Input.Purpose
		}
		if inv.IsParameterSelected("Role") {
			ret["host_ip_address[role]"] = inv.Input.Role
		}
		if inv.IsParameterSelected("Size") {
			ret["host_ip_address[size]"] = convertInt64ToString(inv.Input.Size)
		}
		if inv.IsParameterSelected("User") {
			ret["host_ip_address[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("Version") {
			ret["host_ip_address[version]"] = convertInt64ToString(inv.Input.Version)
		}
		if inv.IsParameterSelected("Vps") {
			ret["host_ip_address[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
	}
}

func (inv *ActionHostIpAddressIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
