package client

import ()

// ActionDnsRecordIndex is a type for action Dns_record#Index
type ActionDnsRecordIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsRecordIndex(client *Client) *ActionDnsRecordIndex {
	return &ActionDnsRecordIndex{
		Client: client,
	}
}

// ActionDnsRecordIndexMetaGlobalInput is a type for action global meta input parameters
type ActionDnsRecordIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionDnsRecordIndexMetaGlobalInput) SetCount(value bool) *ActionDnsRecordIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsRecordIndexMetaGlobalInput) SetIncludes(value string) *ActionDnsRecordIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsRecordIndexMetaGlobalInput) SetNo(value bool) *ActionDnsRecordIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsRecordIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsRecordIndexMetaGlobalInput) SelectParameters(params ...string) *ActionDnsRecordIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsRecordIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsRecordIndexInput is a type for action input parameters
type ActionDnsRecordIndexInput struct {
	DnsZone int64 `json:"dns_zone"`
	FromId  int64 `json:"from_id"`
	Limit   int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDnsZone sets parameter DnsZone to value and selects it for sending
func (in *ActionDnsRecordIndexInput) SetDnsZone(value int64) *ActionDnsRecordIndexInput {
	in.DnsZone = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetDnsZoneNil(false)
	in._selectedParameters["DnsZone"] = nil
	return in
}

// SetDnsZoneNil sets parameter DnsZone to nil and selects it for sending
func (in *ActionDnsRecordIndexInput) SetDnsZoneNil(set bool) *ActionDnsRecordIndexInput {
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
func (in *ActionDnsRecordIndexInput) SetFromId(value int64) *ActionDnsRecordIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDnsRecordIndexInput) SetLimit(value int64) *ActionDnsRecordIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsRecordIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsRecordIndexInput) SelectParameters(params ...string) *ActionDnsRecordIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnsRecordIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnsRecordIndexInput) UnselectParameters(params ...string) *ActionDnsRecordIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnsRecordIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsRecordIndexOutput is a type for action output parameters
type ActionDnsRecordIndexOutput struct {
	Comment              string                   `json:"comment"`
	Content              string                   `json:"content"`
	CreatedAt            string                   `json:"created_at"`
	DnsZone              *ActionDnsZoneShowOutput `json:"dns_zone"`
	DynamicUpdateEnabled bool                     `json:"dynamic_update_enabled"`
	DynamicUpdateUrl     string                   `json:"dynamic_update_url"`
	Enabled              bool                     `json:"enabled"`
	Id                   int64                    `json:"id"`
	Name                 string                   `json:"name"`
	Priority             int64                    `json:"priority"`
	Ttl                  int64                    `json:"ttl"`
	Type                 string                   `json:"type"`
	UpdatedAt            string                   `json:"updated_at"`
}

// Type for action response, including envelope
type ActionDnsRecordIndexResponse struct {
	Action *ActionDnsRecordIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsRecords []*ActionDnsRecordIndexOutput `json:"dns_records"`
	}

	// Action output without the namespace
	Output []*ActionDnsRecordIndexOutput
}

// Prepare the action for invocation
func (action *ActionDnsRecordIndex) Prepare() *ActionDnsRecordIndexInvocation {
	return &ActionDnsRecordIndexInvocation{
		Action: action,
		Path:   "/v7.0/dns_records",
	}
}

// ActionDnsRecordIndexInvocation is used to configure action for invocation
type ActionDnsRecordIndexInvocation struct {
	// Pointer to the action
	Action *ActionDnsRecordIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsRecordIndexInput
	// Global meta input parameters
	MetaInput *ActionDnsRecordIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsRecordIndexInvocation) NewInput() *ActionDnsRecordIndexInput {
	inv.Input = &ActionDnsRecordIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsRecordIndexInvocation) SetInput(input *ActionDnsRecordIndexInput) *ActionDnsRecordIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsRecordIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnsRecordIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsRecordIndexInvocation) NewMetaInput() *ActionDnsRecordIndexMetaGlobalInput {
	inv.MetaInput = &ActionDnsRecordIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsRecordIndexInvocation) SetMetaInput(input *ActionDnsRecordIndexMetaGlobalInput) *ActionDnsRecordIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsRecordIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsRecordIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsRecordIndexInvocation) Call() (*ActionDnsRecordIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDnsRecordIndexInvocation) callAsQuery() (*ActionDnsRecordIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDnsRecordIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsRecords
	}
	return resp, err
}

func (inv *ActionDnsRecordIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("DnsZone") {
			ret["dns_record[dns_zone]"] = convertInt64ToString(inv.Input.DnsZone)
		}
		if inv.IsParameterSelected("FromId") {
			ret["dns_record[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["dns_record[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionDnsRecordIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
