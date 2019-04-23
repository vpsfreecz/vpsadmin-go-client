package client

import (
)

// ActionAuthTokenCreate is a type for action Auth_token#Create
type ActionAuthTokenCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionAuthTokenCreate(client *Client) *ActionAuthTokenCreate {
	return &ActionAuthTokenCreate{
		Client: client,
	}
}

// ActionAuthTokenCreateMetaGlobalInput is a type for action global meta input parameters
type ActionAuthTokenCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionAuthTokenCreateMetaGlobalInput) SetNo(value bool) *ActionAuthTokenCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionAuthTokenCreateMetaGlobalInput) SetIncludes(value string) *ActionAuthTokenCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionAuthTokenCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionAuthTokenCreateMetaGlobalInput) SelectParameters(params ...string) *ActionAuthTokenCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionAuthTokenCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionAuthTokenCreateInput is a type for action input parameters
type ActionAuthTokenCreateInput struct {
	User int64 `json:"user"`
	Label string `json:"label"`
	Lifetime string `json:"lifetime"`
	Interval int64 `json:"interval"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionAuthTokenCreateInput) SetUser(value int64) *ActionAuthTokenCreateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}
// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionAuthTokenCreateInput) SetLabel(value string) *ActionAuthTokenCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}
// SetLifetime sets parameter Lifetime to value and selects it for sending
func (in *ActionAuthTokenCreateInput) SetLifetime(value string) *ActionAuthTokenCreateInput {
	in.Lifetime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Lifetime"] = nil
	return in
}
// SetInterval sets parameter Interval to value and selects it for sending
func (in *ActionAuthTokenCreateInput) SetInterval(value int64) *ActionAuthTokenCreateInput {
	in.Interval = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Interval"] = nil
	return in
}

// SelectParameters sets parameters from ActionAuthTokenCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionAuthTokenCreateInput) SelectParameters(params ...string) *ActionAuthTokenCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionAuthTokenCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionAuthTokenCreateRequest is a type for the entire action request
type ActionAuthTokenCreateRequest struct {
	AuthToken map[string]interface{} `json:"auth_token"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionAuthTokenCreateOutput is a type for action output parameters
type ActionAuthTokenCreateOutput struct {
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
type ActionAuthTokenCreateResponse struct {
	Action *ActionAuthTokenCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		AuthToken *ActionAuthTokenCreateOutput `json:"auth_token"`
	}

	// Action output without the namespace
	Output *ActionAuthTokenCreateOutput
}


// Prepare the action for invocation
func (action *ActionAuthTokenCreate) Prepare() *ActionAuthTokenCreateInvocation {
	return &ActionAuthTokenCreateInvocation{
		Action: action,
		Path: "/v5.0/auth_tokens",
	}
}

// ActionAuthTokenCreateInvocation is used to configure action for invocation
type ActionAuthTokenCreateInvocation struct {
	// Pointer to the action
	Action *ActionAuthTokenCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionAuthTokenCreateInput
	// Global meta input parameters
	MetaInput *ActionAuthTokenCreateMetaGlobalInput
}


// SetInput provides input parameters to send to the API
func (inv *ActionAuthTokenCreateInvocation) SetInput(input *ActionAuthTokenCreateInput) *ActionAuthTokenCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionAuthTokenCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionAuthTokenCreateInvocation) SetMetaInput(input *ActionAuthTokenCreateMetaGlobalInput) *ActionAuthTokenCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionAuthTokenCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionAuthTokenCreateInvocation) Call() (*ActionAuthTokenCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionAuthTokenCreateInvocation) callAsBody() (*ActionAuthTokenCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionAuthTokenCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.AuthToken
	}
	return resp, err
}




func (inv *ActionAuthTokenCreateInvocation) makeAllInputParams() *ActionAuthTokenCreateRequest {
	return &ActionAuthTokenCreateRequest{
		AuthToken: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionAuthTokenCreateInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionAuthTokenCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
