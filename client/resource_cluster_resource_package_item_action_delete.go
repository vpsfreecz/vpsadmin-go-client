package client

import (
	"strings"
)

// ActionClusterResourcePackageItemDelete is a type for action Cluster_resource_package.Item#Delete
type ActionClusterResourcePackageItemDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterResourcePackageItemDelete(client *Client) *ActionClusterResourcePackageItemDelete {
	return &ActionClusterResourcePackageItemDelete{
		Client: client,
	}
}

// ActionClusterResourcePackageItemDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionClusterResourcePackageItemDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionClusterResourcePackageItemDeleteMetaGlobalInput) SetIncludes(value string) *ActionClusterResourcePackageItemDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterResourcePackageItemDeleteMetaGlobalInput) SetNo(value bool) *ActionClusterResourcePackageItemDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourcePackageItemDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourcePackageItemDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionClusterResourcePackageItemDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterResourcePackageItemDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionClusterResourcePackageItemDeleteRequest is a type for the entire action request
type ActionClusterResourcePackageItemDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionClusterResourcePackageItemDeleteResponse struct {
	Action *ActionClusterResourcePackageItemDelete `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionClusterResourcePackageItemDelete) Prepare() *ActionClusterResourcePackageItemDeleteInvocation {
	return &ActionClusterResourcePackageItemDeleteInvocation{
		Action: action,
		Path: "/v6.0/cluster_resource_packages/{cluster_resource_package_id}/items/{item_id}",
	}
}

// ActionClusterResourcePackageItemDeleteInvocation is used to configure action for invocation
type ActionClusterResourcePackageItemDeleteInvocation struct {
	// Pointer to the action
	Action *ActionClusterResourcePackageItemDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionClusterResourcePackageItemDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionClusterResourcePackageItemDeleteInvocation) SetPathParamInt(param string, value int64) *ActionClusterResourcePackageItemDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionClusterResourcePackageItemDeleteInvocation) SetPathParamString(param string, value string) *ActionClusterResourcePackageItemDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionClusterResourcePackageItemDeleteInvocation) NewMetaInput() *ActionClusterResourcePackageItemDeleteMetaGlobalInput {
	inv.MetaInput = &ActionClusterResourcePackageItemDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterResourcePackageItemDeleteInvocation) SetMetaInput(input *ActionClusterResourcePackageItemDeleteMetaGlobalInput) *ActionClusterResourcePackageItemDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterResourcePackageItemDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionClusterResourcePackageItemDeleteInvocation) Call() (*ActionClusterResourcePackageItemDeleteResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionClusterResourcePackageItemDeleteInvocation) callAsBody() (*ActionClusterResourcePackageItemDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionClusterResourcePackageItemDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionClusterResourcePackageItemDeleteInvocation) makeAllInputParams() *ActionClusterResourcePackageItemDeleteRequest {
	return &ActionClusterResourcePackageItemDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionClusterResourcePackageItemDeleteInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
