package client

import (
	"net/url"
	"strings"
)

// ActionNotificationTemplateVariantCreate is a type for action Notification_template.Variant#Create
type ActionNotificationTemplateVariantCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationTemplateVariantCreate(client *Client) *ActionNotificationTemplateVariantCreate {
	return &ActionNotificationTemplateVariantCreate{
		Client: client,
	}
}

// ActionNotificationTemplateVariantCreateMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationTemplateVariantCreateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationTemplateVariantCreateMetaGlobalInput) SetIncludes(value string) *ActionNotificationTemplateVariantCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationTemplateVariantCreateMetaGlobalInput) SetNo(value bool) *ActionNotificationTemplateVariantCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTemplateVariantCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTemplateVariantCreateMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationTemplateVariantCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationTemplateVariantCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTemplateVariantCreateInput is a type for action input parameters
type ActionNotificationTemplateVariantCreateInput struct {
	From       string      "json:\"from\""
	Html       string      "json:\"html\""
	Language   int64       "json:\"language\""
	Options    interface{} "json:\"options\""
	Protocol   string      "json:\"protocol\""
	ReplyTo    string      "json:\"reply_to\""
	ReturnPath string      "json:\"return_path\""
	Subject    string      "json:\"subject\""
	Text       string      "json:\"text\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionNotificationTemplateVariantCreateInput) SetFrom(value string) *ActionNotificationTemplateVariantCreateInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetFromNil(false)
	in._selectedParameters["From"] = nil
	return in
}

// SetFromNil sets parameter From to nil and selects it for sending
func (in *ActionNotificationTemplateVariantCreateInput) SetFromNil(set bool) *ActionNotificationTemplateVariantCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["From"] = nil
		in.SelectParameters("From")
	} else {
		delete(in._nilParameters, "From")
	}
	return in
}

// SetHtml sets parameter Html to value and selects it for sending
func (in *ActionNotificationTemplateVariantCreateInput) SetHtml(value string) *ActionNotificationTemplateVariantCreateInput {
	in.Html = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetHtmlNil(false)
	in._selectedParameters["Html"] = nil
	return in
}

// SetHtmlNil sets parameter Html to nil and selects it for sending
func (in *ActionNotificationTemplateVariantCreateInput) SetHtmlNil(set bool) *ActionNotificationTemplateVariantCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Html"] = nil
		in.SelectParameters("Html")
	} else {
		delete(in._nilParameters, "Html")
	}
	return in
}

// SetLanguage sets parameter Language to value and selects it for sending
func (in *ActionNotificationTemplateVariantCreateInput) SetLanguage(value int64) *ActionNotificationTemplateVariantCreateInput {
	in.Language = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Language"] = nil
	return in
}

// SetOptions sets parameter Options to value and selects it for sending
func (in *ActionNotificationTemplateVariantCreateInput) SetOptions(value interface{}) *ActionNotificationTemplateVariantCreateInput {
	in.Options = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetOptionsNil(false)
	in._selectedParameters["Options"] = nil
	return in
}

// SetOptionsNil sets parameter Options to nil and selects it for sending
func (in *ActionNotificationTemplateVariantCreateInput) SetOptionsNil(set bool) *ActionNotificationTemplateVariantCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Options"] = nil
		in.SelectParameters("Options")
	} else {
		delete(in._nilParameters, "Options")
	}
	return in
}

// SetProtocol sets parameter Protocol to value and selects it for sending
func (in *ActionNotificationTemplateVariantCreateInput) SetProtocol(value string) *ActionNotificationTemplateVariantCreateInput {
	in.Protocol = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Protocol"] = nil
	return in
}

// SetReplyTo sets parameter ReplyTo to value and selects it for sending
func (in *ActionNotificationTemplateVariantCreateInput) SetReplyTo(value string) *ActionNotificationTemplateVariantCreateInput {
	in.ReplyTo = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetReplyToNil(false)
	in._selectedParameters["ReplyTo"] = nil
	return in
}

// SetReplyToNil sets parameter ReplyTo to nil and selects it for sending
func (in *ActionNotificationTemplateVariantCreateInput) SetReplyToNil(set bool) *ActionNotificationTemplateVariantCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["ReplyTo"] = nil
		in.SelectParameters("ReplyTo")
	} else {
		delete(in._nilParameters, "ReplyTo")
	}
	return in
}

// SetReturnPath sets parameter ReturnPath to value and selects it for sending
func (in *ActionNotificationTemplateVariantCreateInput) SetReturnPath(value string) *ActionNotificationTemplateVariantCreateInput {
	in.ReturnPath = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetReturnPathNil(false)
	in._selectedParameters["ReturnPath"] = nil
	return in
}

// SetReturnPathNil sets parameter ReturnPath to nil and selects it for sending
func (in *ActionNotificationTemplateVariantCreateInput) SetReturnPathNil(set bool) *ActionNotificationTemplateVariantCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["ReturnPath"] = nil
		in.SelectParameters("ReturnPath")
	} else {
		delete(in._nilParameters, "ReturnPath")
	}
	return in
}

// SetSubject sets parameter Subject to value and selects it for sending
func (in *ActionNotificationTemplateVariantCreateInput) SetSubject(value string) *ActionNotificationTemplateVariantCreateInput {
	in.Subject = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetSubjectNil(false)
	in._selectedParameters["Subject"] = nil
	return in
}

// SetSubjectNil sets parameter Subject to nil and selects it for sending
func (in *ActionNotificationTemplateVariantCreateInput) SetSubjectNil(set bool) *ActionNotificationTemplateVariantCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Subject"] = nil
		in.SelectParameters("Subject")
	} else {
		delete(in._nilParameters, "Subject")
	}
	return in
}

// SetText sets parameter Text to value and selects it for sending
func (in *ActionNotificationTemplateVariantCreateInput) SetText(value string) *ActionNotificationTemplateVariantCreateInput {
	in.Text = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetTextNil(false)
	in._selectedParameters["Text"] = nil
	return in
}

// SetTextNil sets parameter Text to nil and selects it for sending
func (in *ActionNotificationTemplateVariantCreateInput) SetTextNil(set bool) *ActionNotificationTemplateVariantCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Text"] = nil
		in.SelectParameters("Text")
	} else {
		delete(in._nilParameters, "Text")
	}
	return in
}

// SelectParameters sets parameters from ActionNotificationTemplateVariantCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTemplateVariantCreateInput) SelectParameters(params ...string) *ActionNotificationTemplateVariantCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNotificationTemplateVariantCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNotificationTemplateVariantCreateInput) UnselectParameters(params ...string) *ActionNotificationTemplateVariantCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNotificationTemplateVariantCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTemplateVariantCreateRequest is a type for the entire action request
type ActionNotificationTemplateVariantCreateRequest struct {
	Variant map[string]interface{} "json:\"variant\""
	Meta    map[string]interface{} "json:\"_meta\""
}

// ActionNotificationTemplateVariantCreateOutput is a type for action output parameters
type ActionNotificationTemplateVariantCreateOutput struct {
	CreatedAt  string                    "json:\"created_at\""
	From       string                    "json:\"from\""
	Html       string                    "json:\"html\""
	Id         int64                     "json:\"id\""
	Language   *ActionLanguageShowOutput "json:\"language\""
	Options    interface{}               "json:\"options\""
	Protocol   string                    "json:\"protocol\""
	ReplyTo    string                    "json:\"reply_to\""
	ReturnPath string                    "json:\"return_path\""
	Subject    string                    "json:\"subject\""
	Text       string                    "json:\"text\""
	UpdatedAt  string                    "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionNotificationTemplateVariantCreateResponse struct {
	Action *ActionNotificationTemplateVariantCreate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Variant *ActionNotificationTemplateVariantCreateOutput "json:\"variant\""
	}

	// Action output without the namespace
	Output *ActionNotificationTemplateVariantCreateOutput
}

// Prepare the action for invocation
func (action *ActionNotificationTemplateVariantCreate) Prepare() *ActionNotificationTemplateVariantCreateInvocation {
	return &ActionNotificationTemplateVariantCreateInvocation{
		Action: action,
		Path:   "/v7.0/notification_templates/{notification_template_id}/variants",
	}
}

// ActionNotificationTemplateVariantCreateInvocation is used to configure action for invocation
type ActionNotificationTemplateVariantCreateInvocation struct {
	// Pointer to the action
	Action *ActionNotificationTemplateVariantCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNotificationTemplateVariantCreateInput
	// Global meta input parameters
	MetaInput *ActionNotificationTemplateVariantCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationTemplateVariantCreateInvocation) SetPathParamInt(param string, value int64) *ActionNotificationTemplateVariantCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationTemplateVariantCreateInvocation) SetPathParamString(param string, value string) *ActionNotificationTemplateVariantCreateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNotificationTemplateVariantCreateInvocation) NewInput() *ActionNotificationTemplateVariantCreateInput {
	inv.Input = &ActionNotificationTemplateVariantCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNotificationTemplateVariantCreateInvocation) SetInput(input *ActionNotificationTemplateVariantCreateInput) *ActionNotificationTemplateVariantCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNotificationTemplateVariantCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNotificationTemplateVariantCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationTemplateVariantCreateInvocation) NewMetaInput() *ActionNotificationTemplateVariantCreateMetaGlobalInput {
	inv.MetaInput = &ActionNotificationTemplateVariantCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationTemplateVariantCreateInvocation) SetMetaInput(input *ActionNotificationTemplateVariantCreateMetaGlobalInput) *ActionNotificationTemplateVariantCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationTemplateVariantCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationTemplateVariantCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationTemplateVariantCreateInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("Language") {
			if !inv.IsParameterNil("Language") {
				if inv.Input.Language < 0 {
					verr.Add("language", "not a valid resource id")
				}
			}
		}
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNotificationTemplateVariantCreateInvocation) Call() (*ActionNotificationTemplateVariantCreateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNotificationTemplateVariantCreateInvocation) callAsBody() (*ActionNotificationTemplateVariantCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNotificationTemplateVariantCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Variant
	}
	return resp, err
}

func (inv *ActionNotificationTemplateVariantCreateInvocation) makeAllInputParams() *ActionNotificationTemplateVariantCreateRequest {
	return &ActionNotificationTemplateVariantCreateRequest{
		Variant: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionNotificationTemplateVariantCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("From") {
			if inv.IsParameterNil("From") {
				ret["from"] = nil
			} else {
				ret["from"] = inv.Input.From
			}
		}
		if inv.IsParameterSelected("Html") {
			if inv.IsParameterNil("Html") {
				ret["html"] = nil
			} else {
				ret["html"] = inv.Input.Html
			}
		}
		if inv.IsParameterSelected("Language") {
			ret["language"] = inv.Input.Language
		}
		if inv.IsParameterSelected("Options") {
			if inv.IsParameterNil("Options") {
				ret["options"] = nil
			} else {
				ret["options"] = inv.Input.Options
			}
		}
		if inv.IsParameterSelected("Protocol") {
			ret["protocol"] = inv.Input.Protocol
		}
		if inv.IsParameterSelected("ReplyTo") {
			if inv.IsParameterNil("ReplyTo") {
				ret["reply_to"] = nil
			} else {
				ret["reply_to"] = inv.Input.ReplyTo
			}
		}
		if inv.IsParameterSelected("ReturnPath") {
			if inv.IsParameterNil("ReturnPath") {
				ret["return_path"] = nil
			} else {
				ret["return_path"] = inv.Input.ReturnPath
			}
		}
		if inv.IsParameterSelected("Subject") {
			if inv.IsParameterNil("Subject") {
				ret["subject"] = nil
			} else {
				ret["subject"] = inv.Input.Subject
			}
		}
		if inv.IsParameterSelected("Text") {
			if inv.IsParameterNil("Text") {
				ret["text"] = nil
			} else {
				ret["text"] = inv.Input.Text
			}
		}
	}

	return ret
}

func (inv *ActionNotificationTemplateVariantCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
