package client

import (
	"strings"
)

// ActionDnsServerShow is a type for action Dns_server#Show
type ActionDnsServerShow struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsServerShow(client *Client) *ActionDnsServerShow {
	return &ActionDnsServerShow{
		Client: client,
	}
}

// ActionDnsServerShowMetaGlobalInput is a type for action global meta input parameters
type ActionDnsServerShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsServerShowMetaGlobalInput) SetIncludes(value string) *ActionDnsServerShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsServerShowMetaGlobalInput) SetNo(value bool) *ActionDnsServerShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsServerShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsServerShowMetaGlobalInput) SelectParameters(params ...string) *ActionDnsServerShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsServerShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsServerShowOutput is a type for action output parameters
type ActionDnsServerShowOutput struct {
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
type ActionDnsServerShowResponse struct {
	Action *ActionDnsServerShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsServer *ActionDnsServerShowOutput `json:"dns_server"`
	}

	// Action output without the namespace
	Output *ActionDnsServerShowOutput
}

// Prepare the action for invocation
func (action *ActionDnsServerShow) Prepare() *ActionDnsServerShowInvocation {
	return &ActionDnsServerShowInvocation{
		Action: action,
		Path:   "/v7.0/dns_servers/{dns_server_id}",
	}
}

// ActionDnsServerShowInvocation is used to configure action for invocation
type ActionDnsServerShowInvocation struct {
	// Pointer to the action
	Action *ActionDnsServerShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDnsServerShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDnsServerShowInvocation) SetPathParamInt(param string, value int64) *ActionDnsServerShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDnsServerShowInvocation) SetPathParamString(param string, value string) *ActionDnsServerShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsServerShowInvocation) NewMetaInput() *ActionDnsServerShowMetaGlobalInput {
	inv.MetaInput = &ActionDnsServerShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsServerShowInvocation) SetMetaInput(input *ActionDnsServerShowMetaGlobalInput) *ActionDnsServerShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsServerShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsServerShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsServerShowInvocation) Call() (*ActionDnsServerShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDnsServerShowInvocation) callAsQuery() (*ActionDnsServerShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDnsServerShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsServer
	}
	return resp, err
}

func (inv *ActionDnsServerShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
