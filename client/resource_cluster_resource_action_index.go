package client

import (
)

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
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionClusterResourceIndexInput) SetLimit(value int64) *ActionClusterResourceIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
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

func (in *ActionClusterResourceIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionClusterResourceIndexOutput is a type for action output parameters
type ActionClusterResourceIndexOutput struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Label string `json:"label"`
	Min int64 `json:"min"`
	Max int64 `json:"max"`
	Stepsize int64 `json:"stepsize"`
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
		Path: "/v5.0/cluster_resources",
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
		if inv.IsParameterSelected("Offset") {
			ret["cluster_resource[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["cluster_resource[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionClusterResourceIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

