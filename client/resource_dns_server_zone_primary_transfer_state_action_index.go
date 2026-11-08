package client

import ()

// ActionDnsServerZonePrimaryTransferStateIndex is a type for action Dns_server_zone_primary_transfer_state#Index
type ActionDnsServerZonePrimaryTransferStateIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsServerZonePrimaryTransferStateIndex(client *Client) *ActionDnsServerZonePrimaryTransferStateIndex {
	return &ActionDnsServerZonePrimaryTransferStateIndex{
		Client: client,
	}
}

// ActionDnsServerZonePrimaryTransferStateIndexMetaGlobalInput is a type for action global meta input parameters
type ActionDnsServerZonePrimaryTransferStateIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionDnsServerZonePrimaryTransferStateIndexMetaGlobalInput) SetCount(value bool) *ActionDnsServerZonePrimaryTransferStateIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsServerZonePrimaryTransferStateIndexMetaGlobalInput) SetIncludes(value string) *ActionDnsServerZonePrimaryTransferStateIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsServerZonePrimaryTransferStateIndexMetaGlobalInput) SetNo(value bool) *ActionDnsServerZonePrimaryTransferStateIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsServerZonePrimaryTransferStateIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsServerZonePrimaryTransferStateIndexMetaGlobalInput) SelectParameters(params ...string) *ActionDnsServerZonePrimaryTransferStateIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsServerZonePrimaryTransferStateIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsServerZonePrimaryTransferStateIndexInput is a type for action input parameters
type ActionDnsServerZonePrimaryTransferStateIndexInput struct {
	DnsServerZone   int64  "json:\"dns_server_zone\""
	DnsZone         int64  "json:\"dns_zone\""
	DnsZoneTransfer int64  "json:\"dns_zone_transfer\""
	FailureClass    string "json:\"failure_class\""
	FromId          int64  "json:\"from_id\""
	Limit           int64  "json:\"limit\""
	Status          string "json:\"status\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDnsServerZone sets parameter DnsServerZone to value and selects it for sending
func (in *ActionDnsServerZonePrimaryTransferStateIndexInput) SetDnsServerZone(value int64) *ActionDnsServerZonePrimaryTransferStateIndexInput {
	in.DnsServerZone = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DnsServerZone"] = nil
	return in
}

// SetDnsZone sets parameter DnsZone to value and selects it for sending
func (in *ActionDnsServerZonePrimaryTransferStateIndexInput) SetDnsZone(value int64) *ActionDnsServerZonePrimaryTransferStateIndexInput {
	in.DnsZone = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DnsZone"] = nil
	return in
}

// SetDnsZoneTransfer sets parameter DnsZoneTransfer to value and selects it for sending
func (in *ActionDnsServerZonePrimaryTransferStateIndexInput) SetDnsZoneTransfer(value int64) *ActionDnsServerZonePrimaryTransferStateIndexInput {
	in.DnsZoneTransfer = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DnsZoneTransfer"] = nil
	return in
}

// SetFailureClass sets parameter FailureClass to value and selects it for sending
func (in *ActionDnsServerZonePrimaryTransferStateIndexInput) SetFailureClass(value string) *ActionDnsServerZonePrimaryTransferStateIndexInput {
	in.FailureClass = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetFailureClassNil(false)
	in._selectedParameters["FailureClass"] = nil
	return in
}

// SetFailureClassNil sets parameter FailureClass to nil and selects it for sending
func (in *ActionDnsServerZonePrimaryTransferStateIndexInput) SetFailureClassNil(set bool) *ActionDnsServerZonePrimaryTransferStateIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["FailureClass"] = nil
		in.SelectParameters("FailureClass")
	} else {
		delete(in._nilParameters, "FailureClass")
	}
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionDnsServerZonePrimaryTransferStateIndexInput) SetFromId(value int64) *ActionDnsServerZonePrimaryTransferStateIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDnsServerZonePrimaryTransferStateIndexInput) SetLimit(value int64) *ActionDnsServerZonePrimaryTransferStateIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetStatus sets parameter Status to value and selects it for sending
func (in *ActionDnsServerZonePrimaryTransferStateIndexInput) SetStatus(value string) *ActionDnsServerZonePrimaryTransferStateIndexInput {
	in.Status = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Status"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsServerZonePrimaryTransferStateIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsServerZonePrimaryTransferStateIndexInput) SelectParameters(params ...string) *ActionDnsServerZonePrimaryTransferStateIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnsServerZonePrimaryTransferStateIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnsServerZonePrimaryTransferStateIndexInput) UnselectParameters(params ...string) *ActionDnsServerZonePrimaryTransferStateIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnsServerZonePrimaryTransferStateIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsServerZonePrimaryTransferStateIndexOutput is a type for action output parameters
type ActionDnsServerZonePrimaryTransferStateIndexOutput struct {
	AlertEligibleAt     string                                    "json:\"alert_eligible_at\""
	CreatedAt           string                                    "json:\"created_at\""
	DnsServerZone       *ActionDnsServerZoneShowOutput            "json:\"dns_server_zone\""
	DnsZoneTransfer     *ActionDnsZoneTransferShowOutput          "json:\"dns_zone_transfer\""
	EvidenceCurrent     bool                                      "json:\"evidence_current\""
	FailedSince         string                                    "json:\"failed_since\""
	FailureClass        string                                    "json:\"failure_class\""
	Id                  int64                                     "json:\"id\""
	LastAttemptAt       string                                    "json:\"last_attempt_at\""
	LastAttemptKind     string                                    "json:\"last_attempt_kind\""
	LastSuccessAt       string                                    "json:\"last_success_at\""
	LastTransferLog     *ActionDnsServerZoneTransferLogShowOutput "json:\"last_transfer_log\""
	PrimarySerial       int64                                     "json:\"primary_serial\""
	Reason              string                                    "json:\"reason\""
	ReasonCode          string                                    "json:\"reason_code\""
	ReasonObservedAt    string                                    "json:\"reason_observed_at\""
	SecondarySerial     int64                                     "json:\"secondary_serial\""
	SecondarySourceAddr string                                    "json:\"secondary_source_addr\""
	Status              string                                    "json:\"status\""
	UpdatedAt           string                                    "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionDnsServerZonePrimaryTransferStateIndexResponse struct {
	Action *ActionDnsServerZonePrimaryTransferStateIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsServerZonePrimaryTransferStates []*ActionDnsServerZonePrimaryTransferStateIndexOutput "json:\"dns_server_zone_primary_transfer_states\""
	}

	// Action output without the namespace
	Output []*ActionDnsServerZonePrimaryTransferStateIndexOutput
}

// Prepare the action for invocation
func (action *ActionDnsServerZonePrimaryTransferStateIndex) Prepare() *ActionDnsServerZonePrimaryTransferStateIndexInvocation {
	return &ActionDnsServerZonePrimaryTransferStateIndexInvocation{
		Action: action,
		Path:   "/v7.0/dns_server_zone_primary_transfer_states",
	}
}

// ActionDnsServerZonePrimaryTransferStateIndexInvocation is used to configure action for invocation
type ActionDnsServerZonePrimaryTransferStateIndexInvocation struct {
	// Pointer to the action
	Action *ActionDnsServerZonePrimaryTransferStateIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsServerZonePrimaryTransferStateIndexInput
	// Global meta input parameters
	MetaInput *ActionDnsServerZonePrimaryTransferStateIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsServerZonePrimaryTransferStateIndexInvocation) NewInput() *ActionDnsServerZonePrimaryTransferStateIndexInput {
	inv.Input = &ActionDnsServerZonePrimaryTransferStateIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsServerZonePrimaryTransferStateIndexInvocation) SetInput(input *ActionDnsServerZonePrimaryTransferStateIndexInput) *ActionDnsServerZonePrimaryTransferStateIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsServerZonePrimaryTransferStateIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnsServerZonePrimaryTransferStateIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsServerZonePrimaryTransferStateIndexInvocation) NewMetaInput() *ActionDnsServerZonePrimaryTransferStateIndexMetaGlobalInput {
	inv.MetaInput = &ActionDnsServerZonePrimaryTransferStateIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsServerZonePrimaryTransferStateIndexInvocation) SetMetaInput(input *ActionDnsServerZonePrimaryTransferStateIndexMetaGlobalInput) *ActionDnsServerZonePrimaryTransferStateIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsServerZonePrimaryTransferStateIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsServerZonePrimaryTransferStateIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionDnsServerZonePrimaryTransferStateIndexInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.Input != nil {
		if inv.IsParameterSelected("DnsServerZone") {
			if !inv.IsParameterNil("DnsServerZone") {
				if inv.Input.DnsServerZone < 0 {
					verr.Add("dns_server_zone", t(inv.Action.Client, "validation.invalid_resource_id", nil))
				}
			}
		}
		if inv.IsParameterSelected("DnsZone") {
			if !inv.IsParameterNil("DnsZone") {
				if inv.Input.DnsZone < 0 {
					verr.Add("dns_zone", t(inv.Action.Client, "validation.invalid_resource_id", nil))
				}
			}
		}
		if inv.IsParameterSelected("DnsZoneTransfer") {
			if !inv.IsParameterNil("DnsZoneTransfer") {
				if inv.Input.DnsZoneTransfer < 0 {
					verr.Add("dns_zone_transfer", t(inv.Action.Client, "validation.invalid_resource_id", nil))
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
func (inv *ActionDnsServerZonePrimaryTransferStateIndexInvocation) Call() (*ActionDnsServerZonePrimaryTransferStateIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionDnsServerZonePrimaryTransferStateIndexInvocation) callAsQuery() (*ActionDnsServerZonePrimaryTransferStateIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDnsServerZonePrimaryTransferStateIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsServerZonePrimaryTransferStates
	}
	return resp, err
}

func (inv *ActionDnsServerZonePrimaryTransferStateIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("DnsServerZone") {
			ret["dns_server_zone_primary_transfer_state[dns_server_zone]"] = convertInt64ToString(inv.Input.DnsServerZone)
		}
		if inv.IsParameterSelected("DnsZone") {
			ret["dns_server_zone_primary_transfer_state[dns_zone]"] = convertInt64ToString(inv.Input.DnsZone)
		}
		if inv.IsParameterSelected("DnsZoneTransfer") {
			ret["dns_server_zone_primary_transfer_state[dns_zone_transfer]"] = convertInt64ToString(inv.Input.DnsZoneTransfer)
		}
		if inv.IsParameterSelected("FailureClass") {
			if inv.IsParameterNil("FailureClass") {
				ret["dns_server_zone_primary_transfer_state[failure_class]"] = ""
			} else {
				ret["dns_server_zone_primary_transfer_state[failure_class]"] = inv.Input.FailureClass
			}
		}
		if inv.IsParameterSelected("FromId") {
			ret["dns_server_zone_primary_transfer_state[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["dns_server_zone_primary_transfer_state[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Status") {
			ret["dns_server_zone_primary_transfer_state[status]"] = inv.Input.Status
		}
	}
}

func (inv *ActionDnsServerZonePrimaryTransferStateIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
