package client

import (
)

// ActionMailTemplateIndex is a type for action Mail_template#Index
type ActionMailTemplateIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionMailTemplateIndex(client *Client) *ActionMailTemplateIndex {
	return &ActionMailTemplateIndex{
		Client: client,
	}
}

// ActionMailTemplateIndexMetaGlobalInput is a type for action global meta input parameters
type ActionMailTemplateIndexMetaGlobalInput struct {
	Count bool `json:"count"`
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionMailTemplateIndexMetaGlobalInput) SetCount(value bool) *ActionMailTemplateIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailTemplateIndexMetaGlobalInput) SetIncludes(value string) *ActionMailTemplateIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailTemplateIndexMetaGlobalInput) SetNo(value bool) *ActionMailTemplateIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateIndexMetaGlobalInput) SelectParameters(params ...string) *ActionMailTemplateIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailTemplateIndexInput is a type for action input parameters
type ActionMailTemplateIndexInput struct {
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionMailTemplateIndexInput) SetLimit(value int64) *ActionMailTemplateIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionMailTemplateIndexInput) SetOffset(value int64) *ActionMailTemplateIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateIndexInput) SelectParameters(params ...string) *ActionMailTemplateIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionMailTemplateIndexOutput is a type for action output parameters
type ActionMailTemplateIndexOutput struct {
	CreatedAt string `json:"created_at"`
	Id int64 `json:"id"`
	Label string `json:"label"`
	Name string `json:"name"`
	TemplateId string `json:"template_id"`
	UpdatedAt string `json:"updated_at"`
	UserVisibility string `json:"user_visibility"`
}


// Type for action response, including envelope
type ActionMailTemplateIndexResponse struct {
	Action *ActionMailTemplateIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MailTemplates []*ActionMailTemplateIndexOutput `json:"mail_templates"`
	}

	// Action output without the namespace
	Output []*ActionMailTemplateIndexOutput
}


// Prepare the action for invocation
func (action *ActionMailTemplateIndex) Prepare() *ActionMailTemplateIndexInvocation {
	return &ActionMailTemplateIndexInvocation{
		Action: action,
		Path: "/v6.0/mail_templates",
	}
}

// ActionMailTemplateIndexInvocation is used to configure action for invocation
type ActionMailTemplateIndexInvocation struct {
	// Pointer to the action
	Action *ActionMailTemplateIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMailTemplateIndexInput
	// Global meta input parameters
	MetaInput *ActionMailTemplateIndexMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMailTemplateIndexInvocation) NewInput() *ActionMailTemplateIndexInput {
	inv.Input = &ActionMailTemplateIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMailTemplateIndexInvocation) SetInput(input *ActionMailTemplateIndexInput) *ActionMailTemplateIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMailTemplateIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailTemplateIndexInvocation) NewMetaInput() *ActionMailTemplateIndexMetaGlobalInput {
	inv.MetaInput = &ActionMailTemplateIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailTemplateIndexInvocation) SetMetaInput(input *ActionMailTemplateIndexMetaGlobalInput) *ActionMailTemplateIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailTemplateIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailTemplateIndexInvocation) Call() (*ActionMailTemplateIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMailTemplateIndexInvocation) callAsQuery() (*ActionMailTemplateIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMailTemplateIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MailTemplates
	}
	return resp, err
}



func (inv *ActionMailTemplateIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["mail_template[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["mail_template[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
	}
}

func (inv *ActionMailTemplateIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}

