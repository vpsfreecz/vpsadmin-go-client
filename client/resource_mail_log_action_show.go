package client

import (
	"strings"
)

// ActionMailLogShow is a type for action Mail_log#Show
type ActionMailLogShow struct {
	// Pointer to client
	Client *Client
}

func NewActionMailLogShow(client *Client) *ActionMailLogShow {
	return &ActionMailLogShow{
		Client: client,
	}
}

// ActionMailLogShowMetaGlobalInput is a type for action global meta input parameters
type ActionMailLogShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailLogShowMetaGlobalInput) SetIncludes(value string) *ActionMailLogShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailLogShowMetaGlobalInput) SetNo(value bool) *ActionMailLogShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailLogShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailLogShowMetaGlobalInput) SelectParameters(params ...string) *ActionMailLogShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailLogShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailLogShowOutput is a type for action output parameters
type ActionMailLogShowOutput struct {
	Bcc          string                        `json:"bcc"`
	Cc           string                        `json:"cc"`
	CreatedAt    string                        `json:"created_at"`
	From         string                        `json:"from"`
	Id           int64                         `json:"id"`
	InReplyTo    string                        `json:"in_reply_to"`
	MailTemplate *ActionMailTemplateShowOutput `json:"mail_template"`
	MessageId    string                        `json:"message_id"`
	References   string                        `json:"references"`
	ReplyTo      string                        `json:"reply_to"`
	ReturnPath   string                        `json:"return_path"`
	Subject      string                        `json:"subject"`
	TextHtml     string                        `json:"text_html"`
	TextPlain    string                        `json:"text_plain"`
	To           string                        `json:"to"`
	User         *ActionUserShowOutput         `json:"user"`
}

// Type for action response, including envelope
type ActionMailLogShowResponse struct {
	Action *ActionMailLogShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MailLog *ActionMailLogShowOutput `json:"mail_log"`
	}

	// Action output without the namespace
	Output *ActionMailLogShowOutput
}

// Prepare the action for invocation
func (action *ActionMailLogShow) Prepare() *ActionMailLogShowInvocation {
	return &ActionMailLogShowInvocation{
		Action: action,
		Path:   "/v7.0/mail_logs/{mail_log_id}",
	}
}

// ActionMailLogShowInvocation is used to configure action for invocation
type ActionMailLogShowInvocation struct {
	// Pointer to the action
	Action *ActionMailLogShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMailLogShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailLogShowInvocation) SetPathParamInt(param string, value int64) *ActionMailLogShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailLogShowInvocation) SetPathParamString(param string, value string) *ActionMailLogShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailLogShowInvocation) NewMetaInput() *ActionMailLogShowMetaGlobalInput {
	inv.MetaInput = &ActionMailLogShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailLogShowInvocation) SetMetaInput(input *ActionMailLogShowMetaGlobalInput) *ActionMailLogShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailLogShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMailLogShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailLogShowInvocation) Call() (*ActionMailLogShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMailLogShowInvocation) callAsQuery() (*ActionMailLogShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMailLogShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MailLog
	}
	return resp, err
}

func (inv *ActionMailLogShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
