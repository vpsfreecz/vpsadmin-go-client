package client

import (
	"strings"
)

// ActionVpsConsoleTokenShow is a type for action Vps.Console_token#Show
type ActionVpsConsoleTokenShow struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsConsoleTokenShow(client *Client) *ActionVpsConsoleTokenShow {
	return &ActionVpsConsoleTokenShow{
		Client: client,
	}
}

// ActionVpsConsoleTokenShowMetaGlobalInput is a type for action global meta input parameters
type ActionVpsConsoleTokenShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsConsoleTokenShowMetaGlobalInput) SetNo(value bool) *ActionVpsConsoleTokenShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsConsoleTokenShowMetaGlobalInput) SetIncludes(value string) *ActionVpsConsoleTokenShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsConsoleTokenShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsConsoleTokenShowMetaGlobalInput) SelectParameters(params ...string) *ActionVpsConsoleTokenShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsConsoleTokenShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionVpsConsoleTokenShowOutput is a type for action output parameters
type ActionVpsConsoleTokenShowOutput struct {
	Token string `json:"token"`
	Expiration string `json:"expiration"`
}


// Type for action response, including envelope
type ActionVpsConsoleTokenShowResponse struct {
	Action *ActionVpsConsoleTokenShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ConsoleToken *ActionVpsConsoleTokenShowOutput `json:"console_token"`
	}

	// Action output without the namespace
	Output *ActionVpsConsoleTokenShowOutput
}


// Prepare the action for invocation
func (action *ActionVpsConsoleTokenShow) Prepare() *ActionVpsConsoleTokenShowInvocation {
	return &ActionVpsConsoleTokenShowInvocation{
		Action: action,
		Path: "/v5.0/vpses/:vps_id/console_token",
	}
}

// ActionVpsConsoleTokenShowInvocation is used to configure action for invocation
type ActionVpsConsoleTokenShowInvocation struct {
	// Pointer to the action
	Action *ActionVpsConsoleTokenShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionVpsConsoleTokenShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsConsoleTokenShowInvocation) SetPathParamInt(param string, value int64) *ActionVpsConsoleTokenShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsConsoleTokenShowInvocation) SetPathParamString(param string, value string) *ActionVpsConsoleTokenShowInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsConsoleTokenShowInvocation) SetMetaInput(input *ActionVpsConsoleTokenShowMetaGlobalInput) *ActionVpsConsoleTokenShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsConsoleTokenShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsConsoleTokenShowInvocation) Call() (*ActionVpsConsoleTokenShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsConsoleTokenShowInvocation) callAsQuery() (*ActionVpsConsoleTokenShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsConsoleTokenShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.ConsoleToken
	}
	return resp, err
}




func (inv *ActionVpsConsoleTokenShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

