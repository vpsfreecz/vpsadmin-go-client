package client

import (
	"net/url"
	"strings"
)

// ActionSecurityAdvisoryUpdateDelete is a type for action Security_advisory_update#Delete
type ActionSecurityAdvisoryUpdateDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryUpdateDelete(client *Client) *ActionSecurityAdvisoryUpdateDelete {
	return &ActionSecurityAdvisoryUpdateDelete{
		Client: client,
	}
}

// ActionSecurityAdvisoryUpdateDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryUpdateDeleteMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateDeleteMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryUpdateDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateDeleteMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryUpdateDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryUpdateDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryUpdateDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryUpdateDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryUpdateDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryUpdateDeleteRequest is a type for the entire action request
type ActionSecurityAdvisoryUpdateDeleteRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// Type for action response, including envelope
type ActionSecurityAdvisoryUpdateDeleteResponse struct {
	Action *ActionSecurityAdvisoryUpdateDelete "json:\"-\""
	*Envelope
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryUpdateDelete) Prepare() *ActionSecurityAdvisoryUpdateDeleteInvocation {
	return &ActionSecurityAdvisoryUpdateDeleteInvocation{
		Action: action,
		Path:   "/v7.0/security_advisory_updates/{security_advisory_update_id}",
	}
}

// ActionSecurityAdvisoryUpdateDeleteInvocation is used to configure action for invocation
type ActionSecurityAdvisoryUpdateDeleteInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryUpdateDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryUpdateDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSecurityAdvisoryUpdateDeleteInvocation) SetPathParamInt(param string, value int64) *ActionSecurityAdvisoryUpdateDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSecurityAdvisoryUpdateDeleteInvocation) SetPathParamString(param string, value string) *ActionSecurityAdvisoryUpdateDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryUpdateDeleteInvocation) NewMetaInput() *ActionSecurityAdvisoryUpdateDeleteMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryUpdateDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryUpdateDeleteInvocation) SetMetaInput(input *ActionSecurityAdvisoryUpdateDeleteMetaGlobalInput) *ActionSecurityAdvisoryUpdateDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryUpdateDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryUpdateDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryUpdateDeleteInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSecurityAdvisoryUpdateDeleteInvocation) Call() (*ActionSecurityAdvisoryUpdateDeleteResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionSecurityAdvisoryUpdateDeleteInvocation) callAsBody() (*ActionSecurityAdvisoryUpdateDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSecurityAdvisoryUpdateDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionSecurityAdvisoryUpdateDeleteInvocation) makeAllInputParams() *ActionSecurityAdvisoryUpdateDeleteRequest {
	return &ActionSecurityAdvisoryUpdateDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionSecurityAdvisoryUpdateDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
