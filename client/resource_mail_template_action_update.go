package client

import (
	"strings"
)

// ActionMailTemplateUpdate is a type for action Mail_template#Update
type ActionMailTemplateUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionMailTemplateUpdate(client *Client) *ActionMailTemplateUpdate {
	return &ActionMailTemplateUpdate{
		Client: client,
	}
}

// ActionMailTemplateUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionMailTemplateUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailTemplateUpdateMetaGlobalInput) SetNo(value bool) *ActionMailTemplateUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailTemplateUpdateMetaGlobalInput) SetIncludes(value string) *ActionMailTemplateUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionMailTemplateUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailTemplateUpdateInput is a type for action input parameters
type ActionMailTemplateUpdateInput struct {
	Name string `json:"name"`
	Label string `json:"label"`
	TemplateId string `json:"template_id"`
	UserVisibility string `json:"user_visibility"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionMailTemplateUpdateInput) SetName(value string) *ActionMailTemplateUpdateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}
// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionMailTemplateUpdateInput) SetLabel(value string) *ActionMailTemplateUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}
// SetTemplateId sets parameter TemplateId to value and selects it for sending
func (in *ActionMailTemplateUpdateInput) SetTemplateId(value string) *ActionMailTemplateUpdateInput {
	in.TemplateId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TemplateId"] = nil
	return in
}
// SetUserVisibility sets parameter UserVisibility to value and selects it for sending
func (in *ActionMailTemplateUpdateInput) SetUserVisibility(value string) *ActionMailTemplateUpdateInput {
	in.UserVisibility = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserVisibility"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateUpdateInput) SelectParameters(params ...string) *ActionMailTemplateUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailTemplateUpdateRequest is a type for the entire action request
type ActionMailTemplateUpdateRequest struct {
	MailTemplate map[string]interface{} `json:"mail_template"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionMailTemplateUpdateOutput is a type for action output parameters
type ActionMailTemplateUpdateOutput struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Label string `json:"label"`
	TemplateId string `json:"template_id"`
	UserVisibility string `json:"user_visibility"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


// Type for action response, including envelope
type ActionMailTemplateUpdateResponse struct {
	Action *ActionMailTemplateUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MailTemplate *ActionMailTemplateUpdateOutput `json:"mail_template"`
	}

	// Action output without the namespace
	Output *ActionMailTemplateUpdateOutput
}


// Prepare the action for invocation
func (action *ActionMailTemplateUpdate) Prepare() *ActionMailTemplateUpdateInvocation {
	return &ActionMailTemplateUpdateInvocation{
		Action: action,
		Path: "/v5.0/mail_templates/:mail_template_id",
	}
}

// ActionMailTemplateUpdateInvocation is used to configure action for invocation
type ActionMailTemplateUpdateInvocation struct {
	// Pointer to the action
	Action *ActionMailTemplateUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMailTemplateUpdateInput
	// Global meta input parameters
	MetaInput *ActionMailTemplateUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailTemplateUpdateInvocation) SetPathParamInt(param string, value int64) *ActionMailTemplateUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailTemplateUpdateInvocation) SetPathParamString(param string, value string) *ActionMailTemplateUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMailTemplateUpdateInvocation) NewInput() *ActionMailTemplateUpdateInput {
	inv.Input = &ActionMailTemplateUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMailTemplateUpdateInvocation) SetInput(input *ActionMailTemplateUpdateInput) *ActionMailTemplateUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMailTemplateUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailTemplateUpdateInvocation) NewMetaInput() *ActionMailTemplateUpdateMetaGlobalInput {
	inv.MetaInput = &ActionMailTemplateUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailTemplateUpdateInvocation) SetMetaInput(input *ActionMailTemplateUpdateMetaGlobalInput) *ActionMailTemplateUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailTemplateUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailTemplateUpdateInvocation) Call() (*ActionMailTemplateUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionMailTemplateUpdateInvocation) callAsBody() (*ActionMailTemplateUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMailTemplateUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MailTemplate
	}
	return resp, err
}




func (inv *ActionMailTemplateUpdateInvocation) makeAllInputParams() *ActionMailTemplateUpdateRequest {
	return &ActionMailTemplateUpdateRequest{
		MailTemplate: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionMailTemplateUpdateInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionMailTemplateUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
