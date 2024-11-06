package client

import (
	"strings"
)

// ActionDnsServerZoneShow is a type for action Dns_server_zone#Show
type ActionDnsServerZoneShow struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsServerZoneShow(client *Client) *ActionDnsServerZoneShow {
	return &ActionDnsServerZoneShow{
		Client: client,
	}
}

// ActionDnsServerZoneShowMetaGlobalInput is a type for action global meta input parameters
type ActionDnsServerZoneShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsServerZoneShowMetaGlobalInput) SetIncludes(value string) *ActionDnsServerZoneShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsServerZoneShowMetaGlobalInput) SetNo(value bool) *ActionDnsServerZoneShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsServerZoneShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsServerZoneShowMetaGlobalInput) SelectParameters(params ...string) *ActionDnsServerZoneShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsServerZoneShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsServerZoneShowOutput is a type for action output parameters
type ActionDnsServerZoneShowOutput struct {
	CreatedAt   string                     `json:"created_at"`
	DnsServer   *ActionDnsServerShowOutput `json:"dns_server"`
	DnsZone     *ActionDnsZoneShowOutput   `json:"dns_zone"`
	ExpiresAt   string                     `json:"expires_at"`
	Id          int64                      `json:"id"`
	LastCheckAt string                     `json:"last_check_at"`
	LoadedAt    string                     `json:"loaded_at"`
	RefreshAt   string                     `json:"refresh_at"`
	Serial      int64                      `json:"serial"`
	Type        string                     `json:"type"`
	UpdatedAt   string                     `json:"updated_at"`
}

// Type for action response, including envelope
type ActionDnsServerZoneShowResponse struct {
	Action *ActionDnsServerZoneShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsServerZone *ActionDnsServerZoneShowOutput `json:"dns_server_zone"`
	}

	// Action output without the namespace
	Output *ActionDnsServerZoneShowOutput
}

// Prepare the action for invocation
func (action *ActionDnsServerZoneShow) Prepare() *ActionDnsServerZoneShowInvocation {
	return &ActionDnsServerZoneShowInvocation{
		Action: action,
		Path:   "/v7.0/dns_server_zones/{dns_server_zone_id}",
	}
}

// ActionDnsServerZoneShowInvocation is used to configure action for invocation
type ActionDnsServerZoneShowInvocation struct {
	// Pointer to the action
	Action *ActionDnsServerZoneShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDnsServerZoneShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDnsServerZoneShowInvocation) SetPathParamInt(param string, value int64) *ActionDnsServerZoneShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDnsServerZoneShowInvocation) SetPathParamString(param string, value string) *ActionDnsServerZoneShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsServerZoneShowInvocation) NewMetaInput() *ActionDnsServerZoneShowMetaGlobalInput {
	inv.MetaInput = &ActionDnsServerZoneShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsServerZoneShowInvocation) SetMetaInput(input *ActionDnsServerZoneShowMetaGlobalInput) *ActionDnsServerZoneShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsServerZoneShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsServerZoneShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsServerZoneShowInvocation) Call() (*ActionDnsServerZoneShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDnsServerZoneShowInvocation) callAsQuery() (*ActionDnsServerZoneShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDnsServerZoneShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsServerZone
	}
	return resp, err
}

func (inv *ActionDnsServerZoneShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
