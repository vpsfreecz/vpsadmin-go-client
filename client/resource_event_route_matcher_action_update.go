package client

import (
	"net/url"
	"strings"
)

// ActionEventRouteMatcherUpdate is a type for action Event_route.Matcher#Update
type ActionEventRouteMatcherUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionEventRouteMatcherUpdate(client *Client) *ActionEventRouteMatcherUpdate {
	return &ActionEventRouteMatcherUpdate{
		Client: client,
	}
}

// ActionEventRouteMatcherUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionEventRouteMatcherUpdateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventRouteMatcherUpdateMetaGlobalInput) SetIncludes(value string) *ActionEventRouteMatcherUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventRouteMatcherUpdateMetaGlobalInput) SetNo(value bool) *ActionEventRouteMatcherUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteMatcherUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteMatcherUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionEventRouteMatcherUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventRouteMatcherUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteMatcherUpdateInput is a type for action input parameters
type ActionEventRouteMatcherUpdateInput struct {
	Field    string "json:\"field\""
	Operator string "json:\"operator\""
	Value    string "json:\"value\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetField sets parameter Field to value and selects it for sending
func (in *ActionEventRouteMatcherUpdateInput) SetField(value string) *ActionEventRouteMatcherUpdateInput {
	in.Field = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Field"] = nil
	return in
}

// SetOperator sets parameter Operator to value and selects it for sending
func (in *ActionEventRouteMatcherUpdateInput) SetOperator(value string) *ActionEventRouteMatcherUpdateInput {
	in.Operator = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Operator"] = nil
	return in
}

// SetValue sets parameter Value to value and selects it for sending
func (in *ActionEventRouteMatcherUpdateInput) SetValue(value string) *ActionEventRouteMatcherUpdateInput {
	in.Value = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Value"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteMatcherUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteMatcherUpdateInput) SelectParameters(params ...string) *ActionEventRouteMatcherUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionEventRouteMatcherUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionEventRouteMatcherUpdateInput) UnselectParameters(params ...string) *ActionEventRouteMatcherUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionEventRouteMatcherUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteMatcherUpdateRequest is a type for the entire action request
type ActionEventRouteMatcherUpdateRequest struct {
	Matcher map[string]interface{} "json:\"matcher\""
	Meta    map[string]interface{} "json:\"_meta\""
}

// ActionEventRouteMatcherUpdateOutput is a type for action output parameters
type ActionEventRouteMatcherUpdateOutput struct {
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
type ActionEventRouteMatcherUpdateResponse struct {
	Action *ActionEventRouteMatcherUpdate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Matcher *ActionEventRouteMatcherUpdateOutput "json:\"matcher\""
	}

	// Action output without the namespace
	Output *ActionEventRouteMatcherUpdateOutput
}

// Prepare the action for invocation
func (action *ActionEventRouteMatcherUpdate) Prepare() *ActionEventRouteMatcherUpdateInvocation {
	return &ActionEventRouteMatcherUpdateInvocation{
		Action: action,
		Path:   "/v7.0/event_routes/{event_route_id}/matcher/{matcher_id}",
	}
}

// ActionEventRouteMatcherUpdateInvocation is used to configure action for invocation
type ActionEventRouteMatcherUpdateInvocation struct {
	// Pointer to the action
	Action *ActionEventRouteMatcherUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEventRouteMatcherUpdateInput
	// Global meta input parameters
	MetaInput *ActionEventRouteMatcherUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventRouteMatcherUpdateInvocation) SetPathParamInt(param string, value int64) *ActionEventRouteMatcherUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventRouteMatcherUpdateInvocation) SetPathParamString(param string, value string) *ActionEventRouteMatcherUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEventRouteMatcherUpdateInvocation) NewInput() *ActionEventRouteMatcherUpdateInput {
	inv.Input = &ActionEventRouteMatcherUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEventRouteMatcherUpdateInvocation) SetInput(input *ActionEventRouteMatcherUpdateInput) *ActionEventRouteMatcherUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEventRouteMatcherUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionEventRouteMatcherUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventRouteMatcherUpdateInvocation) NewMetaInput() *ActionEventRouteMatcherUpdateMetaGlobalInput {
	inv.MetaInput = &ActionEventRouteMatcherUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventRouteMatcherUpdateInvocation) SetMetaInput(input *ActionEventRouteMatcherUpdateMetaGlobalInput) *ActionEventRouteMatcherUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventRouteMatcherUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventRouteMatcherUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventRouteMatcherUpdateInvocation) validate() error {
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
func (inv *ActionEventRouteMatcherUpdateInvocation) Call() (*ActionEventRouteMatcherUpdateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionEventRouteMatcherUpdateInvocation) callAsBody() (*ActionEventRouteMatcherUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionEventRouteMatcherUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Matcher
	}
	return resp, err
}

func (inv *ActionEventRouteMatcherUpdateInvocation) makeAllInputParams() *ActionEventRouteMatcherUpdateRequest {
	return &ActionEventRouteMatcherUpdateRequest{
		Matcher: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionEventRouteMatcherUpdateInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionEventRouteMatcherUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
