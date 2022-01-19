package client

import (
	"strings"
)

// ActionUserNamespaceMapShow is a type for action User_namespace_map#Show
type ActionUserNamespaceMapShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserNamespaceMapShow(client *Client) *ActionUserNamespaceMapShow {
	return &ActionUserNamespaceMapShow{
		Client: client,
	}
}

// ActionUserNamespaceMapShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserNamespaceMapShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserNamespaceMapShowMetaGlobalInput) SetIncludes(value string) *ActionUserNamespaceMapShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserNamespaceMapShowMetaGlobalInput) SetNo(value bool) *ActionUserNamespaceMapShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNamespaceMapShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceMapShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserNamespaceMapShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNamespaceMapShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNamespaceMapShowOutput is a type for action output parameters
type ActionUserNamespaceMapShowOutput struct {
	Id            int64                          `json:"id"`
	Label         string                         `json:"label"`
	UserNamespace *ActionUserNamespaceShowOutput `json:"user_namespace"`
}

// Type for action response, including envelope
type ActionUserNamespaceMapShowResponse struct {
	Action *ActionUserNamespaceMapShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserNamespaceMap *ActionUserNamespaceMapShowOutput `json:"user_namespace_map"`
	}

	// Action output without the namespace
	Output *ActionUserNamespaceMapShowOutput
}

// Prepare the action for invocation
func (action *ActionUserNamespaceMapShow) Prepare() *ActionUserNamespaceMapShowInvocation {
	return &ActionUserNamespaceMapShowInvocation{
		Action: action,
		Path:   "/v6.0/user_namespace_maps/{user_namespace_map_id}",
	}
}

// ActionUserNamespaceMapShowInvocation is used to configure action for invocation
type ActionUserNamespaceMapShowInvocation struct {
	// Pointer to the action
	Action *ActionUserNamespaceMapShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserNamespaceMapShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserNamespaceMapShowInvocation) SetPathParamInt(param string, value int64) *ActionUserNamespaceMapShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserNamespaceMapShowInvocation) SetPathParamString(param string, value string) *ActionUserNamespaceMapShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserNamespaceMapShowInvocation) NewMetaInput() *ActionUserNamespaceMapShowMetaGlobalInput {
	inv.MetaInput = &ActionUserNamespaceMapShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserNamespaceMapShowInvocation) SetMetaInput(input *ActionUserNamespaceMapShowMetaGlobalInput) *ActionUserNamespaceMapShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserNamespaceMapShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserNamespaceMapShowInvocation) Call() (*ActionUserNamespaceMapShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserNamespaceMapShowInvocation) callAsQuery() (*ActionUserNamespaceMapShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserNamespaceMapShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserNamespaceMap
	}
	return resp, err
}

func (inv *ActionUserNamespaceMapShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
