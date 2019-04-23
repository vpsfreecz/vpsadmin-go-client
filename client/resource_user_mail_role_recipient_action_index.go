package client

import (
	"strings"
)

// ActionUserMailRoleRecipientIndex is a type for action User.Mail_role_recipient#Index
type ActionUserMailRoleRecipientIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserMailRoleRecipientIndex(client *Client) *ActionUserMailRoleRecipientIndex {
	return &ActionUserMailRoleRecipientIndex{
		Client: client,
	}
}

// ActionUserMailRoleRecipientIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserMailRoleRecipientIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserMailRoleRecipientIndexMetaGlobalInput) SetNo(value bool) *ActionUserMailRoleRecipientIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserMailRoleRecipientIndexMetaGlobalInput) SetCount(value bool) *ActionUserMailRoleRecipientIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserMailRoleRecipientIndexMetaGlobalInput) SetIncludes(value string) *ActionUserMailRoleRecipientIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserMailRoleRecipientIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserMailRoleRecipientIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserMailRoleRecipientIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserMailRoleRecipientIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserMailRoleRecipientIndexInput is a type for action input parameters
type ActionUserMailRoleRecipientIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionUserMailRoleRecipientIndexInput) SetOffset(value int64) *ActionUserMailRoleRecipientIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserMailRoleRecipientIndexInput) SetLimit(value int64) *ActionUserMailRoleRecipientIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserMailRoleRecipientIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserMailRoleRecipientIndexInput) SelectParameters(params ...string) *ActionUserMailRoleRecipientIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserMailRoleRecipientIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionUserMailRoleRecipientIndexOutput is a type for action output parameters
type ActionUserMailRoleRecipientIndexOutput struct {
	Id string `json:"id"`
	Label string `json:"label"`
	Description string `json:"description"`
	To string `json:"to"`
}


// Type for action response, including envelope
type ActionUserMailRoleRecipientIndexResponse struct {
	Action *ActionUserMailRoleRecipientIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MailRoleRecipients []*ActionUserMailRoleRecipientIndexOutput `json:"mail_role_recipients"`
	}

	// Action output without the namespace
	Output []*ActionUserMailRoleRecipientIndexOutput
}


// Prepare the action for invocation
func (action *ActionUserMailRoleRecipientIndex) Prepare() *ActionUserMailRoleRecipientIndexInvocation {
	return &ActionUserMailRoleRecipientIndexInvocation{
		Action: action,
		Path: "/v5.0/users/:user_id/mail_role_recipients",
	}
}

// ActionUserMailRoleRecipientIndexInvocation is used to configure action for invocation
type ActionUserMailRoleRecipientIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserMailRoleRecipientIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserMailRoleRecipientIndexInput
	// Global meta input parameters
	MetaInput *ActionUserMailRoleRecipientIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserMailRoleRecipientIndexInvocation) SetPathParamInt(param string, value int64) *ActionUserMailRoleRecipientIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserMailRoleRecipientIndexInvocation) SetPathParamString(param string, value string) *ActionUserMailRoleRecipientIndexInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserMailRoleRecipientIndexInvocation) SetInput(input *ActionUserMailRoleRecipientIndexInput) *ActionUserMailRoleRecipientIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserMailRoleRecipientIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserMailRoleRecipientIndexInvocation) SetMetaInput(input *ActionUserMailRoleRecipientIndexMetaGlobalInput) *ActionUserMailRoleRecipientIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserMailRoleRecipientIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserMailRoleRecipientIndexInvocation) Call() (*ActionUserMailRoleRecipientIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserMailRoleRecipientIndexInvocation) callAsQuery() (*ActionUserMailRoleRecipientIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserMailRoleRecipientIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MailRoleRecipients
	}
	return resp, err
}



func (inv *ActionUserMailRoleRecipientIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["mail_role_recipient[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["mail_role_recipient[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionUserMailRoleRecipientIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

