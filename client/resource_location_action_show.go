package client

import (
	"strings"
)

// ActionLocationShow is a type for action Location#Show
type ActionLocationShow struct {
	// Pointer to client
	Client *Client
}

func NewActionLocationShow(client *Client) *ActionLocationShow {
	return &ActionLocationShow{
		Client: client,
	}
}

// ActionLocationShowMetaGlobalInput is a type for action global meta input parameters
type ActionLocationShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionLocationShowMetaGlobalInput) SetNo(value bool) *ActionLocationShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionLocationShowMetaGlobalInput) SetIncludes(value string) *ActionLocationShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionLocationShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLocationShowMetaGlobalInput) SelectParameters(params ...string) *ActionLocationShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionLocationShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionLocationShowOutput is a type for action output parameters
type ActionLocationShowOutput struct {
	Id int64 `json:"id"`
	Label string `json:"label"`
	HasIpv6 bool `json:"has_ipv6"`
	VpsOnboot bool `json:"vps_onboot"`
	RemoteConsoleServer string `json:"remote_console_server"`
	Domain string `json:"domain"`
	Environment *ActionEnvironmentShowOutput `json:"environment"`
	MaintenanceLock string `json:"maintenance_lock"`
	MaintenanceLockReason string `json:"maintenance_lock_reason"`
}


// Type for action response, including envelope
type ActionLocationShowResponse struct {
	Action *ActionLocationShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Location *ActionLocationShowOutput `json:"location"`
	}

	// Action output without the namespace
	Output *ActionLocationShowOutput
}


// Prepare the action for invocation
func (action *ActionLocationShow) Prepare() *ActionLocationShowInvocation {
	return &ActionLocationShowInvocation{
		Action: action,
		Path: "/v5.0/locations/:location_id",
	}
}

// ActionLocationShowInvocation is used to configure action for invocation
type ActionLocationShowInvocation struct {
	// Pointer to the action
	Action *ActionLocationShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionLocationShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionLocationShowInvocation) SetPathParamInt(param string, value int64) *ActionLocationShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionLocationShowInvocation) SetPathParamString(param string, value string) *ActionLocationShowInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionLocationShowInvocation) SetMetaInput(input *ActionLocationShowMetaGlobalInput) *ActionLocationShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionLocationShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionLocationShowInvocation) Call() (*ActionLocationShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionLocationShowInvocation) callAsQuery() (*ActionLocationShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionLocationShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Location
	}
	return resp, err
}




func (inv *ActionLocationShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

