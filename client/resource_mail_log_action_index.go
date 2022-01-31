package client

import ()

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
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
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

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailLogIndexMetaGlobalInput) SetNo(value bool) *ActionMailLogIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
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
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
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

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionMailLogIndexInput) SetOffset(value int64) *ActionMailLogIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
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

// UnselectParameters unsets parameters from ActionMailLogIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionMailLogIndexInput) UnselectParameters(params ...string) *ActionMailLogIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
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
		Path:   "/v6.0/mail_logs",
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

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMailLogIndexInvocation) NewInput() *ActionMailLogIndexInput {
	inv.Input = &ActionMailLogIndexInput{}
	return inv.Input
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

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionMailLogIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailLogIndexInvocation) NewMetaInput() *ActionMailLogIndexMetaGlobalInput {
	inv.MetaInput = &ActionMailLogIndexMetaGlobalInput{}
	return inv.MetaInput
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

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMailLogIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
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
		if inv.IsParameterSelected("Limit") {
			ret["mail_log[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["mail_log[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
	}
}

func (inv *ActionMailLogIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
