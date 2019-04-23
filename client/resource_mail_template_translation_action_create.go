package client

import (
	"strings"
)

// ActionMailTemplateTranslationCreate is a type for action Mail_template.Translation#Create
type ActionMailTemplateTranslationCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionMailTemplateTranslationCreate(client *Client) *ActionMailTemplateTranslationCreate {
	return &ActionMailTemplateTranslationCreate{
		Client: client,
	}
}

// ActionMailTemplateTranslationCreateMetaGlobalInput is a type for action global meta input parameters
type ActionMailTemplateTranslationCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailTemplateTranslationCreateMetaGlobalInput) SetNo(value bool) *ActionMailTemplateTranslationCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailTemplateTranslationCreateMetaGlobalInput) SetIncludes(value string) *ActionMailTemplateTranslationCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateTranslationCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateTranslationCreateMetaGlobalInput) SelectParameters(params ...string) *ActionMailTemplateTranslationCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateTranslationCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailTemplateTranslationCreateInput is a type for action input parameters
type ActionMailTemplateTranslationCreateInput struct {
	Language int64 `json:"language"`
	From string `json:"from"`
	ReplyTo string `json:"reply_to"`
	ReturnPath string `json:"return_path"`
	Subject string `json:"subject"`
	TextPlain string `json:"text_plain"`
	TextHtml string `json:"text_html"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLanguage sets parameter Language to value and selects it for sending
func (in *ActionMailTemplateTranslationCreateInput) SetLanguage(value int64) *ActionMailTemplateTranslationCreateInput {
	in.Language = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Language"] = nil
	return in
}
// SetFrom sets parameter From to value and selects it for sending
func (in *ActionMailTemplateTranslationCreateInput) SetFrom(value string) *ActionMailTemplateTranslationCreateInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}
// SetReplyTo sets parameter ReplyTo to value and selects it for sending
func (in *ActionMailTemplateTranslationCreateInput) SetReplyTo(value string) *ActionMailTemplateTranslationCreateInput {
	in.ReplyTo = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ReplyTo"] = nil
	return in
}
// SetReturnPath sets parameter ReturnPath to value and selects it for sending
func (in *ActionMailTemplateTranslationCreateInput) SetReturnPath(value string) *ActionMailTemplateTranslationCreateInput {
	in.ReturnPath = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ReturnPath"] = nil
	return in
}
// SetSubject sets parameter Subject to value and selects it for sending
func (in *ActionMailTemplateTranslationCreateInput) SetSubject(value string) *ActionMailTemplateTranslationCreateInput {
	in.Subject = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Subject"] = nil
	return in
}
// SetTextPlain sets parameter TextPlain to value and selects it for sending
func (in *ActionMailTemplateTranslationCreateInput) SetTextPlain(value string) *ActionMailTemplateTranslationCreateInput {
	in.TextPlain = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TextPlain"] = nil
	return in
}
// SetTextHtml sets parameter TextHtml to value and selects it for sending
func (in *ActionMailTemplateTranslationCreateInput) SetTextHtml(value string) *ActionMailTemplateTranslationCreateInput {
	in.TextHtml = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TextHtml"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateTranslationCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateTranslationCreateInput) SelectParameters(params ...string) *ActionMailTemplateTranslationCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateTranslationCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailTemplateTranslationCreateRequest is a type for the entire action request
type ActionMailTemplateTranslationCreateRequest struct {
	Translation map[string]interface{} `json:"translation"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionMailTemplateTranslationCreateOutput is a type for action output parameters
type ActionMailTemplateTranslationCreateOutput struct {
	Id int64 `json:"id"`
	Language *ActionLanguageShowOutput `json:"language"`
	From string `json:"from"`
	ReplyTo string `json:"reply_to"`
	ReturnPath string `json:"return_path"`
	Subject string `json:"subject"`
	TextPlain string `json:"text_plain"`
	TextHtml string `json:"text_html"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


// Type for action response, including envelope
type ActionMailTemplateTranslationCreateResponse struct {
	Action *ActionMailTemplateTranslationCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Translation *ActionMailTemplateTranslationCreateOutput `json:"translation"`
	}

	// Action output without the namespace
	Output *ActionMailTemplateTranslationCreateOutput
}


// Prepare the action for invocation
func (action *ActionMailTemplateTranslationCreate) Prepare() *ActionMailTemplateTranslationCreateInvocation {
	return &ActionMailTemplateTranslationCreateInvocation{
		Action: action,
		Path: "/v5.0/mail_templates/:mail_template_id/translations",
	}
}

// ActionMailTemplateTranslationCreateInvocation is used to configure action for invocation
type ActionMailTemplateTranslationCreateInvocation struct {
	// Pointer to the action
	Action *ActionMailTemplateTranslationCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMailTemplateTranslationCreateInput
	// Global meta input parameters
	MetaInput *ActionMailTemplateTranslationCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailTemplateTranslationCreateInvocation) SetPathParamInt(param string, value int64) *ActionMailTemplateTranslationCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailTemplateTranslationCreateInvocation) SetPathParamString(param string, value string) *ActionMailTemplateTranslationCreateInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetInput provides input parameters to send to the API
func (inv *ActionMailTemplateTranslationCreateInvocation) SetInput(input *ActionMailTemplateTranslationCreateInput) *ActionMailTemplateTranslationCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMailTemplateTranslationCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailTemplateTranslationCreateInvocation) SetMetaInput(input *ActionMailTemplateTranslationCreateMetaGlobalInput) *ActionMailTemplateTranslationCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailTemplateTranslationCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailTemplateTranslationCreateInvocation) Call() (*ActionMailTemplateTranslationCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionMailTemplateTranslationCreateInvocation) callAsBody() (*ActionMailTemplateTranslationCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMailTemplateTranslationCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Translation
	}
	return resp, err
}




func (inv *ActionMailTemplateTranslationCreateInvocation) makeAllInputParams() *ActionMailTemplateTranslationCreateRequest {
	return &ActionMailTemplateTranslationCreateRequest{
		Translation: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionMailTemplateTranslationCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Language") {
			ret["language"] = inv.Input.Language
		}
		if inv.IsParameterSelected("From") {
			ret["from"] = inv.Input.From
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
		if inv.IsParameterSelected("TextPlain") {
			ret["text_plain"] = inv.Input.TextPlain
		}
		if inv.IsParameterSelected("TextHtml") {
			ret["text_html"] = inv.Input.TextHtml
		}
	}

	return ret
}

func (inv *ActionMailTemplateTranslationCreateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
