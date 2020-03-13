package client

import (
)

// ActionHelpBoxCreate is a type for action Help_box#Create
type ActionHelpBoxCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionHelpBoxCreate(client *Client) *ActionHelpBoxCreate {
	return &ActionHelpBoxCreate{
		Client: client,
	}
}

// ActionHelpBoxCreateMetaGlobalInput is a type for action global meta input parameters
type ActionHelpBoxCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionHelpBoxCreateMetaGlobalInput) SetNo(value bool) *ActionHelpBoxCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionHelpBoxCreateMetaGlobalInput) SetIncludes(value string) *ActionHelpBoxCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionHelpBoxCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionHelpBoxCreateMetaGlobalInput) SelectParameters(params ...string) *ActionHelpBoxCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionHelpBoxCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionHelpBoxCreateInput is a type for action input parameters
type ActionHelpBoxCreateInput struct {
	Page string `json:"page"`
	Action string `json:"action"`
	Language int64 `json:"language"`
	Content string `json:"content"`
	Order int64 `json:"order"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetPage sets parameter Page to value and selects it for sending
func (in *ActionHelpBoxCreateInput) SetPage(value string) *ActionHelpBoxCreateInput {
	in.Page = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Page"] = nil
	return in
}
// SetAction sets parameter Action to value and selects it for sending
func (in *ActionHelpBoxCreateInput) SetAction(value string) *ActionHelpBoxCreateInput {
	in.Action = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Action"] = nil
	return in
}
// SetLanguage sets parameter Language to value and selects it for sending
func (in *ActionHelpBoxCreateInput) SetLanguage(value int64) *ActionHelpBoxCreateInput {
	in.Language = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Language"] = nil
	return in
}
// SetContent sets parameter Content to value and selects it for sending
func (in *ActionHelpBoxCreateInput) SetContent(value string) *ActionHelpBoxCreateInput {
	in.Content = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Content"] = nil
	return in
}
// SetOrder sets parameter Order to value and selects it for sending
func (in *ActionHelpBoxCreateInput) SetOrder(value int64) *ActionHelpBoxCreateInput {
	in.Order = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Order"] = nil
	return in
}

// SelectParameters sets parameters from ActionHelpBoxCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionHelpBoxCreateInput) SelectParameters(params ...string) *ActionHelpBoxCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionHelpBoxCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionHelpBoxCreateRequest is a type for the entire action request
type ActionHelpBoxCreateRequest struct {
	HelpBox map[string]interface{} `json:"help_box"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionHelpBoxCreateOutput is a type for action output parameters
type ActionHelpBoxCreateOutput struct {
	Id int64 `json:"id"`
	Page string `json:"page"`
	Action string `json:"action"`
	Language *ActionLanguageShowOutput `json:"language"`
	Content string `json:"content"`
	Order int64 `json:"order"`
}


// Type for action response, including envelope
type ActionHelpBoxCreateResponse struct {
	Action *ActionHelpBoxCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		HelpBox *ActionHelpBoxCreateOutput `json:"help_box"`
	}

	// Action output without the namespace
	Output *ActionHelpBoxCreateOutput
}


// Prepare the action for invocation
func (action *ActionHelpBoxCreate) Prepare() *ActionHelpBoxCreateInvocation {
	return &ActionHelpBoxCreateInvocation{
		Action: action,
		Path: "/v6.0/help_boxes",
	}
}

// ActionHelpBoxCreateInvocation is used to configure action for invocation
type ActionHelpBoxCreateInvocation struct {
	// Pointer to the action
	Action *ActionHelpBoxCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionHelpBoxCreateInput
	// Global meta input parameters
	MetaInput *ActionHelpBoxCreateMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionHelpBoxCreateInvocation) NewInput() *ActionHelpBoxCreateInput {
	inv.Input = &ActionHelpBoxCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionHelpBoxCreateInvocation) SetInput(input *ActionHelpBoxCreateInput) *ActionHelpBoxCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionHelpBoxCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionHelpBoxCreateInvocation) NewMetaInput() *ActionHelpBoxCreateMetaGlobalInput {
	inv.MetaInput = &ActionHelpBoxCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionHelpBoxCreateInvocation) SetMetaInput(input *ActionHelpBoxCreateMetaGlobalInput) *ActionHelpBoxCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionHelpBoxCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionHelpBoxCreateInvocation) Call() (*ActionHelpBoxCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionHelpBoxCreateInvocation) callAsBody() (*ActionHelpBoxCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionHelpBoxCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.HelpBox
	}
	return resp, err
}




func (inv *ActionHelpBoxCreateInvocation) makeAllInputParams() *ActionHelpBoxCreateRequest {
	return &ActionHelpBoxCreateRequest{
		HelpBox: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionHelpBoxCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Page") {
			ret["page"] = inv.Input.Page
		}
		if inv.IsParameterSelected("Action") {
			ret["action"] = inv.Input.Action
		}
		if inv.IsParameterSelected("Language") {
			ret["language"] = inv.Input.Language
		}
		if inv.IsParameterSelected("Content") {
			ret["content"] = inv.Input.Content
		}
		if inv.IsParameterSelected("Order") {
			ret["order"] = inv.Input.Order
		}
	}

	return ret
}

func (inv *ActionHelpBoxCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
