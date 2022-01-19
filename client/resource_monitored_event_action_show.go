package client

import (
	"strings"
)

// ActionMonitoredEventShow is a type for action Monitored_event#Show
type ActionMonitoredEventShow struct {
	// Pointer to client
	Client *Client
}

func NewActionMonitoredEventShow(client *Client) *ActionMonitoredEventShow {
	return &ActionMonitoredEventShow{
		Client: client,
	}
}

// ActionMonitoredEventShowMetaGlobalInput is a type for action global meta input parameters
type ActionMonitoredEventShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMonitoredEventShowMetaGlobalInput) SetIncludes(value string) *ActionMonitoredEventShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionMonitoredEventShowMetaGlobalInput) SetNo(value bool) *ActionMonitoredEventShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMonitoredEventShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMonitoredEventShowMetaGlobalInput) SelectParameters(params ...string) *ActionMonitoredEventShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMonitoredEventShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionMonitoredEventShowOutput is a type for action output parameters
type ActionMonitoredEventShowOutput struct {
	CreatedAt string `json:"created_at"`
	Id int64 `json:"id"`
	Issue string `json:"issue"`
	Label string `json:"label"`
	Monitor string `json:"monitor"`
	ObjectId int64 `json:"object_id"`
	ObjectName string `json:"object_name"`
	SavedUntil string `json:"saved_until"`
	State string `json:"state"`
	UpdatedAt string `json:"updated_at"`
	User *ActionUserShowOutput `json:"user"`
}


// Type for action response, including envelope
type ActionMonitoredEventShowResponse struct {
	Action *ActionMonitoredEventShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MonitoredEvent *ActionMonitoredEventShowOutput `json:"monitored_event"`
	}

	// Action output without the namespace
	Output *ActionMonitoredEventShowOutput
}


// Prepare the action for invocation
func (action *ActionMonitoredEventShow) Prepare() *ActionMonitoredEventShowInvocation {
	return &ActionMonitoredEventShowInvocation{
		Action: action,
		Path: "/v6.0/monitored_events/{monitored_event_id}",
	}
}

// ActionMonitoredEventShowInvocation is used to configure action for invocation
type ActionMonitoredEventShowInvocation struct {
	// Pointer to the action
	Action *ActionMonitoredEventShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMonitoredEventShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMonitoredEventShowInvocation) SetPathParamInt(param string, value int64) *ActionMonitoredEventShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMonitoredEventShowInvocation) SetPathParamString(param string, value string) *ActionMonitoredEventShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMonitoredEventShowInvocation) NewMetaInput() *ActionMonitoredEventShowMetaGlobalInput {
	inv.MetaInput = &ActionMonitoredEventShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMonitoredEventShowInvocation) SetMetaInput(input *ActionMonitoredEventShowMetaGlobalInput) *ActionMonitoredEventShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMonitoredEventShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMonitoredEventShowInvocation) Call() (*ActionMonitoredEventShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMonitoredEventShowInvocation) callAsQuery() (*ActionMonitoredEventShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMonitoredEventShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MonitoredEvent
	}
	return resp, err
}




func (inv *ActionMonitoredEventShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}

