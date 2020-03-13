package client

import (
	"strings"
)

// ActionUserNamespaceMapUpdate is a type for action User_namespace_map#Update
type ActionUserNamespaceMapUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserNamespaceMapUpdate(client *Client) *ActionUserNamespaceMapUpdate {
	return &ActionUserNamespaceMapUpdate{
		Client: client,
	}
}

// ActionUserNamespaceMapUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionUserNamespaceMapUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserNamespaceMapUpdateMetaGlobalInput) SetNo(value bool) *ActionUserNamespaceMapUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserNamespaceMapUpdateMetaGlobalInput) SetIncludes(value string) *ActionUserNamespaceMapUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNamespaceMapUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceMapUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionUserNamespaceMapUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNamespaceMapUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNamespaceMapUpdateInput is a type for action input parameters
type ActionUserNamespaceMapUpdateInput struct {
	UserNamespace int64 `json:"user_namespace"`
	Label string `json:"label"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetUserNamespace sets parameter UserNamespace to value and selects it for sending
func (in *ActionUserNamespaceMapUpdateInput) SetUserNamespace(value int64) *ActionUserNamespaceMapUpdateInput {
	in.UserNamespace = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserNamespace"] = nil
	return in
}
// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionUserNamespaceMapUpdateInput) SetLabel(value string) *ActionUserNamespaceMapUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNamespaceMapUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceMapUpdateInput) SelectParameters(params ...string) *ActionUserNamespaceMapUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNamespaceMapUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNamespaceMapUpdateRequest is a type for the entire action request
type ActionUserNamespaceMapUpdateRequest struct {
	UserNamespaceMap map[string]interface{} `json:"user_namespace_map"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionUserNamespaceMapUpdateOutput is a type for action output parameters
type ActionUserNamespaceMapUpdateOutput struct {
	Id int64 `json:"id"`
	UserNamespace *ActionUserNamespaceShowOutput `json:"user_namespace"`
	Label string `json:"label"`
}


// Type for action response, including envelope
type ActionUserNamespaceMapUpdateResponse struct {
	Action *ActionUserNamespaceMapUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserNamespaceMap *ActionUserNamespaceMapUpdateOutput `json:"user_namespace_map"`
	}

	// Action output without the namespace
	Output *ActionUserNamespaceMapUpdateOutput
}


// Prepare the action for invocation
func (action *ActionUserNamespaceMapUpdate) Prepare() *ActionUserNamespaceMapUpdateInvocation {
	return &ActionUserNamespaceMapUpdateInvocation{
		Action: action,
		Path: "/v6.0/user_namespace_maps/{user_namespace_map_id}",
	}
}

// ActionUserNamespaceMapUpdateInvocation is used to configure action for invocation
type ActionUserNamespaceMapUpdateInvocation struct {
	// Pointer to the action
	Action *ActionUserNamespaceMapUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserNamespaceMapUpdateInput
	// Global meta input parameters
	MetaInput *ActionUserNamespaceMapUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserNamespaceMapUpdateInvocation) SetPathParamInt(param string, value int64) *ActionUserNamespaceMapUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserNamespaceMapUpdateInvocation) SetPathParamString(param string, value string) *ActionUserNamespaceMapUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserNamespaceMapUpdateInvocation) NewInput() *ActionUserNamespaceMapUpdateInput {
	inv.Input = &ActionUserNamespaceMapUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserNamespaceMapUpdateInvocation) SetInput(input *ActionUserNamespaceMapUpdateInput) *ActionUserNamespaceMapUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserNamespaceMapUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserNamespaceMapUpdateInvocation) NewMetaInput() *ActionUserNamespaceMapUpdateMetaGlobalInput {
	inv.MetaInput = &ActionUserNamespaceMapUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserNamespaceMapUpdateInvocation) SetMetaInput(input *ActionUserNamespaceMapUpdateMetaGlobalInput) *ActionUserNamespaceMapUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserNamespaceMapUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserNamespaceMapUpdateInvocation) Call() (*ActionUserNamespaceMapUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionUserNamespaceMapUpdateInvocation) callAsBody() (*ActionUserNamespaceMapUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserNamespaceMapUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserNamespaceMap
	}
	return resp, err
}




func (inv *ActionUserNamespaceMapUpdateInvocation) makeAllInputParams() *ActionUserNamespaceMapUpdateRequest {
	return &ActionUserNamespaceMapUpdateRequest{
		UserNamespaceMap: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserNamespaceMapUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("UserNamespace") {
			ret["user_namespace"] = inv.Input.UserNamespace
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
	}

	return ret
}

func (inv *ActionUserNamespaceMapUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
