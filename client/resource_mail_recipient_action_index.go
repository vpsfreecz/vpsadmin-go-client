package client

import ()

// ActionMailRecipientIndex is a type for action Mail_recipient#Index
type ActionMailRecipientIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionMailRecipientIndex(client *Client) *ActionMailRecipientIndex {
	return &ActionMailRecipientIndex{
		Client: client,
	}
}

// ActionMailRecipientIndexMetaGlobalInput is a type for action global meta input parameters
type ActionMailRecipientIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionMailRecipientIndexMetaGlobalInput) SetCount(value bool) *ActionMailRecipientIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailRecipientIndexMetaGlobalInput) SetIncludes(value string) *ActionMailRecipientIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailRecipientIndexMetaGlobalInput) SetNo(value bool) *ActionMailRecipientIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailRecipientIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailRecipientIndexMetaGlobalInput) SelectParameters(params ...string) *ActionMailRecipientIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailRecipientIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailRecipientIndexInput is a type for action input parameters
type ActionMailRecipientIndexInput struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionMailRecipientIndexInput) SetLimit(value int64) *ActionMailRecipientIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionMailRecipientIndexInput) SetOffset(value int64) *ActionMailRecipientIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailRecipientIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailRecipientIndexInput) SelectParameters(params ...string) *ActionMailRecipientIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailRecipientIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailRecipientIndexOutput is a type for action output parameters
type ActionMailRecipientIndexOutput struct {
	Bcc   string `json:"bcc"`
	Cc    string `json:"cc"`
	Id    int64  `json:"id"`
	Label string `json:"label"`
	To    string `json:"to"`
}

// Type for action response, including envelope
type ActionMailRecipientIndexResponse struct {
	Action *ActionMailRecipientIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MailRecipients []*ActionMailRecipientIndexOutput `json:"mail_recipients"`
	}

	// Action output without the namespace
	Output []*ActionMailRecipientIndexOutput
}

// Prepare the action for invocation
func (action *ActionMailRecipientIndex) Prepare() *ActionMailRecipientIndexInvocation {
	return &ActionMailRecipientIndexInvocation{
		Action: action,
		Path:   "/v6.0/mail_recipients",
	}
}

// ActionMailRecipientIndexInvocation is used to configure action for invocation
type ActionMailRecipientIndexInvocation struct {
	// Pointer to the action
	Action *ActionMailRecipientIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMailRecipientIndexInput
	// Global meta input parameters
	MetaInput *ActionMailRecipientIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMailRecipientIndexInvocation) NewInput() *ActionMailRecipientIndexInput {
	inv.Input = &ActionMailRecipientIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMailRecipientIndexInvocation) SetInput(input *ActionMailRecipientIndexInput) *ActionMailRecipientIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMailRecipientIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailRecipientIndexInvocation) NewMetaInput() *ActionMailRecipientIndexMetaGlobalInput {
	inv.MetaInput = &ActionMailRecipientIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailRecipientIndexInvocation) SetMetaInput(input *ActionMailRecipientIndexMetaGlobalInput) *ActionMailRecipientIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailRecipientIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailRecipientIndexInvocation) Call() (*ActionMailRecipientIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMailRecipientIndexInvocation) callAsQuery() (*ActionMailRecipientIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMailRecipientIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MailRecipients
	}
	return resp, err
}

func (inv *ActionMailRecipientIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["mail_recipient[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["mail_recipient[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
	}
}

func (inv *ActionMailRecipientIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
