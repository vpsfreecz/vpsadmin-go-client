package client

import (
	"strings"
)

// ActionUserClusterResourcePackageItemShow is a type for action User_cluster_resource_package.Item#Show
type ActionUserClusterResourcePackageItemShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserClusterResourcePackageItemShow(client *Client) *ActionUserClusterResourcePackageItemShow {
	return &ActionUserClusterResourcePackageItemShow{
		Client: client,
	}
}

// ActionUserClusterResourcePackageItemShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserClusterResourcePackageItemShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserClusterResourcePackageItemShowMetaGlobalInput) SetIncludes(value string) *ActionUserClusterResourcePackageItemShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserClusterResourcePackageItemShowMetaGlobalInput) SetNo(value bool) *ActionUserClusterResourcePackageItemShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserClusterResourcePackageItemShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserClusterResourcePackageItemShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserClusterResourcePackageItemShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserClusterResourcePackageItemShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserClusterResourcePackageItemShowOutput is a type for action output parameters
type ActionUserClusterResourcePackageItemShowOutput struct {
	ClusterResource *ActionClusterResourceShowOutput `json:"cluster_resource"`
	Id              int64                            `json:"id"`
	Value           int64                            `json:"value"`
}

// Type for action response, including envelope
type ActionUserClusterResourcePackageItemShowResponse struct {
	Action *ActionUserClusterResourcePackageItemShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Item *ActionUserClusterResourcePackageItemShowOutput `json:"item"`
	}

	// Action output without the namespace
	Output *ActionUserClusterResourcePackageItemShowOutput
}

// Prepare the action for invocation
func (action *ActionUserClusterResourcePackageItemShow) Prepare() *ActionUserClusterResourcePackageItemShowInvocation {
	return &ActionUserClusterResourcePackageItemShowInvocation{
		Action: action,
		Path:   "/v7.0/user_cluster_resource_packages/{user_cluster_resource_package_id}/items/{item_id}",
	}
}

// ActionUserClusterResourcePackageItemShowInvocation is used to configure action for invocation
type ActionUserClusterResourcePackageItemShowInvocation struct {
	// Pointer to the action
	Action *ActionUserClusterResourcePackageItemShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserClusterResourcePackageItemShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserClusterResourcePackageItemShowInvocation) SetPathParamInt(param string, value int64) *ActionUserClusterResourcePackageItemShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserClusterResourcePackageItemShowInvocation) SetPathParamString(param string, value string) *ActionUserClusterResourcePackageItemShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserClusterResourcePackageItemShowInvocation) NewMetaInput() *ActionUserClusterResourcePackageItemShowMetaGlobalInput {
	inv.MetaInput = &ActionUserClusterResourcePackageItemShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserClusterResourcePackageItemShowInvocation) SetMetaInput(input *ActionUserClusterResourcePackageItemShowMetaGlobalInput) *ActionUserClusterResourcePackageItemShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserClusterResourcePackageItemShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserClusterResourcePackageItemShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserClusterResourcePackageItemShowInvocation) Call() (*ActionUserClusterResourcePackageItemShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserClusterResourcePackageItemShowInvocation) callAsQuery() (*ActionUserClusterResourcePackageItemShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserClusterResourcePackageItemShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Item
	}
	return resp, err
}

func (inv *ActionUserClusterResourcePackageItemShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
