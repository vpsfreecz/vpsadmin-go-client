package client

import (
	"net/url"
	"strings"
)

// ActionEventRouteTimeIntervalDelete is a type for action Event_route.Time_interval#Delete
type ActionEventRouteTimeIntervalDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionEventRouteTimeIntervalDelete(client *Client) *ActionEventRouteTimeIntervalDelete {
	return &ActionEventRouteTimeIntervalDelete{
		Client: client,
	}
}

// ActionEventRouteTimeIntervalDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionEventRouteTimeIntervalDeleteMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventRouteTimeIntervalDeleteMetaGlobalInput) SetIncludes(value string) *ActionEventRouteTimeIntervalDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventRouteTimeIntervalDeleteMetaGlobalInput) SetNo(value bool) *ActionEventRouteTimeIntervalDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteTimeIntervalDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteTimeIntervalDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionEventRouteTimeIntervalDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventRouteTimeIntervalDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteTimeIntervalDeleteRequest is a type for the entire action request
type ActionEventRouteTimeIntervalDeleteRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// Type for action response, including envelope
type ActionEventRouteTimeIntervalDeleteResponse struct {
	Action *ActionEventRouteTimeIntervalDelete "json:\"-\""
	*Envelope
}

// Prepare the action for invocation
func (action *ActionEventRouteTimeIntervalDelete) Prepare() *ActionEventRouteTimeIntervalDeleteInvocation {
	return &ActionEventRouteTimeIntervalDeleteInvocation{
		Action: action,
		Path:   "/v7.0/event_routes/{event_route_id}/time_intervals/{time_interval_id}",
	}
}

// ActionEventRouteTimeIntervalDeleteInvocation is used to configure action for invocation
type ActionEventRouteTimeIntervalDeleteInvocation struct {
	// Pointer to the action
	Action *ActionEventRouteTimeIntervalDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionEventRouteTimeIntervalDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventRouteTimeIntervalDeleteInvocation) SetPathParamInt(param string, value int64) *ActionEventRouteTimeIntervalDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventRouteTimeIntervalDeleteInvocation) SetPathParamString(param string, value string) *ActionEventRouteTimeIntervalDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventRouteTimeIntervalDeleteInvocation) NewMetaInput() *ActionEventRouteTimeIntervalDeleteMetaGlobalInput {
	inv.MetaInput = &ActionEventRouteTimeIntervalDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventRouteTimeIntervalDeleteInvocation) SetMetaInput(input *ActionEventRouteTimeIntervalDeleteMetaGlobalInput) *ActionEventRouteTimeIntervalDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventRouteTimeIntervalDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventRouteTimeIntervalDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventRouteTimeIntervalDeleteInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEventRouteTimeIntervalDeleteInvocation) Call() (*ActionEventRouteTimeIntervalDeleteResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionEventRouteTimeIntervalDeleteInvocation) callAsBody() (*ActionEventRouteTimeIntervalDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionEventRouteTimeIntervalDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionEventRouteTimeIntervalDeleteInvocation) makeAllInputParams() *ActionEventRouteTimeIntervalDeleteRequest {
	return &ActionEventRouteTimeIntervalDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionEventRouteTimeIntervalDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
