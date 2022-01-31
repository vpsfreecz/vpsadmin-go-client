package client

import ()

// ActionClusterResourceIndex is a type for action Cluster_resource#Index
type ActionClusterResourceIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterResourceIndex(client *Client) *ActionClusterResourceIndex {
	return &ActionClusterResourceIndex{
		Client: client,
	}
}

// ActionClusterResourceIndexMetaGlobalInput is a type for action global meta input parameters
type ActionClusterResourceIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionClusterResourceIndexMetaGlobalInput) SetCount(value bool) *ActionClusterResourceIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionClusterResourceIndexMetaGlobalInput) SetIncludes(value string) *ActionClusterResourceIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterResourceIndexMetaGlobalInput) SetNo(value bool) *ActionClusterResourceIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourceIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourceIndexMetaGlobalInput) SelectParameters(params ...string) *ActionClusterResourceIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterResourceIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterResourceIndexInput is a type for action input parameters
type ActionClusterResourceIndexInput struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionClusterResourceIndexInput) SetLimit(value int64) *ActionClusterResourceIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionClusterResourceIndexInput) SetOffset(value int64) *ActionClusterResourceIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourceIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourceIndexInput) SelectParameters(params ...string) *ActionClusterResourceIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionClusterResourceIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionClusterResourceIndexInput) UnselectParameters(params ...string) *ActionClusterResourceIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionClusterResourceIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterResourceIndexOutput is a type for action output parameters
type ActionClusterResourceIndexOutput struct {
	Id       int64  `json:"id"`
	Label    string `json:"label"`
	Max      int64  `json:"max"`
	Min      int64  `json:"min"`
	Name     string `json:"name"`
	Stepsize int64  `json:"stepsize"`
}

// Type for action response, including envelope
type ActionClusterResourceIndexResponse struct {
	Action *ActionClusterResourceIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ClusterResources []*ActionClusterResourceIndexOutput `json:"cluster_resources"`
	}

	// Action output without the namespace
	Output []*ActionClusterResourceIndexOutput
}

// Prepare the action for invocation
func (action *ActionClusterResourceIndex) Prepare() *ActionClusterResourceIndexInvocation {
	return &ActionClusterResourceIndexInvocation{
		Action: action,
		Path:   "/v6.0/cluster_resources",
	}
}

// ActionClusterResourceIndexInvocation is used to configure action for invocation
type ActionClusterResourceIndexInvocation struct {
	// Pointer to the action
	Action *ActionClusterResourceIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionClusterResourceIndexInput
	// Global meta input parameters
	MetaInput *ActionClusterResourceIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionClusterResourceIndexInvocation) NewInput() *ActionClusterResourceIndexInput {
	inv.Input = &ActionClusterResourceIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionClusterResourceIndexInvocation) SetInput(input *ActionClusterResourceIndexInput) *ActionClusterResourceIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionClusterResourceIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionClusterResourceIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionClusterResourceIndexInvocation) NewMetaInput() *ActionClusterResourceIndexMetaGlobalInput {
	inv.MetaInput = &ActionClusterResourceIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterResourceIndexInvocation) SetMetaInput(input *ActionClusterResourceIndexMetaGlobalInput) *ActionClusterResourceIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterResourceIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionClusterResourceIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionClusterResourceIndexInvocation) Call() (*ActionClusterResourceIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionClusterResourceIndexInvocation) callAsQuery() (*ActionClusterResourceIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionClusterResourceIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.ClusterResources
	}
	return resp, err
}

func (inv *ActionClusterResourceIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["cluster_resource[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["cluster_resource[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
	}
}

func (inv *ActionClusterResourceIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
