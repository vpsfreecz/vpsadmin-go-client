package client

import (
	"strings"
)

// ActionUserMailTemplateRecipientUpdate is a type for action User.Mail_template_recipient#Update
type ActionUserMailTemplateRecipientUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserMailTemplateRecipientUpdate(client *Client) *ActionUserMailTemplateRecipientUpdate {
	return &ActionUserMailTemplateRecipientUpdate{
		Client: client,
	}
}

// ActionUserMailTemplateRecipientUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionUserMailTemplateRecipientUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserMailTemplateRecipientUpdateMetaGlobalInput) SetIncludes(value string) *ActionUserMailTemplateRecipientUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserMailTemplateRecipientUpdateMetaGlobalInput) SetNo(value bool) *ActionUserMailTemplateRecipientUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserMailTemplateRecipientUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserMailTemplateRecipientUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionUserMailTemplateRecipientUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserMailTemplateRecipientUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserMailTemplateRecipientUpdateInput is a type for action input parameters
type ActionUserMailTemplateRecipientUpdateInput struct {
	Enabled bool   `json:"enabled"`
	To      string `json:"to"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionUserMailTemplateRecipientUpdateInput) SetEnabled(value bool) *ActionUserMailTemplateRecipientUpdateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionUserMailTemplateRecipientUpdateInput) SetTo(value string) *ActionUserMailTemplateRecipientUpdateInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserMailTemplateRecipientUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserMailTemplateRecipientUpdateInput) SelectParameters(params ...string) *ActionUserMailTemplateRecipientUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserMailTemplateRecipientUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserMailTemplateRecipientUpdateInput) UnselectParameters(params ...string) *ActionUserMailTemplateRecipientUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserMailTemplateRecipientUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserMailTemplateRecipientUpdateRequest is a type for the entire action request
type ActionUserMailTemplateRecipientUpdateRequest struct {
	MailTemplateRecipient map[string]interface{} `json:"mail_template_recipient"`
	Meta                  map[string]interface{} `json:"_meta"`
}

// ActionUserMailTemplateRecipientUpdateOutput is a type for action output parameters
type ActionUserMailTemplateRecipientUpdateOutput struct {
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Id          string `json:"id"`
	Label       string `json:"label"`
	To          string `json:"to"`
}

// Type for action response, including envelope
type ActionUserMailTemplateRecipientUpdateResponse struct {
	Action *ActionUserMailTemplateRecipientUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MailTemplateRecipient *ActionUserMailTemplateRecipientUpdateOutput `json:"mail_template_recipient"`
	}

	// Action output without the namespace
	Output *ActionUserMailTemplateRecipientUpdateOutput
}

// Prepare the action for invocation
func (action *ActionUserMailTemplateRecipientUpdate) Prepare() *ActionUserMailTemplateRecipientUpdateInvocation {
	return &ActionUserMailTemplateRecipientUpdateInvocation{
		Action: action,
		Path:   "/v6.0/users/{user_id}/mail_template_recipients/{mail_template_recipient_id}",
	}
}

// ActionUserMailTemplateRecipientUpdateInvocation is used to configure action for invocation
type ActionUserMailTemplateRecipientUpdateInvocation struct {
	// Pointer to the action
	Action *ActionUserMailTemplateRecipientUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserMailTemplateRecipientUpdateInput
	// Global meta input parameters
	MetaInput *ActionUserMailTemplateRecipientUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserMailTemplateRecipientUpdateInvocation) SetPathParamInt(param string, value int64) *ActionUserMailTemplateRecipientUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserMailTemplateRecipientUpdateInvocation) SetPathParamString(param string, value string) *ActionUserMailTemplateRecipientUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserMailTemplateRecipientUpdateInvocation) NewInput() *ActionUserMailTemplateRecipientUpdateInput {
	inv.Input = &ActionUserMailTemplateRecipientUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserMailTemplateRecipientUpdateInvocation) SetInput(input *ActionUserMailTemplateRecipientUpdateInput) *ActionUserMailTemplateRecipientUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserMailTemplateRecipientUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserMailTemplateRecipientUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserMailTemplateRecipientUpdateInvocation) NewMetaInput() *ActionUserMailTemplateRecipientUpdateMetaGlobalInput {
	inv.MetaInput = &ActionUserMailTemplateRecipientUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserMailTemplateRecipientUpdateInvocation) SetMetaInput(input *ActionUserMailTemplateRecipientUpdateMetaGlobalInput) *ActionUserMailTemplateRecipientUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserMailTemplateRecipientUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserMailTemplateRecipientUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserMailTemplateRecipientUpdateInvocation) Call() (*ActionUserMailTemplateRecipientUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionUserMailTemplateRecipientUpdateInvocation) callAsBody() (*ActionUserMailTemplateRecipientUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserMailTemplateRecipientUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MailTemplateRecipient
	}
	return resp, err
}

func (inv *ActionUserMailTemplateRecipientUpdateInvocation) makeAllInputParams() *ActionUserMailTemplateRecipientUpdateRequest {
	return &ActionUserMailTemplateRecipientUpdateRequest{
		MailTemplateRecipient: inv.makeInputParams(),
		Meta:                  inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserMailTemplateRecipientUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Enabled") {
			ret["enabled"] = inv.Input.Enabled
		}
		if inv.IsParameterSelected("To") {
			ret["to"] = inv.Input.To
		}
	}

	return ret
}

func (inv *ActionUserMailTemplateRecipientUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
