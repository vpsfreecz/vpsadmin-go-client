package client

import (
	"strings"
)

// ActionUserClusterResourcePackageUpdate is a type for action User_cluster_resource_package#Update
type ActionUserClusterResourcePackageUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserClusterResourcePackageUpdate(client *Client) *ActionUserClusterResourcePackageUpdate {
	return &ActionUserClusterResourcePackageUpdate{
		Client: client,
	}
}

// ActionUserClusterResourcePackageUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionUserClusterResourcePackageUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserClusterResourcePackageUpdateMetaGlobalInput) SetIncludes(value string) *ActionUserClusterResourcePackageUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserClusterResourcePackageUpdateMetaGlobalInput) SetNo(value bool) *ActionUserClusterResourcePackageUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserClusterResourcePackageUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserClusterResourcePackageUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionUserClusterResourcePackageUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserClusterResourcePackageUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserClusterResourcePackageUpdateInput is a type for action input parameters
type ActionUserClusterResourcePackageUpdateInput struct {
	Comment string `json:"comment"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetComment sets parameter Comment to value and selects it for sending
func (in *ActionUserClusterResourcePackageUpdateInput) SetComment(value string) *ActionUserClusterResourcePackageUpdateInput {
	in.Comment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Comment"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserClusterResourcePackageUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserClusterResourcePackageUpdateInput) SelectParameters(params ...string) *ActionUserClusterResourcePackageUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserClusterResourcePackageUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserClusterResourcePackageUpdateRequest is a type for the entire action request
type ActionUserClusterResourcePackageUpdateRequest struct {
	UserClusterResourcePackage map[string]interface{} `json:"user_cluster_resource_package"`
	Meta                       map[string]interface{} `json:"_meta"`
}

// ActionUserClusterResourcePackageUpdateOutput is a type for action output parameters
type ActionUserClusterResourcePackageUpdateOutput struct {
	AddedBy                *ActionUserShowOutput                   `json:"added_by"`
	ClusterResourcePackage *ActionClusterResourcePackageShowOutput `json:"cluster_resource_package"`
	Comment                string                                  `json:"comment"`
	CreatedAt              string                                  `json:"created_at"`
	Environment            *ActionEnvironmentShowOutput            `json:"environment"`
	Id                     int64                                   `json:"id"`
	IsPersonal             bool                                    `json:"is_personal"`
	Label                  string                                  `json:"label"`
	UpdatedAt              string                                  `json:"updated_at"`
	User                   *ActionUserShowOutput                   `json:"user"`
}

// Type for action response, including envelope
type ActionUserClusterResourcePackageUpdateResponse struct {
	Action *ActionUserClusterResourcePackageUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserClusterResourcePackage *ActionUserClusterResourcePackageUpdateOutput `json:"user_cluster_resource_package"`
	}

	// Action output without the namespace
	Output *ActionUserClusterResourcePackageUpdateOutput
}

// Prepare the action for invocation
func (action *ActionUserClusterResourcePackageUpdate) Prepare() *ActionUserClusterResourcePackageUpdateInvocation {
	return &ActionUserClusterResourcePackageUpdateInvocation{
		Action: action,
		Path:   "/v6.0/user_cluster_resource_packages/{user_cluster_resource_package_id}",
	}
}

// ActionUserClusterResourcePackageUpdateInvocation is used to configure action for invocation
type ActionUserClusterResourcePackageUpdateInvocation struct {
	// Pointer to the action
	Action *ActionUserClusterResourcePackageUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserClusterResourcePackageUpdateInput
	// Global meta input parameters
	MetaInput *ActionUserClusterResourcePackageUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserClusterResourcePackageUpdateInvocation) SetPathParamInt(param string, value int64) *ActionUserClusterResourcePackageUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserClusterResourcePackageUpdateInvocation) SetPathParamString(param string, value string) *ActionUserClusterResourcePackageUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserClusterResourcePackageUpdateInvocation) NewInput() *ActionUserClusterResourcePackageUpdateInput {
	inv.Input = &ActionUserClusterResourcePackageUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserClusterResourcePackageUpdateInvocation) SetInput(input *ActionUserClusterResourcePackageUpdateInput) *ActionUserClusterResourcePackageUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserClusterResourcePackageUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserClusterResourcePackageUpdateInvocation) NewMetaInput() *ActionUserClusterResourcePackageUpdateMetaGlobalInput {
	inv.MetaInput = &ActionUserClusterResourcePackageUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserClusterResourcePackageUpdateInvocation) SetMetaInput(input *ActionUserClusterResourcePackageUpdateMetaGlobalInput) *ActionUserClusterResourcePackageUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserClusterResourcePackageUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserClusterResourcePackageUpdateInvocation) Call() (*ActionUserClusterResourcePackageUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionUserClusterResourcePackageUpdateInvocation) callAsBody() (*ActionUserClusterResourcePackageUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserClusterResourcePackageUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserClusterResourcePackage
	}
	return resp, err
}

func (inv *ActionUserClusterResourcePackageUpdateInvocation) makeAllInputParams() *ActionUserClusterResourcePackageUpdateRequest {
	return &ActionUserClusterResourcePackageUpdateRequest{
		UserClusterResourcePackage: inv.makeInputParams(),
		Meta:                       inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserClusterResourcePackageUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Comment") {
			ret["comment"] = inv.Input.Comment
		}
	}

	return ret
}

func (inv *ActionUserClusterResourcePackageUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
