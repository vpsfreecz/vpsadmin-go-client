package client

import (
	"net/url"
	"strings"
)

// ActionEventRouteTimeIntervalShow is a type for action Event_route.Time_interval#Show
type ActionEventRouteTimeIntervalShow struct {
	// Pointer to client
	Client *Client
}

func NewActionEventRouteTimeIntervalShow(client *Client) *ActionEventRouteTimeIntervalShow {
	return &ActionEventRouteTimeIntervalShow{
		Client: client,
	}
}

// ActionEventRouteTimeIntervalShowMetaGlobalInput is a type for action global meta input parameters
type ActionEventRouteTimeIntervalShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventRouteTimeIntervalShowMetaGlobalInput) SetIncludes(value string) *ActionEventRouteTimeIntervalShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventRouteTimeIntervalShowMetaGlobalInput) SetNo(value bool) *ActionEventRouteTimeIntervalShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteTimeIntervalShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteTimeIntervalShowMetaGlobalInput) SelectParameters(params ...string) *ActionEventRouteTimeIntervalShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventRouteTimeIntervalShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteTimeIntervalShowOutput is a type for action output parameters
type ActionEventRouteTimeIntervalShowOutput struct {
	CreatedAt         string                             "json:\"created_at\""
	EventTimeInterval *ActionEventTimeIntervalShowOutput "json:\"event_time_interval\""
	Id                int64                              "json:\"id\""
	Mode              string                             "json:\"mode\""
	UpdatedAt         string                             "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionEventRouteTimeIntervalShowResponse struct {
	Action *ActionEventRouteTimeIntervalShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		TimeInterval *ActionEventRouteTimeIntervalShowOutput "json:\"time_interval\""
	}

	// Action output without the namespace
	Output *ActionEventRouteTimeIntervalShowOutput
}

// Prepare the action for invocation
func (action *ActionEventRouteTimeIntervalShow) Prepare() *ActionEventRouteTimeIntervalShowInvocation {
	return &ActionEventRouteTimeIntervalShowInvocation{
		Action: action,
		Path:   "/v7.0/event_routes/{event_route_id}/time_intervals/{time_interval_id}",
	}
}

// ActionEventRouteTimeIntervalShowInvocation is used to configure action for invocation
type ActionEventRouteTimeIntervalShowInvocation struct {
	// Pointer to the action
	Action *ActionEventRouteTimeIntervalShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionEventRouteTimeIntervalShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventRouteTimeIntervalShowInvocation) SetPathParamInt(param string, value int64) *ActionEventRouteTimeIntervalShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventRouteTimeIntervalShowInvocation) SetPathParamString(param string, value string) *ActionEventRouteTimeIntervalShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventRouteTimeIntervalShowInvocation) NewMetaInput() *ActionEventRouteTimeIntervalShowMetaGlobalInput {
	inv.MetaInput = &ActionEventRouteTimeIntervalShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventRouteTimeIntervalShowInvocation) SetMetaInput(input *ActionEventRouteTimeIntervalShowMetaGlobalInput) *ActionEventRouteTimeIntervalShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventRouteTimeIntervalShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventRouteTimeIntervalShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventRouteTimeIntervalShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEventRouteTimeIntervalShowInvocation) Call() (*ActionEventRouteTimeIntervalShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionEventRouteTimeIntervalShowInvocation) callAsQuery() (*ActionEventRouteTimeIntervalShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionEventRouteTimeIntervalShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.TimeInterval
	}
	return resp, err
}

func (inv *ActionEventRouteTimeIntervalShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
