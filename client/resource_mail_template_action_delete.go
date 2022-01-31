package client

import (
	"strings"
)

// ActionMailTemplateDelete is a type for action Mail_template#Delete
type ActionMailTemplateDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionMailTemplateDelete(client *Client) *ActionMailTemplateDelete {
	return &ActionMailTemplateDelete{
		Client: client,
	}
}

// ActionMailTemplateDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionMailTemplateDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailTemplateDeleteMetaGlobalInput) SetIncludes(value string) *ActionMailTemplateDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailTemplateDeleteMetaGlobalInput) SetNo(value bool) *ActionMailTemplateDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionMailTemplateDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailTemplateDeleteRequest is a type for the entire action request
type ActionMailTemplateDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionMailTemplateDeleteResponse struct {
	Action *ActionMailTemplateDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionMailTemplateDelete) Prepare() *ActionMailTemplateDeleteInvocation {
	return &ActionMailTemplateDeleteInvocation{
		Action: action,
		Path:   "/v6.0/mail_templates/{mail_template_id}",
	}
}

// ActionMailTemplateDeleteInvocation is used to configure action for invocation
type ActionMailTemplateDeleteInvocation struct {
	// Pointer to the action
	Action *ActionMailTemplateDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMailTemplateDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailTemplateDeleteInvocation) SetPathParamInt(param string, value int64) *ActionMailTemplateDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailTemplateDeleteInvocation) SetPathParamString(param string, value string) *ActionMailTemplateDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailTemplateDeleteInvocation) NewMetaInput() *ActionMailTemplateDeleteMetaGlobalInput {
	inv.MetaInput = &ActionMailTemplateDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailTemplateDeleteInvocation) SetMetaInput(input *ActionMailTemplateDeleteMetaGlobalInput) *ActionMailTemplateDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailTemplateDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMailTemplateDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailTemplateDeleteInvocation) Call() (*ActionMailTemplateDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionMailTemplateDeleteInvocation) callAsBody() (*ActionMailTemplateDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMailTemplateDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionMailTemplateDeleteInvocation) makeAllInputParams() *ActionMailTemplateDeleteRequest {
	return &ActionMailTemplateDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionMailTemplateDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
