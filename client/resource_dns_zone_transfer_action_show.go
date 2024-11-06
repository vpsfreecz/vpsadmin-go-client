package client

import (
	"strings"
)

// ActionDnsZoneTransferShow is a type for action Dns_zone_transfer#Show
type ActionDnsZoneTransferShow struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsZoneTransferShow(client *Client) *ActionDnsZoneTransferShow {
	return &ActionDnsZoneTransferShow{
		Client: client,
	}
}

// ActionDnsZoneTransferShowMetaGlobalInput is a type for action global meta input parameters
type ActionDnsZoneTransferShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsZoneTransferShowMetaGlobalInput) SetIncludes(value string) *ActionDnsZoneTransferShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsZoneTransferShowMetaGlobalInput) SetNo(value bool) *ActionDnsZoneTransferShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsZoneTransferShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsZoneTransferShowMetaGlobalInput) SelectParameters(params ...string) *ActionDnsZoneTransferShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsZoneTransferShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsZoneTransferShowOutput is a type for action output parameters
type ActionDnsZoneTransferShowOutput struct {
	CreatedAt     string                         `json:"created_at"`
	DnsTsigKey    *ActionDnsTsigKeyShowOutput    `json:"dns_tsig_key"`
	DnsZone       *ActionDnsZoneShowOutput       `json:"dns_zone"`
	HostIpAddress *ActionHostIpAddressShowOutput `json:"host_ip_address"`
	Id            int64                          `json:"id"`
	PeerType      string                         `json:"peer_type"`
	UpdatedAt     string                         `json:"updated_at"`
}

// Type for action response, including envelope
type ActionDnsZoneTransferShowResponse struct {
	Action *ActionDnsZoneTransferShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsZoneTransfer *ActionDnsZoneTransferShowOutput `json:"dns_zone_transfer"`
	}

	// Action output without the namespace
	Output *ActionDnsZoneTransferShowOutput
}

// Prepare the action for invocation
func (action *ActionDnsZoneTransferShow) Prepare() *ActionDnsZoneTransferShowInvocation {
	return &ActionDnsZoneTransferShowInvocation{
		Action: action,
		Path:   "/v7.0/dns_zone_transfers/{dns_zone_transfer_id}",
	}
}

// ActionDnsZoneTransferShowInvocation is used to configure action for invocation
type ActionDnsZoneTransferShowInvocation struct {
	// Pointer to the action
	Action *ActionDnsZoneTransferShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDnsZoneTransferShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDnsZoneTransferShowInvocation) SetPathParamInt(param string, value int64) *ActionDnsZoneTransferShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDnsZoneTransferShowInvocation) SetPathParamString(param string, value string) *ActionDnsZoneTransferShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsZoneTransferShowInvocation) NewMetaInput() *ActionDnsZoneTransferShowMetaGlobalInput {
	inv.MetaInput = &ActionDnsZoneTransferShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsZoneTransferShowInvocation) SetMetaInput(input *ActionDnsZoneTransferShowMetaGlobalInput) *ActionDnsZoneTransferShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsZoneTransferShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsZoneTransferShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsZoneTransferShowInvocation) Call() (*ActionDnsZoneTransferShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDnsZoneTransferShowInvocation) callAsQuery() (*ActionDnsZoneTransferShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDnsZoneTransferShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsZoneTransfer
	}
	return resp, err
}

func (inv *ActionDnsZoneTransferShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
