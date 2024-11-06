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
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
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

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserNamespaceMapUpdateMetaGlobalInput) SetNo(value bool) *ActionUserNamespaceMapUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
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
	Label         string `json:"label"`
	UserNamespace int64  `json:"user_namespace"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
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

// SetUserNamespace sets parameter UserNamespace to value and selects it for sending
func (in *ActionUserNamespaceMapUpdateInput) SetUserNamespace(value int64) *ActionUserNamespaceMapUpdateInput {
	in.UserNamespace = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNamespaceNil(false)
	in._selectedParameters["UserNamespace"] = nil
	return in
}

// SetUserNamespaceNil sets parameter UserNamespace to nil and selects it for sending
func (in *ActionUserNamespaceMapUpdateInput) SetUserNamespaceNil(set bool) *ActionUserNamespaceMapUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["UserNamespace"] = nil
		in.SelectParameters("UserNamespace")
	} else {
		delete(in._nilParameters, "UserNamespace")
	}
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

// UnselectParameters unsets parameters from ActionUserNamespaceMapUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserNamespaceMapUpdateInput) UnselectParameters(params ...string) *ActionUserNamespaceMapUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
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
	Meta             map[string]interface{} `json:"_meta"`
}

// ActionUserNamespaceMapUpdateOutput is a type for action output parameters
type ActionUserNamespaceMapUpdateOutput struct {
	Id            int64                          `json:"id"`
	Label         string                         `json:"label"`
	UserNamespace *ActionUserNamespaceShowOutput `json:"user_namespace"`
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
		Path:   "/v7.0/user_namespace_maps/{user_namespace_map_id}",
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

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserNamespaceMapUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
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

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserNamespaceMapUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
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
		Meta:             inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserNamespaceMapUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("UserNamespace") {
			if inv.IsParameterNil("UserNamespace") {
				ret["user_namespace"] = nil
			} else {
				ret["user_namespace"] = inv.Input.UserNamespace
			}
		}
	}

	return ret
}

func (inv *ActionUserNamespaceMapUpdateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
