package client

import (
	"strings"
)

// ActionClusterResourceShow is a type for action Cluster_resource#Show
type ActionClusterResourceShow struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterResourceShow(client *Client) *ActionClusterResourceShow {
	return &ActionClusterResourceShow{
		Client: client,
	}
}

// ActionClusterResourceShowMetaGlobalInput is a type for action global meta input parameters
type ActionClusterResourceShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionClusterResourceShowMetaGlobalInput) SetIncludes(value string) *ActionClusterResourceShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterResourceShowMetaGlobalInput) SetNo(value bool) *ActionClusterResourceShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourceShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourceShowMetaGlobalInput) SelectParameters(params ...string) *ActionClusterResourceShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterResourceShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterResourceShowOutput is a type for action output parameters
type ActionClusterResourceShowOutput struct {
	Id       int64  `json:"id"`
	Label    string `json:"label"`
	Max      int64  `json:"max"`
	Min      int64  `json:"min"`
	Name     string `json:"name"`
	Stepsize int64  `json:"stepsize"`
}

// Type for action response, including envelope
type ActionClusterResourceShowResponse struct {
	Action *ActionClusterResourceShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ClusterResource *ActionClusterResourceShowOutput `json:"cluster_resource"`
	}

	// Action output without the namespace
	Output *ActionClusterResourceShowOutput
}

// Prepare the action for invocation
func (action *ActionClusterResourceShow) Prepare() *ActionClusterResourceShowInvocation {
	return &ActionClusterResourceShowInvocation{
		Action: action,
		Path:   "/v7.0/cluster_resources/{cluster_resource_id}",
	}
}

// ActionClusterResourceShowInvocation is used to configure action for invocation
type ActionClusterResourceShowInvocation struct {
	// Pointer to the action
	Action *ActionClusterResourceShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionClusterResourceShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionClusterResourceShowInvocation) SetPathParamInt(param string, value int64) *ActionClusterResourceShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionClusterResourceShowInvocation) SetPathParamString(param string, value string) *ActionClusterResourceShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionClusterResourceShowInvocation) NewMetaInput() *ActionClusterResourceShowMetaGlobalInput {
	inv.MetaInput = &ActionClusterResourceShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterResourceShowInvocation) SetMetaInput(input *ActionClusterResourceShowMetaGlobalInput) *ActionClusterResourceShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterResourceShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionClusterResourceShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionClusterResourceShowInvocation) Call() (*ActionClusterResourceShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionClusterResourceShowInvocation) callAsQuery() (*ActionClusterResourceShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionClusterResourceShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.ClusterResource
	}
	return resp, err
}

func (inv *ActionClusterResourceShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
