package client

import (
	"strings"
)

// ActionOsFamilyDelete is a type for action Os_family#Delete
type ActionOsFamilyDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionOsFamilyDelete(client *Client) *ActionOsFamilyDelete {
	return &ActionOsFamilyDelete{
		Client: client,
	}
}

// ActionOsFamilyDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionOsFamilyDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOsFamilyDeleteMetaGlobalInput) SetIncludes(value string) *ActionOsFamilyDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOsFamilyDeleteMetaGlobalInput) SetNo(value bool) *ActionOsFamilyDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOsFamilyDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOsFamilyDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionOsFamilyDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOsFamilyDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOsFamilyDeleteRequest is a type for the entire action request
type ActionOsFamilyDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionOsFamilyDeleteResponse struct {
	Action *ActionOsFamilyDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionOsFamilyDelete) Prepare() *ActionOsFamilyDeleteInvocation {
	return &ActionOsFamilyDeleteInvocation{
		Action: action,
		Path:   "/v7.0/os_families/{os_family_id}",
	}
}

// ActionOsFamilyDeleteInvocation is used to configure action for invocation
type ActionOsFamilyDeleteInvocation struct {
	// Pointer to the action
	Action *ActionOsFamilyDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionOsFamilyDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOsFamilyDeleteInvocation) SetPathParamInt(param string, value int64) *ActionOsFamilyDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOsFamilyDeleteInvocation) SetPathParamString(param string, value string) *ActionOsFamilyDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOsFamilyDeleteInvocation) NewMetaInput() *ActionOsFamilyDeleteMetaGlobalInput {
	inv.MetaInput = &ActionOsFamilyDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOsFamilyDeleteInvocation) SetMetaInput(input *ActionOsFamilyDeleteMetaGlobalInput) *ActionOsFamilyDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOsFamilyDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOsFamilyDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOsFamilyDeleteInvocation) Call() (*ActionOsFamilyDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionOsFamilyDeleteInvocation) callAsBody() (*ActionOsFamilyDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOsFamilyDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionOsFamilyDeleteInvocation) makeAllInputParams() *ActionOsFamilyDeleteRequest {
	return &ActionOsFamilyDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionOsFamilyDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
