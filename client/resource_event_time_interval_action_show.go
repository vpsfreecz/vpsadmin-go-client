package client

import (
	"net/url"
	"strings"
)

// ActionEventTimeIntervalShow is a type for action Event_time_interval#Show
type ActionEventTimeIntervalShow struct {
	// Pointer to client
	Client *Client
}

func NewActionEventTimeIntervalShow(client *Client) *ActionEventTimeIntervalShow {
	return &ActionEventTimeIntervalShow{
		Client: client,
	}
}

// ActionEventTimeIntervalShowMetaGlobalInput is a type for action global meta input parameters
type ActionEventTimeIntervalShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventTimeIntervalShowMetaGlobalInput) SetIncludes(value string) *ActionEventTimeIntervalShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventTimeIntervalShowMetaGlobalInput) SetNo(value bool) *ActionEventTimeIntervalShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventTimeIntervalShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventTimeIntervalShowMetaGlobalInput) SelectParameters(params ...string) *ActionEventTimeIntervalShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventTimeIntervalShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventTimeIntervalShowOutput is a type for action output parameters
type ActionEventTimeIntervalShowOutput struct {
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
type ActionEventTimeIntervalShowResponse struct {
	Action *ActionEventTimeIntervalShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		EventTimeInterval *ActionEventTimeIntervalShowOutput "json:\"event_time_interval\""
	}

	// Action output without the namespace
	Output *ActionEventTimeIntervalShowOutput
}

// Prepare the action for invocation
func (action *ActionEventTimeIntervalShow) Prepare() *ActionEventTimeIntervalShowInvocation {
	return &ActionEventTimeIntervalShowInvocation{
		Action: action,
		Path:   "/v7.0/event_time_intervals/{event_time_interval_id}",
	}
}

// ActionEventTimeIntervalShowInvocation is used to configure action for invocation
type ActionEventTimeIntervalShowInvocation struct {
	// Pointer to the action
	Action *ActionEventTimeIntervalShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionEventTimeIntervalShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventTimeIntervalShowInvocation) SetPathParamInt(param string, value int64) *ActionEventTimeIntervalShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventTimeIntervalShowInvocation) SetPathParamString(param string, value string) *ActionEventTimeIntervalShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventTimeIntervalShowInvocation) NewMetaInput() *ActionEventTimeIntervalShowMetaGlobalInput {
	inv.MetaInput = &ActionEventTimeIntervalShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventTimeIntervalShowInvocation) SetMetaInput(input *ActionEventTimeIntervalShowMetaGlobalInput) *ActionEventTimeIntervalShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventTimeIntervalShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventTimeIntervalShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventTimeIntervalShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEventTimeIntervalShowInvocation) Call() (*ActionEventTimeIntervalShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionEventTimeIntervalShowInvocation) callAsQuery() (*ActionEventTimeIntervalShowResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionEventTimeIntervalShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.EventTimeInterval
	}
	return resp, err
}

func (inv *ActionEventTimeIntervalShowInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
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
