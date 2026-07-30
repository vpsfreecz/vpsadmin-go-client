package client

import (
	"net/url"
	"strings"
)

// ActionNotificationTemplateVariantDelete is a type for action Notification_template.Variant#Delete
type ActionNotificationTemplateVariantDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationTemplateVariantDelete(client *Client) *ActionNotificationTemplateVariantDelete {
	return &ActionNotificationTemplateVariantDelete{
		Client: client,
	}
}

// ActionNotificationTemplateVariantDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationTemplateVariantDeleteMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationTemplateVariantDeleteMetaGlobalInput) SetIncludes(value string) *ActionNotificationTemplateVariantDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationTemplateVariantDeleteMetaGlobalInput) SetNo(value bool) *ActionNotificationTemplateVariantDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTemplateVariantDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTemplateVariantDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationTemplateVariantDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationTemplateVariantDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTemplateVariantDeleteRequest is a type for the entire action request
type ActionNotificationTemplateVariantDeleteRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// Type for action response, including envelope
type ActionNotificationTemplateVariantDeleteResponse struct {
	Action *ActionNotificationTemplateVariantDelete "json:\"-\""
	*Envelope
}

// Prepare the action for invocation
func (action *ActionNotificationTemplateVariantDelete) Prepare() *ActionNotificationTemplateVariantDeleteInvocation {
	return &ActionNotificationTemplateVariantDeleteInvocation{
		Action: action,
		Path:   "/v7.0/notification_templates/{notification_template_id}/variants/{variant_id}",
	}
}

// ActionNotificationTemplateVariantDeleteInvocation is used to configure action for invocation
type ActionNotificationTemplateVariantDeleteInvocation struct {
	// Pointer to the action
	Action *ActionNotificationTemplateVariantDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNotificationTemplateVariantDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationTemplateVariantDeleteInvocation) SetPathParamInt(param string, value int64) *ActionNotificationTemplateVariantDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationTemplateVariantDeleteInvocation) SetPathParamString(param string, value string) *ActionNotificationTemplateVariantDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationTemplateVariantDeleteInvocation) NewMetaInput() *ActionNotificationTemplateVariantDeleteMetaGlobalInput {
	inv.MetaInput = &ActionNotificationTemplateVariantDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationTemplateVariantDeleteInvocation) SetMetaInput(input *ActionNotificationTemplateVariantDeleteMetaGlobalInput) *ActionNotificationTemplateVariantDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationTemplateVariantDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationTemplateVariantDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationTemplateVariantDeleteInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNotificationTemplateVariantDeleteInvocation) Call() (*ActionNotificationTemplateVariantDeleteResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNotificationTemplateVariantDeleteInvocation) callAsBody() (*ActionNotificationTemplateVariantDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNotificationTemplateVariantDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionNotificationTemplateVariantDeleteInvocation) makeAllInputParams() *ActionNotificationTemplateVariantDeleteRequest {
	return &ActionNotificationTemplateVariantDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionNotificationTemplateVariantDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
