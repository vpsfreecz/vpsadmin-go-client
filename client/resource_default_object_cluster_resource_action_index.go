package client

import ()

// ActionDefaultObjectClusterResourceIndex is a type for action Default_object_cluster_resource#Index
type ActionDefaultObjectClusterResourceIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionDefaultObjectClusterResourceIndex(client *Client) *ActionDefaultObjectClusterResourceIndex {
	return &ActionDefaultObjectClusterResourceIndex{
		Client: client,
	}
}

// ActionDefaultObjectClusterResourceIndexMetaGlobalInput is a type for action global meta input parameters
type ActionDefaultObjectClusterResourceIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionDefaultObjectClusterResourceIndexMetaGlobalInput) SetCount(value bool) *ActionDefaultObjectClusterResourceIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDefaultObjectClusterResourceIndexMetaGlobalInput) SetIncludes(value string) *ActionDefaultObjectClusterResourceIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDefaultObjectClusterResourceIndexMetaGlobalInput) SetNo(value bool) *ActionDefaultObjectClusterResourceIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDefaultObjectClusterResourceIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDefaultObjectClusterResourceIndexMetaGlobalInput) SelectParameters(params ...string) *ActionDefaultObjectClusterResourceIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDefaultObjectClusterResourceIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDefaultObjectClusterResourceIndexInput is a type for action input parameters
type ActionDefaultObjectClusterResourceIndexInput struct {
	ClassName       string `json:"class_name"`
	ClusterResource int64  `json:"cluster_resource"`
	Environment     int64  `json:"environment"`
	FromId          int64  `json:"from_id"`
	Limit           int64  `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetClassName sets parameter ClassName to value and selects it for sending
func (in *ActionDefaultObjectClusterResourceIndexInput) SetClassName(value string) *ActionDefaultObjectClusterResourceIndexInput {
	in.ClassName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClassName"] = nil
	return in
}

// SetClusterResource sets parameter ClusterResource to value and selects it for sending
func (in *ActionDefaultObjectClusterResourceIndexInput) SetClusterResource(value int64) *ActionDefaultObjectClusterResourceIndexInput {
	in.ClusterResource = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetClusterResourceNil(false)
	in._selectedParameters["ClusterResource"] = nil
	return in
}

// SetClusterResourceNil sets parameter ClusterResource to nil and selects it for sending
func (in *ActionDefaultObjectClusterResourceIndexInput) SetClusterResourceNil(set bool) *ActionDefaultObjectClusterResourceIndexInput {
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
func (in *ActionDefaultObjectClusterResourceIndexInput) SetEnvironment(value int64) *ActionDefaultObjectClusterResourceIndexInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEnvironmentNil(false)
	in._selectedParameters["Environment"] = nil
	return in
}

// SetEnvironmentNil sets parameter Environment to nil and selects it for sending
func (in *ActionDefaultObjectClusterResourceIndexInput) SetEnvironmentNil(set bool) *ActionDefaultObjectClusterResourceIndexInput {
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
func (in *ActionDefaultObjectClusterResourceIndexInput) SetFromId(value int64) *ActionDefaultObjectClusterResourceIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDefaultObjectClusterResourceIndexInput) SetLimit(value int64) *ActionDefaultObjectClusterResourceIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionDefaultObjectClusterResourceIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDefaultObjectClusterResourceIndexInput) SelectParameters(params ...string) *ActionDefaultObjectClusterResourceIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDefaultObjectClusterResourceIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDefaultObjectClusterResourceIndexInput) UnselectParameters(params ...string) *ActionDefaultObjectClusterResourceIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDefaultObjectClusterResourceIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDefaultObjectClusterResourceIndexOutput is a type for action output parameters
type ActionDefaultObjectClusterResourceIndexOutput struct {
	ClassName       string                           `json:"class_name"`
	ClusterResource *ActionClusterResourceShowOutput `json:"cluster_resource"`
	Environment     *ActionEnvironmentShowOutput     `json:"environment"`
	Id              int64                            `json:"id"`
	Value           int64                            `json:"value"`
}

// Type for action response, including envelope
type ActionDefaultObjectClusterResourceIndexResponse struct {
	Action *ActionDefaultObjectClusterResourceIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DefaultObjectClusterResources []*ActionDefaultObjectClusterResourceIndexOutput `json:"default_object_cluster_resources"`
	}

	// Action output without the namespace
	Output []*ActionDefaultObjectClusterResourceIndexOutput
}

// Prepare the action for invocation
func (action *ActionDefaultObjectClusterResourceIndex) Prepare() *ActionDefaultObjectClusterResourceIndexInvocation {
	return &ActionDefaultObjectClusterResourceIndexInvocation{
		Action: action,
		Path:   "/v7.0/default_object_cluster_resources",
	}
}

// ActionDefaultObjectClusterResourceIndexInvocation is used to configure action for invocation
type ActionDefaultObjectClusterResourceIndexInvocation struct {
	// Pointer to the action
	Action *ActionDefaultObjectClusterResourceIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDefaultObjectClusterResourceIndexInput
	// Global meta input parameters
	MetaInput *ActionDefaultObjectClusterResourceIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDefaultObjectClusterResourceIndexInvocation) NewInput() *ActionDefaultObjectClusterResourceIndexInput {
	inv.Input = &ActionDefaultObjectClusterResourceIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDefaultObjectClusterResourceIndexInvocation) SetInput(input *ActionDefaultObjectClusterResourceIndexInput) *ActionDefaultObjectClusterResourceIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDefaultObjectClusterResourceIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDefaultObjectClusterResourceIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDefaultObjectClusterResourceIndexInvocation) NewMetaInput() *ActionDefaultObjectClusterResourceIndexMetaGlobalInput {
	inv.MetaInput = &ActionDefaultObjectClusterResourceIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDefaultObjectClusterResourceIndexInvocation) SetMetaInput(input *ActionDefaultObjectClusterResourceIndexMetaGlobalInput) *ActionDefaultObjectClusterResourceIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDefaultObjectClusterResourceIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDefaultObjectClusterResourceIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDefaultObjectClusterResourceIndexInvocation) Call() (*ActionDefaultObjectClusterResourceIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDefaultObjectClusterResourceIndexInvocation) callAsQuery() (*ActionDefaultObjectClusterResourceIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDefaultObjectClusterResourceIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DefaultObjectClusterResources
	}
	return resp, err
}

func (inv *ActionDefaultObjectClusterResourceIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("ClassName") {
			ret["default_object_cluster_resource[class_name]"] = inv.Input.ClassName
		}
		if inv.IsParameterSelected("ClusterResource") {
			ret["default_object_cluster_resource[cluster_resource]"] = convertInt64ToString(inv.Input.ClusterResource)
		}
		if inv.IsParameterSelected("Environment") {
			ret["default_object_cluster_resource[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("FromId") {
			ret["default_object_cluster_resource[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["default_object_cluster_resource[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionDefaultObjectClusterResourceIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
