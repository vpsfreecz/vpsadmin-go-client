package client

import (
	"strings"
)

// ActionDnsServerDelete is a type for action Dns_server#Delete
type ActionDnsServerDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsServerDelete(client *Client) *ActionDnsServerDelete {
	return &ActionDnsServerDelete{
		Client: client,
	}
}

// ActionDnsServerDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionDnsServerDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsServerDeleteMetaGlobalInput) SetIncludes(value string) *ActionDnsServerDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsServerDeleteMetaGlobalInput) SetNo(value bool) *ActionDnsServerDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsServerDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsServerDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionDnsServerDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsServerDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsServerDeleteRequest is a type for the entire action request
type ActionDnsServerDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionDnsServerDeleteResponse struct {
	Action *ActionDnsServerDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionDnsServerDelete) Prepare() *ActionDnsServerDeleteInvocation {
	return &ActionDnsServerDeleteInvocation{
		Action: action,
		Path:   "/v7.0/dns_servers/{dns_server_id}",
	}
}

// ActionDnsServerDeleteInvocation is used to configure action for invocation
type ActionDnsServerDeleteInvocation struct {
	// Pointer to the action
	Action *ActionDnsServerDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDnsServerDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDnsServerDeleteInvocation) SetPathParamInt(param string, value int64) *ActionDnsServerDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDnsServerDeleteInvocation) SetPathParamString(param string, value string) *ActionDnsServerDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsServerDeleteInvocation) NewMetaInput() *ActionDnsServerDeleteMetaGlobalInput {
	inv.MetaInput = &ActionDnsServerDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsServerDeleteInvocation) SetMetaInput(input *ActionDnsServerDeleteMetaGlobalInput) *ActionDnsServerDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsServerDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsServerDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsServerDeleteInvocation) Call() (*ActionDnsServerDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDnsServerDeleteInvocation) callAsBody() (*ActionDnsServerDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDnsServerDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionDnsServerDeleteInvocation) makeAllInputParams() *ActionDnsServerDeleteRequest {
	return &ActionDnsServerDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionDnsServerDeleteInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
