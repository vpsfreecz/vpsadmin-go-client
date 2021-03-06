package client

import (
)

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
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	IpAddress int64 `json:"ip_address"`
	NetworkInterface int64 `json:"network_interface"`
	Vps int64 `json:"vps"`
	Version int64 `json:"version"`
	Network int64 `json:"network"`
	Location int64 `json:"location"`
	User int64 `json:"user"`
	Role string `json:"role"`
	Addr string `json:"addr"`
	Prefix int64 `json:"prefix"`
	Size int64 `json:"size"`
	MaxTx int64 `json:"max_tx"`
	MaxRx int64 `json:"max_rx"`
	Assigned bool `json:"assigned"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetLimit(value int64) *ActionHostIpAddressIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetIpAddress sets parameter IpAddress to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetIpAddress(value int64) *ActionHostIpAddressIndexInput {
	in.IpAddress = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IpAddress"] = nil
	return in
}
// SetNetworkInterface sets parameter NetworkInterface to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetNetworkInterface(value int64) *ActionHostIpAddressIndexInput {
	in.NetworkInterface = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NetworkInterface"] = nil
	return in
}
// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetVps(value int64) *ActionHostIpAddressIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Vps"] = nil
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
// SetNetwork sets parameter Network to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetNetwork(value int64) *ActionHostIpAddressIndexInput {
	in.Network = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Network"] = nil
	return in
}
// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetLocation(value int64) *ActionHostIpAddressIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}
// SetUser sets parameter User to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetUser(value int64) *ActionHostIpAddressIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
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
// SetAddr sets parameter Addr to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetAddr(value string) *ActionHostIpAddressIndexInput {
	in.Addr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Addr"] = nil
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
// SetSize sets parameter Size to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetSize(value int64) *ActionHostIpAddressIndexInput {
	in.Size = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Size"] = nil
	return in
}
// SetMaxTx sets parameter MaxTx to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetMaxTx(value int64) *ActionHostIpAddressIndexInput {
	in.MaxTx = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxTx"] = nil
	return in
}
// SetMaxRx sets parameter MaxRx to value and selects it for sending
func (in *ActionHostIpAddressIndexInput) SetMaxRx(value int64) *ActionHostIpAddressIndexInput {
	in.MaxRx = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxRx"] = nil
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

func (in *ActionHostIpAddressIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionHostIpAddressIndexOutput is a type for action output parameters
type ActionHostIpAddressIndexOutput struct {
	Id int64 `json:"id"`
	IpAddress *ActionIpAddressShowOutput `json:"ip_address"`
	Addr string `json:"addr"`
	Assigned bool `json:"assigned"`
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
		Path: "/v6.0/host_ip_addresses",
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
		if inv.IsParameterSelected("Offset") {
			ret["host_ip_address[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["host_ip_address[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("IpAddress") {
			ret["host_ip_address[ip_address]"] = convertInt64ToString(inv.Input.IpAddress)
		}
		if inv.IsParameterSelected("NetworkInterface") {
			ret["host_ip_address[network_interface]"] = convertInt64ToString(inv.Input.NetworkInterface)
		}
		if inv.IsParameterSelected("Vps") {
			ret["host_ip_address[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
		if inv.IsParameterSelected("Version") {
			ret["host_ip_address[version]"] = convertInt64ToString(inv.Input.Version)
		}
		if inv.IsParameterSelected("Network") {
			ret["host_ip_address[network]"] = convertInt64ToString(inv.Input.Network)
		}
		if inv.IsParameterSelected("Location") {
			ret["host_ip_address[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("User") {
			ret["host_ip_address[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("Role") {
			ret["host_ip_address[role]"] = inv.Input.Role
		}
		if inv.IsParameterSelected("Addr") {
			ret["host_ip_address[addr]"] = inv.Input.Addr
		}
		if inv.IsParameterSelected("Prefix") {
			ret["host_ip_address[prefix]"] = convertInt64ToString(inv.Input.Prefix)
		}
		if inv.IsParameterSelected("Size") {
			ret["host_ip_address[size]"] = convertInt64ToString(inv.Input.Size)
		}
		if inv.IsParameterSelected("MaxTx") {
			ret["host_ip_address[max_tx]"] = convertInt64ToString(inv.Input.MaxTx)
		}
		if inv.IsParameterSelected("MaxRx") {
			ret["host_ip_address[max_rx]"] = convertInt64ToString(inv.Input.MaxRx)
		}
		if inv.IsParameterSelected("Assigned") {
			ret["host_ip_address[assigned]"] = convertBoolToString(inv.Input.Assigned)
		}
	}
}

func (inv *ActionHostIpAddressIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

