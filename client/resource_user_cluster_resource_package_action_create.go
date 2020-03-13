package client

import (
)

// ActionUserClusterResourcePackageCreate is a type for action User_cluster_resource_package#Create
type ActionUserClusterResourcePackageCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserClusterResourcePackageCreate(client *Client) *ActionUserClusterResourcePackageCreate {
	return &ActionUserClusterResourcePackageCreate{
		Client: client,
	}
}

// ActionUserClusterResourcePackageCreateMetaGlobalInput is a type for action global meta input parameters
type ActionUserClusterResourcePackageCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserClusterResourcePackageCreateMetaGlobalInput) SetNo(value bool) *ActionUserClusterResourcePackageCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserClusterResourcePackageCreateMetaGlobalInput) SetIncludes(value string) *ActionUserClusterResourcePackageCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserClusterResourcePackageCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserClusterResourcePackageCreateMetaGlobalInput) SelectParameters(params ...string) *ActionUserClusterResourcePackageCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserClusterResourcePackageCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserClusterResourcePackageCreateInput is a type for action input parameters
type ActionUserClusterResourcePackageCreateInput struct {
	Environment int64 `json:"environment"`
	User int64 `json:"user"`
	ClusterResourcePackage int64 `json:"cluster_resource_package"`
	Comment string `json:"comment"`
	FromPersonal bool `json:"from_personal"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionUserClusterResourcePackageCreateInput) SetEnvironment(value int64) *ActionUserClusterResourcePackageCreateInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Environment"] = nil
	return in
}
// SetUser sets parameter User to value and selects it for sending
func (in *ActionUserClusterResourcePackageCreateInput) SetUser(value int64) *ActionUserClusterResourcePackageCreateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}
// SetClusterResourcePackage sets parameter ClusterResourcePackage to value and selects it for sending
func (in *ActionUserClusterResourcePackageCreateInput) SetClusterResourcePackage(value int64) *ActionUserClusterResourcePackageCreateInput {
	in.ClusterResourcePackage = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClusterResourcePackage"] = nil
	return in
}
// SetComment sets parameter Comment to value and selects it for sending
func (in *ActionUserClusterResourcePackageCreateInput) SetComment(value string) *ActionUserClusterResourcePackageCreateInput {
	in.Comment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Comment"] = nil
	return in
}
// SetFromPersonal sets parameter FromPersonal to value and selects it for sending
func (in *ActionUserClusterResourcePackageCreateInput) SetFromPersonal(value bool) *ActionUserClusterResourcePackageCreateInput {
	in.FromPersonal = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromPersonal"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserClusterResourcePackageCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserClusterResourcePackageCreateInput) SelectParameters(params ...string) *ActionUserClusterResourcePackageCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserClusterResourcePackageCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserClusterResourcePackageCreateRequest is a type for the entire action request
type ActionUserClusterResourcePackageCreateRequest struct {
	UserClusterResourcePackage map[string]interface{} `json:"user_cluster_resource_package"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionUserClusterResourcePackageCreateOutput is a type for action output parameters
type ActionUserClusterResourcePackageCreateOutput struct {
	Id int64 `json:"id"`
	Environment *ActionEnvironmentShowOutput `json:"environment"`
	User *ActionUserShowOutput `json:"user"`
	ClusterResourcePackage *ActionClusterResourcePackageShowOutput `json:"cluster_resource_package"`
	AddedBy *ActionUserShowOutput `json:"added_by"`
	Label string `json:"label"`
	IsPersonal bool `json:"is_personal"`
	Comment string `json:"comment"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


// Type for action response, including envelope
type ActionUserClusterResourcePackageCreateResponse struct {
	Action *ActionUserClusterResourcePackageCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserClusterResourcePackage *ActionUserClusterResourcePackageCreateOutput `json:"user_cluster_resource_package"`
	}

	// Action output without the namespace
	Output *ActionUserClusterResourcePackageCreateOutput
}


// Prepare the action for invocation
func (action *ActionUserClusterResourcePackageCreate) Prepare() *ActionUserClusterResourcePackageCreateInvocation {
	return &ActionUserClusterResourcePackageCreateInvocation{
		Action: action,
		Path: "/v6.0/user_cluster_resource_packages",
	}
}

// ActionUserClusterResourcePackageCreateInvocation is used to configure action for invocation
type ActionUserClusterResourcePackageCreateInvocation struct {
	// Pointer to the action
	Action *ActionUserClusterResourcePackageCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserClusterResourcePackageCreateInput
	// Global meta input parameters
	MetaInput *ActionUserClusterResourcePackageCreateMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserClusterResourcePackageCreateInvocation) NewInput() *ActionUserClusterResourcePackageCreateInput {
	inv.Input = &ActionUserClusterResourcePackageCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserClusterResourcePackageCreateInvocation) SetInput(input *ActionUserClusterResourcePackageCreateInput) *ActionUserClusterResourcePackageCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserClusterResourcePackageCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserClusterResourcePackageCreateInvocation) NewMetaInput() *ActionUserClusterResourcePackageCreateMetaGlobalInput {
	inv.MetaInput = &ActionUserClusterResourcePackageCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserClusterResourcePackageCreateInvocation) SetMetaInput(input *ActionUserClusterResourcePackageCreateMetaGlobalInput) *ActionUserClusterResourcePackageCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserClusterResourcePackageCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserClusterResourcePackageCreateInvocation) Call() (*ActionUserClusterResourcePackageCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionUserClusterResourcePackageCreateInvocation) callAsBody() (*ActionUserClusterResourcePackageCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserClusterResourcePackageCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserClusterResourcePackage
	}
	return resp, err
}




func (inv *ActionUserClusterResourcePackageCreateInvocation) makeAllInputParams() *ActionUserClusterResourcePackageCreateRequest {
	return &ActionUserClusterResourcePackageCreateRequest{
		UserClusterResourcePackage: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserClusterResourcePackageCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			ret["environment"] = inv.Input.Environment
		}
		if inv.IsParameterSelected("User") {
			ret["user"] = inv.Input.User
		}
		if inv.IsParameterSelected("ClusterResourcePackage") {
			ret["cluster_resource_package"] = inv.Input.ClusterResourcePackage
		}
		if inv.IsParameterSelected("Comment") {
			ret["comment"] = inv.Input.Comment
		}
		if inv.IsParameterSelected("FromPersonal") {
			ret["from_personal"] = inv.Input.FromPersonal
		}
	}

	return ret
}

func (inv *ActionUserClusterResourcePackageCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
