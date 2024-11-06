package client

import (
	"strings"
)

// ActionDnsTsigKeyDelete is a type for action Dns_tsig_key#Delete
type ActionDnsTsigKeyDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsTsigKeyDelete(client *Client) *ActionDnsTsigKeyDelete {
	return &ActionDnsTsigKeyDelete{
		Client: client,
	}
}

// ActionDnsTsigKeyDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionDnsTsigKeyDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsTsigKeyDeleteMetaGlobalInput) SetIncludes(value string) *ActionDnsTsigKeyDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsTsigKeyDeleteMetaGlobalInput) SetNo(value bool) *ActionDnsTsigKeyDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsTsigKeyDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsTsigKeyDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionDnsTsigKeyDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsTsigKeyDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsTsigKeyDeleteRequest is a type for the entire action request
type ActionDnsTsigKeyDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionDnsTsigKeyDeleteResponse struct {
	Action *ActionDnsTsigKeyDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionDnsTsigKeyDelete) Prepare() *ActionDnsTsigKeyDeleteInvocation {
	return &ActionDnsTsigKeyDeleteInvocation{
		Action: action,
		Path:   "/v7.0/dns_tsig_keys/{dns_tsig_key_id}",
	}
}

// ActionDnsTsigKeyDeleteInvocation is used to configure action for invocation
type ActionDnsTsigKeyDeleteInvocation struct {
	// Pointer to the action
	Action *ActionDnsTsigKeyDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDnsTsigKeyDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDnsTsigKeyDeleteInvocation) SetPathParamInt(param string, value int64) *ActionDnsTsigKeyDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDnsTsigKeyDeleteInvocation) SetPathParamString(param string, value string) *ActionDnsTsigKeyDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsTsigKeyDeleteInvocation) NewMetaInput() *ActionDnsTsigKeyDeleteMetaGlobalInput {
	inv.MetaInput = &ActionDnsTsigKeyDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsTsigKeyDeleteInvocation) SetMetaInput(input *ActionDnsTsigKeyDeleteMetaGlobalInput) *ActionDnsTsigKeyDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsTsigKeyDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsTsigKeyDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsTsigKeyDeleteInvocation) Call() (*ActionDnsTsigKeyDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDnsTsigKeyDeleteInvocation) callAsBody() (*ActionDnsTsigKeyDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDnsTsigKeyDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionDnsTsigKeyDeleteInvocation) makeAllInputParams() *ActionDnsTsigKeyDeleteRequest {
	return &ActionDnsTsigKeyDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionDnsTsigKeyDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
