package client

import (
	"strings"
)

// ActionUserNamespaceMapEntryCreate is a type for action User_namespace_map.Entry#Create
type ActionUserNamespaceMapEntryCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserNamespaceMapEntryCreate(client *Client) *ActionUserNamespaceMapEntryCreate {
	return &ActionUserNamespaceMapEntryCreate{
		Client: client,
	}
}

// ActionUserNamespaceMapEntryCreateMetaGlobalInput is a type for action global meta input parameters
type ActionUserNamespaceMapEntryCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserNamespaceMapEntryCreateMetaGlobalInput) SetNo(value bool) *ActionUserNamespaceMapEntryCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserNamespaceMapEntryCreateMetaGlobalInput) SetIncludes(value string) *ActionUserNamespaceMapEntryCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNamespaceMapEntryCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceMapEntryCreateMetaGlobalInput) SelectParameters(params ...string) *ActionUserNamespaceMapEntryCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNamespaceMapEntryCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNamespaceMapEntryCreateInput is a type for action input parameters
type ActionUserNamespaceMapEntryCreateInput struct {
	Kind string `json:"kind"`
	VpsId int64 `json:"vps_id"`
	NsId int64 `json:"ns_id"`
	Count int64 `json:"count"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetKind sets parameter Kind to value and selects it for sending
func (in *ActionUserNamespaceMapEntryCreateInput) SetKind(value string) *ActionUserNamespaceMapEntryCreateInput {
	in.Kind = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Kind"] = nil
	return in
}
// SetVpsId sets parameter VpsId to value and selects it for sending
func (in *ActionUserNamespaceMapEntryCreateInput) SetVpsId(value int64) *ActionUserNamespaceMapEntryCreateInput {
	in.VpsId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["VpsId"] = nil
	return in
}
// SetNsId sets parameter NsId to value and selects it for sending
func (in *ActionUserNamespaceMapEntryCreateInput) SetNsId(value int64) *ActionUserNamespaceMapEntryCreateInput {
	in.NsId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NsId"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserNamespaceMapEntryCreateInput) SetCount(value int64) *ActionUserNamespaceMapEntryCreateInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNamespaceMapEntryCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceMapEntryCreateInput) SelectParameters(params ...string) *ActionUserNamespaceMapEntryCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNamespaceMapEntryCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNamespaceMapEntryCreateRequest is a type for the entire action request
type ActionUserNamespaceMapEntryCreateRequest struct {
	Entry map[string]interface{} `json:"entry"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionUserNamespaceMapEntryCreateOutput is a type for action output parameters
type ActionUserNamespaceMapEntryCreateOutput struct {
	Id int64 `json:"id"`
	Kind string `json:"kind"`
	VpsId int64 `json:"vps_id"`
	NsId int64 `json:"ns_id"`
	Count int64 `json:"count"`
}


// Type for action response, including envelope
type ActionUserNamespaceMapEntryCreateResponse struct {
	Action *ActionUserNamespaceMapEntryCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Entry *ActionUserNamespaceMapEntryCreateOutput `json:"entry"`
	}

	// Action output without the namespace
	Output *ActionUserNamespaceMapEntryCreateOutput
}


// Prepare the action for invocation
func (action *ActionUserNamespaceMapEntryCreate) Prepare() *ActionUserNamespaceMapEntryCreateInvocation {
	return &ActionUserNamespaceMapEntryCreateInvocation{
		Action: action,
		Path: "/v5.0/user_namespace_maps/{user_namespace_map_id}/entries",
	}
}

// ActionUserNamespaceMapEntryCreateInvocation is used to configure action for invocation
type ActionUserNamespaceMapEntryCreateInvocation struct {
	// Pointer to the action
	Action *ActionUserNamespaceMapEntryCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserNamespaceMapEntryCreateInput
	// Global meta input parameters
	MetaInput *ActionUserNamespaceMapEntryCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserNamespaceMapEntryCreateInvocation) SetPathParamInt(param string, value int64) *ActionUserNamespaceMapEntryCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserNamespaceMapEntryCreateInvocation) SetPathParamString(param string, value string) *ActionUserNamespaceMapEntryCreateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserNamespaceMapEntryCreateInvocation) NewInput() *ActionUserNamespaceMapEntryCreateInput {
	inv.Input = &ActionUserNamespaceMapEntryCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserNamespaceMapEntryCreateInvocation) SetInput(input *ActionUserNamespaceMapEntryCreateInput) *ActionUserNamespaceMapEntryCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserNamespaceMapEntryCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserNamespaceMapEntryCreateInvocation) NewMetaInput() *ActionUserNamespaceMapEntryCreateMetaGlobalInput {
	inv.MetaInput = &ActionUserNamespaceMapEntryCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserNamespaceMapEntryCreateInvocation) SetMetaInput(input *ActionUserNamespaceMapEntryCreateMetaGlobalInput) *ActionUserNamespaceMapEntryCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserNamespaceMapEntryCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserNamespaceMapEntryCreateInvocation) Call() (*ActionUserNamespaceMapEntryCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionUserNamespaceMapEntryCreateInvocation) callAsBody() (*ActionUserNamespaceMapEntryCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserNamespaceMapEntryCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Entry
	}
	return resp, err
}




func (inv *ActionUserNamespaceMapEntryCreateInvocation) makeAllInputParams() *ActionUserNamespaceMapEntryCreateRequest {
	return &ActionUserNamespaceMapEntryCreateRequest{
		Entry: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserNamespaceMapEntryCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Kind") {
			ret["kind"] = inv.Input.Kind
		}
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

func (inv *ActionUserNamespaceMapEntryCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
