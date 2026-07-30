package client

import (
	"net/url"
	"strings"
)

// ActionEventRouteMatcherCreate is a type for action Event_route.Matcher#Create
type ActionEventRouteMatcherCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionEventRouteMatcherCreate(client *Client) *ActionEventRouteMatcherCreate {
	return &ActionEventRouteMatcherCreate{
		Client: client,
	}
}

// ActionEventRouteMatcherCreateMetaGlobalInput is a type for action global meta input parameters
type ActionEventRouteMatcherCreateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventRouteMatcherCreateMetaGlobalInput) SetIncludes(value string) *ActionEventRouteMatcherCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventRouteMatcherCreateMetaGlobalInput) SetNo(value bool) *ActionEventRouteMatcherCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteMatcherCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteMatcherCreateMetaGlobalInput) SelectParameters(params ...string) *ActionEventRouteMatcherCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventRouteMatcherCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteMatcherCreateInput is a type for action input parameters
type ActionEventRouteMatcherCreateInput struct {
	Field    string "json:\"field\""
	Operator string "json:\"operator\""
	Value    string "json:\"value\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetField sets parameter Field to value and selects it for sending
func (in *ActionEventRouteMatcherCreateInput) SetField(value string) *ActionEventRouteMatcherCreateInput {
	in.Field = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Field"] = nil
	return in
}

// SetOperator sets parameter Operator to value and selects it for sending
func (in *ActionEventRouteMatcherCreateInput) SetOperator(value string) *ActionEventRouteMatcherCreateInput {
	in.Operator = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Operator"] = nil
	return in
}

// SetValue sets parameter Value to value and selects it for sending
func (in *ActionEventRouteMatcherCreateInput) SetValue(value string) *ActionEventRouteMatcherCreateInput {
	in.Value = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Value"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteMatcherCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteMatcherCreateInput) SelectParameters(params ...string) *ActionEventRouteMatcherCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionEventRouteMatcherCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionEventRouteMatcherCreateInput) UnselectParameters(params ...string) *ActionEventRouteMatcherCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionEventRouteMatcherCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteMatcherCreateRequest is a type for the entire action request
type ActionEventRouteMatcherCreateRequest struct {
	Matcher map[string]interface{} "json:\"matcher\""
	Meta    map[string]interface{} "json:\"_meta\""
}

// ActionEventRouteMatcherCreateOutput is a type for action output parameters
type ActionEventRouteMatcherCreateOutput struct {
	CreatedAt string "json:\"created_at\""
	Field     string "json:\"field\""
	FieldType string "json:\"field_type\""
	Id        int64  "json:\"id\""
	Operator  string "json:\"operator\""
	Summary   string "json:\"summary\""
	UpdatedAt string "json:\"updated_at\""
	Value     string "json:\"value\""
}

// Type for action response, including envelope
type ActionEventRouteMatcherCreateResponse struct {
	Action *ActionEventRouteMatcherCreate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Matcher *ActionEventRouteMatcherCreateOutput "json:\"matcher\""
	}

	// Action output without the namespace
	Output *ActionEventRouteMatcherCreateOutput
}

// Prepare the action for invocation
func (action *ActionEventRouteMatcherCreate) Prepare() *ActionEventRouteMatcherCreateInvocation {
	return &ActionEventRouteMatcherCreateInvocation{
		Action: action,
		Path:   "/v7.0/event_routes/{event_route_id}/matcher",
	}
}

// ActionEventRouteMatcherCreateInvocation is used to configure action for invocation
type ActionEventRouteMatcherCreateInvocation struct {
	// Pointer to the action
	Action *ActionEventRouteMatcherCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEventRouteMatcherCreateInput
	// Global meta input parameters
	MetaInput *ActionEventRouteMatcherCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventRouteMatcherCreateInvocation) SetPathParamInt(param string, value int64) *ActionEventRouteMatcherCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventRouteMatcherCreateInvocation) SetPathParamString(param string, value string) *ActionEventRouteMatcherCreateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEventRouteMatcherCreateInvocation) NewInput() *ActionEventRouteMatcherCreateInput {
	inv.Input = &ActionEventRouteMatcherCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEventRouteMatcherCreateInvocation) SetInput(input *ActionEventRouteMatcherCreateInput) *ActionEventRouteMatcherCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEventRouteMatcherCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionEventRouteMatcherCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventRouteMatcherCreateInvocation) NewMetaInput() *ActionEventRouteMatcherCreateMetaGlobalInput {
	inv.MetaInput = &ActionEventRouteMatcherCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventRouteMatcherCreateInvocation) SetMetaInput(input *ActionEventRouteMatcherCreateMetaGlobalInput) *ActionEventRouteMatcherCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventRouteMatcherCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventRouteMatcherCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventRouteMatcherCreateInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEventRouteMatcherCreateInvocation) Call() (*ActionEventRouteMatcherCreateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionEventRouteMatcherCreateInvocation) callAsBody() (*ActionEventRouteMatcherCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionEventRouteMatcherCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Matcher
	}
	return resp, err
}

func (inv *ActionEventRouteMatcherCreateInvocation) makeAllInputParams() *ActionEventRouteMatcherCreateRequest {
	return &ActionEventRouteMatcherCreateRequest{
		Matcher: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionEventRouteMatcherCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Field") {
			ret["field"] = inv.Input.Field
		}
		if inv.IsParameterSelected("Operator") {
			ret["operator"] = inv.Input.Operator
		}
		if inv.IsParameterSelected("Value") {
			ret["value"] = inv.Input.Value
		}
	}

	return ret
}

func (inv *ActionEventRouteMatcherCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
