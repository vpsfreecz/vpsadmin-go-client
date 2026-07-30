package client

import (
	"net/url"
	"strings"
)

// ActionNotificationReceiverTargetDelete is a type for action Notification_receiver.Target#Delete
type ActionNotificationReceiverTargetDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationReceiverTargetDelete(client *Client) *ActionNotificationReceiverTargetDelete {
	return &ActionNotificationReceiverTargetDelete{
		Client: client,
	}
}

// ActionNotificationReceiverTargetDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationReceiverTargetDeleteMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationReceiverTargetDeleteMetaGlobalInput) SetIncludes(value string) *ActionNotificationReceiverTargetDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationReceiverTargetDeleteMetaGlobalInput) SetNo(value bool) *ActionNotificationReceiverTargetDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationReceiverTargetDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationReceiverTargetDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationReceiverTargetDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationReceiverTargetDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationReceiverTargetDeleteRequest is a type for the entire action request
type ActionNotificationReceiverTargetDeleteRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// Type for action response, including envelope
type ActionNotificationReceiverTargetDeleteResponse struct {
	Action *ActionNotificationReceiverTargetDelete "json:\"-\""
	*Envelope
}

// Prepare the action for invocation
func (action *ActionNotificationReceiverTargetDelete) Prepare() *ActionNotificationReceiverTargetDeleteInvocation {
	return &ActionNotificationReceiverTargetDeleteInvocation{
		Action: action,
		Path:   "/v7.0/notification_receivers/{notification_receiver_id}/target/{target_id}",
	}
}

// ActionNotificationReceiverTargetDeleteInvocation is used to configure action for invocation
type ActionNotificationReceiverTargetDeleteInvocation struct {
	// Pointer to the action
	Action *ActionNotificationReceiverTargetDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNotificationReceiverTargetDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationReceiverTargetDeleteInvocation) SetPathParamInt(param string, value int64) *ActionNotificationReceiverTargetDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationReceiverTargetDeleteInvocation) SetPathParamString(param string, value string) *ActionNotificationReceiverTargetDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationReceiverTargetDeleteInvocation) NewMetaInput() *ActionNotificationReceiverTargetDeleteMetaGlobalInput {
	inv.MetaInput = &ActionNotificationReceiverTargetDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationReceiverTargetDeleteInvocation) SetMetaInput(input *ActionNotificationReceiverTargetDeleteMetaGlobalInput) *ActionNotificationReceiverTargetDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationReceiverTargetDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationReceiverTargetDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationReceiverTargetDeleteInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNotificationReceiverTargetDeleteInvocation) Call() (*ActionNotificationReceiverTargetDeleteResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNotificationReceiverTargetDeleteInvocation) callAsBody() (*ActionNotificationReceiverTargetDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNotificationReceiverTargetDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionNotificationReceiverTargetDeleteInvocation) makeAllInputParams() *ActionNotificationReceiverTargetDeleteRequest {
	return &ActionNotificationReceiverTargetDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionNotificationReceiverTargetDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
