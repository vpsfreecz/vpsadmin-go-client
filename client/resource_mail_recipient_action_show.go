package client

import (
	"strings"
)

// ActionMailRecipientShow is a type for action Mail_recipient#Show
type ActionMailRecipientShow struct {
	// Pointer to client
	Client *Client
}

func NewActionMailRecipientShow(client *Client) *ActionMailRecipientShow {
	return &ActionMailRecipientShow{
		Client: client,
	}
}

// ActionMailRecipientShowMetaGlobalInput is a type for action global meta input parameters
type ActionMailRecipientShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailRecipientShowMetaGlobalInput) SetNo(value bool) *ActionMailRecipientShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailRecipientShowMetaGlobalInput) SetIncludes(value string) *ActionMailRecipientShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailRecipientShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailRecipientShowMetaGlobalInput) SelectParameters(params ...string) *ActionMailRecipientShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailRecipientShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionMailRecipientShowOutput is a type for action output parameters
type ActionMailRecipientShowOutput struct {
	Id int64 `json:"id"`
	Label string `json:"label"`
	To string `json:"to"`
	Cc string `json:"cc"`
	Bcc string `json:"bcc"`
}


// Type for action response, including envelope
type ActionMailRecipientShowResponse struct {
	Action *ActionMailRecipientShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MailRecipient *ActionMailRecipientShowOutput `json:"mail_recipient"`
	}

	// Action output without the namespace
	Output *ActionMailRecipientShowOutput
}


// Prepare the action for invocation
func (action *ActionMailRecipientShow) Prepare() *ActionMailRecipientShowInvocation {
	return &ActionMailRecipientShowInvocation{
		Action: action,
		Path: "/v5.0/mail_recipients/:mail_recipient_id",
	}
}

// ActionMailRecipientShowInvocation is used to configure action for invocation
type ActionMailRecipientShowInvocation struct {
	// Pointer to the action
	Action *ActionMailRecipientShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMailRecipientShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailRecipientShowInvocation) SetPathParamInt(param string, value int64) *ActionMailRecipientShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailRecipientShowInvocation) SetPathParamString(param string, value string) *ActionMailRecipientShowInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailRecipientShowInvocation) NewMetaInput() *ActionMailRecipientShowMetaGlobalInput {
	inv.MetaInput = &ActionMailRecipientShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailRecipientShowInvocation) SetMetaInput(input *ActionMailRecipientShowMetaGlobalInput) *ActionMailRecipientShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailRecipientShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailRecipientShowInvocation) Call() (*ActionMailRecipientShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMailRecipientShowInvocation) callAsQuery() (*ActionMailRecipientShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMailRecipientShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MailRecipient
	}
	return resp, err
}




func (inv *ActionMailRecipientShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

