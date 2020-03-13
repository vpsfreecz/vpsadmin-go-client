package client

import (
)

// ActionMailRecipientCreate is a type for action Mail_recipient#Create
type ActionMailRecipientCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionMailRecipientCreate(client *Client) *ActionMailRecipientCreate {
	return &ActionMailRecipientCreate{
		Client: client,
	}
}

// ActionMailRecipientCreateMetaGlobalInput is a type for action global meta input parameters
type ActionMailRecipientCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailRecipientCreateMetaGlobalInput) SetNo(value bool) *ActionMailRecipientCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailRecipientCreateMetaGlobalInput) SetIncludes(value string) *ActionMailRecipientCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailRecipientCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailRecipientCreateMetaGlobalInput) SelectParameters(params ...string) *ActionMailRecipientCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailRecipientCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailRecipientCreateInput is a type for action input parameters
type ActionMailRecipientCreateInput struct {
	Label string `json:"label"`
	To string `json:"to"`
	Cc string `json:"cc"`
	Bcc string `json:"bcc"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionMailRecipientCreateInput) SetLabel(value string) *ActionMailRecipientCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}
// SetTo sets parameter To to value and selects it for sending
func (in *ActionMailRecipientCreateInput) SetTo(value string) *ActionMailRecipientCreateInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}
// SetCc sets parameter Cc to value and selects it for sending
func (in *ActionMailRecipientCreateInput) SetCc(value string) *ActionMailRecipientCreateInput {
	in.Cc = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Cc"] = nil
	return in
}
// SetBcc sets parameter Bcc to value and selects it for sending
func (in *ActionMailRecipientCreateInput) SetBcc(value string) *ActionMailRecipientCreateInput {
	in.Bcc = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Bcc"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailRecipientCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailRecipientCreateInput) SelectParameters(params ...string) *ActionMailRecipientCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailRecipientCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailRecipientCreateRequest is a type for the entire action request
type ActionMailRecipientCreateRequest struct {
	MailRecipient map[string]interface{} `json:"mail_recipient"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionMailRecipientCreateOutput is a type for action output parameters
type ActionMailRecipientCreateOutput struct {
	Id int64 `json:"id"`
	Label string `json:"label"`
	To string `json:"to"`
	Cc string `json:"cc"`
	Bcc string `json:"bcc"`
}


// Type for action response, including envelope
type ActionMailRecipientCreateResponse struct {
	Action *ActionMailRecipientCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MailRecipient *ActionMailRecipientCreateOutput `json:"mail_recipient"`
	}

	// Action output without the namespace
	Output *ActionMailRecipientCreateOutput
}


// Prepare the action for invocation
func (action *ActionMailRecipientCreate) Prepare() *ActionMailRecipientCreateInvocation {
	return &ActionMailRecipientCreateInvocation{
		Action: action,
		Path: "/v6.0/mail_recipients",
	}
}

// ActionMailRecipientCreateInvocation is used to configure action for invocation
type ActionMailRecipientCreateInvocation struct {
	// Pointer to the action
	Action *ActionMailRecipientCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMailRecipientCreateInput
	// Global meta input parameters
	MetaInput *ActionMailRecipientCreateMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMailRecipientCreateInvocation) NewInput() *ActionMailRecipientCreateInput {
	inv.Input = &ActionMailRecipientCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMailRecipientCreateInvocation) SetInput(input *ActionMailRecipientCreateInput) *ActionMailRecipientCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMailRecipientCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailRecipientCreateInvocation) NewMetaInput() *ActionMailRecipientCreateMetaGlobalInput {
	inv.MetaInput = &ActionMailRecipientCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailRecipientCreateInvocation) SetMetaInput(input *ActionMailRecipientCreateMetaGlobalInput) *ActionMailRecipientCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailRecipientCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailRecipientCreateInvocation) Call() (*ActionMailRecipientCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionMailRecipientCreateInvocation) callAsBody() (*ActionMailRecipientCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMailRecipientCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MailRecipient
	}
	return resp, err
}




func (inv *ActionMailRecipientCreateInvocation) makeAllInputParams() *ActionMailRecipientCreateRequest {
	return &ActionMailRecipientCreateRequest{
		MailRecipient: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionMailRecipientCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("To") {
			ret["to"] = inv.Input.To
		}
		if inv.IsParameterSelected("Cc") {
			ret["cc"] = inv.Input.Cc
		}
		if inv.IsParameterSelected("Bcc") {
			ret["bcc"] = inv.Input.Bcc
		}
	}

	return ret
}

func (inv *ActionMailRecipientCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
