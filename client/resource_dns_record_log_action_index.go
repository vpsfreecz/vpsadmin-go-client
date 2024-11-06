package client

import ()

// ActionDnsRecordLogIndex is a type for action Dns_record_log#Index
type ActionDnsRecordLogIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsRecordLogIndex(client *Client) *ActionDnsRecordLogIndex {
	return &ActionDnsRecordLogIndex{
		Client: client,
	}
}

// ActionDnsRecordLogIndexMetaGlobalInput is a type for action global meta input parameters
type ActionDnsRecordLogIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionDnsRecordLogIndexMetaGlobalInput) SetCount(value bool) *ActionDnsRecordLogIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsRecordLogIndexMetaGlobalInput) SetIncludes(value string) *ActionDnsRecordLogIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsRecordLogIndexMetaGlobalInput) SetNo(value bool) *ActionDnsRecordLogIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsRecordLogIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsRecordLogIndexMetaGlobalInput) SelectParameters(params ...string) *ActionDnsRecordLogIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsRecordLogIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsRecordLogIndexInput is a type for action input parameters
type ActionDnsRecordLogIndexInput struct {
	ChangeType  string `json:"change_type"`
	DnsZone     int64  `json:"dns_zone"`
	DnsZoneName string `json:"dns_zone_name"`
	FromId      int64  `json:"from_id"`
	Limit       int64  `json:"limit"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	User        int64  `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetChangeType sets parameter ChangeType to value and selects it for sending
func (in *ActionDnsRecordLogIndexInput) SetChangeType(value string) *ActionDnsRecordLogIndexInput {
	in.ChangeType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ChangeType"] = nil
	return in
}

// SetDnsZone sets parameter DnsZone to value and selects it for sending
func (in *ActionDnsRecordLogIndexInput) SetDnsZone(value int64) *ActionDnsRecordLogIndexInput {
	in.DnsZone = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetDnsZoneNil(false)
	in._selectedParameters["DnsZone"] = nil
	return in
}

// SetDnsZoneNil sets parameter DnsZone to nil and selects it for sending
func (in *ActionDnsRecordLogIndexInput) SetDnsZoneNil(set bool) *ActionDnsRecordLogIndexInput {
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

// SetDnsZoneName sets parameter DnsZoneName to value and selects it for sending
func (in *ActionDnsRecordLogIndexInput) SetDnsZoneName(value string) *ActionDnsRecordLogIndexInput {
	in.DnsZoneName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DnsZoneName"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionDnsRecordLogIndexInput) SetFromId(value int64) *ActionDnsRecordLogIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDnsRecordLogIndexInput) SetLimit(value int64) *ActionDnsRecordLogIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionDnsRecordLogIndexInput) SetName(value string) *ActionDnsRecordLogIndexInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetType sets parameter Type to value and selects it for sending
func (in *ActionDnsRecordLogIndexInput) SetType(value string) *ActionDnsRecordLogIndexInput {
	in.Type = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Type"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionDnsRecordLogIndexInput) SetUser(value int64) *ActionDnsRecordLogIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionDnsRecordLogIndexInput) SetUserNil(set bool) *ActionDnsRecordLogIndexInput {
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

// SelectParameters sets parameters from ActionDnsRecordLogIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsRecordLogIndexInput) SelectParameters(params ...string) *ActionDnsRecordLogIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnsRecordLogIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnsRecordLogIndexInput) UnselectParameters(params ...string) *ActionDnsRecordLogIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnsRecordLogIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsRecordLogIndexOutput is a type for action output parameters
type ActionDnsRecordLogIndexOutput struct {
	ChangeType       string                            `json:"change_type"`
	CreatedAt        string                            `json:"created_at"`
	DnsZone          *ActionDnsZoneShowOutput          `json:"dns_zone"`
	DnsZoneName      string                            `json:"dns_zone_name"`
	Id               int64                             `json:"id"`
	Name             string                            `json:"name"`
	RawUserId        int64                             `json:"raw_user_id"`
	TransactionChain *ActionTransactionChainShowOutput `json:"transaction_chain"`
	Type             string                            `json:"type"`
	UpdatedAt        string                            `json:"updated_at"`
	User             *ActionUserShowOutput             `json:"user"`
}

// Type for action response, including envelope
type ActionDnsRecordLogIndexResponse struct {
	Action *ActionDnsRecordLogIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsRecordLogs []*ActionDnsRecordLogIndexOutput `json:"dns_record_logs"`
	}

	// Action output without the namespace
	Output []*ActionDnsRecordLogIndexOutput
}

// Prepare the action for invocation
func (action *ActionDnsRecordLogIndex) Prepare() *ActionDnsRecordLogIndexInvocation {
	return &ActionDnsRecordLogIndexInvocation{
		Action: action,
		Path:   "/v7.0/dns_record_logs",
	}
}

// ActionDnsRecordLogIndexInvocation is used to configure action for invocation
type ActionDnsRecordLogIndexInvocation struct {
	// Pointer to the action
	Action *ActionDnsRecordLogIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsRecordLogIndexInput
	// Global meta input parameters
	MetaInput *ActionDnsRecordLogIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsRecordLogIndexInvocation) NewInput() *ActionDnsRecordLogIndexInput {
	inv.Input = &ActionDnsRecordLogIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsRecordLogIndexInvocation) SetInput(input *ActionDnsRecordLogIndexInput) *ActionDnsRecordLogIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsRecordLogIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnsRecordLogIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsRecordLogIndexInvocation) NewMetaInput() *ActionDnsRecordLogIndexMetaGlobalInput {
	inv.MetaInput = &ActionDnsRecordLogIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsRecordLogIndexInvocation) SetMetaInput(input *ActionDnsRecordLogIndexMetaGlobalInput) *ActionDnsRecordLogIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsRecordLogIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsRecordLogIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsRecordLogIndexInvocation) Call() (*ActionDnsRecordLogIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDnsRecordLogIndexInvocation) callAsQuery() (*ActionDnsRecordLogIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDnsRecordLogIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsRecordLogs
	}
	return resp, err
}

func (inv *ActionDnsRecordLogIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("ChangeType") {
			ret["dns_record_log[change_type]"] = inv.Input.ChangeType
		}
		if inv.IsParameterSelected("DnsZone") {
			ret["dns_record_log[dns_zone]"] = convertInt64ToString(inv.Input.DnsZone)
		}
		if inv.IsParameterSelected("DnsZoneName") {
			ret["dns_record_log[dns_zone_name]"] = inv.Input.DnsZoneName
		}
		if inv.IsParameterSelected("FromId") {
			ret["dns_record_log[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["dns_record_log[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Name") {
			ret["dns_record_log[name]"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Type") {
			ret["dns_record_log[type]"] = inv.Input.Type
		}
		if inv.IsParameterSelected("User") {
			ret["dns_record_log[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionDnsRecordLogIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
