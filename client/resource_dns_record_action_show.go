package client

import (
	"strings"
)

// ActionDnsRecordShow is a type for action Dns_record#Show
type ActionDnsRecordShow struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsRecordShow(client *Client) *ActionDnsRecordShow {
	return &ActionDnsRecordShow{
		Client: client,
	}
}

// ActionDnsRecordShowMetaGlobalInput is a type for action global meta input parameters
type ActionDnsRecordShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsRecordShowMetaGlobalInput) SetIncludes(value string) *ActionDnsRecordShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsRecordShowMetaGlobalInput) SetNo(value bool) *ActionDnsRecordShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsRecordShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsRecordShowMetaGlobalInput) SelectParameters(params ...string) *ActionDnsRecordShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsRecordShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsRecordShowOutput is a type for action output parameters
type ActionDnsRecordShowOutput struct {
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
type ActionDnsRecordShowResponse struct {
	Action *ActionDnsRecordShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsRecord *ActionDnsRecordShowOutput `json:"dns_record"`
	}

	// Action output without the namespace
	Output *ActionDnsRecordShowOutput
}

// Prepare the action for invocation
func (action *ActionDnsRecordShow) Prepare() *ActionDnsRecordShowInvocation {
	return &ActionDnsRecordShowInvocation{
		Action: action,
		Path:   "/v7.0/dns_records/{dns_record_id}",
	}
}

// ActionDnsRecordShowInvocation is used to configure action for invocation
type ActionDnsRecordShowInvocation struct {
	// Pointer to the action
	Action *ActionDnsRecordShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDnsRecordShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDnsRecordShowInvocation) SetPathParamInt(param string, value int64) *ActionDnsRecordShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDnsRecordShowInvocation) SetPathParamString(param string, value string) *ActionDnsRecordShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsRecordShowInvocation) NewMetaInput() *ActionDnsRecordShowMetaGlobalInput {
	inv.MetaInput = &ActionDnsRecordShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsRecordShowInvocation) SetMetaInput(input *ActionDnsRecordShowMetaGlobalInput) *ActionDnsRecordShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsRecordShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsRecordShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsRecordShowInvocation) Call() (*ActionDnsRecordShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDnsRecordShowInvocation) callAsQuery() (*ActionDnsRecordShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDnsRecordShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsRecord
	}
	return resp, err
}

func (inv *ActionDnsRecordShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
