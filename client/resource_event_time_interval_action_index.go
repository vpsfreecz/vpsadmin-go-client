package client

import ()

// ActionEventTimeIntervalIndex is a type for action Event_time_interval#Index
type ActionEventTimeIntervalIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionEventTimeIntervalIndex(client *Client) *ActionEventTimeIntervalIndex {
	return &ActionEventTimeIntervalIndex{
		Client: client,
	}
}

// ActionEventTimeIntervalIndexMetaGlobalInput is a type for action global meta input parameters
type ActionEventTimeIntervalIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionEventTimeIntervalIndexMetaGlobalInput) SetCount(value bool) *ActionEventTimeIntervalIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventTimeIntervalIndexMetaGlobalInput) SetIncludes(value string) *ActionEventTimeIntervalIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventTimeIntervalIndexMetaGlobalInput) SetNo(value bool) *ActionEventTimeIntervalIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventTimeIntervalIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventTimeIntervalIndexMetaGlobalInput) SelectParameters(params ...string) *ActionEventTimeIntervalIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventTimeIntervalIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventTimeIntervalIndexInput is a type for action input parameters
type ActionEventTimeIntervalIndexInput struct {
	FromId int64 "json:\"from_id\""
	Limit  int64 "json:\"limit\""
	User   int64 "json:\"user\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionEventTimeIntervalIndexInput) SetFromId(value int64) *ActionEventTimeIntervalIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionEventTimeIntervalIndexInput) SetLimit(value int64) *ActionEventTimeIntervalIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionEventTimeIntervalIndexInput) SetUser(value int64) *ActionEventTimeIntervalIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventTimeIntervalIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventTimeIntervalIndexInput) SelectParameters(params ...string) *ActionEventTimeIntervalIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionEventTimeIntervalIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionEventTimeIntervalIndexInput) UnselectParameters(params ...string) *ActionEventTimeIntervalIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionEventTimeIntervalIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventTimeIntervalIndexOutput is a type for action output parameters
type ActionEventTimeIntervalIndexOutput struct {
	ActiveRouteReferenceCount int64                 "json:\"active_route_reference_count\""
	CreatedAt                 string                "json:\"created_at\""
	DisplaySummary            string                "json:\"display_summary\""
	Id                        int64                 "json:\"id\""
	MatchesNow                bool                  "json:\"matches_now\""
	MuteRouteReferenceCount   int64                 "json:\"mute_route_reference_count\""
	Name                      string                "json:\"name\""
	RouteReferenceCount       int64                 "json:\"route_reference_count\""
	TimeZone                  string                "json:\"time_zone\""
	UpdatedAt                 string                "json:\"updated_at\""
	User                      *ActionUserShowOutput "json:\"user\""
}

// Type for action response, including envelope
type ActionEventTimeIntervalIndexResponse struct {
	Action *ActionEventTimeIntervalIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		EventTimeIntervals []*ActionEventTimeIntervalIndexOutput "json:\"event_time_intervals\""
	}

	// Action output without the namespace
	Output []*ActionEventTimeIntervalIndexOutput
}

// Prepare the action for invocation
func (action *ActionEventTimeIntervalIndex) Prepare() *ActionEventTimeIntervalIndexInvocation {
	return &ActionEventTimeIntervalIndexInvocation{
		Action: action,
		Path:   "/v7.0/event_time_intervals",
	}
}

// ActionEventTimeIntervalIndexInvocation is used to configure action for invocation
type ActionEventTimeIntervalIndexInvocation struct {
	// Pointer to the action
	Action *ActionEventTimeIntervalIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEventTimeIntervalIndexInput
	// Global meta input parameters
	MetaInput *ActionEventTimeIntervalIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEventTimeIntervalIndexInvocation) NewInput() *ActionEventTimeIntervalIndexInput {
	inv.Input = &ActionEventTimeIntervalIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEventTimeIntervalIndexInvocation) SetInput(input *ActionEventTimeIntervalIndexInput) *ActionEventTimeIntervalIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEventTimeIntervalIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionEventTimeIntervalIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventTimeIntervalIndexInvocation) NewMetaInput() *ActionEventTimeIntervalIndexMetaGlobalInput {
	inv.MetaInput = &ActionEventTimeIntervalIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventTimeIntervalIndexInvocation) SetMetaInput(input *ActionEventTimeIntervalIndexMetaGlobalInput) *ActionEventTimeIntervalIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventTimeIntervalIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventTimeIntervalIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventTimeIntervalIndexInvocation) validate() error {
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
func (inv *ActionEventTimeIntervalIndexInvocation) Call() (*ActionEventTimeIntervalIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionEventTimeIntervalIndexInvocation) callAsQuery() (*ActionEventTimeIntervalIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionEventTimeIntervalIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.EventTimeIntervals
	}
	return resp, err
}

func (inv *ActionEventTimeIntervalIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["event_time_interval[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["event_time_interval[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("User") {
			ret["event_time_interval[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionEventTimeIntervalIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
