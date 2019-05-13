package client

import (
	"strings"
)

// ActionMailTemplateShow is a type for action Mail_template#Show
type ActionMailTemplateShow struct {
	// Pointer to client
	Client *Client
}

func NewActionMailTemplateShow(client *Client) *ActionMailTemplateShow {
	return &ActionMailTemplateShow{
		Client: client,
	}
}

// ActionMailTemplateShowMetaGlobalInput is a type for action global meta input parameters
type ActionMailTemplateShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailTemplateShowMetaGlobalInput) SetNo(value bool) *ActionMailTemplateShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailTemplateShowMetaGlobalInput) SetIncludes(value string) *ActionMailTemplateShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateShowMetaGlobalInput) SelectParameters(params ...string) *ActionMailTemplateShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionMailTemplateShowOutput is a type for action output parameters
type ActionMailTemplateShowOutput struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Label string `json:"label"`
	TemplateId string `json:"template_id"`
	UserVisibility string `json:"user_visibility"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


// Type for action response, including envelope
type ActionMailTemplateShowResponse struct {
	Action *ActionMailTemplateShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MailTemplate *ActionMailTemplateShowOutput `json:"mail_template"`
	}

	// Action output without the namespace
	Output *ActionMailTemplateShowOutput
}


// Prepare the action for invocation
func (action *ActionMailTemplateShow) Prepare() *ActionMailTemplateShowInvocation {
	return &ActionMailTemplateShowInvocation{
		Action: action,
		Path: "/v5.0/mail_templates/{mail_template_id}",
	}
}

// ActionMailTemplateShowInvocation is used to configure action for invocation
type ActionMailTemplateShowInvocation struct {
	// Pointer to the action
	Action *ActionMailTemplateShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMailTemplateShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailTemplateShowInvocation) SetPathParamInt(param string, value int64) *ActionMailTemplateShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailTemplateShowInvocation) SetPathParamString(param string, value string) *ActionMailTemplateShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailTemplateShowInvocation) NewMetaInput() *ActionMailTemplateShowMetaGlobalInput {
	inv.MetaInput = &ActionMailTemplateShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailTemplateShowInvocation) SetMetaInput(input *ActionMailTemplateShowMetaGlobalInput) *ActionMailTemplateShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailTemplateShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailTemplateShowInvocation) Call() (*ActionMailTemplateShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMailTemplateShowInvocation) callAsQuery() (*ActionMailTemplateShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMailTemplateShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MailTemplate
	}
	return resp, err
}




func (inv *ActionMailTemplateShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

