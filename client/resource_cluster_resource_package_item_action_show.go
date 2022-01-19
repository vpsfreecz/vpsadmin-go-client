package client

import (
	"strings"
)

// ActionClusterResourcePackageItemShow is a type for action Cluster_resource_package.Item#Show
type ActionClusterResourcePackageItemShow struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterResourcePackageItemShow(client *Client) *ActionClusterResourcePackageItemShow {
	return &ActionClusterResourcePackageItemShow{
		Client: client,
	}
}

// ActionClusterResourcePackageItemShowMetaGlobalInput is a type for action global meta input parameters
type ActionClusterResourcePackageItemShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionClusterResourcePackageItemShowMetaGlobalInput) SetIncludes(value string) *ActionClusterResourcePackageItemShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterResourcePackageItemShowMetaGlobalInput) SetNo(value bool) *ActionClusterResourcePackageItemShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourcePackageItemShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourcePackageItemShowMetaGlobalInput) SelectParameters(params ...string) *ActionClusterResourcePackageItemShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterResourcePackageItemShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionClusterResourcePackageItemShowOutput is a type for action output parameters
type ActionClusterResourcePackageItemShowOutput struct {
	ClusterResource *ActionClusterResourceShowOutput `json:"cluster_resource"`
	Id int64 `json:"id"`
	Value int64 `json:"value"`
}


// Type for action response, including envelope
type ActionClusterResourcePackageItemShowResponse struct {
	Action *ActionClusterResourcePackageItemShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Item *ActionClusterResourcePackageItemShowOutput `json:"item"`
	}

	// Action output without the namespace
	Output *ActionClusterResourcePackageItemShowOutput
}


// Prepare the action for invocation
func (action *ActionClusterResourcePackageItemShow) Prepare() *ActionClusterResourcePackageItemShowInvocation {
	return &ActionClusterResourcePackageItemShowInvocation{
		Action: action,
		Path: "/v6.0/cluster_resource_packages/{cluster_resource_package_id}/items/{item_id}",
	}
}

// ActionClusterResourcePackageItemShowInvocation is used to configure action for invocation
type ActionClusterResourcePackageItemShowInvocation struct {
	// Pointer to the action
	Action *ActionClusterResourcePackageItemShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionClusterResourcePackageItemShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionClusterResourcePackageItemShowInvocation) SetPathParamInt(param string, value int64) *ActionClusterResourcePackageItemShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionClusterResourcePackageItemShowInvocation) SetPathParamString(param string, value string) *ActionClusterResourcePackageItemShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionClusterResourcePackageItemShowInvocation) NewMetaInput() *ActionClusterResourcePackageItemShowMetaGlobalInput {
	inv.MetaInput = &ActionClusterResourcePackageItemShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterResourcePackageItemShowInvocation) SetMetaInput(input *ActionClusterResourcePackageItemShowMetaGlobalInput) *ActionClusterResourcePackageItemShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterResourcePackageItemShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionClusterResourcePackageItemShowInvocation) Call() (*ActionClusterResourcePackageItemShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionClusterResourcePackageItemShowInvocation) callAsQuery() (*ActionClusterResourcePackageItemShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionClusterResourcePackageItemShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Item
	}
	return resp, err
}




func (inv *ActionClusterResourcePackageItemShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}

