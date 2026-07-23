package client

import ()

// ActionEventTypeIndex is a type for action Event_type#Index
type ActionEventTypeIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionEventTypeIndex(client *Client) *ActionEventTypeIndex {
	return &ActionEventTypeIndex{
		Client: client,
	}
}

// ActionEventTypeIndexMetaGlobalInput is a type for action global meta input parameters
type ActionEventTypeIndexMetaGlobalInput struct {
	Count bool "json:\"count\""
	No    bool "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionEventTypeIndexMetaGlobalInput) SetCount(value bool) *ActionEventTypeIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventTypeIndexMetaGlobalInput) SetNo(value bool) *ActionEventTypeIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventTypeIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventTypeIndexMetaGlobalInput) SelectParameters(params ...string) *ActionEventTypeIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventTypeIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventTypeIndexInput is a type for action input parameters
type ActionEventTypeIndexInput struct {
	FromId int64 "json:\"from_id\""
	Limit  int64 "json:\"limit\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionEventTypeIndexInput) SetFromId(value int64) *ActionEventTypeIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionEventTypeIndexInput) SetLimit(value int64) *ActionEventTypeIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventTypeIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventTypeIndexInput) SelectParameters(params ...string) *ActionEventTypeIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionEventTypeIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionEventTypeIndexInput) UnselectParameters(params ...string) *ActionEventTypeIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionEventTypeIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventTypeIndexOutput is a type for action output parameters
type ActionEventTypeIndexOutput struct {
	Category            string "json:\"category\""
	DefaultRouted       bool   "json:\"default_routed\""
	Label               string "json:\"label\""
	Name                string "json:\"name\""
	Severity            string "json:\"severity\""
	SeverityDescription string "json:\"severity_description\""
	Template            string "json:\"template\""
}

// Type for action response, including envelope
type ActionEventTypeIndexResponse struct {
	Action *ActionEventTypeIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		EventTypes []*ActionEventTypeIndexOutput "json:\"event_types\""
	}

	// Action output without the namespace
	Output []*ActionEventTypeIndexOutput
}

// Prepare the action for invocation
func (action *ActionEventTypeIndex) Prepare() *ActionEventTypeIndexInvocation {
	return &ActionEventTypeIndexInvocation{
		Action: action,
		Path:   "/v7.0/event_types",
	}
}

// ActionEventTypeIndexInvocation is used to configure action for invocation
type ActionEventTypeIndexInvocation struct {
	// Pointer to the action
	Action *ActionEventTypeIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEventTypeIndexInput
	// Global meta input parameters
	MetaInput *ActionEventTypeIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEventTypeIndexInvocation) NewInput() *ActionEventTypeIndexInput {
	inv.Input = &ActionEventTypeIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEventTypeIndexInvocation) SetInput(input *ActionEventTypeIndexInput) *ActionEventTypeIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEventTypeIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionEventTypeIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventTypeIndexInvocation) NewMetaInput() *ActionEventTypeIndexMetaGlobalInput {
	inv.MetaInput = &ActionEventTypeIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventTypeIndexInvocation) SetMetaInput(input *ActionEventTypeIndexMetaGlobalInput) *ActionEventTypeIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventTypeIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventTypeIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventTypeIndexInvocation) validate() error {
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
func (inv *ActionEventTypeIndexInvocation) Call() (*ActionEventTypeIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionEventTypeIndexInvocation) callAsQuery() (*ActionEventTypeIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionEventTypeIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.EventTypes
	}
	return resp, err
}

func (inv *ActionEventTypeIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["event_type[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["event_type[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionEventTypeIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
