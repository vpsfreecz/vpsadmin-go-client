package client

import (
)

// ActionMailLogIndex is a type for action Mail_log#Index
type ActionMailLogIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionMailLogIndex(client *Client) *ActionMailLogIndex {
	return &ActionMailLogIndex{
		Client: client,
	}
}

// ActionMailLogIndexMetaGlobalInput is a type for action global meta input parameters
type ActionMailLogIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailLogIndexMetaGlobalInput) SetNo(value bool) *ActionMailLogIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionMailLogIndexMetaGlobalInput) SetCount(value bool) *ActionMailLogIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailLogIndexMetaGlobalInput) SetIncludes(value string) *ActionMailLogIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailLogIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailLogIndexMetaGlobalInput) SelectParameters(params ...string) *ActionMailLogIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailLogIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailLogIndexInput is a type for action input parameters
type ActionMailLogIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionMailLogIndexInput) SetOffset(value int64) *ActionMailLogIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionMailLogIndexInput) SetLimit(value int64) *ActionMailLogIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailLogIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailLogIndexInput) SelectParameters(params ...string) *ActionMailLogIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailLogIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionMailLogIndexOutput is a type for action output parameters
type ActionMailLogIndexOutput struct {
	Id int64 `json:"id"`
	User *ActionUserShowOutput `json:"user"`
	To string `json:"to"`
	Cc string `json:"cc"`
	Bcc string `json:"bcc"`
	From string `json:"from"`
	ReplyTo string `json:"reply_to"`
	ReturnPath string `json:"return_path"`
	MessageId string `json:"message_id"`
	InReplyTo string `json:"in_reply_to"`
	References string `json:"references"`
	Subject string `json:"subject"`
	TextPlain string `json:"text_plain"`
	TextHtml string `json:"text_html"`
	MailTemplate *ActionMailTemplateShowOutput `json:"mail_template"`
	CreatedAt string `json:"created_at"`
}


// Type for action response, including envelope
type ActionMailLogIndexResponse struct {
	Action *ActionMailLogIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MailLogs []*ActionMailLogIndexOutput `json:"mail_logs"`
	}

	// Action output without the namespace
	Output []*ActionMailLogIndexOutput
}


// Prepare the action for invocation
func (action *ActionMailLogIndex) Prepare() *ActionMailLogIndexInvocation {
	return &ActionMailLogIndexInvocation{
		Action: action,
		Path: "/v5.0/mail_logs",
	}
}

// ActionMailLogIndexInvocation is used to configure action for invocation
type ActionMailLogIndexInvocation struct {
	// Pointer to the action
	Action *ActionMailLogIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMailLogIndexInput
	// Global meta input parameters
	MetaInput *ActionMailLogIndexMetaGlobalInput
}


// SetInput provides input parameters to send to the API
func (inv *ActionMailLogIndexInvocation) SetInput(input *ActionMailLogIndexInput) *ActionMailLogIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMailLogIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailLogIndexInvocation) SetMetaInput(input *ActionMailLogIndexMetaGlobalInput) *ActionMailLogIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailLogIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailLogIndexInvocation) Call() (*ActionMailLogIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMailLogIndexInvocation) callAsQuery() (*ActionMailLogIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMailLogIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MailLogs
	}
	return resp, err
}



func (inv *ActionMailLogIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["mail_log[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["mail_log[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionMailLogIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

