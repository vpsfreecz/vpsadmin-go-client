package client

import ()

// ActionDnsServerIndex is a type for action Dns_server#Index
type ActionDnsServerIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsServerIndex(client *Client) *ActionDnsServerIndex {
	return &ActionDnsServerIndex{
		Client: client,
	}
}

// ActionDnsServerIndexMetaGlobalInput is a type for action global meta input parameters
type ActionDnsServerIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionDnsServerIndexMetaGlobalInput) SetCount(value bool) *ActionDnsServerIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsServerIndexMetaGlobalInput) SetIncludes(value string) *ActionDnsServerIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsServerIndexMetaGlobalInput) SetNo(value bool) *ActionDnsServerIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsServerIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsServerIndexMetaGlobalInput) SelectParameters(params ...string) *ActionDnsServerIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsServerIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsServerIndexInput is a type for action input parameters
type ActionDnsServerIndexInput struct {
	FromId int64 `json:"from_id"`
	Limit  int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionDnsServerIndexInput) SetFromId(value int64) *ActionDnsServerIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDnsServerIndexInput) SetLimit(value int64) *ActionDnsServerIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsServerIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsServerIndexInput) SelectParameters(params ...string) *ActionDnsServerIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnsServerIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnsServerIndexInput) UnselectParameters(params ...string) *ActionDnsServerIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnsServerIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsServerIndexOutput is a type for action output parameters
type ActionDnsServerIndexOutput struct {
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
type ActionDnsServerIndexResponse struct {
	Action *ActionDnsServerIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsServers []*ActionDnsServerIndexOutput `json:"dns_servers"`
	}

	// Action output without the namespace
	Output []*ActionDnsServerIndexOutput
}

// Prepare the action for invocation
func (action *ActionDnsServerIndex) Prepare() *ActionDnsServerIndexInvocation {
	return &ActionDnsServerIndexInvocation{
		Action: action,
		Path:   "/v7.0/dns_servers",
	}
}

// ActionDnsServerIndexInvocation is used to configure action for invocation
type ActionDnsServerIndexInvocation struct {
	// Pointer to the action
	Action *ActionDnsServerIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsServerIndexInput
	// Global meta input parameters
	MetaInput *ActionDnsServerIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsServerIndexInvocation) NewInput() *ActionDnsServerIndexInput {
	inv.Input = &ActionDnsServerIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsServerIndexInvocation) SetInput(input *ActionDnsServerIndexInput) *ActionDnsServerIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsServerIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnsServerIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsServerIndexInvocation) NewMetaInput() *ActionDnsServerIndexMetaGlobalInput {
	inv.MetaInput = &ActionDnsServerIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsServerIndexInvocation) SetMetaInput(input *ActionDnsServerIndexMetaGlobalInput) *ActionDnsServerIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsServerIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsServerIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsServerIndexInvocation) Call() (*ActionDnsServerIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDnsServerIndexInvocation) callAsQuery() (*ActionDnsServerIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDnsServerIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsServers
	}
	return resp, err
}

func (inv *ActionDnsServerIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["dns_server[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["dns_server[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionDnsServerIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
