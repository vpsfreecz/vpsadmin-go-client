package client

import ()

// ActionDnsZoneTransferIndex is a type for action Dns_zone_transfer#Index
type ActionDnsZoneTransferIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsZoneTransferIndex(client *Client) *ActionDnsZoneTransferIndex {
	return &ActionDnsZoneTransferIndex{
		Client: client,
	}
}

// ActionDnsZoneTransferIndexMetaGlobalInput is a type for action global meta input parameters
type ActionDnsZoneTransferIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionDnsZoneTransferIndexMetaGlobalInput) SetCount(value bool) *ActionDnsZoneTransferIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsZoneTransferIndexMetaGlobalInput) SetIncludes(value string) *ActionDnsZoneTransferIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsZoneTransferIndexMetaGlobalInput) SetNo(value bool) *ActionDnsZoneTransferIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsZoneTransferIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsZoneTransferIndexMetaGlobalInput) SelectParameters(params ...string) *ActionDnsZoneTransferIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsZoneTransferIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsZoneTransferIndexInput is a type for action input parameters
type ActionDnsZoneTransferIndexInput struct {
	DnsTsigKey    int64  `json:"dns_tsig_key"`
	DnsZone       int64  `json:"dns_zone"`
	FromId        int64  `json:"from_id"`
	HostIpAddress int64  `json:"host_ip_address"`
	Limit         int64  `json:"limit"`
	PeerType      string `json:"peer_type"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDnsTsigKey sets parameter DnsTsigKey to value and selects it for sending
func (in *ActionDnsZoneTransferIndexInput) SetDnsTsigKey(value int64) *ActionDnsZoneTransferIndexInput {
	in.DnsTsigKey = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetDnsTsigKeyNil(false)
	in._selectedParameters["DnsTsigKey"] = nil
	return in
}

// SetDnsTsigKeyNil sets parameter DnsTsigKey to nil and selects it for sending
func (in *ActionDnsZoneTransferIndexInput) SetDnsTsigKeyNil(set bool) *ActionDnsZoneTransferIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["DnsTsigKey"] = nil
		in.SelectParameters("DnsTsigKey")
	} else {
		delete(in._nilParameters, "DnsTsigKey")
	}
	return in
}

// SetDnsZone sets parameter DnsZone to value and selects it for sending
func (in *ActionDnsZoneTransferIndexInput) SetDnsZone(value int64) *ActionDnsZoneTransferIndexInput {
	in.DnsZone = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetDnsZoneNil(false)
	in._selectedParameters["DnsZone"] = nil
	return in
}

// SetDnsZoneNil sets parameter DnsZone to nil and selects it for sending
func (in *ActionDnsZoneTransferIndexInput) SetDnsZoneNil(set bool) *ActionDnsZoneTransferIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["DnsZone"] = nil
		in.SelectParameters("DnsZone")
	} else {
		delete(in._nilParameters, "DnsZone")
	}
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionDnsZoneTransferIndexInput) SetFromId(value int64) *ActionDnsZoneTransferIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetHostIpAddress sets parameter HostIpAddress to value and selects it for sending
func (in *ActionDnsZoneTransferIndexInput) SetHostIpAddress(value int64) *ActionDnsZoneTransferIndexInput {
	in.HostIpAddress = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetHostIpAddressNil(false)
	in._selectedParameters["HostIpAddress"] = nil
	return in
}

// SetHostIpAddressNil sets parameter HostIpAddress to nil and selects it for sending
func (in *ActionDnsZoneTransferIndexInput) SetHostIpAddressNil(set bool) *ActionDnsZoneTransferIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["HostIpAddress"] = nil
		in.SelectParameters("HostIpAddress")
	} else {
		delete(in._nilParameters, "HostIpAddress")
	}
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDnsZoneTransferIndexInput) SetLimit(value int64) *ActionDnsZoneTransferIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetPeerType sets parameter PeerType to value and selects it for sending
func (in *ActionDnsZoneTransferIndexInput) SetPeerType(value string) *ActionDnsZoneTransferIndexInput {
	in.PeerType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["PeerType"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsZoneTransferIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsZoneTransferIndexInput) SelectParameters(params ...string) *ActionDnsZoneTransferIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnsZoneTransferIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnsZoneTransferIndexInput) UnselectParameters(params ...string) *ActionDnsZoneTransferIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnsZoneTransferIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsZoneTransferIndexOutput is a type for action output parameters
type ActionDnsZoneTransferIndexOutput struct {
	CreatedAt     string                         `json:"created_at"`
	DnsTsigKey    *ActionDnsTsigKeyShowOutput    `json:"dns_tsig_key"`
	DnsZone       *ActionDnsZoneShowOutput       `json:"dns_zone"`
	HostIpAddress *ActionHostIpAddressShowOutput `json:"host_ip_address"`
	Id            int64                          `json:"id"`
	PeerType      string                         `json:"peer_type"`
	UpdatedAt     string                         `json:"updated_at"`
}

// Type for action response, including envelope
type ActionDnsZoneTransferIndexResponse struct {
	Action *ActionDnsZoneTransferIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsZoneTransfers []*ActionDnsZoneTransferIndexOutput `json:"dns_zone_transfers"`
	}

	// Action output without the namespace
	Output []*ActionDnsZoneTransferIndexOutput
}

// Prepare the action for invocation
func (action *ActionDnsZoneTransferIndex) Prepare() *ActionDnsZoneTransferIndexInvocation {
	return &ActionDnsZoneTransferIndexInvocation{
		Action: action,
		Path:   "/v7.0/dns_zone_transfers",
	}
}

// ActionDnsZoneTransferIndexInvocation is used to configure action for invocation
type ActionDnsZoneTransferIndexInvocation struct {
	// Pointer to the action
	Action *ActionDnsZoneTransferIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsZoneTransferIndexInput
	// Global meta input parameters
	MetaInput *ActionDnsZoneTransferIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsZoneTransferIndexInvocation) NewInput() *ActionDnsZoneTransferIndexInput {
	inv.Input = &ActionDnsZoneTransferIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsZoneTransferIndexInvocation) SetInput(input *ActionDnsZoneTransferIndexInput) *ActionDnsZoneTransferIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsZoneTransferIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnsZoneTransferIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsZoneTransferIndexInvocation) NewMetaInput() *ActionDnsZoneTransferIndexMetaGlobalInput {
	inv.MetaInput = &ActionDnsZoneTransferIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsZoneTransferIndexInvocation) SetMetaInput(input *ActionDnsZoneTransferIndexMetaGlobalInput) *ActionDnsZoneTransferIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsZoneTransferIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsZoneTransferIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsZoneTransferIndexInvocation) Call() (*ActionDnsZoneTransferIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDnsZoneTransferIndexInvocation) callAsQuery() (*ActionDnsZoneTransferIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDnsZoneTransferIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsZoneTransfers
	}
	return resp, err
}

func (inv *ActionDnsZoneTransferIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("DnsTsigKey") {
			ret["dns_zone_transfer[dns_tsig_key]"] = convertInt64ToString(inv.Input.DnsTsigKey)
		}
		if inv.IsParameterSelected("DnsZone") {
			ret["dns_zone_transfer[dns_zone]"] = convertInt64ToString(inv.Input.DnsZone)
		}
		if inv.IsParameterSelected("FromId") {
			ret["dns_zone_transfer[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("HostIpAddress") {
			ret["dns_zone_transfer[host_ip_address]"] = convertInt64ToString(inv.Input.HostIpAddress)
		}
		if inv.IsParameterSelected("Limit") {
			ret["dns_zone_transfer[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("PeerType") {
			ret["dns_zone_transfer[peer_type]"] = inv.Input.PeerType
		}
	}
}

func (inv *ActionDnsZoneTransferIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
