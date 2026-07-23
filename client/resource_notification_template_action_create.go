package client

import ()

// ActionNotificationTemplateCreate is a type for action Notification_template#Create
type ActionNotificationTemplateCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationTemplateCreate(client *Client) *ActionNotificationTemplateCreate {
	return &ActionNotificationTemplateCreate{
		Client: client,
	}
}

// ActionNotificationTemplateCreateMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationTemplateCreateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationTemplateCreateMetaGlobalInput) SetIncludes(value string) *ActionNotificationTemplateCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationTemplateCreateMetaGlobalInput) SetNo(value bool) *ActionNotificationTemplateCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTemplateCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTemplateCreateMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationTemplateCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationTemplateCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTemplateCreateInput is a type for action input parameters
type ActionNotificationTemplateCreateInput struct {
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
func (in *ActionNotificationTemplateCreateInput) SetLabel(value string) *ActionNotificationTemplateCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionNotificationTemplateCreateInput) SetName(value string) *ActionNotificationTemplateCreateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetTemplateId sets parameter TemplateId to value and selects it for sending
func (in *ActionNotificationTemplateCreateInput) SetTemplateId(value string) *ActionNotificationTemplateCreateInput {
	in.TemplateId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TemplateId"] = nil
	return in
}

// SetUserVisibility sets parameter UserVisibility to value and selects it for sending
func (in *ActionNotificationTemplateCreateInput) SetUserVisibility(value string) *ActionNotificationTemplateCreateInput {
	in.UserVisibility = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserVisibility"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTemplateCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTemplateCreateInput) SelectParameters(params ...string) *ActionNotificationTemplateCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNotificationTemplateCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNotificationTemplateCreateInput) UnselectParameters(params ...string) *ActionNotificationTemplateCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNotificationTemplateCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTemplateCreateRequest is a type for the entire action request
type ActionNotificationTemplateCreateRequest struct {
	NotificationTemplate map[string]interface{} "json:\"notification_template\""
	Meta                 map[string]interface{} "json:\"_meta\""
}

// ActionNotificationTemplateCreateOutput is a type for action output parameters
type ActionNotificationTemplateCreateOutput struct {
	CreatedAt      string "json:\"created_at\""
	Id             int64  "json:\"id\""
	Label          string "json:\"label\""
	Name           string "json:\"name\""
	TemplateId     string "json:\"template_id\""
	UpdatedAt      string "json:\"updated_at\""
	UserVisibility string "json:\"user_visibility\""
}

// Type for action response, including envelope
type ActionNotificationTemplateCreateResponse struct {
	Action *ActionNotificationTemplateCreate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationTemplate *ActionNotificationTemplateCreateOutput "json:\"notification_template\""
	}

	// Action output without the namespace
	Output *ActionNotificationTemplateCreateOutput
}

// Prepare the action for invocation
func (action *ActionNotificationTemplateCreate) Prepare() *ActionNotificationTemplateCreateInvocation {
	return &ActionNotificationTemplateCreateInvocation{
		Action: action,
		Path:   "/v7.0/notification_templates",
	}
}

// ActionNotificationTemplateCreateInvocation is used to configure action for invocation
type ActionNotificationTemplateCreateInvocation struct {
	// Pointer to the action
	Action *ActionNotificationTemplateCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNotificationTemplateCreateInput
	// Global meta input parameters
	MetaInput *ActionNotificationTemplateCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNotificationTemplateCreateInvocation) NewInput() *ActionNotificationTemplateCreateInput {
	inv.Input = &ActionNotificationTemplateCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNotificationTemplateCreateInvocation) SetInput(input *ActionNotificationTemplateCreateInput) *ActionNotificationTemplateCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNotificationTemplateCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNotificationTemplateCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationTemplateCreateInvocation) NewMetaInput() *ActionNotificationTemplateCreateMetaGlobalInput {
	inv.MetaInput = &ActionNotificationTemplateCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationTemplateCreateInvocation) SetMetaInput(input *ActionNotificationTemplateCreateMetaGlobalInput) *ActionNotificationTemplateCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationTemplateCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationTemplateCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationTemplateCreateInvocation) validate() error {
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
func (inv *ActionNotificationTemplateCreateInvocation) Call() (*ActionNotificationTemplateCreateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNotificationTemplateCreateInvocation) callAsBody() (*ActionNotificationTemplateCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNotificationTemplateCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationTemplate
	}
	return resp, err
}

func (inv *ActionNotificationTemplateCreateInvocation) makeAllInputParams() *ActionNotificationTemplateCreateRequest {
	return &ActionNotificationTemplateCreateRequest{
		NotificationTemplate: inv.makeInputParams(),
		Meta:                 inv.makeMetaInputParams(),
	}
}

func (inv *ActionNotificationTemplateCreateInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionNotificationTemplateCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
