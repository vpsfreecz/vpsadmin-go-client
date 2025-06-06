package client

import ()

// ActionUserClusterResourcePackageIndex is a type for action User_cluster_resource_package#Index
type ActionUserClusterResourcePackageIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserClusterResourcePackageIndex(client *Client) *ActionUserClusterResourcePackageIndex {
	return &ActionUserClusterResourcePackageIndex{
		Client: client,
	}
}

// ActionUserClusterResourcePackageIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserClusterResourcePackageIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserClusterResourcePackageIndexMetaGlobalInput) SetCount(value bool) *ActionUserClusterResourcePackageIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserClusterResourcePackageIndexMetaGlobalInput) SetIncludes(value string) *ActionUserClusterResourcePackageIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserClusterResourcePackageIndexMetaGlobalInput) SetNo(value bool) *ActionUserClusterResourcePackageIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserClusterResourcePackageIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserClusterResourcePackageIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserClusterResourcePackageIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserClusterResourcePackageIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserClusterResourcePackageIndexInput is a type for action input parameters
type ActionUserClusterResourcePackageIndexInput struct {
	AddedBy                int64 `json:"added_by"`
	ClusterResourcePackage int64 `json:"cluster_resource_package"`
	Environment            int64 `json:"environment"`
	FromId                 int64 `json:"from_id"`
	Limit                  int64 `json:"limit"`
	User                   int64 `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAddedBy sets parameter AddedBy to value and selects it for sending
func (in *ActionUserClusterResourcePackageIndexInput) SetAddedBy(value int64) *ActionUserClusterResourcePackageIndexInput {
	in.AddedBy = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetAddedByNil(false)
	in._selectedParameters["AddedBy"] = nil
	return in
}

// SetAddedByNil sets parameter AddedBy to nil and selects it for sending
func (in *ActionUserClusterResourcePackageIndexInput) SetAddedByNil(set bool) *ActionUserClusterResourcePackageIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["AddedBy"] = nil
		in.SelectParameters("AddedBy")
	} else {
		delete(in._nilParameters, "AddedBy")
	}
	return in
}

// SetClusterResourcePackage sets parameter ClusterResourcePackage to value and selects it for sending
func (in *ActionUserClusterResourcePackageIndexInput) SetClusterResourcePackage(value int64) *ActionUserClusterResourcePackageIndexInput {
	in.ClusterResourcePackage = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetClusterResourcePackageNil(false)
	in._selectedParameters["ClusterResourcePackage"] = nil
	return in
}

// SetClusterResourcePackageNil sets parameter ClusterResourcePackage to nil and selects it for sending
func (in *ActionUserClusterResourcePackageIndexInput) SetClusterResourcePackageNil(set bool) *ActionUserClusterResourcePackageIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["ClusterResourcePackage"] = nil
		in.SelectParameters("ClusterResourcePackage")
	} else {
		delete(in._nilParameters, "ClusterResourcePackage")
	}
	return in
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionUserClusterResourcePackageIndexInput) SetEnvironment(value int64) *ActionUserClusterResourcePackageIndexInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEnvironmentNil(false)
	in._selectedParameters["Environment"] = nil
	return in
}

// SetEnvironmentNil sets parameter Environment to nil and selects it for sending
func (in *ActionUserClusterResourcePackageIndexInput) SetEnvironmentNil(set bool) *ActionUserClusterResourcePackageIndexInput {
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

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionUserClusterResourcePackageIndexInput) SetFromId(value int64) *ActionUserClusterResourcePackageIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserClusterResourcePackageIndexInput) SetLimit(value int64) *ActionUserClusterResourcePackageIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionUserClusterResourcePackageIndexInput) SetUser(value int64) *ActionUserClusterResourcePackageIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionUserClusterResourcePackageIndexInput) SetUserNil(set bool) *ActionUserClusterResourcePackageIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["User"] = nil
		in.SelectParameters("User")
	} else {
		delete(in._nilParameters, "User")
	}
	return in
}

// SelectParameters sets parameters from ActionUserClusterResourcePackageIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserClusterResourcePackageIndexInput) SelectParameters(params ...string) *ActionUserClusterResourcePackageIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserClusterResourcePackageIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserClusterResourcePackageIndexInput) UnselectParameters(params ...string) *ActionUserClusterResourcePackageIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserClusterResourcePackageIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserClusterResourcePackageIndexOutput is a type for action output parameters
type ActionUserClusterResourcePackageIndexOutput struct {
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
type ActionUserClusterResourcePackageIndexResponse struct {
	Action *ActionUserClusterResourcePackageIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserClusterResourcePackages []*ActionUserClusterResourcePackageIndexOutput `json:"user_cluster_resource_packages"`
	}

	// Action output without the namespace
	Output []*ActionUserClusterResourcePackageIndexOutput
}

// Prepare the action for invocation
func (action *ActionUserClusterResourcePackageIndex) Prepare() *ActionUserClusterResourcePackageIndexInvocation {
	return &ActionUserClusterResourcePackageIndexInvocation{
		Action: action,
		Path:   "/v7.0/user_cluster_resource_packages",
	}
}

// ActionUserClusterResourcePackageIndexInvocation is used to configure action for invocation
type ActionUserClusterResourcePackageIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserClusterResourcePackageIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserClusterResourcePackageIndexInput
	// Global meta input parameters
	MetaInput *ActionUserClusterResourcePackageIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserClusterResourcePackageIndexInvocation) NewInput() *ActionUserClusterResourcePackageIndexInput {
	inv.Input = &ActionUserClusterResourcePackageIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserClusterResourcePackageIndexInvocation) SetInput(input *ActionUserClusterResourcePackageIndexInput) *ActionUserClusterResourcePackageIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserClusterResourcePackageIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserClusterResourcePackageIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserClusterResourcePackageIndexInvocation) NewMetaInput() *ActionUserClusterResourcePackageIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserClusterResourcePackageIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserClusterResourcePackageIndexInvocation) SetMetaInput(input *ActionUserClusterResourcePackageIndexMetaGlobalInput) *ActionUserClusterResourcePackageIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserClusterResourcePackageIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserClusterResourcePackageIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserClusterResourcePackageIndexInvocation) Call() (*ActionUserClusterResourcePackageIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserClusterResourcePackageIndexInvocation) callAsQuery() (*ActionUserClusterResourcePackageIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserClusterResourcePackageIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserClusterResourcePackages
	}
	return resp, err
}

func (inv *ActionUserClusterResourcePackageIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("AddedBy") {
			ret["user_cluster_resource_package[added_by]"] = convertInt64ToString(inv.Input.AddedBy)
		}
		if inv.IsParameterSelected("ClusterResourcePackage") {
			ret["user_cluster_resource_package[cluster_resource_package]"] = convertInt64ToString(inv.Input.ClusterResourcePackage)
		}
		if inv.IsParameterSelected("Environment") {
			ret["user_cluster_resource_package[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("FromId") {
			ret["user_cluster_resource_package[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["user_cluster_resource_package[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("User") {
			ret["user_cluster_resource_package[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionUserClusterResourcePackageIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
