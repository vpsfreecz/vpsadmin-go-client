package client

import (
	"strings"
)

// ActionOsTemplateDelete is a type for action Os_template#Delete
type ActionOsTemplateDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionOsTemplateDelete(client *Client) *ActionOsTemplateDelete {
	return &ActionOsTemplateDelete{
		Client: client,
	}
}

// ActionOsTemplateDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionOsTemplateDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOsTemplateDeleteMetaGlobalInput) SetIncludes(value string) *ActionOsTemplateDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOsTemplateDeleteMetaGlobalInput) SetNo(value bool) *ActionOsTemplateDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOsTemplateDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOsTemplateDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionOsTemplateDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOsTemplateDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOsTemplateDeleteRequest is a type for the entire action request
type ActionOsTemplateDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionOsTemplateDeleteResponse struct {
	Action *ActionOsTemplateDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionOsTemplateDelete) Prepare() *ActionOsTemplateDeleteInvocation {
	return &ActionOsTemplateDeleteInvocation{
		Action: action,
		Path:   "/v7.0/os_templates/{os_template_id}",
	}
}

// ActionOsTemplateDeleteInvocation is used to configure action for invocation
type ActionOsTemplateDeleteInvocation struct {
	// Pointer to the action
	Action *ActionOsTemplateDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionOsTemplateDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOsTemplateDeleteInvocation) SetPathParamInt(param string, value int64) *ActionOsTemplateDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOsTemplateDeleteInvocation) SetPathParamString(param string, value string) *ActionOsTemplateDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOsTemplateDeleteInvocation) NewMetaInput() *ActionOsTemplateDeleteMetaGlobalInput {
	inv.MetaInput = &ActionOsTemplateDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOsTemplateDeleteInvocation) SetMetaInput(input *ActionOsTemplateDeleteMetaGlobalInput) *ActionOsTemplateDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOsTemplateDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOsTemplateDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOsTemplateDeleteInvocation) Call() (*ActionOsTemplateDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionOsTemplateDeleteInvocation) callAsBody() (*ActionOsTemplateDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOsTemplateDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionOsTemplateDeleteInvocation) makeAllInputParams() *ActionOsTemplateDeleteRequest {
	return &ActionOsTemplateDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionOsTemplateDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
