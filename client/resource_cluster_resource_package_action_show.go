package client

import (
	"strings"
)

// ActionClusterResourcePackageShow is a type for action Cluster_resource_package#Show
type ActionClusterResourcePackageShow struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterResourcePackageShow(client *Client) *ActionClusterResourcePackageShow {
	return &ActionClusterResourcePackageShow{
		Client: client,
	}
}

// ActionClusterResourcePackageShowMetaGlobalInput is a type for action global meta input parameters
type ActionClusterResourcePackageShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterResourcePackageShowMetaGlobalInput) SetNo(value bool) *ActionClusterResourcePackageShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionClusterResourcePackageShowMetaGlobalInput) SetIncludes(value string) *ActionClusterResourcePackageShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourcePackageShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourcePackageShowMetaGlobalInput) SelectParameters(params ...string) *ActionClusterResourcePackageShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterResourcePackageShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionClusterResourcePackageShowOutput is a type for action output parameters
type ActionClusterResourcePackageShowOutput struct {
	Id int64 `json:"id"`
	Label string `json:"label"`
	Environment *ActionEnvironmentShowOutput `json:"environment"`
	User *ActionUserShowOutput `json:"user"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


// Type for action response, including envelope
type ActionClusterResourcePackageShowResponse struct {
	Action *ActionClusterResourcePackageShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ClusterResourcePackage *ActionClusterResourcePackageShowOutput `json:"cluster_resource_package"`
	}

	// Action output without the namespace
	Output *ActionClusterResourcePackageShowOutput
}


// Prepare the action for invocation
func (action *ActionClusterResourcePackageShow) Prepare() *ActionClusterResourcePackageShowInvocation {
	return &ActionClusterResourcePackageShowInvocation{
		Action: action,
		Path: "/v5.0/cluster_resource_packages/{cluster_resource_package_id}",
	}
}

// ActionClusterResourcePackageShowInvocation is used to configure action for invocation
type ActionClusterResourcePackageShowInvocation struct {
	// Pointer to the action
	Action *ActionClusterResourcePackageShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionClusterResourcePackageShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionClusterResourcePackageShowInvocation) SetPathParamInt(param string, value int64) *ActionClusterResourcePackageShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionClusterResourcePackageShowInvocation) SetPathParamString(param string, value string) *ActionClusterResourcePackageShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionClusterResourcePackageShowInvocation) NewMetaInput() *ActionClusterResourcePackageShowMetaGlobalInput {
	inv.MetaInput = &ActionClusterResourcePackageShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterResourcePackageShowInvocation) SetMetaInput(input *ActionClusterResourcePackageShowMetaGlobalInput) *ActionClusterResourcePackageShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterResourcePackageShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionClusterResourcePackageShowInvocation) Call() (*ActionClusterResourcePackageShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionClusterResourcePackageShowInvocation) callAsQuery() (*ActionClusterResourcePackageShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionClusterResourcePackageShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.ClusterResourcePackage
	}
	return resp, err
}




func (inv *ActionClusterResourcePackageShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

