package client

import (
	"strings"
)

// ActionUserNamespaceMapEntryUpdate is a type for action User_namespace_map.Entry#Update
type ActionUserNamespaceMapEntryUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserNamespaceMapEntryUpdate(client *Client) *ActionUserNamespaceMapEntryUpdate {
	return &ActionUserNamespaceMapEntryUpdate{
		Client: client,
	}
}

// ActionUserNamespaceMapEntryUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionUserNamespaceMapEntryUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserNamespaceMapEntryUpdateMetaGlobalInput) SetNo(value bool) *ActionUserNamespaceMapEntryUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserNamespaceMapEntryUpdateMetaGlobalInput) SetIncludes(value string) *ActionUserNamespaceMapEntryUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNamespaceMapEntryUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceMapEntryUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionUserNamespaceMapEntryUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNamespaceMapEntryUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNamespaceMapEntryUpdateInput is a type for action input parameters
type ActionUserNamespaceMapEntryUpdateInput struct {
	VpsId int64 `json:"vps_id"`
	NsId int64 `json:"ns_id"`
	Count int64 `json:"count"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetVpsId sets parameter VpsId to value and selects it for sending
func (in *ActionUserNamespaceMapEntryUpdateInput) SetVpsId(value int64) *ActionUserNamespaceMapEntryUpdateInput {
	in.VpsId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["VpsId"] = nil
	return in
}
// SetNsId sets parameter NsId to value and selects it for sending
func (in *ActionUserNamespaceMapEntryUpdateInput) SetNsId(value int64) *ActionUserNamespaceMapEntryUpdateInput {
	in.NsId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NsId"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserNamespaceMapEntryUpdateInput) SetCount(value int64) *ActionUserNamespaceMapEntryUpdateInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNamespaceMapEntryUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceMapEntryUpdateInput) SelectParameters(params ...string) *ActionUserNamespaceMapEntryUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNamespaceMapEntryUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNamespaceMapEntryUpdateRequest is a type for the entire action request
type ActionUserNamespaceMapEntryUpdateRequest struct {
	Entry map[string]interface{} `json:"entry"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionUserNamespaceMapEntryUpdateOutput is a type for action output parameters
type ActionUserNamespaceMapEntryUpdateOutput struct {
	Id int64 `json:"id"`
	Kind string `json:"kind"`
	VpsId int64 `json:"vps_id"`
	NsId int64 `json:"ns_id"`
	Count int64 `json:"count"`
}


// Type for action response, including envelope
type ActionUserNamespaceMapEntryUpdateResponse struct {
	Action *ActionUserNamespaceMapEntryUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Entry *ActionUserNamespaceMapEntryUpdateOutput `json:"entry"`
	}

	// Action output without the namespace
	Output *ActionUserNamespaceMapEntryUpdateOutput
}


// Prepare the action for invocation
func (action *ActionUserNamespaceMapEntryUpdate) Prepare() *ActionUserNamespaceMapEntryUpdateInvocation {
	return &ActionUserNamespaceMapEntryUpdateInvocation{
		Action: action,
		Path: "/v6.0/user_namespace_maps/{user_namespace_map_id}/entries/{entry_id}",
	}
}

// ActionUserNamespaceMapEntryUpdateInvocation is used to configure action for invocation
type ActionUserNamespaceMapEntryUpdateInvocation struct {
	// Pointer to the action
	Action *ActionUserNamespaceMapEntryUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserNamespaceMapEntryUpdateInput
	// Global meta input parameters
	MetaInput *ActionUserNamespaceMapEntryUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) SetPathParamInt(param string, value int64) *ActionUserNamespaceMapEntryUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) SetPathParamString(param string, value string) *ActionUserNamespaceMapEntryUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) NewInput() *ActionUserNamespaceMapEntryUpdateInput {
	inv.Input = &ActionUserNamespaceMapEntryUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) SetInput(input *ActionUserNamespaceMapEntryUpdateInput) *ActionUserNamespaceMapEntryUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) NewMetaInput() *ActionUserNamespaceMapEntryUpdateMetaGlobalInput {
	inv.MetaInput = &ActionUserNamespaceMapEntryUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) SetMetaInput(input *ActionUserNamespaceMapEntryUpdateMetaGlobalInput) *ActionUserNamespaceMapEntryUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) Call() (*ActionUserNamespaceMapEntryUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionUserNamespaceMapEntryUpdateInvocation) callAsBody() (*ActionUserNamespaceMapEntryUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserNamespaceMapEntryUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Entry
	}
	return resp, err
}




func (inv *ActionUserNamespaceMapEntryUpdateInvocation) makeAllInputParams() *ActionUserNamespaceMapEntryUpdateRequest {
	return &ActionUserNamespaceMapEntryUpdateRequest{
		Entry: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserNamespaceMapEntryUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("VpsId") {
			ret["vps_id"] = inv.Input.VpsId
		}
		if inv.IsParameterSelected("NsId") {
			ret["ns_id"] = inv.Input.NsId
		}
		if inv.IsParameterSelected("Count") {
			ret["count"] = inv.Input.Count
		}
	}

	return ret
}

func (inv *ActionUserNamespaceMapEntryUpdateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
