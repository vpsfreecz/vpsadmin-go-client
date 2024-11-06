package client

import ()

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
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
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

// SetNo sets parameter No to value and selects it for sending
func (in *ActionHelpBoxCreateMetaGlobalInput) SetNo(value bool) *ActionHelpBoxCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
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
	Action   string `json:"action"`
	Content  string `json:"content"`
	Language int64  `json:"language"`
	Order    int64  `json:"order"`
	Page     string `json:"page"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
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

// SetContent sets parameter Content to value and selects it for sending
func (in *ActionHelpBoxCreateInput) SetContent(value string) *ActionHelpBoxCreateInput {
	in.Content = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Content"] = nil
	return in
}

// SetLanguage sets parameter Language to value and selects it for sending
func (in *ActionHelpBoxCreateInput) SetLanguage(value int64) *ActionHelpBoxCreateInput {
	in.Language = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLanguageNil(false)
	in._selectedParameters["Language"] = nil
	return in
}

// SetLanguageNil sets parameter Language to nil and selects it for sending
func (in *ActionHelpBoxCreateInput) SetLanguageNil(set bool) *ActionHelpBoxCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Language"] = nil
		in.SelectParameters("Language")
	} else {
		delete(in._nilParameters, "Language")
	}
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

// SetPage sets parameter Page to value and selects it for sending
func (in *ActionHelpBoxCreateInput) SetPage(value string) *ActionHelpBoxCreateInput {
	in.Page = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Page"] = nil
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

// UnselectParameters unsets parameters from ActionHelpBoxCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionHelpBoxCreateInput) UnselectParameters(params ...string) *ActionHelpBoxCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
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
	Meta    map[string]interface{} `json:"_meta"`
}

// ActionHelpBoxCreateOutput is a type for action output parameters
type ActionHelpBoxCreateOutput struct {
	Action   string                    `json:"action"`
	Content  string                    `json:"content"`
	Id       int64                     `json:"id"`
	Language *ActionLanguageShowOutput `json:"language"`
	Order    int64                     `json:"order"`
	Page     string                    `json:"page"`
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
		Path:   "/v7.0/help_boxes",
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

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionHelpBoxCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
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

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionHelpBoxCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
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
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionHelpBoxCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Action") {
			ret["action"] = inv.Input.Action
		}
		if inv.IsParameterSelected("Content") {
			ret["content"] = inv.Input.Content
		}
		if inv.IsParameterSelected("Language") {
			if inv.IsParameterNil("Language") {
				ret["language"] = nil
			} else {
				ret["language"] = inv.Input.Language
			}
		}
		if inv.IsParameterSelected("Order") {
			ret["order"] = inv.Input.Order
		}
		if inv.IsParameterSelected("Page") {
			ret["page"] = inv.Input.Page
		}
	}

	return ret
}

func (inv *ActionHelpBoxCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
