package client

import (
)

// ActionSessionTokenCreate is a type for action Session_token#Create
type ActionSessionTokenCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionSessionTokenCreate(client *Client) *ActionSessionTokenCreate {
	return &ActionSessionTokenCreate{
		Client: client,
	}
}

// ActionSessionTokenCreateMetaGlobalInput is a type for action global meta input parameters
type ActionSessionTokenCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSessionTokenCreateMetaGlobalInput) SetNo(value bool) *ActionSessionTokenCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSessionTokenCreateMetaGlobalInput) SetIncludes(value string) *ActionSessionTokenCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionSessionTokenCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSessionTokenCreateMetaGlobalInput) SelectParameters(params ...string) *ActionSessionTokenCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSessionTokenCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSessionTokenCreateInput is a type for action input parameters
type ActionSessionTokenCreateInput struct {
	User int64 `json:"user"`
	Label string `json:"label"`
	Lifetime string `json:"lifetime"`
	Interval int64 `json:"interval"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionSessionTokenCreateInput) SetUser(value int64) *ActionSessionTokenCreateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}
// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionSessionTokenCreateInput) SetLabel(value string) *ActionSessionTokenCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}
// SetLifetime sets parameter Lifetime to value and selects it for sending
func (in *ActionSessionTokenCreateInput) SetLifetime(value string) *ActionSessionTokenCreateInput {
	in.Lifetime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Lifetime"] = nil
	return in
}
// SetInterval sets parameter Interval to value and selects it for sending
func (in *ActionSessionTokenCreateInput) SetInterval(value int64) *ActionSessionTokenCreateInput {
	in.Interval = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Interval"] = nil
	return in
}

// SelectParameters sets parameters from ActionSessionTokenCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSessionTokenCreateInput) SelectParameters(params ...string) *ActionSessionTokenCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSessionTokenCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSessionTokenCreateRequest is a type for the entire action request
type ActionSessionTokenCreateRequest struct {
	SessionToken map[string]interface{} `json:"session_token"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionSessionTokenCreateOutput is a type for action output parameters
type ActionSessionTokenCreateOutput struct {
	Id int64 `json:"id"`
	User *ActionUserShowOutput `json:"user"`
	Token string `json:"token"`
	ValidTo string `json:"valid_to"`
	Label string `json:"label"`
	Lifetime string `json:"lifetime"`
	Interval int64 `json:"interval"`
	UseCount int64 `json:"use_count"`
	CreatedAt string `json:"created_at"`
}


// Type for action response, including envelope
type ActionSessionTokenCreateResponse struct {
	Action *ActionSessionTokenCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SessionToken *ActionSessionTokenCreateOutput `json:"session_token"`
	}

	// Action output without the namespace
	Output *ActionSessionTokenCreateOutput
}


// Prepare the action for invocation
func (action *ActionSessionTokenCreate) Prepare() *ActionSessionTokenCreateInvocation {
	return &ActionSessionTokenCreateInvocation{
		Action: action,
		Path: "/v5.0/session_tokens",
	}
}

// ActionSessionTokenCreateInvocation is used to configure action for invocation
type ActionSessionTokenCreateInvocation struct {
	// Pointer to the action
	Action *ActionSessionTokenCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSessionTokenCreateInput
	// Global meta input parameters
	MetaInput *ActionSessionTokenCreateMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSessionTokenCreateInvocation) NewInput() *ActionSessionTokenCreateInput {
	inv.Input = &ActionSessionTokenCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSessionTokenCreateInvocation) SetInput(input *ActionSessionTokenCreateInput) *ActionSessionTokenCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSessionTokenCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSessionTokenCreateInvocation) NewMetaInput() *ActionSessionTokenCreateMetaGlobalInput {
	inv.MetaInput = &ActionSessionTokenCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSessionTokenCreateInvocation) SetMetaInput(input *ActionSessionTokenCreateMetaGlobalInput) *ActionSessionTokenCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSessionTokenCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSessionTokenCreateInvocation) Call() (*ActionSessionTokenCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionSessionTokenCreateInvocation) callAsBody() (*ActionSessionTokenCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSessionTokenCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SessionToken
	}
	return resp, err
}




func (inv *ActionSessionTokenCreateInvocation) makeAllInputParams() *ActionSessionTokenCreateRequest {
	return &ActionSessionTokenCreateRequest{
		SessionToken: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionSessionTokenCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("User") {
			ret["user"] = inv.Input.User
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("Lifetime") {
			ret["lifetime"] = inv.Input.Lifetime
		}
		if inv.IsParameterSelected("Interval") {
			ret["interval"] = inv.Input.Interval
		}
	}

	return ret
}

func (inv *ActionSessionTokenCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
