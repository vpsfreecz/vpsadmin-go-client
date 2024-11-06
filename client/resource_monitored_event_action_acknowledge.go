package client

import (
	"strings"
)

// ActionMonitoredEventAcknowledge is a type for action Monitored_event#Acknowledge
type ActionMonitoredEventAcknowledge struct {
	// Pointer to client
	Client *Client
}

func NewActionMonitoredEventAcknowledge(client *Client) *ActionMonitoredEventAcknowledge {
	return &ActionMonitoredEventAcknowledge{
		Client: client,
	}
}

// ActionMonitoredEventAcknowledgeMetaGlobalInput is a type for action global meta input parameters
type ActionMonitoredEventAcknowledgeMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMonitoredEventAcknowledgeMetaGlobalInput) SetIncludes(value string) *ActionMonitoredEventAcknowledgeMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMonitoredEventAcknowledgeMetaGlobalInput) SetNo(value bool) *ActionMonitoredEventAcknowledgeMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMonitoredEventAcknowledgeMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMonitoredEventAcknowledgeMetaGlobalInput) SelectParameters(params ...string) *ActionMonitoredEventAcknowledgeMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMonitoredEventAcknowledgeMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMonitoredEventAcknowledgeInput is a type for action input parameters
type ActionMonitoredEventAcknowledgeInput struct {
	Until string `json:"until"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetUntil sets parameter Until to value and selects it for sending
func (in *ActionMonitoredEventAcknowledgeInput) SetUntil(value string) *ActionMonitoredEventAcknowledgeInput {
	in.Until = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Until"] = nil
	return in
}

// SelectParameters sets parameters from ActionMonitoredEventAcknowledgeInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMonitoredEventAcknowledgeInput) SelectParameters(params ...string) *ActionMonitoredEventAcknowledgeInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionMonitoredEventAcknowledgeInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionMonitoredEventAcknowledgeInput) UnselectParameters(params ...string) *ActionMonitoredEventAcknowledgeInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionMonitoredEventAcknowledgeInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMonitoredEventAcknowledgeRequest is a type for the entire action request
type ActionMonitoredEventAcknowledgeRequest struct {
	MonitoredEvent map[string]interface{} `json:"monitored_event"`
	Meta           map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionMonitoredEventAcknowledgeResponse struct {
	Action *ActionMonitoredEventAcknowledge `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionMonitoredEventAcknowledge) Prepare() *ActionMonitoredEventAcknowledgeInvocation {
	return &ActionMonitoredEventAcknowledgeInvocation{
		Action: action,
		Path:   "/v7.0/monitored_events/{monitored_event_id}/acknowledge",
	}
}

// ActionMonitoredEventAcknowledgeInvocation is used to configure action for invocation
type ActionMonitoredEventAcknowledgeInvocation struct {
	// Pointer to the action
	Action *ActionMonitoredEventAcknowledge

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMonitoredEventAcknowledgeInput
	// Global meta input parameters
	MetaInput *ActionMonitoredEventAcknowledgeMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMonitoredEventAcknowledgeInvocation) SetPathParamInt(param string, value int64) *ActionMonitoredEventAcknowledgeInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMonitoredEventAcknowledgeInvocation) SetPathParamString(param string, value string) *ActionMonitoredEventAcknowledgeInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMonitoredEventAcknowledgeInvocation) NewInput() *ActionMonitoredEventAcknowledgeInput {
	inv.Input = &ActionMonitoredEventAcknowledgeInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMonitoredEventAcknowledgeInvocation) SetInput(input *ActionMonitoredEventAcknowledgeInput) *ActionMonitoredEventAcknowledgeInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMonitoredEventAcknowledgeInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionMonitoredEventAcknowledgeInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMonitoredEventAcknowledgeInvocation) NewMetaInput() *ActionMonitoredEventAcknowledgeMetaGlobalInput {
	inv.MetaInput = &ActionMonitoredEventAcknowledgeMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMonitoredEventAcknowledgeInvocation) SetMetaInput(input *ActionMonitoredEventAcknowledgeMetaGlobalInput) *ActionMonitoredEventAcknowledgeInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMonitoredEventAcknowledgeInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMonitoredEventAcknowledgeInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMonitoredEventAcknowledgeInvocation) Call() (*ActionMonitoredEventAcknowledgeResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionMonitoredEventAcknowledgeInvocation) callAsBody() (*ActionMonitoredEventAcknowledgeResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMonitoredEventAcknowledgeResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionMonitoredEventAcknowledgeInvocation) makeAllInputParams() *ActionMonitoredEventAcknowledgeRequest {
	return &ActionMonitoredEventAcknowledgeRequest{
		MonitoredEvent: inv.makeInputParams(),
		Meta:           inv.makeMetaInputParams(),
	}
}

func (inv *ActionMonitoredEventAcknowledgeInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Until") {
			ret["until"] = inv.Input.Until
		}
	}

	return ret
}

func (inv *ActionMonitoredEventAcknowledgeInvocation) makeMetaInputParams() map[string]interface{} {
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
