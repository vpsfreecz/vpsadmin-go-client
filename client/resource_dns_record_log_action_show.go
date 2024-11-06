package client

import (
	"strings"
)

// ActionDnsRecordLogShow is a type for action Dns_record_log#Show
type ActionDnsRecordLogShow struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsRecordLogShow(client *Client) *ActionDnsRecordLogShow {
	return &ActionDnsRecordLogShow{
		Client: client,
	}
}

// ActionDnsRecordLogShowMetaGlobalInput is a type for action global meta input parameters
type ActionDnsRecordLogShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsRecordLogShowMetaGlobalInput) SetIncludes(value string) *ActionDnsRecordLogShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsRecordLogShowMetaGlobalInput) SetNo(value bool) *ActionDnsRecordLogShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsRecordLogShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsRecordLogShowMetaGlobalInput) SelectParameters(params ...string) *ActionDnsRecordLogShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsRecordLogShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsRecordLogShowOutput is a type for action output parameters
type ActionDnsRecordLogShowOutput struct {
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
type ActionDnsRecordLogShowResponse struct {
	Action *ActionDnsRecordLogShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsRecordLog *ActionDnsRecordLogShowOutput `json:"dns_record_log"`
	}

	// Action output without the namespace
	Output *ActionDnsRecordLogShowOutput
}

// Prepare the action for invocation
func (action *ActionDnsRecordLogShow) Prepare() *ActionDnsRecordLogShowInvocation {
	return &ActionDnsRecordLogShowInvocation{
		Action: action,
		Path:   "/v7.0/dns_record_logs/{dns_record_log_id}",
	}
}

// ActionDnsRecordLogShowInvocation is used to configure action for invocation
type ActionDnsRecordLogShowInvocation struct {
	// Pointer to the action
	Action *ActionDnsRecordLogShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDnsRecordLogShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDnsRecordLogShowInvocation) SetPathParamInt(param string, value int64) *ActionDnsRecordLogShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDnsRecordLogShowInvocation) SetPathParamString(param string, value string) *ActionDnsRecordLogShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsRecordLogShowInvocation) NewMetaInput() *ActionDnsRecordLogShowMetaGlobalInput {
	inv.MetaInput = &ActionDnsRecordLogShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsRecordLogShowInvocation) SetMetaInput(input *ActionDnsRecordLogShowMetaGlobalInput) *ActionDnsRecordLogShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsRecordLogShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsRecordLogShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsRecordLogShowInvocation) Call() (*ActionDnsRecordLogShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDnsRecordLogShowInvocation) callAsQuery() (*ActionDnsRecordLogShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDnsRecordLogShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsRecordLog
	}
	return resp, err
}

func (inv *ActionDnsRecordLogShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
