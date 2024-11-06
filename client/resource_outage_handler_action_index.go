package client

import (
	"strings"
)

// ActionOutageHandlerIndex is a type for action Outage.Handler#Index
type ActionOutageHandlerIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageHandlerIndex(client *Client) *ActionOutageHandlerIndex {
	return &ActionOutageHandlerIndex{
		Client: client,
	}
}

// ActionOutageHandlerIndexMetaGlobalInput is a type for action global meta input parameters
type ActionOutageHandlerIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionOutageHandlerIndexMetaGlobalInput) SetCount(value bool) *ActionOutageHandlerIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageHandlerIndexMetaGlobalInput) SetIncludes(value string) *ActionOutageHandlerIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageHandlerIndexMetaGlobalInput) SetNo(value bool) *ActionOutageHandlerIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageHandlerIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageHandlerIndexMetaGlobalInput) SelectParameters(params ...string) *ActionOutageHandlerIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageHandlerIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageHandlerIndexInput is a type for action input parameters
type ActionOutageHandlerIndexInput struct {
	FromId int64 `json:"from_id"`
	Limit  int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionOutageHandlerIndexInput) SetFromId(value int64) *ActionOutageHandlerIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionOutageHandlerIndexInput) SetLimit(value int64) *ActionOutageHandlerIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageHandlerIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageHandlerIndexInput) SelectParameters(params ...string) *ActionOutageHandlerIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionOutageHandlerIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOutageHandlerIndexInput) UnselectParameters(params ...string) *ActionOutageHandlerIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionOutageHandlerIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageHandlerIndexOutput is a type for action output parameters
type ActionOutageHandlerIndexOutput struct {
	FullName string                `json:"full_name"`
	Id       int64                 `json:"id"`
	Note     string                `json:"note"`
	User     *ActionUserShowOutput `json:"user"`
}

// Type for action response, including envelope
type ActionOutageHandlerIndexResponse struct {
	Action *ActionOutageHandlerIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Handlers []*ActionOutageHandlerIndexOutput `json:"handlers"`
	}

	// Action output without the namespace
	Output []*ActionOutageHandlerIndexOutput
}

// Prepare the action for invocation
func (action *ActionOutageHandlerIndex) Prepare() *ActionOutageHandlerIndexInvocation {
	return &ActionOutageHandlerIndexInvocation{
		Action: action,
		Path:   "/v7.0/outages/{outage_id}/handlers",
	}
}

// ActionOutageHandlerIndexInvocation is used to configure action for invocation
type ActionOutageHandlerIndexInvocation struct {
	// Pointer to the action
	Action *ActionOutageHandlerIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOutageHandlerIndexInput
	// Global meta input parameters
	MetaInput *ActionOutageHandlerIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOutageHandlerIndexInvocation) SetPathParamInt(param string, value int64) *ActionOutageHandlerIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOutageHandlerIndexInvocation) SetPathParamString(param string, value string) *ActionOutageHandlerIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOutageHandlerIndexInvocation) NewInput() *ActionOutageHandlerIndexInput {
	inv.Input = &ActionOutageHandlerIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOutageHandlerIndexInvocation) SetInput(input *ActionOutageHandlerIndexInput) *ActionOutageHandlerIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOutageHandlerIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOutageHandlerIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOutageHandlerIndexInvocation) NewMetaInput() *ActionOutageHandlerIndexMetaGlobalInput {
	inv.MetaInput = &ActionOutageHandlerIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageHandlerIndexInvocation) SetMetaInput(input *ActionOutageHandlerIndexMetaGlobalInput) *ActionOutageHandlerIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageHandlerIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOutageHandlerIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOutageHandlerIndexInvocation) Call() (*ActionOutageHandlerIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOutageHandlerIndexInvocation) callAsQuery() (*ActionOutageHandlerIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOutageHandlerIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Handlers
	}
	return resp, err
}

func (inv *ActionOutageHandlerIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["handler[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["handler[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionOutageHandlerIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
