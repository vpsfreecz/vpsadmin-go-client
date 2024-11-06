package client

import ()

// ActionDnssecRecordIndex is a type for action Dnssec_record#Index
type ActionDnssecRecordIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionDnssecRecordIndex(client *Client) *ActionDnssecRecordIndex {
	return &ActionDnssecRecordIndex{
		Client: client,
	}
}

// ActionDnssecRecordIndexMetaGlobalInput is a type for action global meta input parameters
type ActionDnssecRecordIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionDnssecRecordIndexMetaGlobalInput) SetCount(value bool) *ActionDnssecRecordIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnssecRecordIndexMetaGlobalInput) SetIncludes(value string) *ActionDnssecRecordIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnssecRecordIndexMetaGlobalInput) SetNo(value bool) *ActionDnssecRecordIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnssecRecordIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnssecRecordIndexMetaGlobalInput) SelectParameters(params ...string) *ActionDnssecRecordIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnssecRecordIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnssecRecordIndexInput is a type for action input parameters
type ActionDnssecRecordIndexInput struct {
	DnsZone int64 `json:"dns_zone"`
	FromId  int64 `json:"from_id"`
	Limit   int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDnsZone sets parameter DnsZone to value and selects it for sending
func (in *ActionDnssecRecordIndexInput) SetDnsZone(value int64) *ActionDnssecRecordIndexInput {
	in.DnsZone = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetDnsZoneNil(false)
	in._selectedParameters["DnsZone"] = nil
	return in
}

// SetDnsZoneNil sets parameter DnsZone to nil and selects it for sending
func (in *ActionDnssecRecordIndexInput) SetDnsZoneNil(set bool) *ActionDnssecRecordIndexInput {
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
func (in *ActionDnssecRecordIndexInput) SetFromId(value int64) *ActionDnssecRecordIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDnssecRecordIndexInput) SetLimit(value int64) *ActionDnssecRecordIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnssecRecordIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnssecRecordIndexInput) SelectParameters(params ...string) *ActionDnssecRecordIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnssecRecordIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnssecRecordIndexInput) UnselectParameters(params ...string) *ActionDnssecRecordIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnssecRecordIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnssecRecordIndexOutput is a type for action output parameters
type ActionDnssecRecordIndexOutput struct {
	CreatedAt       string                   `json:"created_at"`
	DnsZone         *ActionDnsZoneShowOutput `json:"dns_zone"`
	DnskeyAlgorithm int64                    `json:"dnskey_algorithm"`
	DnskeyPubkey    string                   `json:"dnskey_pubkey"`
	DsAlgorithm     int64                    `json:"ds_algorithm"`
	DsDigest        string                   `json:"ds_digest"`
	DsDigestType    int64                    `json:"ds_digest_type"`
	Id              int64                    `json:"id"`
	Keyid           int64                    `json:"keyid"`
	UpdatedAt       string                   `json:"updated_at"`
}

// Type for action response, including envelope
type ActionDnssecRecordIndexResponse struct {
	Action *ActionDnssecRecordIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnssecRecords []*ActionDnssecRecordIndexOutput `json:"dnssec_records"`
	}

	// Action output without the namespace
	Output []*ActionDnssecRecordIndexOutput
}

// Prepare the action for invocation
func (action *ActionDnssecRecordIndex) Prepare() *ActionDnssecRecordIndexInvocation {
	return &ActionDnssecRecordIndexInvocation{
		Action: action,
		Path:   "/v7.0/dnssec_records",
	}
}

// ActionDnssecRecordIndexInvocation is used to configure action for invocation
type ActionDnssecRecordIndexInvocation struct {
	// Pointer to the action
	Action *ActionDnssecRecordIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnssecRecordIndexInput
	// Global meta input parameters
	MetaInput *ActionDnssecRecordIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnssecRecordIndexInvocation) NewInput() *ActionDnssecRecordIndexInput {
	inv.Input = &ActionDnssecRecordIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnssecRecordIndexInvocation) SetInput(input *ActionDnssecRecordIndexInput) *ActionDnssecRecordIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnssecRecordIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnssecRecordIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnssecRecordIndexInvocation) NewMetaInput() *ActionDnssecRecordIndexMetaGlobalInput {
	inv.MetaInput = &ActionDnssecRecordIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnssecRecordIndexInvocation) SetMetaInput(input *ActionDnssecRecordIndexMetaGlobalInput) *ActionDnssecRecordIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnssecRecordIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnssecRecordIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnssecRecordIndexInvocation) Call() (*ActionDnssecRecordIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDnssecRecordIndexInvocation) callAsQuery() (*ActionDnssecRecordIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDnssecRecordIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnssecRecords
	}
	return resp, err
}

func (inv *ActionDnssecRecordIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("DnsZone") {
			ret["dnssec_record[dns_zone]"] = convertInt64ToString(inv.Input.DnsZone)
		}
		if inv.IsParameterSelected("FromId") {
			ret["dnssec_record[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["dnssec_record[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionDnssecRecordIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
