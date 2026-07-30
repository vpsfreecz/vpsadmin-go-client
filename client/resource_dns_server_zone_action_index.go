package client

import ()

// ActionDnsServerZoneIndex is a type for action Dns_server_zone#Index
type ActionDnsServerZoneIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsServerZoneIndex(client *Client) *ActionDnsServerZoneIndex {
	return &ActionDnsServerZoneIndex{
		Client: client,
	}
}

// ActionDnsServerZoneIndexMetaGlobalInput is a type for action global meta input parameters
type ActionDnsServerZoneIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionDnsServerZoneIndexMetaGlobalInput) SetCount(value bool) *ActionDnsServerZoneIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsServerZoneIndexMetaGlobalInput) SetIncludes(value string) *ActionDnsServerZoneIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsServerZoneIndexMetaGlobalInput) SetNo(value bool) *ActionDnsServerZoneIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsServerZoneIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsServerZoneIndexMetaGlobalInput) SelectParameters(params ...string) *ActionDnsServerZoneIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsServerZoneIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsServerZoneIndexInput is a type for action input parameters
type ActionDnsServerZoneIndexInput struct {
	DnsServer int64  "json:\"dns_server\""
	DnsZone   int64  "json:\"dns_zone\""
	FromId    int64  "json:\"from_id\""
	Limit     int64  "json:\"limit\""
	Type      string "json:\"type\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDnsServer sets parameter DnsServer to value and selects it for sending
func (in *ActionDnsServerZoneIndexInput) SetDnsServer(value int64) *ActionDnsServerZoneIndexInput {
	in.DnsServer = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DnsServer"] = nil
	return in
}

// SetDnsZone sets parameter DnsZone to value and selects it for sending
func (in *ActionDnsServerZoneIndexInput) SetDnsZone(value int64) *ActionDnsServerZoneIndexInput {
	in.DnsZone = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DnsZone"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionDnsServerZoneIndexInput) SetFromId(value int64) *ActionDnsServerZoneIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDnsServerZoneIndexInput) SetLimit(value int64) *ActionDnsServerZoneIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetType sets parameter Type to value and selects it for sending
func (in *ActionDnsServerZoneIndexInput) SetType(value string) *ActionDnsServerZoneIndexInput {
	in.Type = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Type"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsServerZoneIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsServerZoneIndexInput) SelectParameters(params ...string) *ActionDnsServerZoneIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnsServerZoneIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnsServerZoneIndexInput) UnselectParameters(params ...string) *ActionDnsServerZoneIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnsServerZoneIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsServerZoneIndexOutput is a type for action output parameters
type ActionDnsServerZoneIndexOutput struct {
	CreatedAt               string                     "json:\"created_at\""
	DnsServer               *ActionDnsServerShowOutput "json:\"dns_server\""
	DnsZone                 *ActionDnsZoneShowOutput   "json:\"dns_zone\""
	ExpiresAt               string                     "json:\"expires_at\""
	Id                      int64                      "json:\"id\""
	LastCheckAt             string                     "json:\"last_check_at\""
	LastTransferAt          string                     "json:\"last_transfer_at\""
	LastTransferLogId       int64                      "json:\"last_transfer_log_id\""
	LastTransferPrimaryAddr string                     "json:\"last_transfer_primary_addr\""
	LastTransferReason      string                     "json:\"last_transfer_reason\""
	LastTransferReasonCode  string                     "json:\"last_transfer_reason_code\""
	LastTransferSerial      int64                      "json:\"last_transfer_serial\""
	LastTransferStatus      string                     "json:\"last_transfer_status\""
	LoadedAt                string                     "json:\"loaded_at\""
	RefreshAt               string                     "json:\"refresh_at\""
	Serial                  int64                      "json:\"serial\""
	Type                    string                     "json:\"type\""
	UpdatedAt               string                     "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionDnsServerZoneIndexResponse struct {
	Action *ActionDnsServerZoneIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsServerZones []*ActionDnsServerZoneIndexOutput "json:\"dns_server_zones\""
	}

	// Action output without the namespace
	Output []*ActionDnsServerZoneIndexOutput
}

// Prepare the action for invocation
func (action *ActionDnsServerZoneIndex) Prepare() *ActionDnsServerZoneIndexInvocation {
	return &ActionDnsServerZoneIndexInvocation{
		Action: action,
		Path:   "/v7.0/dns_server_zones",
	}
}

// ActionDnsServerZoneIndexInvocation is used to configure action for invocation
type ActionDnsServerZoneIndexInvocation struct {
	// Pointer to the action
	Action *ActionDnsServerZoneIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsServerZoneIndexInput
	// Global meta input parameters
	MetaInput *ActionDnsServerZoneIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsServerZoneIndexInvocation) NewInput() *ActionDnsServerZoneIndexInput {
	inv.Input = &ActionDnsServerZoneIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsServerZoneIndexInvocation) SetInput(input *ActionDnsServerZoneIndexInput) *ActionDnsServerZoneIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsServerZoneIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnsServerZoneIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsServerZoneIndexInvocation) NewMetaInput() *ActionDnsServerZoneIndexMetaGlobalInput {
	inv.MetaInput = &ActionDnsServerZoneIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsServerZoneIndexInvocation) SetMetaInput(input *ActionDnsServerZoneIndexMetaGlobalInput) *ActionDnsServerZoneIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsServerZoneIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsServerZoneIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionDnsServerZoneIndexInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("DnsServer") {
			if !inv.IsParameterNil("DnsServer") {
				if inv.Input.DnsServer < 0 {
					verr.Add("dns_server", "not a valid resource id")
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
func (inv *ActionDnsServerZoneIndexInvocation) Call() (*ActionDnsServerZoneIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionDnsServerZoneIndexInvocation) callAsQuery() (*ActionDnsServerZoneIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionDnsServerZoneIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsServerZones
	}
	return resp, err
}

func (inv *ActionDnsServerZoneIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("DnsServer") {
			ret["dns_server_zone[dns_server]"] = convertInt64ToString(inv.Input.DnsServer)
		}
		if inv.IsParameterSelected("DnsZone") {
			ret["dns_server_zone[dns_zone]"] = convertInt64ToString(inv.Input.DnsZone)
		}
		if inv.IsParameterSelected("FromId") {
			ret["dns_server_zone[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["dns_server_zone[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Type") {
			ret["dns_server_zone[type]"] = inv.Input.Type
		}
	}

	return nil
}

func (inv *ActionDnsServerZoneIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
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
