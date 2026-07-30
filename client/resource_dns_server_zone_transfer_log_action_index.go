package client

import ()

// ActionDnsServerZoneTransferLogIndex is a type for action Dns_server_zone_transfer_log#Index
type ActionDnsServerZoneTransferLogIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsServerZoneTransferLogIndex(client *Client) *ActionDnsServerZoneTransferLogIndex {
	return &ActionDnsServerZoneTransferLogIndex{
		Client: client,
	}
}

// ActionDnsServerZoneTransferLogIndexMetaGlobalInput is a type for action global meta input parameters
type ActionDnsServerZoneTransferLogIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionDnsServerZoneTransferLogIndexMetaGlobalInput) SetCount(value bool) *ActionDnsServerZoneTransferLogIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsServerZoneTransferLogIndexMetaGlobalInput) SetIncludes(value string) *ActionDnsServerZoneTransferLogIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsServerZoneTransferLogIndexMetaGlobalInput) SetNo(value bool) *ActionDnsServerZoneTransferLogIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsServerZoneTransferLogIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsServerZoneTransferLogIndexMetaGlobalInput) SelectParameters(params ...string) *ActionDnsServerZoneTransferLogIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsServerZoneTransferLogIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsServerZoneTransferLogIndexInput is a type for action input parameters
type ActionDnsServerZoneTransferLogIndexInput struct {
	DnsServerZone int64  "json:\"dns_server_zone\""
	DnsZone       int64  "json:\"dns_zone\""
	FromId        int64  "json:\"from_id\""
	Limit         int64  "json:\"limit\""
	Order         string "json:\"order\""
	PrimaryAddr   string "json:\"primary_addr\""
	ReasonCode    string "json:\"reason_code\""
	Status        string "json:\"status\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDnsServerZone sets parameter DnsServerZone to value and selects it for sending
func (in *ActionDnsServerZoneTransferLogIndexInput) SetDnsServerZone(value int64) *ActionDnsServerZoneTransferLogIndexInput {
	in.DnsServerZone = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DnsServerZone"] = nil
	return in
}

// SetDnsZone sets parameter DnsZone to value and selects it for sending
func (in *ActionDnsServerZoneTransferLogIndexInput) SetDnsZone(value int64) *ActionDnsServerZoneTransferLogIndexInput {
	in.DnsZone = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DnsZone"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionDnsServerZoneTransferLogIndexInput) SetFromId(value int64) *ActionDnsServerZoneTransferLogIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDnsServerZoneTransferLogIndexInput) SetLimit(value int64) *ActionDnsServerZoneTransferLogIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOrder sets parameter Order to value and selects it for sending
func (in *ActionDnsServerZoneTransferLogIndexInput) SetOrder(value string) *ActionDnsServerZoneTransferLogIndexInput {
	in.Order = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Order"] = nil
	return in
}

// SetPrimaryAddr sets parameter PrimaryAddr to value and selects it for sending
func (in *ActionDnsServerZoneTransferLogIndexInput) SetPrimaryAddr(value string) *ActionDnsServerZoneTransferLogIndexInput {
	in.PrimaryAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["PrimaryAddr"] = nil
	return in
}

// SetReasonCode sets parameter ReasonCode to value and selects it for sending
func (in *ActionDnsServerZoneTransferLogIndexInput) SetReasonCode(value string) *ActionDnsServerZoneTransferLogIndexInput {
	in.ReasonCode = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ReasonCode"] = nil
	return in
}

// SetStatus sets parameter Status to value and selects it for sending
func (in *ActionDnsServerZoneTransferLogIndexInput) SetStatus(value string) *ActionDnsServerZoneTransferLogIndexInput {
	in.Status = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Status"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsServerZoneTransferLogIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsServerZoneTransferLogIndexInput) SelectParameters(params ...string) *ActionDnsServerZoneTransferLogIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnsServerZoneTransferLogIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnsServerZoneTransferLogIndexInput) UnselectParameters(params ...string) *ActionDnsServerZoneTransferLogIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnsServerZoneTransferLogIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsServerZoneTransferLogIndexOutput is a type for action output parameters
type ActionDnsServerZoneTransferLogIndexOutput struct {
	CreatedAt     string                         "json:\"created_at\""
	DnsServerZone *ActionDnsServerZoneShowOutput "json:\"dns_server_zone\""
	EventAt       string                         "json:\"event_at\""
	EventKey      string                         "json:\"event_key\""
	Id            int64                          "json:\"id\""
	Message       string                         "json:\"message\""
	PrimaryAddr   string                         "json:\"primary_addr\""
	RawMessage    string                         "json:\"raw_message\""
	Reason        string                         "json:\"reason\""
	ReasonCode    string                         "json:\"reason_code\""
	Serial        int64                          "json:\"serial\""
	SourceCursor  string                         "json:\"source_cursor\""
	Status        string                         "json:\"status\""
	UpdatedAt     string                         "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionDnsServerZoneTransferLogIndexResponse struct {
	Action *ActionDnsServerZoneTransferLogIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsServerZoneTransferLogs []*ActionDnsServerZoneTransferLogIndexOutput "json:\"dns_server_zone_transfer_logs\""
	}

	// Action output without the namespace
	Output []*ActionDnsServerZoneTransferLogIndexOutput
}

// Prepare the action for invocation
func (action *ActionDnsServerZoneTransferLogIndex) Prepare() *ActionDnsServerZoneTransferLogIndexInvocation {
	return &ActionDnsServerZoneTransferLogIndexInvocation{
		Action: action,
		Path:   "/v7.0/dns_server_zone_transfer_logs",
	}
}

// ActionDnsServerZoneTransferLogIndexInvocation is used to configure action for invocation
type ActionDnsServerZoneTransferLogIndexInvocation struct {
	// Pointer to the action
	Action *ActionDnsServerZoneTransferLogIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsServerZoneTransferLogIndexInput
	// Global meta input parameters
	MetaInput *ActionDnsServerZoneTransferLogIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsServerZoneTransferLogIndexInvocation) NewInput() *ActionDnsServerZoneTransferLogIndexInput {
	inv.Input = &ActionDnsServerZoneTransferLogIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsServerZoneTransferLogIndexInvocation) SetInput(input *ActionDnsServerZoneTransferLogIndexInput) *ActionDnsServerZoneTransferLogIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsServerZoneTransferLogIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnsServerZoneTransferLogIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsServerZoneTransferLogIndexInvocation) NewMetaInput() *ActionDnsServerZoneTransferLogIndexMetaGlobalInput {
	inv.MetaInput = &ActionDnsServerZoneTransferLogIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsServerZoneTransferLogIndexInvocation) SetMetaInput(input *ActionDnsServerZoneTransferLogIndexMetaGlobalInput) *ActionDnsServerZoneTransferLogIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsServerZoneTransferLogIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsServerZoneTransferLogIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionDnsServerZoneTransferLogIndexInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("DnsServerZone") {
			if !inv.IsParameterNil("DnsServerZone") {
				if inv.Input.DnsServerZone < 0 {
					verr.Add("dns_server_zone", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("DnsZone") {
			if !inv.IsParameterNil("DnsZone") {
				if inv.Input.DnsZone < 0 {
					verr.Add("dns_zone", "not a valid resource id")
				}
			}
		}
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsServerZoneTransferLogIndexInvocation) Call() (*ActionDnsServerZoneTransferLogIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionDnsServerZoneTransferLogIndexInvocation) callAsQuery() (*ActionDnsServerZoneTransferLogIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionDnsServerZoneTransferLogIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsServerZoneTransferLogs
	}
	return resp, err
}

func (inv *ActionDnsServerZoneTransferLogIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("DnsServerZone") {
			ret["dns_server_zone_transfer_log[dns_server_zone]"] = convertInt64ToString(inv.Input.DnsServerZone)
		}
		if inv.IsParameterSelected("DnsZone") {
			ret["dns_server_zone_transfer_log[dns_zone]"] = convertInt64ToString(inv.Input.DnsZone)
		}
		if inv.IsParameterSelected("FromId") {
			ret["dns_server_zone_transfer_log[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["dns_server_zone_transfer_log[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Order") {
			ret["dns_server_zone_transfer_log[order]"] = inv.Input.Order
		}
		if inv.IsParameterSelected("PrimaryAddr") {
			ret["dns_server_zone_transfer_log[primary_addr]"] = inv.Input.PrimaryAddr
		}
		if inv.IsParameterSelected("ReasonCode") {
			ret["dns_server_zone_transfer_log[reason_code]"] = inv.Input.ReasonCode
		}
		if inv.IsParameterSelected("Status") {
			ret["dns_server_zone_transfer_log[status]"] = inv.Input.Status
		}
	}

	return nil
}

func (inv *ActionDnsServerZoneTransferLogIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			queryValue, err := convertCustomToString(inv.MetaInput.Includes)
			if err != nil {
				return err
			}
			ret["_meta[includes]"] = queryValue
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}

	return nil
}
