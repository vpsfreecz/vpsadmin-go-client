package client

import (
)

// ActionMailTemplateCreate is a type for action Mail_template#Create
type ActionMailTemplateCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionMailTemplateCreate(client *Client) *ActionMailTemplateCreate {
	return &ActionMailTemplateCreate{
		Client: client,
	}
}

// ActionMailTemplateCreateMetaGlobalInput is a type for action global meta input parameters
type ActionMailTemplateCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailTemplateCreateMetaGlobalInput) SetNo(value bool) *ActionMailTemplateCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailTemplateCreateMetaGlobalInput) SetIncludes(value string) *ActionMailTemplateCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateCreateMetaGlobalInput) SelectParameters(params ...string) *ActionMailTemplateCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailTemplateCreateInput is a type for action input parameters
type ActionMailTemplateCreateInput struct {
	Name string `json:"name"`
	Label string `json:"label"`
	TemplateId string `json:"template_id"`
	UserVisibility string `json:"user_visibility"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionMailTemplateCreateInput) SetName(value string) *ActionMailTemplateCreateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}
// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionMailTemplateCreateInput) SetLabel(value string) *ActionMailTemplateCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}
// SetTemplateId sets parameter TemplateId to value and selects it for sending
func (in *ActionMailTemplateCreateInput) SetTemplateId(value string) *ActionMailTemplateCreateInput {
	in.TemplateId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TemplateId"] = nil
	return in
}
// SetUserVisibility sets parameter UserVisibility to value and selects it for sending
func (in *ActionMailTemplateCreateInput) SetUserVisibility(value string) *ActionMailTemplateCreateInput {
	in.UserVisibility = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserVisibility"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateCreateInput) SelectParameters(params ...string) *ActionMailTemplateCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailTemplateCreateRequest is a type for the entire action request
type ActionMailTemplateCreateRequest struct {
	MailTemplate map[string]interface{} `json:"mail_template"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionMailTemplateCreateOutput is a type for action output parameters
type ActionMailTemplateCreateOutput struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Label string `json:"label"`
	TemplateId string `json:"template_id"`
	UserVisibility string `json:"user_visibility"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


// Type for action response, including envelope
type ActionMailTemplateCreateResponse struct {
	Action *ActionMailTemplateCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MailTemplate *ActionMailTemplateCreateOutput `json:"mail_template"`
	}

	// Action output without the namespace
	Output *ActionMailTemplateCreateOutput
}


// Prepare the action for invocation
func (action *ActionMailTemplateCreate) Prepare() *ActionMailTemplateCreateInvocation {
	return &ActionMailTemplateCreateInvocation{
		Action: action,
		Path: "/v6.0/mail_templates",
	}
}

// ActionMailTemplateCreateInvocation is used to configure action for invocation
type ActionMailTemplateCreateInvocation struct {
	// Pointer to the action
	Action *ActionMailTemplateCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMailTemplateCreateInput
	// Global meta input parameters
	MetaInput *ActionMailTemplateCreateMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMailTemplateCreateInvocation) NewInput() *ActionMailTemplateCreateInput {
	inv.Input = &ActionMailTemplateCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMailTemplateCreateInvocation) SetInput(input *ActionMailTemplateCreateInput) *ActionMailTemplateCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMailTemplateCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailTemplateCreateInvocation) NewMetaInput() *ActionMailTemplateCreateMetaGlobalInput {
	inv.MetaInput = &ActionMailTemplateCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailTemplateCreateInvocation) SetMetaInput(input *ActionMailTemplateCreateMetaGlobalInput) *ActionMailTemplateCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailTemplateCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailTemplateCreateInvocation) Call() (*ActionMailTemplateCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionMailTemplateCreateInvocation) callAsBody() (*ActionMailTemplateCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMailTemplateCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MailTemplate
	}
	return resp, err
}




func (inv *ActionMailTemplateCreateInvocation) makeAllInputParams() *ActionMailTemplateCreateRequest {
	return &ActionMailTemplateCreateRequest{
		MailTemplate: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionMailTemplateCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Name") {
			ret["name"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
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

func (inv *ActionMailTemplateCreateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
