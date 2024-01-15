package client

import (
	"strings"
)

// ActionMailboxHandlerCreate is a type for action Mailbox.Handler#Create
type ActionMailboxHandlerCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionMailboxHandlerCreate(client *Client) *ActionMailboxHandlerCreate {
	return &ActionMailboxHandlerCreate{
		Client: client,
	}
}

// ActionMailboxHandlerCreateMetaGlobalInput is a type for action global meta input parameters
type ActionMailboxHandlerCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailboxHandlerCreateMetaGlobalInput) SetIncludes(value string) *ActionMailboxHandlerCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailboxHandlerCreateMetaGlobalInput) SetNo(value bool) *ActionMailboxHandlerCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailboxHandlerCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailboxHandlerCreateMetaGlobalInput) SelectParameters(params ...string) *ActionMailboxHandlerCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailboxHandlerCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailboxHandlerCreateInput is a type for action input parameters
type ActionMailboxHandlerCreateInput struct {
	ClassName string `json:"class_name"`
	Continue  bool   `json:"continue"`
	Order     int64  `json:"order"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetClassName sets parameter ClassName to value and selects it for sending
func (in *ActionMailboxHandlerCreateInput) SetClassName(value string) *ActionMailboxHandlerCreateInput {
	in.ClassName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClassName"] = nil
	return in
}

// SetContinue sets parameter Continue to value and selects it for sending
func (in *ActionMailboxHandlerCreateInput) SetContinue(value bool) *ActionMailboxHandlerCreateInput {
	in.Continue = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Continue"] = nil
	return in
}

// SetOrder sets parameter Order to value and selects it for sending
func (in *ActionMailboxHandlerCreateInput) SetOrder(value int64) *ActionMailboxHandlerCreateInput {
	in.Order = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Order"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailboxHandlerCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailboxHandlerCreateInput) SelectParameters(params ...string) *ActionMailboxHandlerCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionMailboxHandlerCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionMailboxHandlerCreateInput) UnselectParameters(params ...string) *ActionMailboxHandlerCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionMailboxHandlerCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailboxHandlerCreateRequest is a type for the entire action request
type ActionMailboxHandlerCreateRequest struct {
	Handler map[string]interface{} `json:"handler"`
	Meta    map[string]interface{} `json:"_meta"`
}

// ActionMailboxHandlerCreateOutput is a type for action output parameters
type ActionMailboxHandlerCreateOutput struct {
	ClassName string `json:"class_name"`
	Continue  bool   `json:"continue"`
	CreatedAt string `json:"created_at"`
	Id        int64  `json:"id"`
	Order     int64  `json:"order"`
	UpdatedAt string `json:"updated_at"`
}

// Type for action response, including envelope
type ActionMailboxHandlerCreateResponse struct {
	Action *ActionMailboxHandlerCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Handler *ActionMailboxHandlerCreateOutput `json:"handler"`
	}

	// Action output without the namespace
	Output *ActionMailboxHandlerCreateOutput
}

// Prepare the action for invocation
func (action *ActionMailboxHandlerCreate) Prepare() *ActionMailboxHandlerCreateInvocation {
	return &ActionMailboxHandlerCreateInvocation{
		Action: action,
		Path:   "/v6.0/mailboxes/{mailbox_id}/handler",
	}
}

// ActionMailboxHandlerCreateInvocation is used to configure action for invocation
type ActionMailboxHandlerCreateInvocation struct {
	// Pointer to the action
	Action *ActionMailboxHandlerCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMailboxHandlerCreateInput
	// Global meta input parameters
	MetaInput *ActionMailboxHandlerCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailboxHandlerCreateInvocation) SetPathParamInt(param string, value int64) *ActionMailboxHandlerCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailboxHandlerCreateInvocation) SetPathParamString(param string, value string) *ActionMailboxHandlerCreateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMailboxHandlerCreateInvocation) NewInput() *ActionMailboxHandlerCreateInput {
	inv.Input = &ActionMailboxHandlerCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMailboxHandlerCreateInvocation) SetInput(input *ActionMailboxHandlerCreateInput) *ActionMailboxHandlerCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMailboxHandlerCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionMailboxHandlerCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailboxHandlerCreateInvocation) NewMetaInput() *ActionMailboxHandlerCreateMetaGlobalInput {
	inv.MetaInput = &ActionMailboxHandlerCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailboxHandlerCreateInvocation) SetMetaInput(input *ActionMailboxHandlerCreateMetaGlobalInput) *ActionMailboxHandlerCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailboxHandlerCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMailboxHandlerCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailboxHandlerCreateInvocation) Call() (*ActionMailboxHandlerCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionMailboxHandlerCreateInvocation) callAsBody() (*ActionMailboxHandlerCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMailboxHandlerCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Handler
	}
	return resp, err
}

func (inv *ActionMailboxHandlerCreateInvocation) makeAllInputParams() *ActionMailboxHandlerCreateRequest {
	return &ActionMailboxHandlerCreateRequest{
		Handler: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionMailboxHandlerCreateInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionMailboxHandlerCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
