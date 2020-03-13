package client

import (
	"strings"
)

// ActionUserNamespaceMapEntryShow is a type for action User_namespace_map.Entry#Show
type ActionUserNamespaceMapEntryShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserNamespaceMapEntryShow(client *Client) *ActionUserNamespaceMapEntryShow {
	return &ActionUserNamespaceMapEntryShow{
		Client: client,
	}
}

// ActionUserNamespaceMapEntryShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserNamespaceMapEntryShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserNamespaceMapEntryShowMetaGlobalInput) SetNo(value bool) *ActionUserNamespaceMapEntryShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserNamespaceMapEntryShowMetaGlobalInput) SetIncludes(value string) *ActionUserNamespaceMapEntryShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNamespaceMapEntryShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceMapEntryShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserNamespaceMapEntryShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNamespaceMapEntryShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionUserNamespaceMapEntryShowOutput is a type for action output parameters
type ActionUserNamespaceMapEntryShowOutput struct {
	Id int64 `json:"id"`
	Kind string `json:"kind"`
	VpsId int64 `json:"vps_id"`
	NsId int64 `json:"ns_id"`
	Count int64 `json:"count"`
}


// Type for action response, including envelope
type ActionUserNamespaceMapEntryShowResponse struct {
	Action *ActionUserNamespaceMapEntryShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Entry *ActionUserNamespaceMapEntryShowOutput `json:"entry"`
	}

	// Action output without the namespace
	Output *ActionUserNamespaceMapEntryShowOutput
}


// Prepare the action for invocation
func (action *ActionUserNamespaceMapEntryShow) Prepare() *ActionUserNamespaceMapEntryShowInvocation {
	return &ActionUserNamespaceMapEntryShowInvocation{
		Action: action,
		Path: "/v6.0/user_namespace_maps/{user_namespace_map_id}/entries/{entry_id}",
	}
}

// ActionUserNamespaceMapEntryShowInvocation is used to configure action for invocation
type ActionUserNamespaceMapEntryShowInvocation struct {
	// Pointer to the action
	Action *ActionUserNamespaceMapEntryShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserNamespaceMapEntryShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserNamespaceMapEntryShowInvocation) SetPathParamInt(param string, value int64) *ActionUserNamespaceMapEntryShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserNamespaceMapEntryShowInvocation) SetPathParamString(param string, value string) *ActionUserNamespaceMapEntryShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserNamespaceMapEntryShowInvocation) NewMetaInput() *ActionUserNamespaceMapEntryShowMetaGlobalInput {
	inv.MetaInput = &ActionUserNamespaceMapEntryShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserNamespaceMapEntryShowInvocation) SetMetaInput(input *ActionUserNamespaceMapEntryShowMetaGlobalInput) *ActionUserNamespaceMapEntryShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserNamespaceMapEntryShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserNamespaceMapEntryShowInvocation) Call() (*ActionUserNamespaceMapEntryShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserNamespaceMapEntryShowInvocation) callAsQuery() (*ActionUserNamespaceMapEntryShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserNamespaceMapEntryShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Entry
	}
	return resp, err
}




func (inv *ActionUserNamespaceMapEntryShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

