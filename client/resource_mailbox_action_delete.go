package client

import (
	"strings"
)

// ActionMailboxDelete is a type for action Mailbox#Delete
type ActionMailboxDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionMailboxDelete(client *Client) *ActionMailboxDelete {
	return &ActionMailboxDelete{
		Client: client,
	}
}

// ActionMailboxDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionMailboxDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailboxDeleteMetaGlobalInput) SetIncludes(value string) *ActionMailboxDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailboxDeleteMetaGlobalInput) SetNo(value bool) *ActionMailboxDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailboxDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailboxDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionMailboxDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailboxDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailboxDeleteRequest is a type for the entire action request
type ActionMailboxDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionMailboxDeleteResponse struct {
	Action *ActionMailboxDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionMailboxDelete) Prepare() *ActionMailboxDeleteInvocation {
	return &ActionMailboxDeleteInvocation{
		Action: action,
		Path:   "/v7.0/mailboxes/{mailbox_id}",
	}
}

// ActionMailboxDeleteInvocation is used to configure action for invocation
type ActionMailboxDeleteInvocation struct {
	// Pointer to the action
	Action *ActionMailboxDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMailboxDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailboxDeleteInvocation) SetPathParamInt(param string, value int64) *ActionMailboxDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailboxDeleteInvocation) SetPathParamString(param string, value string) *ActionMailboxDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailboxDeleteInvocation) NewMetaInput() *ActionMailboxDeleteMetaGlobalInput {
	inv.MetaInput = &ActionMailboxDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailboxDeleteInvocation) SetMetaInput(input *ActionMailboxDeleteMetaGlobalInput) *ActionMailboxDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailboxDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMailboxDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailboxDeleteInvocation) Call() (*ActionMailboxDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionMailboxDeleteInvocation) callAsBody() (*ActionMailboxDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMailboxDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionMailboxDeleteInvocation) makeAllInputParams() *ActionMailboxDeleteRequest {
	return &ActionMailboxDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionMailboxDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
