package client

import ()

// ActionClusterResourcePackageCreate is a type for action Cluster_resource_package#Create
type ActionClusterResourcePackageCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterResourcePackageCreate(client *Client) *ActionClusterResourcePackageCreate {
	return &ActionClusterResourcePackageCreate{
		Client: client,
	}
}

// ActionClusterResourcePackageCreateMetaGlobalInput is a type for action global meta input parameters
type ActionClusterResourcePackageCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionClusterResourcePackageCreateMetaGlobalInput) SetIncludes(value string) *ActionClusterResourcePackageCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterResourcePackageCreateMetaGlobalInput) SetNo(value bool) *ActionClusterResourcePackageCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourcePackageCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourcePackageCreateMetaGlobalInput) SelectParameters(params ...string) *ActionClusterResourcePackageCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterResourcePackageCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterResourcePackageCreateInput is a type for action input parameters
type ActionClusterResourcePackageCreateInput struct {
	Label string `json:"label"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionClusterResourcePackageCreateInput) SetLabel(value string) *ActionClusterResourcePackageCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourcePackageCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourcePackageCreateInput) SelectParameters(params ...string) *ActionClusterResourcePackageCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterResourcePackageCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterResourcePackageCreateRequest is a type for the entire action request
type ActionClusterResourcePackageCreateRequest struct {
	ClusterResourcePackage map[string]interface{} `json:"cluster_resource_package"`
	Meta                   map[string]interface{} `json:"_meta"`
}

// ActionClusterResourcePackageCreateOutput is a type for action output parameters
type ActionClusterResourcePackageCreateOutput struct {
	CreatedAt   string                       `json:"created_at"`
	Environment *ActionEnvironmentShowOutput `json:"environment"`
	Id          int64                        `json:"id"`
	Label       string                       `json:"label"`
	UpdatedAt   string                       `json:"updated_at"`
	User        *ActionUserShowOutput        `json:"user"`
}

// Type for action response, including envelope
type ActionClusterResourcePackageCreateResponse struct {
	Action *ActionClusterResourcePackageCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ClusterResourcePackage *ActionClusterResourcePackageCreateOutput `json:"cluster_resource_package"`
	}

	// Action output without the namespace
	Output *ActionClusterResourcePackageCreateOutput
}

// Prepare the action for invocation
func (action *ActionClusterResourcePackageCreate) Prepare() *ActionClusterResourcePackageCreateInvocation {
	return &ActionClusterResourcePackageCreateInvocation{
		Action: action,
		Path:   "/v6.0/cluster_resource_packages",
	}
}

// ActionClusterResourcePackageCreateInvocation is used to configure action for invocation
type ActionClusterResourcePackageCreateInvocation struct {
	// Pointer to the action
	Action *ActionClusterResourcePackageCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionClusterResourcePackageCreateInput
	// Global meta input parameters
	MetaInput *ActionClusterResourcePackageCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionClusterResourcePackageCreateInvocation) NewInput() *ActionClusterResourcePackageCreateInput {
	inv.Input = &ActionClusterResourcePackageCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionClusterResourcePackageCreateInvocation) SetInput(input *ActionClusterResourcePackageCreateInput) *ActionClusterResourcePackageCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionClusterResourcePackageCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionClusterResourcePackageCreateInvocation) NewMetaInput() *ActionClusterResourcePackageCreateMetaGlobalInput {
	inv.MetaInput = &ActionClusterResourcePackageCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterResourcePackageCreateInvocation) SetMetaInput(input *ActionClusterResourcePackageCreateMetaGlobalInput) *ActionClusterResourcePackageCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterResourcePackageCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionClusterResourcePackageCreateInvocation) Call() (*ActionClusterResourcePackageCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionClusterResourcePackageCreateInvocation) callAsBody() (*ActionClusterResourcePackageCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionClusterResourcePackageCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.ClusterResourcePackage
	}
	return resp, err
}

func (inv *ActionClusterResourcePackageCreateInvocation) makeAllInputParams() *ActionClusterResourcePackageCreateRequest {
	return &ActionClusterResourcePackageCreateRequest{
		ClusterResourcePackage: inv.makeInputParams(),
		Meta:                   inv.makeMetaInputParams(),
	}
}

func (inv *ActionClusterResourcePackageCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
	}

	return ret
}

func (inv *ActionClusterResourcePackageCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
