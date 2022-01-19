package client

import (
	"strings"
)

// ActionMailTemplateTranslationIndex is a type for action Mail_template.Translation#Index
type ActionMailTemplateTranslationIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionMailTemplateTranslationIndex(client *Client) *ActionMailTemplateTranslationIndex {
	return &ActionMailTemplateTranslationIndex{
		Client: client,
	}
}

// ActionMailTemplateTranslationIndexMetaGlobalInput is a type for action global meta input parameters
type ActionMailTemplateTranslationIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionMailTemplateTranslationIndexMetaGlobalInput) SetCount(value bool) *ActionMailTemplateTranslationIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailTemplateTranslationIndexMetaGlobalInput) SetIncludes(value string) *ActionMailTemplateTranslationIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailTemplateTranslationIndexMetaGlobalInput) SetNo(value bool) *ActionMailTemplateTranslationIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateTranslationIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateTranslationIndexMetaGlobalInput) SelectParameters(params ...string) *ActionMailTemplateTranslationIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateTranslationIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailTemplateTranslationIndexInput is a type for action input parameters
type ActionMailTemplateTranslationIndexInput struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionMailTemplateTranslationIndexInput) SetLimit(value int64) *ActionMailTemplateTranslationIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionMailTemplateTranslationIndexInput) SetOffset(value int64) *ActionMailTemplateTranslationIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateTranslationIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateTranslationIndexInput) SelectParameters(params ...string) *ActionMailTemplateTranslationIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateTranslationIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailTemplateTranslationIndexOutput is a type for action output parameters
type ActionMailTemplateTranslationIndexOutput struct {
	CreatedAt  string                    `json:"created_at"`
	From       string                    `json:"from"`
	Id         int64                     `json:"id"`
	Language   *ActionLanguageShowOutput `json:"language"`
	ReplyTo    string                    `json:"reply_to"`
	ReturnPath string                    `json:"return_path"`
	Subject    string                    `json:"subject"`
	TextHtml   string                    `json:"text_html"`
	TextPlain  string                    `json:"text_plain"`
	UpdatedAt  string                    `json:"updated_at"`
}

// Type for action response, including envelope
type ActionMailTemplateTranslationIndexResponse struct {
	Action *ActionMailTemplateTranslationIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Translations []*ActionMailTemplateTranslationIndexOutput `json:"translations"`
	}

	// Action output without the namespace
	Output []*ActionMailTemplateTranslationIndexOutput
}

// Prepare the action for invocation
func (action *ActionMailTemplateTranslationIndex) Prepare() *ActionMailTemplateTranslationIndexInvocation {
	return &ActionMailTemplateTranslationIndexInvocation{
		Action: action,
		Path:   "/v6.0/mail_templates/{mail_template_id}/translations",
	}
}

// ActionMailTemplateTranslationIndexInvocation is used to configure action for invocation
type ActionMailTemplateTranslationIndexInvocation struct {
	// Pointer to the action
	Action *ActionMailTemplateTranslationIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMailTemplateTranslationIndexInput
	// Global meta input parameters
	MetaInput *ActionMailTemplateTranslationIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailTemplateTranslationIndexInvocation) SetPathParamInt(param string, value int64) *ActionMailTemplateTranslationIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailTemplateTranslationIndexInvocation) SetPathParamString(param string, value string) *ActionMailTemplateTranslationIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMailTemplateTranslationIndexInvocation) NewInput() *ActionMailTemplateTranslationIndexInput {
	inv.Input = &ActionMailTemplateTranslationIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMailTemplateTranslationIndexInvocation) SetInput(input *ActionMailTemplateTranslationIndexInput) *ActionMailTemplateTranslationIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMailTemplateTranslationIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailTemplateTranslationIndexInvocation) NewMetaInput() *ActionMailTemplateTranslationIndexMetaGlobalInput {
	inv.MetaInput = &ActionMailTemplateTranslationIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailTemplateTranslationIndexInvocation) SetMetaInput(input *ActionMailTemplateTranslationIndexMetaGlobalInput) *ActionMailTemplateTranslationIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailTemplateTranslationIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailTemplateTranslationIndexInvocation) Call() (*ActionMailTemplateTranslationIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMailTemplateTranslationIndexInvocation) callAsQuery() (*ActionMailTemplateTranslationIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMailTemplateTranslationIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Translations
	}
	return resp, err
}

func (inv *ActionMailTemplateTranslationIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["translation[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["translation[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
	}
}

func (inv *ActionMailTemplateTranslationIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
