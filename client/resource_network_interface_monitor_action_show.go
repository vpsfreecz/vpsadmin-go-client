package client

import (
	"strings"
)

// ActionNetworkInterfaceMonitorShow is a type for action Network_interface_monitor#Show
type ActionNetworkInterfaceMonitorShow struct {
	// Pointer to client
	Client *Client
}

func NewActionNetworkInterfaceMonitorShow(client *Client) *ActionNetworkInterfaceMonitorShow {
	return &ActionNetworkInterfaceMonitorShow{
		Client: client,
	}
}

// ActionNetworkInterfaceMonitorShowMetaGlobalInput is a type for action global meta input parameters
type ActionNetworkInterfaceMonitorShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNetworkInterfaceMonitorShowMetaGlobalInput) SetIncludes(value string) *ActionNetworkInterfaceMonitorShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNetworkInterfaceMonitorShowMetaGlobalInput) SetNo(value bool) *ActionNetworkInterfaceMonitorShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkInterfaceMonitorShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkInterfaceMonitorShowMetaGlobalInput) SelectParameters(params ...string) *ActionNetworkInterfaceMonitorShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNetworkInterfaceMonitorShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNetworkInterfaceMonitorShowOutput is a type for action output parameters
type ActionNetworkInterfaceMonitorShowOutput struct {
	Bytes            int64                             `json:"bytes"`
	BytesIn          int64                             `json:"bytes_in"`
	BytesOut         int64                             `json:"bytes_out"`
	Delta            int64                             `json:"delta"`
	Id               int64                             `json:"id"`
	NetworkInterface *ActionNetworkInterfaceShowOutput `json:"network_interface"`
	Packets          int64                             `json:"packets"`
	PacketsIn        int64                             `json:"packets_in"`
	PacketsOut       int64                             `json:"packets_out"`
	UpdatedAt        string                            `json:"updated_at"`
}

// Type for action response, including envelope
type ActionNetworkInterfaceMonitorShowResponse struct {
	Action *ActionNetworkInterfaceMonitorShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NetworkInterfaceMonitor *ActionNetworkInterfaceMonitorShowOutput `json:"network_interface_monitor"`
	}

	// Action output without the namespace
	Output *ActionNetworkInterfaceMonitorShowOutput
}

// Prepare the action for invocation
func (action *ActionNetworkInterfaceMonitorShow) Prepare() *ActionNetworkInterfaceMonitorShowInvocation {
	return &ActionNetworkInterfaceMonitorShowInvocation{
		Action: action,
		Path:   "/v7.0/network_interface_monitors/{network_interface_monitor_id}",
	}
}

// ActionNetworkInterfaceMonitorShowInvocation is used to configure action for invocation
type ActionNetworkInterfaceMonitorShowInvocation struct {
	// Pointer to the action
	Action *ActionNetworkInterfaceMonitorShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNetworkInterfaceMonitorShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNetworkInterfaceMonitorShowInvocation) SetPathParamInt(param string, value int64) *ActionNetworkInterfaceMonitorShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNetworkInterfaceMonitorShowInvocation) SetPathParamString(param string, value string) *ActionNetworkInterfaceMonitorShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNetworkInterfaceMonitorShowInvocation) NewMetaInput() *ActionNetworkInterfaceMonitorShowMetaGlobalInput {
	inv.MetaInput = &ActionNetworkInterfaceMonitorShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNetworkInterfaceMonitorShowInvocation) SetMetaInput(input *ActionNetworkInterfaceMonitorShowMetaGlobalInput) *ActionNetworkInterfaceMonitorShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNetworkInterfaceMonitorShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNetworkInterfaceMonitorShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNetworkInterfaceMonitorShowInvocation) Call() (*ActionNetworkInterfaceMonitorShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionNetworkInterfaceMonitorShowInvocation) callAsQuery() (*ActionNetworkInterfaceMonitorShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNetworkInterfaceMonitorShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NetworkInterfaceMonitor
	}
	return resp, err
}

func (inv *ActionNetworkInterfaceMonitorShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
