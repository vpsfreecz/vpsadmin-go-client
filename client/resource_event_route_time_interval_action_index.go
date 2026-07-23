package client

import (
	"net/url"
	"strings"
)

// ActionEventRouteTimeIntervalIndex is a type for action Event_route.Time_interval#Index
type ActionEventRouteTimeIntervalIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionEventRouteTimeIntervalIndex(client *Client) *ActionEventRouteTimeIntervalIndex {
	return &ActionEventRouteTimeIntervalIndex{
		Client: client,
	}
}

// ActionEventRouteTimeIntervalIndexMetaGlobalInput is a type for action global meta input parameters
type ActionEventRouteTimeIntervalIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionEventRouteTimeIntervalIndexMetaGlobalInput) SetCount(value bool) *ActionEventRouteTimeIntervalIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventRouteTimeIntervalIndexMetaGlobalInput) SetIncludes(value string) *ActionEventRouteTimeIntervalIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventRouteTimeIntervalIndexMetaGlobalInput) SetNo(value bool) *ActionEventRouteTimeIntervalIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteTimeIntervalIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteTimeIntervalIndexMetaGlobalInput) SelectParameters(params ...string) *ActionEventRouteTimeIntervalIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventRouteTimeIntervalIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteTimeIntervalIndexInput is a type for action input parameters
type ActionEventRouteTimeIntervalIndexInput struct {
	FromId int64 "json:\"from_id\""
	Limit  int64 "json:\"limit\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionEventRouteTimeIntervalIndexInput) SetFromId(value int64) *ActionEventRouteTimeIntervalIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionEventRouteTimeIntervalIndexInput) SetLimit(value int64) *ActionEventRouteTimeIntervalIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteTimeIntervalIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteTimeIntervalIndexInput) SelectParameters(params ...string) *ActionEventRouteTimeIntervalIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionEventRouteTimeIntervalIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionEventRouteTimeIntervalIndexInput) UnselectParameters(params ...string) *ActionEventRouteTimeIntervalIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionEventRouteTimeIntervalIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteTimeIntervalIndexOutput is a type for action output parameters
type ActionEventRouteTimeIntervalIndexOutput struct {
	CreatedAt         string                             "json:\"created_at\""
	EventTimeInterval *ActionEventTimeIntervalShowOutput "json:\"event_time_interval\""
	Id                int64                              "json:\"id\""
	Mode              string                             "json:\"mode\""
	UpdatedAt         string                             "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionEventRouteTimeIntervalIndexResponse struct {
	Action *ActionEventRouteTimeIntervalIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		TimeIntervals []*ActionEventRouteTimeIntervalIndexOutput "json:\"time_intervals\""
	}

	// Action output without the namespace
	Output []*ActionEventRouteTimeIntervalIndexOutput
}

// Prepare the action for invocation
func (action *ActionEventRouteTimeIntervalIndex) Prepare() *ActionEventRouteTimeIntervalIndexInvocation {
	return &ActionEventRouteTimeIntervalIndexInvocation{
		Action: action,
		Path:   "/v7.0/event_routes/{event_route_id}/time_intervals",
	}
}

// ActionEventRouteTimeIntervalIndexInvocation is used to configure action for invocation
type ActionEventRouteTimeIntervalIndexInvocation struct {
	// Pointer to the action
	Action *ActionEventRouteTimeIntervalIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEventRouteTimeIntervalIndexInput
	// Global meta input parameters
	MetaInput *ActionEventRouteTimeIntervalIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventRouteTimeIntervalIndexInvocation) SetPathParamInt(param string, value int64) *ActionEventRouteTimeIntervalIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventRouteTimeIntervalIndexInvocation) SetPathParamString(param string, value string) *ActionEventRouteTimeIntervalIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEventRouteTimeIntervalIndexInvocation) NewInput() *ActionEventRouteTimeIntervalIndexInput {
	inv.Input = &ActionEventRouteTimeIntervalIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEventRouteTimeIntervalIndexInvocation) SetInput(input *ActionEventRouteTimeIntervalIndexInput) *ActionEventRouteTimeIntervalIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEventRouteTimeIntervalIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionEventRouteTimeIntervalIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventRouteTimeIntervalIndexInvocation) NewMetaInput() *ActionEventRouteTimeIntervalIndexMetaGlobalInput {
	inv.MetaInput = &ActionEventRouteTimeIntervalIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventRouteTimeIntervalIndexInvocation) SetMetaInput(input *ActionEventRouteTimeIntervalIndexMetaGlobalInput) *ActionEventRouteTimeIntervalIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventRouteTimeIntervalIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventRouteTimeIntervalIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventRouteTimeIntervalIndexInvocation) validate() error {
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
func (inv *ActionEventRouteTimeIntervalIndexInvocation) Call() (*ActionEventRouteTimeIntervalIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionEventRouteTimeIntervalIndexInvocation) callAsQuery() (*ActionEventRouteTimeIntervalIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionEventRouteTimeIntervalIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.TimeIntervals
	}
	return resp, err
}

func (inv *ActionEventRouteTimeIntervalIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["time_interval[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["time_interval[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionEventRouteTimeIntervalIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
