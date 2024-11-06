package client

import (
	"strings"
)

// ActionDnssecRecordShow is a type for action Dnssec_record#Show
type ActionDnssecRecordShow struct {
	// Pointer to client
	Client *Client
}

func NewActionDnssecRecordShow(client *Client) *ActionDnssecRecordShow {
	return &ActionDnssecRecordShow{
		Client: client,
	}
}

// ActionDnssecRecordShowMetaGlobalInput is a type for action global meta input parameters
type ActionDnssecRecordShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnssecRecordShowMetaGlobalInput) SetIncludes(value string) *ActionDnssecRecordShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnssecRecordShowMetaGlobalInput) SetNo(value bool) *ActionDnssecRecordShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnssecRecordShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnssecRecordShowMetaGlobalInput) SelectParameters(params ...string) *ActionDnssecRecordShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnssecRecordShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnssecRecordShowOutput is a type for action output parameters
type ActionDnssecRecordShowOutput struct {
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
type ActionDnssecRecordShowResponse struct {
	Action *ActionDnssecRecordShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnssecRecord *ActionDnssecRecordShowOutput `json:"dnssec_record"`
	}

	// Action output without the namespace
	Output *ActionDnssecRecordShowOutput
}

// Prepare the action for invocation
func (action *ActionDnssecRecordShow) Prepare() *ActionDnssecRecordShowInvocation {
	return &ActionDnssecRecordShowInvocation{
		Action: action,
		Path:   "/v7.0/dnssec_records/{dnssec_record_id}",
	}
}

// ActionDnssecRecordShowInvocation is used to configure action for invocation
type ActionDnssecRecordShowInvocation struct {
	// Pointer to the action
	Action *ActionDnssecRecordShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDnssecRecordShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDnssecRecordShowInvocation) SetPathParamInt(param string, value int64) *ActionDnssecRecordShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDnssecRecordShowInvocation) SetPathParamString(param string, value string) *ActionDnssecRecordShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnssecRecordShowInvocation) NewMetaInput() *ActionDnssecRecordShowMetaGlobalInput {
	inv.MetaInput = &ActionDnssecRecordShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnssecRecordShowInvocation) SetMetaInput(input *ActionDnssecRecordShowMetaGlobalInput) *ActionDnssecRecordShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnssecRecordShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnssecRecordShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnssecRecordShowInvocation) Call() (*ActionDnssecRecordShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDnssecRecordShowInvocation) callAsQuery() (*ActionDnssecRecordShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDnssecRecordShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnssecRecord
	}
	return resp, err
}

func (inv *ActionDnssecRecordShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
