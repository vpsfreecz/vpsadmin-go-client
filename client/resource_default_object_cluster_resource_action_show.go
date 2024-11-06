package client

import (
	"strings"
)

// ActionDefaultObjectClusterResourceShow is a type for action Default_object_cluster_resource#Show
type ActionDefaultObjectClusterResourceShow struct {
	// Pointer to client
	Client *Client
}

func NewActionDefaultObjectClusterResourceShow(client *Client) *ActionDefaultObjectClusterResourceShow {
	return &ActionDefaultObjectClusterResourceShow{
		Client: client,
	}
}

// ActionDefaultObjectClusterResourceShowMetaGlobalInput is a type for action global meta input parameters
type ActionDefaultObjectClusterResourceShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDefaultObjectClusterResourceShowMetaGlobalInput) SetIncludes(value string) *ActionDefaultObjectClusterResourceShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDefaultObjectClusterResourceShowMetaGlobalInput) SetNo(value bool) *ActionDefaultObjectClusterResourceShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDefaultObjectClusterResourceShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDefaultObjectClusterResourceShowMetaGlobalInput) SelectParameters(params ...string) *ActionDefaultObjectClusterResourceShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDefaultObjectClusterResourceShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDefaultObjectClusterResourceShowOutput is a type for action output parameters
type ActionDefaultObjectClusterResourceShowOutput struct {
	ClassName       string                           `json:"class_name"`
	ClusterResource *ActionClusterResourceShowOutput `json:"cluster_resource"`
	Environment     *ActionEnvironmentShowOutput     `json:"environment"`
	Id              int64                            `json:"id"`
	Value           int64                            `json:"value"`
}

// Type for action response, including envelope
type ActionDefaultObjectClusterResourceShowResponse struct {
	Action *ActionDefaultObjectClusterResourceShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DefaultObjectClusterResource *ActionDefaultObjectClusterResourceShowOutput `json:"default_object_cluster_resource"`
	}

	// Action output without the namespace
	Output *ActionDefaultObjectClusterResourceShowOutput
}

// Prepare the action for invocation
func (action *ActionDefaultObjectClusterResourceShow) Prepare() *ActionDefaultObjectClusterResourceShowInvocation {
	return &ActionDefaultObjectClusterResourceShowInvocation{
		Action: action,
		Path:   "/v7.0/default_object_cluster_resources/{default_object_cluster_resource_id}",
	}
}

// ActionDefaultObjectClusterResourceShowInvocation is used to configure action for invocation
type ActionDefaultObjectClusterResourceShowInvocation struct {
	// Pointer to the action
	Action *ActionDefaultObjectClusterResourceShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDefaultObjectClusterResourceShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDefaultObjectClusterResourceShowInvocation) SetPathParamInt(param string, value int64) *ActionDefaultObjectClusterResourceShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDefaultObjectClusterResourceShowInvocation) SetPathParamString(param string, value string) *ActionDefaultObjectClusterResourceShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDefaultObjectClusterResourceShowInvocation) NewMetaInput() *ActionDefaultObjectClusterResourceShowMetaGlobalInput {
	inv.MetaInput = &ActionDefaultObjectClusterResourceShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDefaultObjectClusterResourceShowInvocation) SetMetaInput(input *ActionDefaultObjectClusterResourceShowMetaGlobalInput) *ActionDefaultObjectClusterResourceShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDefaultObjectClusterResourceShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDefaultObjectClusterResourceShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDefaultObjectClusterResourceShowInvocation) Call() (*ActionDefaultObjectClusterResourceShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDefaultObjectClusterResourceShowInvocation) callAsQuery() (*ActionDefaultObjectClusterResourceShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDefaultObjectClusterResourceShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DefaultObjectClusterResource
	}
	return resp, err
}

func (inv *ActionDefaultObjectClusterResourceShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
