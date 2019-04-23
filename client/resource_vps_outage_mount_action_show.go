package client

import (
	"strings"
)

// ActionVpsOutageMountShow is a type for action Vps_outage_mount#Show
type ActionVpsOutageMountShow struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsOutageMountShow(client *Client) *ActionVpsOutageMountShow {
	return &ActionVpsOutageMountShow{
		Client: client,
	}
}

// ActionVpsOutageMountShowMetaGlobalInput is a type for action global meta input parameters
type ActionVpsOutageMountShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsOutageMountShowMetaGlobalInput) SetNo(value bool) *ActionVpsOutageMountShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsOutageMountShowMetaGlobalInput) SetIncludes(value string) *ActionVpsOutageMountShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsOutageMountShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsOutageMountShowMetaGlobalInput) SelectParameters(params ...string) *ActionVpsOutageMountShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsOutageMountShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionVpsOutageMountShowOutput is a type for action output parameters
type ActionVpsOutageMountShowOutput struct {
	Id int64 `json:"id"`
	VpsOutage *ActionVpsOutageShowOutput `json:"vps_outage"`
	Mount *ActionVpsMountShowOutput `json:"mount"`
}


// Type for action response, including envelope
type ActionVpsOutageMountShowResponse struct {
	Action *ActionVpsOutageMountShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		VpsOutageMount *ActionVpsOutageMountShowOutput `json:"vps_outage_mount"`
	}

	// Action output without the namespace
	Output *ActionVpsOutageMountShowOutput
}


// Prepare the action for invocation
func (action *ActionVpsOutageMountShow) Prepare() *ActionVpsOutageMountShowInvocation {
	return &ActionVpsOutageMountShowInvocation{
		Action: action,
		Path: "/v5.0/vps_outage_mounts/:vps_outage_mount_id",
	}
}

// ActionVpsOutageMountShowInvocation is used to configure action for invocation
type ActionVpsOutageMountShowInvocation struct {
	// Pointer to the action
	Action *ActionVpsOutageMountShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionVpsOutageMountShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsOutageMountShowInvocation) SetPathParamInt(param string, value int64) *ActionVpsOutageMountShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsOutageMountShowInvocation) SetPathParamString(param string, value string) *ActionVpsOutageMountShowInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsOutageMountShowInvocation) SetMetaInput(input *ActionVpsOutageMountShowMetaGlobalInput) *ActionVpsOutageMountShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsOutageMountShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsOutageMountShowInvocation) Call() (*ActionVpsOutageMountShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsOutageMountShowInvocation) callAsQuery() (*ActionVpsOutageMountShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsOutageMountShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.VpsOutageMount
	}
	return resp, err
}




func (inv *ActionVpsOutageMountShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

