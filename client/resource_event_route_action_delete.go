package client

import (
	"net/url"
	"strings"
)

// ActionEventRouteDelete is a type for action Event_route#Delete
type ActionEventRouteDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionEventRouteDelete(client *Client) *ActionEventRouteDelete {
	return &ActionEventRouteDelete{
		Client: client,
	}
}

// ActionEventRouteDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionEventRouteDeleteMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventRouteDeleteMetaGlobalInput) SetIncludes(value string) *ActionEventRouteDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventRouteDeleteMetaGlobalInput) SetNo(value bool) *ActionEventRouteDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionEventRouteDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventRouteDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteDeleteRequest is a type for the entire action request
type ActionEventRouteDeleteRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// Type for action response, including envelope
type ActionEventRouteDeleteResponse struct {
	Action *ActionEventRouteDelete "json:\"-\""
	*Envelope
}

// Prepare the action for invocation
func (action *ActionEventRouteDelete) Prepare() *ActionEventRouteDeleteInvocation {
	return &ActionEventRouteDeleteInvocation{
		Action: action,
		Path:   "/v7.0/event_routes/{event_route_id}",
	}
}

// ActionEventRouteDeleteInvocation is used to configure action for invocation
type ActionEventRouteDeleteInvocation struct {
	// Pointer to the action
	Action *ActionEventRouteDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionEventRouteDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventRouteDeleteInvocation) SetPathParamInt(param string, value int64) *ActionEventRouteDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventRouteDeleteInvocation) SetPathParamString(param string, value string) *ActionEventRouteDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventRouteDeleteInvocation) NewMetaInput() *ActionEventRouteDeleteMetaGlobalInput {
	inv.MetaInput = &ActionEventRouteDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventRouteDeleteInvocation) SetMetaInput(input *ActionEventRouteDeleteMetaGlobalInput) *ActionEventRouteDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventRouteDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventRouteDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventRouteDeleteInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEventRouteDeleteInvocation) Call() (*ActionEventRouteDeleteResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionEventRouteDeleteInvocation) callAsBody() (*ActionEventRouteDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionEventRouteDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionEventRouteDeleteInvocation) makeAllInputParams() *ActionEventRouteDeleteRequest {
	return &ActionEventRouteDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionEventRouteDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
