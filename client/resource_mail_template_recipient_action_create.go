package client

import (
	"strings"
)

// ActionMailTemplateRecipientCreate is a type for action Mail_template.Recipient#Create
type ActionMailTemplateRecipientCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionMailTemplateRecipientCreate(client *Client) *ActionMailTemplateRecipientCreate {
	return &ActionMailTemplateRecipientCreate{
		Client: client,
	}
}

// ActionMailTemplateRecipientCreateMetaGlobalInput is a type for action global meta input parameters
type ActionMailTemplateRecipientCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailTemplateRecipientCreateMetaGlobalInput) SetNo(value bool) *ActionMailTemplateRecipientCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailTemplateRecipientCreateMetaGlobalInput) SetIncludes(value string) *ActionMailTemplateRecipientCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateRecipientCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateRecipientCreateMetaGlobalInput) SelectParameters(params ...string) *ActionMailTemplateRecipientCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateRecipientCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailTemplateRecipientCreateInput is a type for action input parameters
type ActionMailTemplateRecipientCreateInput struct {
	MailRecipient int64 `json:"mail_recipient"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetMailRecipient sets parameter MailRecipient to value and selects it for sending
func (in *ActionMailTemplateRecipientCreateInput) SetMailRecipient(value int64) *ActionMailTemplateRecipientCreateInput {
	in.MailRecipient = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MailRecipient"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateRecipientCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateRecipientCreateInput) SelectParameters(params ...string) *ActionMailTemplateRecipientCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateRecipientCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailTemplateRecipientCreateRequest is a type for the entire action request
type ActionMailTemplateRecipientCreateRequest struct {
	Recipient map[string]interface{} `json:"recipient"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionMailTemplateRecipientCreateOutput is a type for action output parameters
type ActionMailTemplateRecipientCreateOutput struct {
	Id int64 `json:"id"`
	MailRecipient *ActionMailRecipientShowOutput `json:"mail_recipient"`
}


// Type for action response, including envelope
type ActionMailTemplateRecipientCreateResponse struct {
	Action *ActionMailTemplateRecipientCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Recipient *ActionMailTemplateRecipientCreateOutput `json:"recipient"`
	}

	// Action output without the namespace
	Output *ActionMailTemplateRecipientCreateOutput
}


// Prepare the action for invocation
func (action *ActionMailTemplateRecipientCreate) Prepare() *ActionMailTemplateRecipientCreateInvocation {
	return &ActionMailTemplateRecipientCreateInvocation{
		Action: action,
		Path: "/v5.0/mail_templates/:mail_template_id/recipients",
	}
}

// ActionMailTemplateRecipientCreateInvocation is used to configure action for invocation
type ActionMailTemplateRecipientCreateInvocation struct {
	// Pointer to the action
	Action *ActionMailTemplateRecipientCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMailTemplateRecipientCreateInput
	// Global meta input parameters
	MetaInput *ActionMailTemplateRecipientCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailTemplateRecipientCreateInvocation) SetPathParamInt(param string, value int64) *ActionMailTemplateRecipientCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailTemplateRecipientCreateInvocation) SetPathParamString(param string, value string) *ActionMailTemplateRecipientCreateInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMailTemplateRecipientCreateInvocation) NewInput() *ActionMailTemplateRecipientCreateInput {
	inv.Input = &ActionMailTemplateRecipientCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMailTemplateRecipientCreateInvocation) SetInput(input *ActionMailTemplateRecipientCreateInput) *ActionMailTemplateRecipientCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMailTemplateRecipientCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailTemplateRecipientCreateInvocation) NewMetaInput() *ActionMailTemplateRecipientCreateMetaGlobalInput {
	inv.MetaInput = &ActionMailTemplateRecipientCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailTemplateRecipientCreateInvocation) SetMetaInput(input *ActionMailTemplateRecipientCreateMetaGlobalInput) *ActionMailTemplateRecipientCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailTemplateRecipientCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailTemplateRecipientCreateInvocation) Call() (*ActionMailTemplateRecipientCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionMailTemplateRecipientCreateInvocation) callAsBody() (*ActionMailTemplateRecipientCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMailTemplateRecipientCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Recipient
	}
	return resp, err
}




func (inv *ActionMailTemplateRecipientCreateInvocation) makeAllInputParams() *ActionMailTemplateRecipientCreateRequest {
	return &ActionMailTemplateRecipientCreateRequest{
		Recipient: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionMailTemplateRecipientCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("MailRecipient") {
			ret["mail_recipient"] = inv.Input.MailRecipient
		}
	}

	return ret
}

func (inv *ActionMailTemplateRecipientCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
