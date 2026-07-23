package client

import (
	"net/url"
	"strings"
)

// ActionEventRouteMatcherShow is a type for action Event_route.Matcher#Show
type ActionEventRouteMatcherShow struct {
	// Pointer to client
	Client *Client
}

func NewActionEventRouteMatcherShow(client *Client) *ActionEventRouteMatcherShow {
	return &ActionEventRouteMatcherShow{
		Client: client,
	}
}

// ActionEventRouteMatcherShowMetaGlobalInput is a type for action global meta input parameters
type ActionEventRouteMatcherShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventRouteMatcherShowMetaGlobalInput) SetIncludes(value string) *ActionEventRouteMatcherShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventRouteMatcherShowMetaGlobalInput) SetNo(value bool) *ActionEventRouteMatcherShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteMatcherShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteMatcherShowMetaGlobalInput) SelectParameters(params ...string) *ActionEventRouteMatcherShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventRouteMatcherShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteMatcherShowOutput is a type for action output parameters
type ActionEventRouteMatcherShowOutput struct {
	CreatedAt string "json:\"created_at\""
	Field     string "json:\"field\""
	FieldType string "json:\"field_type\""
	Id        int64  "json:\"id\""
	Operator  string "json:\"operator\""
	Summary   string "json:\"summary\""
	UpdatedAt string "json:\"updated_at\""
	Value     string "json:\"value\""
}

// Type for action response, including envelope
type ActionEventRouteMatcherShowResponse struct {
	Action *ActionEventRouteMatcherShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Matcher *ActionEventRouteMatcherShowOutput "json:\"matcher\""
	}

	// Action output without the namespace
	Output *ActionEventRouteMatcherShowOutput
}

// Prepare the action for invocation
func (action *ActionEventRouteMatcherShow) Prepare() *ActionEventRouteMatcherShowInvocation {
	return &ActionEventRouteMatcherShowInvocation{
		Action: action,
		Path:   "/v7.0/event_routes/{event_route_id}/matcher/{matcher_id}",
	}
}

// ActionEventRouteMatcherShowInvocation is used to configure action for invocation
type ActionEventRouteMatcherShowInvocation struct {
	// Pointer to the action
	Action *ActionEventRouteMatcherShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionEventRouteMatcherShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventRouteMatcherShowInvocation) SetPathParamInt(param string, value int64) *ActionEventRouteMatcherShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventRouteMatcherShowInvocation) SetPathParamString(param string, value string) *ActionEventRouteMatcherShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventRouteMatcherShowInvocation) NewMetaInput() *ActionEventRouteMatcherShowMetaGlobalInput {
	inv.MetaInput = &ActionEventRouteMatcherShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventRouteMatcherShowInvocation) SetMetaInput(input *ActionEventRouteMatcherShowMetaGlobalInput) *ActionEventRouteMatcherShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventRouteMatcherShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventRouteMatcherShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventRouteMatcherShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEventRouteMatcherShowInvocation) Call() (*ActionEventRouteMatcherShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionEventRouteMatcherShowInvocation) callAsQuery() (*ActionEventRouteMatcherShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionEventRouteMatcherShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Matcher
	}
	return resp, err
}

func (inv *ActionEventRouteMatcherShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
