package client

import (
	"strings"
)

// ActionUserClusterResourcePackageDelete is a type for action User_cluster_resource_package#Delete
type ActionUserClusterResourcePackageDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionUserClusterResourcePackageDelete(client *Client) *ActionUserClusterResourcePackageDelete {
	return &ActionUserClusterResourcePackageDelete{
		Client: client,
	}
}

// ActionUserClusterResourcePackageDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionUserClusterResourcePackageDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserClusterResourcePackageDeleteMetaGlobalInput) SetIncludes(value string) *ActionUserClusterResourcePackageDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserClusterResourcePackageDeleteMetaGlobalInput) SetNo(value bool) *ActionUserClusterResourcePackageDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserClusterResourcePackageDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserClusterResourcePackageDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionUserClusterResourcePackageDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserClusterResourcePackageDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserClusterResourcePackageDeleteRequest is a type for the entire action request
type ActionUserClusterResourcePackageDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionUserClusterResourcePackageDeleteResponse struct {
	Action *ActionUserClusterResourcePackageDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionUserClusterResourcePackageDelete) Prepare() *ActionUserClusterResourcePackageDeleteInvocation {
	return &ActionUserClusterResourcePackageDeleteInvocation{
		Action: action,
		Path:   "/v6.0/user_cluster_resource_packages/{user_cluster_resource_package_id}",
	}
}

// ActionUserClusterResourcePackageDeleteInvocation is used to configure action for invocation
type ActionUserClusterResourcePackageDeleteInvocation struct {
	// Pointer to the action
	Action *ActionUserClusterResourcePackageDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserClusterResourcePackageDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserClusterResourcePackageDeleteInvocation) SetPathParamInt(param string, value int64) *ActionUserClusterResourcePackageDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserClusterResourcePackageDeleteInvocation) SetPathParamString(param string, value string) *ActionUserClusterResourcePackageDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserClusterResourcePackageDeleteInvocation) NewMetaInput() *ActionUserClusterResourcePackageDeleteMetaGlobalInput {
	inv.MetaInput = &ActionUserClusterResourcePackageDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserClusterResourcePackageDeleteInvocation) SetMetaInput(input *ActionUserClusterResourcePackageDeleteMetaGlobalInput) *ActionUserClusterResourcePackageDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserClusterResourcePackageDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserClusterResourcePackageDeleteInvocation) Call() (*ActionUserClusterResourcePackageDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionUserClusterResourcePackageDeleteInvocation) callAsBody() (*ActionUserClusterResourcePackageDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserClusterResourcePackageDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionUserClusterResourcePackageDeleteInvocation) makeAllInputParams() *ActionUserClusterResourcePackageDeleteRequest {
	return &ActionUserClusterResourcePackageDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserClusterResourcePackageDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
