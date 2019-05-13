package client

import (
	"strings"
)

// ActionMailRecipientUpdate is a type for action Mail_recipient#Update
type ActionMailRecipientUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionMailRecipientUpdate(client *Client) *ActionMailRecipientUpdate {
	return &ActionMailRecipientUpdate{
		Client: client,
	}
}

// ActionMailRecipientUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionMailRecipientUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailRecipientUpdateMetaGlobalInput) SetNo(value bool) *ActionMailRecipientUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailRecipientUpdateMetaGlobalInput) SetIncludes(value string) *ActionMailRecipientUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailRecipientUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailRecipientUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionMailRecipientUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailRecipientUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailRecipientUpdateInput is a type for action input parameters
type ActionMailRecipientUpdateInput struct {
	Label string `json:"label"`
	To string `json:"to"`
	Cc string `json:"cc"`
	Bcc string `json:"bcc"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionMailRecipientUpdateInput) SetLabel(value string) *ActionMailRecipientUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}
// SetTo sets parameter To to value and selects it for sending
func (in *ActionMailRecipientUpdateInput) SetTo(value string) *ActionMailRecipientUpdateInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}
// SetCc sets parameter Cc to value and selects it for sending
func (in *ActionMailRecipientUpdateInput) SetCc(value string) *ActionMailRecipientUpdateInput {
	in.Cc = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Cc"] = nil
	return in
}
// SetBcc sets parameter Bcc to value and selects it for sending
func (in *ActionMailRecipientUpdateInput) SetBcc(value string) *ActionMailRecipientUpdateInput {
	in.Bcc = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Bcc"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailRecipientUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailRecipientUpdateInput) SelectParameters(params ...string) *ActionMailRecipientUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailRecipientUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailRecipientUpdateRequest is a type for the entire action request
type ActionMailRecipientUpdateRequest struct {
	MailRecipient map[string]interface{} `json:"mail_recipient"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionMailRecipientUpdateOutput is a type for action output parameters
type ActionMailRecipientUpdateOutput struct {
	Id int64 `json:"id"`
	Label string `json:"label"`
	To string `json:"to"`
	Cc string `json:"cc"`
	Bcc string `json:"bcc"`
}


// Type for action response, including envelope
type ActionMailRecipientUpdateResponse struct {
	Action *ActionMailRecipientUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MailRecipient *ActionMailRecipientUpdateOutput `json:"mail_recipient"`
	}

	// Action output without the namespace
	Output *ActionMailRecipientUpdateOutput
}


// Prepare the action for invocation
func (action *ActionMailRecipientUpdate) Prepare() *ActionMailRecipientUpdateInvocation {
	return &ActionMailRecipientUpdateInvocation{
		Action: action,
		Path: "/v5.0/mail_recipients/{mail_recipient_id}",
	}
}

// ActionMailRecipientUpdateInvocation is used to configure action for invocation
type ActionMailRecipientUpdateInvocation struct {
	// Pointer to the action
	Action *ActionMailRecipientUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMailRecipientUpdateInput
	// Global meta input parameters
	MetaInput *ActionMailRecipientUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailRecipientUpdateInvocation) SetPathParamInt(param string, value int64) *ActionMailRecipientUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailRecipientUpdateInvocation) SetPathParamString(param string, value string) *ActionMailRecipientUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMailRecipientUpdateInvocation) NewInput() *ActionMailRecipientUpdateInput {
	inv.Input = &ActionMailRecipientUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMailRecipientUpdateInvocation) SetInput(input *ActionMailRecipientUpdateInput) *ActionMailRecipientUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMailRecipientUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailRecipientUpdateInvocation) NewMetaInput() *ActionMailRecipientUpdateMetaGlobalInput {
	inv.MetaInput = &ActionMailRecipientUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailRecipientUpdateInvocation) SetMetaInput(input *ActionMailRecipientUpdateMetaGlobalInput) *ActionMailRecipientUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailRecipientUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailRecipientUpdateInvocation) Call() (*ActionMailRecipientUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionMailRecipientUpdateInvocation) callAsBody() (*ActionMailRecipientUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMailRecipientUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MailRecipient
	}
	return resp, err
}




func (inv *ActionMailRecipientUpdateInvocation) makeAllInputParams() *ActionMailRecipientUpdateRequest {
	return &ActionMailRecipientUpdateRequest{
		MailRecipient: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionMailRecipientUpdateInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionMailRecipientUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
