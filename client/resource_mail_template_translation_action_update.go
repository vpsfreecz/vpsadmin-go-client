package client

import (
	"strings"
)

// ActionMailTemplateTranslationUpdate is a type for action Mail_template.Translation#Update
type ActionMailTemplateTranslationUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionMailTemplateTranslationUpdate(client *Client) *ActionMailTemplateTranslationUpdate {
	return &ActionMailTemplateTranslationUpdate{
		Client: client,
	}
}

// ActionMailTemplateTranslationUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionMailTemplateTranslationUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailTemplateTranslationUpdateMetaGlobalInput) SetIncludes(value string) *ActionMailTemplateTranslationUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailTemplateTranslationUpdateMetaGlobalInput) SetNo(value bool) *ActionMailTemplateTranslationUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateTranslationUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateTranslationUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionMailTemplateTranslationUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateTranslationUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailTemplateTranslationUpdateInput is a type for action input parameters
type ActionMailTemplateTranslationUpdateInput struct {
	From       string `json:"from"`
	Language   int64  `json:"language"`
	ReplyTo    string `json:"reply_to"`
	ReturnPath string `json:"return_path"`
	Subject    string `json:"subject"`
	TextHtml   string `json:"text_html"`
	TextPlain  string `json:"text_plain"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionMailTemplateTranslationUpdateInput) SetFrom(value string) *ActionMailTemplateTranslationUpdateInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetLanguage sets parameter Language to value and selects it for sending
func (in *ActionMailTemplateTranslationUpdateInput) SetLanguage(value int64) *ActionMailTemplateTranslationUpdateInput {
	in.Language = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLanguageNil(false)
	in._selectedParameters["Language"] = nil
	return in
}

// SetLanguageNil sets parameter Language to nil and selects it for sending
func (in *ActionMailTemplateTranslationUpdateInput) SetLanguageNil(set bool) *ActionMailTemplateTranslationUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Language"] = nil
		in.SelectParameters("Language")
	} else {
		delete(in._nilParameters, "Language")
	}
	return in
}

// SetReplyTo sets parameter ReplyTo to value and selects it for sending
func (in *ActionMailTemplateTranslationUpdateInput) SetReplyTo(value string) *ActionMailTemplateTranslationUpdateInput {
	in.ReplyTo = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ReplyTo"] = nil
	return in
}

// SetReturnPath sets parameter ReturnPath to value and selects it for sending
func (in *ActionMailTemplateTranslationUpdateInput) SetReturnPath(value string) *ActionMailTemplateTranslationUpdateInput {
	in.ReturnPath = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ReturnPath"] = nil
	return in
}

// SetSubject sets parameter Subject to value and selects it for sending
func (in *ActionMailTemplateTranslationUpdateInput) SetSubject(value string) *ActionMailTemplateTranslationUpdateInput {
	in.Subject = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Subject"] = nil
	return in
}

// SetTextHtml sets parameter TextHtml to value and selects it for sending
func (in *ActionMailTemplateTranslationUpdateInput) SetTextHtml(value string) *ActionMailTemplateTranslationUpdateInput {
	in.TextHtml = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TextHtml"] = nil
	return in
}

// SetTextPlain sets parameter TextPlain to value and selects it for sending
func (in *ActionMailTemplateTranslationUpdateInput) SetTextPlain(value string) *ActionMailTemplateTranslationUpdateInput {
	in.TextPlain = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TextPlain"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateTranslationUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateTranslationUpdateInput) SelectParameters(params ...string) *ActionMailTemplateTranslationUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionMailTemplateTranslationUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionMailTemplateTranslationUpdateInput) UnselectParameters(params ...string) *ActionMailTemplateTranslationUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionMailTemplateTranslationUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailTemplateTranslationUpdateRequest is a type for the entire action request
type ActionMailTemplateTranslationUpdateRequest struct {
	Translation map[string]interface{} `json:"translation"`
	Meta        map[string]interface{} `json:"_meta"`
}

// ActionMailTemplateTranslationUpdateOutput is a type for action output parameters
type ActionMailTemplateTranslationUpdateOutput struct {
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
type ActionMailTemplateTranslationUpdateResponse struct {
	Action *ActionMailTemplateTranslationUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Translation *ActionMailTemplateTranslationUpdateOutput `json:"translation"`
	}

	// Action output without the namespace
	Output *ActionMailTemplateTranslationUpdateOutput
}

// Prepare the action for invocation
func (action *ActionMailTemplateTranslationUpdate) Prepare() *ActionMailTemplateTranslationUpdateInvocation {
	return &ActionMailTemplateTranslationUpdateInvocation{
		Action: action,
		Path:   "/v7.0/mail_templates/{mail_template_id}/translations/{translation_id}",
	}
}

// ActionMailTemplateTranslationUpdateInvocation is used to configure action for invocation
type ActionMailTemplateTranslationUpdateInvocation struct {
	// Pointer to the action
	Action *ActionMailTemplateTranslationUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMailTemplateTranslationUpdateInput
	// Global meta input parameters
	MetaInput *ActionMailTemplateTranslationUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailTemplateTranslationUpdateInvocation) SetPathParamInt(param string, value int64) *ActionMailTemplateTranslationUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailTemplateTranslationUpdateInvocation) SetPathParamString(param string, value string) *ActionMailTemplateTranslationUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMailTemplateTranslationUpdateInvocation) NewInput() *ActionMailTemplateTranslationUpdateInput {
	inv.Input = &ActionMailTemplateTranslationUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMailTemplateTranslationUpdateInvocation) SetInput(input *ActionMailTemplateTranslationUpdateInput) *ActionMailTemplateTranslationUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMailTemplateTranslationUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionMailTemplateTranslationUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailTemplateTranslationUpdateInvocation) NewMetaInput() *ActionMailTemplateTranslationUpdateMetaGlobalInput {
	inv.MetaInput = &ActionMailTemplateTranslationUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailTemplateTranslationUpdateInvocation) SetMetaInput(input *ActionMailTemplateTranslationUpdateMetaGlobalInput) *ActionMailTemplateTranslationUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailTemplateTranslationUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMailTemplateTranslationUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailTemplateTranslationUpdateInvocation) Call() (*ActionMailTemplateTranslationUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionMailTemplateTranslationUpdateInvocation) callAsBody() (*ActionMailTemplateTranslationUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMailTemplateTranslationUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Translation
	}
	return resp, err
}

func (inv *ActionMailTemplateTranslationUpdateInvocation) makeAllInputParams() *ActionMailTemplateTranslationUpdateRequest {
	return &ActionMailTemplateTranslationUpdateRequest{
		Translation: inv.makeInputParams(),
		Meta:        inv.makeMetaInputParams(),
	}
}

func (inv *ActionMailTemplateTranslationUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("From") {
			ret["from"] = inv.Input.From
		}
		if inv.IsParameterSelected("Language") {
			if inv.IsParameterNil("Language") {
				ret["language"] = nil
			} else {
				ret["language"] = inv.Input.Language
			}
		}
		if inv.IsParameterSelected("ReplyTo") {
			ret["reply_to"] = inv.Input.ReplyTo
		}
		if inv.IsParameterSelected("ReturnPath") {
			ret["return_path"] = inv.Input.ReturnPath
		}
		if inv.IsParameterSelected("Subject") {
			ret["subject"] = inv.Input.Subject
		}
		if inv.IsParameterSelected("TextHtml") {
			ret["text_html"] = inv.Input.TextHtml
		}
		if inv.IsParameterSelected("TextPlain") {
			ret["text_plain"] = inv.Input.TextPlain
		}
	}

	return ret
}

func (inv *ActionMailTemplateTranslationUpdateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
