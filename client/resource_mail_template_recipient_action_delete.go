package client

import (
	"strings"
)

// ActionMailTemplateRecipientDelete is a type for action Mail_template.Recipient#Delete
type ActionMailTemplateRecipientDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionMailTemplateRecipientDelete(client *Client) *ActionMailTemplateRecipientDelete {
	return &ActionMailTemplateRecipientDelete{
		Client: client,
	}
}

// ActionMailTemplateRecipientDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionMailTemplateRecipientDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailTemplateRecipientDeleteMetaGlobalInput) SetIncludes(value string) *ActionMailTemplateRecipientDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailTemplateRecipientDeleteMetaGlobalInput) SetNo(value bool) *ActionMailTemplateRecipientDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateRecipientDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateRecipientDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionMailTemplateRecipientDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateRecipientDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionMailTemplateRecipientDeleteRequest is a type for the entire action request
type ActionMailTemplateRecipientDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionMailTemplateRecipientDeleteResponse struct {
	Action *ActionMailTemplateRecipientDelete `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionMailTemplateRecipientDelete) Prepare() *ActionMailTemplateRecipientDeleteInvocation {
	return &ActionMailTemplateRecipientDeleteInvocation{
		Action: action,
		Path: "/v6.0/mail_templates/{mail_template_id}/recipients/{recipient_id}",
	}
}

// ActionMailTemplateRecipientDeleteInvocation is used to configure action for invocation
type ActionMailTemplateRecipientDeleteInvocation struct {
	// Pointer to the action
	Action *ActionMailTemplateRecipientDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMailTemplateRecipientDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailTemplateRecipientDeleteInvocation) SetPathParamInt(param string, value int64) *ActionMailTemplateRecipientDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailTemplateRecipientDeleteInvocation) SetPathParamString(param string, value string) *ActionMailTemplateRecipientDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailTemplateRecipientDeleteInvocation) NewMetaInput() *ActionMailTemplateRecipientDeleteMetaGlobalInput {
	inv.MetaInput = &ActionMailTemplateRecipientDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailTemplateRecipientDeleteInvocation) SetMetaInput(input *ActionMailTemplateRecipientDeleteMetaGlobalInput) *ActionMailTemplateRecipientDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailTemplateRecipientDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailTemplateRecipientDeleteInvocation) Call() (*ActionMailTemplateRecipientDeleteResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionMailTemplateRecipientDeleteInvocation) callAsBody() (*ActionMailTemplateRecipientDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMailTemplateRecipientDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionMailTemplateRecipientDeleteInvocation) makeAllInputParams() *ActionMailTemplateRecipientDeleteRequest {
	return &ActionMailTemplateRecipientDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionMailTemplateRecipientDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
