package client

import (
	"strings"
)

// ActionDefaultObjectClusterResourceDelete is a type for action Default_object_cluster_resource#Delete
type ActionDefaultObjectClusterResourceDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionDefaultObjectClusterResourceDelete(client *Client) *ActionDefaultObjectClusterResourceDelete {
	return &ActionDefaultObjectClusterResourceDelete{
		Client: client,
	}
}

// ActionDefaultObjectClusterResourceDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionDefaultObjectClusterResourceDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDefaultObjectClusterResourceDeleteMetaGlobalInput) SetIncludes(value string) *ActionDefaultObjectClusterResourceDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDefaultObjectClusterResourceDeleteMetaGlobalInput) SetNo(value bool) *ActionDefaultObjectClusterResourceDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDefaultObjectClusterResourceDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDefaultObjectClusterResourceDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionDefaultObjectClusterResourceDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDefaultObjectClusterResourceDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDefaultObjectClusterResourceDeleteRequest is a type for the entire action request
type ActionDefaultObjectClusterResourceDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionDefaultObjectClusterResourceDeleteResponse struct {
	Action *ActionDefaultObjectClusterResourceDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionDefaultObjectClusterResourceDelete) Prepare() *ActionDefaultObjectClusterResourceDeleteInvocation {
	return &ActionDefaultObjectClusterResourceDeleteInvocation{
		Action: action,
		Path:   "/v7.0/default_object_cluster_resources/{default_object_cluster_resource_id}",
	}
}

// ActionDefaultObjectClusterResourceDeleteInvocation is used to configure action for invocation
type ActionDefaultObjectClusterResourceDeleteInvocation struct {
	// Pointer to the action
	Action *ActionDefaultObjectClusterResourceDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDefaultObjectClusterResourceDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDefaultObjectClusterResourceDeleteInvocation) SetPathParamInt(param string, value int64) *ActionDefaultObjectClusterResourceDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDefaultObjectClusterResourceDeleteInvocation) SetPathParamString(param string, value string) *ActionDefaultObjectClusterResourceDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDefaultObjectClusterResourceDeleteInvocation) NewMetaInput() *ActionDefaultObjectClusterResourceDeleteMetaGlobalInput {
	inv.MetaInput = &ActionDefaultObjectClusterResourceDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDefaultObjectClusterResourceDeleteInvocation) SetMetaInput(input *ActionDefaultObjectClusterResourceDeleteMetaGlobalInput) *ActionDefaultObjectClusterResourceDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDefaultObjectClusterResourceDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDefaultObjectClusterResourceDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDefaultObjectClusterResourceDeleteInvocation) Call() (*ActionDefaultObjectClusterResourceDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDefaultObjectClusterResourceDeleteInvocation) callAsBody() (*ActionDefaultObjectClusterResourceDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDefaultObjectClusterResourceDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionDefaultObjectClusterResourceDeleteInvocation) makeAllInputParams() *ActionDefaultObjectClusterResourceDeleteRequest {
	return &ActionDefaultObjectClusterResourceDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionDefaultObjectClusterResourceDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
