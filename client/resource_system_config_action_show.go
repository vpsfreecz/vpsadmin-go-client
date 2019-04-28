package client

import (
	"strings"
)

// ActionSystemConfigShow is a type for action System_config#Show
type ActionSystemConfigShow struct {
	// Pointer to client
	Client *Client
}

func NewActionSystemConfigShow(client *Client) *ActionSystemConfigShow {
	return &ActionSystemConfigShow{
		Client: client,
	}
}

// ActionSystemConfigShowMetaGlobalInput is a type for action global meta input parameters
type ActionSystemConfigShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSystemConfigShowMetaGlobalInput) SetNo(value bool) *ActionSystemConfigShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSystemConfigShowMetaGlobalInput) SetIncludes(value string) *ActionSystemConfigShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionSystemConfigShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSystemConfigShowMetaGlobalInput) SelectParameters(params ...string) *ActionSystemConfigShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSystemConfigShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionSystemConfigShowOutput is a type for action output parameters
type ActionSystemConfigShowOutput struct {
	Category string `json:"category"`
	Name string `json:"name"`
	Type string `json:"type"`
	Label string `json:"label"`
	Description string `json:"description"`
	MinUserLevel int64 `json:"min_user_level"`
}


// Type for action response, including envelope
type ActionSystemConfigShowResponse struct {
	Action *ActionSystemConfigShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SystemConfig *ActionSystemConfigShowOutput `json:"system_config"`
	}

	// Action output without the namespace
	Output *ActionSystemConfigShowOutput
}


// Prepare the action for invocation
func (action *ActionSystemConfigShow) Prepare() *ActionSystemConfigShowInvocation {
	return &ActionSystemConfigShowInvocation{
		Action: action,
		Path: "/v5.0/system_configs/:category/:name",
	}
}

// ActionSystemConfigShowInvocation is used to configure action for invocation
type ActionSystemConfigShowInvocation struct {
	// Pointer to the action
	Action *ActionSystemConfigShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionSystemConfigShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSystemConfigShowInvocation) SetPathParamInt(param string, value int64) *ActionSystemConfigShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSystemConfigShowInvocation) SetPathParamString(param string, value string) *ActionSystemConfigShowInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSystemConfigShowInvocation) NewMetaInput() *ActionSystemConfigShowMetaGlobalInput {
	inv.MetaInput = &ActionSystemConfigShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSystemConfigShowInvocation) SetMetaInput(input *ActionSystemConfigShowMetaGlobalInput) *ActionSystemConfigShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSystemConfigShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSystemConfigShowInvocation) Call() (*ActionSystemConfigShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionSystemConfigShowInvocation) callAsQuery() (*ActionSystemConfigShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionSystemConfigShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SystemConfig
	}
	return resp, err
}




func (inv *ActionSystemConfigShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

