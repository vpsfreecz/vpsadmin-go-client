package client

import (
	"net/url"
	"strings"
)

// ActionEventRouteMatcherDelete is a type for action Event_route.Matcher#Delete
type ActionEventRouteMatcherDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionEventRouteMatcherDelete(client *Client) *ActionEventRouteMatcherDelete {
	return &ActionEventRouteMatcherDelete{
		Client: client,
	}
}

// ActionEventRouteMatcherDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionEventRouteMatcherDeleteMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventRouteMatcherDeleteMetaGlobalInput) SetIncludes(value string) *ActionEventRouteMatcherDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventRouteMatcherDeleteMetaGlobalInput) SetNo(value bool) *ActionEventRouteMatcherDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteMatcherDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteMatcherDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionEventRouteMatcherDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventRouteMatcherDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteMatcherDeleteRequest is a type for the entire action request
type ActionEventRouteMatcherDeleteRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// Type for action response, including envelope
type ActionEventRouteMatcherDeleteResponse struct {
	Action *ActionEventRouteMatcherDelete "json:\"-\""
	*Envelope
}

// Prepare the action for invocation
func (action *ActionEventRouteMatcherDelete) Prepare() *ActionEventRouteMatcherDeleteInvocation {
	return &ActionEventRouteMatcherDeleteInvocation{
		Action: action,
		Path:   "/v7.0/event_routes/{event_route_id}/matcher/{matcher_id}",
	}
}

// ActionEventRouteMatcherDeleteInvocation is used to configure action for invocation
type ActionEventRouteMatcherDeleteInvocation struct {
	// Pointer to the action
	Action *ActionEventRouteMatcherDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionEventRouteMatcherDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventRouteMatcherDeleteInvocation) SetPathParamInt(param string, value int64) *ActionEventRouteMatcherDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventRouteMatcherDeleteInvocation) SetPathParamString(param string, value string) *ActionEventRouteMatcherDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventRouteMatcherDeleteInvocation) NewMetaInput() *ActionEventRouteMatcherDeleteMetaGlobalInput {
	inv.MetaInput = &ActionEventRouteMatcherDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventRouteMatcherDeleteInvocation) SetMetaInput(input *ActionEventRouteMatcherDeleteMetaGlobalInput) *ActionEventRouteMatcherDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventRouteMatcherDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventRouteMatcherDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventRouteMatcherDeleteInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEventRouteMatcherDeleteInvocation) Call() (*ActionEventRouteMatcherDeleteResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionEventRouteMatcherDeleteInvocation) callAsBody() (*ActionEventRouteMatcherDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionEventRouteMatcherDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionEventRouteMatcherDeleteInvocation) makeAllInputParams() *ActionEventRouteMatcherDeleteRequest {
	return &ActionEventRouteMatcherDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionEventRouteMatcherDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
