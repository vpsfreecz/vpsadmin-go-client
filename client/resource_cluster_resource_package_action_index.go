package client

import (
)

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
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	Environment int64 `json:"environment"`
	User int64 `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionClusterResourcePackageIndexInput) SetOffset(value int64) *ActionClusterResourcePackageIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
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
// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionClusterResourcePackageIndexInput) SetEnvironment(value int64) *ActionClusterResourcePackageIndexInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Environment"] = nil
	return in
}
// SetUser sets parameter User to value and selects it for sending
func (in *ActionClusterResourcePackageIndexInput) SetUser(value int64) *ActionClusterResourcePackageIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
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

func (in *ActionClusterResourcePackageIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionClusterResourcePackageIndexOutput is a type for action output parameters
type ActionClusterResourcePackageIndexOutput struct {
	Id int64 `json:"id"`
	Label string `json:"label"`
	Environment *ActionEnvironmentShowOutput `json:"environment"`
	User *ActionUserShowOutput `json:"user"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
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
		Path: "/v6.0/cluster_resource_packages",
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
		if inv.IsParameterSelected("Offset") {
			ret["cluster_resource_package[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["cluster_resource_package[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Environment") {
			ret["cluster_resource_package[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("User") {
			ret["cluster_resource_package[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionClusterResourcePackageIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

