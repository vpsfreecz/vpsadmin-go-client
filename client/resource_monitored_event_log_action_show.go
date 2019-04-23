package client

import (
	"strings"
)

// ActionMonitoredEventLogShow is a type for action Monitored_event.Log#Show
type ActionMonitoredEventLogShow struct {
	// Pointer to client
	Client *Client
}

func NewActionMonitoredEventLogShow(client *Client) *ActionMonitoredEventLogShow {
	return &ActionMonitoredEventLogShow{
		Client: client,
	}
}

// ActionMonitoredEventLogShowMetaGlobalInput is a type for action global meta input parameters
type ActionMonitoredEventLogShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMonitoredEventLogShowMetaGlobalInput) SetNo(value bool) *ActionMonitoredEventLogShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMonitoredEventLogShowMetaGlobalInput) SetIncludes(value string) *ActionMonitoredEventLogShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionMonitoredEventLogShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMonitoredEventLogShowMetaGlobalInput) SelectParameters(params ...string) *ActionMonitoredEventLogShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMonitoredEventLogShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionMonitoredEventLogShowOutput is a type for action output parameters
type ActionMonitoredEventLogShowOutput struct {
	Id int64 `json:"id"`
	Passed bool `json:"passed"`
	CreatedAt string `json:"created_at"`
}


// Type for action response, including envelope
type ActionMonitoredEventLogShowResponse struct {
	Action *ActionMonitoredEventLogShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Log *ActionMonitoredEventLogShowOutput `json:"log"`
	}

	// Action output without the namespace
	Output *ActionMonitoredEventLogShowOutput
}


// Prepare the action for invocation
func (action *ActionMonitoredEventLogShow) Prepare() *ActionMonitoredEventLogShowInvocation {
	return &ActionMonitoredEventLogShowInvocation{
		Action: action,
		Path: "/v5.0/monitored_events/:monitored_event_id/logs/:log_id",
	}
}

// ActionMonitoredEventLogShowInvocation is used to configure action for invocation
type ActionMonitoredEventLogShowInvocation struct {
	// Pointer to the action
	Action *ActionMonitoredEventLogShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMonitoredEventLogShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMonitoredEventLogShowInvocation) SetPathParamInt(param string, value int64) *ActionMonitoredEventLogShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMonitoredEventLogShowInvocation) SetPathParamString(param string, value string) *ActionMonitoredEventLogShowInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMonitoredEventLogShowInvocation) SetMetaInput(input *ActionMonitoredEventLogShowMetaGlobalInput) *ActionMonitoredEventLogShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMonitoredEventLogShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMonitoredEventLogShowInvocation) Call() (*ActionMonitoredEventLogShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMonitoredEventLogShowInvocation) callAsQuery() (*ActionMonitoredEventLogShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMonitoredEventLogShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Log
	}
	return resp, err
}




func (inv *ActionMonitoredEventLogShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

