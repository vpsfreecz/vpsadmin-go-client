package client

import (
	"strings"
)

// ActionUserMailTemplateRecipientIndex is a type for action User.Mail_template_recipient#Index
type ActionUserMailTemplateRecipientIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserMailTemplateRecipientIndex(client *Client) *ActionUserMailTemplateRecipientIndex {
	return &ActionUserMailTemplateRecipientIndex{
		Client: client,
	}
}

// ActionUserMailTemplateRecipientIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserMailTemplateRecipientIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserMailTemplateRecipientIndexMetaGlobalInput) SetCount(value bool) *ActionUserMailTemplateRecipientIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserMailTemplateRecipientIndexMetaGlobalInput) SetIncludes(value string) *ActionUserMailTemplateRecipientIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserMailTemplateRecipientIndexMetaGlobalInput) SetNo(value bool) *ActionUserMailTemplateRecipientIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserMailTemplateRecipientIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserMailTemplateRecipientIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserMailTemplateRecipientIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserMailTemplateRecipientIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserMailTemplateRecipientIndexInput is a type for action input parameters
type ActionUserMailTemplateRecipientIndexInput struct {
	FromId int64 `json:"from_id"`
	Limit  int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionUserMailTemplateRecipientIndexInput) SetFromId(value int64) *ActionUserMailTemplateRecipientIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserMailTemplateRecipientIndexInput) SetLimit(value int64) *ActionUserMailTemplateRecipientIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserMailTemplateRecipientIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserMailTemplateRecipientIndexInput) SelectParameters(params ...string) *ActionUserMailTemplateRecipientIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserMailTemplateRecipientIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserMailTemplateRecipientIndexInput) UnselectParameters(params ...string) *ActionUserMailTemplateRecipientIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserMailTemplateRecipientIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserMailTemplateRecipientIndexOutput is a type for action output parameters
type ActionUserMailTemplateRecipientIndexOutput struct {
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Id          string `json:"id"`
	Label       string `json:"label"`
	To          string `json:"to"`
}

// Type for action response, including envelope
type ActionUserMailTemplateRecipientIndexResponse struct {
	Action *ActionUserMailTemplateRecipientIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MailTemplateRecipients []*ActionUserMailTemplateRecipientIndexOutput `json:"mail_template_recipients"`
	}

	// Action output without the namespace
	Output []*ActionUserMailTemplateRecipientIndexOutput
}

// Prepare the action for invocation
func (action *ActionUserMailTemplateRecipientIndex) Prepare() *ActionUserMailTemplateRecipientIndexInvocation {
	return &ActionUserMailTemplateRecipientIndexInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/mail_template_recipients",
	}
}

// ActionUserMailTemplateRecipientIndexInvocation is used to configure action for invocation
type ActionUserMailTemplateRecipientIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserMailTemplateRecipientIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserMailTemplateRecipientIndexInput
	// Global meta input parameters
	MetaInput *ActionUserMailTemplateRecipientIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserMailTemplateRecipientIndexInvocation) SetPathParamInt(param string, value int64) *ActionUserMailTemplateRecipientIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserMailTemplateRecipientIndexInvocation) SetPathParamString(param string, value string) *ActionUserMailTemplateRecipientIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserMailTemplateRecipientIndexInvocation) NewInput() *ActionUserMailTemplateRecipientIndexInput {
	inv.Input = &ActionUserMailTemplateRecipientIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserMailTemplateRecipientIndexInvocation) SetInput(input *ActionUserMailTemplateRecipientIndexInput) *ActionUserMailTemplateRecipientIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserMailTemplateRecipientIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserMailTemplateRecipientIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserMailTemplateRecipientIndexInvocation) NewMetaInput() *ActionUserMailTemplateRecipientIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserMailTemplateRecipientIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserMailTemplateRecipientIndexInvocation) SetMetaInput(input *ActionUserMailTemplateRecipientIndexMetaGlobalInput) *ActionUserMailTemplateRecipientIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserMailTemplateRecipientIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserMailTemplateRecipientIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserMailTemplateRecipientIndexInvocation) Call() (*ActionUserMailTemplateRecipientIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserMailTemplateRecipientIndexInvocation) callAsQuery() (*ActionUserMailTemplateRecipientIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserMailTemplateRecipientIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MailTemplateRecipients
	}
	return resp, err
}

func (inv *ActionUserMailTemplateRecipientIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["mail_template_recipient[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["mail_template_recipient[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionUserMailTemplateRecipientIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
