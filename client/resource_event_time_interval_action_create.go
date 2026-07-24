package client

import ()

// ActionEventTimeIntervalCreate is a type for action Event_time_interval#Create
type ActionEventTimeIntervalCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionEventTimeIntervalCreate(client *Client) *ActionEventTimeIntervalCreate {
	return &ActionEventTimeIntervalCreate{
		Client: client,
	}
}

// ActionEventTimeIntervalCreateMetaGlobalInput is a type for action global meta input parameters
type ActionEventTimeIntervalCreateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventTimeIntervalCreateMetaGlobalInput) SetIncludes(value string) *ActionEventTimeIntervalCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventTimeIntervalCreateMetaGlobalInput) SetNo(value bool) *ActionEventTimeIntervalCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventTimeIntervalCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventTimeIntervalCreateMetaGlobalInput) SelectParameters(params ...string) *ActionEventTimeIntervalCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventTimeIntervalCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventTimeIntervalCreateInput is a type for action input parameters
type ActionEventTimeIntervalCreateInput struct {
	Name     string      "json:\"name\""
	Specs    interface{} "json:\"specs\""
	TimeZone string      "json:\"time_zone\""
	User     int64       "json:\"user\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionEventTimeIntervalCreateInput) SetName(value string) *ActionEventTimeIntervalCreateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetSpecs sets parameter Specs to value and selects it for sending
func (in *ActionEventTimeIntervalCreateInput) SetSpecs(value interface{}) *ActionEventTimeIntervalCreateInput {
	in.Specs = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Specs"] = nil
	return in
}

// SetTimeZone sets parameter TimeZone to value and selects it for sending
func (in *ActionEventTimeIntervalCreateInput) SetTimeZone(value string) *ActionEventTimeIntervalCreateInput {
	in.TimeZone = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TimeZone"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionEventTimeIntervalCreateInput) SetUser(value int64) *ActionEventTimeIntervalCreateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventTimeIntervalCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventTimeIntervalCreateInput) SelectParameters(params ...string) *ActionEventTimeIntervalCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionEventTimeIntervalCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionEventTimeIntervalCreateInput) UnselectParameters(params ...string) *ActionEventTimeIntervalCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionEventTimeIntervalCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventTimeIntervalCreateRequest is a type for the entire action request
type ActionEventTimeIntervalCreateRequest struct {
	EventTimeInterval map[string]interface{} "json:\"event_time_interval\""
	Meta              map[string]interface{} "json:\"_meta\""
}

// ActionEventTimeIntervalCreateOutput is a type for action output parameters
type ActionEventTimeIntervalCreateOutput struct {
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
type ActionEventTimeIntervalCreateResponse struct {
	Action *ActionEventTimeIntervalCreate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		EventTimeInterval *ActionEventTimeIntervalCreateOutput "json:\"event_time_interval\""
	}

	// Action output without the namespace
	Output *ActionEventTimeIntervalCreateOutput
}

// Prepare the action for invocation
func (action *ActionEventTimeIntervalCreate) Prepare() *ActionEventTimeIntervalCreateInvocation {
	return &ActionEventTimeIntervalCreateInvocation{
		Action: action,
		Path:   "/v7.0/event_time_intervals",
	}
}

// ActionEventTimeIntervalCreateInvocation is used to configure action for invocation
type ActionEventTimeIntervalCreateInvocation struct {
	// Pointer to the action
	Action *ActionEventTimeIntervalCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEventTimeIntervalCreateInput
	// Global meta input parameters
	MetaInput *ActionEventTimeIntervalCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEventTimeIntervalCreateInvocation) NewInput() *ActionEventTimeIntervalCreateInput {
	inv.Input = &ActionEventTimeIntervalCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEventTimeIntervalCreateInvocation) SetInput(input *ActionEventTimeIntervalCreateInput) *ActionEventTimeIntervalCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEventTimeIntervalCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionEventTimeIntervalCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventTimeIntervalCreateInvocation) NewMetaInput() *ActionEventTimeIntervalCreateMetaGlobalInput {
	inv.MetaInput = &ActionEventTimeIntervalCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventTimeIntervalCreateInvocation) SetMetaInput(input *ActionEventTimeIntervalCreateMetaGlobalInput) *ActionEventTimeIntervalCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventTimeIntervalCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventTimeIntervalCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventTimeIntervalCreateInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("User") {
			if !inv.IsParameterNil("User") {
				if inv.Input.User < 0 {
					verr.Add("user", "not a valid resource id")
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
func (inv *ActionEventTimeIntervalCreateInvocation) Call() (*ActionEventTimeIntervalCreateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionEventTimeIntervalCreateInvocation) callAsBody() (*ActionEventTimeIntervalCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionEventTimeIntervalCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.EventTimeInterval
	}
	return resp, err
}

func (inv *ActionEventTimeIntervalCreateInvocation) makeAllInputParams() *ActionEventTimeIntervalCreateRequest {
	return &ActionEventTimeIntervalCreateRequest{
		EventTimeInterval: inv.makeInputParams(),
		Meta:              inv.makeMetaInputParams(),
	}
}

func (inv *ActionEventTimeIntervalCreateInvocation) makeInputParams() map[string]interface{} {
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
		if inv.IsParameterSelected("User") {
			ret["user"] = inv.Input.User
		}
	}

	return ret
}

func (inv *ActionEventTimeIntervalCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
