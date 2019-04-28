package client

import (
	"strings"
)

// ActionUserNamespaceMapEntryIndex is a type for action User_namespace_map.Entry#Index
type ActionUserNamespaceMapEntryIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserNamespaceMapEntryIndex(client *Client) *ActionUserNamespaceMapEntryIndex {
	return &ActionUserNamespaceMapEntryIndex{
		Client: client,
	}
}

// ActionUserNamespaceMapEntryIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserNamespaceMapEntryIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserNamespaceMapEntryIndexMetaGlobalInput) SetNo(value bool) *ActionUserNamespaceMapEntryIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserNamespaceMapEntryIndexMetaGlobalInput) SetCount(value bool) *ActionUserNamespaceMapEntryIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserNamespaceMapEntryIndexMetaGlobalInput) SetIncludes(value string) *ActionUserNamespaceMapEntryIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNamespaceMapEntryIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceMapEntryIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserNamespaceMapEntryIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNamespaceMapEntryIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNamespaceMapEntryIndexInput is a type for action input parameters
type ActionUserNamespaceMapEntryIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionUserNamespaceMapEntryIndexInput) SetOffset(value int64) *ActionUserNamespaceMapEntryIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserNamespaceMapEntryIndexInput) SetLimit(value int64) *ActionUserNamespaceMapEntryIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNamespaceMapEntryIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceMapEntryIndexInput) SelectParameters(params ...string) *ActionUserNamespaceMapEntryIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNamespaceMapEntryIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionUserNamespaceMapEntryIndexOutput is a type for action output parameters
type ActionUserNamespaceMapEntryIndexOutput struct {
	Id int64 `json:"id"`
	Kind string `json:"kind"`
	VpsId int64 `json:"vps_id"`
	NsId int64 `json:"ns_id"`
	Count int64 `json:"count"`
}


// Type for action response, including envelope
type ActionUserNamespaceMapEntryIndexResponse struct {
	Action *ActionUserNamespaceMapEntryIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Entries []*ActionUserNamespaceMapEntryIndexOutput `json:"entries"`
	}

	// Action output without the namespace
	Output []*ActionUserNamespaceMapEntryIndexOutput
}


// Prepare the action for invocation
func (action *ActionUserNamespaceMapEntryIndex) Prepare() *ActionUserNamespaceMapEntryIndexInvocation {
	return &ActionUserNamespaceMapEntryIndexInvocation{
		Action: action,
		Path: "/v5.0/user_namespace_maps/:user_namespace_map_id/entries",
	}
}

// ActionUserNamespaceMapEntryIndexInvocation is used to configure action for invocation
type ActionUserNamespaceMapEntryIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserNamespaceMapEntryIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserNamespaceMapEntryIndexInput
	// Global meta input parameters
	MetaInput *ActionUserNamespaceMapEntryIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserNamespaceMapEntryIndexInvocation) SetPathParamInt(param string, value int64) *ActionUserNamespaceMapEntryIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserNamespaceMapEntryIndexInvocation) SetPathParamString(param string, value string) *ActionUserNamespaceMapEntryIndexInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserNamespaceMapEntryIndexInvocation) NewInput() *ActionUserNamespaceMapEntryIndexInput {
	inv.Input = &ActionUserNamespaceMapEntryIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserNamespaceMapEntryIndexInvocation) SetInput(input *ActionUserNamespaceMapEntryIndexInput) *ActionUserNamespaceMapEntryIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserNamespaceMapEntryIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserNamespaceMapEntryIndexInvocation) NewMetaInput() *ActionUserNamespaceMapEntryIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserNamespaceMapEntryIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserNamespaceMapEntryIndexInvocation) SetMetaInput(input *ActionUserNamespaceMapEntryIndexMetaGlobalInput) *ActionUserNamespaceMapEntryIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserNamespaceMapEntryIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserNamespaceMapEntryIndexInvocation) Call() (*ActionUserNamespaceMapEntryIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserNamespaceMapEntryIndexInvocation) callAsQuery() (*ActionUserNamespaceMapEntryIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserNamespaceMapEntryIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Entries
	}
	return resp, err
}



func (inv *ActionUserNamespaceMapEntryIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["entry[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["entry[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionUserNamespaceMapEntryIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

