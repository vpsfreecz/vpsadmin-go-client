package client

import (
	"strings"
)

// ActionUserNamespaceShow is a type for action User_namespace#Show
type ActionUserNamespaceShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserNamespaceShow(client *Client) *ActionUserNamespaceShow {
	return &ActionUserNamespaceShow{
		Client: client,
	}
}

// ActionUserNamespaceShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserNamespaceShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserNamespaceShowMetaGlobalInput) SetNo(value bool) *ActionUserNamespaceShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserNamespaceShowMetaGlobalInput) SetIncludes(value string) *ActionUserNamespaceShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNamespaceShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserNamespaceShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNamespaceShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionUserNamespaceShowOutput is a type for action output parameters
type ActionUserNamespaceShowOutput struct {
	Id int64 `json:"id"`
	User *ActionUserShowOutput `json:"user"`
	Offset int64 `json:"offset"`
	BlockCount int64 `json:"block_count"`
	Size int64 `json:"size"`
}


// Type for action response, including envelope
type ActionUserNamespaceShowResponse struct {
	Action *ActionUserNamespaceShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserNamespace *ActionUserNamespaceShowOutput `json:"user_namespace"`
	}

	// Action output without the namespace
	Output *ActionUserNamespaceShowOutput
}


// Prepare the action for invocation
func (action *ActionUserNamespaceShow) Prepare() *ActionUserNamespaceShowInvocation {
	return &ActionUserNamespaceShowInvocation{
		Action: action,
		Path: "/v5.0/user_namespaces/:user_namespace_id",
	}
}

// ActionUserNamespaceShowInvocation is used to configure action for invocation
type ActionUserNamespaceShowInvocation struct {
	// Pointer to the action
	Action *ActionUserNamespaceShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserNamespaceShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserNamespaceShowInvocation) SetPathParamInt(param string, value int64) *ActionUserNamespaceShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserNamespaceShowInvocation) SetPathParamString(param string, value string) *ActionUserNamespaceShowInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserNamespaceShowInvocation) SetMetaInput(input *ActionUserNamespaceShowMetaGlobalInput) *ActionUserNamespaceShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserNamespaceShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserNamespaceShowInvocation) Call() (*ActionUserNamespaceShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserNamespaceShowInvocation) callAsQuery() (*ActionUserNamespaceShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserNamespaceShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserNamespace
	}
	return resp, err
}




func (inv *ActionUserNamespaceShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

