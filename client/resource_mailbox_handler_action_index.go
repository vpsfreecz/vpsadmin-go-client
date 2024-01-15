package client

import (
	"strings"
)

// ActionMailboxHandlerIndex is a type for action Mailbox.Handler#Index
type ActionMailboxHandlerIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionMailboxHandlerIndex(client *Client) *ActionMailboxHandlerIndex {
	return &ActionMailboxHandlerIndex{
		Client: client,
	}
}

// ActionMailboxHandlerIndexMetaGlobalInput is a type for action global meta input parameters
type ActionMailboxHandlerIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionMailboxHandlerIndexMetaGlobalInput) SetCount(value bool) *ActionMailboxHandlerIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailboxHandlerIndexMetaGlobalInput) SetIncludes(value string) *ActionMailboxHandlerIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailboxHandlerIndexMetaGlobalInput) SetNo(value bool) *ActionMailboxHandlerIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailboxHandlerIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailboxHandlerIndexMetaGlobalInput) SelectParameters(params ...string) *ActionMailboxHandlerIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailboxHandlerIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailboxHandlerIndexInput is a type for action input parameters
type ActionMailboxHandlerIndexInput struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionMailboxHandlerIndexInput) SetLimit(value int64) *ActionMailboxHandlerIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionMailboxHandlerIndexInput) SetOffset(value int64) *ActionMailboxHandlerIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailboxHandlerIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailboxHandlerIndexInput) SelectParameters(params ...string) *ActionMailboxHandlerIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionMailboxHandlerIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionMailboxHandlerIndexInput) UnselectParameters(params ...string) *ActionMailboxHandlerIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionMailboxHandlerIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailboxHandlerIndexOutput is a type for action output parameters
type ActionMailboxHandlerIndexOutput struct {
	ClassName string `json:"class_name"`
	Continue  bool   `json:"continue"`
	CreatedAt string `json:"created_at"`
	Id        int64  `json:"id"`
	Order     int64  `json:"order"`
	UpdatedAt string `json:"updated_at"`
}

// Type for action response, including envelope
type ActionMailboxHandlerIndexResponse struct {
	Action *ActionMailboxHandlerIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Handlers []*ActionMailboxHandlerIndexOutput `json:"handlers"`
	}

	// Action output without the namespace
	Output []*ActionMailboxHandlerIndexOutput
}

// Prepare the action for invocation
func (action *ActionMailboxHandlerIndex) Prepare() *ActionMailboxHandlerIndexInvocation {
	return &ActionMailboxHandlerIndexInvocation{
		Action: action,
		Path:   "/v6.0/mailboxes/{mailbox_id}/handler",
	}
}

// ActionMailboxHandlerIndexInvocation is used to configure action for invocation
type ActionMailboxHandlerIndexInvocation struct {
	// Pointer to the action
	Action *ActionMailboxHandlerIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMailboxHandlerIndexInput
	// Global meta input parameters
	MetaInput *ActionMailboxHandlerIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailboxHandlerIndexInvocation) SetPathParamInt(param string, value int64) *ActionMailboxHandlerIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailboxHandlerIndexInvocation) SetPathParamString(param string, value string) *ActionMailboxHandlerIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMailboxHandlerIndexInvocation) NewInput() *ActionMailboxHandlerIndexInput {
	inv.Input = &ActionMailboxHandlerIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMailboxHandlerIndexInvocation) SetInput(input *ActionMailboxHandlerIndexInput) *ActionMailboxHandlerIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMailboxHandlerIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionMailboxHandlerIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailboxHandlerIndexInvocation) NewMetaInput() *ActionMailboxHandlerIndexMetaGlobalInput {
	inv.MetaInput = &ActionMailboxHandlerIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailboxHandlerIndexInvocation) SetMetaInput(input *ActionMailboxHandlerIndexMetaGlobalInput) *ActionMailboxHandlerIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailboxHandlerIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMailboxHandlerIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailboxHandlerIndexInvocation) Call() (*ActionMailboxHandlerIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMailboxHandlerIndexInvocation) callAsQuery() (*ActionMailboxHandlerIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMailboxHandlerIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Handlers
	}
	return resp, err
}

func (inv *ActionMailboxHandlerIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["handler[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["handler[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
	}
}

func (inv *ActionMailboxHandlerIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
