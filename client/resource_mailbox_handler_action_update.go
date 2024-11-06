package client

import (
	"strings"
)

// ActionMailboxHandlerUpdate is a type for action Mailbox.Handler#Update
type ActionMailboxHandlerUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionMailboxHandlerUpdate(client *Client) *ActionMailboxHandlerUpdate {
	return &ActionMailboxHandlerUpdate{
		Client: client,
	}
}

// ActionMailboxHandlerUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionMailboxHandlerUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailboxHandlerUpdateMetaGlobalInput) SetIncludes(value string) *ActionMailboxHandlerUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailboxHandlerUpdateMetaGlobalInput) SetNo(value bool) *ActionMailboxHandlerUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailboxHandlerUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailboxHandlerUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionMailboxHandlerUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailboxHandlerUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailboxHandlerUpdateInput is a type for action input parameters
type ActionMailboxHandlerUpdateInput struct {
	ClassName string `json:"class_name"`
	Continue  bool   `json:"continue"`
	Order     int64  `json:"order"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetClassName sets parameter ClassName to value and selects it for sending
func (in *ActionMailboxHandlerUpdateInput) SetClassName(value string) *ActionMailboxHandlerUpdateInput {
	in.ClassName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClassName"] = nil
	return in
}

// SetContinue sets parameter Continue to value and selects it for sending
func (in *ActionMailboxHandlerUpdateInput) SetContinue(value bool) *ActionMailboxHandlerUpdateInput {
	in.Continue = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Continue"] = nil
	return in
}

// SetOrder sets parameter Order to value and selects it for sending
func (in *ActionMailboxHandlerUpdateInput) SetOrder(value int64) *ActionMailboxHandlerUpdateInput {
	in.Order = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Order"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailboxHandlerUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailboxHandlerUpdateInput) SelectParameters(params ...string) *ActionMailboxHandlerUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionMailboxHandlerUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionMailboxHandlerUpdateInput) UnselectParameters(params ...string) *ActionMailboxHandlerUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionMailboxHandlerUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailboxHandlerUpdateRequest is a type for the entire action request
type ActionMailboxHandlerUpdateRequest struct {
	Handler map[string]interface{} `json:"handler"`
	Meta    map[string]interface{} `json:"_meta"`
}

// ActionMailboxHandlerUpdateOutput is a type for action output parameters
type ActionMailboxHandlerUpdateOutput struct {
	ClassName string `json:"class_name"`
	Continue  bool   `json:"continue"`
	CreatedAt string `json:"created_at"`
	Id        int64  `json:"id"`
	Order     int64  `json:"order"`
	UpdatedAt string `json:"updated_at"`
}

// Type for action response, including envelope
type ActionMailboxHandlerUpdateResponse struct {
	Action *ActionMailboxHandlerUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Handler *ActionMailboxHandlerUpdateOutput `json:"handler"`
	}

	// Action output without the namespace
	Output *ActionMailboxHandlerUpdateOutput
}

// Prepare the action for invocation
func (action *ActionMailboxHandlerUpdate) Prepare() *ActionMailboxHandlerUpdateInvocation {
	return &ActionMailboxHandlerUpdateInvocation{
		Action: action,
		Path:   "/v7.0/mailboxes/{mailbox_id}/handler/{handler_id}",
	}
}

// ActionMailboxHandlerUpdateInvocation is used to configure action for invocation
type ActionMailboxHandlerUpdateInvocation struct {
	// Pointer to the action
	Action *ActionMailboxHandlerUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMailboxHandlerUpdateInput
	// Global meta input parameters
	MetaInput *ActionMailboxHandlerUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailboxHandlerUpdateInvocation) SetPathParamInt(param string, value int64) *ActionMailboxHandlerUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailboxHandlerUpdateInvocation) SetPathParamString(param string, value string) *ActionMailboxHandlerUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMailboxHandlerUpdateInvocation) NewInput() *ActionMailboxHandlerUpdateInput {
	inv.Input = &ActionMailboxHandlerUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMailboxHandlerUpdateInvocation) SetInput(input *ActionMailboxHandlerUpdateInput) *ActionMailboxHandlerUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMailboxHandlerUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionMailboxHandlerUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailboxHandlerUpdateInvocation) NewMetaInput() *ActionMailboxHandlerUpdateMetaGlobalInput {
	inv.MetaInput = &ActionMailboxHandlerUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailboxHandlerUpdateInvocation) SetMetaInput(input *ActionMailboxHandlerUpdateMetaGlobalInput) *ActionMailboxHandlerUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailboxHandlerUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMailboxHandlerUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailboxHandlerUpdateInvocation) Call() (*ActionMailboxHandlerUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionMailboxHandlerUpdateInvocation) callAsBody() (*ActionMailboxHandlerUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMailboxHandlerUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Handler
	}
	return resp, err
}

func (inv *ActionMailboxHandlerUpdateInvocation) makeAllInputParams() *ActionMailboxHandlerUpdateRequest {
	return &ActionMailboxHandlerUpdateRequest{
		Handler: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionMailboxHandlerUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("ClassName") {
			ret["class_name"] = inv.Input.ClassName
		}
		if inv.IsParameterSelected("Continue") {
			ret["continue"] = inv.Input.Continue
		}
		if inv.IsParameterSelected("Order") {
			ret["order"] = inv.Input.Order
		}
	}

	return ret
}

func (inv *ActionMailboxHandlerUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
