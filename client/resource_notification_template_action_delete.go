package client

import (
	"net/url"
	"strings"
)

// ActionNotificationTemplateDelete is a type for action Notification_template#Delete
type ActionNotificationTemplateDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationTemplateDelete(client *Client) *ActionNotificationTemplateDelete {
	return &ActionNotificationTemplateDelete{
		Client: client,
	}
}

// ActionNotificationTemplateDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationTemplateDeleteMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationTemplateDeleteMetaGlobalInput) SetIncludes(value string) *ActionNotificationTemplateDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationTemplateDeleteMetaGlobalInput) SetNo(value bool) *ActionNotificationTemplateDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTemplateDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTemplateDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationTemplateDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationTemplateDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTemplateDeleteRequest is a type for the entire action request
type ActionNotificationTemplateDeleteRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// Type for action response, including envelope
type ActionNotificationTemplateDeleteResponse struct {
	Action *ActionNotificationTemplateDelete "json:\"-\""
	*Envelope
}

// Prepare the action for invocation
func (action *ActionNotificationTemplateDelete) Prepare() *ActionNotificationTemplateDeleteInvocation {
	return &ActionNotificationTemplateDeleteInvocation{
		Action: action,
		Path:   "/v7.0/notification_templates/{notification_template_id}",
	}
}

// ActionNotificationTemplateDeleteInvocation is used to configure action for invocation
type ActionNotificationTemplateDeleteInvocation struct {
	// Pointer to the action
	Action *ActionNotificationTemplateDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNotificationTemplateDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationTemplateDeleteInvocation) SetPathParamInt(param string, value int64) *ActionNotificationTemplateDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationTemplateDeleteInvocation) SetPathParamString(param string, value string) *ActionNotificationTemplateDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationTemplateDeleteInvocation) NewMetaInput() *ActionNotificationTemplateDeleteMetaGlobalInput {
	inv.MetaInput = &ActionNotificationTemplateDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationTemplateDeleteInvocation) SetMetaInput(input *ActionNotificationTemplateDeleteMetaGlobalInput) *ActionNotificationTemplateDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationTemplateDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationTemplateDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationTemplateDeleteInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNotificationTemplateDeleteInvocation) Call() (*ActionNotificationTemplateDeleteResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNotificationTemplateDeleteInvocation) callAsBody() (*ActionNotificationTemplateDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNotificationTemplateDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionNotificationTemplateDeleteInvocation) makeAllInputParams() *ActionNotificationTemplateDeleteRequest {
	return &ActionNotificationTemplateDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionNotificationTemplateDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
