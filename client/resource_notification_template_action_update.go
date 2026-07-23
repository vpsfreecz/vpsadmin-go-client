package client

import (
	"net/url"
	"strings"
)

// ActionNotificationTemplateUpdate is a type for action Notification_template#Update
type ActionNotificationTemplateUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationTemplateUpdate(client *Client) *ActionNotificationTemplateUpdate {
	return &ActionNotificationTemplateUpdate{
		Client: client,
	}
}

// ActionNotificationTemplateUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationTemplateUpdateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationTemplateUpdateMetaGlobalInput) SetIncludes(value string) *ActionNotificationTemplateUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationTemplateUpdateMetaGlobalInput) SetNo(value bool) *ActionNotificationTemplateUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTemplateUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTemplateUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationTemplateUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationTemplateUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTemplateUpdateInput is a type for action input parameters
type ActionNotificationTemplateUpdateInput struct {
	Label          string "json:\"label\""
	Name           string "json:\"name\""
	TemplateId     string "json:\"template_id\""
	UserVisibility string "json:\"user_visibility\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionNotificationTemplateUpdateInput) SetLabel(value string) *ActionNotificationTemplateUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionNotificationTemplateUpdateInput) SetName(value string) *ActionNotificationTemplateUpdateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetTemplateId sets parameter TemplateId to value and selects it for sending
func (in *ActionNotificationTemplateUpdateInput) SetTemplateId(value string) *ActionNotificationTemplateUpdateInput {
	in.TemplateId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TemplateId"] = nil
	return in
}

// SetUserVisibility sets parameter UserVisibility to value and selects it for sending
func (in *ActionNotificationTemplateUpdateInput) SetUserVisibility(value string) *ActionNotificationTemplateUpdateInput {
	in.UserVisibility = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserVisibility"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTemplateUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTemplateUpdateInput) SelectParameters(params ...string) *ActionNotificationTemplateUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNotificationTemplateUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNotificationTemplateUpdateInput) UnselectParameters(params ...string) *ActionNotificationTemplateUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNotificationTemplateUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTemplateUpdateRequest is a type for the entire action request
type ActionNotificationTemplateUpdateRequest struct {
	NotificationTemplate map[string]interface{} "json:\"notification_template\""
	Meta                 map[string]interface{} "json:\"_meta\""
}

// ActionNotificationTemplateUpdateOutput is a type for action output parameters
type ActionNotificationTemplateUpdateOutput struct {
	CreatedAt      string "json:\"created_at\""
	Id             int64  "json:\"id\""
	Label          string "json:\"label\""
	Name           string "json:\"name\""
	TemplateId     string "json:\"template_id\""
	UpdatedAt      string "json:\"updated_at\""
	UserVisibility string "json:\"user_visibility\""
}

// Type for action response, including envelope
type ActionNotificationTemplateUpdateResponse struct {
	Action *ActionNotificationTemplateUpdate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationTemplate *ActionNotificationTemplateUpdateOutput "json:\"notification_template\""
	}

	// Action output without the namespace
	Output *ActionNotificationTemplateUpdateOutput
}

// Prepare the action for invocation
func (action *ActionNotificationTemplateUpdate) Prepare() *ActionNotificationTemplateUpdateInvocation {
	return &ActionNotificationTemplateUpdateInvocation{
		Action: action,
		Path:   "/v7.0/notification_templates/{notification_template_id}",
	}
}

// ActionNotificationTemplateUpdateInvocation is used to configure action for invocation
type ActionNotificationTemplateUpdateInvocation struct {
	// Pointer to the action
	Action *ActionNotificationTemplateUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNotificationTemplateUpdateInput
	// Global meta input parameters
	MetaInput *ActionNotificationTemplateUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationTemplateUpdateInvocation) SetPathParamInt(param string, value int64) *ActionNotificationTemplateUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationTemplateUpdateInvocation) SetPathParamString(param string, value string) *ActionNotificationTemplateUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNotificationTemplateUpdateInvocation) NewInput() *ActionNotificationTemplateUpdateInput {
	inv.Input = &ActionNotificationTemplateUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNotificationTemplateUpdateInvocation) SetInput(input *ActionNotificationTemplateUpdateInput) *ActionNotificationTemplateUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNotificationTemplateUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNotificationTemplateUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationTemplateUpdateInvocation) NewMetaInput() *ActionNotificationTemplateUpdateMetaGlobalInput {
	inv.MetaInput = &ActionNotificationTemplateUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationTemplateUpdateInvocation) SetMetaInput(input *ActionNotificationTemplateUpdateMetaGlobalInput) *ActionNotificationTemplateUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationTemplateUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationTemplateUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationTemplateUpdateInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNotificationTemplateUpdateInvocation) Call() (*ActionNotificationTemplateUpdateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNotificationTemplateUpdateInvocation) callAsBody() (*ActionNotificationTemplateUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNotificationTemplateUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationTemplate
	}
	return resp, err
}

func (inv *ActionNotificationTemplateUpdateInvocation) makeAllInputParams() *ActionNotificationTemplateUpdateRequest {
	return &ActionNotificationTemplateUpdateRequest{
		NotificationTemplate: inv.makeInputParams(),
		Meta:                 inv.makeMetaInputParams(),
	}
}

func (inv *ActionNotificationTemplateUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("Name") {
			ret["name"] = inv.Input.Name
		}
		if inv.IsParameterSelected("TemplateId") {
			ret["template_id"] = inv.Input.TemplateId
		}
		if inv.IsParameterSelected("UserVisibility") {
			ret["user_visibility"] = inv.Input.UserVisibility
		}
	}

	return ret
}

func (inv *ActionNotificationTemplateUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
