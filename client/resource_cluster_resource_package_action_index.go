package client

import ()

// ActionClusterResourcePackageIndex is a type for action Cluster_resource_package#Index
type ActionClusterResourcePackageIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterResourcePackageIndex(client *Client) *ActionClusterResourcePackageIndex {
	return &ActionClusterResourcePackageIndex{
		Client: client,
	}
}

// ActionClusterResourcePackageIndexMetaGlobalInput is a type for action global meta input parameters
type ActionClusterResourcePackageIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionClusterResourcePackageIndexMetaGlobalInput) SetCount(value bool) *ActionClusterResourcePackageIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionClusterResourcePackageIndexMetaGlobalInput) SetIncludes(value string) *ActionClusterResourcePackageIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterResourcePackageIndexMetaGlobalInput) SetNo(value bool) *ActionClusterResourcePackageIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourcePackageIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourcePackageIndexMetaGlobalInput) SelectParameters(params ...string) *ActionClusterResourcePackageIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterResourcePackageIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterResourcePackageIndexInput is a type for action input parameters
type ActionClusterResourcePackageIndexInput struct {
	Environment int64 `json:"environment"`
	FromId      int64 `json:"from_id"`
	Limit       int64 `json:"limit"`
	User        int64 `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionClusterResourcePackageIndexInput) SetEnvironment(value int64) *ActionClusterResourcePackageIndexInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEnvironmentNil(false)
	in._selectedParameters["Environment"] = nil
	return in
}

// SetEnvironmentNil sets parameter Environment to nil and selects it for sending
func (in *ActionClusterResourcePackageIndexInput) SetEnvironmentNil(set bool) *ActionClusterResourcePackageIndexInput {
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
func (in *ActionClusterResourcePackageIndexInput) SetFromId(value int64) *ActionClusterResourcePackageIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionClusterResourcePackageIndexInput) SetLimit(value int64) *ActionClusterResourcePackageIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionClusterResourcePackageIndexInput) SetUser(value int64) *ActionClusterResourcePackageIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionClusterResourcePackageIndexInput) SetUserNil(set bool) *ActionClusterResourcePackageIndexInput {
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

// SelectParameters sets parameters from ActionClusterResourcePackageIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourcePackageIndexInput) SelectParameters(params ...string) *ActionClusterResourcePackageIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionClusterResourcePackageIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionClusterResourcePackageIndexInput) UnselectParameters(params ...string) *ActionClusterResourcePackageIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionClusterResourcePackageIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterResourcePackageIndexOutput is a type for action output parameters
type ActionClusterResourcePackageIndexOutput struct {
	CreatedAt   string                       `json:"created_at"`
	Environment *ActionEnvironmentShowOutput `json:"environment"`
	Id          int64                        `json:"id"`
	Label       string                       `json:"label"`
	UpdatedAt   string                       `json:"updated_at"`
	User        *ActionUserShowOutput        `json:"user"`
}

// Type for action response, including envelope
type ActionClusterResourcePackageIndexResponse struct {
	Action *ActionClusterResourcePackageIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ClusterResourcePackages []*ActionClusterResourcePackageIndexOutput `json:"cluster_resource_packages"`
	}

	// Action output without the namespace
	Output []*ActionClusterResourcePackageIndexOutput
}

// Prepare the action for invocation
func (action *ActionClusterResourcePackageIndex) Prepare() *ActionClusterResourcePackageIndexInvocation {
	return &ActionClusterResourcePackageIndexInvocation{
		Action: action,
		Path:   "/v7.0/cluster_resource_packages",
	}
}

// ActionClusterResourcePackageIndexInvocation is used to configure action for invocation
type ActionClusterResourcePackageIndexInvocation struct {
	// Pointer to the action
	Action *ActionClusterResourcePackageIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionClusterResourcePackageIndexInput
	// Global meta input parameters
	MetaInput *ActionClusterResourcePackageIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionClusterResourcePackageIndexInvocation) NewInput() *ActionClusterResourcePackageIndexInput {
	inv.Input = &ActionClusterResourcePackageIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionClusterResourcePackageIndexInvocation) SetInput(input *ActionClusterResourcePackageIndexInput) *ActionClusterResourcePackageIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionClusterResourcePackageIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionClusterResourcePackageIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionClusterResourcePackageIndexInvocation) NewMetaInput() *ActionClusterResourcePackageIndexMetaGlobalInput {
	inv.MetaInput = &ActionClusterResourcePackageIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterResourcePackageIndexInvocation) SetMetaInput(input *ActionClusterResourcePackageIndexMetaGlobalInput) *ActionClusterResourcePackageIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterResourcePackageIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionClusterResourcePackageIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionClusterResourcePackageIndexInvocation) Call() (*ActionClusterResourcePackageIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionClusterResourcePackageIndexInvocation) callAsQuery() (*ActionClusterResourcePackageIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionClusterResourcePackageIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.ClusterResourcePackages
	}
	return resp, err
}

func (inv *ActionClusterResourcePackageIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			ret["cluster_resource_package[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("FromId") {
			ret["cluster_resource_package[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["cluster_resource_package[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("User") {
			ret["cluster_resource_package[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionClusterResourcePackageIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
