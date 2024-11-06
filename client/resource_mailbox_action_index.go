package client

import ()

// ActionMailboxIndex is a type for action Mailbox#Index
type ActionMailboxIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionMailboxIndex(client *Client) *ActionMailboxIndex {
	return &ActionMailboxIndex{
		Client: client,
	}
}

// ActionMailboxIndexMetaGlobalInput is a type for action global meta input parameters
type ActionMailboxIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionMailboxIndexMetaGlobalInput) SetCount(value bool) *ActionMailboxIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailboxIndexMetaGlobalInput) SetIncludes(value string) *ActionMailboxIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailboxIndexMetaGlobalInput) SetNo(value bool) *ActionMailboxIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailboxIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailboxIndexMetaGlobalInput) SelectParameters(params ...string) *ActionMailboxIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailboxIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailboxIndexInput is a type for action input parameters
type ActionMailboxIndexInput struct {
	FromId int64 `json:"from_id"`
	Limit  int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionMailboxIndexInput) SetFromId(value int64) *ActionMailboxIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionMailboxIndexInput) SetLimit(value int64) *ActionMailboxIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailboxIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailboxIndexInput) SelectParameters(params ...string) *ActionMailboxIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionMailboxIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionMailboxIndexInput) UnselectParameters(params ...string) *ActionMailboxIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionMailboxIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailboxIndexOutput is a type for action output parameters
type ActionMailboxIndexOutput struct {
	CreatedAt string `json:"created_at"`
	EnableSsl bool   `json:"enable_ssl"`
	Id        int64  `json:"id"`
	Label     string `json:"label"`
	Port      int64  `json:"port"`
	Server    string `json:"server"`
	UpdatedAt string `json:"updated_at"`
	User      string `json:"user"`
}

// Type for action response, including envelope
type ActionMailboxIndexResponse struct {
	Action *ActionMailboxIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Mailboxes []*ActionMailboxIndexOutput `json:"mailboxes"`
	}

	// Action output without the namespace
	Output []*ActionMailboxIndexOutput
}

// Prepare the action for invocation
func (action *ActionMailboxIndex) Prepare() *ActionMailboxIndexInvocation {
	return &ActionMailboxIndexInvocation{
		Action: action,
		Path:   "/v7.0/mailboxes",
	}
}

// ActionMailboxIndexInvocation is used to configure action for invocation
type ActionMailboxIndexInvocation struct {
	// Pointer to the action
	Action *ActionMailboxIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMailboxIndexInput
	// Global meta input parameters
	MetaInput *ActionMailboxIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMailboxIndexInvocation) NewInput() *ActionMailboxIndexInput {
	inv.Input = &ActionMailboxIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMailboxIndexInvocation) SetInput(input *ActionMailboxIndexInput) *ActionMailboxIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMailboxIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionMailboxIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailboxIndexInvocation) NewMetaInput() *ActionMailboxIndexMetaGlobalInput {
	inv.MetaInput = &ActionMailboxIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailboxIndexInvocation) SetMetaInput(input *ActionMailboxIndexMetaGlobalInput) *ActionMailboxIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailboxIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMailboxIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailboxIndexInvocation) Call() (*ActionMailboxIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMailboxIndexInvocation) callAsQuery() (*ActionMailboxIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMailboxIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Mailboxes
	}
	return resp, err
}

func (inv *ActionMailboxIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["mailbox[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["mailbox[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionMailboxIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
