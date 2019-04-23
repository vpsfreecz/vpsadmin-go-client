package client

import (
	"strings"
)

// ActionUserClusterResourcePackageShow is a type for action User_cluster_resource_package#Show
type ActionUserClusterResourcePackageShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserClusterResourcePackageShow(client *Client) *ActionUserClusterResourcePackageShow {
	return &ActionUserClusterResourcePackageShow{
		Client: client,
	}
}

// ActionUserClusterResourcePackageShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserClusterResourcePackageShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserClusterResourcePackageShowMetaGlobalInput) SetNo(value bool) *ActionUserClusterResourcePackageShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserClusterResourcePackageShowMetaGlobalInput) SetIncludes(value string) *ActionUserClusterResourcePackageShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserClusterResourcePackageShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserClusterResourcePackageShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserClusterResourcePackageShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserClusterResourcePackageShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionUserClusterResourcePackageShowOutput is a type for action output parameters
type ActionUserClusterResourcePackageShowOutput struct {
	Id int64 `json:"id"`
	Environment *ActionEnvironmentShowOutput `json:"environment"`
	User *ActionUserShowOutput `json:"user"`
	ClusterResourcePackage *ActionClusterResourcePackageShowOutput `json:"cluster_resource_package"`
	AddedBy *ActionUserShowOutput `json:"added_by"`
	Label string `json:"label"`
	IsPersonal bool `json:"is_personal"`
	Comment string `json:"comment"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


// Type for action response, including envelope
type ActionUserClusterResourcePackageShowResponse struct {
	Action *ActionUserClusterResourcePackageShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserClusterResourcePackage *ActionUserClusterResourcePackageShowOutput `json:"user_cluster_resource_package"`
	}

	// Action output without the namespace
	Output *ActionUserClusterResourcePackageShowOutput
}


// Prepare the action for invocation
func (action *ActionUserClusterResourcePackageShow) Prepare() *ActionUserClusterResourcePackageShowInvocation {
	return &ActionUserClusterResourcePackageShowInvocation{
		Action: action,
		Path: "/v5.0/user_cluster_resource_packages/:user_cluster_resource_package_id",
	}
}

// ActionUserClusterResourcePackageShowInvocation is used to configure action for invocation
type ActionUserClusterResourcePackageShowInvocation struct {
	// Pointer to the action
	Action *ActionUserClusterResourcePackageShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserClusterResourcePackageShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserClusterResourcePackageShowInvocation) SetPathParamInt(param string, value int64) *ActionUserClusterResourcePackageShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserClusterResourcePackageShowInvocation) SetPathParamString(param string, value string) *ActionUserClusterResourcePackageShowInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserClusterResourcePackageShowInvocation) SetMetaInput(input *ActionUserClusterResourcePackageShowMetaGlobalInput) *ActionUserClusterResourcePackageShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserClusterResourcePackageShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserClusterResourcePackageShowInvocation) Call() (*ActionUserClusterResourcePackageShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserClusterResourcePackageShowInvocation) callAsQuery() (*ActionUserClusterResourcePackageShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserClusterResourcePackageShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserClusterResourcePackage
	}
	return resp, err
}




func (inv *ActionUserClusterResourcePackageShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

