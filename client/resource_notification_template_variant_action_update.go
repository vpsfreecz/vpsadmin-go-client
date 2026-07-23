package client

import (
	"net/url"
	"strings"
)

// ActionNotificationTemplateVariantUpdate is a type for action Notification_template.Variant#Update
type ActionNotificationTemplateVariantUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationTemplateVariantUpdate(client *Client) *ActionNotificationTemplateVariantUpdate {
	return &ActionNotificationTemplateVariantUpdate{
		Client: client,
	}
}

// ActionNotificationTemplateVariantUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationTemplateVariantUpdateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationTemplateVariantUpdateMetaGlobalInput) SetIncludes(value string) *ActionNotificationTemplateVariantUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationTemplateVariantUpdateMetaGlobalInput) SetNo(value bool) *ActionNotificationTemplateVariantUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTemplateVariantUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTemplateVariantUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationTemplateVariantUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationTemplateVariantUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTemplateVariantUpdateInput is a type for action input parameters
type ActionNotificationTemplateVariantUpdateInput struct {
	From       string "json:\"from\""
	Html       string "json:\"html\""
	Language   int64  "json:\"language\""
	Protocol   string "json:\"protocol\""
	ReplyTo    string "json:\"reply_to\""
	ReturnPath string "json:\"return_path\""
	Subject    string "json:\"subject\""
	Text       string "json:\"text\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionNotificationTemplateVariantUpdateInput) SetFrom(value string) *ActionNotificationTemplateVariantUpdateInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetFromNil(false)
	in._selectedParameters["From"] = nil
	return in
}

// SetFromNil sets parameter From to nil and selects it for sending
func (in *ActionNotificationTemplateVariantUpdateInput) SetFromNil(set bool) *ActionNotificationTemplateVariantUpdateInput {
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
func (in *ActionNotificationTemplateVariantUpdateInput) SetHtml(value string) *ActionNotificationTemplateVariantUpdateInput {
	in.Html = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetHtmlNil(false)
	in._selectedParameters["Html"] = nil
	return in
}

// SetHtmlNil sets parameter Html to nil and selects it for sending
func (in *ActionNotificationTemplateVariantUpdateInput) SetHtmlNil(set bool) *ActionNotificationTemplateVariantUpdateInput {
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
func (in *ActionNotificationTemplateVariantUpdateInput) SetLanguage(value int64) *ActionNotificationTemplateVariantUpdateInput {
	in.Language = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Language"] = nil
	return in
}

// SetProtocol sets parameter Protocol to value and selects it for sending
func (in *ActionNotificationTemplateVariantUpdateInput) SetProtocol(value string) *ActionNotificationTemplateVariantUpdateInput {
	in.Protocol = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Protocol"] = nil
	return in
}

// SetReplyTo sets parameter ReplyTo to value and selects it for sending
func (in *ActionNotificationTemplateVariantUpdateInput) SetReplyTo(value string) *ActionNotificationTemplateVariantUpdateInput {
	in.ReplyTo = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetReplyToNil(false)
	in._selectedParameters["ReplyTo"] = nil
	return in
}

// SetReplyToNil sets parameter ReplyTo to nil and selects it for sending
func (in *ActionNotificationTemplateVariantUpdateInput) SetReplyToNil(set bool) *ActionNotificationTemplateVariantUpdateInput {
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
func (in *ActionNotificationTemplateVariantUpdateInput) SetReturnPath(value string) *ActionNotificationTemplateVariantUpdateInput {
	in.ReturnPath = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetReturnPathNil(false)
	in._selectedParameters["ReturnPath"] = nil
	return in
}

// SetReturnPathNil sets parameter ReturnPath to nil and selects it for sending
func (in *ActionNotificationTemplateVariantUpdateInput) SetReturnPathNil(set bool) *ActionNotificationTemplateVariantUpdateInput {
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
func (in *ActionNotificationTemplateVariantUpdateInput) SetSubject(value string) *ActionNotificationTemplateVariantUpdateInput {
	in.Subject = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetSubjectNil(false)
	in._selectedParameters["Subject"] = nil
	return in
}

// SetSubjectNil sets parameter Subject to nil and selects it for sending
func (in *ActionNotificationTemplateVariantUpdateInput) SetSubjectNil(set bool) *ActionNotificationTemplateVariantUpdateInput {
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
func (in *ActionNotificationTemplateVariantUpdateInput) SetText(value string) *ActionNotificationTemplateVariantUpdateInput {
	in.Text = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetTextNil(false)
	in._selectedParameters["Text"] = nil
	return in
}

// SetTextNil sets parameter Text to nil and selects it for sending
func (in *ActionNotificationTemplateVariantUpdateInput) SetTextNil(set bool) *ActionNotificationTemplateVariantUpdateInput {
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

// SelectParameters sets parameters from ActionNotificationTemplateVariantUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTemplateVariantUpdateInput) SelectParameters(params ...string) *ActionNotificationTemplateVariantUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNotificationTemplateVariantUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNotificationTemplateVariantUpdateInput) UnselectParameters(params ...string) *ActionNotificationTemplateVariantUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNotificationTemplateVariantUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTemplateVariantUpdateRequest is a type for the entire action request
type ActionNotificationTemplateVariantUpdateRequest struct {
	Variant map[string]interface{} "json:\"variant\""
	Meta    map[string]interface{} "json:\"_meta\""
}

// ActionNotificationTemplateVariantUpdateOutput is a type for action output parameters
type ActionNotificationTemplateVariantUpdateOutput struct {
	CreatedAt  string                    "json:\"created_at\""
	From       string                    "json:\"from\""
	Html       string                    "json:\"html\""
	Id         int64                     "json:\"id\""
	Language   *ActionLanguageShowOutput "json:\"language\""
	Protocol   string                    "json:\"protocol\""
	ReplyTo    string                    "json:\"reply_to\""
	ReturnPath string                    "json:\"return_path\""
	Subject    string                    "json:\"subject\""
	Text       string                    "json:\"text\""
	UpdatedAt  string                    "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionNotificationTemplateVariantUpdateResponse struct {
	Action *ActionNotificationTemplateVariantUpdate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Variant *ActionNotificationTemplateVariantUpdateOutput "json:\"variant\""
	}

	// Action output without the namespace
	Output *ActionNotificationTemplateVariantUpdateOutput
}

// Prepare the action for invocation
func (action *ActionNotificationTemplateVariantUpdate) Prepare() *ActionNotificationTemplateVariantUpdateInvocation {
	return &ActionNotificationTemplateVariantUpdateInvocation{
		Action: action,
		Path:   "/v7.0/notification_templates/{notification_template_id}/variants/{variant_id}",
	}
}

// ActionNotificationTemplateVariantUpdateInvocation is used to configure action for invocation
type ActionNotificationTemplateVariantUpdateInvocation struct {
	// Pointer to the action
	Action *ActionNotificationTemplateVariantUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNotificationTemplateVariantUpdateInput
	// Global meta input parameters
	MetaInput *ActionNotificationTemplateVariantUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationTemplateVariantUpdateInvocation) SetPathParamInt(param string, value int64) *ActionNotificationTemplateVariantUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationTemplateVariantUpdateInvocation) SetPathParamString(param string, value string) *ActionNotificationTemplateVariantUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNotificationTemplateVariantUpdateInvocation) NewInput() *ActionNotificationTemplateVariantUpdateInput {
	inv.Input = &ActionNotificationTemplateVariantUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNotificationTemplateVariantUpdateInvocation) SetInput(input *ActionNotificationTemplateVariantUpdateInput) *ActionNotificationTemplateVariantUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNotificationTemplateVariantUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNotificationTemplateVariantUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationTemplateVariantUpdateInvocation) NewMetaInput() *ActionNotificationTemplateVariantUpdateMetaGlobalInput {
	inv.MetaInput = &ActionNotificationTemplateVariantUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationTemplateVariantUpdateInvocation) SetMetaInput(input *ActionNotificationTemplateVariantUpdateMetaGlobalInput) *ActionNotificationTemplateVariantUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationTemplateVariantUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationTemplateVariantUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationTemplateVariantUpdateInvocation) validate() error {
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
func (inv *ActionNotificationTemplateVariantUpdateInvocation) Call() (*ActionNotificationTemplateVariantUpdateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNotificationTemplateVariantUpdateInvocation) callAsBody() (*ActionNotificationTemplateVariantUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNotificationTemplateVariantUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Variant
	}
	return resp, err
}

func (inv *ActionNotificationTemplateVariantUpdateInvocation) makeAllInputParams() *ActionNotificationTemplateVariantUpdateRequest {
	return &ActionNotificationTemplateVariantUpdateRequest{
		Variant: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionNotificationTemplateVariantUpdateInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionNotificationTemplateVariantUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
