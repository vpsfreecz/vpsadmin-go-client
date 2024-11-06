package client

import (
	"strings"
)

// ActionDefaultObjectClusterResourceUpdate is a type for action Default_object_cluster_resource#Update
type ActionDefaultObjectClusterResourceUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionDefaultObjectClusterResourceUpdate(client *Client) *ActionDefaultObjectClusterResourceUpdate {
	return &ActionDefaultObjectClusterResourceUpdate{
		Client: client,
	}
}

// ActionDefaultObjectClusterResourceUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionDefaultObjectClusterResourceUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDefaultObjectClusterResourceUpdateMetaGlobalInput) SetIncludes(value string) *ActionDefaultObjectClusterResourceUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDefaultObjectClusterResourceUpdateMetaGlobalInput) SetNo(value bool) *ActionDefaultObjectClusterResourceUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDefaultObjectClusterResourceUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDefaultObjectClusterResourceUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionDefaultObjectClusterResourceUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDefaultObjectClusterResourceUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDefaultObjectClusterResourceUpdateInput is a type for action input parameters
type ActionDefaultObjectClusterResourceUpdateInput struct {
	Value int64 `json:"value"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetValue sets parameter Value to value and selects it for sending
func (in *ActionDefaultObjectClusterResourceUpdateInput) SetValue(value int64) *ActionDefaultObjectClusterResourceUpdateInput {
	in.Value = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Value"] = nil
	return in
}

// SelectParameters sets parameters from ActionDefaultObjectClusterResourceUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDefaultObjectClusterResourceUpdateInput) SelectParameters(params ...string) *ActionDefaultObjectClusterResourceUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDefaultObjectClusterResourceUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDefaultObjectClusterResourceUpdateInput) UnselectParameters(params ...string) *ActionDefaultObjectClusterResourceUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDefaultObjectClusterResourceUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDefaultObjectClusterResourceUpdateRequest is a type for the entire action request
type ActionDefaultObjectClusterResourceUpdateRequest struct {
	DefaultObjectClusterResource map[string]interface{} `json:"default_object_cluster_resource"`
	Meta                         map[string]interface{} `json:"_meta"`
}

// ActionDefaultObjectClusterResourceUpdateOutput is a type for action output parameters
type ActionDefaultObjectClusterResourceUpdateOutput struct {
	ClassName       string                           `json:"class_name"`
	ClusterResource *ActionClusterResourceShowOutput `json:"cluster_resource"`
	Environment     *ActionEnvironmentShowOutput     `json:"environment"`
	Id              int64                            `json:"id"`
	Value           int64                            `json:"value"`
}

// Type for action response, including envelope
type ActionDefaultObjectClusterResourceUpdateResponse struct {
	Action *ActionDefaultObjectClusterResourceUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DefaultObjectClusterResource *ActionDefaultObjectClusterResourceUpdateOutput `json:"default_object_cluster_resource"`
	}

	// Action output without the namespace
	Output *ActionDefaultObjectClusterResourceUpdateOutput
}

// Prepare the action for invocation
func (action *ActionDefaultObjectClusterResourceUpdate) Prepare() *ActionDefaultObjectClusterResourceUpdateInvocation {
	return &ActionDefaultObjectClusterResourceUpdateInvocation{
		Action: action,
		Path:   "/v7.0/default_object_cluster_resources/{default_object_cluster_resource_id}",
	}
}

// ActionDefaultObjectClusterResourceUpdateInvocation is used to configure action for invocation
type ActionDefaultObjectClusterResourceUpdateInvocation struct {
	// Pointer to the action
	Action *ActionDefaultObjectClusterResourceUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDefaultObjectClusterResourceUpdateInput
	// Global meta input parameters
	MetaInput *ActionDefaultObjectClusterResourceUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDefaultObjectClusterResourceUpdateInvocation) SetPathParamInt(param string, value int64) *ActionDefaultObjectClusterResourceUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDefaultObjectClusterResourceUpdateInvocation) SetPathParamString(param string, value string) *ActionDefaultObjectClusterResourceUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDefaultObjectClusterResourceUpdateInvocation) NewInput() *ActionDefaultObjectClusterResourceUpdateInput {
	inv.Input = &ActionDefaultObjectClusterResourceUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDefaultObjectClusterResourceUpdateInvocation) SetInput(input *ActionDefaultObjectClusterResourceUpdateInput) *ActionDefaultObjectClusterResourceUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDefaultObjectClusterResourceUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDefaultObjectClusterResourceUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDefaultObjectClusterResourceUpdateInvocation) NewMetaInput() *ActionDefaultObjectClusterResourceUpdateMetaGlobalInput {
	inv.MetaInput = &ActionDefaultObjectClusterResourceUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDefaultObjectClusterResourceUpdateInvocation) SetMetaInput(input *ActionDefaultObjectClusterResourceUpdateMetaGlobalInput) *ActionDefaultObjectClusterResourceUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDefaultObjectClusterResourceUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDefaultObjectClusterResourceUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDefaultObjectClusterResourceUpdateInvocation) Call() (*ActionDefaultObjectClusterResourceUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDefaultObjectClusterResourceUpdateInvocation) callAsBody() (*ActionDefaultObjectClusterResourceUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDefaultObjectClusterResourceUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DefaultObjectClusterResource
	}
	return resp, err
}

func (inv *ActionDefaultObjectClusterResourceUpdateInvocation) makeAllInputParams() *ActionDefaultObjectClusterResourceUpdateRequest {
	return &ActionDefaultObjectClusterResourceUpdateRequest{
		DefaultObjectClusterResource: inv.makeInputParams(),
		Meta:                         inv.makeMetaInputParams(),
	}
}

func (inv *ActionDefaultObjectClusterResourceUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Value") {
			ret["value"] = inv.Input.Value
		}
	}

	return ret
}

func (inv *ActionDefaultObjectClusterResourceUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
