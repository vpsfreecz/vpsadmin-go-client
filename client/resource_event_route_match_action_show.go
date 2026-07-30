package client

import (
	"net/url"
	"strings"
)

// ActionEventRouteMatchShow is a type for action Event.Route_match#Show
type ActionEventRouteMatchShow struct {
	// Pointer to client
	Client *Client
}

func NewActionEventRouteMatchShow(client *Client) *ActionEventRouteMatchShow {
	return &ActionEventRouteMatchShow{
		Client: client,
	}
}

// ActionEventRouteMatchShowMetaGlobalInput is a type for action global meta input parameters
type ActionEventRouteMatchShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventRouteMatchShowMetaGlobalInput) SetIncludes(value string) *ActionEventRouteMatchShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventRouteMatchShowMetaGlobalInput) SetNo(value bool) *ActionEventRouteMatchShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteMatchShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteMatchShowMetaGlobalInput) SelectParameters(params ...string) *ActionEventRouteMatchShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventRouteMatchShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteMatchShowOutput is a type for action output parameters
type ActionEventRouteMatchShowOutput struct {
	CreatedAt            string      "json:\"created_at\""
	EventRouteId         int64       "json:\"event_route_id\""
	EventRouteLabel      string      "json:\"event_route_label\""
	Id                   int64       "json:\"id\""
	MatchOrder           int64       "json:\"match_order\""
	RouteOwnerId         int64       "json:\"route_owner_id\""
	RouteOwnerLogin      string      "json:\"route_owner_login\""
	Source               string      "json:\"source\""
	SubjectRelation      string      "json:\"subject_relation\""
	TimeIntervalSnapshot interface{} "json:\"time_interval_snapshot\""
	TimeIntervalState    string      "json:\"time_interval_state\""
	UpdatedAt            string      "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionEventRouteMatchShowResponse struct {
	Action *ActionEventRouteMatchShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		RouteMatch *ActionEventRouteMatchShowOutput "json:\"route_match\""
	}

	// Action output without the namespace
	Output *ActionEventRouteMatchShowOutput
}

// Prepare the action for invocation
func (action *ActionEventRouteMatchShow) Prepare() *ActionEventRouteMatchShowInvocation {
	return &ActionEventRouteMatchShowInvocation{
		Action: action,
		Path:   "/v7.0/events/{event_id}/route_matches/{route_match_id}",
	}
}

// ActionEventRouteMatchShowInvocation is used to configure action for invocation
type ActionEventRouteMatchShowInvocation struct {
	// Pointer to the action
	Action *ActionEventRouteMatchShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionEventRouteMatchShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventRouteMatchShowInvocation) SetPathParamInt(param string, value int64) *ActionEventRouteMatchShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventRouteMatchShowInvocation) SetPathParamString(param string, value string) *ActionEventRouteMatchShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventRouteMatchShowInvocation) NewMetaInput() *ActionEventRouteMatchShowMetaGlobalInput {
	inv.MetaInput = &ActionEventRouteMatchShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventRouteMatchShowInvocation) SetMetaInput(input *ActionEventRouteMatchShowMetaGlobalInput) *ActionEventRouteMatchShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventRouteMatchShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventRouteMatchShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventRouteMatchShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEventRouteMatchShowInvocation) Call() (*ActionEventRouteMatchShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionEventRouteMatchShowInvocation) callAsQuery() (*ActionEventRouteMatchShowResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionEventRouteMatchShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.RouteMatch
	}
	return resp, err
}

func (inv *ActionEventRouteMatchShowInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			queryValue, err := convertCustomToString(inv.MetaInput.Includes)
			if err != nil {
				return err
			}
			ret["_meta[includes]"] = queryValue
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}

	return nil
}
