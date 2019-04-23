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
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMonitoredEventAcknowledgeMetaGlobalInput) SetIncludes(value string) *ActionMonitoredEventAcknowledgeMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
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

func (in *ActionMonitoredEventAcknowledgeInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMonitoredEventAcknowledgeRequest is a type for the entire action request
type ActionMonitoredEventAcknowledgeRequest struct {
	MonitoredEvent map[string]interface{} `json:"monitored_event"`
	Meta map[string]interface{} `json:"_meta"`
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
		Path: "/v5.0/monitored_events/:monitored_event_id/acknowledge",
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
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
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
		Meta: inv.makeMetaInputParams(),
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
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
