package client

import (
	"net/url"
	"strings"
)

// ActionEventRouteMatcherIndex is a type for action Event_route.Matcher#Index
type ActionEventRouteMatcherIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionEventRouteMatcherIndex(client *Client) *ActionEventRouteMatcherIndex {
	return &ActionEventRouteMatcherIndex{
		Client: client,
	}
}

// ActionEventRouteMatcherIndexMetaGlobalInput is a type for action global meta input parameters
type ActionEventRouteMatcherIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionEventRouteMatcherIndexMetaGlobalInput) SetCount(value bool) *ActionEventRouteMatcherIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventRouteMatcherIndexMetaGlobalInput) SetIncludes(value string) *ActionEventRouteMatcherIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventRouteMatcherIndexMetaGlobalInput) SetNo(value bool) *ActionEventRouteMatcherIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteMatcherIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteMatcherIndexMetaGlobalInput) SelectParameters(params ...string) *ActionEventRouteMatcherIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventRouteMatcherIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteMatcherIndexInput is a type for action input parameters
type ActionEventRouteMatcherIndexInput struct {
	FromId int64 "json:\"from_id\""
	Limit  int64 "json:\"limit\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionEventRouteMatcherIndexInput) SetFromId(value int64) *ActionEventRouteMatcherIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionEventRouteMatcherIndexInput) SetLimit(value int64) *ActionEventRouteMatcherIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteMatcherIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteMatcherIndexInput) SelectParameters(params ...string) *ActionEventRouteMatcherIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionEventRouteMatcherIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionEventRouteMatcherIndexInput) UnselectParameters(params ...string) *ActionEventRouteMatcherIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionEventRouteMatcherIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteMatcherIndexOutput is a type for action output parameters
type ActionEventRouteMatcherIndexOutput struct {
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
type ActionEventRouteMatcherIndexResponse struct {
	Action *ActionEventRouteMatcherIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Matchers []*ActionEventRouteMatcherIndexOutput "json:\"matchers\""
	}

	// Action output without the namespace
	Output []*ActionEventRouteMatcherIndexOutput
}

// Prepare the action for invocation
func (action *ActionEventRouteMatcherIndex) Prepare() *ActionEventRouteMatcherIndexInvocation {
	return &ActionEventRouteMatcherIndexInvocation{
		Action: action,
		Path:   "/v7.0/event_routes/{event_route_id}/matcher",
	}
}

// ActionEventRouteMatcherIndexInvocation is used to configure action for invocation
type ActionEventRouteMatcherIndexInvocation struct {
	// Pointer to the action
	Action *ActionEventRouteMatcherIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEventRouteMatcherIndexInput
	// Global meta input parameters
	MetaInput *ActionEventRouteMatcherIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventRouteMatcherIndexInvocation) SetPathParamInt(param string, value int64) *ActionEventRouteMatcherIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventRouteMatcherIndexInvocation) SetPathParamString(param string, value string) *ActionEventRouteMatcherIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEventRouteMatcherIndexInvocation) NewInput() *ActionEventRouteMatcherIndexInput {
	inv.Input = &ActionEventRouteMatcherIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEventRouteMatcherIndexInvocation) SetInput(input *ActionEventRouteMatcherIndexInput) *ActionEventRouteMatcherIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEventRouteMatcherIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionEventRouteMatcherIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventRouteMatcherIndexInvocation) NewMetaInput() *ActionEventRouteMatcherIndexMetaGlobalInput {
	inv.MetaInput = &ActionEventRouteMatcherIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventRouteMatcherIndexInvocation) SetMetaInput(input *ActionEventRouteMatcherIndexMetaGlobalInput) *ActionEventRouteMatcherIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventRouteMatcherIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventRouteMatcherIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventRouteMatcherIndexInvocation) validate() error {
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
func (inv *ActionEventRouteMatcherIndexInvocation) Call() (*ActionEventRouteMatcherIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionEventRouteMatcherIndexInvocation) callAsQuery() (*ActionEventRouteMatcherIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionEventRouteMatcherIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Matchers
	}
	return resp, err
}

func (inv *ActionEventRouteMatcherIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["matcher[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["matcher[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}

	return nil
}

func (inv *ActionEventRouteMatcherIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
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
