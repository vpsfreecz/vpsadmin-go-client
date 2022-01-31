package client

import (
	"strings"
)

// ActionMailTemplateTranslationShow is a type for action Mail_template.Translation#Show
type ActionMailTemplateTranslationShow struct {
	// Pointer to client
	Client *Client
}

func NewActionMailTemplateTranslationShow(client *Client) *ActionMailTemplateTranslationShow {
	return &ActionMailTemplateTranslationShow{
		Client: client,
	}
}

// ActionMailTemplateTranslationShowMetaGlobalInput is a type for action global meta input parameters
type ActionMailTemplateTranslationShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailTemplateTranslationShowMetaGlobalInput) SetIncludes(value string) *ActionMailTemplateTranslationShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailTemplateTranslationShowMetaGlobalInput) SetNo(value bool) *ActionMailTemplateTranslationShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateTranslationShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateTranslationShowMetaGlobalInput) SelectParameters(params ...string) *ActionMailTemplateTranslationShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateTranslationShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailTemplateTranslationShowOutput is a type for action output parameters
type ActionMailTemplateTranslationShowOutput struct {
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
type ActionMailTemplateTranslationShowResponse struct {
	Action *ActionMailTemplateTranslationShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Translation *ActionMailTemplateTranslationShowOutput `json:"translation"`
	}

	// Action output without the namespace
	Output *ActionMailTemplateTranslationShowOutput
}

// Prepare the action for invocation
func (action *ActionMailTemplateTranslationShow) Prepare() *ActionMailTemplateTranslationShowInvocation {
	return &ActionMailTemplateTranslationShowInvocation{
		Action: action,
		Path:   "/v6.0/mail_templates/{mail_template_id}/translations/{translation_id}",
	}
}

// ActionMailTemplateTranslationShowInvocation is used to configure action for invocation
type ActionMailTemplateTranslationShowInvocation struct {
	// Pointer to the action
	Action *ActionMailTemplateTranslationShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMailTemplateTranslationShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailTemplateTranslationShowInvocation) SetPathParamInt(param string, value int64) *ActionMailTemplateTranslationShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailTemplateTranslationShowInvocation) SetPathParamString(param string, value string) *ActionMailTemplateTranslationShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailTemplateTranslationShowInvocation) NewMetaInput() *ActionMailTemplateTranslationShowMetaGlobalInput {
	inv.MetaInput = &ActionMailTemplateTranslationShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailTemplateTranslationShowInvocation) SetMetaInput(input *ActionMailTemplateTranslationShowMetaGlobalInput) *ActionMailTemplateTranslationShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailTemplateTranslationShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMailTemplateTranslationShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailTemplateTranslationShowInvocation) Call() (*ActionMailTemplateTranslationShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMailTemplateTranslationShowInvocation) callAsQuery() (*ActionMailTemplateTranslationShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMailTemplateTranslationShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Translation
	}
	return resp, err
}

func (inv *ActionMailTemplateTranslationShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
