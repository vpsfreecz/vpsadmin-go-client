package client

import (
	"strings"
)

// ActionIpTrafficMonitorShow is a type for action Ip_traffic_monitor#Show
type ActionIpTrafficMonitorShow struct {
	// Pointer to client
	Client *Client
}

func NewActionIpTrafficMonitorShow(client *Client) *ActionIpTrafficMonitorShow {
	return &ActionIpTrafficMonitorShow{
		Client: client,
	}
}

// ActionIpTrafficMonitorShowMetaGlobalInput is a type for action global meta input parameters
type ActionIpTrafficMonitorShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpTrafficMonitorShowMetaGlobalInput) SetNo(value bool) *ActionIpTrafficMonitorShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpTrafficMonitorShowMetaGlobalInput) SetIncludes(value string) *ActionIpTrafficMonitorShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpTrafficMonitorShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpTrafficMonitorShowMetaGlobalInput) SelectParameters(params ...string) *ActionIpTrafficMonitorShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpTrafficMonitorShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionIpTrafficMonitorShowOutput is a type for action output parameters
type ActionIpTrafficMonitorShowOutput struct {
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
type ActionIpTrafficMonitorShowResponse struct {
	Action *ActionIpTrafficMonitorShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpTrafficMonitor *ActionIpTrafficMonitorShowOutput `json:"ip_traffic_monitor"`
	}

	// Action output without the namespace
	Output *ActionIpTrafficMonitorShowOutput
}


// Prepare the action for invocation
func (action *ActionIpTrafficMonitorShow) Prepare() *ActionIpTrafficMonitorShowInvocation {
	return &ActionIpTrafficMonitorShowInvocation{
		Action: action,
		Path: "/v5.0/ip_traffic_monitors/{ip_traffic_monitor_id}",
	}
}

// ActionIpTrafficMonitorShowInvocation is used to configure action for invocation
type ActionIpTrafficMonitorShowInvocation struct {
	// Pointer to the action
	Action *ActionIpTrafficMonitorShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionIpTrafficMonitorShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpTrafficMonitorShowInvocation) SetPathParamInt(param string, value int64) *ActionIpTrafficMonitorShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpTrafficMonitorShowInvocation) SetPathParamString(param string, value string) *ActionIpTrafficMonitorShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpTrafficMonitorShowInvocation) NewMetaInput() *ActionIpTrafficMonitorShowMetaGlobalInput {
	inv.MetaInput = &ActionIpTrafficMonitorShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpTrafficMonitorShowInvocation) SetMetaInput(input *ActionIpTrafficMonitorShowMetaGlobalInput) *ActionIpTrafficMonitorShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpTrafficMonitorShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIpTrafficMonitorShowInvocation) Call() (*ActionIpTrafficMonitorShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionIpTrafficMonitorShowInvocation) callAsQuery() (*ActionIpTrafficMonitorShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIpTrafficMonitorShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IpTrafficMonitor
	}
	return resp, err
}




func (inv *ActionIpTrafficMonitorShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

