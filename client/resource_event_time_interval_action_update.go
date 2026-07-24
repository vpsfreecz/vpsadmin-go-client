package client

import (
	"net/url"
	"strings"
)

// ActionEventTimeIntervalUpdate is a type for action Event_time_interval#Update
type ActionEventTimeIntervalUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionEventTimeIntervalUpdate(client *Client) *ActionEventTimeIntervalUpdate {
	return &ActionEventTimeIntervalUpdate{
		Client: client,
	}
}

// ActionEventTimeIntervalUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionEventTimeIntervalUpdateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventTimeIntervalUpdateMetaGlobalInput) SetIncludes(value string) *ActionEventTimeIntervalUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventTimeIntervalUpdateMetaGlobalInput) SetNo(value bool) *ActionEventTimeIntervalUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventTimeIntervalUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventTimeIntervalUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionEventTimeIntervalUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventTimeIntervalUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventTimeIntervalUpdateInput is a type for action input parameters
type ActionEventTimeIntervalUpdateInput struct {
	Name     string      "json:\"name\""
	Specs    interface{} "json:\"specs\""
	TimeZone string      "json:\"time_zone\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionEventTimeIntervalUpdateInput) SetName(value string) *ActionEventTimeIntervalUpdateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetSpecs sets parameter Specs to value and selects it for sending
func (in *ActionEventTimeIntervalUpdateInput) SetSpecs(value interface{}) *ActionEventTimeIntervalUpdateInput {
	in.Specs = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Specs"] = nil
	return in
}

// SetTimeZone sets parameter TimeZone to value and selects it for sending
func (in *ActionEventTimeIntervalUpdateInput) SetTimeZone(value string) *ActionEventTimeIntervalUpdateInput {
	in.TimeZone = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TimeZone"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventTimeIntervalUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventTimeIntervalUpdateInput) SelectParameters(params ...string) *ActionEventTimeIntervalUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionEventTimeIntervalUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionEventTimeIntervalUpdateInput) UnselectParameters(params ...string) *ActionEventTimeIntervalUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionEventTimeIntervalUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventTimeIntervalUpdateRequest is a type for the entire action request
type ActionEventTimeIntervalUpdateRequest struct {
	EventTimeInterval map[string]interface{} "json:\"event_time_interval\""
	Meta              map[string]interface{} "json:\"_meta\""
}

// ActionEventTimeIntervalUpdateOutput is a type for action output parameters
type ActionEventTimeIntervalUpdateOutput struct {
	ActiveRouteReferenceCount int64                 "json:\"active_route_reference_count\""
	CreatedAt                 string                "json:\"created_at\""
	DisplaySummary            string                "json:\"display_summary\""
	Id                        int64                 "json:\"id\""
	MatchesNow                bool                  "json:\"matches_now\""
	MuteRouteReferenceCount   int64                 "json:\"mute_route_reference_count\""
	Name                      string                "json:\"name\""
	RouteReferenceCount       int64                 "json:\"route_reference_count\""
	Specs                     interface{}           "json:\"specs\""
	TimeZone                  string                "json:\"time_zone\""
	UpdatedAt                 string                "json:\"updated_at\""
	User                      *ActionUserShowOutput "json:\"user\""
}

// Type for action response, including envelope
type ActionEventTimeIntervalUpdateResponse struct {
	Action *ActionEventTimeIntervalUpdate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		EventTimeInterval *ActionEventTimeIntervalUpdateOutput "json:\"event_time_interval\""
	}

	// Action output without the namespace
	Output *ActionEventTimeIntervalUpdateOutput
}

// Prepare the action for invocation
func (action *ActionEventTimeIntervalUpdate) Prepare() *ActionEventTimeIntervalUpdateInvocation {
	return &ActionEventTimeIntervalUpdateInvocation{
		Action: action,
		Path:   "/v7.0/event_time_intervals/{event_time_interval_id}",
	}
}

// ActionEventTimeIntervalUpdateInvocation is used to configure action for invocation
type ActionEventTimeIntervalUpdateInvocation struct {
	// Pointer to the action
	Action *ActionEventTimeIntervalUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEventTimeIntervalUpdateInput
	// Global meta input parameters
	MetaInput *ActionEventTimeIntervalUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventTimeIntervalUpdateInvocation) SetPathParamInt(param string, value int64) *ActionEventTimeIntervalUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventTimeIntervalUpdateInvocation) SetPathParamString(param string, value string) *ActionEventTimeIntervalUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEventTimeIntervalUpdateInvocation) NewInput() *ActionEventTimeIntervalUpdateInput {
	inv.Input = &ActionEventTimeIntervalUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEventTimeIntervalUpdateInvocation) SetInput(input *ActionEventTimeIntervalUpdateInput) *ActionEventTimeIntervalUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEventTimeIntervalUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionEventTimeIntervalUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventTimeIntervalUpdateInvocation) NewMetaInput() *ActionEventTimeIntervalUpdateMetaGlobalInput {
	inv.MetaInput = &ActionEventTimeIntervalUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventTimeIntervalUpdateInvocation) SetMetaInput(input *ActionEventTimeIntervalUpdateMetaGlobalInput) *ActionEventTimeIntervalUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventTimeIntervalUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventTimeIntervalUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventTimeIntervalUpdateInvocation) validate() error {
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
func (inv *ActionEventTimeIntervalUpdateInvocation) Call() (*ActionEventTimeIntervalUpdateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionEventTimeIntervalUpdateInvocation) callAsBody() (*ActionEventTimeIntervalUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionEventTimeIntervalUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.EventTimeInterval
	}
	return resp, err
}

func (inv *ActionEventTimeIntervalUpdateInvocation) makeAllInputParams() *ActionEventTimeIntervalUpdateRequest {
	return &ActionEventTimeIntervalUpdateRequest{
		EventTimeInterval: inv.makeInputParams(),
		Meta:              inv.makeMetaInputParams(),
	}
}

func (inv *ActionEventTimeIntervalUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Name") {
			ret["name"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Specs") {
			ret["specs"] = inv.Input.Specs
		}
		if inv.IsParameterSelected("TimeZone") {
			ret["time_zone"] = inv.Input.TimeZone
		}
	}

	return ret
}

func (inv *ActionEventTimeIntervalUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
