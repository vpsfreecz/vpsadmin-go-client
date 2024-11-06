package client

import (
	"strings"
)

// ActionMailTemplateRecipientIndex is a type for action Mail_template.Recipient#Index
type ActionMailTemplateRecipientIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionMailTemplateRecipientIndex(client *Client) *ActionMailTemplateRecipientIndex {
	return &ActionMailTemplateRecipientIndex{
		Client: client,
	}
}

// ActionMailTemplateRecipientIndexMetaGlobalInput is a type for action global meta input parameters
type ActionMailTemplateRecipientIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionMailTemplateRecipientIndexMetaGlobalInput) SetCount(value bool) *ActionMailTemplateRecipientIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailTemplateRecipientIndexMetaGlobalInput) SetIncludes(value string) *ActionMailTemplateRecipientIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailTemplateRecipientIndexMetaGlobalInput) SetNo(value bool) *ActionMailTemplateRecipientIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateRecipientIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateRecipientIndexMetaGlobalInput) SelectParameters(params ...string) *ActionMailTemplateRecipientIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateRecipientIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailTemplateRecipientIndexInput is a type for action input parameters
type ActionMailTemplateRecipientIndexInput struct {
	FromId int64 `json:"from_id"`
	Limit  int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionMailTemplateRecipientIndexInput) SetFromId(value int64) *ActionMailTemplateRecipientIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionMailTemplateRecipientIndexInput) SetLimit(value int64) *ActionMailTemplateRecipientIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateRecipientIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateRecipientIndexInput) SelectParameters(params ...string) *ActionMailTemplateRecipientIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionMailTemplateRecipientIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionMailTemplateRecipientIndexInput) UnselectParameters(params ...string) *ActionMailTemplateRecipientIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionMailTemplateRecipientIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailTemplateRecipientIndexOutput is a type for action output parameters
type ActionMailTemplateRecipientIndexOutput struct {
	Id            int64                          `json:"id"`
	MailRecipient *ActionMailRecipientShowOutput `json:"mail_recipient"`
}

// Type for action response, including envelope
type ActionMailTemplateRecipientIndexResponse struct {
	Action *ActionMailTemplateRecipientIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Recipients []*ActionMailTemplateRecipientIndexOutput `json:"recipients"`
	}

	// Action output without the namespace
	Output []*ActionMailTemplateRecipientIndexOutput
}

// Prepare the action for invocation
func (action *ActionMailTemplateRecipientIndex) Prepare() *ActionMailTemplateRecipientIndexInvocation {
	return &ActionMailTemplateRecipientIndexInvocation{
		Action: action,
		Path:   "/v7.0/mail_templates/{mail_template_id}/recipients",
	}
}

// ActionMailTemplateRecipientIndexInvocation is used to configure action for invocation
type ActionMailTemplateRecipientIndexInvocation struct {
	// Pointer to the action
	Action *ActionMailTemplateRecipientIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMailTemplateRecipientIndexInput
	// Global meta input parameters
	MetaInput *ActionMailTemplateRecipientIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailTemplateRecipientIndexInvocation) SetPathParamInt(param string, value int64) *ActionMailTemplateRecipientIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailTemplateRecipientIndexInvocation) SetPathParamString(param string, value string) *ActionMailTemplateRecipientIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMailTemplateRecipientIndexInvocation) NewInput() *ActionMailTemplateRecipientIndexInput {
	inv.Input = &ActionMailTemplateRecipientIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMailTemplateRecipientIndexInvocation) SetInput(input *ActionMailTemplateRecipientIndexInput) *ActionMailTemplateRecipientIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMailTemplateRecipientIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionMailTemplateRecipientIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailTemplateRecipientIndexInvocation) NewMetaInput() *ActionMailTemplateRecipientIndexMetaGlobalInput {
	inv.MetaInput = &ActionMailTemplateRecipientIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailTemplateRecipientIndexInvocation) SetMetaInput(input *ActionMailTemplateRecipientIndexMetaGlobalInput) *ActionMailTemplateRecipientIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailTemplateRecipientIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMailTemplateRecipientIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailTemplateRecipientIndexInvocation) Call() (*ActionMailTemplateRecipientIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMailTemplateRecipientIndexInvocation) callAsQuery() (*ActionMailTemplateRecipientIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMailTemplateRecipientIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Recipients
	}
	return resp, err
}

func (inv *ActionMailTemplateRecipientIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["recipient[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["recipient[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionMailTemplateRecipientIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
