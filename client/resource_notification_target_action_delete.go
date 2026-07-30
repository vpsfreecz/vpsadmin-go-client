package client

import (
	"net/url"
	"strings"
)

// ActionNotificationTargetDelete is a type for action Notification_target#Delete
type ActionNotificationTargetDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationTargetDelete(client *Client) *ActionNotificationTargetDelete {
	return &ActionNotificationTargetDelete{
		Client: client,
	}
}

// ActionNotificationTargetDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationTargetDeleteMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationTargetDeleteMetaGlobalInput) SetIncludes(value string) *ActionNotificationTargetDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationTargetDeleteMetaGlobalInput) SetNo(value bool) *ActionNotificationTargetDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTargetDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTargetDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationTargetDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationTargetDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTargetDeleteRequest is a type for the entire action request
type ActionNotificationTargetDeleteRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// Type for action response, including envelope
type ActionNotificationTargetDeleteResponse struct {
	Action *ActionNotificationTargetDelete "json:\"-\""
	*Envelope
}

// Prepare the action for invocation
func (action *ActionNotificationTargetDelete) Prepare() *ActionNotificationTargetDeleteInvocation {
	return &ActionNotificationTargetDeleteInvocation{
		Action: action,
		Path:   "/v7.0/notification_targets/{notification_target_id}",
	}
}

// ActionNotificationTargetDeleteInvocation is used to configure action for invocation
type ActionNotificationTargetDeleteInvocation struct {
	// Pointer to the action
	Action *ActionNotificationTargetDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNotificationTargetDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationTargetDeleteInvocation) SetPathParamInt(param string, value int64) *ActionNotificationTargetDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationTargetDeleteInvocation) SetPathParamString(param string, value string) *ActionNotificationTargetDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationTargetDeleteInvocation) NewMetaInput() *ActionNotificationTargetDeleteMetaGlobalInput {
	inv.MetaInput = &ActionNotificationTargetDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationTargetDeleteInvocation) SetMetaInput(input *ActionNotificationTargetDeleteMetaGlobalInput) *ActionNotificationTargetDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationTargetDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationTargetDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationTargetDeleteInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNotificationTargetDeleteInvocation) Call() (*ActionNotificationTargetDeleteResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNotificationTargetDeleteInvocation) callAsBody() (*ActionNotificationTargetDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNotificationTargetDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionNotificationTargetDeleteInvocation) makeAllInputParams() *ActionNotificationTargetDeleteRequest {
	return &ActionNotificationTargetDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionNotificationTargetDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
