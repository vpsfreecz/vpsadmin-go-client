package client

import (
	"net/url"
	"strings"
)

// ActionEventTimeIntervalDelete is a type for action Event_time_interval#Delete
type ActionEventTimeIntervalDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionEventTimeIntervalDelete(client *Client) *ActionEventTimeIntervalDelete {
	return &ActionEventTimeIntervalDelete{
		Client: client,
	}
}

// ActionEventTimeIntervalDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionEventTimeIntervalDeleteMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventTimeIntervalDeleteMetaGlobalInput) SetIncludes(value string) *ActionEventTimeIntervalDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventTimeIntervalDeleteMetaGlobalInput) SetNo(value bool) *ActionEventTimeIntervalDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventTimeIntervalDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventTimeIntervalDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionEventTimeIntervalDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventTimeIntervalDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventTimeIntervalDeleteRequest is a type for the entire action request
type ActionEventTimeIntervalDeleteRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// Type for action response, including envelope
type ActionEventTimeIntervalDeleteResponse struct {
	Action *ActionEventTimeIntervalDelete "json:\"-\""
	*Envelope
}

// Prepare the action for invocation
func (action *ActionEventTimeIntervalDelete) Prepare() *ActionEventTimeIntervalDeleteInvocation {
	return &ActionEventTimeIntervalDeleteInvocation{
		Action: action,
		Path:   "/v7.0/event_time_intervals/{event_time_interval_id}",
	}
}

// ActionEventTimeIntervalDeleteInvocation is used to configure action for invocation
type ActionEventTimeIntervalDeleteInvocation struct {
	// Pointer to the action
	Action *ActionEventTimeIntervalDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionEventTimeIntervalDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventTimeIntervalDeleteInvocation) SetPathParamInt(param string, value int64) *ActionEventTimeIntervalDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventTimeIntervalDeleteInvocation) SetPathParamString(param string, value string) *ActionEventTimeIntervalDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventTimeIntervalDeleteInvocation) NewMetaInput() *ActionEventTimeIntervalDeleteMetaGlobalInput {
	inv.MetaInput = &ActionEventTimeIntervalDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventTimeIntervalDeleteInvocation) SetMetaInput(input *ActionEventTimeIntervalDeleteMetaGlobalInput) *ActionEventTimeIntervalDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventTimeIntervalDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventTimeIntervalDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventTimeIntervalDeleteInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEventTimeIntervalDeleteInvocation) Call() (*ActionEventTimeIntervalDeleteResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionEventTimeIntervalDeleteInvocation) callAsBody() (*ActionEventTimeIntervalDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionEventTimeIntervalDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionEventTimeIntervalDeleteInvocation) makeAllInputParams() *ActionEventTimeIntervalDeleteRequest {
	return &ActionEventTimeIntervalDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionEventTimeIntervalDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
