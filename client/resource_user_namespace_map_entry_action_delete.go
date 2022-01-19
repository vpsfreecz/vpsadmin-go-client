package client

import (
	"strings"
)

// ActionUserNamespaceMapEntryDelete is a type for action User_namespace_map.Entry#Delete
type ActionUserNamespaceMapEntryDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionUserNamespaceMapEntryDelete(client *Client) *ActionUserNamespaceMapEntryDelete {
	return &ActionUserNamespaceMapEntryDelete{
		Client: client,
	}
}

// ActionUserNamespaceMapEntryDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionUserNamespaceMapEntryDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserNamespaceMapEntryDeleteMetaGlobalInput) SetIncludes(value string) *ActionUserNamespaceMapEntryDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserNamespaceMapEntryDeleteMetaGlobalInput) SetNo(value bool) *ActionUserNamespaceMapEntryDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNamespaceMapEntryDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceMapEntryDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionUserNamespaceMapEntryDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNamespaceMapEntryDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNamespaceMapEntryDeleteRequest is a type for the entire action request
type ActionUserNamespaceMapEntryDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionUserNamespaceMapEntryDeleteResponse struct {
	Action *ActionUserNamespaceMapEntryDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionUserNamespaceMapEntryDelete) Prepare() *ActionUserNamespaceMapEntryDeleteInvocation {
	return &ActionUserNamespaceMapEntryDeleteInvocation{
		Action: action,
		Path:   "/v6.0/user_namespace_maps/{user_namespace_map_id}/entries/{entry_id}",
	}
}

// ActionUserNamespaceMapEntryDeleteInvocation is used to configure action for invocation
type ActionUserNamespaceMapEntryDeleteInvocation struct {
	// Pointer to the action
	Action *ActionUserNamespaceMapEntryDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserNamespaceMapEntryDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserNamespaceMapEntryDeleteInvocation) SetPathParamInt(param string, value int64) *ActionUserNamespaceMapEntryDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserNamespaceMapEntryDeleteInvocation) SetPathParamString(param string, value string) *ActionUserNamespaceMapEntryDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserNamespaceMapEntryDeleteInvocation) NewMetaInput() *ActionUserNamespaceMapEntryDeleteMetaGlobalInput {
	inv.MetaInput = &ActionUserNamespaceMapEntryDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserNamespaceMapEntryDeleteInvocation) SetMetaInput(input *ActionUserNamespaceMapEntryDeleteMetaGlobalInput) *ActionUserNamespaceMapEntryDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserNamespaceMapEntryDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserNamespaceMapEntryDeleteInvocation) Call() (*ActionUserNamespaceMapEntryDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionUserNamespaceMapEntryDeleteInvocation) callAsBody() (*ActionUserNamespaceMapEntryDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserNamespaceMapEntryDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionUserNamespaceMapEntryDeleteInvocation) makeAllInputParams() *ActionUserNamespaceMapEntryDeleteRequest {
	return &ActionUserNamespaceMapEntryDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserNamespaceMapEntryDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
