package client

import (
	"strings"
)

// ActionIpAddressShow is a type for action Ip_address#Show
type ActionIpAddressShow struct {
	// Pointer to client
	Client *Client
}

func NewActionIpAddressShow(client *Client) *ActionIpAddressShow {
	return &ActionIpAddressShow{
		Client: client,
	}
}

// ActionIpAddressShowMetaGlobalInput is a type for action global meta input parameters
type ActionIpAddressShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpAddressShowMetaGlobalInput) SetNo(value bool) *ActionIpAddressShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpAddressShowMetaGlobalInput) SetIncludes(value string) *ActionIpAddressShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpAddressShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpAddressShowMetaGlobalInput) SelectParameters(params ...string) *ActionIpAddressShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpAddressShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionIpAddressShowOutput is a type for action output parameters
type ActionIpAddressShowOutput struct {
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
type ActionIpAddressShowResponse struct {
	Action *ActionIpAddressShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpAddress *ActionIpAddressShowOutput `json:"ip_address"`
	}

	// Action output without the namespace
	Output *ActionIpAddressShowOutput
}


// Prepare the action for invocation
func (action *ActionIpAddressShow) Prepare() *ActionIpAddressShowInvocation {
	return &ActionIpAddressShowInvocation{
		Action: action,
		Path: "/v5.0/ip_addresses/:ip_address_id",
	}
}

// ActionIpAddressShowInvocation is used to configure action for invocation
type ActionIpAddressShowInvocation struct {
	// Pointer to the action
	Action *ActionIpAddressShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionIpAddressShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpAddressShowInvocation) SetPathParamInt(param string, value int64) *ActionIpAddressShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpAddressShowInvocation) SetPathParamString(param string, value string) *ActionIpAddressShowInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpAddressShowInvocation) SetMetaInput(input *ActionIpAddressShowMetaGlobalInput) *ActionIpAddressShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpAddressShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIpAddressShowInvocation) Call() (*ActionIpAddressShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionIpAddressShowInvocation) callAsQuery() (*ActionIpAddressShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIpAddressShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IpAddress
	}
	return resp, err
}




func (inv *ActionIpAddressShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

