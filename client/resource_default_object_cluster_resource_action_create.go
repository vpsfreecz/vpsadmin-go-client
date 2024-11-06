package client

import ()

// ActionDefaultObjectClusterResourceCreate is a type for action Default_object_cluster_resource#Create
type ActionDefaultObjectClusterResourceCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionDefaultObjectClusterResourceCreate(client *Client) *ActionDefaultObjectClusterResourceCreate {
	return &ActionDefaultObjectClusterResourceCreate{
		Client: client,
	}
}

// ActionDefaultObjectClusterResourceCreateMetaGlobalInput is a type for action global meta input parameters
type ActionDefaultObjectClusterResourceCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDefaultObjectClusterResourceCreateMetaGlobalInput) SetIncludes(value string) *ActionDefaultObjectClusterResourceCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDefaultObjectClusterResourceCreateMetaGlobalInput) SetNo(value bool) *ActionDefaultObjectClusterResourceCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDefaultObjectClusterResourceCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDefaultObjectClusterResourceCreateMetaGlobalInput) SelectParameters(params ...string) *ActionDefaultObjectClusterResourceCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDefaultObjectClusterResourceCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDefaultObjectClusterResourceCreateInput is a type for action input parameters
type ActionDefaultObjectClusterResourceCreateInput struct {
	ClassName       string `json:"class_name"`
	ClusterResource int64  `json:"cluster_resource"`
	Environment     int64  `json:"environment"`
	Value           int64  `json:"value"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetClassName sets parameter ClassName to value and selects it for sending
func (in *ActionDefaultObjectClusterResourceCreateInput) SetClassName(value string) *ActionDefaultObjectClusterResourceCreateInput {
	in.ClassName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClassName"] = nil
	return in
}

// SetClusterResource sets parameter ClusterResource to value and selects it for sending
func (in *ActionDefaultObjectClusterResourceCreateInput) SetClusterResource(value int64) *ActionDefaultObjectClusterResourceCreateInput {
	in.ClusterResource = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetClusterResourceNil(false)
	in._selectedParameters["ClusterResource"] = nil
	return in
}

// SetClusterResourceNil sets parameter ClusterResource to nil and selects it for sending
func (in *ActionDefaultObjectClusterResourceCreateInput) SetClusterResourceNil(set bool) *ActionDefaultObjectClusterResourceCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["ClusterResource"] = nil
		in.SelectParameters("ClusterResource")
	} else {
		delete(in._nilParameters, "ClusterResource")
	}
	return in
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionDefaultObjectClusterResourceCreateInput) SetEnvironment(value int64) *ActionDefaultObjectClusterResourceCreateInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEnvironmentNil(false)
	in._selectedParameters["Environment"] = nil
	return in
}

// SetEnvironmentNil sets parameter Environment to nil and selects it for sending
func (in *ActionDefaultObjectClusterResourceCreateInput) SetEnvironmentNil(set bool) *ActionDefaultObjectClusterResourceCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Environment"] = nil
		in.SelectParameters("Environment")
	} else {
		delete(in._nilParameters, "Environment")
	}
	return in
}

// SetValue sets parameter Value to value and selects it for sending
func (in *ActionDefaultObjectClusterResourceCreateInput) SetValue(value int64) *ActionDefaultObjectClusterResourceCreateInput {
	in.Value = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Value"] = nil
	return in
}

// SelectParameters sets parameters from ActionDefaultObjectClusterResourceCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDefaultObjectClusterResourceCreateInput) SelectParameters(params ...string) *ActionDefaultObjectClusterResourceCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDefaultObjectClusterResourceCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDefaultObjectClusterResourceCreateInput) UnselectParameters(params ...string) *ActionDefaultObjectClusterResourceCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDefaultObjectClusterResourceCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDefaultObjectClusterResourceCreateRequest is a type for the entire action request
type ActionDefaultObjectClusterResourceCreateRequest struct {
	DefaultObjectClusterResource map[string]interface{} `json:"default_object_cluster_resource"`
	Meta                         map[string]interface{} `json:"_meta"`
}

// ActionDefaultObjectClusterResourceCreateOutput is a type for action output parameters
type ActionDefaultObjectClusterResourceCreateOutput struct {
	ClassName       string                           `json:"class_name"`
	ClusterResource *ActionClusterResourceShowOutput `json:"cluster_resource"`
	Environment     *ActionEnvironmentShowOutput     `json:"environment"`
	Id              int64                            `json:"id"`
	Value           int64                            `json:"value"`
}

// Type for action response, including envelope
type ActionDefaultObjectClusterResourceCreateResponse struct {
	Action *ActionDefaultObjectClusterResourceCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DefaultObjectClusterResource *ActionDefaultObjectClusterResourceCreateOutput `json:"default_object_cluster_resource"`
	}

	// Action output without the namespace
	Output *ActionDefaultObjectClusterResourceCreateOutput
}

// Prepare the action for invocation
func (action *ActionDefaultObjectClusterResourceCreate) Prepare() *ActionDefaultObjectClusterResourceCreateInvocation {
	return &ActionDefaultObjectClusterResourceCreateInvocation{
		Action: action,
		Path:   "/v7.0/default_object_cluster_resources",
	}
}

// ActionDefaultObjectClusterResourceCreateInvocation is used to configure action for invocation
type ActionDefaultObjectClusterResourceCreateInvocation struct {
	// Pointer to the action
	Action *ActionDefaultObjectClusterResourceCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDefaultObjectClusterResourceCreateInput
	// Global meta input parameters
	MetaInput *ActionDefaultObjectClusterResourceCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDefaultObjectClusterResourceCreateInvocation) NewInput() *ActionDefaultObjectClusterResourceCreateInput {
	inv.Input = &ActionDefaultObjectClusterResourceCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDefaultObjectClusterResourceCreateInvocation) SetInput(input *ActionDefaultObjectClusterResourceCreateInput) *ActionDefaultObjectClusterResourceCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDefaultObjectClusterResourceCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDefaultObjectClusterResourceCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDefaultObjectClusterResourceCreateInvocation) NewMetaInput() *ActionDefaultObjectClusterResourceCreateMetaGlobalInput {
	inv.MetaInput = &ActionDefaultObjectClusterResourceCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDefaultObjectClusterResourceCreateInvocation) SetMetaInput(input *ActionDefaultObjectClusterResourceCreateMetaGlobalInput) *ActionDefaultObjectClusterResourceCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDefaultObjectClusterResourceCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDefaultObjectClusterResourceCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDefaultObjectClusterResourceCreateInvocation) Call() (*ActionDefaultObjectClusterResourceCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDefaultObjectClusterResourceCreateInvocation) callAsBody() (*ActionDefaultObjectClusterResourceCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDefaultObjectClusterResourceCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DefaultObjectClusterResource
	}
	return resp, err
}

func (inv *ActionDefaultObjectClusterResourceCreateInvocation) makeAllInputParams() *ActionDefaultObjectClusterResourceCreateRequest {
	return &ActionDefaultObjectClusterResourceCreateRequest{
		DefaultObjectClusterResource: inv.makeInputParams(),
		Meta:                         inv.makeMetaInputParams(),
	}
}

func (inv *ActionDefaultObjectClusterResourceCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("ClassName") {
			ret["class_name"] = inv.Input.ClassName
		}
		if inv.IsParameterSelected("ClusterResource") {
			if inv.IsParameterNil("ClusterResource") {
				ret["cluster_resource"] = nil
			} else {
				ret["cluster_resource"] = inv.Input.ClusterResource
			}
		}
		if inv.IsParameterSelected("Environment") {
			if inv.IsParameterNil("Environment") {
				ret["environment"] = nil
			} else {
				ret["environment"] = inv.Input.Environment
			}
		}
		if inv.IsParameterSelected("Value") {
			ret["value"] = inv.Input.Value
		}
	}

	return ret
}

func (inv *ActionDefaultObjectClusterResourceCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
