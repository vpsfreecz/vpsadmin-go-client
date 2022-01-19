package client

import (
	"strings"
)

// ActionIpTrafficShow is a type for action Ip_traffic#Show
type ActionIpTrafficShow struct {
	// Pointer to client
	Client *Client
}

func NewActionIpTrafficShow(client *Client) *ActionIpTrafficShow {
	return &ActionIpTrafficShow{
		Client: client,
	}
}

// ActionIpTrafficShowMetaGlobalInput is a type for action global meta input parameters
type ActionIpTrafficShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpTrafficShowMetaGlobalInput) SetIncludes(value string) *ActionIpTrafficShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpTrafficShowMetaGlobalInput) SetNo(value bool) *ActionIpTrafficShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpTrafficShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpTrafficShowMetaGlobalInput) SelectParameters(params ...string) *ActionIpTrafficShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpTrafficShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpTrafficShowOutput is a type for action output parameters
type ActionIpTrafficShowOutput struct {
	BytesIn    int64                      `json:"bytes_in"`
	BytesOut   int64                      `json:"bytes_out"`
	CreatedAt  string                     `json:"created_at"`
	Id         int64                      `json:"id"`
	IpAddress  *ActionIpAddressShowOutput `json:"ip_address"`
	PacketsIn  int64                      `json:"packets_in"`
	PacketsOut int64                      `json:"packets_out"`
	Protocol   string                     `json:"protocol"`
	Role       string                     `json:"role"`
	User       *ActionUserShowOutput      `json:"user"`
}

// Type for action response, including envelope
type ActionIpTrafficShowResponse struct {
	Action *ActionIpTrafficShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpTraffic *ActionIpTrafficShowOutput `json:"ip_traffic"`
	}

	// Action output without the namespace
	Output *ActionIpTrafficShowOutput
}

// Prepare the action for invocation
func (action *ActionIpTrafficShow) Prepare() *ActionIpTrafficShowInvocation {
	return &ActionIpTrafficShowInvocation{
		Action: action,
		Path:   "/v6.0/ip_traffics/{ip_traffic_id}",
	}
}

// ActionIpTrafficShowInvocation is used to configure action for invocation
type ActionIpTrafficShowInvocation struct {
	// Pointer to the action
	Action *ActionIpTrafficShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionIpTrafficShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpTrafficShowInvocation) SetPathParamInt(param string, value int64) *ActionIpTrafficShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpTrafficShowInvocation) SetPathParamString(param string, value string) *ActionIpTrafficShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpTrafficShowInvocation) NewMetaInput() *ActionIpTrafficShowMetaGlobalInput {
	inv.MetaInput = &ActionIpTrafficShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpTrafficShowInvocation) SetMetaInput(input *ActionIpTrafficShowMetaGlobalInput) *ActionIpTrafficShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpTrafficShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIpTrafficShowInvocation) Call() (*ActionIpTrafficShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionIpTrafficShowInvocation) callAsQuery() (*ActionIpTrafficShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIpTrafficShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IpTraffic
	}
	return resp, err
}

func (inv *ActionIpTrafficShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
