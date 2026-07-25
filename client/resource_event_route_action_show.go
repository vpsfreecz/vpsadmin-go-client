package client

import (
	"net/url"
	"strings"
)

// ActionEventRouteShow is a type for action Event_route#Show
type ActionEventRouteShow struct {
	// Pointer to client
	Client *Client
}

func NewActionEventRouteShow(client *Client) *ActionEventRouteShow {
	return &ActionEventRouteShow{
		Client: client,
	}
}

// ActionEventRouteShowMetaGlobalInput is a type for action global meta input parameters
type ActionEventRouteShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventRouteShowMetaGlobalInput) SetIncludes(value string) *ActionEventRouteShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventRouteShowMetaGlobalInput) SetNo(value bool) *ActionEventRouteShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteShowMetaGlobalInput) SelectParameters(params ...string) *ActionEventRouteShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventRouteShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteShowOutput is a type for action output parameters
type ActionEventRouteShowOutput struct {
	Continue               bool                  "json:\"continue\""
	CreatedAt              string                "json:\"created_at\""
	DisplayLabel           string                "json:\"display_label\""
	Enabled                bool                  "json:\"enabled\""
	EventType              string                "json:\"event_type\""
	EventTypePattern       string                "json:\"event_type_pattern\""
	ExpiresAt              string                "json:\"expires_at\""
	GroupBy                interface{}           "json:\"group_by\""
	GroupIntervalSeconds   int64                 "json:\"group_interval_seconds\""
	GroupWaitSeconds       int64                 "json:\"group_wait_seconds\""
	GroupingEnabled        bool                  "json:\"grouping_enabled\""
	GroupingSummary        string                "json:\"grouping_summary\""
	HitCount               int64                 "json:\"hit_count\""
	Id                     int64                 "json:\"id\""
	Label                  string                "json:\"label\""
	MatcherCount           int64                 "json:\"matcher_count\""
	MatcherSummary         string                "json:\"matcher_summary\""
	NotificationReceiverId int64                 "json:\"notification_receiver_id\""
	ParentId               int64                 "json:\"parent_id\""
	Position               int64                 "json:\"position\""
	SingleUse              bool                  "json:\"single_use\""
	SpentAt                string                "json:\"spent_at\""
	SubjectScope           string                "json:\"subject_scope\""
	UpdatedAt              string                "json:\"updated_at\""
	User                   *ActionUserShowOutput "json:\"user\""
}

// Type for action response, including envelope
type ActionEventRouteShowResponse struct {
	Action *ActionEventRouteShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		EventRoute *ActionEventRouteShowOutput "json:\"event_route\""
	}

	// Action output without the namespace
	Output *ActionEventRouteShowOutput
}

// Prepare the action for invocation
func (action *ActionEventRouteShow) Prepare() *ActionEventRouteShowInvocation {
	return &ActionEventRouteShowInvocation{
		Action: action,
		Path:   "/v7.0/event_routes/{event_route_id}",
	}
}

// ActionEventRouteShowInvocation is used to configure action for invocation
type ActionEventRouteShowInvocation struct {
	// Pointer to the action
	Action *ActionEventRouteShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionEventRouteShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventRouteShowInvocation) SetPathParamInt(param string, value int64) *ActionEventRouteShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventRouteShowInvocation) SetPathParamString(param string, value string) *ActionEventRouteShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventRouteShowInvocation) NewMetaInput() *ActionEventRouteShowMetaGlobalInput {
	inv.MetaInput = &ActionEventRouteShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventRouteShowInvocation) SetMetaInput(input *ActionEventRouteShowMetaGlobalInput) *ActionEventRouteShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventRouteShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventRouteShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventRouteShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEventRouteShowInvocation) Call() (*ActionEventRouteShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionEventRouteShowInvocation) callAsQuery() (*ActionEventRouteShowResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionEventRouteShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.EventRoute
	}
	return resp, err
}

func (inv *ActionEventRouteShowInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
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
