package client

import (
)

// ActionDnsResolverIndex is a type for action Dns_resolver#Index
type ActionDnsResolverIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsResolverIndex(client *Client) *ActionDnsResolverIndex {
	return &ActionDnsResolverIndex{
		Client: client,
	}
}

// ActionDnsResolverIndexMetaGlobalInput is a type for action global meta input parameters
type ActionDnsResolverIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsResolverIndexMetaGlobalInput) SetNo(value bool) *ActionDnsResolverIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionDnsResolverIndexMetaGlobalInput) SetCount(value bool) *ActionDnsResolverIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsResolverIndexMetaGlobalInput) SetIncludes(value string) *ActionDnsResolverIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsResolverIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsResolverIndexMetaGlobalInput) SelectParameters(params ...string) *ActionDnsResolverIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsResolverIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsResolverIndexInput is a type for action input parameters
type ActionDnsResolverIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	Vps int64 `json:"vps"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionDnsResolverIndexInput) SetOffset(value int64) *ActionDnsResolverIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDnsResolverIndexInput) SetLimit(value int64) *ActionDnsResolverIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionDnsResolverIndexInput) SetVps(value int64) *ActionDnsResolverIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Vps"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsResolverIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsResolverIndexInput) SelectParameters(params ...string) *ActionDnsResolverIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsResolverIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionDnsResolverIndexOutput is a type for action output parameters
type ActionDnsResolverIndexOutput struct {
	Id int64 `json:"id"`
	IpAddr string `json:"ip_addr"`
	Label string `json:"label"`
	IsUniversal bool `json:"is_universal"`
	Location *ActionLocationShowOutput `json:"location"`
}


// Type for action response, including envelope
type ActionDnsResolverIndexResponse struct {
	Action *ActionDnsResolverIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsResolvers []*ActionDnsResolverIndexOutput `json:"dns_resolvers"`
	}

	// Action output without the namespace
	Output []*ActionDnsResolverIndexOutput
}


// Prepare the action for invocation
func (action *ActionDnsResolverIndex) Prepare() *ActionDnsResolverIndexInvocation {
	return &ActionDnsResolverIndexInvocation{
		Action: action,
		Path: "/v6.0/dns_resolvers",
	}
}

// ActionDnsResolverIndexInvocation is used to configure action for invocation
type ActionDnsResolverIndexInvocation struct {
	// Pointer to the action
	Action *ActionDnsResolverIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsResolverIndexInput
	// Global meta input parameters
	MetaInput *ActionDnsResolverIndexMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsResolverIndexInvocation) NewInput() *ActionDnsResolverIndexInput {
	inv.Input = &ActionDnsResolverIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsResolverIndexInvocation) SetInput(input *ActionDnsResolverIndexInput) *ActionDnsResolverIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsResolverIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsResolverIndexInvocation) NewMetaInput() *ActionDnsResolverIndexMetaGlobalInput {
	inv.MetaInput = &ActionDnsResolverIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsResolverIndexInvocation) SetMetaInput(input *ActionDnsResolverIndexMetaGlobalInput) *ActionDnsResolverIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsResolverIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsResolverIndexInvocation) Call() (*ActionDnsResolverIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDnsResolverIndexInvocation) callAsQuery() (*ActionDnsResolverIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDnsResolverIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsResolvers
	}
	return resp, err
}



func (inv *ActionDnsResolverIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["dns_resolver[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["dns_resolver[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Vps") {
			ret["dns_resolver[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
	}
}

func (inv *ActionDnsResolverIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

