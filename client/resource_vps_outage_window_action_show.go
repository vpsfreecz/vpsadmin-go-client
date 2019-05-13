package client

import (
	"strings"
)

// ActionVpsOutageWindowShow is a type for action Vps.Outage_window#Show
type ActionVpsOutageWindowShow struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsOutageWindowShow(client *Client) *ActionVpsOutageWindowShow {
	return &ActionVpsOutageWindowShow{
		Client: client,
	}
}

// ActionVpsOutageWindowShowMetaGlobalInput is a type for action global meta input parameters
type ActionVpsOutageWindowShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsOutageWindowShowMetaGlobalInput) SetNo(value bool) *ActionVpsOutageWindowShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsOutageWindowShowMetaGlobalInput) SetIncludes(value string) *ActionVpsOutageWindowShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsOutageWindowShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsOutageWindowShowMetaGlobalInput) SelectParameters(params ...string) *ActionVpsOutageWindowShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsOutageWindowShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionVpsOutageWindowShowOutput is a type for action output parameters
type ActionVpsOutageWindowShowOutput struct {
	Weekday int64 `json:"weekday"`
	IsOpen bool `json:"is_open"`
	OpensAt int64 `json:"opens_at"`
	ClosesAt int64 `json:"closes_at"`
}


// Type for action response, including envelope
type ActionVpsOutageWindowShowResponse struct {
	Action *ActionVpsOutageWindowShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OutageWindow *ActionVpsOutageWindowShowOutput `json:"outage_window"`
	}

	// Action output without the namespace
	Output *ActionVpsOutageWindowShowOutput
}


// Prepare the action for invocation
func (action *ActionVpsOutageWindowShow) Prepare() *ActionVpsOutageWindowShowInvocation {
	return &ActionVpsOutageWindowShowInvocation{
		Action: action,
		Path: "/v5.0/vpses/{vps_id}/outage_windows/{outage_window_id}",
	}
}

// ActionVpsOutageWindowShowInvocation is used to configure action for invocation
type ActionVpsOutageWindowShowInvocation struct {
	// Pointer to the action
	Action *ActionVpsOutageWindowShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionVpsOutageWindowShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsOutageWindowShowInvocation) SetPathParamInt(param string, value int64) *ActionVpsOutageWindowShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsOutageWindowShowInvocation) SetPathParamString(param string, value string) *ActionVpsOutageWindowShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsOutageWindowShowInvocation) NewMetaInput() *ActionVpsOutageWindowShowMetaGlobalInput {
	inv.MetaInput = &ActionVpsOutageWindowShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsOutageWindowShowInvocation) SetMetaInput(input *ActionVpsOutageWindowShowMetaGlobalInput) *ActionVpsOutageWindowShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsOutageWindowShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsOutageWindowShowInvocation) Call() (*ActionVpsOutageWindowShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsOutageWindowShowInvocation) callAsQuery() (*ActionVpsOutageWindowShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsOutageWindowShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OutageWindow
	}
	return resp, err
}




func (inv *ActionVpsOutageWindowShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

