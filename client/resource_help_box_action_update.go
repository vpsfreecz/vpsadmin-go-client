package client

import (
	"strings"
)

// ActionHelpBoxUpdate is a type for action Help_box#Update
type ActionHelpBoxUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionHelpBoxUpdate(client *Client) *ActionHelpBoxUpdate {
	return &ActionHelpBoxUpdate{
		Client: client,
	}
}

// ActionHelpBoxUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionHelpBoxUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionHelpBoxUpdateMetaGlobalInput) SetIncludes(value string) *ActionHelpBoxUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionHelpBoxUpdateMetaGlobalInput) SetNo(value bool) *ActionHelpBoxUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionHelpBoxUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionHelpBoxUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionHelpBoxUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionHelpBoxUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionHelpBoxUpdateInput is a type for action input parameters
type ActionHelpBoxUpdateInput struct {
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
func (in *ActionHelpBoxUpdateInput) SetAction(value string) *ActionHelpBoxUpdateInput {
	in.Action = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Action"] = nil
	return in
}

// SetContent sets parameter Content to value and selects it for sending
func (in *ActionHelpBoxUpdateInput) SetContent(value string) *ActionHelpBoxUpdateInput {
	in.Content = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Content"] = nil
	return in
}

// SetLanguage sets parameter Language to value and selects it for sending
func (in *ActionHelpBoxUpdateInput) SetLanguage(value int64) *ActionHelpBoxUpdateInput {
	in.Language = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLanguageNil(false)
	in._selectedParameters["Language"] = nil
	return in
}

// SetLanguageNil sets parameter Language to nil and selects it for sending
func (in *ActionHelpBoxUpdateInput) SetLanguageNil(set bool) *ActionHelpBoxUpdateInput {
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
func (in *ActionHelpBoxUpdateInput) SetOrder(value int64) *ActionHelpBoxUpdateInput {
	in.Order = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Order"] = nil
	return in
}

// SetPage sets parameter Page to value and selects it for sending
func (in *ActionHelpBoxUpdateInput) SetPage(value string) *ActionHelpBoxUpdateInput {
	in.Page = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Page"] = nil
	return in
}

// SelectParameters sets parameters from ActionHelpBoxUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionHelpBoxUpdateInput) SelectParameters(params ...string) *ActionHelpBoxUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionHelpBoxUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionHelpBoxUpdateInput) UnselectParameters(params ...string) *ActionHelpBoxUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionHelpBoxUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionHelpBoxUpdateRequest is a type for the entire action request
type ActionHelpBoxUpdateRequest struct {
	HelpBox map[string]interface{} `json:"help_box"`
	Meta    map[string]interface{} `json:"_meta"`
}

// ActionHelpBoxUpdateOutput is a type for action output parameters
type ActionHelpBoxUpdateOutput struct {
	Action   string                    `json:"action"`
	Content  string                    `json:"content"`
	Id       int64                     `json:"id"`
	Language *ActionLanguageShowOutput `json:"language"`
	Order    int64                     `json:"order"`
	Page     string                    `json:"page"`
}

// Type for action response, including envelope
type ActionHelpBoxUpdateResponse struct {
	Action *ActionHelpBoxUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		HelpBox *ActionHelpBoxUpdateOutput `json:"help_box"`
	}

	// Action output without the namespace
	Output *ActionHelpBoxUpdateOutput
}

// Prepare the action for invocation
func (action *ActionHelpBoxUpdate) Prepare() *ActionHelpBoxUpdateInvocation {
	return &ActionHelpBoxUpdateInvocation{
		Action: action,
		Path:   "/v6.0/help_boxes/{help_box_id}",
	}
}

// ActionHelpBoxUpdateInvocation is used to configure action for invocation
type ActionHelpBoxUpdateInvocation struct {
	// Pointer to the action
	Action *ActionHelpBoxUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionHelpBoxUpdateInput
	// Global meta input parameters
	MetaInput *ActionHelpBoxUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionHelpBoxUpdateInvocation) SetPathParamInt(param string, value int64) *ActionHelpBoxUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionHelpBoxUpdateInvocation) SetPathParamString(param string, value string) *ActionHelpBoxUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionHelpBoxUpdateInvocation) NewInput() *ActionHelpBoxUpdateInput {
	inv.Input = &ActionHelpBoxUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionHelpBoxUpdateInvocation) SetInput(input *ActionHelpBoxUpdateInput) *ActionHelpBoxUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionHelpBoxUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionHelpBoxUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionHelpBoxUpdateInvocation) NewMetaInput() *ActionHelpBoxUpdateMetaGlobalInput {
	inv.MetaInput = &ActionHelpBoxUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionHelpBoxUpdateInvocation) SetMetaInput(input *ActionHelpBoxUpdateMetaGlobalInput) *ActionHelpBoxUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionHelpBoxUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionHelpBoxUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionHelpBoxUpdateInvocation) Call() (*ActionHelpBoxUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionHelpBoxUpdateInvocation) callAsBody() (*ActionHelpBoxUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionHelpBoxUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.HelpBox
	}
	return resp, err
}

func (inv *ActionHelpBoxUpdateInvocation) makeAllInputParams() *ActionHelpBoxUpdateRequest {
	return &ActionHelpBoxUpdateRequest{
		HelpBox: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionHelpBoxUpdateInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionHelpBoxUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
