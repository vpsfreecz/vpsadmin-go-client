package client

import (
	"strings"
)

// ActionClusterResourcePackageItemIndex is a type for action Cluster_resource_package.Item#Index
type ActionClusterResourcePackageItemIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterResourcePackageItemIndex(client *Client) *ActionClusterResourcePackageItemIndex {
	return &ActionClusterResourcePackageItemIndex{
		Client: client,
	}
}

// ActionClusterResourcePackageItemIndexMetaGlobalInput is a type for action global meta input parameters
type ActionClusterResourcePackageItemIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterResourcePackageItemIndexMetaGlobalInput) SetNo(value bool) *ActionClusterResourcePackageItemIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionClusterResourcePackageItemIndexMetaGlobalInput) SetCount(value bool) *ActionClusterResourcePackageItemIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionClusterResourcePackageItemIndexMetaGlobalInput) SetIncludes(value string) *ActionClusterResourcePackageItemIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourcePackageItemIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourcePackageItemIndexMetaGlobalInput) SelectParameters(params ...string) *ActionClusterResourcePackageItemIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterResourcePackageItemIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterResourcePackageItemIndexInput is a type for action input parameters
type ActionClusterResourcePackageItemIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionClusterResourcePackageItemIndexInput) SetOffset(value int64) *ActionClusterResourcePackageItemIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionClusterResourcePackageItemIndexInput) SetLimit(value int64) *ActionClusterResourcePackageItemIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourcePackageItemIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourcePackageItemIndexInput) SelectParameters(params ...string) *ActionClusterResourcePackageItemIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterResourcePackageItemIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionClusterResourcePackageItemIndexOutput is a type for action output parameters
type ActionClusterResourcePackageItemIndexOutput struct {
	Id int64 `json:"id"`
	ClusterResource *ActionClusterResourceShowOutput `json:"cluster_resource"`
	Value int64 `json:"value"`
}


// Type for action response, including envelope
type ActionClusterResourcePackageItemIndexResponse struct {
	Action *ActionClusterResourcePackageItemIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Items []*ActionClusterResourcePackageItemIndexOutput `json:"items"`
	}

	// Action output without the namespace
	Output []*ActionClusterResourcePackageItemIndexOutput
}


// Prepare the action for invocation
func (action *ActionClusterResourcePackageItemIndex) Prepare() *ActionClusterResourcePackageItemIndexInvocation {
	return &ActionClusterResourcePackageItemIndexInvocation{
		Action: action,
		Path: "/v5.0/cluster_resource_packages/{cluster_resource_package_id}/items",
	}
}

// ActionClusterResourcePackageItemIndexInvocation is used to configure action for invocation
type ActionClusterResourcePackageItemIndexInvocation struct {
	// Pointer to the action
	Action *ActionClusterResourcePackageItemIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionClusterResourcePackageItemIndexInput
	// Global meta input parameters
	MetaInput *ActionClusterResourcePackageItemIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionClusterResourcePackageItemIndexInvocation) SetPathParamInt(param string, value int64) *ActionClusterResourcePackageItemIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionClusterResourcePackageItemIndexInvocation) SetPathParamString(param string, value string) *ActionClusterResourcePackageItemIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionClusterResourcePackageItemIndexInvocation) NewInput() *ActionClusterResourcePackageItemIndexInput {
	inv.Input = &ActionClusterResourcePackageItemIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionClusterResourcePackageItemIndexInvocation) SetInput(input *ActionClusterResourcePackageItemIndexInput) *ActionClusterResourcePackageItemIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionClusterResourcePackageItemIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionClusterResourcePackageItemIndexInvocation) NewMetaInput() *ActionClusterResourcePackageItemIndexMetaGlobalInput {
	inv.MetaInput = &ActionClusterResourcePackageItemIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterResourcePackageItemIndexInvocation) SetMetaInput(input *ActionClusterResourcePackageItemIndexMetaGlobalInput) *ActionClusterResourcePackageItemIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterResourcePackageItemIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionClusterResourcePackageItemIndexInvocation) Call() (*ActionClusterResourcePackageItemIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionClusterResourcePackageItemIndexInvocation) callAsQuery() (*ActionClusterResourcePackageItemIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionClusterResourcePackageItemIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Items
	}
	return resp, err
}



func (inv *ActionClusterResourcePackageItemIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["item[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["item[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionClusterResourcePackageItemIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

