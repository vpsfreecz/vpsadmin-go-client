package client

import (
	"strings"
)

// ActionClusterResourcePackageItemCreate is a type for action Cluster_resource_package.Item#Create
type ActionClusterResourcePackageItemCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterResourcePackageItemCreate(client *Client) *ActionClusterResourcePackageItemCreate {
	return &ActionClusterResourcePackageItemCreate{
		Client: client,
	}
}

// ActionClusterResourcePackageItemCreateMetaGlobalInput is a type for action global meta input parameters
type ActionClusterResourcePackageItemCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionClusterResourcePackageItemCreateMetaGlobalInput) SetIncludes(value string) *ActionClusterResourcePackageItemCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterResourcePackageItemCreateMetaGlobalInput) SetNo(value bool) *ActionClusterResourcePackageItemCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourcePackageItemCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourcePackageItemCreateMetaGlobalInput) SelectParameters(params ...string) *ActionClusterResourcePackageItemCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterResourcePackageItemCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterResourcePackageItemCreateInput is a type for action input parameters
type ActionClusterResourcePackageItemCreateInput struct {
	ClusterResource int64 `json:"cluster_resource"`
	Value int64 `json:"value"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetClusterResource sets parameter ClusterResource to value and selects it for sending
func (in *ActionClusterResourcePackageItemCreateInput) SetClusterResource(value int64) *ActionClusterResourcePackageItemCreateInput {
	in.ClusterResource = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClusterResource"] = nil
	return in
}
// SetValue sets parameter Value to value and selects it for sending
func (in *ActionClusterResourcePackageItemCreateInput) SetValue(value int64) *ActionClusterResourcePackageItemCreateInput {
	in.Value = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Value"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourcePackageItemCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourcePackageItemCreateInput) SelectParameters(params ...string) *ActionClusterResourcePackageItemCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterResourcePackageItemCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterResourcePackageItemCreateRequest is a type for the entire action request
type ActionClusterResourcePackageItemCreateRequest struct {
	Item map[string]interface{} `json:"item"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionClusterResourcePackageItemCreateOutput is a type for action output parameters
type ActionClusterResourcePackageItemCreateOutput struct {
	ClusterResource *ActionClusterResourceShowOutput `json:"cluster_resource"`
	Id int64 `json:"id"`
	Value int64 `json:"value"`
}


// Type for action response, including envelope
type ActionClusterResourcePackageItemCreateResponse struct {
	Action *ActionClusterResourcePackageItemCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Item *ActionClusterResourcePackageItemCreateOutput `json:"item"`
	}

	// Action output without the namespace
	Output *ActionClusterResourcePackageItemCreateOutput
}


// Prepare the action for invocation
func (action *ActionClusterResourcePackageItemCreate) Prepare() *ActionClusterResourcePackageItemCreateInvocation {
	return &ActionClusterResourcePackageItemCreateInvocation{
		Action: action,
		Path: "/v6.0/cluster_resource_packages/{cluster_resource_package_id}/items",
	}
}

// ActionClusterResourcePackageItemCreateInvocation is used to configure action for invocation
type ActionClusterResourcePackageItemCreateInvocation struct {
	// Pointer to the action
	Action *ActionClusterResourcePackageItemCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionClusterResourcePackageItemCreateInput
	// Global meta input parameters
	MetaInput *ActionClusterResourcePackageItemCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionClusterResourcePackageItemCreateInvocation) SetPathParamInt(param string, value int64) *ActionClusterResourcePackageItemCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionClusterResourcePackageItemCreateInvocation) SetPathParamString(param string, value string) *ActionClusterResourcePackageItemCreateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionClusterResourcePackageItemCreateInvocation) NewInput() *ActionClusterResourcePackageItemCreateInput {
	inv.Input = &ActionClusterResourcePackageItemCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionClusterResourcePackageItemCreateInvocation) SetInput(input *ActionClusterResourcePackageItemCreateInput) *ActionClusterResourcePackageItemCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionClusterResourcePackageItemCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionClusterResourcePackageItemCreateInvocation) NewMetaInput() *ActionClusterResourcePackageItemCreateMetaGlobalInput {
	inv.MetaInput = &ActionClusterResourcePackageItemCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterResourcePackageItemCreateInvocation) SetMetaInput(input *ActionClusterResourcePackageItemCreateMetaGlobalInput) *ActionClusterResourcePackageItemCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterResourcePackageItemCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionClusterResourcePackageItemCreateInvocation) Call() (*ActionClusterResourcePackageItemCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionClusterResourcePackageItemCreateInvocation) callAsBody() (*ActionClusterResourcePackageItemCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionClusterResourcePackageItemCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Item
	}
	return resp, err
}




func (inv *ActionClusterResourcePackageItemCreateInvocation) makeAllInputParams() *ActionClusterResourcePackageItemCreateRequest {
	return &ActionClusterResourcePackageItemCreateRequest{
		Item: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionClusterResourcePackageItemCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("ClusterResource") {
			ret["cluster_resource"] = inv.Input.ClusterResource
		}
		if inv.IsParameterSelected("Value") {
			ret["value"] = inv.Input.Value
		}
	}

	return ret
}

func (inv *ActionClusterResourcePackageItemCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
