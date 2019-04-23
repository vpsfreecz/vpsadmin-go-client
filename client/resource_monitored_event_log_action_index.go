package client

import (
	"strings"
)

// ActionMonitoredEventLogIndex is a type for action Monitored_event.Log#Index
type ActionMonitoredEventLogIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionMonitoredEventLogIndex(client *Client) *ActionMonitoredEventLogIndex {
	return &ActionMonitoredEventLogIndex{
		Client: client,
	}
}

// ActionMonitoredEventLogIndexMetaGlobalInput is a type for action global meta input parameters
type ActionMonitoredEventLogIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMonitoredEventLogIndexMetaGlobalInput) SetNo(value bool) *ActionMonitoredEventLogIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionMonitoredEventLogIndexMetaGlobalInput) SetCount(value bool) *ActionMonitoredEventLogIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMonitoredEventLogIndexMetaGlobalInput) SetIncludes(value string) *ActionMonitoredEventLogIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionMonitoredEventLogIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMonitoredEventLogIndexMetaGlobalInput) SelectParameters(params ...string) *ActionMonitoredEventLogIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMonitoredEventLogIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMonitoredEventLogIndexInput is a type for action input parameters
type ActionMonitoredEventLogIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	Passed bool `json:"passed"`
	Order string `json:"order"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionMonitoredEventLogIndexInput) SetOffset(value int64) *ActionMonitoredEventLogIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionMonitoredEventLogIndexInput) SetLimit(value int64) *ActionMonitoredEventLogIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetPassed sets parameter Passed to value and selects it for sending
func (in *ActionMonitoredEventLogIndexInput) SetPassed(value bool) *ActionMonitoredEventLogIndexInput {
	in.Passed = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Passed"] = nil
	return in
}
// SetOrder sets parameter Order to value and selects it for sending
func (in *ActionMonitoredEventLogIndexInput) SetOrder(value string) *ActionMonitoredEventLogIndexInput {
	in.Order = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Order"] = nil
	return in
}

// SelectParameters sets parameters from ActionMonitoredEventLogIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMonitoredEventLogIndexInput) SelectParameters(params ...string) *ActionMonitoredEventLogIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMonitoredEventLogIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionMonitoredEventLogIndexOutput is a type for action output parameters
type ActionMonitoredEventLogIndexOutput struct {
	Id int64 `json:"id"`
	Passed bool `json:"passed"`
	CreatedAt string `json:"created_at"`
}


// Type for action response, including envelope
type ActionMonitoredEventLogIndexResponse struct {
	Action *ActionMonitoredEventLogIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Logs []*ActionMonitoredEventLogIndexOutput `json:"logs"`
	}

	// Action output without the namespace
	Output []*ActionMonitoredEventLogIndexOutput
}


// Prepare the action for invocation
func (action *ActionMonitoredEventLogIndex) Prepare() *ActionMonitoredEventLogIndexInvocation {
	return &ActionMonitoredEventLogIndexInvocation{
		Action: action,
		Path: "/v5.0/monitored_events/:monitored_event_id/logs",
	}
}

// ActionMonitoredEventLogIndexInvocation is used to configure action for invocation
type ActionMonitoredEventLogIndexInvocation struct {
	// Pointer to the action
	Action *ActionMonitoredEventLogIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMonitoredEventLogIndexInput
	// Global meta input parameters
	MetaInput *ActionMonitoredEventLogIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMonitoredEventLogIndexInvocation) SetPathParamInt(param string, value int64) *ActionMonitoredEventLogIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMonitoredEventLogIndexInvocation) SetPathParamString(param string, value string) *ActionMonitoredEventLogIndexInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetInput provides input parameters to send to the API
func (inv *ActionMonitoredEventLogIndexInvocation) SetInput(input *ActionMonitoredEventLogIndexInput) *ActionMonitoredEventLogIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMonitoredEventLogIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMonitoredEventLogIndexInvocation) SetMetaInput(input *ActionMonitoredEventLogIndexMetaGlobalInput) *ActionMonitoredEventLogIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMonitoredEventLogIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMonitoredEventLogIndexInvocation) Call() (*ActionMonitoredEventLogIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMonitoredEventLogIndexInvocation) callAsQuery() (*ActionMonitoredEventLogIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMonitoredEventLogIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Logs
	}
	return resp, err
}



func (inv *ActionMonitoredEventLogIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["log[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["log[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Passed") {
			ret["log[passed]"] = convertBoolToString(inv.Input.Passed)
		}
		if inv.IsParameterSelected("Order") {
			ret["log[order]"] = inv.Input.Order
		}
	}
}

func (inv *ActionMonitoredEventLogIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

