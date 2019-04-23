package client

import (
)

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
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
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
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetOffset(value int64) *ActionIpAddressIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
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
// SetNetworkInterface sets parameter NetworkInterface to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetNetworkInterface(value int64) *ActionIpAddressIndexInput {
	in.NetworkInterface = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NetworkInterface"] = nil
	return in
}
// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetVps(value int64) *ActionIpAddressIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Vps"] = nil
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
// SetNetwork sets parameter Network to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetNetwork(value int64) *ActionIpAddressIndexInput {
	in.Network = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Network"] = nil
	return in
}
// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetLocation(value int64) *ActionIpAddressIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}
// SetUser sets parameter User to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetUser(value int64) *ActionIpAddressIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
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
// SetAddr sets parameter Addr to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetAddr(value string) *ActionIpAddressIndexInput {
	in.Addr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Addr"] = nil
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
// SetSize sets parameter Size to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetSize(value int64) *ActionIpAddressIndexInput {
	in.Size = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Size"] = nil
	return in
}
// SetMaxTx sets parameter MaxTx to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetMaxTx(value int64) *ActionIpAddressIndexInput {
	in.MaxTx = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxTx"] = nil
	return in
}
// SetMaxRx sets parameter MaxRx to value and selects it for sending
func (in *ActionIpAddressIndexInput) SetMaxRx(value int64) *ActionIpAddressIndexInput {
	in.MaxRx = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxRx"] = nil
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

func (in *ActionIpAddressIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionIpAddressIndexOutput is a type for action output parameters
type ActionIpAddressIndexOutput struct {
	Id int64 `json:"id"`
	NetworkInterface *ActionNetworkInterfaceShowOutput `json:"network_interface"`
	Network *ActionNetworkShowOutput `json:"network"`
	User *ActionUserShowOutput `json:"user"`
	Addr string `json:"addr"`
	Prefix int64 `json:"prefix"`
	Size int64 `json:"size"`
	RouteVia *ActionHostIpAddressShowOutput `json:"route_via"`
	MaxTx int64 `json:"max_tx"`
	MaxRx int64 `json:"max_rx"`
	ClassId int64 `json:"class_id"`
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
		Path: "/v5.0/ip_addresses",
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
		if inv.IsParameterSelected("Offset") {
			ret["ip_address[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["ip_address[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("NetworkInterface") {
			ret["ip_address[network_interface]"] = convertInt64ToString(inv.Input.NetworkInterface)
		}
		if inv.IsParameterSelected("Vps") {
			ret["ip_address[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
		if inv.IsParameterSelected("Version") {
			ret["ip_address[version]"] = convertInt64ToString(inv.Input.Version)
		}
		if inv.IsParameterSelected("Network") {
			ret["ip_address[network]"] = convertInt64ToString(inv.Input.Network)
		}
		if inv.IsParameterSelected("Location") {
			ret["ip_address[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("User") {
			ret["ip_address[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("Role") {
			ret["ip_address[role]"] = inv.Input.Role
		}
		if inv.IsParameterSelected("Addr") {
			ret["ip_address[addr]"] = inv.Input.Addr
		}
		if inv.IsParameterSelected("Prefix") {
			ret["ip_address[prefix]"] = convertInt64ToString(inv.Input.Prefix)
		}
		if inv.IsParameterSelected("Size") {
			ret["ip_address[size]"] = convertInt64ToString(inv.Input.Size)
		}
		if inv.IsParameterSelected("MaxTx") {
			ret["ip_address[max_tx]"] = convertInt64ToString(inv.Input.MaxTx)
		}
		if inv.IsParameterSelected("MaxRx") {
			ret["ip_address[max_rx]"] = convertInt64ToString(inv.Input.MaxRx)
		}
	}
}

func (inv *ActionIpAddressIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

