package client

import (
	"strings"
)

// ActionUserMailRoleRecipientShow is a type for action User.Mail_role_recipient#Show
type ActionUserMailRoleRecipientShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserMailRoleRecipientShow(client *Client) *ActionUserMailRoleRecipientShow {
	return &ActionUserMailRoleRecipientShow{
		Client: client,
	}
}

// ActionUserMailRoleRecipientShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserMailRoleRecipientShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserMailRoleRecipientShowMetaGlobalInput) SetNo(value bool) *ActionUserMailRoleRecipientShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserMailRoleRecipientShowMetaGlobalInput) SetIncludes(value string) *ActionUserMailRoleRecipientShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserMailRoleRecipientShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserMailRoleRecipientShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserMailRoleRecipientShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserMailRoleRecipientShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionUserMailRoleRecipientShowOutput is a type for action output parameters
type ActionUserMailRoleRecipientShowOutput struct {
	Id string `json:"id"`
	Label string `json:"label"`
	Description string `json:"description"`
	To string `json:"to"`
}


// Type for action response, including envelope
type ActionUserMailRoleRecipientShowResponse struct {
	Action *ActionUserMailRoleRecipientShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MailRoleRecipient *ActionUserMailRoleRecipientShowOutput `json:"mail_role_recipient"`
	}

	// Action output without the namespace
	Output *ActionUserMailRoleRecipientShowOutput
}


// Prepare the action for invocation
func (action *ActionUserMailRoleRecipientShow) Prepare() *ActionUserMailRoleRecipientShowInvocation {
	return &ActionUserMailRoleRecipientShowInvocation{
		Action: action,
		Path: "/v6.0/users/{user_id}/mail_role_recipients/{mail_role_recipient_id}",
	}
}

// ActionUserMailRoleRecipientShowInvocation is used to configure action for invocation
type ActionUserMailRoleRecipientShowInvocation struct {
	// Pointer to the action
	Action *ActionUserMailRoleRecipientShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserMailRoleRecipientShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserMailRoleRecipientShowInvocation) SetPathParamInt(param string, value int64) *ActionUserMailRoleRecipientShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserMailRoleRecipientShowInvocation) SetPathParamString(param string, value string) *ActionUserMailRoleRecipientShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserMailRoleRecipientShowInvocation) NewMetaInput() *ActionUserMailRoleRecipientShowMetaGlobalInput {
	inv.MetaInput = &ActionUserMailRoleRecipientShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserMailRoleRecipientShowInvocation) SetMetaInput(input *ActionUserMailRoleRecipientShowMetaGlobalInput) *ActionUserMailRoleRecipientShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserMailRoleRecipientShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserMailRoleRecipientShowInvocation) Call() (*ActionUserMailRoleRecipientShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserMailRoleRecipientShowInvocation) callAsQuery() (*ActionUserMailRoleRecipientShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserMailRoleRecipientShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MailRoleRecipient
	}
	return resp, err
}




func (inv *ActionUserMailRoleRecipientShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

