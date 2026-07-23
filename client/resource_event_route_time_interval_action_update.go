package client

import (
	"net/url"
	"strings"
)

// ActionEventRouteTimeIntervalUpdate is a type for action Event_route.Time_interval#Update
type ActionEventRouteTimeIntervalUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionEventRouteTimeIntervalUpdate(client *Client) *ActionEventRouteTimeIntervalUpdate {
	return &ActionEventRouteTimeIntervalUpdate{
		Client: client,
	}
}

// ActionEventRouteTimeIntervalUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionEventRouteTimeIntervalUpdateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventRouteTimeIntervalUpdateMetaGlobalInput) SetIncludes(value string) *ActionEventRouteTimeIntervalUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventRouteTimeIntervalUpdateMetaGlobalInput) SetNo(value bool) *ActionEventRouteTimeIntervalUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteTimeIntervalUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteTimeIntervalUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionEventRouteTimeIntervalUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventRouteTimeIntervalUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteTimeIntervalUpdateInput is a type for action input parameters
type ActionEventRouteTimeIntervalUpdateInput struct {
	Mode string "json:\"mode\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetMode sets parameter Mode to value and selects it for sending
func (in *ActionEventRouteTimeIntervalUpdateInput) SetMode(value string) *ActionEventRouteTimeIntervalUpdateInput {
	in.Mode = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Mode"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteTimeIntervalUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteTimeIntervalUpdateInput) SelectParameters(params ...string) *ActionEventRouteTimeIntervalUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionEventRouteTimeIntervalUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionEventRouteTimeIntervalUpdateInput) UnselectParameters(params ...string) *ActionEventRouteTimeIntervalUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionEventRouteTimeIntervalUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteTimeIntervalUpdateRequest is a type for the entire action request
type ActionEventRouteTimeIntervalUpdateRequest struct {
	TimeInterval map[string]interface{} "json:\"time_interval\""
	Meta         map[string]interface{} "json:\"_meta\""
}

// ActionEventRouteTimeIntervalUpdateOutput is a type for action output parameters
type ActionEventRouteTimeIntervalUpdateOutput struct {
	CreatedAt         string                             "json:\"created_at\""
	EventTimeInterval *ActionEventTimeIntervalShowOutput "json:\"event_time_interval\""
	Id                int64                              "json:\"id\""
	Mode              string                             "json:\"mode\""
	UpdatedAt         string                             "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionEventRouteTimeIntervalUpdateResponse struct {
	Action *ActionEventRouteTimeIntervalUpdate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		TimeInterval *ActionEventRouteTimeIntervalUpdateOutput "json:\"time_interval\""
	}

	// Action output without the namespace
	Output *ActionEventRouteTimeIntervalUpdateOutput
}

// Prepare the action for invocation
func (action *ActionEventRouteTimeIntervalUpdate) Prepare() *ActionEventRouteTimeIntervalUpdateInvocation {
	return &ActionEventRouteTimeIntervalUpdateInvocation{
		Action: action,
		Path:   "/v7.0/event_routes/{event_route_id}/time_intervals/{time_interval_id}",
	}
}

// ActionEventRouteTimeIntervalUpdateInvocation is used to configure action for invocation
type ActionEventRouteTimeIntervalUpdateInvocation struct {
	// Pointer to the action
	Action *ActionEventRouteTimeIntervalUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEventRouteTimeIntervalUpdateInput
	// Global meta input parameters
	MetaInput *ActionEventRouteTimeIntervalUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventRouteTimeIntervalUpdateInvocation) SetPathParamInt(param string, value int64) *ActionEventRouteTimeIntervalUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventRouteTimeIntervalUpdateInvocation) SetPathParamString(param string, value string) *ActionEventRouteTimeIntervalUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEventRouteTimeIntervalUpdateInvocation) NewInput() *ActionEventRouteTimeIntervalUpdateInput {
	inv.Input = &ActionEventRouteTimeIntervalUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEventRouteTimeIntervalUpdateInvocation) SetInput(input *ActionEventRouteTimeIntervalUpdateInput) *ActionEventRouteTimeIntervalUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEventRouteTimeIntervalUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionEventRouteTimeIntervalUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventRouteTimeIntervalUpdateInvocation) NewMetaInput() *ActionEventRouteTimeIntervalUpdateMetaGlobalInput {
	inv.MetaInput = &ActionEventRouteTimeIntervalUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventRouteTimeIntervalUpdateInvocation) SetMetaInput(input *ActionEventRouteTimeIntervalUpdateMetaGlobalInput) *ActionEventRouteTimeIntervalUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventRouteTimeIntervalUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventRouteTimeIntervalUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventRouteTimeIntervalUpdateInvocation) validate() error {
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
func (inv *ActionEventRouteTimeIntervalUpdateInvocation) Call() (*ActionEventRouteTimeIntervalUpdateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionEventRouteTimeIntervalUpdateInvocation) callAsBody() (*ActionEventRouteTimeIntervalUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionEventRouteTimeIntervalUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.TimeInterval
	}
	return resp, err
}

func (inv *ActionEventRouteTimeIntervalUpdateInvocation) makeAllInputParams() *ActionEventRouteTimeIntervalUpdateRequest {
	return &ActionEventRouteTimeIntervalUpdateRequest{
		TimeInterval: inv.makeInputParams(),
		Meta:         inv.makeMetaInputParams(),
	}
}

func (inv *ActionEventRouteTimeIntervalUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Mode") {
			ret["mode"] = inv.Input.Mode
		}
	}

	return ret
}

func (inv *ActionEventRouteTimeIntervalUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
