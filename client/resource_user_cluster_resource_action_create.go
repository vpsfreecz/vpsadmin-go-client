package client

import (
	"strings"
)

// ActionUserClusterResourceCreate is a type for action User.Cluster_resource#Create
type ActionUserClusterResourceCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserClusterResourceCreate(client *Client) *ActionUserClusterResourceCreate {
	return &ActionUserClusterResourceCreate{
		Client: client,
	}
}

// ActionUserClusterResourceCreateMetaGlobalInput is a type for action global meta input parameters
type ActionUserClusterResourceCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserClusterResourceCreateMetaGlobalInput) SetNo(value bool) *ActionUserClusterResourceCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserClusterResourceCreateMetaGlobalInput) SetIncludes(value string) *ActionUserClusterResourceCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserClusterResourceCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserClusterResourceCreateMetaGlobalInput) SelectParameters(params ...string) *ActionUserClusterResourceCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserClusterResourceCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserClusterResourceCreateInput is a type for action input parameters
type ActionUserClusterResourceCreateInput struct {
	Environment int64 `json:"environment"`
	ClusterResource int64 `json:"cluster_resource"`
	Value int64 `json:"value"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionUserClusterResourceCreateInput) SetEnvironment(value int64) *ActionUserClusterResourceCreateInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Environment"] = nil
	return in
}
// SetClusterResource sets parameter ClusterResource to value and selects it for sending
func (in *ActionUserClusterResourceCreateInput) SetClusterResource(value int64) *ActionUserClusterResourceCreateInput {
	in.ClusterResource = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClusterResource"] = nil
	return in
}
// SetValue sets parameter Value to value and selects it for sending
func (in *ActionUserClusterResourceCreateInput) SetValue(value int64) *ActionUserClusterResourceCreateInput {
	in.Value = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Value"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserClusterResourceCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserClusterResourceCreateInput) SelectParameters(params ...string) *ActionUserClusterResourceCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserClusterResourceCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserClusterResourceCreateRequest is a type for the entire action request
type ActionUserClusterResourceCreateRequest struct {
	ClusterResource map[string]interface{} `json:"cluster_resource"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionUserClusterResourceCreateOutput is a type for action output parameters
type ActionUserClusterResourceCreateOutput struct {
	Id int64 `json:"id"`
	Environment *ActionEnvironmentShowOutput `json:"environment"`
	ClusterResource *ActionClusterResourceShowOutput `json:"cluster_resource"`
	Value int64 `json:"value"`
	Used int64 `json:"used"`
	Free int64 `json:"free"`
}


// Type for action response, including envelope
type ActionUserClusterResourceCreateResponse struct {
	Action *ActionUserClusterResourceCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ClusterResource *ActionUserClusterResourceCreateOutput `json:"cluster_resource"`
	}

	// Action output without the namespace
	Output *ActionUserClusterResourceCreateOutput
}


// Prepare the action for invocation
func (action *ActionUserClusterResourceCreate) Prepare() *ActionUserClusterResourceCreateInvocation {
	return &ActionUserClusterResourceCreateInvocation{
		Action: action,
		Path: "/v5.0/users/:user_id/cluster_resources",
	}
}

// ActionUserClusterResourceCreateInvocation is used to configure action for invocation
type ActionUserClusterResourceCreateInvocation struct {
	// Pointer to the action
	Action *ActionUserClusterResourceCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserClusterResourceCreateInput
	// Global meta input parameters
	MetaInput *ActionUserClusterResourceCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserClusterResourceCreateInvocation) SetPathParamInt(param string, value int64) *ActionUserClusterResourceCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserClusterResourceCreateInvocation) SetPathParamString(param string, value string) *ActionUserClusterResourceCreateInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserClusterResourceCreateInvocation) SetInput(input *ActionUserClusterResourceCreateInput) *ActionUserClusterResourceCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserClusterResourceCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserClusterResourceCreateInvocation) SetMetaInput(input *ActionUserClusterResourceCreateMetaGlobalInput) *ActionUserClusterResourceCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserClusterResourceCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserClusterResourceCreateInvocation) Call() (*ActionUserClusterResourceCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionUserClusterResourceCreateInvocation) callAsBody() (*ActionUserClusterResourceCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserClusterResourceCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.ClusterResource
	}
	return resp, err
}




func (inv *ActionUserClusterResourceCreateInvocation) makeAllInputParams() *ActionUserClusterResourceCreateRequest {
	return &ActionUserClusterResourceCreateRequest{
		ClusterResource: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserClusterResourceCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			ret["environment"] = inv.Input.Environment
		}
		if inv.IsParameterSelected("ClusterResource") {
			ret["cluster_resource"] = inv.Input.ClusterResource
		}
		if inv.IsParameterSelected("Value") {
			ret["value"] = inv.Input.Value
		}
	}

	return ret
}

func (inv *ActionUserClusterResourceCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
