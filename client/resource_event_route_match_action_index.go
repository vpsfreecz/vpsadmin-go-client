package client

import (
	"net/url"
	"strings"
)

// ActionEventRouteMatchIndex is a type for action Event.Route_match#Index
type ActionEventRouteMatchIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionEventRouteMatchIndex(client *Client) *ActionEventRouteMatchIndex {
	return &ActionEventRouteMatchIndex{
		Client: client,
	}
}

// ActionEventRouteMatchIndexMetaGlobalInput is a type for action global meta input parameters
type ActionEventRouteMatchIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionEventRouteMatchIndexMetaGlobalInput) SetCount(value bool) *ActionEventRouteMatchIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventRouteMatchIndexMetaGlobalInput) SetIncludes(value string) *ActionEventRouteMatchIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventRouteMatchIndexMetaGlobalInput) SetNo(value bool) *ActionEventRouteMatchIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteMatchIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteMatchIndexMetaGlobalInput) SelectParameters(params ...string) *ActionEventRouteMatchIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventRouteMatchIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteMatchIndexInput is a type for action input parameters
type ActionEventRouteMatchIndexInput struct {
	FromId int64 "json:\"from_id\""
	Limit  int64 "json:\"limit\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionEventRouteMatchIndexInput) SetFromId(value int64) *ActionEventRouteMatchIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionEventRouteMatchIndexInput) SetLimit(value int64) *ActionEventRouteMatchIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteMatchIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteMatchIndexInput) SelectParameters(params ...string) *ActionEventRouteMatchIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionEventRouteMatchIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionEventRouteMatchIndexInput) UnselectParameters(params ...string) *ActionEventRouteMatchIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionEventRouteMatchIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteMatchIndexOutput is a type for action output parameters
type ActionEventRouteMatchIndexOutput struct {
	CreatedAt         string "json:\"created_at\""
	EventRouteId      int64  "json:\"event_route_id\""
	EventRouteLabel   string "json:\"event_route_label\""
	Id                int64  "json:\"id\""
	MatchOrder        int64  "json:\"match_order\""
	RouteOwnerId      int64  "json:\"route_owner_id\""
	RouteOwnerLogin   string "json:\"route_owner_login\""
	Source            string "json:\"source\""
	SubjectRelation   string "json:\"subject_relation\""
	TimeIntervalState string "json:\"time_interval_state\""
	UpdatedAt         string "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionEventRouteMatchIndexResponse struct {
	Action *ActionEventRouteMatchIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		RouteMatches []*ActionEventRouteMatchIndexOutput "json:\"route_matches\""
	}

	// Action output without the namespace
	Output []*ActionEventRouteMatchIndexOutput
}

// Prepare the action for invocation
func (action *ActionEventRouteMatchIndex) Prepare() *ActionEventRouteMatchIndexInvocation {
	return &ActionEventRouteMatchIndexInvocation{
		Action: action,
		Path:   "/v7.0/events/{event_id}/route_matches",
	}
}

// ActionEventRouteMatchIndexInvocation is used to configure action for invocation
type ActionEventRouteMatchIndexInvocation struct {
	// Pointer to the action
	Action *ActionEventRouteMatchIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEventRouteMatchIndexInput
	// Global meta input parameters
	MetaInput *ActionEventRouteMatchIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventRouteMatchIndexInvocation) SetPathParamInt(param string, value int64) *ActionEventRouteMatchIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventRouteMatchIndexInvocation) SetPathParamString(param string, value string) *ActionEventRouteMatchIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEventRouteMatchIndexInvocation) NewInput() *ActionEventRouteMatchIndexInput {
	inv.Input = &ActionEventRouteMatchIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEventRouteMatchIndexInvocation) SetInput(input *ActionEventRouteMatchIndexInput) *ActionEventRouteMatchIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEventRouteMatchIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionEventRouteMatchIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventRouteMatchIndexInvocation) NewMetaInput() *ActionEventRouteMatchIndexMetaGlobalInput {
	inv.MetaInput = &ActionEventRouteMatchIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventRouteMatchIndexInvocation) SetMetaInput(input *ActionEventRouteMatchIndexMetaGlobalInput) *ActionEventRouteMatchIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventRouteMatchIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventRouteMatchIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventRouteMatchIndexInvocation) validate() error {
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
func (inv *ActionEventRouteMatchIndexInvocation) Call() (*ActionEventRouteMatchIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionEventRouteMatchIndexInvocation) callAsQuery() (*ActionEventRouteMatchIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionEventRouteMatchIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.RouteMatches
	}
	return resp, err
}

func (inv *ActionEventRouteMatchIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["route_match[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["route_match[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionEventRouteMatchIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
