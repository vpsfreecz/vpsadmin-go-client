package client

import (
	"strings"
)

// ActionClusterResourcePackageDelete is a type for action Cluster_resource_package#Delete
type ActionClusterResourcePackageDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterResourcePackageDelete(client *Client) *ActionClusterResourcePackageDelete {
	return &ActionClusterResourcePackageDelete{
		Client: client,
	}
}

// ActionClusterResourcePackageDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionClusterResourcePackageDeleteMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterResourcePackageDeleteMetaGlobalInput) SetNo(value bool) *ActionClusterResourcePackageDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionClusterResourcePackageDeleteMetaGlobalInput) SetIncludes(value string) *ActionClusterResourcePackageDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourcePackageDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourcePackageDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionClusterResourcePackageDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterResourcePackageDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionClusterResourcePackageDeleteRequest is a type for the entire action request
type ActionClusterResourcePackageDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionClusterResourcePackageDeleteResponse struct {
	Action *ActionClusterResourcePackageDelete `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionClusterResourcePackageDelete) Prepare() *ActionClusterResourcePackageDeleteInvocation {
	return &ActionClusterResourcePackageDeleteInvocation{
		Action: action,
		Path: "/v5.0/cluster_resource_packages/{cluster_resource_package_id}",
	}
}

// ActionClusterResourcePackageDeleteInvocation is used to configure action for invocation
type ActionClusterResourcePackageDeleteInvocation struct {
	// Pointer to the action
	Action *ActionClusterResourcePackageDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionClusterResourcePackageDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionClusterResourcePackageDeleteInvocation) SetPathParamInt(param string, value int64) *ActionClusterResourcePackageDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionClusterResourcePackageDeleteInvocation) SetPathParamString(param string, value string) *ActionClusterResourcePackageDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionClusterResourcePackageDeleteInvocation) NewMetaInput() *ActionClusterResourcePackageDeleteMetaGlobalInput {
	inv.MetaInput = &ActionClusterResourcePackageDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterResourcePackageDeleteInvocation) SetMetaInput(input *ActionClusterResourcePackageDeleteMetaGlobalInput) *ActionClusterResourcePackageDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterResourcePackageDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionClusterResourcePackageDeleteInvocation) Call() (*ActionClusterResourcePackageDeleteResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionClusterResourcePackageDeleteInvocation) callAsBody() (*ActionClusterResourcePackageDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionClusterResourcePackageDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionClusterResourcePackageDeleteInvocation) makeAllInputParams() *ActionClusterResourcePackageDeleteRequest {
	return &ActionClusterResourcePackageDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionClusterResourcePackageDeleteInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
