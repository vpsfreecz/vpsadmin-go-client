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
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
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

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserMailRoleRecipientIndexMetaGlobalInput) SetNo(value bool) *ActionUserMailRoleRecipientIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
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
	FromId int64 `json:"from_id"`
	Limit  int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionUserMailRoleRecipientIndexInput) SetFromId(value int64) *ActionUserMailRoleRecipientIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
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

// UnselectParameters unsets parameters from ActionUserMailRoleRecipientIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserMailRoleRecipientIndexInput) UnselectParameters(params ...string) *ActionUserMailRoleRecipientIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
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
	Description string `json:"description"`
	Id          string `json:"id"`
	Label       string `json:"label"`
	To          string `json:"to"`
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
		Path:   "/v7.0/users/{user_id}/mail_role_recipients",
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
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserMailRoleRecipientIndexInvocation) NewInput() *ActionUserMailRoleRecipientIndexInput {
	inv.Input = &ActionUserMailRoleRecipientIndexInput{}
	return inv.Input
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

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserMailRoleRecipientIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserMailRoleRecipientIndexInvocation) NewMetaInput() *ActionUserMailRoleRecipientIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserMailRoleRecipientIndexMetaGlobalInput{}
	return inv.MetaInput
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

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserMailRoleRecipientIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
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
		if inv.IsParameterSelected("FromId") {
			ret["mail_role_recipient[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["mail_role_recipient[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionUserMailRoleRecipientIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
