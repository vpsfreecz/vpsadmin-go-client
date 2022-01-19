package client

import (
	"strings"
)

// ActionExportHostShow is a type for action Export.Host#Show
type ActionExportHostShow struct {
	// Pointer to client
	Client *Client
}

func NewActionExportHostShow(client *Client) *ActionExportHostShow {
	return &ActionExportHostShow{
		Client: client,
	}
}

// ActionExportHostShowMetaGlobalInput is a type for action global meta input parameters
type ActionExportHostShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionExportHostShowMetaGlobalInput) SetIncludes(value string) *ActionExportHostShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionExportHostShowMetaGlobalInput) SetNo(value bool) *ActionExportHostShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionExportHostShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionExportHostShowMetaGlobalInput) SelectParameters(params ...string) *ActionExportHostShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionExportHostShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionExportHostShowOutput is a type for action output parameters
type ActionExportHostShowOutput struct {
	Id int64 `json:"id"`
	IpAddress *ActionIpAddressShowOutput `json:"ip_address"`
	RootSquash bool `json:"root_squash"`
	Rw bool `json:"rw"`
	SubtreeCheck bool `json:"subtree_check"`
	Sync bool `json:"sync"`
}


// Type for action response, including envelope
type ActionExportHostShowResponse struct {
	Action *ActionExportHostShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Host *ActionExportHostShowOutput `json:"host"`
	}

	// Action output without the namespace
	Output *ActionExportHostShowOutput
}


// Prepare the action for invocation
func (action *ActionExportHostShow) Prepare() *ActionExportHostShowInvocation {
	return &ActionExportHostShowInvocation{
		Action: action,
		Path: "/v6.0/exports/{export_id}/hosts/{host_id}",
	}
}

// ActionExportHostShowInvocation is used to configure action for invocation
type ActionExportHostShowInvocation struct {
	// Pointer to the action
	Action *ActionExportHostShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionExportHostShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionExportHostShowInvocation) SetPathParamInt(param string, value int64) *ActionExportHostShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionExportHostShowInvocation) SetPathParamString(param string, value string) *ActionExportHostShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionExportHostShowInvocation) NewMetaInput() *ActionExportHostShowMetaGlobalInput {
	inv.MetaInput = &ActionExportHostShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionExportHostShowInvocation) SetMetaInput(input *ActionExportHostShowMetaGlobalInput) *ActionExportHostShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionExportHostShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionExportHostShowInvocation) Call() (*ActionExportHostShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionExportHostShowInvocation) callAsQuery() (*ActionExportHostShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionExportHostShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Host
	}
	return resp, err
}




func (inv *ActionExportHostShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}

