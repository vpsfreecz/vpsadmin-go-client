package client

import (
	"net/url"
	"strings"
)

// ActionDnsServerZoneTransferLogShow is a type for action Dns_server_zone_transfer_log#Show
type ActionDnsServerZoneTransferLogShow struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsServerZoneTransferLogShow(client *Client) *ActionDnsServerZoneTransferLogShow {
	return &ActionDnsServerZoneTransferLogShow{
		Client: client,
	}
}

// ActionDnsServerZoneTransferLogShowMetaGlobalInput is a type for action global meta input parameters
type ActionDnsServerZoneTransferLogShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsServerZoneTransferLogShowMetaGlobalInput) SetIncludes(value string) *ActionDnsServerZoneTransferLogShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsServerZoneTransferLogShowMetaGlobalInput) SetNo(value bool) *ActionDnsServerZoneTransferLogShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsServerZoneTransferLogShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsServerZoneTransferLogShowMetaGlobalInput) SelectParameters(params ...string) *ActionDnsServerZoneTransferLogShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsServerZoneTransferLogShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsServerZoneTransferLogShowOutput is a type for action output parameters
type ActionDnsServerZoneTransferLogShowOutput struct {
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
type ActionDnsServerZoneTransferLogShowResponse struct {
	Action *ActionDnsServerZoneTransferLogShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsServerZoneTransferLog *ActionDnsServerZoneTransferLogShowOutput "json:\"dns_server_zone_transfer_log\""
	}

	// Action output without the namespace
	Output *ActionDnsServerZoneTransferLogShowOutput
}

// Prepare the action for invocation
func (action *ActionDnsServerZoneTransferLogShow) Prepare() *ActionDnsServerZoneTransferLogShowInvocation {
	return &ActionDnsServerZoneTransferLogShowInvocation{
		Action: action,
		Path:   "/v7.0/dns_server_zone_transfer_logs/{dns_server_zone_transfer_log_id}",
	}
}

// ActionDnsServerZoneTransferLogShowInvocation is used to configure action for invocation
type ActionDnsServerZoneTransferLogShowInvocation struct {
	// Pointer to the action
	Action *ActionDnsServerZoneTransferLogShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDnsServerZoneTransferLogShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDnsServerZoneTransferLogShowInvocation) SetPathParamInt(param string, value int64) *ActionDnsServerZoneTransferLogShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDnsServerZoneTransferLogShowInvocation) SetPathParamString(param string, value string) *ActionDnsServerZoneTransferLogShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsServerZoneTransferLogShowInvocation) NewMetaInput() *ActionDnsServerZoneTransferLogShowMetaGlobalInput {
	inv.MetaInput = &ActionDnsServerZoneTransferLogShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsServerZoneTransferLogShowInvocation) SetMetaInput(input *ActionDnsServerZoneTransferLogShowMetaGlobalInput) *ActionDnsServerZoneTransferLogShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsServerZoneTransferLogShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsServerZoneTransferLogShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionDnsServerZoneTransferLogShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsServerZoneTransferLogShowInvocation) Call() (*ActionDnsServerZoneTransferLogShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionDnsServerZoneTransferLogShowInvocation) callAsQuery() (*ActionDnsServerZoneTransferLogShowResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionDnsServerZoneTransferLogShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsServerZoneTransferLog
	}
	return resp, err
}

func (inv *ActionDnsServerZoneTransferLogShowInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
	if inv.MetaInput != nil {
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
