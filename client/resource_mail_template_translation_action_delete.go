package client

import (
	"strings"
)

// ActionMailTemplateTranslationDelete is a type for action Mail_template.Translation#Delete
type ActionMailTemplateTranslationDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionMailTemplateTranslationDelete(client *Client) *ActionMailTemplateTranslationDelete {
	return &ActionMailTemplateTranslationDelete{
		Client: client,
	}
}

// ActionMailTemplateTranslationDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionMailTemplateTranslationDeleteMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailTemplateTranslationDeleteMetaGlobalInput) SetNo(value bool) *ActionMailTemplateTranslationDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailTemplateTranslationDeleteMetaGlobalInput) SetIncludes(value string) *ActionMailTemplateTranslationDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateTranslationDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateTranslationDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionMailTemplateTranslationDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateTranslationDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionMailTemplateTranslationDeleteRequest is a type for the entire action request
type ActionMailTemplateTranslationDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionMailTemplateTranslationDeleteResponse struct {
	Action *ActionMailTemplateTranslationDelete `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionMailTemplateTranslationDelete) Prepare() *ActionMailTemplateTranslationDeleteInvocation {
	return &ActionMailTemplateTranslationDeleteInvocation{
		Action: action,
		Path: "/v5.0/mail_templates/{mail_template_id}/translations/{translation_id}",
	}
}

// ActionMailTemplateTranslationDeleteInvocation is used to configure action for invocation
type ActionMailTemplateTranslationDeleteInvocation struct {
	// Pointer to the action
	Action *ActionMailTemplateTranslationDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMailTemplateTranslationDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailTemplateTranslationDeleteInvocation) SetPathParamInt(param string, value int64) *ActionMailTemplateTranslationDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailTemplateTranslationDeleteInvocation) SetPathParamString(param string, value string) *ActionMailTemplateTranslationDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailTemplateTranslationDeleteInvocation) NewMetaInput() *ActionMailTemplateTranslationDeleteMetaGlobalInput {
	inv.MetaInput = &ActionMailTemplateTranslationDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailTemplateTranslationDeleteInvocation) SetMetaInput(input *ActionMailTemplateTranslationDeleteMetaGlobalInput) *ActionMailTemplateTranslationDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailTemplateTranslationDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailTemplateTranslationDeleteInvocation) Call() (*ActionMailTemplateTranslationDeleteResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionMailTemplateTranslationDeleteInvocation) callAsBody() (*ActionMailTemplateTranslationDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMailTemplateTranslationDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionMailTemplateTranslationDeleteInvocation) makeAllInputParams() *ActionMailTemplateTranslationDeleteRequest {
	return &ActionMailTemplateTranslationDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionMailTemplateTranslationDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
