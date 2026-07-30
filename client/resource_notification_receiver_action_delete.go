package client

import (
	"net/url"
	"strings"
)

// ActionNotificationReceiverDelete is a type for action Notification_receiver#Delete
type ActionNotificationReceiverDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationReceiverDelete(client *Client) *ActionNotificationReceiverDelete {
	return &ActionNotificationReceiverDelete{
		Client: client,
	}
}

// ActionNotificationReceiverDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationReceiverDeleteMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationReceiverDeleteMetaGlobalInput) SetIncludes(value string) *ActionNotificationReceiverDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationReceiverDeleteMetaGlobalInput) SetNo(value bool) *ActionNotificationReceiverDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationReceiverDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationReceiverDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationReceiverDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationReceiverDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationReceiverDeleteRequest is a type for the entire action request
type ActionNotificationReceiverDeleteRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// Type for action response, including envelope
type ActionNotificationReceiverDeleteResponse struct {
	Action *ActionNotificationReceiverDelete "json:\"-\""
	*Envelope
}

// Prepare the action for invocation
func (action *ActionNotificationReceiverDelete) Prepare() *ActionNotificationReceiverDeleteInvocation {
	return &ActionNotificationReceiverDeleteInvocation{
		Action: action,
		Path:   "/v7.0/notification_receivers/{notification_receiver_id}",
	}
}

// ActionNotificationReceiverDeleteInvocation is used to configure action for invocation
type ActionNotificationReceiverDeleteInvocation struct {
	// Pointer to the action
	Action *ActionNotificationReceiverDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNotificationReceiverDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationReceiverDeleteInvocation) SetPathParamInt(param string, value int64) *ActionNotificationReceiverDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationReceiverDeleteInvocation) SetPathParamString(param string, value string) *ActionNotificationReceiverDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationReceiverDeleteInvocation) NewMetaInput() *ActionNotificationReceiverDeleteMetaGlobalInput {
	inv.MetaInput = &ActionNotificationReceiverDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationReceiverDeleteInvocation) SetMetaInput(input *ActionNotificationReceiverDeleteMetaGlobalInput) *ActionNotificationReceiverDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationReceiverDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationReceiverDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationReceiverDeleteInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNotificationReceiverDeleteInvocation) Call() (*ActionNotificationReceiverDeleteResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNotificationReceiverDeleteInvocation) callAsBody() (*ActionNotificationReceiverDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNotificationReceiverDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionNotificationReceiverDeleteInvocation) makeAllInputParams() *ActionNotificationReceiverDeleteRequest {
	return &ActionNotificationReceiverDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionNotificationReceiverDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
