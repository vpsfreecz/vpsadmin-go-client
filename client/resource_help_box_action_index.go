package client

import ()

// ActionHelpBoxIndex is a type for action Help_box#Index
type ActionHelpBoxIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionHelpBoxIndex(client *Client) *ActionHelpBoxIndex {
	return &ActionHelpBoxIndex{
		Client: client,
	}
}

// ActionHelpBoxIndexMetaGlobalInput is a type for action global meta input parameters
type ActionHelpBoxIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionHelpBoxIndexMetaGlobalInput) SetCount(value bool) *ActionHelpBoxIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionHelpBoxIndexMetaGlobalInput) SetIncludes(value string) *ActionHelpBoxIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionHelpBoxIndexMetaGlobalInput) SetNo(value bool) *ActionHelpBoxIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionHelpBoxIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionHelpBoxIndexMetaGlobalInput) SelectParameters(params ...string) *ActionHelpBoxIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionHelpBoxIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionHelpBoxIndexInput is a type for action input parameters
type ActionHelpBoxIndexInput struct {
	Action   string `json:"action"`
	FromId   int64  `json:"from_id"`
	Language int64  `json:"language"`
	Limit    int64  `json:"limit"`
	Page     string `json:"page"`
	View     bool   `json:"view"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAction sets parameter Action to value and selects it for sending
func (in *ActionHelpBoxIndexInput) SetAction(value string) *ActionHelpBoxIndexInput {
	in.Action = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Action"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionHelpBoxIndexInput) SetFromId(value int64) *ActionHelpBoxIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLanguage sets parameter Language to value and selects it for sending
func (in *ActionHelpBoxIndexInput) SetLanguage(value int64) *ActionHelpBoxIndexInput {
	in.Language = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLanguageNil(false)
	in._selectedParameters["Language"] = nil
	return in
}

// SetLanguageNil sets parameter Language to nil and selects it for sending
func (in *ActionHelpBoxIndexInput) SetLanguageNil(set bool) *ActionHelpBoxIndexInput {
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

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionHelpBoxIndexInput) SetLimit(value int64) *ActionHelpBoxIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetPage sets parameter Page to value and selects it for sending
func (in *ActionHelpBoxIndexInput) SetPage(value string) *ActionHelpBoxIndexInput {
	in.Page = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Page"] = nil
	return in
}

// SetView sets parameter View to value and selects it for sending
func (in *ActionHelpBoxIndexInput) SetView(value bool) *ActionHelpBoxIndexInput {
	in.View = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["View"] = nil
	return in
}

// SelectParameters sets parameters from ActionHelpBoxIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionHelpBoxIndexInput) SelectParameters(params ...string) *ActionHelpBoxIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionHelpBoxIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionHelpBoxIndexInput) UnselectParameters(params ...string) *ActionHelpBoxIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionHelpBoxIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionHelpBoxIndexOutput is a type for action output parameters
type ActionHelpBoxIndexOutput struct {
	Action   string                    `json:"action"`
	Content  string                    `json:"content"`
	Id       int64                     `json:"id"`
	Language *ActionLanguageShowOutput `json:"language"`
	Order    int64                     `json:"order"`
	Page     string                    `json:"page"`
}

// Type for action response, including envelope
type ActionHelpBoxIndexResponse struct {
	Action *ActionHelpBoxIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		HelpBoxes []*ActionHelpBoxIndexOutput `json:"help_boxes"`
	}

	// Action output without the namespace
	Output []*ActionHelpBoxIndexOutput
}

// Prepare the action for invocation
func (action *ActionHelpBoxIndex) Prepare() *ActionHelpBoxIndexInvocation {
	return &ActionHelpBoxIndexInvocation{
		Action: action,
		Path:   "/v7.0/help_boxes",
	}
}

// ActionHelpBoxIndexInvocation is used to configure action for invocation
type ActionHelpBoxIndexInvocation struct {
	// Pointer to the action
	Action *ActionHelpBoxIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionHelpBoxIndexInput
	// Global meta input parameters
	MetaInput *ActionHelpBoxIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionHelpBoxIndexInvocation) NewInput() *ActionHelpBoxIndexInput {
	inv.Input = &ActionHelpBoxIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionHelpBoxIndexInvocation) SetInput(input *ActionHelpBoxIndexInput) *ActionHelpBoxIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionHelpBoxIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionHelpBoxIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionHelpBoxIndexInvocation) NewMetaInput() *ActionHelpBoxIndexMetaGlobalInput {
	inv.MetaInput = &ActionHelpBoxIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionHelpBoxIndexInvocation) SetMetaInput(input *ActionHelpBoxIndexMetaGlobalInput) *ActionHelpBoxIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionHelpBoxIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionHelpBoxIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionHelpBoxIndexInvocation) Call() (*ActionHelpBoxIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionHelpBoxIndexInvocation) callAsQuery() (*ActionHelpBoxIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionHelpBoxIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.HelpBoxes
	}
	return resp, err
}

func (inv *ActionHelpBoxIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Action") {
			ret["help_box[action]"] = inv.Input.Action
		}
		if inv.IsParameterSelected("FromId") {
			ret["help_box[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Language") {
			ret["help_box[language]"] = convertInt64ToString(inv.Input.Language)
		}
		if inv.IsParameterSelected("Limit") {
			ret["help_box[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Page") {
			ret["help_box[page]"] = inv.Input.Page
		}
		if inv.IsParameterSelected("View") {
			ret["help_box[view]"] = convertBoolToString(inv.Input.View)
		}
	}
}

func (inv *ActionHelpBoxIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
