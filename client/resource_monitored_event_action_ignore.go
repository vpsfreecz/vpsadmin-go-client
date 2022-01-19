package client

import (
	"strings"
)

// ActionMonitoredEventIgnore is a type for action Monitored_event#Ignore
type ActionMonitoredEventIgnore struct {
	// Pointer to client
	Client *Client
}

func NewActionMonitoredEventIgnore(client *Client) *ActionMonitoredEventIgnore {
	return &ActionMonitoredEventIgnore{
		Client: client,
	}
}

// ActionMonitoredEventIgnoreMetaGlobalInput is a type for action global meta input parameters
type ActionMonitoredEventIgnoreMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMonitoredEventIgnoreMetaGlobalInput) SetIncludes(value string) *ActionMonitoredEventIgnoreMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionMonitoredEventIgnoreMetaGlobalInput) SetNo(value bool) *ActionMonitoredEventIgnoreMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMonitoredEventIgnoreMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMonitoredEventIgnoreMetaGlobalInput) SelectParameters(params ...string) *ActionMonitoredEventIgnoreMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMonitoredEventIgnoreMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMonitoredEventIgnoreInput is a type for action input parameters
type ActionMonitoredEventIgnoreInput struct {
	Until string `json:"until"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetUntil sets parameter Until to value and selects it for sending
func (in *ActionMonitoredEventIgnoreInput) SetUntil(value string) *ActionMonitoredEventIgnoreInput {
	in.Until = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Until"] = nil
	return in
}

// SelectParameters sets parameters from ActionMonitoredEventIgnoreInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMonitoredEventIgnoreInput) SelectParameters(params ...string) *ActionMonitoredEventIgnoreInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMonitoredEventIgnoreInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMonitoredEventIgnoreRequest is a type for the entire action request
type ActionMonitoredEventIgnoreRequest struct {
	MonitoredEvent map[string]interface{} `json:"monitored_event"`
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionMonitoredEventIgnoreResponse struct {
	Action *ActionMonitoredEventIgnore `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionMonitoredEventIgnore) Prepare() *ActionMonitoredEventIgnoreInvocation {
	return &ActionMonitoredEventIgnoreInvocation{
		Action: action,
		Path: "/v6.0/monitored_events/{monitored_event_id}/ignore",
	}
}

// ActionMonitoredEventIgnoreInvocation is used to configure action for invocation
type ActionMonitoredEventIgnoreInvocation struct {
	// Pointer to the action
	Action *ActionMonitoredEventIgnore

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMonitoredEventIgnoreInput
	// Global meta input parameters
	MetaInput *ActionMonitoredEventIgnoreMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMonitoredEventIgnoreInvocation) SetPathParamInt(param string, value int64) *ActionMonitoredEventIgnoreInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMonitoredEventIgnoreInvocation) SetPathParamString(param string, value string) *ActionMonitoredEventIgnoreInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMonitoredEventIgnoreInvocation) NewInput() *ActionMonitoredEventIgnoreInput {
	inv.Input = &ActionMonitoredEventIgnoreInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMonitoredEventIgnoreInvocation) SetInput(input *ActionMonitoredEventIgnoreInput) *ActionMonitoredEventIgnoreInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMonitoredEventIgnoreInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMonitoredEventIgnoreInvocation) NewMetaInput() *ActionMonitoredEventIgnoreMetaGlobalInput {
	inv.MetaInput = &ActionMonitoredEventIgnoreMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMonitoredEventIgnoreInvocation) SetMetaInput(input *ActionMonitoredEventIgnoreMetaGlobalInput) *ActionMonitoredEventIgnoreInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMonitoredEventIgnoreInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMonitoredEventIgnoreInvocation) Call() (*ActionMonitoredEventIgnoreResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionMonitoredEventIgnoreInvocation) callAsBody() (*ActionMonitoredEventIgnoreResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMonitoredEventIgnoreResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionMonitoredEventIgnoreInvocation) makeAllInputParams() *ActionMonitoredEventIgnoreRequest {
	return &ActionMonitoredEventIgnoreRequest{
		MonitoredEvent: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionMonitoredEventIgnoreInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Until") {
			ret["until"] = inv.Input.Until
		}
	}

	return ret
}

func (inv *ActionMonitoredEventIgnoreInvocation) makeMetaInputParams() map[string]interface{} {
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
