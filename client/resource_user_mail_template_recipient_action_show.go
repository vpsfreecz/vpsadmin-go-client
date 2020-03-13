package client

import (
	"strings"
)

// ActionUserMailTemplateRecipientShow is a type for action User.Mail_template_recipient#Show
type ActionUserMailTemplateRecipientShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserMailTemplateRecipientShow(client *Client) *ActionUserMailTemplateRecipientShow {
	return &ActionUserMailTemplateRecipientShow{
		Client: client,
	}
}

// ActionUserMailTemplateRecipientShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserMailTemplateRecipientShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserMailTemplateRecipientShowMetaGlobalInput) SetNo(value bool) *ActionUserMailTemplateRecipientShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserMailTemplateRecipientShowMetaGlobalInput) SetIncludes(value string) *ActionUserMailTemplateRecipientShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserMailTemplateRecipientShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserMailTemplateRecipientShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserMailTemplateRecipientShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserMailTemplateRecipientShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionUserMailTemplateRecipientShowOutput is a type for action output parameters
type ActionUserMailTemplateRecipientShowOutput struct {
	Id string `json:"id"`
	Label string `json:"label"`
	Description string `json:"description"`
	To string `json:"to"`
	Enabled bool `json:"enabled"`
}


// Type for action response, including envelope
type ActionUserMailTemplateRecipientShowResponse struct {
	Action *ActionUserMailTemplateRecipientShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MailTemplateRecipient *ActionUserMailTemplateRecipientShowOutput `json:"mail_template_recipient"`
	}

	// Action output without the namespace
	Output *ActionUserMailTemplateRecipientShowOutput
}


// Prepare the action for invocation
func (action *ActionUserMailTemplateRecipientShow) Prepare() *ActionUserMailTemplateRecipientShowInvocation {
	return &ActionUserMailTemplateRecipientShowInvocation{
		Action: action,
		Path: "/v6.0/users/{user_id}/mail_template_recipients/{mail_template_recipient_id}",
	}
}

// ActionUserMailTemplateRecipientShowInvocation is used to configure action for invocation
type ActionUserMailTemplateRecipientShowInvocation struct {
	// Pointer to the action
	Action *ActionUserMailTemplateRecipientShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserMailTemplateRecipientShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserMailTemplateRecipientShowInvocation) SetPathParamInt(param string, value int64) *ActionUserMailTemplateRecipientShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserMailTemplateRecipientShowInvocation) SetPathParamString(param string, value string) *ActionUserMailTemplateRecipientShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserMailTemplateRecipientShowInvocation) NewMetaInput() *ActionUserMailTemplateRecipientShowMetaGlobalInput {
	inv.MetaInput = &ActionUserMailTemplateRecipientShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserMailTemplateRecipientShowInvocation) SetMetaInput(input *ActionUserMailTemplateRecipientShowMetaGlobalInput) *ActionUserMailTemplateRecipientShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserMailTemplateRecipientShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserMailTemplateRecipientShowInvocation) Call() (*ActionUserMailTemplateRecipientShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserMailTemplateRecipientShowInvocation) callAsQuery() (*ActionUserMailTemplateRecipientShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserMailTemplateRecipientShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MailTemplateRecipient
	}
	return resp, err
}




func (inv *ActionUserMailTemplateRecipientShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

