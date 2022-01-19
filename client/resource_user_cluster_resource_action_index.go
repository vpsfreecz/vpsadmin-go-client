package client

import (
	"strings"
)

// ActionUserClusterResourceIndex is a type for action User.Cluster_resource#Index
type ActionUserClusterResourceIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserClusterResourceIndex(client *Client) *ActionUserClusterResourceIndex {
	return &ActionUserClusterResourceIndex{
		Client: client,
	}
}

// ActionUserClusterResourceIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserClusterResourceIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserClusterResourceIndexMetaGlobalInput) SetCount(value bool) *ActionUserClusterResourceIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserClusterResourceIndexMetaGlobalInput) SetIncludes(value string) *ActionUserClusterResourceIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserClusterResourceIndexMetaGlobalInput) SetNo(value bool) *ActionUserClusterResourceIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserClusterResourceIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserClusterResourceIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserClusterResourceIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserClusterResourceIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserClusterResourceIndexInput is a type for action input parameters
type ActionUserClusterResourceIndexInput struct {
	Environment int64 `json:"environment"`
	Limit       int64 `json:"limit"`
	Offset      int64 `json:"offset"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionUserClusterResourceIndexInput) SetEnvironment(value int64) *ActionUserClusterResourceIndexInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Environment"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserClusterResourceIndexInput) SetLimit(value int64) *ActionUserClusterResourceIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionUserClusterResourceIndexInput) SetOffset(value int64) *ActionUserClusterResourceIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserClusterResourceIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserClusterResourceIndexInput) SelectParameters(params ...string) *ActionUserClusterResourceIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserClusterResourceIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserClusterResourceIndexOutput is a type for action output parameters
type ActionUserClusterResourceIndexOutput struct {
	ClusterResource *ActionClusterResourceShowOutput `json:"cluster_resource"`
	Environment     *ActionEnvironmentShowOutput     `json:"environment"`
	Free            int64                            `json:"free"`
	Id              int64                            `json:"id"`
	Used            int64                            `json:"used"`
	Value           int64                            `json:"value"`
}

// Type for action response, including envelope
type ActionUserClusterResourceIndexResponse struct {
	Action *ActionUserClusterResourceIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ClusterResources []*ActionUserClusterResourceIndexOutput `json:"cluster_resources"`
	}

	// Action output without the namespace
	Output []*ActionUserClusterResourceIndexOutput
}

// Prepare the action for invocation
func (action *ActionUserClusterResourceIndex) Prepare() *ActionUserClusterResourceIndexInvocation {
	return &ActionUserClusterResourceIndexInvocation{
		Action: action,
		Path:   "/v6.0/users/{user_id}/cluster_resources",
	}
}

// ActionUserClusterResourceIndexInvocation is used to configure action for invocation
type ActionUserClusterResourceIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserClusterResourceIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserClusterResourceIndexInput
	// Global meta input parameters
	MetaInput *ActionUserClusterResourceIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserClusterResourceIndexInvocation) SetPathParamInt(param string, value int64) *ActionUserClusterResourceIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserClusterResourceIndexInvocation) SetPathParamString(param string, value string) *ActionUserClusterResourceIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserClusterResourceIndexInvocation) NewInput() *ActionUserClusterResourceIndexInput {
	inv.Input = &ActionUserClusterResourceIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserClusterResourceIndexInvocation) SetInput(input *ActionUserClusterResourceIndexInput) *ActionUserClusterResourceIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserClusterResourceIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserClusterResourceIndexInvocation) NewMetaInput() *ActionUserClusterResourceIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserClusterResourceIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserClusterResourceIndexInvocation) SetMetaInput(input *ActionUserClusterResourceIndexMetaGlobalInput) *ActionUserClusterResourceIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserClusterResourceIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserClusterResourceIndexInvocation) Call() (*ActionUserClusterResourceIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserClusterResourceIndexInvocation) callAsQuery() (*ActionUserClusterResourceIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserClusterResourceIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.ClusterResources
	}
	return resp, err
}

func (inv *ActionUserClusterResourceIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			ret["cluster_resource[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("Limit") {
			ret["cluster_resource[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["cluster_resource[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
	}
}

func (inv *ActionUserClusterResourceIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
