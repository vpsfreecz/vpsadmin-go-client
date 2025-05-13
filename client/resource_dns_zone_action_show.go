package client

import (
	"strings"
)

// ActionDnsZoneShow is a type for action Dns_zone#Show
type ActionDnsZoneShow struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsZoneShow(client *Client) *ActionDnsZoneShow {
	return &ActionDnsZoneShow{
		Client: client,
	}
}

// ActionDnsZoneShowMetaGlobalInput is a type for action global meta input parameters
type ActionDnsZoneShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsZoneShowMetaGlobalInput) SetIncludes(value string) *ActionDnsZoneShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsZoneShowMetaGlobalInput) SetNo(value bool) *ActionDnsZoneShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsZoneShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsZoneShowMetaGlobalInput) SelectParameters(params ...string) *ActionDnsZoneShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsZoneShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsZoneShowOutput is a type for action output parameters
type ActionDnsZoneShowOutput struct {
	CreatedAt             string                `json:"created_at"`
	DefaultTtl            int64                 `json:"default_ttl"`
	DnssecEnabled         bool                  `json:"dnssec_enabled"`
	Email                 string                `json:"email"`
	Enabled               bool                  `json:"enabled"`
	Id                    int64                 `json:"id"`
	Label                 string                `json:"label"`
	Name                  string                `json:"name"`
	ReverseNetworkAddress string                `json:"reverse_network_address"`
	ReverseNetworkPrefix  string                `json:"reverse_network_prefix"`
	Role                  string                `json:"role"`
	Serial                int64                 `json:"serial"`
	Source                string                `json:"source"`
	UpdatedAt             string                `json:"updated_at"`
	User                  *ActionUserShowOutput `json:"user"`
}

// Type for action response, including envelope
type ActionDnsZoneShowResponse struct {
	Action *ActionDnsZoneShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsZone *ActionDnsZoneShowOutput `json:"dns_zone"`
	}

	// Action output without the namespace
	Output *ActionDnsZoneShowOutput
}

// Prepare the action for invocation
func (action *ActionDnsZoneShow) Prepare() *ActionDnsZoneShowInvocation {
	return &ActionDnsZoneShowInvocation{
		Action: action,
		Path:   "/v7.0/dns_zones/{dns_zone_id}",
	}
}

// ActionDnsZoneShowInvocation is used to configure action for invocation
type ActionDnsZoneShowInvocation struct {
	// Pointer to the action
	Action *ActionDnsZoneShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDnsZoneShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDnsZoneShowInvocation) SetPathParamInt(param string, value int64) *ActionDnsZoneShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDnsZoneShowInvocation) SetPathParamString(param string, value string) *ActionDnsZoneShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsZoneShowInvocation) NewMetaInput() *ActionDnsZoneShowMetaGlobalInput {
	inv.MetaInput = &ActionDnsZoneShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsZoneShowInvocation) SetMetaInput(input *ActionDnsZoneShowMetaGlobalInput) *ActionDnsZoneShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsZoneShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsZoneShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsZoneShowInvocation) Call() (*ActionDnsZoneShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDnsZoneShowInvocation) callAsQuery() (*ActionDnsZoneShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDnsZoneShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsZone
	}
	return resp, err
}

func (inv *ActionDnsZoneShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
