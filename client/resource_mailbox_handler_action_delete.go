package client

import (
	"strings"
)

// ActionMailboxHandlerDelete is a type for action Mailbox.Handler#Delete
type ActionMailboxHandlerDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionMailboxHandlerDelete(client *Client) *ActionMailboxHandlerDelete {
	return &ActionMailboxHandlerDelete{
		Client: client,
	}
}

// ActionMailboxHandlerDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionMailboxHandlerDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailboxHandlerDeleteMetaGlobalInput) SetIncludes(value string) *ActionMailboxHandlerDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailboxHandlerDeleteMetaGlobalInput) SetNo(value bool) *ActionMailboxHandlerDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailboxHandlerDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailboxHandlerDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionMailboxHandlerDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailboxHandlerDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailboxHandlerDeleteRequest is a type for the entire action request
type ActionMailboxHandlerDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionMailboxHandlerDeleteResponse struct {
	Action *ActionMailboxHandlerDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionMailboxHandlerDelete) Prepare() *ActionMailboxHandlerDeleteInvocation {
	return &ActionMailboxHandlerDeleteInvocation{
		Action: action,
		Path:   "/v7.0/mailboxes/{mailbox_id}/handler/{handler_id}",
	}
}

// ActionMailboxHandlerDeleteInvocation is used to configure action for invocation
type ActionMailboxHandlerDeleteInvocation struct {
	// Pointer to the action
	Action *ActionMailboxHandlerDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMailboxHandlerDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailboxHandlerDeleteInvocation) SetPathParamInt(param string, value int64) *ActionMailboxHandlerDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailboxHandlerDeleteInvocation) SetPathParamString(param string, value string) *ActionMailboxHandlerDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailboxHandlerDeleteInvocation) NewMetaInput() *ActionMailboxHandlerDeleteMetaGlobalInput {
	inv.MetaInput = &ActionMailboxHandlerDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailboxHandlerDeleteInvocation) SetMetaInput(input *ActionMailboxHandlerDeleteMetaGlobalInput) *ActionMailboxHandlerDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailboxHandlerDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMailboxHandlerDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailboxHandlerDeleteInvocation) Call() (*ActionMailboxHandlerDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionMailboxHandlerDeleteInvocation) callAsBody() (*ActionMailboxHandlerDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMailboxHandlerDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionMailboxHandlerDeleteInvocation) makeAllInputParams() *ActionMailboxHandlerDeleteRequest {
	return &ActionMailboxHandlerDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionMailboxHandlerDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
