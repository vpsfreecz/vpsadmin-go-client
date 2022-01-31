package client

import (
	"strings"
)

// ActionDnsResolverShow is a type for action Dns_resolver#Show
type ActionDnsResolverShow struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsResolverShow(client *Client) *ActionDnsResolverShow {
	return &ActionDnsResolverShow{
		Client: client,
	}
}

// ActionDnsResolverShowMetaGlobalInput is a type for action global meta input parameters
type ActionDnsResolverShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsResolverShowMetaGlobalInput) SetIncludes(value string) *ActionDnsResolverShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsResolverShowMetaGlobalInput) SetNo(value bool) *ActionDnsResolverShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsResolverShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsResolverShowMetaGlobalInput) SelectParameters(params ...string) *ActionDnsResolverShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsResolverShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsResolverShowOutput is a type for action output parameters
type ActionDnsResolverShowOutput struct {
	Id          int64                     `json:"id"`
	IpAddr      string                    `json:"ip_addr"`
	IsUniversal bool                      `json:"is_universal"`
	Label       string                    `json:"label"`
	Location    *ActionLocationShowOutput `json:"location"`
}

// Type for action response, including envelope
type ActionDnsResolverShowResponse struct {
	Action *ActionDnsResolverShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsResolver *ActionDnsResolverShowOutput `json:"dns_resolver"`
	}

	// Action output without the namespace
	Output *ActionDnsResolverShowOutput
}

// Prepare the action for invocation
func (action *ActionDnsResolverShow) Prepare() *ActionDnsResolverShowInvocation {
	return &ActionDnsResolverShowInvocation{
		Action: action,
		Path:   "/v6.0/dns_resolvers/{dns_resolver_id}",
	}
}

// ActionDnsResolverShowInvocation is used to configure action for invocation
type ActionDnsResolverShowInvocation struct {
	// Pointer to the action
	Action *ActionDnsResolverShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDnsResolverShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDnsResolverShowInvocation) SetPathParamInt(param string, value int64) *ActionDnsResolverShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDnsResolverShowInvocation) SetPathParamString(param string, value string) *ActionDnsResolverShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsResolverShowInvocation) NewMetaInput() *ActionDnsResolverShowMetaGlobalInput {
	inv.MetaInput = &ActionDnsResolverShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsResolverShowInvocation) SetMetaInput(input *ActionDnsResolverShowMetaGlobalInput) *ActionDnsResolverShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsResolverShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsResolverShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsResolverShowInvocation) Call() (*ActionDnsResolverShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDnsResolverShowInvocation) callAsQuery() (*ActionDnsResolverShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDnsResolverShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsResolver
	}
	return resp, err
}

func (inv *ActionDnsResolverShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
