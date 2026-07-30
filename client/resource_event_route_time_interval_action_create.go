package client

import (
	"net/url"
	"strings"
)

// ActionEventRouteTimeIntervalCreate is a type for action Event_route.Time_interval#Create
type ActionEventRouteTimeIntervalCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionEventRouteTimeIntervalCreate(client *Client) *ActionEventRouteTimeIntervalCreate {
	return &ActionEventRouteTimeIntervalCreate{
		Client: client,
	}
}

// ActionEventRouteTimeIntervalCreateMetaGlobalInput is a type for action global meta input parameters
type ActionEventRouteTimeIntervalCreateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventRouteTimeIntervalCreateMetaGlobalInput) SetIncludes(value string) *ActionEventRouteTimeIntervalCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventRouteTimeIntervalCreateMetaGlobalInput) SetNo(value bool) *ActionEventRouteTimeIntervalCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteTimeIntervalCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteTimeIntervalCreateMetaGlobalInput) SelectParameters(params ...string) *ActionEventRouteTimeIntervalCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventRouteTimeIntervalCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteTimeIntervalCreateInput is a type for action input parameters
type ActionEventRouteTimeIntervalCreateInput struct {
	EventTimeInterval int64  "json:\"event_time_interval\""
	Mode              string "json:\"mode\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEventTimeInterval sets parameter EventTimeInterval to value and selects it for sending
func (in *ActionEventRouteTimeIntervalCreateInput) SetEventTimeInterval(value int64) *ActionEventRouteTimeIntervalCreateInput {
	in.EventTimeInterval = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EventTimeInterval"] = nil
	return in
}

// SetMode sets parameter Mode to value and selects it for sending
func (in *ActionEventRouteTimeIntervalCreateInput) SetMode(value string) *ActionEventRouteTimeIntervalCreateInput {
	in.Mode = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Mode"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteTimeIntervalCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteTimeIntervalCreateInput) SelectParameters(params ...string) *ActionEventRouteTimeIntervalCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionEventRouteTimeIntervalCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionEventRouteTimeIntervalCreateInput) UnselectParameters(params ...string) *ActionEventRouteTimeIntervalCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionEventRouteTimeIntervalCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteTimeIntervalCreateRequest is a type for the entire action request
type ActionEventRouteTimeIntervalCreateRequest struct {
	TimeInterval map[string]interface{} "json:\"time_interval\""
	Meta         map[string]interface{} "json:\"_meta\""
}

// ActionEventRouteTimeIntervalCreateOutput is a type for action output parameters
type ActionEventRouteTimeIntervalCreateOutput struct {
	CreatedAt         string                             "json:\"created_at\""
	EventTimeInterval *ActionEventTimeIntervalShowOutput "json:\"event_time_interval\""
	Id                int64                              "json:\"id\""
	Mode              string                             "json:\"mode\""
	UpdatedAt         string                             "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionEventRouteTimeIntervalCreateResponse struct {
	Action *ActionEventRouteTimeIntervalCreate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		TimeInterval *ActionEventRouteTimeIntervalCreateOutput "json:\"time_interval\""
	}

	// Action output without the namespace
	Output *ActionEventRouteTimeIntervalCreateOutput
}

// Prepare the action for invocation
func (action *ActionEventRouteTimeIntervalCreate) Prepare() *ActionEventRouteTimeIntervalCreateInvocation {
	return &ActionEventRouteTimeIntervalCreateInvocation{
		Action: action,
		Path:   "/v7.0/event_routes/{event_route_id}/time_intervals",
	}
}

// ActionEventRouteTimeIntervalCreateInvocation is used to configure action for invocation
type ActionEventRouteTimeIntervalCreateInvocation struct {
	// Pointer to the action
	Action *ActionEventRouteTimeIntervalCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEventRouteTimeIntervalCreateInput
	// Global meta input parameters
	MetaInput *ActionEventRouteTimeIntervalCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventRouteTimeIntervalCreateInvocation) SetPathParamInt(param string, value int64) *ActionEventRouteTimeIntervalCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventRouteTimeIntervalCreateInvocation) SetPathParamString(param string, value string) *ActionEventRouteTimeIntervalCreateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEventRouteTimeIntervalCreateInvocation) NewInput() *ActionEventRouteTimeIntervalCreateInput {
	inv.Input = &ActionEventRouteTimeIntervalCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEventRouteTimeIntervalCreateInvocation) SetInput(input *ActionEventRouteTimeIntervalCreateInput) *ActionEventRouteTimeIntervalCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEventRouteTimeIntervalCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionEventRouteTimeIntervalCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventRouteTimeIntervalCreateInvocation) NewMetaInput() *ActionEventRouteTimeIntervalCreateMetaGlobalInput {
	inv.MetaInput = &ActionEventRouteTimeIntervalCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventRouteTimeIntervalCreateInvocation) SetMetaInput(input *ActionEventRouteTimeIntervalCreateMetaGlobalInput) *ActionEventRouteTimeIntervalCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventRouteTimeIntervalCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventRouteTimeIntervalCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventRouteTimeIntervalCreateInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("EventTimeInterval") {
			if !inv.IsParameterNil("EventTimeInterval") {
				if inv.Input.EventTimeInterval < 0 {
					verr.Add("event_time_interval", "not a valid resource id")
				}
			}
		}
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEventRouteTimeIntervalCreateInvocation) Call() (*ActionEventRouteTimeIntervalCreateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionEventRouteTimeIntervalCreateInvocation) callAsBody() (*ActionEventRouteTimeIntervalCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionEventRouteTimeIntervalCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.TimeInterval
	}
	return resp, err
}

func (inv *ActionEventRouteTimeIntervalCreateInvocation) makeAllInputParams() *ActionEventRouteTimeIntervalCreateRequest {
	return &ActionEventRouteTimeIntervalCreateRequest{
		TimeInterval: inv.makeInputParams(),
		Meta:         inv.makeMetaInputParams(),
	}
}

func (inv *ActionEventRouteTimeIntervalCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("EventTimeInterval") {
			ret["event_time_interval"] = inv.Input.EventTimeInterval
		}
		if inv.IsParameterSelected("Mode") {
			ret["mode"] = inv.Input.Mode
		}
	}

	return ret
}

func (inv *ActionEventRouteTimeIntervalCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
