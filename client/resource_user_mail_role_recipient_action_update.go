package client

import (
	"strings"
)

// ActionUserMailRoleRecipientUpdate is a type for action User.Mail_role_recipient#Update
type ActionUserMailRoleRecipientUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserMailRoleRecipientUpdate(client *Client) *ActionUserMailRoleRecipientUpdate {
	return &ActionUserMailRoleRecipientUpdate{
		Client: client,
	}
}

// ActionUserMailRoleRecipientUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionUserMailRoleRecipientUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserMailRoleRecipientUpdateMetaGlobalInput) SetNo(value bool) *ActionUserMailRoleRecipientUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserMailRoleRecipientUpdateMetaGlobalInput) SetIncludes(value string) *ActionUserMailRoleRecipientUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserMailRoleRecipientUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserMailRoleRecipientUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionUserMailRoleRecipientUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserMailRoleRecipientUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserMailRoleRecipientUpdateInput is a type for action input parameters
type ActionUserMailRoleRecipientUpdateInput struct {
	To string `json:"to"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionUserMailRoleRecipientUpdateInput) SetTo(value string) *ActionUserMailRoleRecipientUpdateInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserMailRoleRecipientUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserMailRoleRecipientUpdateInput) SelectParameters(params ...string) *ActionUserMailRoleRecipientUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserMailRoleRecipientUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserMailRoleRecipientUpdateRequest is a type for the entire action request
type ActionUserMailRoleRecipientUpdateRequest struct {
	MailRoleRecipient map[string]interface{} `json:"mail_role_recipient"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionUserMailRoleRecipientUpdateOutput is a type for action output parameters
type ActionUserMailRoleRecipientUpdateOutput struct {
	Id string `json:"id"`
	Label string `json:"label"`
	Description string `json:"description"`
	To string `json:"to"`
}


// Type for action response, including envelope
type ActionUserMailRoleRecipientUpdateResponse struct {
	Action *ActionUserMailRoleRecipientUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MailRoleRecipient *ActionUserMailRoleRecipientUpdateOutput `json:"mail_role_recipient"`
	}

	// Action output without the namespace
	Output *ActionUserMailRoleRecipientUpdateOutput
}


// Prepare the action for invocation
func (action *ActionUserMailRoleRecipientUpdate) Prepare() *ActionUserMailRoleRecipientUpdateInvocation {
	return &ActionUserMailRoleRecipientUpdateInvocation{
		Action: action,
		Path: "/v5.0/users/:user_id/mail_role_recipients/:mail_role_recipient_id",
	}
}

// ActionUserMailRoleRecipientUpdateInvocation is used to configure action for invocation
type ActionUserMailRoleRecipientUpdateInvocation struct {
	// Pointer to the action
	Action *ActionUserMailRoleRecipientUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserMailRoleRecipientUpdateInput
	// Global meta input parameters
	MetaInput *ActionUserMailRoleRecipientUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserMailRoleRecipientUpdateInvocation) SetPathParamInt(param string, value int64) *ActionUserMailRoleRecipientUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserMailRoleRecipientUpdateInvocation) SetPathParamString(param string, value string) *ActionUserMailRoleRecipientUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserMailRoleRecipientUpdateInvocation) NewInput() *ActionUserMailRoleRecipientUpdateInput {
	inv.Input = &ActionUserMailRoleRecipientUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserMailRoleRecipientUpdateInvocation) SetInput(input *ActionUserMailRoleRecipientUpdateInput) *ActionUserMailRoleRecipientUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserMailRoleRecipientUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserMailRoleRecipientUpdateInvocation) NewMetaInput() *ActionUserMailRoleRecipientUpdateMetaGlobalInput {
	inv.MetaInput = &ActionUserMailRoleRecipientUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserMailRoleRecipientUpdateInvocation) SetMetaInput(input *ActionUserMailRoleRecipientUpdateMetaGlobalInput) *ActionUserMailRoleRecipientUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserMailRoleRecipientUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserMailRoleRecipientUpdateInvocation) Call() (*ActionUserMailRoleRecipientUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionUserMailRoleRecipientUpdateInvocation) callAsBody() (*ActionUserMailRoleRecipientUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserMailRoleRecipientUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MailRoleRecipient
	}
	return resp, err
}




func (inv *ActionUserMailRoleRecipientUpdateInvocation) makeAllInputParams() *ActionUserMailRoleRecipientUpdateRequest {
	return &ActionUserMailRoleRecipientUpdateRequest{
		MailRoleRecipient: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserMailRoleRecipientUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("To") {
			ret["to"] = inv.Input.To
		}
	}

	return ret
}

func (inv *ActionUserMailRoleRecipientUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
