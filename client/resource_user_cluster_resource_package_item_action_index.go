package client

import (
	"strings"
)

// ActionUserClusterResourcePackageItemIndex is a type for action User_cluster_resource_package.Item#Index
type ActionUserClusterResourcePackageItemIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserClusterResourcePackageItemIndex(client *Client) *ActionUserClusterResourcePackageItemIndex {
	return &ActionUserClusterResourcePackageItemIndex{
		Client: client,
	}
}

// ActionUserClusterResourcePackageItemIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserClusterResourcePackageItemIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserClusterResourcePackageItemIndexMetaGlobalInput) SetCount(value bool) *ActionUserClusterResourcePackageItemIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserClusterResourcePackageItemIndexMetaGlobalInput) SetIncludes(value string) *ActionUserClusterResourcePackageItemIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserClusterResourcePackageItemIndexMetaGlobalInput) SetNo(value bool) *ActionUserClusterResourcePackageItemIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserClusterResourcePackageItemIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserClusterResourcePackageItemIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserClusterResourcePackageItemIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserClusterResourcePackageItemIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserClusterResourcePackageItemIndexInput is a type for action input parameters
type ActionUserClusterResourcePackageItemIndexInput struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserClusterResourcePackageItemIndexInput) SetLimit(value int64) *ActionUserClusterResourcePackageItemIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionUserClusterResourcePackageItemIndexInput) SetOffset(value int64) *ActionUserClusterResourcePackageItemIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserClusterResourcePackageItemIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserClusterResourcePackageItemIndexInput) SelectParameters(params ...string) *ActionUserClusterResourcePackageItemIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserClusterResourcePackageItemIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserClusterResourcePackageItemIndexOutput is a type for action output parameters
type ActionUserClusterResourcePackageItemIndexOutput struct {
	ClusterResource *ActionClusterResourceShowOutput `json:"cluster_resource"`
	Id              int64                            `json:"id"`
	Value           int64                            `json:"value"`
}

// Type for action response, including envelope
type ActionUserClusterResourcePackageItemIndexResponse struct {
	Action *ActionUserClusterResourcePackageItemIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Items []*ActionUserClusterResourcePackageItemIndexOutput `json:"items"`
	}

	// Action output without the namespace
	Output []*ActionUserClusterResourcePackageItemIndexOutput
}

// Prepare the action for invocation
func (action *ActionUserClusterResourcePackageItemIndex) Prepare() *ActionUserClusterResourcePackageItemIndexInvocation {
	return &ActionUserClusterResourcePackageItemIndexInvocation{
		Action: action,
		Path:   "/v6.0/user_cluster_resource_packages/{user_cluster_resource_package_id}/items",
	}
}

// ActionUserClusterResourcePackageItemIndexInvocation is used to configure action for invocation
type ActionUserClusterResourcePackageItemIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserClusterResourcePackageItemIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserClusterResourcePackageItemIndexInput
	// Global meta input parameters
	MetaInput *ActionUserClusterResourcePackageItemIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserClusterResourcePackageItemIndexInvocation) SetPathParamInt(param string, value int64) *ActionUserClusterResourcePackageItemIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserClusterResourcePackageItemIndexInvocation) SetPathParamString(param string, value string) *ActionUserClusterResourcePackageItemIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserClusterResourcePackageItemIndexInvocation) NewInput() *ActionUserClusterResourcePackageItemIndexInput {
	inv.Input = &ActionUserClusterResourcePackageItemIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserClusterResourcePackageItemIndexInvocation) SetInput(input *ActionUserClusterResourcePackageItemIndexInput) *ActionUserClusterResourcePackageItemIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserClusterResourcePackageItemIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserClusterResourcePackageItemIndexInvocation) NewMetaInput() *ActionUserClusterResourcePackageItemIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserClusterResourcePackageItemIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserClusterResourcePackageItemIndexInvocation) SetMetaInput(input *ActionUserClusterResourcePackageItemIndexMetaGlobalInput) *ActionUserClusterResourcePackageItemIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserClusterResourcePackageItemIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserClusterResourcePackageItemIndexInvocation) Call() (*ActionUserClusterResourcePackageItemIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserClusterResourcePackageItemIndexInvocation) callAsQuery() (*ActionUserClusterResourcePackageItemIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserClusterResourcePackageItemIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Items
	}
	return resp, err
}

func (inv *ActionUserClusterResourcePackageItemIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["item[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["item[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
	}
}

func (inv *ActionUserClusterResourcePackageItemIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
