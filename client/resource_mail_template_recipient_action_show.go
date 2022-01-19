package client

import (
	"strings"
)

// ActionMailTemplateRecipientShow is a type for action Mail_template.Recipient#Show
type ActionMailTemplateRecipientShow struct {
	// Pointer to client
	Client *Client
}

func NewActionMailTemplateRecipientShow(client *Client) *ActionMailTemplateRecipientShow {
	return &ActionMailTemplateRecipientShow{
		Client: client,
	}
}

// ActionMailTemplateRecipientShowMetaGlobalInput is a type for action global meta input parameters
type ActionMailTemplateRecipientShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailTemplateRecipientShowMetaGlobalInput) SetIncludes(value string) *ActionMailTemplateRecipientShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailTemplateRecipientShowMetaGlobalInput) SetNo(value bool) *ActionMailTemplateRecipientShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateRecipientShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateRecipientShowMetaGlobalInput) SelectParameters(params ...string) *ActionMailTemplateRecipientShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateRecipientShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionMailTemplateRecipientShowOutput is a type for action output parameters
type ActionMailTemplateRecipientShowOutput struct {
	Id int64 `json:"id"`
	MailRecipient *ActionMailRecipientShowOutput `json:"mail_recipient"`
}


// Type for action response, including envelope
type ActionMailTemplateRecipientShowResponse struct {
	Action *ActionMailTemplateRecipientShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Recipient *ActionMailTemplateRecipientShowOutput `json:"recipient"`
	}

	// Action output without the namespace
	Output *ActionMailTemplateRecipientShowOutput
}


// Prepare the action for invocation
func (action *ActionMailTemplateRecipientShow) Prepare() *ActionMailTemplateRecipientShowInvocation {
	return &ActionMailTemplateRecipientShowInvocation{
		Action: action,
		Path: "/v6.0/mail_templates/{mail_template_id}/recipients/{recipient_id}",
	}
}

// ActionMailTemplateRecipientShowInvocation is used to configure action for invocation
type ActionMailTemplateRecipientShowInvocation struct {
	// Pointer to the action
	Action *ActionMailTemplateRecipientShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMailTemplateRecipientShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailTemplateRecipientShowInvocation) SetPathParamInt(param string, value int64) *ActionMailTemplateRecipientShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailTemplateRecipientShowInvocation) SetPathParamString(param string, value string) *ActionMailTemplateRecipientShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailTemplateRecipientShowInvocation) NewMetaInput() *ActionMailTemplateRecipientShowMetaGlobalInput {
	inv.MetaInput = &ActionMailTemplateRecipientShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailTemplateRecipientShowInvocation) SetMetaInput(input *ActionMailTemplateRecipientShowMetaGlobalInput) *ActionMailTemplateRecipientShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailTemplateRecipientShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailTemplateRecipientShowInvocation) Call() (*ActionMailTemplateRecipientShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMailTemplateRecipientShowInvocation) callAsQuery() (*ActionMailTemplateRecipientShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMailTemplateRecipientShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Recipient
	}
	return resp, err
}




func (inv *ActionMailTemplateRecipientShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}

