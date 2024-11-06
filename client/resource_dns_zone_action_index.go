package client

import ()

// ActionDnsZoneIndex is a type for action Dns_zone#Index
type ActionDnsZoneIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsZoneIndex(client *Client) *ActionDnsZoneIndex {
	return &ActionDnsZoneIndex{
		Client: client,
	}
}

// ActionDnsZoneIndexMetaGlobalInput is a type for action global meta input parameters
type ActionDnsZoneIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionDnsZoneIndexMetaGlobalInput) SetCount(value bool) *ActionDnsZoneIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsZoneIndexMetaGlobalInput) SetIncludes(value string) *ActionDnsZoneIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsZoneIndexMetaGlobalInput) SetNo(value bool) *ActionDnsZoneIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsZoneIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsZoneIndexMetaGlobalInput) SelectParameters(params ...string) *ActionDnsZoneIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsZoneIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsZoneIndexInput is a type for action input parameters
type ActionDnsZoneIndexInput struct {
	DnssecEnabled bool   `json:"dnssec_enabled"`
	Enabled       bool   `json:"enabled"`
	FromId        int64  `json:"from_id"`
	Limit         int64  `json:"limit"`
	Role          string `json:"role"`
	Source        string `json:"source"`
	User          int64  `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDnssecEnabled sets parameter DnssecEnabled to value and selects it for sending
func (in *ActionDnsZoneIndexInput) SetDnssecEnabled(value bool) *ActionDnsZoneIndexInput {
	in.DnssecEnabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DnssecEnabled"] = nil
	return in
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionDnsZoneIndexInput) SetEnabled(value bool) *ActionDnsZoneIndexInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionDnsZoneIndexInput) SetFromId(value int64) *ActionDnsZoneIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDnsZoneIndexInput) SetLimit(value int64) *ActionDnsZoneIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetRole sets parameter Role to value and selects it for sending
func (in *ActionDnsZoneIndexInput) SetRole(value string) *ActionDnsZoneIndexInput {
	in.Role = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Role"] = nil
	return in
}

// SetSource sets parameter Source to value and selects it for sending
func (in *ActionDnsZoneIndexInput) SetSource(value string) *ActionDnsZoneIndexInput {
	in.Source = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Source"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionDnsZoneIndexInput) SetUser(value int64) *ActionDnsZoneIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionDnsZoneIndexInput) SetUserNil(set bool) *ActionDnsZoneIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["User"] = nil
		in.SelectParameters("User")
	} else {
		delete(in._nilParameters, "User")
	}
	return in
}

// SelectParameters sets parameters from ActionDnsZoneIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsZoneIndexInput) SelectParameters(params ...string) *ActionDnsZoneIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnsZoneIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnsZoneIndexInput) UnselectParameters(params ...string) *ActionDnsZoneIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnsZoneIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsZoneIndexOutput is a type for action output parameters
type ActionDnsZoneIndexOutput struct {
	CreatedAt             string                `json:"created_at"`
	DefaultTtl            int64                 `json:"default_ttl"`
	DnssecEnabled         bool                  `json:"dnssec_enabled"`
	Email                 string                `json:"email"`
	Enabled               bool                  `json:"enabled"`
	Id                    int64                 `json:"id"`
	Label                 string                `json:"label"`
	Managed               bool                  `json:"managed"`
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
type ActionDnsZoneIndexResponse struct {
	Action *ActionDnsZoneIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsZones []*ActionDnsZoneIndexOutput `json:"dns_zones"`
	}

	// Action output without the namespace
	Output []*ActionDnsZoneIndexOutput
}

// Prepare the action for invocation
func (action *ActionDnsZoneIndex) Prepare() *ActionDnsZoneIndexInvocation {
	return &ActionDnsZoneIndexInvocation{
		Action: action,
		Path:   "/v7.0/dns_zones",
	}
}

// ActionDnsZoneIndexInvocation is used to configure action for invocation
type ActionDnsZoneIndexInvocation struct {
	// Pointer to the action
	Action *ActionDnsZoneIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsZoneIndexInput
	// Global meta input parameters
	MetaInput *ActionDnsZoneIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsZoneIndexInvocation) NewInput() *ActionDnsZoneIndexInput {
	inv.Input = &ActionDnsZoneIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsZoneIndexInvocation) SetInput(input *ActionDnsZoneIndexInput) *ActionDnsZoneIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsZoneIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnsZoneIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsZoneIndexInvocation) NewMetaInput() *ActionDnsZoneIndexMetaGlobalInput {
	inv.MetaInput = &ActionDnsZoneIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsZoneIndexInvocation) SetMetaInput(input *ActionDnsZoneIndexMetaGlobalInput) *ActionDnsZoneIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsZoneIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsZoneIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsZoneIndexInvocation) Call() (*ActionDnsZoneIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDnsZoneIndexInvocation) callAsQuery() (*ActionDnsZoneIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDnsZoneIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsZones
	}
	return resp, err
}

func (inv *ActionDnsZoneIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("DnssecEnabled") {
			ret["dns_zone[dnssec_enabled]"] = convertBoolToString(inv.Input.DnssecEnabled)
		}
		if inv.IsParameterSelected("Enabled") {
			ret["dns_zone[enabled]"] = convertBoolToString(inv.Input.Enabled)
		}
		if inv.IsParameterSelected("FromId") {
			ret["dns_zone[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["dns_zone[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Role") {
			ret["dns_zone[role]"] = inv.Input.Role
		}
		if inv.IsParameterSelected("Source") {
			ret["dns_zone[source]"] = inv.Input.Source
		}
		if inv.IsParameterSelected("User") {
			ret["dns_zone[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionDnsZoneIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
