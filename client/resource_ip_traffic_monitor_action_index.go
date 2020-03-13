package client

import (
)

// ActionIpTrafficMonitorIndex is a type for action Ip_traffic_monitor#Index
type ActionIpTrafficMonitorIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIpTrafficMonitorIndex(client *Client) *ActionIpTrafficMonitorIndex {
	return &ActionIpTrafficMonitorIndex{
		Client: client,
	}
}

// ActionIpTrafficMonitorIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIpTrafficMonitorIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpTrafficMonitorIndexMetaGlobalInput) SetNo(value bool) *ActionIpTrafficMonitorIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionIpTrafficMonitorIndexMetaGlobalInput) SetCount(value bool) *ActionIpTrafficMonitorIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpTrafficMonitorIndexMetaGlobalInput) SetIncludes(value string) *ActionIpTrafficMonitorIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpTrafficMonitorIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpTrafficMonitorIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIpTrafficMonitorIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpTrafficMonitorIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpTrafficMonitorIndexInput is a type for action input parameters
type ActionIpTrafficMonitorIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	IpAddress int64 `json:"ip_address"`
	User int64 `json:"user"`
	IpVersion int64 `json:"ip_version"`
	Environment int64 `json:"environment"`
	Location int64 `json:"location"`
	Network int64 `json:"network"`
	Node int64 `json:"node"`
	Vps int64 `json:"vps"`
	Order string `json:"order"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionIpTrafficMonitorIndexInput) SetOffset(value int64) *ActionIpTrafficMonitorIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIpTrafficMonitorIndexInput) SetLimit(value int64) *ActionIpTrafficMonitorIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetIpAddress sets parameter IpAddress to value and selects it for sending
func (in *ActionIpTrafficMonitorIndexInput) SetIpAddress(value int64) *ActionIpTrafficMonitorIndexInput {
	in.IpAddress = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IpAddress"] = nil
	return in
}
// SetUser sets parameter User to value and selects it for sending
func (in *ActionIpTrafficMonitorIndexInput) SetUser(value int64) *ActionIpTrafficMonitorIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}
// SetIpVersion sets parameter IpVersion to value and selects it for sending
func (in *ActionIpTrafficMonitorIndexInput) SetIpVersion(value int64) *ActionIpTrafficMonitorIndexInput {
	in.IpVersion = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IpVersion"] = nil
	return in
}
// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionIpTrafficMonitorIndexInput) SetEnvironment(value int64) *ActionIpTrafficMonitorIndexInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Environment"] = nil
	return in
}
// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionIpTrafficMonitorIndexInput) SetLocation(value int64) *ActionIpTrafficMonitorIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}
// SetNetwork sets parameter Network to value and selects it for sending
func (in *ActionIpTrafficMonitorIndexInput) SetNetwork(value int64) *ActionIpTrafficMonitorIndexInput {
	in.Network = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Network"] = nil
	return in
}
// SetNode sets parameter Node to value and selects it for sending
func (in *ActionIpTrafficMonitorIndexInput) SetNode(value int64) *ActionIpTrafficMonitorIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}
// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionIpTrafficMonitorIndexInput) SetVps(value int64) *ActionIpTrafficMonitorIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Vps"] = nil
	return in
}
// SetOrder sets parameter Order to value and selects it for sending
func (in *ActionIpTrafficMonitorIndexInput) SetOrder(value string) *ActionIpTrafficMonitorIndexInput {
	in.Order = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Order"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpTrafficMonitorIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpTrafficMonitorIndexInput) SelectParameters(params ...string) *ActionIpTrafficMonitorIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpTrafficMonitorIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionIpTrafficMonitorIndexOutput is a type for action output parameters
type ActionIpTrafficMonitorIndexOutput struct {
	Id int64 `json:"id"`
	IpAddress *ActionIpAddressShowOutput `json:"ip_address"`
	User *ActionUserShowOutput `json:"user"`
	Packets int64 `json:"packets"`
	PacketsIn int64 `json:"packets_in"`
	PacketsOut int64 `json:"packets_out"`
	Bytes int64 `json:"bytes"`
	BytesIn int64 `json:"bytes_in"`
	BytesOut int64 `json:"bytes_out"`
	PublicPackets int64 `json:"public_packets"`
	PublicPacketsIn int64 `json:"public_packets_in"`
	PublicPacketsOut int64 `json:"public_packets_out"`
	PublicBytes int64 `json:"public_bytes"`
	PublicBytesIn int64 `json:"public_bytes_in"`
	PublicBytesOut int64 `json:"public_bytes_out"`
	PublicTcpPackets int64 `json:"public_tcp_packets"`
	PublicTcpPacketsIn int64 `json:"public_tcp_packets_in"`
	PublicTcpPacketsOut int64 `json:"public_tcp_packets_out"`
	PublicTcpBytes int64 `json:"public_tcp_bytes"`
	PublicTcpBytesIn int64 `json:"public_tcp_bytes_in"`
	PublicTcpBytesOut int64 `json:"public_tcp_bytes_out"`
	PublicUdpPackets int64 `json:"public_udp_packets"`
	PublicUdpPacketsIn int64 `json:"public_udp_packets_in"`
	PublicUdpPacketsOut int64 `json:"public_udp_packets_out"`
	PublicUdpBytes int64 `json:"public_udp_bytes"`
	PublicUdpBytesIn int64 `json:"public_udp_bytes_in"`
	PublicUdpBytesOut int64 `json:"public_udp_bytes_out"`
	PublicOtherPackets int64 `json:"public_other_packets"`
	PublicOtherPacketsIn int64 `json:"public_other_packets_in"`
	PublicOtherPacketsOut int64 `json:"public_other_packets_out"`
	PublicOtherBytes int64 `json:"public_other_bytes"`
	PublicOtherBytesIn int64 `json:"public_other_bytes_in"`
	PublicOtherBytesOut int64 `json:"public_other_bytes_out"`
	PrivatePackets int64 `json:"private_packets"`
	PrivatePacketsIn int64 `json:"private_packets_in"`
	PrivatePacketsOut int64 `json:"private_packets_out"`
	PrivateBytes int64 `json:"private_bytes"`
	PrivateBytesIn int64 `json:"private_bytes_in"`
	PrivateBytesOut int64 `json:"private_bytes_out"`
	PrivateTcpPackets int64 `json:"private_tcp_packets"`
	PrivateTcpPacketsIn int64 `json:"private_tcp_packets_in"`
	PrivateTcpPacketsOut int64 `json:"private_tcp_packets_out"`
	PrivateTcpBytes int64 `json:"private_tcp_bytes"`
	PrivateTcpBytesIn int64 `json:"private_tcp_bytes_in"`
	PrivateTcpBytesOut int64 `json:"private_tcp_bytes_out"`
	PrivateUdpPackets int64 `json:"private_udp_packets"`
	PrivateUdpPacketsIn int64 `json:"private_udp_packets_in"`
	PrivateUdpPacketsOut int64 `json:"private_udp_packets_out"`
	PrivateUdpBytes int64 `json:"private_udp_bytes"`
	PrivateUdpBytesIn int64 `json:"private_udp_bytes_in"`
	PrivateUdpBytesOut int64 `json:"private_udp_bytes_out"`
	PrivateOtherPackets int64 `json:"private_other_packets"`
	PrivateOtherPacketsIn int64 `json:"private_other_packets_in"`
	PrivateOtherPacketsOut int64 `json:"private_other_packets_out"`
	PrivateOtherBytes int64 `json:"private_other_bytes"`
	PrivateOtherBytesIn int64 `json:"private_other_bytes_in"`
	PrivateOtherBytesOut int64 `json:"private_other_bytes_out"`
	UpdatedAt string `json:"updated_at"`
	Delta int64 `json:"delta"`
}


// Type for action response, including envelope
type ActionIpTrafficMonitorIndexResponse struct {
	Action *ActionIpTrafficMonitorIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpTrafficMonitors []*ActionIpTrafficMonitorIndexOutput `json:"ip_traffic_monitors"`
	}

	// Action output without the namespace
	Output []*ActionIpTrafficMonitorIndexOutput
}


// Prepare the action for invocation
func (action *ActionIpTrafficMonitorIndex) Prepare() *ActionIpTrafficMonitorIndexInvocation {
	return &ActionIpTrafficMonitorIndexInvocation{
		Action: action,
		Path: "/v6.0/ip_traffic_monitors",
	}
}

// ActionIpTrafficMonitorIndexInvocation is used to configure action for invocation
type ActionIpTrafficMonitorIndexInvocation struct {
	// Pointer to the action
	Action *ActionIpTrafficMonitorIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpTrafficMonitorIndexInput
	// Global meta input parameters
	MetaInput *ActionIpTrafficMonitorIndexMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpTrafficMonitorIndexInvocation) NewInput() *ActionIpTrafficMonitorIndexInput {
	inv.Input = &ActionIpTrafficMonitorIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpTrafficMonitorIndexInvocation) SetInput(input *ActionIpTrafficMonitorIndexInput) *ActionIpTrafficMonitorIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpTrafficMonitorIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpTrafficMonitorIndexInvocation) NewMetaInput() *ActionIpTrafficMonitorIndexMetaGlobalInput {
	inv.MetaInput = &ActionIpTrafficMonitorIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpTrafficMonitorIndexInvocation) SetMetaInput(input *ActionIpTrafficMonitorIndexMetaGlobalInput) *ActionIpTrafficMonitorIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpTrafficMonitorIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIpTrafficMonitorIndexInvocation) Call() (*ActionIpTrafficMonitorIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionIpTrafficMonitorIndexInvocation) callAsQuery() (*ActionIpTrafficMonitorIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIpTrafficMonitorIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IpTrafficMonitors
	}
	return resp, err
}



func (inv *ActionIpTrafficMonitorIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["ip_traffic_monitor[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["ip_traffic_monitor[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("IpAddress") {
			ret["ip_traffic_monitor[ip_address]"] = convertInt64ToString(inv.Input.IpAddress)
		}
		if inv.IsParameterSelected("User") {
			ret["ip_traffic_monitor[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("IpVersion") {
			ret["ip_traffic_monitor[ip_version]"] = convertInt64ToString(inv.Input.IpVersion)
		}
		if inv.IsParameterSelected("Environment") {
			ret["ip_traffic_monitor[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("Location") {
			ret["ip_traffic_monitor[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("Network") {
			ret["ip_traffic_monitor[network]"] = convertInt64ToString(inv.Input.Network)
		}
		if inv.IsParameterSelected("Node") {
			ret["ip_traffic_monitor[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("Vps") {
			ret["ip_traffic_monitor[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
		if inv.IsParameterSelected("Order") {
			ret["ip_traffic_monitor[order]"] = inv.Input.Order
		}
	}
}

func (inv *ActionIpTrafficMonitorIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

