package client

import (
	"strings"
)

// ActionUserClusterResourceShow is a type for action User.Cluster_resource#Show
type ActionUserClusterResourceShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserClusterResourceShow(client *Client) *ActionUserClusterResourceShow {
	return &ActionUserClusterResourceShow{
		Client: client,
	}
}

// ActionUserClusterResourceShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserClusterResourceShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserClusterResourceShowMetaGlobalInput) SetNo(value bool) *ActionUserClusterResourceShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserClusterResourceShowMetaGlobalInput) SetIncludes(value string) *ActionUserClusterResourceShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserClusterResourceShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserClusterResourceShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserClusterResourceShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserClusterResourceShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionUserClusterResourceShowOutput is a type for action output parameters
type ActionUserClusterResourceShowOutput struct {
	Id int64 `json:"id"`
	Environment *ActionEnvironmentShowOutput `json:"environment"`
	ClusterResource *ActionClusterResourceShowOutput `json:"cluster_resource"`
	Value int64 `json:"value"`
	Used int64 `json:"used"`
	Free int64 `json:"free"`
}


// Type for action response, including envelope
type ActionUserClusterResourceShowResponse struct {
	Action *ActionUserClusterResourceShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ClusterResource *ActionUserClusterResourceShowOutput `json:"cluster_resource"`
	}

	// Action output without the namespace
	Output *ActionUserClusterResourceShowOutput
}


// Prepare the action for invocation
func (action *ActionUserClusterResourceShow) Prepare() *ActionUserClusterResourceShowInvocation {
	return &ActionUserClusterResourceShowInvocation{
		Action: action,
		Path: "/v5.0/users/{user_id}/cluster_resources/{cluster_resource_id}",
	}
}

// ActionUserClusterResourceShowInvocation is used to configure action for invocation
type ActionUserClusterResourceShowInvocation struct {
	// Pointer to the action
	Action *ActionUserClusterResourceShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserClusterResourceShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserClusterResourceShowInvocation) SetPathParamInt(param string, value int64) *ActionUserClusterResourceShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserClusterResourceShowInvocation) SetPathParamString(param string, value string) *ActionUserClusterResourceShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserClusterResourceShowInvocation) NewMetaInput() *ActionUserClusterResourceShowMetaGlobalInput {
	inv.MetaInput = &ActionUserClusterResourceShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserClusterResourceShowInvocation) SetMetaInput(input *ActionUserClusterResourceShowMetaGlobalInput) *ActionUserClusterResourceShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserClusterResourceShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserClusterResourceShowInvocation) Call() (*ActionUserClusterResourceShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserClusterResourceShowInvocation) callAsQuery() (*ActionUserClusterResourceShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserClusterResourceShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.ClusterResource
	}
	return resp, err
}




func (inv *ActionUserClusterResourceShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

