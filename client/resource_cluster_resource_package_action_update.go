package client

import (
	"strings"
)

// ActionClusterResourcePackageUpdate is a type for action Cluster_resource_package#Update
type ActionClusterResourcePackageUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterResourcePackageUpdate(client *Client) *ActionClusterResourcePackageUpdate {
	return &ActionClusterResourcePackageUpdate{
		Client: client,
	}
}

// ActionClusterResourcePackageUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionClusterResourcePackageUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterResourcePackageUpdateMetaGlobalInput) SetNo(value bool) *ActionClusterResourcePackageUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionClusterResourcePackageUpdateMetaGlobalInput) SetIncludes(value string) *ActionClusterResourcePackageUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourcePackageUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourcePackageUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionClusterResourcePackageUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterResourcePackageUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterResourcePackageUpdateInput is a type for action input parameters
type ActionClusterResourcePackageUpdateInput struct {
	Label string `json:"label"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionClusterResourcePackageUpdateInput) SetLabel(value string) *ActionClusterResourcePackageUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourcePackageUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourcePackageUpdateInput) SelectParameters(params ...string) *ActionClusterResourcePackageUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterResourcePackageUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterResourcePackageUpdateRequest is a type for the entire action request
type ActionClusterResourcePackageUpdateRequest struct {
	ClusterResourcePackage map[string]interface{} `json:"cluster_resource_package"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionClusterResourcePackageUpdateOutput is a type for action output parameters
type ActionClusterResourcePackageUpdateOutput struct {
	Id int64 `json:"id"`
	Label string `json:"label"`
	Environment *ActionEnvironmentShowOutput `json:"environment"`
	User *ActionUserShowOutput `json:"user"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


// Type for action response, including envelope
type ActionClusterResourcePackageUpdateResponse struct {
	Action *ActionClusterResourcePackageUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ClusterResourcePackage *ActionClusterResourcePackageUpdateOutput `json:"cluster_resource_package"`
	}

	// Action output without the namespace
	Output *ActionClusterResourcePackageUpdateOutput
}


// Prepare the action for invocation
func (action *ActionClusterResourcePackageUpdate) Prepare() *ActionClusterResourcePackageUpdateInvocation {
	return &ActionClusterResourcePackageUpdateInvocation{
		Action: action,
		Path: "/v5.0/cluster_resource_packages/:cluster_resource_package_id",
	}
}

// ActionClusterResourcePackageUpdateInvocation is used to configure action for invocation
type ActionClusterResourcePackageUpdateInvocation struct {
	// Pointer to the action
	Action *ActionClusterResourcePackageUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionClusterResourcePackageUpdateInput
	// Global meta input parameters
	MetaInput *ActionClusterResourcePackageUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionClusterResourcePackageUpdateInvocation) SetPathParamInt(param string, value int64) *ActionClusterResourcePackageUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionClusterResourcePackageUpdateInvocation) SetPathParamString(param string, value string) *ActionClusterResourcePackageUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionClusterResourcePackageUpdateInvocation) NewInput() *ActionClusterResourcePackageUpdateInput {
	inv.Input = &ActionClusterResourcePackageUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionClusterResourcePackageUpdateInvocation) SetInput(input *ActionClusterResourcePackageUpdateInput) *ActionClusterResourcePackageUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionClusterResourcePackageUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionClusterResourcePackageUpdateInvocation) NewMetaInput() *ActionClusterResourcePackageUpdateMetaGlobalInput {
	inv.MetaInput = &ActionClusterResourcePackageUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterResourcePackageUpdateInvocation) SetMetaInput(input *ActionClusterResourcePackageUpdateMetaGlobalInput) *ActionClusterResourcePackageUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterResourcePackageUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionClusterResourcePackageUpdateInvocation) Call() (*ActionClusterResourcePackageUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionClusterResourcePackageUpdateInvocation) callAsBody() (*ActionClusterResourcePackageUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionClusterResourcePackageUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.ClusterResourcePackage
	}
	return resp, err
}




func (inv *ActionClusterResourcePackageUpdateInvocation) makeAllInputParams() *ActionClusterResourcePackageUpdateRequest {
	return &ActionClusterResourcePackageUpdateRequest{
		ClusterResourcePackage: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionClusterResourcePackageUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
	}

	return ret
}

func (inv *ActionClusterResourcePackageUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
